// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta2

import (
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// HorizontalPodAutoscalerList is a list of horizontal pod autoscaler objects.
type HorizontalPodAutoscalerList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// items is the list of horizontal pod autoscaler objects.
	Items HorizontalPodAutoscalerTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// metadata is the standard list metadata.
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewHorizontalPodAutoscalerList registers a new resource with the given unique name, arguments, and options.
func NewHorizontalPodAutoscalerList(ctx *pulumi.Context,
	name string, args *HorizontalPodAutoscalerListArgs, opts ...pulumi.ResourceOption) (*HorizontalPodAutoscalerList, error) {
	if args == nil || args.Items == nil {
		return nil, errors.New("missing required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("autoscaling/v2beta2")
	args.Kind = pulumi.StringPtr("HorizontalPodAutoscalerList")
	var resource HorizontalPodAutoscalerList
	err := ctx.RegisterResource("kubernetes:autoscaling/v2beta2:HorizontalPodAutoscalerList", name, args, &resource, opts...)
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
	err := ctx.ReadResource("kubernetes:autoscaling/v2beta2:HorizontalPodAutoscalerList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HorizontalPodAutoscalerList resources.
type horizontalPodAutoscalerListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is the list of horizontal pod autoscaler objects.
	Items []HorizontalPodAutoscalerType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// metadata is the standard list metadata.
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type HorizontalPodAutoscalerListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is the list of horizontal pod autoscaler objects.
	Items HorizontalPodAutoscalerTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// metadata is the standard list metadata.
	Metadata metav1.ListMetaPtrInput
}

func (HorizontalPodAutoscalerListState) ElementType() reflect.Type {
	return reflect.TypeOf((*horizontalPodAutoscalerListState)(nil)).Elem()
}

type horizontalPodAutoscalerListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is the list of horizontal pod autoscaler objects.
	Items []HorizontalPodAutoscalerType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// metadata is the standard list metadata.
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a HorizontalPodAutoscalerList resource.
type HorizontalPodAutoscalerListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is the list of horizontal pod autoscaler objects.
	Items HorizontalPodAutoscalerTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// metadata is the standard list metadata.
	Metadata metav1.ListMetaPtrInput
}

func (HorizontalPodAutoscalerListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*horizontalPodAutoscalerListArgs)(nil)).Elem()
}
