// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// configuration of a horizontal pod autoscaler.
type HorizontalPodAutoscaler struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// behaviour of autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec HorizontalPodAutoscalerSpecPtrOutput `pulumi:"spec"`
	// current information about the autoscaler.
	Status HorizontalPodAutoscalerStatusPtrOutput `pulumi:"status"`
}

// NewHorizontalPodAutoscaler registers a new resource with the given unique name, arguments, and options.
func NewHorizontalPodAutoscaler(ctx *pulumi.Context,
	name string, args *HorizontalPodAutoscalerArgs, opts ...pulumi.ResourceOption) (*HorizontalPodAutoscaler, error) {
	if args == nil {
		args = &HorizontalPodAutoscalerArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("autoscaling/v1")
	args.Kind = pulumi.StringPtr("HorizontalPodAutoscaler")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:autoscaling/v2beta1:HorizontalPodAutoscaler"),
		},
		{
			Type: pulumi.String("kubernetes:autoscaling/v2beta2:HorizontalPodAutoscaler"),
		},
	})
	opts = append(opts, aliases)
	var resource HorizontalPodAutoscaler
	err := ctx.RegisterResource("kubernetes:autoscaling/v1:HorizontalPodAutoscaler", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHorizontalPodAutoscaler gets an existing HorizontalPodAutoscaler resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHorizontalPodAutoscaler(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HorizontalPodAutoscalerState, opts ...pulumi.ResourceOption) (*HorizontalPodAutoscaler, error) {
	var resource HorizontalPodAutoscaler
	err := ctx.ReadResource("kubernetes:autoscaling/v1:HorizontalPodAutoscaler", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HorizontalPodAutoscaler resources.
type horizontalPodAutoscalerState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// behaviour of autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec *HorizontalPodAutoscalerSpec `pulumi:"spec"`
	// current information about the autoscaler.
	Status *HorizontalPodAutoscalerStatus `pulumi:"status"`
}

type HorizontalPodAutoscalerState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// behaviour of autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec HorizontalPodAutoscalerSpecPtrInput
	// current information about the autoscaler.
	Status HorizontalPodAutoscalerStatusPtrInput
}

func (HorizontalPodAutoscalerState) ElementType() reflect.Type {
	return reflect.TypeOf((*horizontalPodAutoscalerState)(nil)).Elem()
}

type horizontalPodAutoscalerArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// behaviour of autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec *HorizontalPodAutoscalerSpec `pulumi:"spec"`
}

// The set of arguments for constructing a HorizontalPodAutoscaler resource.
type HorizontalPodAutoscalerArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// behaviour of autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec HorizontalPodAutoscalerSpecPtrInput
}

func (HorizontalPodAutoscalerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*horizontalPodAutoscalerArgs)(nil)).Elem()
}

type HorizontalPodAutoscalerInput interface {
	pulumi.Input

	ToHorizontalPodAutoscalerOutput() HorizontalPodAutoscalerOutput
	ToHorizontalPodAutoscalerOutputWithContext(ctx context.Context) HorizontalPodAutoscalerOutput
}

func (*HorizontalPodAutoscaler) ElementType() reflect.Type {
	return reflect.TypeOf((*HorizontalPodAutoscaler)(nil))
}

func (i *HorizontalPodAutoscaler) ToHorizontalPodAutoscalerOutput() HorizontalPodAutoscalerOutput {
	return i.ToHorizontalPodAutoscalerOutputWithContext(context.Background())
}

func (i *HorizontalPodAutoscaler) ToHorizontalPodAutoscalerOutputWithContext(ctx context.Context) HorizontalPodAutoscalerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HorizontalPodAutoscalerOutput)
}

func (i *HorizontalPodAutoscaler) ToHorizontalPodAutoscalerPtrOutput() HorizontalPodAutoscalerPtrOutput {
	return i.ToHorizontalPodAutoscalerPtrOutputWithContext(context.Background())
}

