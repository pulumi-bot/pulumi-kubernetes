// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ResourceQuota sets aggregate quota restrictions enforced per namespace
type ResourceQuota struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec defines the desired quota. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec ResourceQuotaSpecPtrOutput `pulumi:"spec"`
	// Status defines the actual enforced quota and its current usage. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status ResourceQuotaStatusPtrOutput `pulumi:"status"`
}

// NewResourceQuota registers a new resource with the given unique name, arguments, and options.
func NewResourceQuota(ctx *pulumi.Context,
	name string, args *ResourceQuotaArgs, opts ...pulumi.ResourceOption) (*ResourceQuota, error) {
	if args == nil {
		args = &ResourceQuotaArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("ResourceQuota")
	var resource ResourceQuota
	err := ctx.RegisterResource("kubernetes:core/v1:ResourceQuota", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceQuota gets an existing ResourceQuota resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceQuota(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceQuotaState, opts ...pulumi.ResourceOption) (*ResourceQuota, error) {
	var resource ResourceQuota
	err := ctx.ReadResource("kubernetes:core/v1:ResourceQuota", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceQuota resources.
type resourceQuotaState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec defines the desired quota. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *ResourceQuotaSpec `pulumi:"spec"`
	// Status defines the actual enforced quota and its current usage. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status *ResourceQuotaStatus `pulumi:"status"`
}

type ResourceQuotaState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Spec defines the desired quota. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec ResourceQuotaSpecPtrInput
	// Status defines the actual enforced quota and its current usage. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status ResourceQuotaStatusPtrInput
}

func (ResourceQuotaState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceQuotaState)(nil)).Elem()
}

type resourceQuotaArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec defines the desired quota. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *ResourceQuotaSpec `pulumi:"spec"`
}

// The set of arguments for constructing a ResourceQuota resource.
type ResourceQuotaArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Spec defines the desired quota. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec ResourceQuotaSpecPtrInput
}

func (ResourceQuotaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceQuotaArgs)(nil)).Elem()
}

type ResourceQuotaInput interface {
	pulumi.Input

	ToResourceQuotaOutput() ResourceQuotaOutput
	ToResourceQuotaOutputWithContext(ctx context.Context) ResourceQuotaOutput
}

func (*ResourceQuota) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceQuota)(nil))
}

func (i *ResourceQuota) ToResourceQuotaOutput() ResourceQuotaOutput {
	return i.ToResourceQuotaOutputWithContext(context.Background())
}

func (i *ResourceQuota) ToResourceQuotaOutputWithContext(ctx context.Context) ResourceQuotaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceQuotaOutput)
}

func (i *ResourceQuota) ToResourceQuotaPtrOutput() ResourceQuotaPtrOutput {
	return i.ToResourceQuotaPtrOutputWithContext(context.Background())
}

func (i *ResourceQuota) ToResourceQuotaPtrOutputWithContext(ctx context.Context) ResourceQuotaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceQuotaPtrOutput)
}

type ResourceQuotaPtrInput interface {
	pulumi.Input

	ToResourceQuotaPtrOutput() ResourceQuotaPtrOutput
	ToResourceQuotaPtrOutputWithContext(ctx context.Context) ResourceQuotaPtrOutput
}

type resourceQuotaPtrType ResourceQuotaArgs

func (*resourceQuotaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceQuota)(nil))
}

func (i *resourceQuotaPtrType) ToResourceQuotaPtrOutput() ResourceQuotaPtrOutput {
	return i.ToResourceQuotaPtrOutputWithContext(context.Background())
}

func (i *resourceQuotaPtrType) ToResourceQuotaPtrOutputWithContext(ctx context.Context) ResourceQuotaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceQuotaOutput).ToResourceQuotaPtrOutput()
}

// ResourceQuotaArrayInput is an input type that accepts ResourceQuotaArray and ResourceQuotaArrayOutput values.
// You can construct a concrete instance of `ResourceQuotaArrayInput` via:
//
//          ResourceQuotaArray{ ResourceQuotaArgs{...} }
type ResourceQuotaArrayInput interface {
	pulumi.Input

	ToResourceQuotaArrayOutput() ResourceQuotaArrayOutput
	ToResourceQuotaArrayOutputWithContext(context.Context) ResourceQuotaArrayOutput
}

type ResourceQuotaArray []ResourceQuotaInput

