# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['EksKubeConfigArgs', 'EksKubeConfig']

@pulumi.input_type
class EksKubeConfigArgs:
    def __init__(__self__, *,
                 cluster_endpoint: pulumi.Input[str],
                 cluster_name: pulumi.Input[str],
                 certificate_data: Optional[pulumi.Input[str]] = None,
                 profile_name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EksKubeConfig resource.
        :param pulumi.Input[str] cluster_endpoint: Endpoint for your Kubernetes API server.
        :param pulumi.Input[str] cluster_name: Name of the EKS cluster you want to generate the kubeconfig for
        :param pulumi.Input[str] certificate_data: Base64 encoded certificate data required to communicate with your cluster.
        :param pulumi.Input[str] profile_name: AWS Profile name that you want the kubeconfig to use
        :param pulumi.Input[str] role_arn: Role arn that you want the kubeconfig to use. Optional
        """
        pulumi.set(__self__, "cluster_endpoint", cluster_endpoint)
        pulumi.set(__self__, "cluster_name", cluster_name)
        if certificate_data is not None:
            pulumi.set(__self__, "certificate_data", certificate_data)
        if profile_name is not None:
            pulumi.set(__self__, "profile_name", profile_name)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)

    @property
    @pulumi.getter(name="clusterEndpoint")
    def cluster_endpoint(self) -> pulumi.Input[str]:
        """
        Endpoint for your Kubernetes API server.
        """
        return pulumi.get(self, "cluster_endpoint")

    @cluster_endpoint.setter
    def cluster_endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_endpoint", value)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Input[str]:
        """
        Name of the EKS cluster you want to generate the kubeconfig for
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="certificateData")
    def certificate_data(self) -> Optional[pulumi.Input[str]]:
        """
        Base64 encoded certificate data required to communicate with your cluster.
        """
        return pulumi.get(self, "certificate_data")

    @certificate_data.setter
    def certificate_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_data", value)

    @property
    @pulumi.getter(name="profileName")
    def profile_name(self) -> Optional[pulumi.Input[str]]:
        """
        AWS Profile name that you want the kubeconfig to use
        """
        return pulumi.get(self, "profile_name")

    @profile_name.setter
    def profile_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "profile_name", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Role arn that you want the kubeconfig to use. Optional
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)


class EksKubeConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_data: Optional[pulumi.Input[str]] = None,
                 cluster_endpoint: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 profile_name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a EksKubeConfig resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_data: Base64 encoded certificate data required to communicate with your cluster.
        :param pulumi.Input[str] cluster_endpoint: Endpoint for your Kubernetes API server.
        :param pulumi.Input[str] cluster_name: Name of the EKS cluster you want to generate the kubeconfig for
        :param pulumi.Input[str] profile_name: AWS Profile name that you want the kubeconfig to use
        :param pulumi.Input[str] role_arn: Role arn that you want the kubeconfig to use. Optional
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EksKubeConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a EksKubeConfig resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param EksKubeConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EksKubeConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_data: Optional[pulumi.Input[str]] = None,
                 cluster_endpoint: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 profile_name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EksKubeConfigArgs.__new__(EksKubeConfigArgs)

            __props__.__dict__["certificate_data"] = certificate_data
            if cluster_endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_endpoint'")
            __props__.__dict__["cluster_endpoint"] = cluster_endpoint
            if cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_name'")
            __props__.__dict__["cluster_name"] = cluster_name
            __props__.__dict__["profile_name"] = profile_name
            __props__.__dict__["role_arn"] = role_arn
            __props__.__dict__["kubeconfig"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["kubeconfig"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(EksKubeConfig, __self__).__init__(
            'kubeconfig:index:EksKubeConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'EksKubeConfig':
        """
        Get an existing EksKubeConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EksKubeConfigArgs.__new__(EksKubeConfigArgs)

        __props__.__dict__["certificate_data"] = None
        __props__.__dict__["cluster_endpoint"] = None
        __props__.__dict__["cluster_name"] = None
        __props__.__dict__["kubeconfig"] = None
        __props__.__dict__["profile_name"] = None
        __props__.__dict__["role_arn"] = None
        return EksKubeConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="certificateData")
    def certificate_data(self) -> pulumi.Output[str]:
        """
        Base64 encoded certificate data required to communicate with your cluster.
        """
        return pulumi.get(self, "certificate_data")

    @property
    @pulumi.getter(name="clusterEndpoint")
    def cluster_endpoint(self) -> pulumi.Output[str]:
        """
        Endpoint for your Kubernetes API server.
        """
        return pulumi.get(self, "cluster_endpoint")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[str]:
        """
        Name of the EKS cluster you want to generate the kubeconfig for
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter
    def kubeconfig(self) -> pulumi.Output[str]:
        """
        Generated Kubeconfig for working with your EKS cluster
        """
        return pulumi.get(self, "kubeconfig")

    @property
    @pulumi.getter(name="profileName")
    def profile_name(self) -> pulumi.Output[Optional[str]]:
        """
        AWS Profile name that you want the kubeconfig to use. Optional
        """
        return pulumi.get(self, "profile_name")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        Role arn that you want the kubeconfig to use. Optional
        """
        return pulumi.get(self, "role_arn")

