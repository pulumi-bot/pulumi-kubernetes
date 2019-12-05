// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// EndpointSliceList represents a list of endpoint slices
type EndpointSliceList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`

	// List of endpoint slices
	Items EndpointSlicePropertyArrayOutput `pulumi:"items"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`

	// Standard list metadata.
	Metadata metaV1.ListMetaOutput `pulumi:"metadata"`

}

// EndpointSliceListArgs is the set of arguments needed to create a EndpointSliceList resource.
type EndpointSliceListArgs struct {
	// List of endpoint slices
	Items EndpointSlicePropertyArrayInput `pulumi:"items"`

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

	// Standard list metadata.
	Metadata metaV1.ListMetaInput `pulumi:"metadata"`

}

// NewEndpointSliceList creates a EndpointSliceList resource with the given unique name, arguments, and options.
func NewEndpointSliceList(ctx *pulumi.Context, name string, args *EndpointSliceListArgs, opts ...pulumi.ResourceOption) (*EndpointSliceList, error) {
	inputs := map[string]pulumi.Input{}
	if args == nil || args.Items == nil {
		return nil, errors.New("missing required argument 'Items'")
	}
	if args != nil {
		args.ApiVersion = pulumi.String("discovery.k8s.io/v1alpha1")
		args.Kind = pulumi.String("EndpointSliceList")
		inputs["items"] = args.Items.ToEndpointSlicePropertyArrayOutput()
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
	var resource EndpointSliceList
	err := ctx.RegisterResource("kubernetes:discovery.k8s.io/v1alpha1:EndpointSliceList", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointSliceList gets an existing EndpointSliceList resource's state with the given name and ID.
func GetEndpointSliceList(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*EndpointSliceList, error) {
	var resource EndpointSliceList
	err := ctx.ReadResource("kubernetes:discovery.k8s.io/v1alpha1:EndpointSliceList", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

