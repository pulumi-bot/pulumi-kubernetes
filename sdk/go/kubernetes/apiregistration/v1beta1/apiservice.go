// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// APIService represents a server for a particular GroupVersion. Name must be "version.group".
type APIService struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec contains information for locating and communicating with a server
	Spec APIServiceSpecPtrOutput `pulumi:"spec"`
	// Status contains derived information about an API server
	Status APIServiceStatusPtrOutput `pulumi:"status"`
}

// NewAPIService registers a new resource with the given unique name, arguments, and options.
func NewAPIService(ctx *pulumi.Context,
	name string, args *APIServiceArgs, opts ...pulumi.ResourceOption) (*APIService, error) {
	if args == nil {
		args = &APIServiceArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("apiregistration.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("APIService")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:apiregistration.k8s.io/v1:APIService"),
		},
		{
			Type: pulumi.String("kubernetes:apiregistration/v1:APIService"),
		},
		{
			Type: pulumi.String("kubernetes:apiregistration/v1beta1:APIService"),
		},
	})
	opts = append(opts, aliases)
	var resource APIService
	err := ctx.RegisterResource("kubernetes:apiregistration.k8s.io/v1beta1:APIService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAPIService gets an existing APIService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAPIService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *APIServiceState, opts ...pulumi.ResourceOption) (*APIService, error) {
	var resource APIService
	err := ctx.ReadResource("kubernetes:apiregistration.k8s.io/v1beta1:APIService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering APIService resources.
type apiserviceState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec contains information for locating and communicating with a server
	Spec *APIServiceSpec `pulumi:"spec"`
	// Status contains derived information about an API server
	Status *APIServiceStatus `pulumi:"status"`
}

type APIServiceState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Spec contains information for locating and communicating with a server
	Spec APIServiceSpecPtrInput
	// Status contains derived information about an API server
	Status APIServiceStatusPtrInput
}

func (APIServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiserviceState)(nil)).Elem()
}

type apiserviceArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec contains information for locating and communicating with a server
	Spec *APIServiceSpec `pulumi:"spec"`
}

// The set of arguments for constructing a APIService resource.
type APIServiceArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Spec contains information for locating and communicating with a server
	Spec APIServiceSpecPtrInput
}

func (APIServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiserviceArgs)(nil)).Elem()
}

type APIServiceInput interface {
	pulumi.Input

	ToAPIServiceOutput() APIServiceOutput
	ToAPIServiceOutputWithContext(ctx context.Context) APIServiceOutput
}

func (*APIService) ElementType() reflect.Type {
	return reflect.TypeOf((*APIService)(nil))
}

func (i *APIService) ToAPIServiceOutput() APIServiceOutput {
	return i.ToAPIServiceOutputWithContext(context.Background())
}

func (i *APIService) ToAPIServiceOutputWithContext(ctx context.Context) APIServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIServiceOutput)
}

func (i *APIService) ToAPIServicePtrOutput() APIServicePtrOutput {
	return i.ToAPIServicePtrOutputWithContext(context.Background())
}

func (i *APIService) ToAPIServicePtrOutputWithContext(ctx context.Context) APIServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIServicePtrOutput)
}

type APIServicePtrInput interface {
	pulumi.Input

	ToAPIServicePtrOutput() APIServicePtrOutput
	ToAPIServicePtrOutputWithContext(ctx context.Context) APIServicePtrOutput
}

type apiservicePtrType APIServiceArgs

func (*apiservicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**APIService)(nil))
}

func (i *apiservicePtrType) ToAPIServicePtrOutput() APIServicePtrOutput {
	return i.ToAPIServicePtrOutputWithContext(context.Background())
}

func (i *apiservicePtrType) ToAPIServicePtrOutputWithContext(ctx context.Context) APIServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIServicePtrOutput)
}

// APIServiceArrayInput is an input type that accepts APIServiceArray and APIServiceArrayOutput values.
// You can construct a concrete instance of `APIServiceArrayInput` via:
//
//          APIServiceArray{ APIServiceArgs{...} }
type APIServiceArrayInput interface {
	pulumi.Input

	ToAPIServiceArrayOutput() APIServiceArrayOutput
	ToAPIServiceArrayOutputWithContext(context.Context) APIServiceArrayOutput
}

