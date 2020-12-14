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

// RuntimeClassList is a list of RuntimeClass objects.
type RuntimeClassList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Items is a list of schema objects.
	Items RuntimeClassTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewRuntimeClassList registers a new resource with the given unique name, arguments, and options.
func NewRuntimeClassList(ctx *pulumi.Context,
	name string, args *RuntimeClassListArgs, opts ...pulumi.ResourceOption) (*RuntimeClassList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("node.k8s.io/v1")
	args.Kind = pulumi.StringPtr("RuntimeClassList")
	var resource RuntimeClassList
	err := ctx.RegisterResource("kubernetes:node.k8s.io/v1:RuntimeClassList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuntimeClassList gets an existing RuntimeClassList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuntimeClassList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuntimeClassListState, opts ...pulumi.ResourceOption) (*RuntimeClassList, error) {
	var resource RuntimeClassList
	err := ctx.ReadResource("kubernetes:node.k8s.io/v1:RuntimeClassList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuntimeClassList resources.
type runtimeClassListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is a list of schema objects.
	Items []RuntimeClassType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type RuntimeClassListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is a list of schema objects.
	Items RuntimeClassTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (RuntimeClassListState) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeClassListState)(nil)).Elem()
}

type runtimeClassListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is a list of schema objects.
	Items []RuntimeClassType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a RuntimeClassList resource.
type RuntimeClassListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is a list of schema objects.
	Items RuntimeClassTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (RuntimeClassListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeClassListArgs)(nil)).Elem()
}

type RuntimeClassListInput interface {
	pulumi.Input

	ToRuntimeClassListOutput() RuntimeClassListOutput
	ToRuntimeClassListOutputWithContext(ctx context.Context) RuntimeClassListOutput
}

func (*RuntimeClassList) ElementType() reflect.Type {
	return reflect.TypeOf((*RuntimeClassList)(nil))
}

func (i *RuntimeClassList) ToRuntimeClassListOutput() RuntimeClassListOutput {
	return i.ToRuntimeClassListOutputWithContext(context.Background())
}

func (i *RuntimeClassList) ToRuntimeClassListOutputWithContext(ctx context.Context) RuntimeClassListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeClassListOutput)
}

type RuntimeClassListOutput struct {
	*pulumi.OutputState
}

func (RuntimeClassListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuntimeClassList)(nil))
}

func (o RuntimeClassListOutput) ToRuntimeClassListOutput() RuntimeClassListOutput {
	return o
}

func (o RuntimeClassListOutput) ToRuntimeClassListOutputWithContext(ctx context.Context) RuntimeClassListOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RuntimeClassListOutput{})
}
