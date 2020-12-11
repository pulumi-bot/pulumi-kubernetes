// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// AuditSink represents a cluster level audit sink
type AuditSink struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec defines the audit configuration spec
	Spec AuditSinkSpecPtrOutput `pulumi:"spec"`
}

// NewAuditSink registers a new resource with the given unique name, arguments, and options.
func NewAuditSink(ctx *pulumi.Context,
	name string, args *AuditSinkArgs, opts ...pulumi.ResourceOption) (*AuditSink, error) {
	if args == nil {
		args = &AuditSinkArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("auditregistration.k8s.io/v1alpha1")
	args.Kind = pulumi.StringPtr("AuditSink")
	var resource AuditSink
	err := ctx.RegisterResource("kubernetes:auditregistration.k8s.io/v1alpha1:AuditSink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuditSink gets an existing AuditSink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuditSink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuditSinkState, opts ...pulumi.ResourceOption) (*AuditSink, error) {
	var resource AuditSink
	err := ctx.ReadResource("kubernetes:auditregistration.k8s.io/v1alpha1:AuditSink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuditSink resources.
type auditSinkState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec defines the audit configuration spec
	Spec *AuditSinkSpec `pulumi:"spec"`
}

type AuditSinkState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Spec defines the audit configuration spec
	Spec AuditSinkSpecPtrInput
}

func (AuditSinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*auditSinkState)(nil)).Elem()
}

type auditSinkArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec defines the audit configuration spec
	Spec *AuditSinkSpec `pulumi:"spec"`
}

// The set of arguments for constructing a AuditSink resource.
type AuditSinkArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Spec defines the audit configuration spec
	Spec AuditSinkSpecPtrInput
}

func (AuditSinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*auditSinkArgs)(nil)).Elem()
}

type AuditSinkInput interface {
	pulumi.Input

	ToAuditSinkOutput() AuditSinkOutput
	ToAuditSinkOutputWithContext(ctx context.Context) AuditSinkOutput
}

func (*AuditSink) ElementType() reflect.Type {
	return reflect.TypeOf((*AuditSink)(nil))
}

func (i *AuditSink) ToAuditSinkOutput() AuditSinkOutput {
	return i.ToAuditSinkOutputWithContext(context.Background())
}

func (i *AuditSink) ToAuditSinkOutputWithContext(ctx context.Context) AuditSinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditSinkOutput)
}

func (i *AuditSink) ToAuditSinkPtrOutput() AuditSinkPtrOutput {
	return i.ToAuditSinkPtrOutputWithContext(context.Background())
}

func (i *AuditSink) ToAuditSinkPtrOutputWithContext(ctx context.Context) AuditSinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditSinkPtrOutput)
}

type AuditSinkPtrInput interface {
	pulumi.Input

	ToAuditSinkPtrOutput() AuditSinkPtrOutput
	ToAuditSinkPtrOutputWithContext(ctx context.Context) AuditSinkPtrOutput
}

type AuditSinkOutput struct {
	*pulumi.OutputState
}

func (AuditSinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuditSink)(nil))
}

func (o AuditSinkOutput) ToAuditSinkOutput() AuditSinkOutput {
	return o
}

func (o AuditSinkOutput) ToAuditSinkOutputWithContext(ctx context.Context) AuditSinkOutput {
	return o
}

type AuditSinkPtrOutput struct {
	*pulumi.OutputState
}

func (AuditSinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuditSink)(nil))
}

func (o AuditSinkPtrOutput) ToAuditSinkPtrOutput() AuditSinkPtrOutput {
	return o
}

func (o AuditSinkPtrOutput) ToAuditSinkPtrOutputWithContext(ctx context.Context) AuditSinkPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AuditSinkOutput{})
	pulumi.RegisterOutputType(AuditSinkPtrOutput{})
}
