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

func (*PersistentVolumeClaim) ElementType() reflect.Type {
	return reflect.TypeOf((*PersistentVolumeClaim)(nil))
}

func (i *PersistentVolumeClaim) ToPersistentVolumeClaimOutput() PersistentVolumeClaimOutput {
	return i.ToPersistentVolumeClaimOutputWithContext(context.Background())
}

func (i *PersistentVolumeClaim) ToPersistentVolumeClaimOutputWithContext(ctx context.Context) PersistentVolumeClaimOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumeClaimOutput)
}

func (i *PersistentVolumeClaim) ToPersistentVolumeClaimPtrOutput() PersistentVolumeClaimPtrOutput {
	return i.ToPersistentVolumeClaimPtrOutputWithContext(context.Background())
}

func (i *PersistentVolumeClaim) ToPersistentVolumeClaimPtrOutputWithContext(ctx context.Context) PersistentVolumeClaimPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumeClaimPtrOutput)
}

type PersistentVolumeClaimPtrInput interface {
	pulumi.Input

	ToPersistentVolumeClaimPtrOutput() PersistentVolumeClaimPtrOutput
	ToPersistentVolumeClaimPtrOutputWithContext(ctx context.Context) PersistentVolumeClaimPtrOutput
}

type persistentVolumeClaimPtrType PersistentVolumeClaimArgs

func (*persistentVolumeClaimPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistentVolumeClaim)(nil))
}

func (i *persistentVolumeClaimPtrType) ToPersistentVolumeClaimPtrOutput() PersistentVolumeClaimPtrOutput {
	return i.ToPersistentVolumeClaimPtrOutputWithContext(context.Background())
}

func (i *persistentVolumeClaimPtrType) ToPersistentVolumeClaimPtrOutputWithContext(ctx context.Context) PersistentVolumeClaimPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumeClaimOutput).ToPersistentVolumeClaimPtrOutput()
}

// PersistentVolumeClaimArrayInput is an input type that accepts PersistentVolumeClaimArray and PersistentVolumeClaimArrayOutput values.
// You can construct a concrete instance of `PersistentVolumeClaimArrayInput` via:
//
//          PersistentVolumeClaimArray{ PersistentVolumeClaimArgs{...} }
type PersistentVolumeClaimArrayInput interface {
	pulumi.Input

	ToPersistentVolumeClaimArrayOutput() PersistentVolumeClaimArrayOutput
	ToPersistentVolumeClaimArrayOutputWithContext(context.Context) PersistentVolumeClaimArrayOutput
}

type PersistentVolumeClaimArray []PersistentVolumeClaimInput

func (PersistentVolumeClaimArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PersistentVolumeClaim)(nil))
}

func (i PersistentVolumeClaimArray) ToPersistentVolumeClaimArrayOutput() PersistentVolumeClaimArrayOutput {
	return i.ToPersistentVolumeClaimArrayOutputWithContext(context.Background())
}

func (i PersistentVolumeClaimArray) ToPersistentVolumeClaimArrayOutputWithContext(ctx context.Context) PersistentVolumeClaimArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumeClaimArrayOutput)
}

// PersistentVolumeClaimMapInput is an input type that accepts PersistentVolumeClaimMap and PersistentVolumeClaimMapOutput values.
// You can construct a concrete instance of `PersistentVolumeClaimMapInput` via:
//
//          PersistentVolumeClaimMap{ "key": PersistentVolumeClaimArgs{...} }
type PersistentVolumeClaimMapInput interface {
	pulumi.Input

	ToPersistentVolumeClaimMapOutput() PersistentVolumeClaimMapOutput
	ToPersistentVolumeClaimMapOutputWithContext(context.Context) PersistentVolumeClaimMapOutput
}

type PersistentVolumeClaimMap map[string]PersistentVolumeClaimInput

func (PersistentVolumeClaimMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PersistentVolumeClaim)(nil))
}

