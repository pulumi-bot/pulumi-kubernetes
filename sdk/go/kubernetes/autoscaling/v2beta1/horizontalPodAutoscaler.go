// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// HorizontalPodAutoscaler is the configuration for a horizontal pod autoscaler, which automatically manages the replica count of any resource implementing the scale subresource based on the metrics specified.
type HorizontalPodAutoscaler struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec HorizontalPodAutoscalerSpecPtrOutput `pulumi:"spec"`
	// status is the current information about the autoscaler.
	Status HorizontalPodAutoscalerStatusPtrOutput `pulumi:"status"`
}

// NewHorizontalPodAutoscaler registers a new resource with the given unique name, arguments, and options.
func NewHorizontalPodAutoscaler(ctx *pulumi.Context,
	name string, args *HorizontalPodAutoscalerArgs, opts ...pulumi.ResourceOption) (*HorizontalPodAutoscaler, error) {
	if args == nil {
		args = &HorizontalPodAutoscalerArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("autoscaling/v2beta1")
	args.Kind = pulumi.StringPtr("HorizontalPodAutoscaler")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:autoscaling/v1:HorizontalPodAutoscaler"),
		},
		{
			Type: pulumi.String("kubernetes:autoscaling/v2beta2:HorizontalPodAutoscaler"),
		},
	})
	opts = append(opts, aliases)
	var resource HorizontalPodAutoscaler
	err := ctx.RegisterResource("kubernetes:autoscaling/v2beta1:HorizontalPodAutoscaler", name, args, &resource, opts...)
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
	err := ctx.ReadResource("kubernetes:autoscaling/v2beta1:HorizontalPodAutoscaler", name, id, state, &resource, opts...)
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
	// metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec *HorizontalPodAutoscalerSpec `pulumi:"spec"`
	// status is the current information about the autoscaler.
	Status *HorizontalPodAutoscalerStatus `pulumi:"status"`
}

type HorizontalPodAutoscalerState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec HorizontalPodAutoscalerSpecPtrInput
	// status is the current information about the autoscaler.
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
	// metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec *HorizontalPodAutoscalerSpec `pulumi:"spec"`
}

// The set of arguments for constructing a HorizontalPodAutoscaler resource.
type HorizontalPodAutoscalerArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
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

type HorizontalPodAutoscalerPtrInput interface {
	pulumi.Input

	ToHorizontalPodAutoscalerPtrOutput() HorizontalPodAutoscalerPtrOutput
	ToHorizontalPodAutoscalerPtrOutputWithContext(ctx context.Context) HorizontalPodAutoscalerPtrOutput
}

func (HorizontalPodAutoscaler) ElementType() reflect.Type {
	return reflect.TypeOf((*HorizontalPodAutoscaler)(nil)).Elem()
}

func (i HorizontalPodAutoscaler) ToHorizontalPodAutoscalerOutput() HorizontalPodAutoscalerOutput {
	return i.ToHorizontalPodAutoscalerOutputWithContext(context.Background())
}

func (i HorizontalPodAutoscaler) ToHorizontalPodAutoscalerOutputWithContext(ctx context.Context) HorizontalPodAutoscalerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HorizontalPodAutoscalerOutput)
}

func (i HorizontalPodAutoscaler) ToHorizontalPodAutoscalerPtrOutput() HorizontalPodAutoscalerPtrOutput {
	return i.ToHorizontalPodAutoscalerPtrOutputWithContext(context.Background())
}

func (i HorizontalPodAutoscaler) ToHorizontalPodAutoscalerPtrOutputWithContext(ctx context.Context) HorizontalPodAutoscalerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HorizontalPodAutoscalerPtrOutput)
}

type HorizontalPodAutoscalerOutput struct {
	*pulumi.OutputState
}

func (HorizontalPodAutoscalerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HorizontalPodAutoscalerOutput)(nil)).Elem()
}

func (o HorizontalPodAutoscalerOutput) ToHorizontalPodAutoscalerOutput() HorizontalPodAutoscalerOutput {
	return o
}

func (o HorizontalPodAutoscalerOutput) ToHorizontalPodAutoscalerOutputWithContext(ctx context.Context) HorizontalPodAutoscalerOutput {
	return o
}

type HorizontalPodAutoscalerPtrOutput struct {
	*pulumi.OutputState
}

func (HorizontalPodAutoscalerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HorizontalPodAutoscaler)(nil)).Elem()
}

func (o HorizontalPodAutoscalerPtrOutput) ToHorizontalPodAutoscalerPtrOutput() HorizontalPodAutoscalerPtrOutput {
	return o
}

func (o HorizontalPodAutoscalerPtrOutput) ToHorizontalPodAutoscalerPtrOutputWithContext(ctx context.Context) HorizontalPodAutoscalerPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HorizontalPodAutoscalerOutput{})
	pulumi.RegisterOutputType(HorizontalPodAutoscalerPtrOutput{})
}
