// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// PersistentVolumeClaim is a user's request for and claim to a persistent volume
type PersistentVolumeClaim struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Spec PersistentVolumeClaimSpecPtrOutput `pulumi:"spec"`
	// Status represents the current information/status of a persistent volume claim. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Status PersistentVolumeClaimStatusPtrOutput `pulumi:"status"`
}

// NewPersistentVolumeClaim registers a new resource with the given unique name, arguments, and options.
func NewPersistentVolumeClaim(ctx *pulumi.Context,
	name string, args *PersistentVolumeClaimArgs, opts ...pulumi.ResourceOption) (*PersistentVolumeClaim, error) {
	if args == nil {
		args = &PersistentVolumeClaimArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("PersistentVolumeClaim")
	var resource PersistentVolumeClaim
	err := ctx.RegisterResource("kubernetes:core/v1:PersistentVolumeClaim", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPersistentVolumeClaim gets an existing PersistentVolumeClaim resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPersistentVolumeClaim(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PersistentVolumeClaimState, opts ...pulumi.ResourceOption) (*PersistentVolumeClaim, error) {
	var resource PersistentVolumeClaim
	err := ctx.ReadResource("kubernetes:core/v1:PersistentVolumeClaim", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PersistentVolumeClaim resources.
type persistentVolumeClaimState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Spec *PersistentVolumeClaimSpec `pulumi:"spec"`
	// Status represents the current information/status of a persistent volume claim. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Status *PersistentVolumeClaimStatus `pulumi:"status"`
}

type PersistentVolumeClaimState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Spec PersistentVolumeClaimSpecPtrInput
	// Status represents the current information/status of a persistent volume claim. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Status PersistentVolumeClaimStatusPtrInput
}

func (PersistentVolumeClaimState) ElementType() reflect.Type {
	return reflect.TypeOf((*persistentVolumeClaimState)(nil)).Elem()
}

type persistentVolumeClaimArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Spec *PersistentVolumeClaimSpec `pulumi:"spec"`
}

// The set of arguments for constructing a PersistentVolumeClaim resource.
type PersistentVolumeClaimArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Spec PersistentVolumeClaimSpecPtrInput
}

func (PersistentVolumeClaimArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*persistentVolumeClaimArgs)(nil)).Elem()
}

type PersistentVolumeClaimInput interface {
	pulumi.Input

	ToPersistentVolumeClaimOutput() PersistentVolumeClaimOutput
	ToPersistentVolumeClaimOutputWithContext(ctx context.Context) PersistentVolumeClaimOutput
}

type PersistentVolumeClaimPtrInput interface {
	pulumi.Input

	ToPersistentVolumeClaimPtrOutput() PersistentVolumeClaimPtrOutput
	ToPersistentVolumeClaimPtrOutputWithContext(ctx context.Context) PersistentVolumeClaimPtrOutput
}

func (PersistentVolumeClaim) ElementType() reflect.Type {
	return reflect.TypeOf((*PersistentVolumeClaim)(nil)).Elem()
}

func (i PersistentVolumeClaim) ToPersistentVolumeClaimOutput() PersistentVolumeClaimOutput {
	return i.ToPersistentVolumeClaimOutputWithContext(context.Background())
}

func (i PersistentVolumeClaim) ToPersistentVolumeClaimOutputWithContext(ctx context.Context) PersistentVolumeClaimOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumeClaimOutput)
}

func (i PersistentVolumeClaim) ToPersistentVolumeClaimPtrOutput() PersistentVolumeClaimPtrOutput {
	return i.ToPersistentVolumeClaimPtrOutputWithContext(context.Background())
}

func (i PersistentVolumeClaim) ToPersistentVolumeClaimPtrOutputWithContext(ctx context.Context) PersistentVolumeClaimPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumeClaimPtrOutput)
}

type PersistentVolumeClaimOutput struct {
	*pulumi.OutputState
}

func (PersistentVolumeClaimOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PersistentVolumeClaimOutput)(nil)).Elem()
}

func (o PersistentVolumeClaimOutput) ToPersistentVolumeClaimOutput() PersistentVolumeClaimOutput {
	return o
}

func (o PersistentVolumeClaimOutput) ToPersistentVolumeClaimOutputWithContext(ctx context.Context) PersistentVolumeClaimOutput {
	return o
}

type PersistentVolumeClaimPtrOutput struct {
	*pulumi.OutputState
}

func (PersistentVolumeClaimPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistentVolumeClaim)(nil)).Elem()
}

func (o PersistentVolumeClaimPtrOutput) ToPersistentVolumeClaimPtrOutput() PersistentVolumeClaimPtrOutput {
	return o
}

func (o PersistentVolumeClaimPtrOutput) ToPersistentVolumeClaimPtrOutputWithContext(ctx context.Context) PersistentVolumeClaimPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PersistentVolumeClaimOutput{})
	pulumi.RegisterOutputType(PersistentVolumeClaimPtrOutput{})
}
