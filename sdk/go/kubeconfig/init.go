// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubeconfig

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pierskarsenbarg/pulumi-kubeconfig/sdk/go/kubeconfig/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "kubeconfig:index:AksKubeConfig":
		r = &AksKubeConfig{}
	case "kubeconfig:index:EksKubeConfig":
		r = &EksKubeConfig{}
	case "kubeconfig:index:GkeKubeConfig":
		r = &GkeKubeConfig{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:kubeconfig" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"kubeconfig",
		"index",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"kubeconfig",
		&pkg{version},
	)
}
