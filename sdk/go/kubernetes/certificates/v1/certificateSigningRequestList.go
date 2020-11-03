// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// CertificateSigningRequestList is a collection of CertificateSigningRequest objects
type CertificateSigningRequestList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// items is a collection of CertificateSigningRequest objects
	Items CertificateSigningRequestTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput   `pulumi:"kind"`
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewCertificateSigningRequestList registers a new resource with the given unique name, arguments, and options.
func NewCertificateSigningRequestList(ctx *pulumi.Context,
	name string, args *CertificateSigningRequestListArgs, opts ...pulumi.ResourceOption) (*CertificateSigningRequestList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("certificates.k8s.io/v1")
	args.Kind = pulumi.StringPtr("CertificateSigningRequestList")
	var resource CertificateSigningRequestList
	err := ctx.RegisterResource("kubernetes:certificates.k8s.io/v1:CertificateSigningRequestList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateSigningRequestList gets an existing CertificateSigningRequestList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateSigningRequestList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateSigningRequestListState, opts ...pulumi.ResourceOption) (*CertificateSigningRequestList, error) {
	var resource CertificateSigningRequestList
	err := ctx.ReadResource("kubernetes:certificates.k8s.io/v1:CertificateSigningRequestList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateSigningRequestList resources.
type certificateSigningRequestListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is a collection of CertificateSigningRequest objects
	Items []CertificateSigningRequestType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string          `pulumi:"kind"`
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type CertificateSigningRequestListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is a collection of CertificateSigningRequest objects
	Items CertificateSigningRequestTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ListMetaPtrInput
}

func (CertificateSigningRequestListState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateSigningRequestListState)(nil)).Elem()
}

type certificateSigningRequestListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is a collection of CertificateSigningRequest objects
	Items []CertificateSigningRequestType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string          `pulumi:"kind"`
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a CertificateSigningRequestList resource.
type CertificateSigningRequestListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is a collection of CertificateSigningRequest objects
	Items CertificateSigningRequestTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ListMetaPtrInput
}

func (CertificateSigningRequestListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateSigningRequestListArgs)(nil)).Elem()
}