func (ResourceQuotaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceQuota)(nil))
}

func (i ResourceQuotaArray) ToResourceQuotaArrayOutput() ResourceQuotaArrayOutput {
	return i.ToResourceQuotaArrayOutputWithContext(context.Background())
}

func (i ResourceQuotaArray) ToResourceQuotaArrayOutputWithContext(ctx context.Context) ResourceQuotaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceQuotaArrayOutput)
}

// ResourceQuotaMapInput is an input type that accepts ResourceQuotaMap and ResourceQuotaMapOutput values.
// You can construct a concrete instance of `ResourceQuotaMapInput` via:
//
//          ResourceQuotaMap{ "key": ResourceQuotaArgs{...} }
type ResourceQuotaMapInput interface {
	pulumi.Input

	ToResourceQuotaMapOutput() ResourceQuotaMapOutput
	ToResourceQuotaMapOutputWithContext(context.Context) ResourceQuotaMapOutput
}

type ResourceQuotaMap map[string]ResourceQuotaInput

func (ResourceQuotaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceQuota)(nil))
}

func (i ResourceQuotaMap) ToResourceQuotaMapOutput() ResourceQuotaMapOutput {
	return i.ToResourceQuotaMapOutputWithContext(context.Background())
}

func (i ResourceQuotaMap) ToResourceQuotaMapOutputWithContext(ctx context.Context) ResourceQuotaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceQuotaMapOutput)
}

type ResourceQuotaOutput struct {
	*pulumi.OutputState
}

func (ResourceQuotaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceQuota)(nil))
}

func (o ResourceQuotaOutput) ToResourceQuotaOutput() ResourceQuotaOutput {
	return o
}

func (o ResourceQuotaOutput) ToResourceQuotaOutputWithContext(ctx context.Context) ResourceQuotaOutput {
	return o
}

func (o ResourceQuotaOutput) ToResourceQuotaPtrOutput() ResourceQuotaPtrOutput {
	return o.ToResourceQuotaPtrOutputWithContext(context.Background())
}

func (o ResourceQuotaOutput) ToResourceQuotaPtrOutputWithContext(ctx context.Context) ResourceQuotaPtrOutput {
	return o.ApplyT(func(v ResourceQuota) *ResourceQuota {
		return &v
	}).(ResourceQuotaPtrOutput)
}

type ResourceQuotaPtrOutput struct {
	*pulumi.OutputState
}

func (ResourceQuotaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceQuota)(nil))
}

func (o ResourceQuotaPtrOutput) ToResourceQuotaPtrOutput() ResourceQuotaPtrOutput {
	return o
}

func (o ResourceQuotaPtrOutput) ToResourceQuotaPtrOutputWithContext(ctx context.Context) ResourceQuotaPtrOutput {
	return o
}

type ResourceQuotaArrayOutput struct{ *pulumi.OutputState }

func (ResourceQuotaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceQuota)(nil))
}

func (o ResourceQuotaArrayOutput) ToResourceQuotaArrayOutput() ResourceQuotaArrayOutput {
	return o
}

func (o ResourceQuotaArrayOutput) ToResourceQuotaArrayOutputWithContext(ctx context.Context) ResourceQuotaArrayOutput {
	return o
}

func (o ResourceQuotaArrayOutput) Index(i pulumi.IntInput) ResourceQuotaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceQuota {
		return vs[0].([]ResourceQuota)[vs[1].(int)]
	}).(ResourceQuotaOutput)
}

type ResourceQuotaMapOutput struct{ *pulumi.OutputState }

func (ResourceQuotaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceQuota)(nil))
}

func (o ResourceQuotaMapOutput) ToResourceQuotaMapOutput() ResourceQuotaMapOutput {
	return o
}

func (o ResourceQuotaMapOutput) ToResourceQuotaMapOutputWithContext(ctx context.Context) ResourceQuotaMapOutput {
	return o
}

func (o ResourceQuotaMapOutput) MapIndex(k pulumi.StringInput) ResourceQuotaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResourceQuota {
		return vs[0].(map[string]ResourceQuota)[vs[1].(string)]
	}).(ResourceQuotaOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceQuotaOutput{})
	pulumi.RegisterOutputType(ResourceQuotaPtrOutput{})
	pulumi.RegisterOutputType(ResourceQuotaArrayOutput{})
	pulumi.RegisterOutputType(ResourceQuotaMapOutput{})
}
