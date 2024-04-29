package pkg

import (
	"encoding/base64"

	"github.com/pulumi/pulumi-azure-native-sdk/containerservice/v2"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AksKubeConfig struct{}

type AksKubeConfigArgs struct {
	ResourceGroupName pulumi.String `pulumi:"resourceGroupName"`
	ClusterName       pulumi.String `pulumi:"clusterName"`
	IsAdmin           pulumi.Bool   `pulumi:"isAdmin,optional"`
}

func (kca *AksKubeConfigArgs) Annotate(a infer.Annotator) {
	a.Describe(&kca.ResourceGroupName, "Name of the resource group that the cluster is part of.")
	a.Describe(&kca.ClusterName, "Name of the AKS cluster you want the Kubeconfig from.")
	a.Describe(&kca.IsAdmin, "Specify whether you want to retrieve the admin kubeconfig or the user kubeconfig. Default value is false")
}

type AksKubeConfigState struct {
	pulumi.ResourceState
	KubeConfig pulumi.StringOutput `pulumi:"kubeconfig"`
}

func (kca *AksKubeConfigState) Annotate(a infer.Annotator) {
	a.Describe(&kca.KubeConfig, "Kubeconfig returned from AKS cluster")
}

func (k *AksKubeConfig) Construct(ctx *pulumi.Context, name, typ string, args AksKubeConfigArgs, opts pulumi.ResourceOption) (
	*AksKubeConfigState, error) {
	comp := &AksKubeConfigState{}
	err := ctx.RegisterComponentResource(typ, name, comp, opts)

	if err != nil {
		return nil, err
	}

	var configs containerservice.CredentialResultResponseArrayOutput

	if args.IsAdmin {
		configs = containerservice.ListManagedClusterAdminCredentialsOutput(ctx,
			containerservice.ListManagedClusterAdminCredentialsOutputArgs{
				ResourceGroupName: args.ResourceGroupName,
				ResourceName:      args.ClusterName,
			}).Kubeconfigs()
	} else {
		configs = containerservice.ListManagedClusterUserCredentialsOutput(ctx,
			containerservice.ListManagedClusterUserCredentialsOutputArgs{
				ResourceGroupName: args.ResourceGroupName,
				ResourceName:      args.ClusterName,
			}).Kubeconfigs()
	}

	kubeconfig := configs.Index(pulumi.Int(0)).Value().
		ApplyT(func(encoded string) string {
			kubeconfig, err := base64.StdEncoding.DecodeString(encoded)
			if err != nil {
				return ""
			}
			return string(kubeconfig[:])
		}).(pulumi.StringOutput)

	comp.KubeConfig = pulumi.ToSecret(kubeconfig).(pulumi.StringOutput)
	return comp, nil
}
