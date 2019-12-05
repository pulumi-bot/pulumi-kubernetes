// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// ServiceList holds a list of services.
type ServiceList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`

	// List of services
	Items ServicePropertyArrayOutput `pulumi:"items"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`

	// Standard list metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metaV1.ListMetaOutput `pulumi:"metadata"`

}

// ServiceListArgs is the set of arguments needed to create a ServiceList resource.
type ServiceListArgs struct {
	// List of services
	Items ServicePropertyArrayInput `pulumi:"items"`

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringInput `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringInput `pulumi:"kind"`

	// Standard list metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metaV1.ListMetaInput `pulumi:"metadata"`

}

// NewServiceList creates a ServiceList resource with the given unique name, arguments, and options.
func NewServiceList(ctx *pulumi.Context, name string, args *ServiceListArgs, opts ...pulumi.ResourceOption) (*ServiceList, error) {
	inputs := map[string]pulumi.Input{}
	if args == nil || args.Items == nil {
		return nil, errors.New("missing required argument 'Items'")
	}
	if args != nil {
		args.ApiVersion = pulumi.String("v1")
		args.Kind = pulumi.String("ServiceList")
		inputs["items"] = args.Items.ToServicePropertyArrayOutput()
		if i := args.ApiVersion; i != nil {
			inputs["apiVersion"] = i.ToStringOutput()
		}
		if i := args.Kind; i != nil {
			inputs["kind"] = i.ToStringOutput()
		}
		if i := args.Metadata; i != nil {
			inputs["metadata"] = i.ToListMetaOutput()
		}
	}
	var resource ServiceList
	err := ctx.RegisterResource("kubernetes:core/v1:ServiceList", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceList gets an existing ServiceList resource's state with the given name and ID.
func GetServiceList(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*ServiceList, error) {
	var resource ServiceList
	err := ctx.ReadResource("kubernetes:core/v1:ServiceList", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

