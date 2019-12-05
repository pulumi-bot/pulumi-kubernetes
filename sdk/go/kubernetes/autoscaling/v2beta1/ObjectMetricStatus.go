// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// ObjectMetricStatus indicates the current value of a metric describing a kubernetes object (for
// example, hits-per-second on an Ingress object).
type ObjectMetricStatus struct {
	// averageValue is the current value of the average of the metric across all relevant pods (as a
	// quantity)
	AverageValue *string `pulumi:"averageValue"`

	// currentValue is the current value of the metric (as a quantity).
	CurrentValue string `pulumi:"currentValue"`

	// metricName is the name of the metric in question.
	MetricName string `pulumi:"metricName"`

	// selector is the string-encoded form of a standard kubernetes label selector for the given metric
	// When set in the ObjectMetricSource, it is passed as an additional parameter to the metrics
	// server for more specific metrics scoping. When unset, just the metricName will be used to gather
	// metrics.
	Selector *metaV1.LabelSelector `pulumi:"selector"`

	// target is the described Kubernetes object.
	Target CrossVersionObjectReference `pulumi:"target"`

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
	// currentValue is the current value of the metric (as a quantity).
	CurrentValue pulumi.StringInput `pulumi:"currentValue"`

	// metricName is the name of the metric in question.
	MetricName pulumi.StringInput `pulumi:"metricName"`

	// target is the described Kubernetes object.
	Target CrossVersionObjectReferenceInput `pulumi:"target"`

	// averageValue is the current value of the average of the metric across all relevant pods (as a
	// quantity)
	AverageValue pulumi.StringInput `pulumi:"averageValue"`

	// selector is the string-encoded form of a standard kubernetes label selector for the given metric
	// When set in the ObjectMetricSource, it is passed as an additional parameter to the metrics
	// server for more specific metrics scoping. When unset, just the metricName will be used to gather
	// metrics.
	Selector metaV1.LabelSelectorInput `pulumi:"selector"`

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

func (o ObjectMetricStatusOutput) AverageValue() pulumi.StringOutput {
	return o.Apply(func(v ObjectMetricStatus) *string {
		return v.AverageValue
	}).(pulumi.StringOutput)
}

func (o ObjectMetricStatusOutput) CurrentValue() pulumi.StringOutput {
	return o.Apply(func(v ObjectMetricStatus) string {
		return v.CurrentValue
	}).(pulumi.StringOutput)
}

func (o ObjectMetricStatusOutput) MetricName() pulumi.StringOutput {
	return o.Apply(func(v ObjectMetricStatus) string {
		return v.MetricName
	}).(pulumi.StringOutput)
}

func (o ObjectMetricStatusOutput) Selector() metaV1.LabelSelectorOutput {
	return o.Apply(func(v ObjectMetricStatus) *metaV1.LabelSelector {
		return v.Selector
	}).(metaV1.LabelSelectorOutput)
}

func (o ObjectMetricStatusOutput) Target() CrossVersionObjectReferenceOutput {
	return o.Apply(func(v ObjectMetricStatus) CrossVersionObjectReference {
		return v.Target
	}).(CrossVersionObjectReferenceOutput)
}