type APIServiceArray []APIServiceInput

func (APIServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*APIService)(nil))
}

func (i APIServiceArray) ToAPIServiceArrayOutput() APIServiceArrayOutput {
	return i.ToAPIServiceArrayOutputWithContext(context.Background())
}

func (i APIServiceArray) ToAPIServiceArrayOutputWithContext(ctx context.Context) APIServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIServiceArrayOutput)
}

// APIServiceMapInput is an input type that accepts APIServiceMap and APIServiceMapOutput values.
// You can construct a concrete instance of `APIServiceMapInput` via:
//
//          APIServiceMap{ "key": APIServiceArgs{...} }
type APIServiceMapInput interface {
	pulumi.Input

	ToAPIServiceMapOutput() APIServiceMapOutput
	ToAPIServiceMapOutputWithContext(context.Context) APIServiceMapOutput
}

type APIServiceMap map[string]APIServiceInput

func (APIServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*APIService)(nil))
}

func (i APIServiceMap) ToAPIServiceMapOutput() APIServiceMapOutput {
	return i.ToAPIServiceMapOutputWithContext(context.Background())
}

func (i APIServiceMap) ToAPIServiceMapOutputWithContext(ctx context.Context) APIServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIServiceMapOutput)
}

type APIServiceOutput struct {
	*pulumi.OutputState
}

func (APIServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*APIService)(nil))
}

func (o APIServiceOutput) ToAPIServiceOutput() APIServiceOutput {
	return o
}

func (o APIServiceOutput) ToAPIServiceOutputWithContext(ctx context.Context) APIServiceOutput {
	return o
}

func (o APIServiceOutput) ToAPIServicePtrOutput() APIServicePtrOutput {
	return o.ToAPIServicePtrOutputWithContext(context.Background())
}

func (o APIServiceOutput) ToAPIServicePtrOutputWithContext(ctx context.Context) APIServicePtrOutput {
	return o.ApplyT(func(v APIService) *APIService {
		return &v
	}).(APIServicePtrOutput)
}

type APIServicePtrOutput struct {
	*pulumi.OutputState
}

func (APIServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**APIService)(nil))
}

func (o APIServicePtrOutput) ToAPIServicePtrOutput() APIServicePtrOutput {
	return o
}

func (o APIServicePtrOutput) ToAPIServicePtrOutputWithContext(ctx context.Context) APIServicePtrOutput {
	return o
}

type APIServiceArrayOutput struct{ *pulumi.OutputState }

func (APIServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]APIService)(nil))
}

func (o APIServiceArrayOutput) ToAPIServiceArrayOutput() APIServiceArrayOutput {
	return o
}

func (o APIServiceArrayOutput) ToAPIServiceArrayOutputWithContext(ctx context.Context) APIServiceArrayOutput {
	return o
}

func (o APIServiceArrayOutput) Index(i pulumi.IntInput) APIServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) APIService {
		return vs[0].([]APIService)[vs[1].(int)]
	}).(APIServiceOutput)
}

type APIServiceMapOutput struct{ *pulumi.OutputState }

func (APIServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]APIService)(nil))
}

func (o APIServiceMapOutput) ToAPIServiceMapOutput() APIServiceMapOutput {
	return o
}

func (o APIServiceMapOutput) ToAPIServiceMapOutputWithContext(ctx context.Context) APIServiceMapOutput {
	return o
}

func (o APIServiceMapOutput) MapIndex(k pulumi.StringInput) APIServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) APIService {
		return vs[0].(map[string]APIService)[vs[1].(string)]
	}).(APIServiceOutput)
}

func init() {
	pulumi.RegisterOutputType(APIServiceOutput{})
	pulumi.RegisterOutputType(APIServicePtrOutput{})
	pulumi.RegisterOutputType(APIServiceArrayOutput{})
	pulumi.RegisterOutputType(APIServiceMapOutput{})
}
