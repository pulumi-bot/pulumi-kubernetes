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

// EventList is a list of Event objects.
type EventList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// items is a list of schema objects.
	Items EventTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewEventList registers a new resource with the given unique name, arguments, and options.
func NewEventList(ctx *pulumi.Context,
	name string, args *EventListArgs, opts ...pulumi.ResourceOption) (*EventList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("events.k8s.io/v1")
	args.Kind = pulumi.StringPtr("EventList")
	var resource EventList
	err := ctx.RegisterResource("kubernetes:events.k8s.io/v1:EventList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventList gets an existing EventList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventListState, opts ...pulumi.ResourceOption) (*EventList, error) {
	var resource EventList
	err := ctx.ReadResource("kubernetes:events.k8s.io/v1:EventList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventList resources.
type eventListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is a list of schema objects.
	Items []EventType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type EventListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is a list of schema objects.
	Items EventTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (EventListState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventListState)(nil)).Elem()
}

type eventListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is a list of schema objects.
	Items []EventType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a EventList resource.
type EventListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is a list of schema objects.
	Items EventTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (EventListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventListArgs)(nil)).Elem()
}

type EventListInput interface {
	pulumi.Input

	ToEventListOutput() EventListOutput
	ToEventListOutputWithContext(ctx context.Context) EventListOutput
}

type EventListPtrInput interface {
	pulumi.Input

	ToEventListPtrOutput() EventListPtrOutput
	ToEventListPtrOutputWithContext(ctx context.Context) EventListPtrOutput
}

func (EventList) ElementType() reflect.Type {
	return reflect.TypeOf((*EventList)(nil)).Elem()
}

func (i EventList) ToEventListOutput() EventListOutput {
	return i.ToEventListOutputWithContext(context.Background())
}

func (i EventList) ToEventListOutputWithContext(ctx context.Context) EventListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventListOutput)
}

func (i EventList) ToEventListPtrOutput() EventListPtrOutput {
	return i.ToEventListPtrOutputWithContext(context.Background())
}

func (i EventList) ToEventListPtrOutputWithContext(ctx context.Context) EventListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventListPtrOutput)
}

type EventListOutput struct {
	*pulumi.OutputState
}

func (EventListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventListOutput)(nil)).Elem()
}

func (o EventListOutput) ToEventListOutput() EventListOutput {
	return o
}

func (o EventListOutput) ToEventListOutputWithContext(ctx context.Context) EventListOutput {
	return o
}

type EventListPtrOutput struct {
	*pulumi.OutputState
}

func (EventListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventList)(nil)).Elem()
}

func (o EventListPtrOutput) ToEventListPtrOutput() EventListPtrOutput {
	return o
}

func (o EventListPtrOutput) ToEventListPtrOutputWithContext(ctx context.Context) EventListPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EventListOutput{})
	pulumi.RegisterOutputType(EventListPtrOutput{})
}
