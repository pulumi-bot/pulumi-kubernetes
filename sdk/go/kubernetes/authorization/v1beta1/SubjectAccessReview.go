// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// SubjectAccessReview checks whether or not a user or group can perform an action.
type SubjectAccessReview struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`

	
	Metadata metaV1.ObjectMetaOutput `pulumi:"metadata"`

	// Spec holds information about the request being evaluated
	Spec SubjectAccessReviewSpecOutput `pulumi:"spec"`

	// Status is filled in by the server and indicates whether the request is allowed or not
	Status SubjectAccessReviewStatusOutput `pulumi:"status"`

}

// SubjectAccessReviewArgs is the set of arguments needed to create a SubjectAccessReview resource.
type SubjectAccessReviewArgs struct {
	// Spec holds information about the request being evaluated
	Spec SubjectAccessReviewSpecInput `pulumi:"spec"`

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringInput `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringInput `pulumi:"kind"`

	
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

}

// NewSubjectAccessReview creates a SubjectAccessReview resource with the given unique name, arguments, and options.
func NewSubjectAccessReview(ctx *pulumi.Context, name string, args *SubjectAccessReviewArgs, opts ...pulumi.ResourceOption) (*SubjectAccessReview, error) {
	inputs := map[string]pulumi.Input{}
	if args == nil || args.Spec == nil {
		return nil, errors.New("missing required argument 'Spec'")
	}
	if args != nil {
		args.ApiVersion = pulumi.String("authorization.k8s.io/v1beta1")
		args.Kind = pulumi.String("SubjectAccessReview")
		inputs["spec"] = args.Spec.ToSubjectAccessReviewSpecOutput()
		if i := args.ApiVersion; i != nil {
			inputs["apiVersion"] = i.ToStringOutput()
		}
		if i := args.Kind; i != nil {
			inputs["kind"] = i.ToStringOutput()
		}
		if i := args.Metadata; i != nil {
			inputs["metadata"] = i.ToObjectMetaOutput()
		}
	}
	var resource SubjectAccessReview
	err := ctx.RegisterResource("kubernetes:authorization.k8s.io/v1beta1:SubjectAccessReview", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubjectAccessReview gets an existing SubjectAccessReview resource's state with the given name and ID.
func GetSubjectAccessReview(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*SubjectAccessReview, error) {
	var resource SubjectAccessReview
	err := ctx.ReadResource("kubernetes:authorization.k8s.io/v1beta1:SubjectAccessReview", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

