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

// LocalSubjectAccessReview checks whether or not a user or group can perform an action in a given namespace. Having a namespace scoped resource makes it much easier to grant namespace scoped policy that includes permissions checking.
type LocalSubjectAccessReview struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec holds information about the request being evaluated.  spec.namespace must be equal to the namespace you made the request against.  If empty, it is defaulted.
	Spec SubjectAccessReviewSpecOutput `pulumi:"spec"`
	// Status is filled in by the server and indicates whether the request is allowed or not
	Status SubjectAccessReviewStatusPtrOutput `pulumi:"status"`
}

// NewLocalSubjectAccessReview registers a new resource with the given unique name, arguments, and options.
func NewLocalSubjectAccessReview(ctx *pulumi.Context,
	name string, args *LocalSubjectAccessReviewArgs, opts ...pulumi.ResourceOption) (*LocalSubjectAccessReview, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	args.ApiVersion = pulumi.StringPtr("authorization.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("LocalSubjectAccessReview")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:authorization.k8s.io/v1:LocalSubjectAccessReview"),
		},
	})
	opts = append(opts, aliases)
	var resource LocalSubjectAccessReview
	err := ctx.RegisterResource("kubernetes:authorization.k8s.io/v1beta1:LocalSubjectAccessReview", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalSubjectAccessReview gets an existing LocalSubjectAccessReview resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalSubjectAccessReview(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalSubjectAccessReviewState, opts ...pulumi.ResourceOption) (*LocalSubjectAccessReview, error) {
	var resource LocalSubjectAccessReview
	err := ctx.ReadResource("kubernetes:authorization.k8s.io/v1beta1:LocalSubjectAccessReview", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalSubjectAccessReview resources.
type localSubjectAccessReviewState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec holds information about the request being evaluated.  spec.namespace must be equal to the namespace you made the request against.  If empty, it is defaulted.
	Spec *SubjectAccessReviewSpec `pulumi:"spec"`
	// Status is filled in by the server and indicates whether the request is allowed or not
	Status *SubjectAccessReviewStatus `pulumi:"status"`
}

type LocalSubjectAccessReviewState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Spec holds information about the request being evaluated.  spec.namespace must be equal to the namespace you made the request against.  If empty, it is defaulted.
	Spec SubjectAccessReviewSpecPtrInput
	// Status is filled in by the server and indicates whether the request is allowed or not
	Status SubjectAccessReviewStatusPtrInput
}

func (LocalSubjectAccessReviewState) ElementType() reflect.Type {
	return reflect.TypeOf((*localSubjectAccessReviewState)(nil)).Elem()
}

type localSubjectAccessReviewArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec holds information about the request being evaluated.  spec.namespace must be equal to the namespace you made the request against.  If empty, it is defaulted.
	Spec SubjectAccessReviewSpec `pulumi:"spec"`
}

// The set of arguments for constructing a LocalSubjectAccessReview resource.
type LocalSubjectAccessReviewArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Spec holds information about the request being evaluated.  spec.namespace must be equal to the namespace you made the request against.  If empty, it is defaulted.
	Spec SubjectAccessReviewSpecInput
}

func (LocalSubjectAccessReviewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localSubjectAccessReviewArgs)(nil)).Elem()
}

type LocalSubjectAccessReviewInput interface {
	pulumi.Input

	ToLocalSubjectAccessReviewOutput() LocalSubjectAccessReviewOutput
	ToLocalSubjectAccessReviewOutputWithContext(ctx context.Context) LocalSubjectAccessReviewOutput
}

func (*LocalSubjectAccessReview) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalSubjectAccessReview)(nil))
}

func (i *LocalSubjectAccessReview) ToLocalSubjectAccessReviewOutput() LocalSubjectAccessReviewOutput {
	return i.ToLocalSubjectAccessReviewOutputWithContext(context.Background())
}

func (i *LocalSubjectAccessReview) ToLocalSubjectAccessReviewOutputWithContext(ctx context.Context) LocalSubjectAccessReviewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalSubjectAccessReviewOutput)
}

