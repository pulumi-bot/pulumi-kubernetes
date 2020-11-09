// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// APIServiceList is a list of APIService objects.
type APIServiceList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput    `pulumi:"apiVersion"`
	Items      APIServiceTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput   `pulumi:"kind"`
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewAPIServiceList registers a new resource with the given unique name, arguments, and options.
func NewAPIServiceList(ctx *pulumi.Context,
	name string, args *APIServiceListArgs, opts ...pulumi.ResourceOption) (*APIServiceList, error) {
	if args == nil || args.Items == nil {
		return nil, errors.New("missing required argument 'Items'")
	}
	if args == nil {
		args = &APIServiceListArgs{}
	}
	args.ApiVersion = pulumi.StringPtr("apiregistration.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("APIServiceList")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:apiregistration/v1beta1:APIServiceList"),
		},
	})
	opts = append(opts, aliases)
	var resource APIServiceList
	err := ctx.RegisterResource("kubernetes:apiregistration.k8s.io/v1beta1:APIServiceList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAPIServiceList gets an existing APIServiceList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAPIServiceList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *APIServiceListState, opts ...pulumi.ResourceOption) (*APIServiceList, error) {
	var resource APIServiceList
	err := ctx.ReadResource("kubernetes:apiregistration.k8s.io/v1beta1:APIServiceList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering APIServiceList resources.
type apiserviceListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string          `pulumi:"apiVersion"`
	Items      []APIServiceType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string          `pulumi:"kind"`
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type APIServiceListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	Items      APIServiceTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ListMetaPtrInput
}

func (APIServiceListState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiserviceListState)(nil)).Elem()
}

type apiserviceListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string          `pulumi:"apiVersion"`
	Items      []APIServiceType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string          `pulumi:"kind"`
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a APIServiceList resource.
type APIServiceListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	Items      APIServiceTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ListMetaPtrInput
}

func (APIServiceListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiserviceListArgs)(nil)).Elem()
}

type APIServiceListInput interface {
	pulumi.Input

	ToAPIServiceListOutput() APIServiceListOutput
	ToAPIServiceListOutputWithContext(ctx context.Context) APIServiceListOutput
}

func (APIServiceList) ElementType() reflect.Type {
	return reflect.TypeOf((*APIServiceList)(nil)).Elem()
}

func (i APIServiceList) ToAPIServiceListOutput() APIServiceListOutput {
	return i.ToAPIServiceListOutputWithContext(context.Background())
}

func (i APIServiceList) ToAPIServiceListOutputWithContext(ctx context.Context) APIServiceListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIServiceListOutput)
}

type APIServiceListOutput struct {
	*pulumi.OutputState
}

func (APIServiceListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*APIServiceListOutput)(nil)).Elem()
}

func (o APIServiceListOutput) ToAPIServiceListOutput() APIServiceListOutput {
	return o
}

func (o APIServiceListOutput) ToAPIServiceListOutputWithContext(ctx context.Context) APIServiceListOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(APIServiceListOutput{})
}
