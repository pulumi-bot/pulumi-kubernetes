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

// TokenRequest requests a token for a given service account.
type TokenRequest struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput      `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput  `pulumi:"metadata"`
	Spec     TokenRequestSpecOutput      `pulumi:"spec"`
	Status   TokenRequestStatusPtrOutput `pulumi:"status"`
}

// NewTokenRequest registers a new resource with the given unique name, arguments, and options.
func NewTokenRequest(ctx *pulumi.Context,
	name string, args *TokenRequestArgs, opts ...pulumi.ResourceOption) (*TokenRequest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	args.ApiVersion = pulumi.StringPtr("authentication.k8s.io/v1")
	args.Kind = pulumi.StringPtr("TokenRequest")
	var resource TokenRequest
	err := ctx.RegisterResource("kubernetes:authentication.k8s.io/v1:TokenRequest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTokenRequest gets an existing TokenRequest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTokenRequest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TokenRequestState, opts ...pulumi.ResourceOption) (*TokenRequest, error) {
	var resource TokenRequest
	err := ctx.ReadResource("kubernetes:authentication.k8s.io/v1:TokenRequest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TokenRequest resources.
type tokenRequestState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string             `pulumi:"kind"`
	Metadata *metav1.ObjectMeta  `pulumi:"metadata"`
	Spec     *TokenRequestSpec   `pulumi:"spec"`
	Status   *TokenRequestStatus `pulumi:"status"`
}

type TokenRequestState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	Spec     TokenRequestSpecPtrInput
	Status   TokenRequestStatusPtrInput
}

func (TokenRequestState) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenRequestState)(nil)).Elem()
}

type tokenRequestArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	Spec     TokenRequestSpec   `pulumi:"spec"`
}

// The set of arguments for constructing a TokenRequest resource.
type TokenRequestArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	Spec     TokenRequestSpecInput
}

func (TokenRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenRequestArgs)(nil)).Elem()
}

type TokenRequestInput interface {
	pulumi.Input

	ToTokenRequestOutput() TokenRequestOutput
	ToTokenRequestOutputWithContext(ctx context.Context) TokenRequestOutput
}

type TokenRequestPtrInput interface {
	pulumi.Input

	ToTokenRequestPtrOutput() TokenRequestPtrOutput
	ToTokenRequestPtrOutputWithContext(ctx context.Context) TokenRequestPtrOutput
}

func (TokenRequest) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenRequest)(nil)).Elem()
}

func (i TokenRequest) ToTokenRequestOutput() TokenRequestOutput {
	return i.ToTokenRequestOutputWithContext(context.Background())
}

func (i TokenRequest) ToTokenRequestOutputWithContext(ctx context.Context) TokenRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenRequestOutput)
}

func (i TokenRequest) ToTokenRequestPtrOutput() TokenRequestPtrOutput {
	return i.ToTokenRequestPtrOutputWithContext(context.Background())
}

func (i TokenRequest) ToTokenRequestPtrOutputWithContext(ctx context.Context) TokenRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenRequestPtrOutput)
}

type TokenRequestOutput struct {
	*pulumi.OutputState
}

func (TokenRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenRequestOutput)(nil)).Elem()
}

func (o TokenRequestOutput) ToTokenRequestOutput() TokenRequestOutput {
	return o
}

func (o TokenRequestOutput) ToTokenRequestOutputWithContext(ctx context.Context) TokenRequestOutput {
	return o
}

type TokenRequestPtrOutput struct {
	*pulumi.OutputState
}

func (TokenRequestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenRequest)(nil)).Elem()
}

func (o TokenRequestPtrOutput) ToTokenRequestPtrOutput() TokenRequestPtrOutput {
	return o
}

func (o TokenRequestPtrOutput) ToTokenRequestPtrOutputWithContext(ctx context.Context) TokenRequestPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TokenRequestOutput{})
	pulumi.RegisterOutputType(TokenRequestPtrOutput{})
}
