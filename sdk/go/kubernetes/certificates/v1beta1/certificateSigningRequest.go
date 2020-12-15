// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes a certificate signing request
type CertificateSigningRequest struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// The certificate request itself and any additional information.
	Spec CertificateSigningRequestSpecPtrOutput `pulumi:"spec"`
	// Derived information about the request.
	Status CertificateSigningRequestStatusPtrOutput `pulumi:"status"`
}

// NewCertificateSigningRequest registers a new resource with the given unique name, arguments, and options.
func NewCertificateSigningRequest(ctx *pulumi.Context,
	name string, args *CertificateSigningRequestArgs, opts ...pulumi.ResourceOption) (*CertificateSigningRequest, error) {
	if args == nil {
		args = &CertificateSigningRequestArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("certificates.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("CertificateSigningRequest")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:certificates.k8s.io/v1:CertificateSigningRequest"),
		},
	})
	opts = append(opts, aliases)
	var resource CertificateSigningRequest
	err := ctx.RegisterResource("kubernetes:certificates.k8s.io/v1beta1:CertificateSigningRequest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateSigningRequest gets an existing CertificateSigningRequest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateSigningRequest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateSigningRequestState, opts ...pulumi.ResourceOption) (*CertificateSigningRequest, error) {
	var resource CertificateSigningRequest
	err := ctx.ReadResource("kubernetes:certificates.k8s.io/v1beta1:CertificateSigningRequest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateSigningRequest resources.
type certificateSigningRequestState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// The certificate request itself and any additional information.
	Spec *CertificateSigningRequestSpec `pulumi:"spec"`
	// Derived information about the request.
	Status *CertificateSigningRequestStatus `pulumi:"status"`
}

type CertificateSigningRequestState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// The certificate request itself and any additional information.
	Spec CertificateSigningRequestSpecPtrInput
	// Derived information about the request.
	Status CertificateSigningRequestStatusPtrInput
}

func (CertificateSigningRequestState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateSigningRequestState)(nil)).Elem()
}

type certificateSigningRequestArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// The certificate request itself and any additional information.
	Spec *CertificateSigningRequestSpec `pulumi:"spec"`
}

// The set of arguments for constructing a CertificateSigningRequest resource.
type CertificateSigningRequestArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// The certificate request itself and any additional information.
	Spec CertificateSigningRequestSpecPtrInput
}

func (CertificateSigningRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateSigningRequestArgs)(nil)).Elem()
}

type CertificateSigningRequestInput interface {
	pulumi.Input

	ToCertificateSigningRequestOutput() CertificateSigningRequestOutput
	ToCertificateSigningRequestOutputWithContext(ctx context.Context) CertificateSigningRequestOutput
}

func (*CertificateSigningRequest) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateSigningRequest)(nil))
}

func (i *CertificateSigningRequest) ToCertificateSigningRequestOutput() CertificateSigningRequestOutput {
	return i.ToCertificateSigningRequestOutputWithContext(context.Background())
}

func (i *CertificateSigningRequest) ToCertificateSigningRequestOutputWithContext(ctx context.Context) CertificateSigningRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateSigningRequestOutput)
}

func (i *CertificateSigningRequest) ToCertificateSigningRequestPtrOutput() CertificateSigningRequestPtrOutput {
	return i.ToCertificateSigningRequestPtrOutputWithContext(context.Background())
}

func (i *CertificateSigningRequest) ToCertificateSigningRequestPtrOutputWithContext(ctx context.Context) CertificateSigningRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateSigningRequestPtrOutput)
}

type CertificateSigningRequestPtrInput interface {
	pulumi.Input

	ToCertificateSigningRequestPtrOutput() CertificateSigningRequestPtrOutput
	ToCertificateSigningRequestPtrOutputWithContext(ctx context.Context) CertificateSigningRequestPtrOutput
}

type CertificateSigningRequestOutput struct {
	*pulumi.OutputState
}

func (CertificateSigningRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateSigningRequest)(nil))
}

func (o CertificateSigningRequestOutput) ToCertificateSigningRequestOutput() CertificateSigningRequestOutput {
	return o
}

func (o CertificateSigningRequestOutput) ToCertificateSigningRequestOutputWithContext(ctx context.Context) CertificateSigningRequestOutput {
	return o
}

type CertificateSigningRequestPtrOutput struct {
	*pulumi.OutputState
}

func (CertificateSigningRequestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateSigningRequest)(nil))
}

func (o CertificateSigningRequestPtrOutput) ToCertificateSigningRequestPtrOutput() CertificateSigningRequestPtrOutput {
	return o
}

func (o CertificateSigningRequestPtrOutput) ToCertificateSigningRequestPtrOutputWithContext(ctx context.Context) CertificateSigningRequestPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CertificateSigningRequestOutput{})
	pulumi.RegisterOutputType(CertificateSigningRequestPtrOutput{})
}
