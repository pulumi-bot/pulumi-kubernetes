// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// DaemonSetCondition describes the state of a DaemonSet at a certain point.
type DaemonSetCondition struct {
	// Last time the condition transitioned from one status to another.
	LastTransitionTime *string `pulumi:"lastTransitionTime"`

	// A human readable message indicating details about the transition.
	Message *string `pulumi:"message"`

	// The reason for the condition's last transition.
	Reason *string `pulumi:"reason"`

	// Status of the condition, one of True, False, Unknown.
	Status string `pulumi:"status"`

	// Type of DaemonSet condition.
	Type string `pulumi:"type"`

}

var _DaemonSetConditionType = reflect.TypeOf((*DaemonSetCondition)(nil)).Elem()

// DaemonSetConditionInput represents an input type that resolves to a DaemonSetCondition.
type DaemonSetConditionInput interface {
	ElementType() reflect.Type

	ToDaemonSetConditionOutput() DaemonSetConditionOutput
	ToDaemonSetConditionOutputWithContext(ctx context.Context) DaemonSetConditionOutput
}

// DaemonSetConditionArgs is a DaemonSetConditionInput whose fields are all Input types.
type DaemonSetConditionArgs struct {
	// Status of the condition, one of True, False, Unknown.
	Status pulumi.StringInput `pulumi:"status"`

	// Type of DaemonSet condition.
	Type pulumi.StringInput `pulumi:"type"`

	// Last time the condition transitioned from one status to another.
	LastTransitionTime pulumi.StringInput `pulumi:"lastTransitionTime"`

	// A human readable message indicating details about the transition.
	Message pulumi.StringInput `pulumi:"message"`

	// The reason for the condition's last transition.
	Reason pulumi.StringInput `pulumi:"reason"`

}

func (a DaemonSetConditionArgs) ElementType() reflect.Type {
	return _DaemonSetConditionType
}

func (a DaemonSetConditionArgs) ToDaemonSetConditionOutput() DaemonSetConditionOutput {
	return pulumi.ToOutput(a).(DaemonSetConditionOutput)
}

func (a DaemonSetConditionArgs) ToDaemonSetConditionOutputWithContext(ctx context.Context) DaemonSetConditionOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DaemonSetConditionOutput)
}

// DaemonSetConditionOutput is an output type that resolves to a Input.
type DaemonSetConditionOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(DaemonSetConditionOutput{}) }

func (DaemonSetConditionOutput) ElementType() reflect.Type {
	return _DaemonSetConditionType
}

func (o DaemonSetConditionOutput) LastTransitionTime() pulumi.StringOutput {
	return o.Apply(func(v DaemonSetCondition) *string {
		return v.LastTransitionTime
	}).(pulumi.StringOutput)
}

func (o DaemonSetConditionOutput) Message() pulumi.StringOutput {
	return o.Apply(func(v DaemonSetCondition) *string {
		return v.Message
	}).(pulumi.StringOutput)
}

func (o DaemonSetConditionOutput) Reason() pulumi.StringOutput {
	return o.Apply(func(v DaemonSetCondition) *string {
		return v.Reason
	}).(pulumi.StringOutput)
}

func (o DaemonSetConditionOutput) Status() pulumi.StringOutput {
	return o.Apply(func(v DaemonSetCondition) string {
		return v.Status
	}).(pulumi.StringOutput)
}

func (o DaemonSetConditionOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v DaemonSetCondition) string {
		return v.Type
	}).(pulumi.StringOutput)
}

var _DaemonSetConditionArrayType = reflect.TypeOf((*[]DaemonSetCondition)(nil)).Elem()

type DaemonSetConditionArrayInput interface {
	ElementType() reflect.Type

	ToDaemonSetConditionArrayOutput() DaemonSetConditionArrayOutput
	ToDaemonSetConditionArrayOutputWithContext(ctx context.Context) DaemonSetConditionArrayOutput
}

type DaemonSetConditionArray []DaemonSetConditionInput

func (a DaemonSetConditionArray) ElementType() reflect.Type {
	return _DaemonSetConditionArrayType
}

func (a DaemonSetConditionArray) ToDaemonSetConditionArrayOutput() DaemonSetConditionArrayOutput {
	return pulumi.ToOutput(a).(DaemonSetConditionArrayOutput)
}

func (a DaemonSetConditionArray) ToDaemonSetConditionArrayOutputWithContext(ctx context.Context) DaemonSetConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DaemonSetConditionArrayOutput)
}

type DaemonSetConditionArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(DaemonSetConditionArrayOutput{}) }

func (DaemonSetConditionArrayOutput) ElementType() reflect.Type {
	return _DaemonSetConditionArrayType
}

func (o DaemonSetConditionArrayOutput) Index(i pulumi.IntInput) DaemonSetConditionOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) DaemonSetCondition {
		return vs[0].([]DaemonSetCondition)[vs[1].(int)]
	}).(DaemonSetConditionOutput)
}
