// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"fmt"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// MutatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and may change the object. Deprecated in v1.16, planned for removal in v1.19. Use admissionregistration.k8s.io/v1 MutatingWebhookConfiguration instead.
type MutatingWebhookConfiguration struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks MutatingWebhookArrayOutput `pulumi:"webhooks"`
}

// NewMutatingWebhookConfiguration registers a new resource with the given unique name, arguments, and options.
func NewMutatingWebhookConfiguration(ctx *pulumi.Context,
	name string, args *MutatingWebhookConfigurationArgs, opts ...pulumi.ResourceOption) (*MutatingWebhookConfiguration, error) {
	if args == nil {
		args = &MutatingWebhookConfigurationArgs{}
	}
	args.ApiVersion = pulumi.StringPtr("admissionregistration.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("MutatingWebhookConfiguration")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:admissionregistration.k8s.io/v1:MutatingWebhookConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource MutatingWebhookConfiguration
	err := ctx.RegisterResource("kubernetes:admissionregistration.k8s.io/v1beta1:MutatingWebhookConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMutatingWebhookConfiguration gets an existing MutatingWebhookConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMutatingWebhookConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MutatingWebhookConfigurationState, opts ...pulumi.ResourceOption) (*MutatingWebhookConfiguration, error) {
	var resource MutatingWebhookConfiguration
	err := ctx.ReadResource("kubernetes:admissionregistration.k8s.io/v1beta1:MutatingWebhookConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MutatingWebhookConfiguration resources.
type mutatingWebhookConfigurationState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks []MutatingWebhook `pulumi:"webhooks"`
}

type MutatingWebhookConfigurationState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata metav1.ObjectMetaPtrInput
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks MutatingWebhookArrayInput
}

func (MutatingWebhookConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*mutatingWebhookConfigurationState)(nil)).Elem()
}

type mutatingWebhookConfigurationArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks []MutatingWebhook `pulumi:"webhooks"`
}

// The set of arguments for constructing a MutatingWebhookConfiguration resource.
type MutatingWebhookConfigurationArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata metav1.ObjectMetaPtrInput
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks MutatingWebhookArrayInput
}

func (MutatingWebhookConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mutatingWebhookConfigurationArgs)(nil)).Elem()
}

type MutatingWebhookConfigurationInput interface {
	pulumi.Input

	ToMutatingWebhookConfigurationOutput() MutatingWebhookConfigurationOutput
	ToMutatingWebhookConfigurationOutputWithContext(ctx context.Context) MutatingWebhookConfigurationOutput
}

func (MutatingWebhookConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*MutatingWebhookConfiguration)(nil)).Elem()
}

func (i MutatingWebhookConfiguration) ToMutatingWebhookConfigurationOutput() MutatingWebhookConfigurationOutput {
	return i.ToMutatingWebhookConfigurationOutputWithContext(context.Background())
}

func (i MutatingWebhookConfiguration) ToMutatingWebhookConfigurationOutputWithContext(ctx context.Context) MutatingWebhookConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MutatingWebhookConfigurationOutput)
}

type MutatingWebhookConfigurationOutput struct {
	*pulumi.OutputState
}

func (MutatingWebhookConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MutatingWebhookConfigurationOutput)(nil)).Elem()
}

func (o MutatingWebhookConfigurationOutput) ToMutatingWebhookConfigurationOutput() MutatingWebhookConfigurationOutput {
	return o
}

func (o MutatingWebhookConfigurationOutput) ToMutatingWebhookConfigurationOutputWithContext(ctx context.Context) MutatingWebhookConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MutatingWebhookConfigurationOutput{})
}
