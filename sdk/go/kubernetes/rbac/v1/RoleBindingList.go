// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// RoleBindingList is a collection of RoleBindings
type RoleBindingList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`

	// Items is a list of RoleBindings
	Items RoleBindingPropertyArrayOutput `pulumi:"items"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`

	// Standard object's metadata.
	Metadata metaV1.ListMetaOutput `pulumi:"metadata"`

}

// RoleBindingListArgs is the set of arguments needed to create a RoleBindingList resource.
type RoleBindingListArgs struct {
	// Items is a list of RoleBindings
	Items RoleBindingPropertyArrayInput `pulumi:"items"`

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

	// Standard object's metadata.
	Metadata metaV1.ListMetaInput `pulumi:"metadata"`

}

// NewRoleBindingList creates a RoleBindingList resource with the given unique name, arguments, and options.
func NewRoleBindingList(ctx *pulumi.Context, name string, args *RoleBindingListArgs, opts ...pulumi.ResourceOption) (*RoleBindingList, error) {
	inputs := map[string]pulumi.Input{}
	if args == nil || args.Items == nil {
		return nil, errors.New("missing required argument 'Items'")
	}
	if args != nil {
		args.ApiVersion = pulumi.String("rbac.authorization.k8s.io/v1")
		args.Kind = pulumi.String("RoleBindingList")
		inputs["items"] = args.Items.ToRoleBindingPropertyArrayOutput()
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
	var resource RoleBindingList
	err := ctx.RegisterResource("kubernetes:rbac.authorization.k8s.io/v1:RoleBindingList", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleBindingList gets an existing RoleBindingList resource's state with the given name and ID.
func GetRoleBindingList(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*RoleBindingList, error) {
	var resource RoleBindingList
	err := ctx.ReadResource("kubernetes:rbac.authorization.k8s.io/v1:RoleBindingList", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

