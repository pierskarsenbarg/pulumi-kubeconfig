package pkg

import (
	"context"
	"encoding/json"
	"fmt"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type GkeKubeConfig struct {
}

type GkeKubeConfigArgs struct {
	ClusterName     string `pulumi:"clusterName"`
	ClusterEndpoint string `pulumi:"clusterEndpoint"`
	CertificateData string `pulumi:"certificateData"`
	Project         string `pulumi:"project"`
	Zone            string `pulumi:"zone"`
}

func (gke *GkeKubeConfigArgs) Annotate(a infer.Annotator) {
	a.Describe(&gke.ClusterName, "Name of the GKE cluster you want to generate the kubeconfig for")
	a.Describe(&gke.ClusterEndpoint, "Endpoint for your Kubernetes API server.")
	a.Describe(&gke.CertificateData, "Base64 encoded certificate data required to communicate with your cluster.")
}

type GkeKubeConfigState struct {
	ClusterName     string `pulumi:"clusterName"`
	ClusterEndpoint string `pulumi:"clusterEndpoint"`
	CertificateData string `pulumi:"certificateData,optional"`
	KubeConfig      string `pulumi:"kubeconfig" provider:"secret"`
}

func (gke *GkeKubeConfigState) Annotate(a infer.Annotator) {
	a.Describe(&gke.KubeConfig, "Generated Kubeconfig for working with your GKE cluster")
}

func (g *GkeKubeConfig) Create(ctx context.Context, req infer.CreateRequest[GkeKubeConfigArgs]) (infer.CreateResponse[GkeKubeConfigState], error) {
	if req.DryRun {
		return infer.CreateResponse[GkeKubeConfigState]{}, nil
	}

	kubeConfig, err := buildGkeConfig(req.Inputs)
	if err != nil {
		return infer.CreateResponse[GkeKubeConfigState]{
			ID: req.Name,
			Output: GkeKubeConfigState{
				KubeConfig: "",
			},
		}, nil
	}

	return infer.CreateResponse[GkeKubeConfigState]{
		ID: req.Name,
		Output: GkeKubeConfigState{
			KubeConfig:      kubeConfig,
			ClusterName:     req.Inputs.ClusterName,
			ClusterEndpoint: req.Inputs.ClusterEndpoint,
			CertificateData: req.Inputs.CertificateData,
		},
	}, nil
}

func (*GkeKubeConfig) Diff(ctx context.Context, id string, olds GkeKubeConfigState, news GkeKubeConfigArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}
	if news.CertificateData != olds.CertificateData {
		diff["certificateData"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.ClusterEndpoint != olds.ClusterEndpoint {
		diff["clusterEndpoint"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.ClusterName != olds.ClusterName {
		diff["clusterName"] = p.PropertyDiff{Kind: p.Update}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func buildGkeConfig(input GkeKubeConfigArgs) (string, error) {
	kubeconfig, err := json.Marshal(&KubeConfig{
		APIVersion: "v1",
		Clusters: []Clusters{
			{
				Cluster: Cluster{
					Server:                   fmt.Sprintf("https://%s", input.ClusterEndpoint),
					CertificateAuthorityData: input.CertificateData,
				},
				Name: "kubernetes",
			},
		},
		Contexts: []Contexts{
			{
				Context: Context{
					Cluster: "kubernetes",
					User:    "gke",
				},
				Name: "gke",
			},
		},
		Kind:           "Config",
		CurrentContext: "gke",
		Users: []Users{
			{
				Name: "gke",
				User: User{
					Exec: Exec{
						APIVersion:        "client.authentication.k8s.io/v1beta1",
						Command:           "gke-gcloud-auth-plugin",
						InstallHint:       "Install gke-gcloud-auth-plugin for use with kubectl by following https://cloud.google.com/blog/products/containers-kubernetes/kubectl-auth-changes-in-gke",
						ProvideCluserInfo: true,
					},
				},
			},
		},
	})

	if err != nil {
		return "", nil
	}

	return string(kubeconfig), nil
}
