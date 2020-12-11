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

// list of horizontal pod autoscaler objects.
type HorizontalPodAutoscalerList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// list of horizontal pod autoscaler objects.
	Items HorizontalPodAutoscalerTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata.
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewHorizontalPodAutoscalerList registers a new resource with the given unique name, arguments, and options.
func NewHorizontalPodAutoscalerList(ctx *pulumi.Context,
	name string, args *HorizontalPodAutoscalerListArgs, opts ...pulumi.ResourceOption) (*HorizontalPodAutoscalerList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("autoscaling/v1")
	args.Kind = pulumi.StringPtr("HorizontalPodAutoscalerList")
	var resource HorizontalPodAutoscalerList
	err := ctx.RegisterResource("kubernetes:autoscaling/v1:HorizontalPodAutoscalerList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHorizontalPodAutoscalerList gets an existing HorizontalPodAutoscalerList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHorizontalPodAutoscalerList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HorizontalPodAutoscalerListState, opts ...pulumi.ResourceOption) (*HorizontalPodAutoscalerList, error) {
	var resource HorizontalPodAutoscalerList
	err := ctx.ReadResource("kubernetes:autoscaling/v1:HorizontalPodAutoscalerList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HorizontalPodAutoscalerList resources.
type horizontalPodAutoscalerListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// list of horizontal pod autoscaler objects.
	Items []HorizontalPodAutoscalerType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata.
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type HorizontalPodAutoscalerListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// list of horizontal pod autoscaler objects.
	Items HorizontalPodAutoscalerTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata.
	Metadata metav1.ListMetaPtrInput
}

func (HorizontalPodAutoscalerListState) ElementType() reflect.Type {
	return reflect.TypeOf((*horizontalPodAutoscalerListState)(nil)).Elem()
}

type horizontalPodAutoscalerListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// list of horizontal pod autoscaler objects.
	Items []HorizontalPodAutoscalerType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata.
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a HorizontalPodAutoscalerList resource.
type HorizontalPodAutoscalerListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// list of horizontal pod autoscaler objects.
	Items HorizontalPodAutoscalerTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata.
	Metadata metav1.ListMetaPtrInput
}

func (HorizontalPodAutoscalerListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*horizontalPodAutoscalerListArgs)(nil)).Elem()
}

type HorizontalPodAutoscalerListInput interface {
	pulumi.Input

	ToHorizontalPodAutoscalerListOutput() HorizontalPodAutoscalerListOutput
	ToHorizontalPodAutoscalerListOutputWithContext(ctx context.Context) HorizontalPodAutoscalerListOutput
}

type HorizontalPodAutoscalerListPtrInput interface {
	pulumi.Input

	ToHorizontalPodAutoscalerListPtrOutput() HorizontalPodAutoscalerListPtrOutput
	ToHorizontalPodAutoscalerListPtrOutputWithContext(ctx context.Context) HorizontalPodAutoscalerListPtrOutput
}

func (HorizontalPodAutoscalerList) ElementType() reflect.Type {
	return reflect.TypeOf((*HorizontalPodAutoscalerList)(nil)).Elem()
}

func (i HorizontalPodAutoscalerList) ToHorizontalPodAutoscalerListOutput() HorizontalPodAutoscalerListOutput {
	return i.ToHorizontalPodAutoscalerListOutputWithContext(context.Background())
}

func (i HorizontalPodAutoscalerList) ToHorizontalPodAutoscalerListOutputWithContext(ctx context.Context) HorizontalPodAutoscalerListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HorizontalPodAutoscalerListOutput)
}

func (i HorizontalPodAutoscalerList) ToHorizontalPodAutoscalerListPtrOutput() HorizontalPodAutoscalerListPtrOutput {
	return i.ToHorizontalPodAutoscalerListPtrOutputWithContext(context.Background())
}

func (i HorizontalPodAutoscalerList) ToHorizontalPodAutoscalerListPtrOutputWithContext(ctx context.Context) HorizontalPodAutoscalerListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HorizontalPodAutoscalerListPtrOutput)
}

type HorizontalPodAutoscalerListOutput struct {
	*pulumi.OutputState
}

func (HorizontalPodAutoscalerListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HorizontalPodAutoscalerListOutput)(nil)).Elem()
}

func (o HorizontalPodAutoscalerListOutput) ToHorizontalPodAutoscalerListOutput() HorizontalPodAutoscalerListOutput {
	return o
}

func (o HorizontalPodAutoscalerListOutput) ToHorizontalPodAutoscalerListOutputWithContext(ctx context.Context) HorizontalPodAutoscalerListOutput {
	return o
}

type HorizontalPodAutoscalerListPtrOutput struct {
	*pulumi.OutputState
}

func (HorizontalPodAutoscalerListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HorizontalPodAutoscalerList)(nil)).Elem()
}

func (o HorizontalPodAutoscalerListPtrOutput) ToHorizontalPodAutoscalerListPtrOutput() HorizontalPodAutoscalerListPtrOutput {
	return o
}

func (o HorizontalPodAutoscalerListPtrOutput) ToHorizontalPodAutoscalerListPtrOutputWithContext(ctx context.Context) HorizontalPodAutoscalerListPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HorizontalPodAutoscalerListOutput{})
	pulumi.RegisterOutputType(HorizontalPodAutoscalerListPtrOutput{})
}
