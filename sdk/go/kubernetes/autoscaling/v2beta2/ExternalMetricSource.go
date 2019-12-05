// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ExternalMetricSource indicates how to scale on a metric not associated with any Kubernetes object
// (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside
// of cluster).
type ExternalMetricSource struct {
	// metric identifies the target metric by name and selector
	Metric MetricIdentifier `pulumi:"metric"`

	// target specifies the target value for the given metric
	Target MetricTarget `pulumi:"target"`

}

var _ExternalMetricSourceType = reflect.TypeOf((*ExternalMetricSource)(nil)).Elem()

// ExternalMetricSourceInput represents an input type that resolves to a ExternalMetricSource.
type ExternalMetricSourceInput interface {
	ElementType() reflect.Type

	ToExternalMetricSourceOutput() ExternalMetricSourceOutput
	ToExternalMetricSourceOutputWithContext(ctx context.Context) ExternalMetricSourceOutput
}

// ExternalMetricSourceArgs is a ExternalMetricSourceInput whose fields are all Input types.
type ExternalMetricSourceArgs struct {
	// metric identifies the target metric by name and selector
	Metric MetricIdentifierInput `pulumi:"metric"`

	// target specifies the target value for the given metric
	Target MetricTargetInput `pulumi:"target"`

}

func (a ExternalMetricSourceArgs) ElementType() reflect.Type {
	return _ExternalMetricSourceType
}

func (a ExternalMetricSourceArgs) ToExternalMetricSourceOutput() ExternalMetricSourceOutput {
	return pulumi.ToOutput(a).(ExternalMetricSourceOutput)
}

func (a ExternalMetricSourceArgs) ToExternalMetricSourceOutputWithContext(ctx context.Context) ExternalMetricSourceOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ExternalMetricSourceOutput)
}

// ExternalMetricSourceOutput is an output type that resolves to a Input.
type ExternalMetricSourceOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ExternalMetricSourceOutput{}) }

func (ExternalMetricSourceOutput) ElementType() reflect.Type {
	return _ExternalMetricSourceType
}

func (o ExternalMetricSourceOutput) Metric() MetricIdentifierOutput {
	return o.Apply(func(v ExternalMetricSource) MetricIdentifier {
		return v.Metric
	}).(MetricIdentifierOutput)
}

func (o ExternalMetricSourceOutput) Target() MetricTargetOutput {
	return o.Apply(func(v ExternalMetricSource) MetricTarget {
		return v.Target
	}).(MetricTargetOutput)
}

