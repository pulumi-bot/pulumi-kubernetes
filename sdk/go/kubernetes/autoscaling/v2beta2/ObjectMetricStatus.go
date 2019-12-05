// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ObjectMetricStatus indicates the current value of a metric describing a kubernetes object (for
// example, hits-per-second on an Ingress object).
type ObjectMetricStatus struct {
	// current contains the current value for the given metric
	Current MetricValueStatus `pulumi:"current"`

	
	DescribedObject CrossVersionObjectReference `pulumi:"describedObject"`

	// metric identifies the target metric by name and selector
	Metric MetricIdentifier `pulumi:"metric"`

}

var _ObjectMetricStatusType = reflect.TypeOf((*ObjectMetricStatus)(nil)).Elem()

// ObjectMetricStatusInput represents an input type that resolves to a ObjectMetricStatus.
type ObjectMetricStatusInput interface {
	ElementType() reflect.Type

	ToObjectMetricStatusOutput() ObjectMetricStatusOutput
	ToObjectMetricStatusOutputWithContext(ctx context.Context) ObjectMetricStatusOutput
}

// ObjectMetricStatusArgs is a ObjectMetricStatusInput whose fields are all Input types.
type ObjectMetricStatusArgs struct {
	// current contains the current value for the given metric
	Current MetricValueStatusInput `pulumi:"current"`

	
	DescribedObject CrossVersionObjectReferenceInput `pulumi:"describedObject"`

	// metric identifies the target metric by name and selector
	Metric MetricIdentifierInput `pulumi:"metric"`

}

func (a ObjectMetricStatusArgs) ElementType() reflect.Type {
	return _ObjectMetricStatusType
}

func (a ObjectMetricStatusArgs) ToObjectMetricStatusOutput() ObjectMetricStatusOutput {
	return pulumi.ToOutput(a).(ObjectMetricStatusOutput)
}

func (a ObjectMetricStatusArgs) ToObjectMetricStatusOutputWithContext(ctx context.Context) ObjectMetricStatusOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ObjectMetricStatusOutput)
}

// ObjectMetricStatusOutput is an output type that resolves to a Input.
type ObjectMetricStatusOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ObjectMetricStatusOutput{}) }

func (ObjectMetricStatusOutput) ElementType() reflect.Type {
	return _ObjectMetricStatusType
}

func (o ObjectMetricStatusOutput) Current() MetricValueStatusOutput {
	return o.Apply(func(v ObjectMetricStatus) MetricValueStatus {
		return v.Current
	}).(MetricValueStatusOutput)
}

func (o ObjectMetricStatusOutput) DescribedObject() CrossVersionObjectReferenceOutput {
	return o.Apply(func(v ObjectMetricStatus) CrossVersionObjectReference {
		return v.DescribedObject
	}).(CrossVersionObjectReferenceOutput)
}

func (o ObjectMetricStatusOutput) Metric() MetricIdentifierOutput {
	return o.Apply(func(v ObjectMetricStatus) MetricIdentifier {
		return v.Metric
	}).(MetricIdentifierOutput)
}

