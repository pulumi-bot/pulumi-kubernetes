// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ReplicaSetList is a collection of ReplicaSets.
type ReplicaSetList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// List of ReplicaSets. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Items ReplicaSetTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewReplicaSetList registers a new resource with the given unique name, arguments, and options.
func NewReplicaSetList(ctx *pulumi.Context,
	name string, args *ReplicaSetListArgs, opts ...pulumi.ResourceOption) (*ReplicaSetList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("apps/v1beta2")
	args.Kind = pulumi.StringPtr("ReplicaSetList")
	var resource ReplicaSetList
	err := ctx.RegisterResource("kubernetes:apps/v1beta2:ReplicaSetList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicaSetList gets an existing ReplicaSetList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicaSetList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicaSetListState, opts ...pulumi.ResourceOption) (*ReplicaSetList, error) {
	var resource ReplicaSetList
	err := ctx.ReadResource("kubernetes:apps/v1beta2:ReplicaSetList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicaSetList resources.
type replicaSetListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of ReplicaSets. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Items []ReplicaSetType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type ReplicaSetListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of ReplicaSets. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Items ReplicaSetTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ReplicaSetListState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicaSetListState)(nil)).Elem()
}

type replicaSetListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of ReplicaSets. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Items []ReplicaSetType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ReplicaSetList resource.
type ReplicaSetListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of ReplicaSets. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Items ReplicaSetTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ReplicaSetListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicaSetListArgs)(nil)).Elem()
}

type ReplicaSetListInput interface {
	pulumi.Input

	ToReplicaSetListOutput() ReplicaSetListOutput
	ToReplicaSetListOutputWithContext(ctx context.Context) ReplicaSetListOutput
}

func (*ReplicaSetList) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicaSetList)(nil))
}

func (i *ReplicaSetList) ToReplicaSetListOutput() ReplicaSetListOutput {
	return i.ToReplicaSetListOutputWithContext(context.Background())
}

func (i *ReplicaSetList) ToReplicaSetListOutputWithContext(ctx context.Context) ReplicaSetListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaSetListOutput)
}

func (i *ReplicaSetList) ToReplicaSetListPtrOutput() ReplicaSetListPtrOutput {
	return i.ToReplicaSetListPtrOutputWithContext(context.Background())
}

func (i *ReplicaSetList) ToReplicaSetListPtrOutputWithContext(ctx context.Context) ReplicaSetListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaSetListPtrOutput)
}

type ReplicaSetListPtrInput interface {
	pulumi.Input

	ToReplicaSetListPtrOutput() ReplicaSetListPtrOutput
	ToReplicaSetListPtrOutputWithContext(ctx context.Context) ReplicaSetListPtrOutput
}

type replicaSetListPtrType ReplicaSetListArgs

func (*replicaSetListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicaSetList)(nil))
}

func (i *replicaSetListPtrType) ToReplicaSetListPtrOutput() ReplicaSetListPtrOutput {
	return i.ToReplicaSetListPtrOutputWithContext(context.Background())
}

func (i *replicaSetListPtrType) ToReplicaSetListPtrOutputWithContext(ctx context.Context) ReplicaSetListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaSetListPtrOutput)
}

// ReplicaSetListArrayInput is an input type that accepts ReplicaSetListArray and ReplicaSetListArrayOutput values.
// You can construct a concrete instance of `ReplicaSetListArrayInput` via:
//
//          ReplicaSetListArray{ ReplicaSetListArgs{...} }
type ReplicaSetListArrayInput interface {
	pulumi.Input

	ToReplicaSetListArrayOutput() ReplicaSetListArrayOutput
	ToReplicaSetListArrayOutputWithContext(context.Context) ReplicaSetListArrayOutput
}

type ReplicaSetListArray []ReplicaSetListInput

func (ReplicaSetListArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ReplicaSetList)(nil))
}

func (i ReplicaSetListArray) ToReplicaSetListArrayOutput() ReplicaSetListArrayOutput {
	return i.ToReplicaSetListArrayOutputWithContext(context.Background())
}

