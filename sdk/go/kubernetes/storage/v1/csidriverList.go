// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// CSIDriverList is a collection of CSIDriver objects.
type CSIDriverList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// items is the list of CSIDriver
	Items CSIDriverTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewCSIDriverList registers a new resource with the given unique name, arguments, and options.
func NewCSIDriverList(ctx *pulumi.Context,
	name string, args *CSIDriverListArgs, opts ...pulumi.ResourceOption) (*CSIDriverList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("storage.k8s.io/v1")
	args.Kind = pulumi.StringPtr("CSIDriverList")
	var resource CSIDriverList
	err := ctx.RegisterResource("kubernetes:storage.k8s.io/v1:CSIDriverList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCSIDriverList gets an existing CSIDriverList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCSIDriverList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CSIDriverListState, opts ...pulumi.ResourceOption) (*CSIDriverList, error) {
	var resource CSIDriverList
	err := ctx.ReadResource("kubernetes:storage.k8s.io/v1:CSIDriverList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CSIDriverList resources.
type csidriverListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is the list of CSIDriver
	Items []CSIDriverType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type CSIDriverListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is the list of CSIDriver
	Items CSIDriverTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (CSIDriverListState) ElementType() reflect.Type {
	return reflect.TypeOf((*csidriverListState)(nil)).Elem()
}

type csidriverListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is the list of CSIDriver
	Items []CSIDriverType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a CSIDriverList resource.
type CSIDriverListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is the list of CSIDriver
	Items CSIDriverTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (CSIDriverListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*csidriverListArgs)(nil)).Elem()
}

type CSIDriverListInput interface {
	pulumi.Input

	ToCSIDriverListOutput() CSIDriverListOutput
	ToCSIDriverListOutputWithContext(ctx context.Context) CSIDriverListOutput
}

func (*CSIDriverList) ElementType() reflect.Type {
	return reflect.TypeOf((*CSIDriverList)(nil))
}

func (i *CSIDriverList) ToCSIDriverListOutput() CSIDriverListOutput {
	return i.ToCSIDriverListOutputWithContext(context.Background())
}

func (i *CSIDriverList) ToCSIDriverListOutputWithContext(ctx context.Context) CSIDriverListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIDriverListOutput)
}

type CSIDriverListOutput struct {
	*pulumi.OutputState
}

func (CSIDriverListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CSIDriverList)(nil))
}

func (o CSIDriverListOutput) ToCSIDriverListOutput() CSIDriverListOutput {
	return o
}

func (o CSIDriverListOutput) ToCSIDriverListOutputWithContext(ctx context.Context) CSIDriverListOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CSIDriverListOutput{})
}
