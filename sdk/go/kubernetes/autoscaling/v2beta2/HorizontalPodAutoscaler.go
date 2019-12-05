// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// HorizontalPodAutoscaler is the configuration for a horizontal pod autoscaler, which automatically
// manages the replica count of any resource implementing the scale subresource based on the metrics
// specified.
type HorizontalPodAutoscaler struct {
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

	// metadata is the standard object metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaOutput `pulumi:"metadata"`

	// spec is the specification for the behaviour of the autoscaler. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec HorizontalPodAutoscalerSpecOutput `pulumi:"spec"`

	// status is the current information about the autoscaler.
	Status HorizontalPodAutoscalerStatusOutput `pulumi:"status"`

}

// HorizontalPodAutoscalerArgs is the set of arguments needed to create a HorizontalPodAutoscaler resource.
type HorizontalPodAutoscalerArgs struct {
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

	// metadata is the standard object metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// spec is the specification for the behaviour of the autoscaler. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec HorizontalPodAutoscalerSpecInput `pulumi:"spec"`

}

// NewHorizontalPodAutoscaler creates a HorizontalPodAutoscaler resource with the given unique name, arguments, and options.
func NewHorizontalPodAutoscaler(ctx *pulumi.Context, name string, args *HorizontalPodAutoscalerArgs, opts ...pulumi.ResourceOption) (*HorizontalPodAutoscaler, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		args.ApiVersion = pulumi.String("autoscaling/v2beta2")
		args.Kind = pulumi.String("HorizontalPodAutoscaler")
		if i := args.ApiVersion; i != nil {
			inputs["apiVersion"] = i.ToStringOutput()
		}
		if i := args.Kind; i != nil {
			inputs["kind"] = i.ToStringOutput()
		}
		if i := args.Metadata; i != nil {
			inputs["metadata"] = i.ToObjectMetaOutput()
		}
		if i := args.Spec; i != nil {
			inputs["spec"] = i.ToHorizontalPodAutoscalerSpecOutput()
		}
	}
	var resource HorizontalPodAutoscaler
	err := ctx.RegisterResource("kubernetes:autoscaling/v2beta2:HorizontalPodAutoscaler", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHorizontalPodAutoscaler gets an existing HorizontalPodAutoscaler resource's state with the given name and ID.
func GetHorizontalPodAutoscaler(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*HorizontalPodAutoscaler, error) {
	var resource HorizontalPodAutoscaler
	err := ctx.ReadResource("kubernetes:autoscaling/v2beta2:HorizontalPodAutoscaler", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// HorizontalPodAutoscaler is the configuration for a horizontal pod autoscaler, which automatically
// manages the replica count of any resource implementing the scale subresource based on the metrics
// specified.
type HorizontalPodAutoscalerProperty struct {
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

	// metadata is the standard object metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metaV1.ObjectMeta `pulumi:"metadata"`

	// spec is the specification for the behaviour of the autoscaler. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec *HorizontalPodAutoscalerSpec `pulumi:"spec"`

	// status is the current information about the autoscaler.
	Status HorizontalPodAutoscalerStatus `pulumi:"status"`

}

var _HorizontalPodAutoscalerPropertyType = reflect.TypeOf((*HorizontalPodAutoscalerProperty)(nil)).Elem()

// HorizontalPodAutoscalerPropertyInput represents an input type that resolves to a HorizontalPodAutoscalerProperty.
type HorizontalPodAutoscalerPropertyInput interface {
	ElementType() reflect.Type

	ToHorizontalPodAutoscalerPropertyOutput() HorizontalPodAutoscalerPropertyOutput
	ToHorizontalPodAutoscalerPropertyOutputWithContext(ctx context.Context) HorizontalPodAutoscalerPropertyOutput
}

// HorizontalPodAutoscalerPropertyArgs is a HorizontalPodAutoscalerPropertyInput whose fields are all Input types.
type HorizontalPodAutoscalerPropertyArgs struct {
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

	// metadata is the standard object metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// spec is the specification for the behaviour of the autoscaler. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	Spec HorizontalPodAutoscalerSpecInput `pulumi:"spec"`

}

func (a HorizontalPodAutoscalerPropertyArgs) ElementType() reflect.Type {
	return _HorizontalPodAutoscalerPropertyType
}

func (a HorizontalPodAutoscalerPropertyArgs) ToHorizontalPodAutoscalerPropertyOutput() HorizontalPodAutoscalerPropertyOutput {
	return pulumi.ToOutput(a).(HorizontalPodAutoscalerPropertyOutput)
}

func (a HorizontalPodAutoscalerPropertyArgs) ToHorizontalPodAutoscalerPropertyOutputWithContext(ctx context.Context) HorizontalPodAutoscalerPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, a).(HorizontalPodAutoscalerPropertyOutput)
}

// HorizontalPodAutoscalerPropertyOutput is an output type that resolves to a Input.
type HorizontalPodAutoscalerPropertyOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(HorizontalPodAutoscalerPropertyOutput{}) }

