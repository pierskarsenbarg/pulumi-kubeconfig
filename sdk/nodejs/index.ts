// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { AksKubeConfigArgs } from "./aksKubeConfig";
export type AksKubeConfig = import("./aksKubeConfig").AksKubeConfig;
export const AksKubeConfig: typeof import("./aksKubeConfig").AksKubeConfig = null as any;
utilities.lazyLoad(exports, ["AksKubeConfig"], () => require("./aksKubeConfig"));

export { EksKubeConfigArgs } from "./eksKubeConfig";
export type EksKubeConfig = import("./eksKubeConfig").EksKubeConfig;
export const EksKubeConfig: typeof import("./eksKubeConfig").EksKubeConfig = null as any;
utilities.lazyLoad(exports, ["EksKubeConfig"], () => require("./eksKubeConfig"));

export { GkeKubeConfigArgs } from "./gkeKubeConfig";
export type GkeKubeConfig = import("./gkeKubeConfig").GkeKubeConfig;
export const GkeKubeConfig: typeof import("./gkeKubeConfig").GkeKubeConfig = null as any;
utilities.lazyLoad(exports, ["GkeKubeConfig"], () => require("./gkeKubeConfig"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "kubeconfig:index:AksKubeConfig":
                return new AksKubeConfig(name, <any>undefined, { urn })
            case "kubeconfig:index:EksKubeConfig":
                return new EksKubeConfig(name, <any>undefined, { urn })
            case "kubeconfig:index:GkeKubeConfig":
                return new GkeKubeConfig(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("kubeconfig", "index", _module)
pulumi.runtime.registerResourcePackage("kubeconfig", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:kubeconfig") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});