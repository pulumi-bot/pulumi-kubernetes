// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// PodPreset is a policy resource that defines additional runtime requirements for a Pod.
type PodPreset struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	Spec     PodPresetSpecPtrOutput     `pulumi:"spec"`
}

// NewPodPreset registers a new resource with the given unique name, arguments, and options.
func NewPodPreset(ctx *pulumi.Context,
	name string, args *PodPresetArgs, opts ...pulumi.ResourceOption) (*PodPreset, error) {
	if args == nil {
		args = &PodPresetArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("settings.k8s.io/v1alpha1")
	args.Kind = pulumi.StringPtr("PodPreset")
	var resource PodPreset
	err := ctx.RegisterResource("kubernetes:settings.k8s.io/v1alpha1:PodPreset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPodPreset gets an existing PodPreset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPodPreset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PodPresetState, opts ...pulumi.ResourceOption) (*PodPreset, error) {
	var resource PodPreset
	err := ctx.ReadResource("kubernetes:settings.k8s.io/v1alpha1:PodPreset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PodPreset resources.
type podPresetState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	Spec     *PodPresetSpec     `pulumi:"spec"`
}

type PodPresetState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	Spec     PodPresetSpecPtrInput
}

func (PodPresetState) ElementType() reflect.Type {
	return reflect.TypeOf((*podPresetState)(nil)).Elem()
}

type podPresetArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	Spec     *PodPresetSpec     `pulumi:"spec"`
}

// The set of arguments for constructing a PodPreset resource.
type PodPresetArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	Spec     PodPresetSpecPtrInput
}

func (PodPresetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*podPresetArgs)(nil)).Elem()
}

type PodPresetInput interface {
	pulumi.Input

	ToPodPresetOutput() PodPresetOutput
	ToPodPresetOutputWithContext(ctx context.Context) PodPresetOutput
}

func (*PodPreset) ElementType() reflect.Type {
	return reflect.TypeOf((*PodPreset)(nil))
}

func (i *PodPreset) ToPodPresetOutput() PodPresetOutput {
	return i.ToPodPresetOutputWithContext(context.Background())
}

func (i *PodPreset) ToPodPresetOutputWithContext(ctx context.Context) PodPresetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodPresetOutput)
}

func (i *PodPreset) ToPodPresetPtrOutput() PodPresetPtrOutput {
	return i.ToPodPresetPtrOutputWithContext(context.Background())
}

func (i *PodPreset) ToPodPresetPtrOutputWithContext(ctx context.Context) PodPresetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodPresetPtrOutput)
}

type PodPresetPtrInput interface {
	pulumi.Input

	ToPodPresetPtrOutput() PodPresetPtrOutput
	ToPodPresetPtrOutputWithContext(ctx context.Context) PodPresetPtrOutput
}

type PodPresetOutput struct {
	*pulumi.OutputState
}

func (PodPresetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PodPreset)(nil))
}

func (o PodPresetOutput) ToPodPresetOutput() PodPresetOutput {
	return o
}

func (o PodPresetOutput) ToPodPresetOutputWithContext(ctx context.Context) PodPresetOutput {
	return o
}

type PodPresetPtrOutput struct {
	*pulumi.OutputState
}

func (PodPresetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PodPreset)(nil))
}

func (o PodPresetPtrOutput) ToPodPresetPtrOutput() PodPresetPtrOutput {
	return o
}

func (o PodPresetPtrOutput) ToPodPresetPtrOutputWithContext(ctx context.Context) PodPresetPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PodPresetOutput{})
	pulumi.RegisterOutputType(PodPresetPtrOutput{})
}
