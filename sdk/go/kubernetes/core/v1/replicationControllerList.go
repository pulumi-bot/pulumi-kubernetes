// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ReplicationControllerList is a collection of replication controllers.
type ReplicationControllerList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// List of replication controllers. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Items ReplicationControllerTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewReplicationControllerList registers a new resource with the given unique name, arguments, and options.
func NewReplicationControllerList(ctx *pulumi.Context,
	name string, args *ReplicationControllerListArgs, opts ...pulumi.ResourceOption) (*ReplicationControllerList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("ReplicationControllerList")
	var resource ReplicationControllerList
	err := ctx.RegisterResource("kubernetes:core/v1:ReplicationControllerList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationControllerList gets an existing ReplicationControllerList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationControllerList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationControllerListState, opts ...pulumi.ResourceOption) (*ReplicationControllerList, error) {
	var resource ReplicationControllerList
	err := ctx.ReadResource("kubernetes:core/v1:ReplicationControllerList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationControllerList resources.
type replicationControllerListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of replication controllers. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Items []ReplicationControllerType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type ReplicationControllerListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of replication controllers. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Items ReplicationControllerTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ReplicationControllerListState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationControllerListState)(nil)).Elem()
}

type replicationControllerListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of replication controllers. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Items []ReplicationControllerType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ReplicationControllerList resource.
type ReplicationControllerListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of replication controllers. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Items ReplicationControllerTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ReplicationControllerListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationControllerListArgs)(nil)).Elem()
}

type ReplicationControllerListInput interface {
	pulumi.Input

	ToReplicationControllerListOutput() ReplicationControllerListOutput
	ToReplicationControllerListOutputWithContext(ctx context.Context) ReplicationControllerListOutput
}

func (*ReplicationControllerList) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationControllerList)(nil))
}

func (i *ReplicationControllerList) ToReplicationControllerListOutput() ReplicationControllerListOutput {
	return i.ToReplicationControllerListOutputWithContext(context.Background())
}

func (i *ReplicationControllerList) ToReplicationControllerListOutputWithContext(ctx context.Context) ReplicationControllerListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationControllerListOutput)
}

func (i *ReplicationControllerList) ToReplicationControllerListPtrOutput() ReplicationControllerListPtrOutput {
	return i.ToReplicationControllerListPtrOutputWithContext(context.Background())
}

func (i *ReplicationControllerList) ToReplicationControllerListPtrOutputWithContext(ctx context.Context) ReplicationControllerListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationControllerListPtrOutput)
}

type ReplicationControllerListPtrInput interface {
	pulumi.Input

	ToReplicationControllerListPtrOutput() ReplicationControllerListPtrOutput
	ToReplicationControllerListPtrOutputWithContext(ctx context.Context) ReplicationControllerListPtrOutput
}

type ReplicationControllerListOutput struct {
	*pulumi.OutputState
}

func (ReplicationControllerListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationControllerList)(nil))
}

func (o ReplicationControllerListOutput) ToReplicationControllerListOutput() ReplicationControllerListOutput {
	return o
}

func (o ReplicationControllerListOutput) ToReplicationControllerListOutputWithContext(ctx context.Context) ReplicationControllerListOutput {
	return o
}

type ReplicationControllerListPtrOutput struct {
	*pulumi.OutputState
}

func (ReplicationControllerListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationControllerList)(nil))
}

func (o ReplicationControllerListPtrOutput) ToReplicationControllerListPtrOutput() ReplicationControllerListPtrOutput {
	return o
}

func (o ReplicationControllerListPtrOutput) ToReplicationControllerListPtrOutputWithContext(ctx context.Context) ReplicationControllerListPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ReplicationControllerListOutput{})
	pulumi.RegisterOutputType(ReplicationControllerListPtrOutput{})
}