func (i *HorizontalPodAutoscaler) ToHorizontalPodAutoscalerPtrOutputWithContext(ctx context.Context) HorizontalPodAutoscalerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HorizontalPodAutoscalerPtrOutput)
}

type HorizontalPodAutoscalerPtrInput interface {
	pulumi.Input

	ToHorizontalPodAutoscalerPtrOutput() HorizontalPodAutoscalerPtrOutput
	ToHorizontalPodAutoscalerPtrOutputWithContext(ctx context.Context) HorizontalPodAutoscalerPtrOutput
}

type horizontalPodAutoscalerPtrType HorizontalPodAutoscalerArgs

func (*horizontalPodAutoscalerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HorizontalPodAutoscaler)(nil))
}

func (i *horizontalPodAutoscalerPtrType) ToHorizontalPodAutoscalerPtrOutput() HorizontalPodAutoscalerPtrOutput {
	return i.ToHorizontalPodAutoscalerPtrOutputWithContext(context.Background())
}

func (i *horizontalPodAutoscalerPtrType) ToHorizontalPodAutoscalerPtrOutputWithContext(ctx context.Context) HorizontalPodAutoscalerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HorizontalPodAutoscalerOutput).ToHorizontalPodAutoscalerPtrOutput()
}

// HorizontalPodAutoscalerArrayInput is an input type that accepts HorizontalPodAutoscalerArray and HorizontalPodAutoscalerArrayOutput values.
// You can construct a concrete instance of `HorizontalPodAutoscalerArrayInput` via:
//
//          HorizontalPodAutoscalerArray{ HorizontalPodAutoscalerArgs{...} }
type HorizontalPodAutoscalerArrayInput interface {
	pulumi.Input

	ToHorizontalPodAutoscalerArrayOutput() HorizontalPodAutoscalerArrayOutput
	ToHorizontalPodAutoscalerArrayOutputWithContext(context.Context) HorizontalPodAutoscalerArrayOutput
}

type HorizontalPodAutoscalerArray []HorizontalPodAutoscalerInput

func (HorizontalPodAutoscalerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HorizontalPodAutoscaler)(nil))
}

func (i HorizontalPodAutoscalerArray) ToHorizontalPodAutoscalerArrayOutput() HorizontalPodAutoscalerArrayOutput {
	return i.ToHorizontalPodAutoscalerArrayOutputWithContext(context.Background())
}

func (i HorizontalPodAutoscalerArray) ToHorizontalPodAutoscalerArrayOutputWithContext(ctx context.Context) HorizontalPodAutoscalerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HorizontalPodAutoscalerArrayOutput)
}

// HorizontalPodAutoscalerMapInput is an input type that accepts HorizontalPodAutoscalerMap and HorizontalPodAutoscalerMapOutput values.
// You can construct a concrete instance of `HorizontalPodAutoscalerMapInput` via:
//
//          HorizontalPodAutoscalerMap{ "key": HorizontalPodAutoscalerArgs{...} }
type HorizontalPodAutoscalerMapInput interface {
	pulumi.Input

	ToHorizontalPodAutoscalerMapOutput() HorizontalPodAutoscalerMapOutput
	ToHorizontalPodAutoscalerMapOutputWithContext(context.Context) HorizontalPodAutoscalerMapOutput
}

type HorizontalPodAutoscalerMap map[string]HorizontalPodAutoscalerInput

func (HorizontalPodAutoscalerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]HorizontalPodAutoscaler)(nil))
}

func (i HorizontalPodAutoscalerMap) ToHorizontalPodAutoscalerMapOutput() HorizontalPodAutoscalerMapOutput {
	return i.ToHorizontalPodAutoscalerMapOutputWithContext(context.Background())
}

func (i HorizontalPodAutoscalerMap) ToHorizontalPodAutoscalerMapOutputWithContext(ctx context.Context) HorizontalPodAutoscalerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HorizontalPodAutoscalerMapOutput)
}

