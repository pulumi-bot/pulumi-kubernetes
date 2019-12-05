// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ResourceMetricStatus indicates the current value of a resource metric known to Kubernetes, as
// specified in requests and limits, describing each pod in the current scale target (e.g. CPU or
// memory).  Such metrics are built in to Kubernetes, and have special scaling options on top of
// those available to normal per-pod metrics using the "pods" source.
type ResourceMetricStatus struct {
	// current contains the current value for the given metric
	Current MetricValueStatus `pulumi:"current"`

	// Name is the name of the resource in question.
	Name string `pulumi:"name"`

}

var _ResourceMetricStatusType = reflect.TypeOf((*ResourceMetricStatus)(nil)).Elem()

// ResourceMetricStatusInput represents an input type that resolves to a ResourceMetricStatus.
type ResourceMetricStatusInput interface {
	ElementType() reflect.Type

	ToResourceMetricStatusOutput() ResourceMetricStatusOutput
	ToResourceMetricStatusOutputWithContext(ctx context.Context) ResourceMetricStatusOutput
}

// ResourceMetricStatusArgs is a ResourceMetricStatusInput whose fields are all Input types.
type ResourceMetricStatusArgs struct {
	// current contains the current value for the given metric
	Current MetricValueStatusInput `pulumi:"current"`

	// Name is the name of the resource in question.
	Name pulumi.StringInput `pulumi:"name"`

}

func (a ResourceMetricStatusArgs) ElementType() reflect.Type {
	return _ResourceMetricStatusType
}

func (a ResourceMetricStatusArgs) ToResourceMetricStatusOutput() ResourceMetricStatusOutput {
	return pulumi.ToOutput(a).(ResourceMetricStatusOutput)
}

func (a ResourceMetricStatusArgs) ToResourceMetricStatusOutputWithContext(ctx context.Context) ResourceMetricStatusOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ResourceMetricStatusOutput)
}

// ResourceMetricStatusOutput is an output type that resolves to a Input.
type ResourceMetricStatusOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ResourceMetricStatusOutput{}) }

func (ResourceMetricStatusOutput) ElementType() reflect.Type {
	return _ResourceMetricStatusType
}

func (o ResourceMetricStatusOutput) Current() MetricValueStatusOutput {
	return o.Apply(func(v ResourceMetricStatus) MetricValueStatus {
		return v.Current
	}).(MetricValueStatusOutput)
}

func (o ResourceMetricStatusOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v ResourceMetricStatus) string {
		return v.Name
	}).(pulumi.StringOutput)
}

