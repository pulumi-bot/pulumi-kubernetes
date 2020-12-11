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

// StatefulSetList is a collection of StatefulSets.
type StatefulSetList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Items      StatefulSetTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput   `pulumi:"kind"`
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewStatefulSetList registers a new resource with the given unique name, arguments, and options.
func NewStatefulSetList(ctx *pulumi.Context,
	name string, args *StatefulSetListArgs, opts ...pulumi.ResourceOption) (*StatefulSetList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("apps/v1")
	args.Kind = pulumi.StringPtr("StatefulSetList")
	var resource StatefulSetList
	err := ctx.RegisterResource("kubernetes:apps/v1:StatefulSetList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStatefulSetList gets an existing StatefulSetList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStatefulSetList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StatefulSetListState, opts ...pulumi.ResourceOption) (*StatefulSetList, error) {
	var resource StatefulSetList
	err := ctx.ReadResource("kubernetes:apps/v1:StatefulSetList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StatefulSetList resources.
type statefulSetListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string           `pulumi:"apiVersion"`
	Items      []StatefulSetType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string          `pulumi:"kind"`
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type StatefulSetListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	Items      StatefulSetTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ListMetaPtrInput
}

func (StatefulSetListState) ElementType() reflect.Type {
	return reflect.TypeOf((*statefulSetListState)(nil)).Elem()
}

type statefulSetListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string           `pulumi:"apiVersion"`
	Items      []StatefulSetType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string          `pulumi:"kind"`
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a StatefulSetList resource.
type StatefulSetListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	Items      StatefulSetTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ListMetaPtrInput
}

func (StatefulSetListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*statefulSetListArgs)(nil)).Elem()
}

type StatefulSetListInput interface {
	pulumi.Input

	ToStatefulSetListOutput() StatefulSetListOutput
	ToStatefulSetListOutputWithContext(ctx context.Context) StatefulSetListOutput
}

func (*StatefulSetList) ElementType() reflect.Type {
	return reflect.TypeOf((*StatefulSetList)(nil))
}

func (i *StatefulSetList) ToStatefulSetListOutput() StatefulSetListOutput {
	return i.ToStatefulSetListOutputWithContext(context.Background())
}

func (i *StatefulSetList) ToStatefulSetListOutputWithContext(ctx context.Context) StatefulSetListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatefulSetListOutput)
}

func (i *StatefulSetList) ToStatefulSetListPtrOutput() StatefulSetListPtrOutput {
	return i.ToStatefulSetListPtrOutputWithContext(context.Background())
}

func (i *StatefulSetList) ToStatefulSetListPtrOutputWithContext(ctx context.Context) StatefulSetListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatefulSetListPtrOutput)
}

type StatefulSetListPtrInput interface {
	pulumi.Input

	ToStatefulSetListPtrOutput() StatefulSetListPtrOutput
	ToStatefulSetListPtrOutputWithContext(ctx context.Context) StatefulSetListPtrOutput
}

type StatefulSetListOutput struct {
	*pulumi.OutputState
}

func (StatefulSetListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatefulSetList)(nil))
}

func (o StatefulSetListOutput) ToStatefulSetListOutput() StatefulSetListOutput {
	return o
}

func (o StatefulSetListOutput) ToStatefulSetListOutputWithContext(ctx context.Context) StatefulSetListOutput {
	return o
}

type StatefulSetListPtrOutput struct {
	*pulumi.OutputState
}

func (StatefulSetListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StatefulSetList)(nil))
}

func (o StatefulSetListPtrOutput) ToStatefulSetListPtrOutput() StatefulSetListPtrOutput {
	return o
}

func (o StatefulSetListPtrOutput) ToStatefulSetListPtrOutputWithContext(ctx context.Context) StatefulSetListPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StatefulSetListOutput{})
	pulumi.RegisterOutputType(StatefulSetListPtrOutput{})
}
