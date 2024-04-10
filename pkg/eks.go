package pkg

import (
	"encoding/json"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type EksKubeConfig struct {
}

type EksKubeConfigArgs struct {
	ClusterName     string `pulumi:"clusterName"`
	ClusterEndpoint string `pulumi:"clusterEndpoint"`
	CertificateData string `pulumi:"certificateData,optional"`
	RoleArn         string `pulumi:"roleArn,optional"`
	ProfileName     string `pulumi:"profileName,optional"`
}

func (eks *EksKubeConfigArgs) Annotate(a infer.Annotator) {
	a.Describe(&eks.ClusterName, "Name of the EKS cluster you want to generate the kubeconfig for")
	a.Describe(&eks.ClusterEndpoint, "Endpoint for your Kubernetes API server.")
	a.Describe(&eks.CertificateData, "Base64 encoded certificate data required to communicate with your cluster.")
	a.Describe(&eks.RoleArn, "Role arn that you want the kubeconfig to use. Optional")
	a.Describe(&eks.ProfileName, "AWS Profile name that you want the kubeconfig to use")
}

type EksKubeConfigState struct {
	ClusterName     string `pulumi:"clusterName"`
	ClusterEndpoint string `pulumi:"clusterEndpoint"`
	CertificateData string `pulumi:"certificateData"`
	RoleArn         string `pulumi:"roleArn,optional"`
	ProfileName     string `pulumi:"profileName,optional"`
	KubeConfig      string `pulumi:"kubeconfig" provider:"secret"`
}

func (eks *EksKubeConfigState) Annotate(a infer.Annotator) {
	a.Describe(&eks.KubeConfig, "Generated Kubeconfig for working with your EKS cluster")
	a.Describe(&eks.ClusterName, "Name of the EKS cluster you want to generate the kubeconfig for")
	a.Describe(&eks.ClusterEndpoint, "Endpoint for your Kubernetes API server.")
	a.Describe(&eks.CertificateData, "Base64 encoded certificate data required to communicate with your cluster.")
	a.Describe(&eks.RoleArn, "Role arn that you want the kubeconfig to use. Optional")
	a.Describe(&eks.ProfileName, "AWS Profile name that you want the kubeconfig to use. Optional")
}

func (e *EksKubeConfig) Create(ctx p.Context, name string, input EksKubeConfigArgs, preview bool) (
	id string, output EksKubeConfigState, err error) {
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
		ProfileName:     input.ProfileName,
	}, nil
}

func (*EksKubeConfig) Diff(ctx p.Context, id string, olds EksKubeConfigState, news EksKubeConfigArgs) (p.DiffResponse, error) {
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
	if news.ProfileName != olds.ProfileName {
		diff["profileName"] = p.PropertyDiff{Kind: p.Update}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (*EksKubeConfig) Update(ctx p.Context, id string, olds EksKubeConfigState, news EksKubeConfigArgs, preview bool) (EksKubeConfigState, error) {

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
		ProfileName:     news.ProfileName,
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

	env := []Env{
		{
			Name:  "KUBERNETES_EXEC_INFO",
			Value: "{\"apiVersion\": \"client.authentication.k8s.io/v1beta1\"}",
		},
	}

	if len(input.ProfileName) > 0 {
		env = append(env, Env{
			Name:  "AWS_PROFILE",
			Value: input.ProfileName,
		},
		)
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
		return "", nil
	}

	return string(kubeconfig), nil
}
