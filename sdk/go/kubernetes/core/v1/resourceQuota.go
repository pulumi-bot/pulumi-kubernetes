// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ResourceQuota sets aggregate quota restrictions enforced per namespace
type ResourceQuota struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec defines the desired quota. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec ResourceQuotaSpecPtrOutput `pulumi:"spec"`
	// Status defines the actual enforced quota and its current usage. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status ResourceQuotaStatusPtrOutput `pulumi:"status"`
}

// NewResourceQuota registers a new resource with the given unique name, arguments, and options.
func NewResourceQuota(ctx *pulumi.Context,
	name string, args *ResourceQuotaArgs, opts ...pulumi.ResourceOption) (*ResourceQuota, error) {
	if args == nil {
		args = &ResourceQuotaArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("ResourceQuota")
	var resource ResourceQuota
	err := ctx.RegisterResource("kubernetes:core/v1:ResourceQuota", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceQuota gets an existing ResourceQuota resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceQuota(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceQuotaState, opts ...pulumi.ResourceOption) (*ResourceQuota, error) {
	var resource ResourceQuota
	err := ctx.ReadResource("kubernetes:core/v1:ResourceQuota", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceQuota resources.
type resourceQuotaState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec defines the desired quota. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *ResourceQuotaSpec `pulumi:"spec"`
	// Status defines the actual enforced quota and its current usage. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status *ResourceQuotaStatus `pulumi:"status"`
}

type ResourceQuotaState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Spec defines the desired quota. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec ResourceQuotaSpecPtrInput
	// Status defines the actual enforced quota and its current usage. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status ResourceQuotaStatusPtrInput
}

func (ResourceQuotaState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceQuotaState)(nil)).Elem()
}

type resourceQuotaArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec defines the desired quota. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *ResourceQuotaSpec `pulumi:"spec"`
}

// The set of arguments for constructing a ResourceQuota resource.
type ResourceQuotaArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Spec defines the desired quota. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec ResourceQuotaSpecPtrInput
}

func (ResourceQuotaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceQuotaArgs)(nil)).Elem()
}

type ResourceQuotaInput interface {
	pulumi.Input

	ToResourceQuotaOutput() ResourceQuotaOutput
	ToResourceQuotaOutputWithContext(ctx context.Context) ResourceQuotaOutput
}

func (*ResourceQuota) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceQuota)(nil))
}

func (i *ResourceQuota) ToResourceQuotaOutput() ResourceQuotaOutput {
	return i.ToResourceQuotaOutputWithContext(context.Background())
}

func (i *ResourceQuota) ToResourceQuotaOutputWithContext(ctx context.Context) ResourceQuotaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceQuotaOutput)
}

func (i *ResourceQuota) ToResourceQuotaPtrOutput() ResourceQuotaPtrOutput {
	return i.ToResourceQuotaPtrOutputWithContext(context.Background())
}

func (i *ResourceQuota) ToResourceQuotaPtrOutputWithContext(ctx context.Context) ResourceQuotaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceQuotaPtrOutput)
}

type ResourceQuotaPtrInput interface {
	pulumi.Input

	ToResourceQuotaPtrOutput() ResourceQuotaPtrOutput
	ToResourceQuotaPtrOutputWithContext(ctx context.Context) ResourceQuotaPtrOutput
}

type ResourceQuotaOutput struct {
	*pulumi.OutputState
}

func (ResourceQuotaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceQuota)(nil))
}

func (o ResourceQuotaOutput) ToResourceQuotaOutput() ResourceQuotaOutput {
	return o
}

func (o ResourceQuotaOutput) ToResourceQuotaOutputWithContext(ctx context.Context) ResourceQuotaOutput {
	return o
}

type ResourceQuotaPtrOutput struct {
	*pulumi.OutputState
}

func (ResourceQuotaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceQuota)(nil))
}

func (o ResourceQuotaPtrOutput) ToResourceQuotaPtrOutput() ResourceQuotaPtrOutput {
	return o
}

func (o ResourceQuotaPtrOutput) ToResourceQuotaPtrOutputWithContext(ctx context.Context) ResourceQuotaPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResourceQuotaOutput{})
	pulumi.RegisterOutputType(ResourceQuotaPtrOutput{})
}