type HorizontalPodAutoscalerOutput struct {
	*pulumi.OutputState
}

func (HorizontalPodAutoscalerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HorizontalPodAutoscaler)(nil))
}

func (o HorizontalPodAutoscalerOutput) ToHorizontalPodAutoscalerOutput() HorizontalPodAutoscalerOutput {
	return o
}

func (o HorizontalPodAutoscalerOutput) ToHorizontalPodAutoscalerOutputWithContext(ctx context.Context) HorizontalPodAutoscalerOutput {
	return o
}

func (o HorizontalPodAutoscalerOutput) ToHorizontalPodAutoscalerPtrOutput() HorizontalPodAutoscalerPtrOutput {
	return o.ToHorizontalPodAutoscalerPtrOutputWithContext(context.Background())
}

func (o HorizontalPodAutoscalerOutput) ToHorizontalPodAutoscalerPtrOutputWithContext(ctx context.Context) HorizontalPodAutoscalerPtrOutput {
	return o.ApplyT(func(v HorizontalPodAutoscaler) *HorizontalPodAutoscaler {
		return &v
	}).(HorizontalPodAutoscalerPtrOutput)
}

type HorizontalPodAutoscalerPtrOutput struct {
	*pulumi.OutputState
}

func (HorizontalPodAutoscalerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HorizontalPodAutoscaler)(nil))
}

func (o HorizontalPodAutoscalerPtrOutput) ToHorizontalPodAutoscalerPtrOutput() HorizontalPodAutoscalerPtrOutput {
	return o
}

func (o HorizontalPodAutoscalerPtrOutput) ToHorizontalPodAutoscalerPtrOutputWithContext(ctx context.Context) HorizontalPodAutoscalerPtrOutput {
	return o
}

type HorizontalPodAutoscalerArrayOutput struct{ *pulumi.OutputState }

func (HorizontalPodAutoscalerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HorizontalPodAutoscaler)(nil))
}

func (o HorizontalPodAutoscalerArrayOutput) ToHorizontalPodAutoscalerArrayOutput() HorizontalPodAutoscalerArrayOutput {
	return o
}

func (o HorizontalPodAutoscalerArrayOutput) ToHorizontalPodAutoscalerArrayOutputWithContext(ctx context.Context) HorizontalPodAutoscalerArrayOutput {
	return o
}

func (o HorizontalPodAutoscalerArrayOutput) Index(i pulumi.IntInput) HorizontalPodAutoscalerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HorizontalPodAutoscaler {
		return vs[0].([]HorizontalPodAutoscaler)[vs[1].(int)]
	}).(HorizontalPodAutoscalerOutput)
}

type HorizontalPodAutoscalerMapOutput struct{ *pulumi.OutputState }

func (HorizontalPodAutoscalerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]HorizontalPodAutoscaler)(nil))
}

func (o HorizontalPodAutoscalerMapOutput) ToHorizontalPodAutoscalerMapOutput() HorizontalPodAutoscalerMapOutput {
	return o
}

func (o HorizontalPodAutoscalerMapOutput) ToHorizontalPodAutoscalerMapOutputWithContext(ctx context.Context) HorizontalPodAutoscalerMapOutput {
	return o
}

func (o HorizontalPodAutoscalerMapOutput) MapIndex(k pulumi.StringInput) HorizontalPodAutoscalerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) HorizontalPodAutoscaler {
		return vs[0].(map[string]HorizontalPodAutoscaler)[vs[1].(string)]
	}).(HorizontalPodAutoscalerOutput)
}

func init() {
	pulumi.RegisterOutputType(HorizontalPodAutoscalerOutput{})
	pulumi.RegisterOutputType(HorizontalPodAutoscalerPtrOutput{})
	pulumi.RegisterOutputType(HorizontalPodAutoscalerArrayOutput{})
	pulumi.RegisterOutputType(HorizontalPodAutoscalerMapOutput{})
}
