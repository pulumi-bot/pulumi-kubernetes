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

// ConfigMapList is a resource containing a list of ConfigMap objects.
type ConfigMapList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Items is the list of ConfigMaps.
	Items ConfigMapTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewConfigMapList registers a new resource with the given unique name, arguments, and options.
func NewConfigMapList(ctx *pulumi.Context,
	name string, args *ConfigMapListArgs, opts ...pulumi.ResourceOption) (*ConfigMapList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("ConfigMapList")
	var resource ConfigMapList
	err := ctx.RegisterResource("kubernetes:core/v1:ConfigMapList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigMapList gets an existing ConfigMapList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigMapList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigMapListState, opts ...pulumi.ResourceOption) (*ConfigMapList, error) {
	var resource ConfigMapList
	err := ctx.ReadResource("kubernetes:core/v1:ConfigMapList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigMapList resources.
type configMapListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is the list of ConfigMaps.
	Items []ConfigMapType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type ConfigMapListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is the list of ConfigMaps.
	Items ConfigMapTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (ConfigMapListState) ElementType() reflect.Type {
	return reflect.TypeOf((*configMapListState)(nil)).Elem()
}

type configMapListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is the list of ConfigMaps.
	Items []ConfigMapType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ConfigMapList resource.
type ConfigMapListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is the list of ConfigMaps.
	Items ConfigMapTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (ConfigMapListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configMapListArgs)(nil)).Elem()
}

type ConfigMapListInput interface {
	pulumi.Input

	ToConfigMapListOutput() ConfigMapListOutput
	ToConfigMapListOutputWithContext(ctx context.Context) ConfigMapListOutput
}

func (*ConfigMapList) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigMapList)(nil))
}

func (i *ConfigMapList) ToConfigMapListOutput() ConfigMapListOutput {
	return i.ToConfigMapListOutputWithContext(context.Background())
}

func (i *ConfigMapList) ToConfigMapListOutputWithContext(ctx context.Context) ConfigMapListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigMapListOutput)
}

func (i *ConfigMapList) ToConfigMapListPtrOutput() ConfigMapListPtrOutput {
	return i.ToConfigMapListPtrOutputWithContext(context.Background())
}

func (i *ConfigMapList) ToConfigMapListPtrOutputWithContext(ctx context.Context) ConfigMapListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigMapListPtrOutput)
}

type ConfigMapListPtrInput interface {
	pulumi.Input

	ToConfigMapListPtrOutput() ConfigMapListPtrOutput
	ToConfigMapListPtrOutputWithContext(ctx context.Context) ConfigMapListPtrOutput
}

type ConfigMapListOutput struct {
	*pulumi.OutputState
}

func (ConfigMapListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigMapList)(nil))
}

func (o ConfigMapListOutput) ToConfigMapListOutput() ConfigMapListOutput {
	return o
}

func (o ConfigMapListOutput) ToConfigMapListOutputWithContext(ctx context.Context) ConfigMapListOutput {
	return o
}

type ConfigMapListPtrOutput struct {
	*pulumi.OutputState
}

func (ConfigMapListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigMapList)(nil))
}

func (o ConfigMapListPtrOutput) ToConfigMapListPtrOutput() ConfigMapListPtrOutput {
	return o
}

func (o ConfigMapListPtrOutput) ToConfigMapListPtrOutputWithContext(ctx context.Context) ConfigMapListPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConfigMapListOutput{})
	pulumi.RegisterOutputType(ConfigMapListPtrOutput{})
}
