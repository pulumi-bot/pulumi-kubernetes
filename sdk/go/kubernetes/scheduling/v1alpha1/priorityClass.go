// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// DEPRECATED - This group version of PriorityClass is deprecated by scheduling.k8s.io/v1/PriorityClass. PriorityClass defines mapping from a priority class name to the priority integer value. The value can be any valid integer.
type PriorityClass struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// description is an arbitrary string that usually provides guidelines on when this priority class should be used.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
	GlobalDefault pulumi.BoolPtrOutput `pulumi:"globalDefault"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is beta-level, gated by the NonPreemptingPriority feature-gate.
	PreemptionPolicy pulumi.StringPtrOutput `pulumi:"preemptionPolicy"`
	// The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
	Value pulumi.IntOutput `pulumi:"value"`
}

// NewPriorityClass registers a new resource with the given unique name, arguments, and options.
func NewPriorityClass(ctx *pulumi.Context,
	name string, args *PriorityClassArgs, opts ...pulumi.ResourceOption) (*PriorityClass, error) {
	if args == nil || args.Value == nil {
		return nil, errors.New("missing required argument 'Value'")
	}
	if args == nil {
		args = &PriorityClassArgs{}
	}
	args.ApiVersion = pulumi.StringPtr("scheduling.k8s.io/v1alpha1")
	args.Kind = pulumi.StringPtr("PriorityClass")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:scheduling.k8s.io/v1:PriorityClass"),
		},
		{
			Type: pulumi.String("kubernetes:scheduling.k8s.io/v1beta1:PriorityClass"),
		},
	})
	opts = append(opts, aliases)
	var resource PriorityClass
	err := ctx.RegisterResource("kubernetes:scheduling.k8s.io/v1alpha1:PriorityClass", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPriorityClass gets an existing PriorityClass resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPriorityClass(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PriorityClassState, opts ...pulumi.ResourceOption) (*PriorityClass, error) {
	var resource PriorityClass
	err := ctx.ReadResource("kubernetes:scheduling.k8s.io/v1alpha1:PriorityClass", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PriorityClass resources.
type priorityClassState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// description is an arbitrary string that usually provides guidelines on when this priority class should be used.
	Description *string `pulumi:"description"`
	// globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
	GlobalDefault *bool `pulumi:"globalDefault"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is beta-level, gated by the NonPreemptingPriority feature-gate.
	PreemptionPolicy *string `pulumi:"preemptionPolicy"`
	// The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
	Value *int `pulumi:"value"`
}

type PriorityClassState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// description is an arbitrary string that usually provides guidelines on when this priority class should be used.
	Description pulumi.StringPtrInput
	// globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
	GlobalDefault pulumi.BoolPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is beta-level, gated by the NonPreemptingPriority feature-gate.
	PreemptionPolicy pulumi.StringPtrInput
	// The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
	Value pulumi.IntPtrInput
}

func (PriorityClassState) ElementType() reflect.Type {
	return reflect.TypeOf((*priorityClassState)(nil)).Elem()
}

type priorityClassArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// description is an arbitrary string that usually provides guidelines on when this priority class should be used.
	Description *string `pulumi:"description"`
	// globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
	GlobalDefault *bool `pulumi:"globalDefault"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is beta-level, gated by the NonPreemptingPriority feature-gate.
	PreemptionPolicy *string `pulumi:"preemptionPolicy"`
	// The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
	Value int `pulumi:"value"`
}

// The set of arguments for constructing a PriorityClass resource.
type PriorityClassArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// description is an arbitrary string that usually provides guidelines on when this priority class should be used.
	Description pulumi.StringPtrInput
	// globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
	GlobalDefault pulumi.BoolPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is beta-level, gated by the NonPreemptingPriority feature-gate.
	PreemptionPolicy pulumi.StringPtrInput
	// The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
	Value pulumi.IntInput
}

func (PriorityClassArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*priorityClassArgs)(nil)).Elem()
}

type PriorityClassInput interface {
	pulumi.Input

	ToPriorityClassOutput() PriorityClassOutput
	ToPriorityClassOutputWithContext(ctx context.Context) PriorityClassOutput
}

func (PriorityClass) ElementType() reflect.Type {
	return reflect.TypeOf((*PriorityClass)(nil)).Elem()
}

func (i PriorityClass) ToPriorityClassOutput() PriorityClassOutput {
	return i.ToPriorityClassOutputWithContext(context.Background())
}

func (i PriorityClass) ToPriorityClassOutputWithContext(ctx context.Context) PriorityClassOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PriorityClassOutput)
}

type PriorityClassOutput struct {
	*pulumi.OutputState
}

func (PriorityClassOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PriorityClassOutput)(nil)).Elem()
}

func (o PriorityClassOutput) ToPriorityClassOutput() PriorityClassOutput {
	return o
}

func (o PriorityClassOutput) ToPriorityClassOutputWithContext(ctx context.Context) PriorityClassOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PriorityClassOutput{})
}