func (i PersistentVolumeClaimMap) ToPersistentVolumeClaimMapOutput() PersistentVolumeClaimMapOutput {
	return i.ToPersistentVolumeClaimMapOutputWithContext(context.Background())
}

func (i PersistentVolumeClaimMap) ToPersistentVolumeClaimMapOutputWithContext(ctx context.Context) PersistentVolumeClaimMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumeClaimMapOutput)
}

type PersistentVolumeClaimOutput struct {
	*pulumi.OutputState
}

func (PersistentVolumeClaimOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PersistentVolumeClaim)(nil))
}

func (o PersistentVolumeClaimOutput) ToPersistentVolumeClaimOutput() PersistentVolumeClaimOutput {
	return o
}

func (o PersistentVolumeClaimOutput) ToPersistentVolumeClaimOutputWithContext(ctx context.Context) PersistentVolumeClaimOutput {
	return o
}

func (o PersistentVolumeClaimOutput) ToPersistentVolumeClaimPtrOutput() PersistentVolumeClaimPtrOutput {
	return o.ToPersistentVolumeClaimPtrOutputWithContext(context.Background())
}

func (o PersistentVolumeClaimOutput) ToPersistentVolumeClaimPtrOutputWithContext(ctx context.Context) PersistentVolumeClaimPtrOutput {
	return o.ApplyT(func(v PersistentVolumeClaim) *PersistentVolumeClaim {
		return &v
	}).(PersistentVolumeClaimPtrOutput)
}

type PersistentVolumeClaimPtrOutput struct {
	*pulumi.OutputState
}

func (PersistentVolumeClaimPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistentVolumeClaim)(nil))
}

func (o PersistentVolumeClaimPtrOutput) ToPersistentVolumeClaimPtrOutput() PersistentVolumeClaimPtrOutput {
	return o
}

func (o PersistentVolumeClaimPtrOutput) ToPersistentVolumeClaimPtrOutputWithContext(ctx context.Context) PersistentVolumeClaimPtrOutput {
	return o
}

type PersistentVolumeClaimArrayOutput struct{ *pulumi.OutputState }

func (PersistentVolumeClaimArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PersistentVolumeClaim)(nil))
}

func (o PersistentVolumeClaimArrayOutput) ToPersistentVolumeClaimArrayOutput() PersistentVolumeClaimArrayOutput {
	return o
}

func (o PersistentVolumeClaimArrayOutput) ToPersistentVolumeClaimArrayOutputWithContext(ctx context.Context) PersistentVolumeClaimArrayOutput {
	return o
}

func (o PersistentVolumeClaimArrayOutput) Index(i pulumi.IntInput) PersistentVolumeClaimOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PersistentVolumeClaim {
		return vs[0].([]PersistentVolumeClaim)[vs[1].(int)]
	}).(PersistentVolumeClaimOutput)
}

type PersistentVolumeClaimMapOutput struct{ *pulumi.OutputState }

func (PersistentVolumeClaimMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PersistentVolumeClaim)(nil))
}

func (o PersistentVolumeClaimMapOutput) ToPersistentVolumeClaimMapOutput() PersistentVolumeClaimMapOutput {
	return o
}

func (o PersistentVolumeClaimMapOutput) ToPersistentVolumeClaimMapOutputWithContext(ctx context.Context) PersistentVolumeClaimMapOutput {
	return o
}

func (o PersistentVolumeClaimMapOutput) MapIndex(k pulumi.StringInput) PersistentVolumeClaimOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PersistentVolumeClaim {
		return vs[0].(map[string]PersistentVolumeClaim)[vs[1].(string)]
	}).(PersistentVolumeClaimOutput)
}

func init() {
	pulumi.RegisterOutputType(PersistentVolumeClaimOutput{})
	pulumi.RegisterOutputType(PersistentVolumeClaimPtrOutput{})
	pulumi.RegisterOutputType(PersistentVolumeClaimArrayOutput{})
	pulumi.RegisterOutputType(PersistentVolumeClaimMapOutput{})
}
