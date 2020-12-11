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

// ValidatingWebhookConfigurationList is a list of ValidatingWebhookConfiguration.
type ValidatingWebhookConfigurationList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// List of ValidatingWebhookConfiguration.
	Items ValidatingWebhookConfigurationTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewValidatingWebhookConfigurationList registers a new resource with the given unique name, arguments, and options.
func NewValidatingWebhookConfigurationList(ctx *pulumi.Context,
	name string, args *ValidatingWebhookConfigurationListArgs, opts ...pulumi.ResourceOption) (*ValidatingWebhookConfigurationList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("admissionregistration.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("ValidatingWebhookConfigurationList")
	var resource ValidatingWebhookConfigurationList
	err := ctx.RegisterResource("kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingWebhookConfigurationList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetValidatingWebhookConfigurationList gets an existing ValidatingWebhookConfigurationList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetValidatingWebhookConfigurationList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ValidatingWebhookConfigurationListState, opts ...pulumi.ResourceOption) (*ValidatingWebhookConfigurationList, error) {
	var resource ValidatingWebhookConfigurationList
	err := ctx.ReadResource("kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingWebhookConfigurationList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ValidatingWebhookConfigurationList resources.
type validatingWebhookConfigurationListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of ValidatingWebhookConfiguration.
	Items []ValidatingWebhookConfigurationType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type ValidatingWebhookConfigurationListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of ValidatingWebhookConfiguration.
	Items ValidatingWebhookConfigurationTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ValidatingWebhookConfigurationListState) ElementType() reflect.Type {
	return reflect.TypeOf((*validatingWebhookConfigurationListState)(nil)).Elem()
}

type validatingWebhookConfigurationListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of ValidatingWebhookConfiguration.
	Items []ValidatingWebhookConfigurationType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ValidatingWebhookConfigurationList resource.
type ValidatingWebhookConfigurationListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of ValidatingWebhookConfiguration.
	Items ValidatingWebhookConfigurationTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ValidatingWebhookConfigurationListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*validatingWebhookConfigurationListArgs)(nil)).Elem()
}

type ValidatingWebhookConfigurationListInput interface {
	pulumi.Input

	ToValidatingWebhookConfigurationListOutput() ValidatingWebhookConfigurationListOutput
	ToValidatingWebhookConfigurationListOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationListOutput
}

func (*ValidatingWebhookConfigurationList) ElementType() reflect.Type {
	return reflect.TypeOf((*ValidatingWebhookConfigurationList)(nil))
}

func (i *ValidatingWebhookConfigurationList) ToValidatingWebhookConfigurationListOutput() ValidatingWebhookConfigurationListOutput {
	return i.ToValidatingWebhookConfigurationListOutputWithContext(context.Background())
}

func (i *ValidatingWebhookConfigurationList) ToValidatingWebhookConfigurationListOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ValidatingWebhookConfigurationListOutput)
}

type ValidatingWebhookConfigurationListOutput struct {
	*pulumi.OutputState
}

func (ValidatingWebhookConfigurationListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ValidatingWebhookConfigurationList)(nil))
}

func (o ValidatingWebhookConfigurationListOutput) ToValidatingWebhookConfigurationListOutput() ValidatingWebhookConfigurationListOutput {
	return o
}

func (o ValidatingWebhookConfigurationListOutput) ToValidatingWebhookConfigurationListOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationListOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ValidatingWebhookConfigurationListOutput{})
}
