// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// MetricSpec specifies how to scale based on a single metric (only `type` and one other matching
// field should be set at once).
type MetricSpec struct {
	// external refers to a global metric that is not associated with any Kubernetes object. It allows
	// autoscaling based on information coming from components running outside of cluster (for example
	// length of queue in cloud messaging service, or QPS from loadbalancer running outside of
	// cluster).
	External *ExternalMetricSource `pulumi:"external"`

	// object refers to a metric describing a single kubernetes object (for example, hits-per-second on
	// an Ingress object).
	Object *ObjectMetricSource `pulumi:"object"`

	// pods refers to a metric describing each pod in the current scale target (for example,
	// transactions-processed-per-second).  The values will be averaged together before being compared
	// to the target value.
	Pods *PodsMetricSource `pulumi:"pods"`

	// resource refers to a resource metric (such as those specified in requests and limits) known to
	// Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics
	// are built in to Kubernetes, and have special scaling options on top of those available to normal
	// per-pod metrics using the "pods" source.
	Resource *ResourceMetricSource `pulumi:"resource"`

	// type is the type of metric source.  It should be one of "Object", "Pods" or "Resource", each
	// mapping to a matching field in the object.
	Type string `pulumi:"type"`

}

var _MetricSpecType = reflect.TypeOf((*MetricSpec)(nil)).Elem()

// MetricSpecInput represents an input type that resolves to a MetricSpec.
type MetricSpecInput interface {
	ElementType() reflect.Type

	ToMetricSpecOutput() MetricSpecOutput
	ToMetricSpecOutputWithContext(ctx context.Context) MetricSpecOutput
}

// MetricSpecArgs is a MetricSpecInput whose fields are all Input types.
type MetricSpecArgs struct {
	// type is the type of metric source.  It should be one of "Object", "Pods" or "Resource", each
	// mapping to a matching field in the object.
	Type pulumi.StringInput `pulumi:"type"`

	// external refers to a global metric that is not associated with any Kubernetes object. It allows
	// autoscaling based on information coming from components running outside of cluster (for example
	// length of queue in cloud messaging service, or QPS from loadbalancer running outside of
	// cluster).
	External ExternalMetricSourceInput `pulumi:"external"`

	// object refers to a metric describing a single kubernetes object (for example, hits-per-second on
	// an Ingress object).
	Object ObjectMetricSourceInput `pulumi:"object"`

	// pods refers to a metric describing each pod in the current scale target (for example,
	// transactions-processed-per-second).  The values will be averaged together before being compared
	// to the target value.
	Pods PodsMetricSourceInput `pulumi:"pods"`

	// resource refers to a resource metric (such as those specified in requests and limits) known to
	// Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics
	// are built in to Kubernetes, and have special scaling options on top of those available to normal
	// per-pod metrics using the "pods" source.
	Resource ResourceMetricSourceInput `pulumi:"resource"`

}

func (a MetricSpecArgs) ElementType() reflect.Type {
	return _MetricSpecType
}

func (a MetricSpecArgs) ToMetricSpecOutput() MetricSpecOutput {
	return pulumi.ToOutput(a).(MetricSpecOutput)
}

func (a MetricSpecArgs) ToMetricSpecOutputWithContext(ctx context.Context) MetricSpecOutput {
	return pulumi.ToOutputWithContext(ctx, a).(MetricSpecOutput)
}

// MetricSpecOutput is an output type that resolves to a Input.
type MetricSpecOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(MetricSpecOutput{}) }

func (MetricSpecOutput) ElementType() reflect.Type {
	return _MetricSpecType
}

func (o MetricSpecOutput) External() ExternalMetricSourceOutput {
	return o.Apply(func(v MetricSpec) *ExternalMetricSource {
		return v.External
	}).(ExternalMetricSourceOutput)
}

func (o MetricSpecOutput) Object() ObjectMetricSourceOutput {
	return o.Apply(func(v MetricSpec) *ObjectMetricSource {
		return v.Object
	}).(ObjectMetricSourceOutput)
}

func (o MetricSpecOutput) Pods() PodsMetricSourceOutput {
	return o.Apply(func(v MetricSpec) *PodsMetricSource {
		return v.Pods
	}).(PodsMetricSourceOutput)
}

func (o MetricSpecOutput) Resource() ResourceMetricSourceOutput {
	return o.Apply(func(v MetricSpec) *ResourceMetricSource {
		return v.Resource
	}).(ResourceMetricSourceOutput)
}

func (o MetricSpecOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v MetricSpec) string {
		return v.Type
	}).(pulumi.StringOutput)
}

var _MetricSpecArrayType = reflect.TypeOf((*[]MetricSpec)(nil)).Elem()

type MetricSpecArrayInput interface {
	ElementType() reflect.Type

	ToMetricSpecArrayOutput() MetricSpecArrayOutput
	ToMetricSpecArrayOutputWithContext(ctx context.Context) MetricSpecArrayOutput
}

type MetricSpecArray []MetricSpecInput

func (a MetricSpecArray) ElementType() reflect.Type {
	return _MetricSpecArrayType
}

func (a MetricSpecArray) ToMetricSpecArrayOutput() MetricSpecArrayOutput {
	return pulumi.ToOutput(a).(MetricSpecArrayOutput)
}

func (a MetricSpecArray) ToMetricSpecArrayOutputWithContext(ctx context.Context) MetricSpecArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(MetricSpecArrayOutput)
}

type MetricSpecArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(MetricSpecArrayOutput{}) }

func (MetricSpecArrayOutput) ElementType() reflect.Type {
	return _MetricSpecArrayType
}

func (o MetricSpecArrayOutput) Index(i pulumi.IntInput) MetricSpecOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) MetricSpec {
		return vs[0].([]MetricSpec)[vs[1].(int)]
	}).(MetricSpecOutput)
}
