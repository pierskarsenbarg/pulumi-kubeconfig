// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PiersKarsenbarg.Pulumi.Kubeconfig
{
    [KubeconfigResourceType("kubeconfig:index:EksKubeConfig")]
    public partial class EksKubeConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Base64 encoded certificate data required to communicate with your cluster.
        /// </summary>
        [Output("certificateData")]
        public Output<string> CertificateData { get; private set; } = null!;

        /// <summary>
        /// Endpoint for your Kubernetes API server.
        /// </summary>
        [Output("clusterEndpoint")]
        public Output<string> ClusterEndpoint { get; private set; } = null!;

        /// <summary>
        /// Name of the EKS cluster you want to generate the kubeconfig for
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// Generated Kubeconfig for working with your EKS cluster
        /// </summary>
        [Output("kubeconfig")]
        public Output<string> Kubeconfig { get; private set; } = null!;

        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        /// <summary>
        /// Role arn that you want the kubeconfig to use. Optional
        /// </summary>
        [Output("roleArn")]
        public Output<string?> RoleArn { get; private set; } = null!;


        /// <summary>
        /// Create a EksKubeConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EksKubeConfig(string name, EksKubeConfigArgs args, CustomResourceOptions? options = null)
            : base("kubeconfig:index:EksKubeConfig", name, args ?? new EksKubeConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EksKubeConfig(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubeconfig:index:EksKubeConfig", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pierskarsenbarg/pulumi-kubeconfig",
                AdditionalSecretOutputs =
                {
                    "kubeconfig",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EksKubeConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EksKubeConfig Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EksKubeConfig(name, id, options);
        }
    }

    public sealed class EksKubeConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Base64 encoded certificate data required to communicate with your cluster.
        /// </summary>
        [Input("certificateData")]
        public Input<string>? CertificateData { get; set; }

        /// <summary>
        /// Endpoint for your Kubernetes API server.
        /// </summary>
        [Input("clusterEndpoint", required: true)]
        public Input<string> ClusterEndpoint { get; set; } = null!;

        /// <summary>
        /// Name of the EKS cluster you want to generate the kubeconfig for
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// Region that the EKS cluster is in. Optional
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Role arn that you want the kubeconfig to use. Optional
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        public EksKubeConfigArgs()
        {
        }
        public static new EksKubeConfigArgs Empty => new EksKubeConfigArgs();
    }
}