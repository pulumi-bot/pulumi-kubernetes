// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// PriorityClassList is a collection of priority classes.
type PriorityClassList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`

	// items is the list of PriorityClasses
	Items PriorityClassPropertyArrayOutput `pulumi:"items"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`

	// Standard list metadata More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ListMetaOutput `pulumi:"metadata"`

}

// PriorityClassListArgs is the set of arguments needed to create a PriorityClassList resource.
type PriorityClassListArgs struct {
	// items is the list of PriorityClasses
	Items PriorityClassPropertyArrayInput `pulumi:"items"`

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

	// Standard list metadata More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ListMetaInput `pulumi:"metadata"`

}

// NewPriorityClassList creates a PriorityClassList resource with the given unique name, arguments, and options.
func NewPriorityClassList(ctx *pulumi.Context, name string, args *PriorityClassListArgs, opts ...pulumi.ResourceOption) (*PriorityClassList, error) {
	inputs := map[string]pulumi.Input{}
	if args == nil || args.Items == nil {
		return nil, errors.New("missing required argument 'Items'")
	}
	if args != nil {
		args.ApiVersion = pulumi.String("scheduling.k8s.io/v1")
		args.Kind = pulumi.String("PriorityClassList")
		inputs["items"] = args.Items.ToPriorityClassPropertyArrayOutput()
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
	var resource PriorityClassList
	err := ctx.RegisterResource("kubernetes:scheduling.k8s.io/v1:PriorityClassList", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPriorityClassList gets an existing PriorityClassList resource's state with the given name and ID.
func GetPriorityClassList(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*PriorityClassList, error) {
	var resource PriorityClassList
	err := ctx.ReadResource("kubernetes:scheduling.k8s.io/v1:PriorityClassList", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

