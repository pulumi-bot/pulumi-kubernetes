// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ValidatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and object without changing it.
type ValidatingWebhookConfiguration struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks ValidatingWebhookArrayOutput `pulumi:"webhooks"`
}

// NewValidatingWebhookConfiguration registers a new resource with the given unique name, arguments, and options.
func NewValidatingWebhookConfiguration(ctx *pulumi.Context,
	name string, args *ValidatingWebhookConfigurationArgs, opts ...pulumi.ResourceOption) (*ValidatingWebhookConfiguration, error) {
	if args == nil {
		args = &ValidatingWebhookConfigurationArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("admissionregistration.k8s.io/v1")
	args.Kind = pulumi.StringPtr("ValidatingWebhookConfiguration")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingWebhookConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource ValidatingWebhookConfiguration
	err := ctx.RegisterResource("kubernetes:admissionregistration.k8s.io/v1:ValidatingWebhookConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetValidatingWebhookConfiguration gets an existing ValidatingWebhookConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetValidatingWebhookConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ValidatingWebhookConfigurationState, opts ...pulumi.ResourceOption) (*ValidatingWebhookConfiguration, error) {
	var resource ValidatingWebhookConfiguration
	err := ctx.ReadResource("kubernetes:admissionregistration.k8s.io/v1:ValidatingWebhookConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ValidatingWebhookConfiguration resources.
type validatingWebhookConfigurationState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks []ValidatingWebhook `pulumi:"webhooks"`
}

type ValidatingWebhookConfigurationState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata metav1.ObjectMetaPtrInput
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks ValidatingWebhookArrayInput
}

func (ValidatingWebhookConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*validatingWebhookConfigurationState)(nil)).Elem()
}

type validatingWebhookConfigurationArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks []ValidatingWebhook `pulumi:"webhooks"`
}

// The set of arguments for constructing a ValidatingWebhookConfiguration resource.
type ValidatingWebhookConfigurationArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata metav1.ObjectMetaPtrInput
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks ValidatingWebhookArrayInput
}

func (ValidatingWebhookConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*validatingWebhookConfigurationArgs)(nil)).Elem()
}

type ValidatingWebhookConfigurationInput interface {
	pulumi.Input

	ToValidatingWebhookConfigurationOutput() ValidatingWebhookConfigurationOutput
	ToValidatingWebhookConfigurationOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationOutput
}

func (*ValidatingWebhookConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*ValidatingWebhookConfiguration)(nil))
}

func (i *ValidatingWebhookConfiguration) ToValidatingWebhookConfigurationOutput() ValidatingWebhookConfigurationOutput {
	return i.ToValidatingWebhookConfigurationOutputWithContext(context.Background())
}

func (i *ValidatingWebhookConfiguration) ToValidatingWebhookConfigurationOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ValidatingWebhookConfigurationOutput)
}

type ValidatingWebhookConfigurationOutput struct {
	*pulumi.OutputState
}

func (ValidatingWebhookConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ValidatingWebhookConfiguration)(nil))
}

func (o ValidatingWebhookConfigurationOutput) ToValidatingWebhookConfigurationOutput() ValidatingWebhookConfigurationOutput {
	return o
}

func (o ValidatingWebhookConfigurationOutput) ToValidatingWebhookConfigurationOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ValidatingWebhookConfigurationOutput{})
}
