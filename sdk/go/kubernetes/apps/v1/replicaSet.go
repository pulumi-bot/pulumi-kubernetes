// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ReplicaSet ensures that a specified number of pod replicas are running at any given time.
type ReplicaSet struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// If the Labels of a ReplicaSet are empty, they are defaulted to be the same as the Pod(s) that the ReplicaSet manages. Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec defines the specification of the desired behavior of the ReplicaSet. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec ReplicaSetSpecPtrOutput `pulumi:"spec"`
	// Status is the most recently observed status of the ReplicaSet. This data may be out of date by some window of time. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status ReplicaSetStatusPtrOutput `pulumi:"status"`
}

// NewReplicaSet registers a new resource with the given unique name, arguments, and options.
func NewReplicaSet(ctx *pulumi.Context,
	name string, args *ReplicaSetArgs, opts ...pulumi.ResourceOption) (*ReplicaSet, error) {
	if args == nil {
		args = &ReplicaSetArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("apps/v1")
	args.Kind = pulumi.StringPtr("ReplicaSet")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:apps/v1beta2:ReplicaSet"),
		},
		{
			Type: pulumi.String("kubernetes:extensions/v1beta1:ReplicaSet"),
		},
	})
	opts = append(opts, aliases)
	var resource ReplicaSet
	err := ctx.RegisterResource("kubernetes:apps/v1:ReplicaSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicaSet gets an existing ReplicaSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicaSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicaSetState, opts ...pulumi.ResourceOption) (*ReplicaSet, error) {
	var resource ReplicaSet
	err := ctx.ReadResource("kubernetes:apps/v1:ReplicaSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicaSet resources.
type replicaSetState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// If the Labels of a ReplicaSet are empty, they are defaulted to be the same as the Pod(s) that the ReplicaSet manages. Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec defines the specification of the desired behavior of the ReplicaSet. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *ReplicaSetSpec `pulumi:"spec"`
	// Status is the most recently observed status of the ReplicaSet. This data may be out of date by some window of time. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status *ReplicaSetStatus `pulumi:"status"`
}

type ReplicaSetState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// If the Labels of a ReplicaSet are empty, they are defaulted to be the same as the Pod(s) that the ReplicaSet manages. Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Spec defines the specification of the desired behavior of the ReplicaSet. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec ReplicaSetSpecPtrInput
	// Status is the most recently observed status of the ReplicaSet. This data may be out of date by some window of time. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status ReplicaSetStatusPtrInput
}

func (ReplicaSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicaSetState)(nil)).Elem()
}

type replicaSetArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// If the Labels of a ReplicaSet are empty, they are defaulted to be the same as the Pod(s) that the ReplicaSet manages. Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec defines the specification of the desired behavior of the ReplicaSet. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *ReplicaSetSpec `pulumi:"spec"`
}

// The set of arguments for constructing a ReplicaSet resource.
type ReplicaSetArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// If the Labels of a ReplicaSet are empty, they are defaulted to be the same as the Pod(s) that the ReplicaSet manages. Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Spec defines the specification of the desired behavior of the ReplicaSet. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec ReplicaSetSpecPtrInput
}

func (ReplicaSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicaSetArgs)(nil)).Elem()
}

type ReplicaSetInput interface {
	pulumi.Input

	ToReplicaSetOutput() ReplicaSetOutput
	ToReplicaSetOutputWithContext(ctx context.Context) ReplicaSetOutput
}

func (*ReplicaSet) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicaSet)(nil))
}

func (i *ReplicaSet) ToReplicaSetOutput() ReplicaSetOutput {
	return i.ToReplicaSetOutputWithContext(context.Background())
}

func (i *ReplicaSet) ToReplicaSetOutputWithContext(ctx context.Context) ReplicaSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaSetOutput)
}

func (i *ReplicaSet) ToReplicaSetPtrOutput() ReplicaSetPtrOutput {
	return i.ToReplicaSetPtrOutputWithContext(context.Background())
}

func (i *ReplicaSet) ToReplicaSetPtrOutputWithContext(ctx context.Context) ReplicaSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaSetPtrOutput)
}

type ReplicaSetPtrInput interface {
	pulumi.Input

	ToReplicaSetPtrOutput() ReplicaSetPtrOutput
	ToReplicaSetPtrOutputWithContext(ctx context.Context) ReplicaSetPtrOutput
}

type replicaSetPtrType ReplicaSetArgs

func (*replicaSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicaSet)(nil))
}

func (i *replicaSetPtrType) ToReplicaSetPtrOutput() ReplicaSetPtrOutput {
	return i.ToReplicaSetPtrOutputWithContext(context.Background())
}

func (i *replicaSetPtrType) ToReplicaSetPtrOutputWithContext(ctx context.Context) ReplicaSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaSetOutput).ToReplicaSetPtrOutput()
}

