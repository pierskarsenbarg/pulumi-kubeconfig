package pkg

import (
	"context"
	"encoding/json"
	"os"

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
	Profile         string `pulumi:"profile,optional"`
}

func (eks *EksKubeConfigArgs) Annotate(a infer.Annotator) {
	a.Describe(&eks.ClusterName, "Name of the EKS cluster you want to generate the kubeconfig for")
	a.Describe(&eks.ClusterEndpoint, "Endpoint for your Kubernetes API server.")
	a.Describe(&eks.CertificateData, "Base64 encoded certificate data required to communicate with your cluster.")
	a.Describe(&eks.RoleArn, "Role arn that you want the kubeconfig to use. Optional")
	a.Describe(&eks.Region, "Region that the EKS cluster is in. Optional")
	a.Describe(&eks.Profile, "AWS Profile name. This will overwrite any environment variables set.")
}

type EksKubeConfigState struct {
	ClusterName     string `pulumi:"clusterName"`
	ClusterEndpoint string `pulumi:"clusterEndpoint"`
	CertificateData string `pulumi:"certificateData"`
	RoleArn         string `pulumi:"roleArn,optional"`
	Region          string `pulumi:"region,optional"`
	KubeConfig      string `pulumi:"kubeconfig" provider:"secret"`
	Profile         string `pulumi:"profile,optional"`
}

func (eks *EksKubeConfigState) Annotate(a infer.Annotator) {
	a.Describe(&eks.KubeConfig, "Generated Kubeconfig for working with your EKS cluster")
	a.Describe(&eks.ClusterName, "Name of the EKS cluster you want to generate the kubeconfig for")
	a.Describe(&eks.ClusterEndpoint, "Endpoint for your Kubernetes API server.")
	a.Describe(&eks.CertificateData, "Base64 encoded certificate data required to communicate with your cluster.")
	a.Describe(&eks.RoleArn, "Role arn that you want the kubeconfig to use. Optional")
	a.Describe(&eks.Profile, "AWS Profile name. This will overwrite any environment variables set.")
}

func (e *EksKubeConfig) Create(ctx context.Context, req infer.CreateRequest[EksKubeConfigArgs]) (infer.CreateResponse[EksKubeConfigState], error) {
	if req.DryRun {
		return infer.CreateResponse[EksKubeConfigState]{}, nil
	}

	kubeConfig, err := buildEksConfig(req.Inputs)
	if err != nil {
		return infer.CreateResponse[EksKubeConfigState]{
			ID: req.Name,
			Output: EksKubeConfigState{
				KubeConfig: kubeConfig,
			},
		}, nil
	}

	return infer.CreateResponse[EksKubeConfigState]{
		ID: req.Name,
		Output: EksKubeConfigState{
			KubeConfig:      kubeConfig,
			ClusterName:     req.Inputs.ClusterName,
			ClusterEndpoint: req.Inputs.ClusterEndpoint,
			CertificateData: req.Inputs.CertificateData,
			RoleArn:         req.Inputs.RoleArn,
			Region:          req.Inputs.Region,
			Profile:         req.Inputs.Profile,
		},
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
	if news.Profile != olds.Profile {
		diff["profile"] = p.PropertyDiff{Kind: p.Update}
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
		Profile:         news.Profile,
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

	le, ok := os.LookupEnv("AWS_PROFILE")
	if len(le) > 0 && ok {
		env = append(env, Env{
			Name:  "AWS_PROFILE",
			Value: os.Getenv("AWS_PROFILE"),
		})
	} else if len(input.Profile) > 0 {
		env = append(env, Env{
			Name:  "AWS_PROFILE",
			Value: input.Profile,
		})
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
