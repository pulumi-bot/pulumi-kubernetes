// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// MutatingWebhookConfigurationList is a list of MutatingWebhookConfiguration.
type MutatingWebhookConfigurationList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// List of MutatingWebhookConfiguration.
	Items MutatingWebhookConfigurationTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewMutatingWebhookConfigurationList registers a new resource with the given unique name, arguments, and options.
func NewMutatingWebhookConfigurationList(ctx *pulumi.Context,
	name string, args *MutatingWebhookConfigurationListArgs, opts ...pulumi.ResourceOption) (*MutatingWebhookConfigurationList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("admissionregistration.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("MutatingWebhookConfigurationList")
	var resource MutatingWebhookConfigurationList
	err := ctx.RegisterResource("kubernetes:admissionregistration.k8s.io/v1beta1:MutatingWebhookConfigurationList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMutatingWebhookConfigurationList gets an existing MutatingWebhookConfigurationList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMutatingWebhookConfigurationList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MutatingWebhookConfigurationListState, opts ...pulumi.ResourceOption) (*MutatingWebhookConfigurationList, error) {
	var resource MutatingWebhookConfigurationList
	err := ctx.ReadResource("kubernetes:admissionregistration.k8s.io/v1beta1:MutatingWebhookConfigurationList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MutatingWebhookConfigurationList resources.
type mutatingWebhookConfigurationListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of MutatingWebhookConfiguration.
	Items []MutatingWebhookConfigurationType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type MutatingWebhookConfigurationListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of MutatingWebhookConfiguration.
	Items MutatingWebhookConfigurationTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (MutatingWebhookConfigurationListState) ElementType() reflect.Type {
	return reflect.TypeOf((*mutatingWebhookConfigurationListState)(nil)).Elem()
}

type mutatingWebhookConfigurationListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// List of MutatingWebhookConfiguration.
	Items []MutatingWebhookConfigurationType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a MutatingWebhookConfigurationList resource.
type MutatingWebhookConfigurationListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// List of MutatingWebhookConfiguration.
	Items MutatingWebhookConfigurationTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (MutatingWebhookConfigurationListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mutatingWebhookConfigurationListArgs)(nil)).Elem()
}
