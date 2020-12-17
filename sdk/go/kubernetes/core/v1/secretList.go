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

// SecretList is a list of Secret.
type SecretList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Items is a list of secret objects. More info: https://kubernetes.io/docs/concepts/configuration/secret
	Items SecretTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewSecretList registers a new resource with the given unique name, arguments, and options.
func NewSecretList(ctx *pulumi.Context,
	name string, args *SecretListArgs, opts ...pulumi.ResourceOption) (*SecretList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("SecretList")
	var resource SecretList
	err := ctx.RegisterResource("kubernetes:core/v1:SecretList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretList gets an existing SecretList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretListState, opts ...pulumi.ResourceOption) (*SecretList, error) {
	var resource SecretList
	err := ctx.ReadResource("kubernetes:core/v1:SecretList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretList resources.
type secretListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is a list of secret objects. More info: https://kubernetes.io/docs/concepts/configuration/secret
	Items []SecretType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type SecretListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is a list of secret objects. More info: https://kubernetes.io/docs/concepts/configuration/secret
	Items SecretTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (SecretListState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretListState)(nil)).Elem()
}

type secretListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is a list of secret objects. More info: https://kubernetes.io/docs/concepts/configuration/secret
	Items []SecretType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a SecretList resource.
type SecretListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is a list of secret objects. More info: https://kubernetes.io/docs/concepts/configuration/secret
	Items SecretTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (SecretListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretListArgs)(nil)).Elem()
}

type SecretListInput interface {
	pulumi.Input

	ToSecretListOutput() SecretListOutput
	ToSecretListOutputWithContext(ctx context.Context) SecretListOutput
}

func (*SecretList) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretList)(nil))
}

func (i *SecretList) ToSecretListOutput() SecretListOutput {
	return i.ToSecretListOutputWithContext(context.Background())
}

func (i *SecretList) ToSecretListOutputWithContext(ctx context.Context) SecretListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretListOutput)
}

func (i *SecretList) ToSecretListPtrOutput() SecretListPtrOutput {
	return i.ToSecretListPtrOutputWithContext(context.Background())
}

func (i *SecretList) ToSecretListPtrOutputWithContext(ctx context.Context) SecretListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretListPtrOutput)
}

type SecretListPtrInput interface {
	pulumi.Input

	ToSecretListPtrOutput() SecretListPtrOutput
	ToSecretListPtrOutputWithContext(ctx context.Context) SecretListPtrOutput
}

type secretListPtrType SecretListArgs

func (*secretListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretList)(nil))
}

func (i *secretListPtrType) ToSecretListPtrOutput() SecretListPtrOutput {
	return i.ToSecretListPtrOutputWithContext(context.Background())
}

func (i *secretListPtrType) ToSecretListPtrOutputWithContext(ctx context.Context) SecretListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretListOutput).ToSecretListPtrOutput()
}

// SecretListArrayInput is an input type that accepts SecretListArray and SecretListArrayOutput values.
// You can construct a concrete instance of `SecretListArrayInput` via:
//
//          SecretListArray{ SecretListArgs{...} }
type SecretListArrayInput interface {
	pulumi.Input

	ToSecretListArrayOutput() SecretListArrayOutput
	ToSecretListArrayOutputWithContext(context.Context) SecretListArrayOutput
}

type SecretListArray []SecretListInput

func (SecretListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretList)(nil))
}

func (i SecretListArray) ToSecretListArrayOutput() SecretListArrayOutput {
	return i.ToSecretListArrayOutputWithContext(context.Background())
}

func (i SecretListArray) ToSecretListArrayOutputWithContext(ctx context.Context) SecretListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretListArrayOutput)
}

// SecretListMapInput is an input type that accepts SecretListMap and SecretListMapOutput values.
// You can construct a concrete instance of `SecretListMapInput` via:
//
//          SecretListMap{ "key": SecretListArgs{...} }
type SecretListMapInput interface {
	pulumi.Input

	ToSecretListMapOutput() SecretListMapOutput
	ToSecretListMapOutputWithContext(context.Context) SecretListMapOutput
}

type SecretListMap map[string]SecretListInput

func (SecretListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SecretList)(nil))
}

func (i SecretListMap) ToSecretListMapOutput() SecretListMapOutput {
	return i.ToSecretListMapOutputWithContext(context.Background())
}

func (i SecretListMap) ToSecretListMapOutputWithContext(ctx context.Context) SecretListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretListMapOutput)
}

type SecretListOutput struct {
	*pulumi.OutputState
}

func (SecretListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretList)(nil))
}

func (o SecretListOutput) ToSecretListOutput() SecretListOutput {
	return o
}

func (o SecretListOutput) ToSecretListOutputWithContext(ctx context.Context) SecretListOutput {
	return o
}

func (o SecretListOutput) ToSecretListPtrOutput() SecretListPtrOutput {
	return o.ToSecretListPtrOutputWithContext(context.Background())
}

func (o SecretListOutput) ToSecretListPtrOutputWithContext(ctx context.Context) SecretListPtrOutput {
	return o.ApplyT(func(v SecretList) *SecretList {
		return &v
	}).(SecretListPtrOutput)
}

type SecretListPtrOutput struct {
	*pulumi.OutputState
}

func (SecretListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretList)(nil))
}

func (o SecretListPtrOutput) ToSecretListPtrOutput() SecretListPtrOutput {
	return o
}

func (o SecretListPtrOutput) ToSecretListPtrOutputWithContext(ctx context.Context) SecretListPtrOutput {
	return o
}

type SecretListArrayOutput struct{ *pulumi.OutputState }

func (SecretListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretList)(nil))
}

func (o SecretListArrayOutput) ToSecretListArrayOutput() SecretListArrayOutput {
	return o
}

func (o SecretListArrayOutput) ToSecretListArrayOutputWithContext(ctx context.Context) SecretListArrayOutput {
	return o
}

func (o SecretListArrayOutput) Index(i pulumi.IntInput) SecretListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretList {
		return vs[0].([]SecretList)[vs[1].(int)]
	}).(SecretListOutput)
}

type SecretListMapOutput struct{ *pulumi.OutputState }

func (SecretListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SecretList)(nil))
}

func (o SecretListMapOutput) ToSecretListMapOutput() SecretListMapOutput {
	return o
}

func (o SecretListMapOutput) ToSecretListMapOutputWithContext(ctx context.Context) SecretListMapOutput {
	return o
}

func (o SecretListMapOutput) MapIndex(k pulumi.StringInput) SecretListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SecretList {
		return vs[0].(map[string]SecretList)[vs[1].(string)]
	}).(SecretListOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretListOutput{})
	pulumi.RegisterOutputType(SecretListPtrOutput{})
	pulumi.RegisterOutputType(SecretListArrayOutput{})
	pulumi.RegisterOutputType(SecretListMapOutput{})
}
