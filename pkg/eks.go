package pkg

import (
	"context"
	"encoding/json"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type EksKubeConfig struct{}

type EksKubeConfigArgs struct {
	ClusterName     string `pulumi:"clusterName"`
	ClusterEndpoint string `pulumi:"clusterEndpoint"`
	CertificateData string `pulumi:"certificateData,optional"`
	RoleArn         string `pulumi:"roleArn,optional"`
	Region          string `pulumi:"region,optional"`
}

func (eks *EksKubeConfigArgs) Annotate(a infer.Annotator) {
	a.Describe(&eks.ClusterName, "Name of the EKS cluster you want to generate the kubeconfig for")
	a.Describe(&eks.ClusterEndpoint, "Endpoint for your Kubernetes API server.")
	a.Describe(&eks.CertificateData, "Base64 encoded certificate data required to communicate with your cluster.")
	a.Describe(&eks.RoleArn, "Role arn that you want the kubeconfig to use. Optional")
	a.Describe(&eks.Region, "Region that the EKS cluster is in. Optional")
}

type EksKubeConfigState struct {
	ClusterName     string `pulumi:"clusterName"`
	ClusterEndpoint string `pulumi:"clusterEndpoint"`
	CertificateData string `pulumi:"certificateData"`
	RoleArn         string `pulumi:"roleArn,optional"`
	Region          string `pulumi:"region,optional"`
	KubeConfig      string `pulumi:"kubeconfig" provider:"secret"`
}

func (eks *EksKubeConfigState) Annotate(a infer.Annotator) {
	a.Describe(&eks.KubeConfig, "Generated Kubeconfig for working with your EKS cluster")
	a.Describe(&eks.ClusterName, "Name of the EKS cluster you want to generate the kubeconfig for")
	a.Describe(&eks.ClusterEndpoint, "Endpoint for your Kubernetes API server.")
	a.Describe(&eks.CertificateData, "Base64 encoded certificate data required to communicate with your cluster.")
	a.Describe(&eks.RoleArn, "Role arn that you want the kubeconfig to use. Optional")
}

func (e *EksKubeConfig) Create(ctx context.Context, name string, input EksKubeConfigArgs, preview bool) (
	id string, output EksKubeConfigState, err error,
) {
	if preview {
		return "", EksKubeConfigState{}, nil
	}

	kubeConfig, err := buildEksConfig(input)
	if err != nil {
		return name, EksKubeConfigState{
			KubeConfig: kubeConfig,
		}, nil
	}

	return name, EksKubeConfigState{
		KubeConfig:      kubeConfig,
		ClusterName:     input.ClusterName,
		ClusterEndpoint: input.ClusterEndpoint,
		CertificateData: input.CertificateData,
		RoleArn:         input.RoleArn,
		Region: 		 input.Region,
	}, nil
}

func (*EksKubeConfig) Diff(ctx context.Context, id string, olds EksKubeConfigState, news EksKubeConfigArgs) (p.DiffResponse, error) {
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
	if news.RoleArn != olds.RoleArn {
		diff["roleArn"] = p.PropertyDiff{Kind: p.Update}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (*EksKubeConfig) Update(ctx context.Context, id string, olds EksKubeConfigState, news EksKubeConfigArgs, preview bool) (EksKubeConfigState, error) {
	kubeConfig, err := buildEksConfig(news)
	if err != nil {
		return EksKubeConfigState{}, err
	}

	return EksKubeConfigState{
		KubeConfig:      kubeConfig,
		ClusterName:     news.ClusterName,
		ClusterEndpoint: news.ClusterEndpoint,
		CertificateData: news.CertificateData,
		RoleArn:         news.RoleArn,
	}, nil
}

func buildEksConfig(input EksKubeConfigArgs) (string, error) {
	cmdArgs := []string{
		"eks",
		"get-token",
		"--cluster-name",
		input.ClusterName,
	}

	if len(input.RoleArn) > 0 {
		cmdArgs = append(cmdArgs, "--role", input.RoleArn)
	}

	if len(input.Region) > 0 {
		cmdArgs = append(cmdArgs, "--region", input.Region)
	}

	env := []Env{
		{
			Name:  "KUBERNETES_EXEC_INFO",
			Value: "{\"apiVersion\": \"client.authentication.k8s.io/v1beta1\"}",
		},
	}

	kubeconfig, err := json.Marshal(&KubeConfig{
		APIVersion: "v1",
		Clusters: []Clusters{
			{
				Cluster: Cluster{
					Server:                   input.ClusterEndpoint,
					CertificateAuthorityData: input.CertificateData,
				},
				Name: "kubernetes",
			},
		},
		Contexts: []Contexts{
			{
				Context: Context{
					Cluster: "kubernetes",
					User:    "aws",
				},
				Name: "aws",
			},
		},
		Kind:           "Config",
		CurrentContext: "aws",
		Users: []Users{
			{
				Name: "aws",
				User: User{
					Exec: Exec{
						APIVersion: "client.authentication.k8s.io/v1beta1",
						Command:    "aws",
						Args:       cmdArgs,
						Env:        env,
					},
				},
			},
		},
	})
	if err != nil {
		return "", err
	}

	return string(kubeconfig), nil
}
