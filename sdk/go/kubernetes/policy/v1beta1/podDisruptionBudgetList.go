// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// PodDisruptionBudgetList is a collection of PodDisruptionBudgets.
type PodDisruptionBudgetList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput             `pulumi:"apiVersion"`
	Items      PodDisruptionBudgetTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput   `pulumi:"kind"`
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewPodDisruptionBudgetList registers a new resource with the given unique name, arguments, and options.
func NewPodDisruptionBudgetList(ctx *pulumi.Context,
	name string, args *PodDisruptionBudgetListArgs, opts ...pulumi.ResourceOption) (*PodDisruptionBudgetList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("policy/v1beta1")
	args.Kind = pulumi.StringPtr("PodDisruptionBudgetList")
	var resource PodDisruptionBudgetList
	err := ctx.RegisterResource("kubernetes:policy/v1beta1:PodDisruptionBudgetList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPodDisruptionBudgetList gets an existing PodDisruptionBudgetList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPodDisruptionBudgetList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PodDisruptionBudgetListState, opts ...pulumi.ResourceOption) (*PodDisruptionBudgetList, error) {
	var resource PodDisruptionBudgetList
	err := ctx.ReadResource("kubernetes:policy/v1beta1:PodDisruptionBudgetList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PodDisruptionBudgetList resources.
type podDisruptionBudgetListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string                   `pulumi:"apiVersion"`
	Items      []PodDisruptionBudgetType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string          `pulumi:"kind"`
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type PodDisruptionBudgetListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	Items      PodDisruptionBudgetTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ListMetaPtrInput
}

func (PodDisruptionBudgetListState) ElementType() reflect.Type {
	return reflect.TypeOf((*podDisruptionBudgetListState)(nil)).Elem()
}

type podDisruptionBudgetListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string                   `pulumi:"apiVersion"`
	Items      []PodDisruptionBudgetType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string          `pulumi:"kind"`
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a PodDisruptionBudgetList resource.
type PodDisruptionBudgetListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	Items      PodDisruptionBudgetTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ListMetaPtrInput
}

func (PodDisruptionBudgetListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*podDisruptionBudgetListArgs)(nil)).Elem()
}

type PodDisruptionBudgetListInput interface {
	pulumi.Input

	ToPodDisruptionBudgetListOutput() PodDisruptionBudgetListOutput
	ToPodDisruptionBudgetListOutputWithContext(ctx context.Context) PodDisruptionBudgetListOutput
}

func (*PodDisruptionBudgetList) ElementType() reflect.Type {
	return reflect.TypeOf((*PodDisruptionBudgetList)(nil))
}

func (i *PodDisruptionBudgetList) ToPodDisruptionBudgetListOutput() PodDisruptionBudgetListOutput {
	return i.ToPodDisruptionBudgetListOutputWithContext(context.Background())
}

func (i *PodDisruptionBudgetList) ToPodDisruptionBudgetListOutputWithContext(ctx context.Context) PodDisruptionBudgetListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodDisruptionBudgetListOutput)
}

func (i *PodDisruptionBudgetList) ToPodDisruptionBudgetListPtrOutput() PodDisruptionBudgetListPtrOutput {
	return i.ToPodDisruptionBudgetListPtrOutputWithContext(context.Background())
}

func (i *PodDisruptionBudgetList) ToPodDisruptionBudgetListPtrOutputWithContext(ctx context.Context) PodDisruptionBudgetListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodDisruptionBudgetListPtrOutput)
}

type PodDisruptionBudgetListPtrInput interface {
	pulumi.Input

	ToPodDisruptionBudgetListPtrOutput() PodDisruptionBudgetListPtrOutput
	ToPodDisruptionBudgetListPtrOutputWithContext(ctx context.Context) PodDisruptionBudgetListPtrOutput
}

type podDisruptionBudgetListPtrType PodDisruptionBudgetListArgs

func (*podDisruptionBudgetListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PodDisruptionBudgetList)(nil))
}

func (i *podDisruptionBudgetListPtrType) ToPodDisruptionBudgetListPtrOutput() PodDisruptionBudgetListPtrOutput {
	return i.ToPodDisruptionBudgetListPtrOutputWithContext(context.Background())
}

func (i *podDisruptionBudgetListPtrType) ToPodDisruptionBudgetListPtrOutputWithContext(ctx context.Context) PodDisruptionBudgetListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodDisruptionBudgetListOutput).ToPodDisruptionBudgetListPtrOutput()
}

// PodDisruptionBudgetListArrayInput is an input type that accepts PodDisruptionBudgetListArray and PodDisruptionBudgetListArrayOutput values.
// You can construct a concrete instance of `PodDisruptionBudgetListArrayInput` via:
//
//          PodDisruptionBudgetListArray{ PodDisruptionBudgetListArgs{...} }
type PodDisruptionBudgetListArrayInput interface {
	pulumi.Input

	ToPodDisruptionBudgetListArrayOutput() PodDisruptionBudgetListArrayOutput
	ToPodDisruptionBudgetListArrayOutputWithContext(context.Context) PodDisruptionBudgetListArrayOutput
}

type PodDisruptionBudgetListArray []PodDisruptionBudgetListInput

func (PodDisruptionBudgetListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PodDisruptionBudgetList)(nil))
}