func (HorizontalPodAutoscalerPropertyOutput) ElementType() reflect.Type {
	return _HorizontalPodAutoscalerPropertyType
}

func (o HorizontalPodAutoscalerPropertyOutput) ApiVersion() pulumi.StringOutput {
	return o.Apply(func(v HorizontalPodAutoscalerProperty) *string {
		return v.ApiVersion
	}).(pulumi.StringOutput)
}

func (o HorizontalPodAutoscalerPropertyOutput) Kind() pulumi.StringOutput {
	return o.Apply(func(v HorizontalPodAutoscalerProperty) *string {
		return v.Kind
	}).(pulumi.StringOutput)
}

func (o HorizontalPodAutoscalerPropertyOutput) Metadata() metaV1.ObjectMetaOutput {
	return o.Apply(func(v HorizontalPodAutoscalerProperty) *metaV1.ObjectMeta {
		return v.Metadata
	}).(metaV1.ObjectMetaOutput)
}

func (o HorizontalPodAutoscalerPropertyOutput) Spec() HorizontalPodAutoscalerSpecOutput {
	return o.Apply(func(v HorizontalPodAutoscalerProperty) *HorizontalPodAutoscalerSpec {
		return v.Spec
	}).(HorizontalPodAutoscalerSpecOutput)
}

func (o HorizontalPodAutoscalerPropertyOutput) Status() HorizontalPodAutoscalerStatusOutput {
	return o.Apply(func(v HorizontalPodAutoscalerProperty) HorizontalPodAutoscalerStatus {
		return v.Status
	}).(HorizontalPodAutoscalerStatusOutput)
}

var _HorizontalPodAutoscalerPropertyArrayType = reflect.TypeOf((*[]HorizontalPodAutoscalerProperty)(nil)).Elem()

type HorizontalPodAutoscalerPropertyArrayInput interface {
	ElementType() reflect.Type

	ToHorizontalPodAutoscalerPropertyArrayOutput() HorizontalPodAutoscalerPropertyArrayOutput
	ToHorizontalPodAutoscalerPropertyArrayOutputWithContext(ctx context.Context) HorizontalPodAutoscalerPropertyArrayOutput
}

type HorizontalPodAutoscalerPropertyArray []HorizontalPodAutoscalerPropertyInput

func (a HorizontalPodAutoscalerPropertyArray) ElementType() reflect.Type {
	return _HorizontalPodAutoscalerPropertyArrayType
}

func (a HorizontalPodAutoscalerPropertyArray) ToHorizontalPodAutoscalerPropertyArrayOutput() HorizontalPodAutoscalerPropertyArrayOutput {
	return pulumi.ToOutput(a).(HorizontalPodAutoscalerPropertyArrayOutput)
}

func (a HorizontalPodAutoscalerPropertyArray) ToHorizontalPodAutoscalerPropertyArrayOutputWithContext(ctx context.Context) HorizontalPodAutoscalerPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(HorizontalPodAutoscalerPropertyArrayOutput)
}

type HorizontalPodAutoscalerPropertyArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(HorizontalPodAutoscalerPropertyArrayOutput{}) }

func (HorizontalPodAutoscalerPropertyArrayOutput) ElementType() reflect.Type {
	return _HorizontalPodAutoscalerPropertyArrayType
}

func (o HorizontalPodAutoscalerPropertyArrayOutput) Index(i pulumi.IntInput) HorizontalPodAutoscalerPropertyOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) HorizontalPodAutoscalerProperty {
		return vs[0].([]HorizontalPodAutoscalerProperty)[vs[1].(int)]
	}).(HorizontalPodAutoscalerPropertyOutput)
}
