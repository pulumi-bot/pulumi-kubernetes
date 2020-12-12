// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// StorageClassList is a collection of storage classes.
type StorageClassList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Items is the list of StorageClasses
	Items StorageClassTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewStorageClassList registers a new resource with the given unique name, arguments, and options.
func NewStorageClassList(ctx *pulumi.Context,
	name string, args *StorageClassListArgs, opts ...pulumi.ResourceOption) (*StorageClassList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("storage.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("StorageClassList")
	var resource StorageClassList
	err := ctx.RegisterResource("kubernetes:storage.k8s.io/v1beta1:StorageClassList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageClassList gets an existing StorageClassList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageClassList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageClassListState, opts ...pulumi.ResourceOption) (*StorageClassList, error) {
	var resource StorageClassList
	err := ctx.ReadResource("kubernetes:storage.k8s.io/v1beta1:StorageClassList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageClassList resources.
type storageClassListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is the list of StorageClasses
	Items []StorageClassType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type StorageClassListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is the list of StorageClasses
	Items StorageClassTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (StorageClassListState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageClassListState)(nil)).Elem()
}

type storageClassListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is the list of StorageClasses
	Items []StorageClassType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a StorageClassList resource.
type StorageClassListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is the list of StorageClasses
	Items StorageClassTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (StorageClassListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageClassListArgs)(nil)).Elem()
}

type StorageClassListInput interface {
	pulumi.Input

	ToStorageClassListOutput() StorageClassListOutput
	ToStorageClassListOutputWithContext(ctx context.Context) StorageClassListOutput
}

func (*StorageClassList) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageClassList)(nil))
}

func (i *StorageClassList) ToStorageClassListOutput() StorageClassListOutput {
	return i.ToStorageClassListOutputWithContext(context.Background())
}

func (i *StorageClassList) ToStorageClassListOutputWithContext(ctx context.Context) StorageClassListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageClassListOutput)
}

func (i *StorageClassList) ToStorageClassListPtrOutput() StorageClassListPtrOutput {
	return i.ToStorageClassListPtrOutputWithContext(context.Background())
}

func (i *StorageClassList) ToStorageClassListPtrOutputWithContext(ctx context.Context) StorageClassListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageClassListPtrOutput)
}

type StorageClassListPtrInput interface {
	pulumi.Input

	ToStorageClassListPtrOutput() StorageClassListPtrOutput
	ToStorageClassListPtrOutputWithContext(ctx context.Context) StorageClassListPtrOutput
}

type StorageClassListOutput struct {
	*pulumi.OutputState
}

func (StorageClassListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageClassList)(nil))
}

func (o StorageClassListOutput) ToStorageClassListOutput() StorageClassListOutput {
	return o
}

func (o StorageClassListOutput) ToStorageClassListOutputWithContext(ctx context.Context) StorageClassListOutput {
	return o
}

type StorageClassListPtrOutput struct {
	*pulumi.OutputState
}

func (StorageClassListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageClassList)(nil))
}

func (o StorageClassListPtrOutput) ToStorageClassListPtrOutput() StorageClassListPtrOutput {
	return o
}

func (o StorageClassListPtrOutput) ToStorageClassListPtrOutputWithContext(ctx context.Context) StorageClassListPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StorageClassListOutput{})
	pulumi.RegisterOutputType(StorageClassListPtrOutput{})
}