func (i ReplicaSetListArray) ToReplicaSetListArrayOutputWithContext(ctx context.Context) ReplicaSetListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaSetListArrayOutput)
}

// ReplicaSetListMapInput is an input type that accepts ReplicaSetListMap and ReplicaSetListMapOutput values.
// You can construct a concrete instance of `ReplicaSetListMapInput` via:
//
//          ReplicaSetListMap{ "key": ReplicaSetListArgs{...} }
type ReplicaSetListMapInput interface {
	pulumi.Input

	ToReplicaSetListMapOutput() ReplicaSetListMapOutput
	ToReplicaSetListMapOutputWithContext(context.Context) ReplicaSetListMapOutput
}

type ReplicaSetListMap map[string]ReplicaSetListInput

func (ReplicaSetListMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ReplicaSetList)(nil))
}

func (i ReplicaSetListMap) ToReplicaSetListMapOutput() ReplicaSetListMapOutput {
	return i.ToReplicaSetListMapOutputWithContext(context.Background())
}

func (i ReplicaSetListMap) ToReplicaSetListMapOutputWithContext(ctx context.Context) ReplicaSetListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaSetListMapOutput)
}

type ReplicaSetListOutput struct {
	*pulumi.OutputState
}

func (ReplicaSetListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicaSetList)(nil))
}

func (o ReplicaSetListOutput) ToReplicaSetListOutput() ReplicaSetListOutput {
	return o
}

func (o ReplicaSetListOutput) ToReplicaSetListOutputWithContext(ctx context.Context) ReplicaSetListOutput {
	return o
}

func (o ReplicaSetListOutput) ToReplicaSetListPtrOutput() ReplicaSetListPtrOutput {
	return o.ToReplicaSetListPtrOutputWithContext(context.Background())
}

func (o ReplicaSetListOutput) ToReplicaSetListPtrOutputWithContext(ctx context.Context) ReplicaSetListPtrOutput {
	return o.ApplyT(func(v ReplicaSetList) *ReplicaSetList {
		return &v
	}).(ReplicaSetListPtrOutput)
}

type ReplicaSetListPtrOutput struct {
	*pulumi.OutputState
}

func (ReplicaSetListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicaSetList)(nil))
}

func (o ReplicaSetListPtrOutput) ToReplicaSetListPtrOutput() ReplicaSetListPtrOutput {
	return o
}

func (o ReplicaSetListPtrOutput) ToReplicaSetListPtrOutputWithContext(ctx context.Context) ReplicaSetListPtrOutput {
	return o
}

type ReplicaSetListArrayOutput struct{ *pulumi.OutputState }

func (ReplicaSetListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReplicaSetList)(nil))
}

func (o ReplicaSetListArrayOutput) ToReplicaSetListArrayOutput() ReplicaSetListArrayOutput {
	return o
}

func (o ReplicaSetListArrayOutput) ToReplicaSetListArrayOutputWithContext(ctx context.Context) ReplicaSetListArrayOutput {
	return o
}

func (o ReplicaSetListArrayOutput) Index(i pulumi.IntInput) ReplicaSetListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReplicaSetList {
		return vs[0].([]ReplicaSetList)[vs[1].(int)]
	}).(ReplicaSetListOutput)
}

type ReplicaSetListMapOutput struct{ *pulumi.OutputState }

func (ReplicaSetListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReplicaSetList)(nil))
}

func (o ReplicaSetListMapOutput) ToReplicaSetListMapOutput() ReplicaSetListMapOutput {
	return o
}

func (o ReplicaSetListMapOutput) ToReplicaSetListMapOutputWithContext(ctx context.Context) ReplicaSetListMapOutput {
	return o
}

func (o ReplicaSetListMapOutput) MapIndex(k pulumi.StringInput) ReplicaSetListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReplicaSetList {
		return vs[0].(map[string]ReplicaSetList)[vs[1].(string)]
	}).(ReplicaSetListOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicaSetListOutput{})
	pulumi.RegisterOutputType(ReplicaSetListPtrOutput{})
	pulumi.RegisterOutputType(ReplicaSetListArrayOutput{})
	pulumi.RegisterOutputType(ReplicaSetListMapOutput{})
}