// ReplicaSetArrayInput is an input type that accepts ReplicaSetArray and ReplicaSetArrayOutput values.
// You can construct a concrete instance of `ReplicaSetArrayInput` via:
//
//          ReplicaSetArray{ ReplicaSetArgs{...} }
type ReplicaSetArrayInput interface {
	pulumi.Input

	ToReplicaSetArrayOutput() ReplicaSetArrayOutput
	ToReplicaSetArrayOutputWithContext(context.Context) ReplicaSetArrayOutput
}

type ReplicaSetArray []ReplicaSetInput

func (ReplicaSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReplicaSet)(nil))
}

func (i ReplicaSetArray) ToReplicaSetArrayOutput() ReplicaSetArrayOutput {
	return i.ToReplicaSetArrayOutputWithContext(context.Background())
}

func (i ReplicaSetArray) ToReplicaSetArrayOutputWithContext(ctx context.Context) ReplicaSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaSetArrayOutput)
}

// ReplicaSetMapInput is an input type that accepts ReplicaSetMap and ReplicaSetMapOutput values.
// You can construct a concrete instance of `ReplicaSetMapInput` via:
//
//          ReplicaSetMap{ "key": ReplicaSetArgs{...} }
type ReplicaSetMapInput interface {
	pulumi.Input

	ToReplicaSetMapOutput() ReplicaSetMapOutput
	ToReplicaSetMapOutputWithContext(context.Context) ReplicaSetMapOutput
}

type ReplicaSetMap map[string]ReplicaSetInput

func (ReplicaSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReplicaSet)(nil))
}

func (i ReplicaSetMap) ToReplicaSetMapOutput() ReplicaSetMapOutput {
	return i.ToReplicaSetMapOutputWithContext(context.Background())
}

func (i ReplicaSetMap) ToReplicaSetMapOutputWithContext(ctx context.Context) ReplicaSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaSetMapOutput)
}

type ReplicaSetOutput struct {
	*pulumi.OutputState
}

func (ReplicaSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicaSet)(nil))
}

func (o ReplicaSetOutput) ToReplicaSetOutput() ReplicaSetOutput {
	return o
}

func (o ReplicaSetOutput) ToReplicaSetOutputWithContext(ctx context.Context) ReplicaSetOutput {
	return o
}

func (o ReplicaSetOutput) ToReplicaSetPtrOutput() ReplicaSetPtrOutput {
	return o.ToReplicaSetPtrOutputWithContext(context.Background())
}

func (o ReplicaSetOutput) ToReplicaSetPtrOutputWithContext(ctx context.Context) ReplicaSetPtrOutput {
	return o.ApplyT(func(v ReplicaSet) *ReplicaSet {
		return &v
	}).(ReplicaSetPtrOutput)
}

type ReplicaSetPtrOutput struct {
	*pulumi.OutputState
}

func (ReplicaSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicaSet)(nil))
}

func (o ReplicaSetPtrOutput) ToReplicaSetPtrOutput() ReplicaSetPtrOutput {
	return o
}

func (o ReplicaSetPtrOutput) ToReplicaSetPtrOutputWithContext(ctx context.Context) ReplicaSetPtrOutput {
	return o
}

type ReplicaSetArrayOutput struct{ *pulumi.OutputState }

func (ReplicaSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReplicaSet)(nil))
}

func (o ReplicaSetArrayOutput) ToReplicaSetArrayOutput() ReplicaSetArrayOutput {
	return o
}

func (o ReplicaSetArrayOutput) ToReplicaSetArrayOutputWithContext(ctx context.Context) ReplicaSetArrayOutput {
	return o
}

func (o ReplicaSetArrayOutput) Index(i pulumi.IntInput) ReplicaSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReplicaSet {
		return vs[0].([]ReplicaSet)[vs[1].(int)]
	}).(ReplicaSetOutput)
}

type ReplicaSetMapOutput struct{ *pulumi.OutputState }

func (ReplicaSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReplicaSet)(nil))
}

func (o ReplicaSetMapOutput) ToReplicaSetMapOutput() ReplicaSetMapOutput {
	return o
}

func (o ReplicaSetMapOutput) ToReplicaSetMapOutputWithContext(ctx context.Context) ReplicaSetMapOutput {
	return o
}

func (o ReplicaSetMapOutput) MapIndex(k pulumi.StringInput) ReplicaSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReplicaSet {
		return vs[0].(map[string]ReplicaSet)[vs[1].(string)]
	}).(ReplicaSetOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicaSetOutput{})
	pulumi.RegisterOutputType(ReplicaSetPtrOutput{})
	pulumi.RegisterOutputType(ReplicaSetArrayOutput{})
	pulumi.RegisterOutputType(ReplicaSetMapOutput{})
}
