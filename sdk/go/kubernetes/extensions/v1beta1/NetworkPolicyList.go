// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// DEPRECATED 1.9 - This group version of NetworkPolicyList is deprecated by
// networking/v1/NetworkPolicyList. Network Policy List is a list of NetworkPolicy objects.
type NetworkPolicyList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`

	// Items is a list of schema objects.
	Items NetworkPolicyPropertyArrayOutput `pulumi:"items"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`

	// Standard list metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ListMetaOutput `pulumi:"metadata"`

}

// NetworkPolicyListArgs is the set of arguments needed to create a NetworkPolicyList resource.
type NetworkPolicyListArgs struct {
	// Items is a list of schema objects.
	Items NetworkPolicyPropertyArrayInput `pulumi:"items"`

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
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ListMetaInput `pulumi:"metadata"`

}

// NewNetworkPolicyList creates a NetworkPolicyList resource with the given unique name, arguments, and options.
func NewNetworkPolicyList(ctx *pulumi.Context, name string, args *NetworkPolicyListArgs, opts ...pulumi.ResourceOption) (*NetworkPolicyList, error) {
	inputs := map[string]pulumi.Input{}
	if args == nil || args.Items == nil {
		return nil, errors.New("missing required argument 'Items'")
	}
	if args != nil {
		args.ApiVersion = pulumi.String("extensions/v1beta1")
		args.Kind = pulumi.String("NetworkPolicyList")
		inputs["items"] = args.Items.ToNetworkPolicyPropertyArrayOutput()
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
	var resource NetworkPolicyList
	err := ctx.RegisterResource("kubernetes:extensions/v1beta1:NetworkPolicyList", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkPolicyList gets an existing NetworkPolicyList resource's state with the given name and ID.
func GetNetworkPolicyList(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*NetworkPolicyList, error) {
	var resource NetworkPolicyList
	err := ctx.ReadResource("kubernetes:extensions/v1beta1:NetworkPolicyList", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

