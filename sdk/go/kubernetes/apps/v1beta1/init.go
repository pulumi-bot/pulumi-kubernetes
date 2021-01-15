// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "kubernetes:apps/v1beta1:ControllerRevision":
		r, err = NewControllerRevision(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:apps/v1beta1:ControllerRevisionList":
		r, err = NewControllerRevisionList(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:apps/v1beta1:Deployment":
		r, err = NewDeployment(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:apps/v1beta1:DeploymentList":
		r, err = NewDeploymentList(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:apps/v1beta1:StatefulSet":
		r, err = NewStatefulSet(ctx, name, nil, pulumi.URN_(urn))
	case "kubernetes:apps/v1beta1:StatefulSetList":
		r, err = NewStatefulSetList(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := kubernetes.PkgVersion()
	if err != nil {
		panic(fmt.Errorf("failed to determine package version: %v", err))
	}
	pulumi.RegisterResourceModule(
		"kubernetes",
		"apps/v1beta1",
		&module{version},
	)
}