func (i PodDisruptionBudgetListArray) ToPodDisruptionBudgetListArrayOutput() PodDisruptionBudgetListArrayOutput {
	return i.ToPodDisruptionBudgetListArrayOutputWithContext(context.Background())
}

func (i PodDisruptionBudgetListArray) ToPodDisruptionBudgetListArrayOutputWithContext(ctx context.Context) PodDisruptionBudgetListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodDisruptionBudgetListArrayOutput)
}

// PodDisruptionBudgetListMapInput is an input type that accepts PodDisruptionBudgetListMap and PodDisruptionBudgetListMapOutput values.
// You can construct a concrete instance of `PodDisruptionBudgetListMapInput` via:
//
//          PodDisruptionBudgetListMap{ "key": PodDisruptionBudgetListArgs{...} }
type PodDisruptionBudgetListMapInput interface {
	pulumi.Input

	ToPodDisruptionBudgetListMapOutput() PodDisruptionBudgetListMapOutput
	ToPodDisruptionBudgetListMapOutputWithContext(context.Context) PodDisruptionBudgetListMapOutput
}

type PodDisruptionBudgetListMap map[string]PodDisruptionBudgetListInput

func (PodDisruptionBudgetListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PodDisruptionBudgetList)(nil))
}

func (i PodDisruptionBudgetListMap) ToPodDisruptionBudgetListMapOutput() PodDisruptionBudgetListMapOutput {
	return i.ToPodDisruptionBudgetListMapOutputWithContext(context.Background())
}

func (i PodDisruptionBudgetListMap) ToPodDisruptionBudgetListMapOutputWithContext(ctx context.Context) PodDisruptionBudgetListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodDisruptionBudgetListMapOutput)
}

type PodDisruptionBudgetListOutput struct {
	*pulumi.OutputState
}

func (PodDisruptionBudgetListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PodDisruptionBudgetList)(nil))
}

func (o PodDisruptionBudgetListOutput) ToPodDisruptionBudgetListOutput() PodDisruptionBudgetListOutput {
	return o
}

func (o PodDisruptionBudgetListOutput) ToPodDisruptionBudgetListOutputWithContext(ctx context.Context) PodDisruptionBudgetListOutput {
	return o
}

func (o PodDisruptionBudgetListOutput) ToPodDisruptionBudgetListPtrOutput() PodDisruptionBudgetListPtrOutput {
	return o.ToPodDisruptionBudgetListPtrOutputWithContext(context.Background())
}

func (o PodDisruptionBudgetListOutput) ToPodDisruptionBudgetListPtrOutputWithContext(ctx context.Context) PodDisruptionBudgetListPtrOutput {
	return o.ApplyT(func(v PodDisruptionBudgetList) *PodDisruptionBudgetList {
		return &v
	}).(PodDisruptionBudgetListPtrOutput)
}

type PodDisruptionBudgetListPtrOutput struct {
	*pulumi.OutputState
}

func (PodDisruptionBudgetListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PodDisruptionBudgetList)(nil))
}

func (o PodDisruptionBudgetListPtrOutput) ToPodDisruptionBudgetListPtrOutput() PodDisruptionBudgetListPtrOutput {
	return o
}

func (o PodDisruptionBudgetListPtrOutput) ToPodDisruptionBudgetListPtrOutputWithContext(ctx context.Context) PodDisruptionBudgetListPtrOutput {
	return o
}

type PodDisruptionBudgetListArrayOutput struct{ *pulumi.OutputState }

func (PodDisruptionBudgetListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PodDisruptionBudgetList)(nil))
}

func (o PodDisruptionBudgetListArrayOutput) ToPodDisruptionBudgetListArrayOutput() PodDisruptionBudgetListArrayOutput {
	return o
}

func (o PodDisruptionBudgetListArrayOutput) ToPodDisruptionBudgetListArrayOutputWithContext(ctx context.Context) PodDisruptionBudgetListArrayOutput {
	return o
}

func (o PodDisruptionBudgetListArrayOutput) Index(i pulumi.IntInput) PodDisruptionBudgetListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PodDisruptionBudgetList {
		return vs[0].([]PodDisruptionBudgetList)[vs[1].(int)]
	}).(PodDisruptionBudgetListOutput)
}

type PodDisruptionBudgetListMapOutput struct{ *pulumi.OutputState }

func (PodDisruptionBudgetListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PodDisruptionBudgetList)(nil))
}

func (o PodDisruptionBudgetListMapOutput) ToPodDisruptionBudgetListMapOutput() PodDisruptionBudgetListMapOutput {
	return o
}

func (o PodDisruptionBudgetListMapOutput) ToPodDisruptionBudgetListMapOutputWithContext(ctx context.Context) PodDisruptionBudgetListMapOutput {
	return o
}

func (o PodDisruptionBudgetListMapOutput) MapIndex(k pulumi.StringInput) PodDisruptionBudgetListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PodDisruptionBudgetList {
		return vs[0].(map[string]PodDisruptionBudgetList)[vs[1].(string)]
	}).(PodDisruptionBudgetListOutput)
}

func init() {
	pulumi.RegisterOutputType(PodDisruptionBudgetListOutput{})
	pulumi.RegisterOutputType(PodDisruptionBudgetListPtrOutput{})
	pulumi.RegisterOutputType(PodDisruptionBudgetListArrayOutput{})
	pulumi.RegisterOutputType(PodDisruptionBudgetListMapOutput{})
}
