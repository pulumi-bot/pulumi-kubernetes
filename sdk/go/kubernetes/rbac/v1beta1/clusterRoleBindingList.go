// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ClusterRoleBindingList is a collection of ClusterRoleBindings. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 ClusterRoleBindingList, and will no longer be served in v1.22.
type ClusterRoleBindingList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Items is a list of ClusterRoleBindings
	Items ClusterRoleBindingTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata.
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewClusterRoleBindingList registers a new resource with the given unique name, arguments, and options.
func NewClusterRoleBindingList(ctx *pulumi.Context,
	name string, args *ClusterRoleBindingListArgs, opts ...pulumi.ResourceOption) (*ClusterRoleBindingList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("rbac.authorization.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("ClusterRoleBindingList")
	var resource ClusterRoleBindingList
	err := ctx.RegisterResource("kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRoleBindingList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterRoleBindingList gets an existing ClusterRoleBindingList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterRoleBindingList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterRoleBindingListState, opts ...pulumi.ResourceOption) (*ClusterRoleBindingList, error) {
	var resource ClusterRoleBindingList
	err := ctx.ReadResource("kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRoleBindingList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterRoleBindingList resources.
type clusterRoleBindingListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is a list of ClusterRoleBindings
	Items []ClusterRoleBindingType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type ClusterRoleBindingListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is a list of ClusterRoleBindings
	Items ClusterRoleBindingTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ListMetaPtrInput
}

func (ClusterRoleBindingListState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterRoleBindingListState)(nil)).Elem()
}

type clusterRoleBindingListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is a list of ClusterRoleBindings
	Items []ClusterRoleBindingType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ClusterRoleBindingList resource.
type ClusterRoleBindingListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is a list of ClusterRoleBindings
	Items ClusterRoleBindingTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ListMetaPtrInput
}

func (ClusterRoleBindingListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterRoleBindingListArgs)(nil)).Elem()
}

type ClusterRoleBindingListInput interface {
	pulumi.Input

	ToClusterRoleBindingListOutput() ClusterRoleBindingListOutput
	ToClusterRoleBindingListOutputWithContext(ctx context.Context) ClusterRoleBindingListOutput
}

func (ClusterRoleBindingList) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterRoleBindingList)(nil)).Elem()
}

func (i ClusterRoleBindingList) ToClusterRoleBindingListOutput() ClusterRoleBindingListOutput {
	return i.ToClusterRoleBindingListOutputWithContext(context.Background())
}

func (i ClusterRoleBindingList) ToClusterRoleBindingListOutputWithContext(ctx context.Context) ClusterRoleBindingListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoleBindingListOutput)
}

type ClusterRoleBindingListOutput struct {
	*pulumi.OutputState
}

func (ClusterRoleBindingListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterRoleBindingListOutput)(nil)).Elem()
}

func (o ClusterRoleBindingListOutput) ToClusterRoleBindingListOutput() ClusterRoleBindingListOutput {
	return o
}

func (o ClusterRoleBindingListOutput) ToClusterRoleBindingListOutputWithContext(ctx context.Context) ClusterRoleBindingListOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ClusterRoleBindingListOutput{})
}
