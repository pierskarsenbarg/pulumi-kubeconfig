package main

import (
	"fmt"
	"os"

	"github.com/pierskarsenbarg/pulumi-kubeconfig/pkg"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	nodejsgen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	pythongen "github.com/pulumi/pulumi/pkg/v3/codegen/python"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

var Version string

func main() {
	err := p.RunProvider("kubeconfig", Version, provider())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}

func provider() p.Provider {
	return infer.Provider(infer.Options{
		Metadata: schema.Metadata{

			DisplayName: "kubeconfig",
			Description: "Kubeconfig provider",
			Keywords: []string{"pulumi", "kubeconfig", "category/infrastructure", "kind/native"},
			License: "Apache-2.0",
			Homepage: "https://github.com/pierskarenbarg/pulumi-kubeconfig",
			Repository: "https://github.com/pierskarenbarg/pulumi-kubeconfig",
			LanguageMap: map[string]any{
				"go": gen.GoPackageInfo{
					ImportBasePath: "github.com/pierskarsenbarg/pulumi-kubeconfig/sdk/go/kubeconfig",
				},
				"nodejs": nodejsgen.NodePackageInfo{
					PackageName: "@pierskarsenbarg/kubeconfig",
					Dependencies: map[string]string{
						"@pulumi/pulumi":       "^3.0.0",
						"@pulumi/kubernetes":   "^4.0.0",
						"@pulumi/azure-native": "^2.0.0",
						"@pulumi/aws":          "^6.0.0",
					},
					DevDependencies: map[string]string{
						"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
						"@types/mime": "^2.0.0",
					},
				},
				"csharp": dotnetgen.CSharpPackageInfo{
					RootNamespace: "PiersKarsenbarg",
					PackageReferences: map[string]string{
						"Pulumi":             "3.*",
						"Pulumi.Kubernetes":  "4.*",
						"Pulumi.AzureNative": "2.*",
						"Pulumi.Aws":         "6.*",
					},
				},
				"python": pythongen.PackageInfo{
					Requires: map[string]string{
						"pulumi":              ">=3.0.0,<4.0.0",
						"pulumi_kubernetes":   ">=4.0.0,<5.0.0",
						"pulumi_azure_native": ">=2.0.0,<3.0.0",
						"pulumni_aws":         ">=6.0.0,<7.0.0",
					},
					PackageName: "pierskarsenbarg_pulumi_kubeconfig",
				},
			},
			PluginDownloadURL: "github://api.github.com/pierskarsenbarg/pulumi-kubeconfig",
			Publisher:         "Piers Karsenbarg",
		},
		Resources: []infer.InferredResource{
			infer.Resource[*pkg.EksKubeConfig, pkg.EksKubeConfigArgs, pkg.EksKubeConfigState](),
			infer.Resource[*pkg.GkeKubeConfig, pkg.GkeKubeConfigArgs, pkg.GkeKubeConfigState](),
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"pkg": "index",
		},
		Functions: []infer.InferredFunction{},
		Components: []infer.InferredComponent{
			infer.Component[*pkg.AksKubeConfig, pkg.AksKubeConfigArgs, *pkg.AksKubeConfigState](),
		},
	})
}
