// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// FlowSchemaList is a list of FlowSchema objects.
type FlowSchemaList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// `items` is a list of FlowSchemas.
	Items FlowSchemaTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// `metadata` is the standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewFlowSchemaList registers a new resource with the given unique name, arguments, and options.
func NewFlowSchemaList(ctx *pulumi.Context,
	name string, args *FlowSchemaListArgs, opts ...pulumi.ResourceOption) (*FlowSchemaList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("flowcontrol.apiserver.k8s.io/v1alpha1")
	args.Kind = pulumi.StringPtr("FlowSchemaList")
	var resource FlowSchemaList
	err := ctx.RegisterResource("kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:FlowSchemaList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlowSchemaList gets an existing FlowSchemaList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlowSchemaList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowSchemaListState, opts ...pulumi.ResourceOption) (*FlowSchemaList, error) {
	var resource FlowSchemaList
	err := ctx.ReadResource("kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:FlowSchemaList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlowSchemaList resources.
type flowSchemaListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// `items` is a list of FlowSchemas.
	Items []FlowSchemaType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// `metadata` is the standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type FlowSchemaListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// `items` is a list of FlowSchemas.
	Items FlowSchemaTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// `metadata` is the standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (FlowSchemaListState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowSchemaListState)(nil)).Elem()
}

type flowSchemaListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// `items` is a list of FlowSchemas.
	Items []FlowSchemaType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// `metadata` is the standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a FlowSchemaList resource.
type FlowSchemaListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// `items` is a list of FlowSchemas.
	Items FlowSchemaTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// `metadata` is the standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (FlowSchemaListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowSchemaListArgs)(nil)).Elem()
}

type FlowSchemaListInput interface {
	pulumi.Input

	ToFlowSchemaListOutput() FlowSchemaListOutput
	ToFlowSchemaListOutputWithContext(ctx context.Context) FlowSchemaListOutput
}

func (FlowSchemaList) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowSchemaList)(nil)).Elem()
}

func (i FlowSchemaList) ToFlowSchemaListOutput() FlowSchemaListOutput {
	return i.ToFlowSchemaListOutputWithContext(context.Background())
}

func (i FlowSchemaList) ToFlowSchemaListOutputWithContext(ctx context.Context) FlowSchemaListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowSchemaListOutput)
}

type FlowSchemaListOutput struct {
	*pulumi.OutputState
}

func (FlowSchemaListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowSchemaListOutput)(nil)).Elem()
}

func (o FlowSchemaListOutput) ToFlowSchemaListOutput() FlowSchemaListOutput {
	return o
}

func (o FlowSchemaListOutput) ToFlowSchemaListOutputWithContext(ctx context.Context) FlowSchemaListOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FlowSchemaListOutput{})
}
