// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// represents a scaling request for a resource.
type Scale struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`

	// Standard object metadata; More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata *metaV1.ObjectMeta `pulumi:"metadata"`

	// defines the behavior of the scale. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec *ScaleSpec `pulumi:"spec"`

	// current status of the scale. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	// Read-only.
	Status ScaleStatus `pulumi:"status"`

}

var _ScaleType = reflect.TypeOf((*Scale)(nil)).Elem()

// ScaleInput represents an input type that resolves to a Scale.
type ScaleInput interface {
	ElementType() reflect.Type

	ToScaleOutput() ScaleOutput
	ToScaleOutputWithContext(ctx context.Context) ScaleOutput
}

// ScaleArgs is a ScaleInput whose fields are all Input types.
type ScaleArgs struct {
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

	// Standard object metadata; More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// defines the behavior of the scale. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec ScaleSpecInput `pulumi:"spec"`

}

func (a ScaleArgs) ElementType() reflect.Type {
	return _ScaleType
}

func (a ScaleArgs) ToScaleOutput() ScaleOutput {
	return pulumi.ToOutput(a).(ScaleOutput)
}

func (a ScaleArgs) ToScaleOutputWithContext(ctx context.Context) ScaleOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ScaleOutput)
}

// ScaleOutput is an output type that resolves to a Input.
type ScaleOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ScaleOutput{}) }

func (ScaleOutput) ElementType() reflect.Type {
	return _ScaleType
}

func (o ScaleOutput) ApiVersion() pulumi.StringOutput {
	return o.Apply(func(v Scale) *string {
		return v.ApiVersion
	}).(pulumi.StringOutput)
}

func (o ScaleOutput) Kind() pulumi.StringOutput {
	return o.Apply(func(v Scale) *string {
		return v.Kind
	}).(pulumi.StringOutput)
}

func (o ScaleOutput) Metadata() metaV1.ObjectMetaOutput {
	return o.Apply(func(v Scale) *metaV1.ObjectMeta {
		return v.Metadata
	}).(metaV1.ObjectMetaOutput)
}

func (o ScaleOutput) Spec() ScaleSpecOutput {
	return o.Apply(func(v Scale) *ScaleSpec {
		return v.Spec
	}).(ScaleSpecOutput)
}

func (o ScaleOutput) Status() ScaleStatusOutput {
	return o.Apply(func(v Scale) ScaleStatus {
		return v.Status
	}).(ScaleStatusOutput)
}