func (i *LocalSubjectAccessReview) ToLocalSubjectAccessReviewPtrOutput() LocalSubjectAccessReviewPtrOutput {
	return i.ToLocalSubjectAccessReviewPtrOutputWithContext(context.Background())
}

func (i *LocalSubjectAccessReview) ToLocalSubjectAccessReviewPtrOutputWithContext(ctx context.Context) LocalSubjectAccessReviewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalSubjectAccessReviewPtrOutput)
}

type LocalSubjectAccessReviewPtrInput interface {
	pulumi.Input

	ToLocalSubjectAccessReviewPtrOutput() LocalSubjectAccessReviewPtrOutput
	ToLocalSubjectAccessReviewPtrOutputWithContext(ctx context.Context) LocalSubjectAccessReviewPtrOutput
}

type localSubjectAccessReviewPtrType LocalSubjectAccessReviewArgs

func (*localSubjectAccessReviewPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalSubjectAccessReview)(nil))
}

func (i *localSubjectAccessReviewPtrType) ToLocalSubjectAccessReviewPtrOutput() LocalSubjectAccessReviewPtrOutput {
	return i.ToLocalSubjectAccessReviewPtrOutputWithContext(context.Background())
}

func (i *localSubjectAccessReviewPtrType) ToLocalSubjectAccessReviewPtrOutputWithContext(ctx context.Context) LocalSubjectAccessReviewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalSubjectAccessReviewOutput).ToLocalSubjectAccessReviewPtrOutput()
}

// LocalSubjectAccessReviewArrayInput is an input type that accepts LocalSubjectAccessReviewArray and LocalSubjectAccessReviewArrayOutput values.
// You can construct a concrete instance of `LocalSubjectAccessReviewArrayInput` via:
//
//          LocalSubjectAccessReviewArray{ LocalSubjectAccessReviewArgs{...} }
type LocalSubjectAccessReviewArrayInput interface {
	pulumi.Input

	ToLocalSubjectAccessReviewArrayOutput() LocalSubjectAccessReviewArrayOutput
	ToLocalSubjectAccessReviewArrayOutputWithContext(context.Context) LocalSubjectAccessReviewArrayOutput
}

type LocalSubjectAccessReviewArray []LocalSubjectAccessReviewInput

func (LocalSubjectAccessReviewArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LocalSubjectAccessReview)(nil))
}

func (i LocalSubjectAccessReviewArray) ToLocalSubjectAccessReviewArrayOutput() LocalSubjectAccessReviewArrayOutput {
	return i.ToLocalSubjectAccessReviewArrayOutputWithContext(context.Background())
}

func (i LocalSubjectAccessReviewArray) ToLocalSubjectAccessReviewArrayOutputWithContext(ctx context.Context) LocalSubjectAccessReviewArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalSubjectAccessReviewArrayOutput)
}

// LocalSubjectAccessReviewMapInput is an input type that accepts LocalSubjectAccessReviewMap and LocalSubjectAccessReviewMapOutput values.
// You can construct a concrete instance of `LocalSubjectAccessReviewMapInput` via:
//
//          LocalSubjectAccessReviewMap{ "key": LocalSubjectAccessReviewArgs{...} }
type LocalSubjectAccessReviewMapInput interface {
	pulumi.Input

	ToLocalSubjectAccessReviewMapOutput() LocalSubjectAccessReviewMapOutput
	ToLocalSubjectAccessReviewMapOutputWithContext(context.Context) LocalSubjectAccessReviewMapOutput
}

type LocalSubjectAccessReviewMap map[string]LocalSubjectAccessReviewInput

func (LocalSubjectAccessReviewMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LocalSubjectAccessReview)(nil))
}

func (i LocalSubjectAccessReviewMap) ToLocalSubjectAccessReviewMapOutput() LocalSubjectAccessReviewMapOutput {
	return i.ToLocalSubjectAccessReviewMapOutputWithContext(context.Background())
}

