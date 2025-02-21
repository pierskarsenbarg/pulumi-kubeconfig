// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class EksKubeConfig extends pulumi.CustomResource {
    /**
     * Get an existing EksKubeConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EksKubeConfig {
        return new EksKubeConfig(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubeconfig:index:EksKubeConfig';

    /**
     * Returns true if the given object is an instance of EksKubeConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EksKubeConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EksKubeConfig.__pulumiType;
    }

    /**
     * Base64 encoded certificate data required to communicate with your cluster.
     */
    public readonly certificateData!: pulumi.Output<string>;
    /**
     * Endpoint for your Kubernetes API server.
     */
    public readonly clusterEndpoint!: pulumi.Output<string>;
    /**
     * Name of the EKS cluster you want to generate the kubeconfig for
     */
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * Generated Kubeconfig for working with your EKS cluster
     */
    public /*out*/ readonly kubeconfig!: pulumi.Output<string>;
    /**
     * AWS Profile name. This will overwrite any environment variables set.
     */
    public readonly profile!: pulumi.Output<string | undefined>;
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * Role arn that you want the kubeconfig to use. Optional
     */
    public readonly roleArn!: pulumi.Output<string | undefined>;

    /**
     * Create a EksKubeConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EksKubeConfigArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.clusterEndpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterEndpoint'");
            }
            if ((!args || args.clusterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterName'");
            }
            resourceInputs["certificateData"] = args ? args.certificateData : undefined;
            resourceInputs["clusterEndpoint"] = args ? args.clusterEndpoint : undefined;
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["profile"] = args ? args.profile : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["kubeconfig"] = undefined /*out*/;
        } else {
            resourceInputs["certificateData"] = undefined /*out*/;
            resourceInputs["clusterEndpoint"] = undefined /*out*/;
            resourceInputs["clusterName"] = undefined /*out*/;
            resourceInputs["kubeconfig"] = undefined /*out*/;
            resourceInputs["profile"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["kubeconfig"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(EksKubeConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a EksKubeConfig resource.
 */
export interface EksKubeConfigArgs {
    /**
     * Base64 encoded certificate data required to communicate with your cluster.
     */
    certificateData?: pulumi.Input<string>;
    /**
     * Endpoint for your Kubernetes API server.
     */
    clusterEndpoint: pulumi.Input<string>;
    /**
     * Name of the EKS cluster you want to generate the kubeconfig for
     */
    clusterName: pulumi.Input<string>;
    /**
     * AWS Profile name. This will overwrite any environment variables set.
     */
    profile?: pulumi.Input<string>;
    /**
     * Region that the EKS cluster is in. Optional
     */
    region?: pulumi.Input<string>;
    /**
     * Role arn that you want the kubeconfig to use. Optional
     */
    roleArn?: pulumi.Input<string>;
}