func (i LocalSubjectAccessReviewMap) ToLocalSubjectAccessReviewMapOutputWithContext(ctx context.Context) LocalSubjectAccessReviewMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalSubjectAccessReviewMapOutput)
}

type LocalSubjectAccessReviewOutput struct {
	*pulumi.OutputState
}

func (LocalSubjectAccessReviewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalSubjectAccessReview)(nil))
}

func (o LocalSubjectAccessReviewOutput) ToLocalSubjectAccessReviewOutput() LocalSubjectAccessReviewOutput {
	return o
}

func (o LocalSubjectAccessReviewOutput) ToLocalSubjectAccessReviewOutputWithContext(ctx context.Context) LocalSubjectAccessReviewOutput {
	return o
}

func (o LocalSubjectAccessReviewOutput) ToLocalSubjectAccessReviewPtrOutput() LocalSubjectAccessReviewPtrOutput {
	return o.ToLocalSubjectAccessReviewPtrOutputWithContext(context.Background())
}

func (o LocalSubjectAccessReviewOutput) ToLocalSubjectAccessReviewPtrOutputWithContext(ctx context.Context) LocalSubjectAccessReviewPtrOutput {
	return o.ApplyT(func(v LocalSubjectAccessReview) *LocalSubjectAccessReview {
		return &v
	}).(LocalSubjectAccessReviewPtrOutput)
}

type LocalSubjectAccessReviewPtrOutput struct {
	*pulumi.OutputState
}

func (LocalSubjectAccessReviewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalSubjectAccessReview)(nil))
}

func (o LocalSubjectAccessReviewPtrOutput) ToLocalSubjectAccessReviewPtrOutput() LocalSubjectAccessReviewPtrOutput {
	return o
}

func (o LocalSubjectAccessReviewPtrOutput) ToLocalSubjectAccessReviewPtrOutputWithContext(ctx context.Context) LocalSubjectAccessReviewPtrOutput {
	return o
}

type LocalSubjectAccessReviewArrayOutput struct{ *pulumi.OutputState }

func (LocalSubjectAccessReviewArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LocalSubjectAccessReview)(nil))
}

func (o LocalSubjectAccessReviewArrayOutput) ToLocalSubjectAccessReviewArrayOutput() LocalSubjectAccessReviewArrayOutput {
	return o
}

func (o LocalSubjectAccessReviewArrayOutput) ToLocalSubjectAccessReviewArrayOutputWithContext(ctx context.Context) LocalSubjectAccessReviewArrayOutput {
	return o
}

func (o LocalSubjectAccessReviewArrayOutput) Index(i pulumi.IntInput) LocalSubjectAccessReviewOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LocalSubjectAccessReview {
		return vs[0].([]LocalSubjectAccessReview)[vs[1].(int)]
	}).(LocalSubjectAccessReviewOutput)
}

type LocalSubjectAccessReviewMapOutput struct{ *pulumi.OutputState }

func (LocalSubjectAccessReviewMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LocalSubjectAccessReview)(nil))
}

func (o LocalSubjectAccessReviewMapOutput) ToLocalSubjectAccessReviewMapOutput() LocalSubjectAccessReviewMapOutput {
	return o
}

func (o LocalSubjectAccessReviewMapOutput) ToLocalSubjectAccessReviewMapOutputWithContext(ctx context.Context) LocalSubjectAccessReviewMapOutput {
	return o
}

func (o LocalSubjectAccessReviewMapOutput) MapIndex(k pulumi.StringInput) LocalSubjectAccessReviewOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LocalSubjectAccessReview {
		return vs[0].(map[string]LocalSubjectAccessReview)[vs[1].(string)]
	}).(LocalSubjectAccessReviewOutput)
}

func init() {
	pulumi.RegisterOutputType(LocalSubjectAccessReviewOutput{})
	pulumi.RegisterOutputType(LocalSubjectAccessReviewPtrOutput{})
	pulumi.RegisterOutputType(LocalSubjectAccessReviewArrayOutput{})
	pulumi.RegisterOutputType(LocalSubjectAccessReviewMapOutput{})
}
