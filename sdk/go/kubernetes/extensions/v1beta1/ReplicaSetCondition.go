// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ReplicaSetCondition describes the state of a replica set at a certain point.
type ReplicaSetCondition struct {
	// The last time the condition transitioned from one status to another.
	LastTransitionTime *string `pulumi:"lastTransitionTime"`

	// A human readable message indicating details about the transition.
	Message *string `pulumi:"message"`

	// The reason for the condition's last transition.
	Reason *string `pulumi:"reason"`

	// Status of the condition, one of True, False, Unknown.
	Status string `pulumi:"status"`

	// Type of replica set condition.
	Type string `pulumi:"type"`

}

var _ReplicaSetConditionType = reflect.TypeOf((*ReplicaSetCondition)(nil)).Elem()

// ReplicaSetConditionInput represents an input type that resolves to a ReplicaSetCondition.
type ReplicaSetConditionInput interface {
	ElementType() reflect.Type

	ToReplicaSetConditionOutput() ReplicaSetConditionOutput
	ToReplicaSetConditionOutputWithContext(ctx context.Context) ReplicaSetConditionOutput
}

// ReplicaSetConditionArgs is a ReplicaSetConditionInput whose fields are all Input types.
type ReplicaSetConditionArgs struct {
	// Status of the condition, one of True, False, Unknown.
	Status pulumi.StringInput `pulumi:"status"`

	// Type of replica set condition.
	Type pulumi.StringInput `pulumi:"type"`

	// The last time the condition transitioned from one status to another.
	LastTransitionTime pulumi.StringInput `pulumi:"lastTransitionTime"`

	// A human readable message indicating details about the transition.
	Message pulumi.StringInput `pulumi:"message"`

	// The reason for the condition's last transition.
	Reason pulumi.StringInput `pulumi:"reason"`

}

func (a ReplicaSetConditionArgs) ElementType() reflect.Type {
	return _ReplicaSetConditionType
}

func (a ReplicaSetConditionArgs) ToReplicaSetConditionOutput() ReplicaSetConditionOutput {
	return pulumi.ToOutput(a).(ReplicaSetConditionOutput)
}

func (a ReplicaSetConditionArgs) ToReplicaSetConditionOutputWithContext(ctx context.Context) ReplicaSetConditionOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ReplicaSetConditionOutput)
}

// ReplicaSetConditionOutput is an output type that resolves to a Input.
type ReplicaSetConditionOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ReplicaSetConditionOutput{}) }

func (ReplicaSetConditionOutput) ElementType() reflect.Type {
	return _ReplicaSetConditionType
}

func (o ReplicaSetConditionOutput) LastTransitionTime() pulumi.StringOutput {
	return o.Apply(func(v ReplicaSetCondition) *string {
		return v.LastTransitionTime
	}).(pulumi.StringOutput)
}

func (o ReplicaSetConditionOutput) Message() pulumi.StringOutput {
	return o.Apply(func(v ReplicaSetCondition) *string {
		return v.Message
	}).(pulumi.StringOutput)
}

func (o ReplicaSetConditionOutput) Reason() pulumi.StringOutput {
	return o.Apply(func(v ReplicaSetCondition) *string {
		return v.Reason
	}).(pulumi.StringOutput)
}

func (o ReplicaSetConditionOutput) Status() pulumi.StringOutput {
	return o.Apply(func(v ReplicaSetCondition) string {
		return v.Status
	}).(pulumi.StringOutput)
}

func (o ReplicaSetConditionOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v ReplicaSetCondition) string {
		return v.Type
	}).(pulumi.StringOutput)
}

var _ReplicaSetConditionArrayType = reflect.TypeOf((*[]ReplicaSetCondition)(nil)).Elem()

type ReplicaSetConditionArrayInput interface {
	ElementType() reflect.Type

	ToReplicaSetConditionArrayOutput() ReplicaSetConditionArrayOutput
	ToReplicaSetConditionArrayOutputWithContext(ctx context.Context) ReplicaSetConditionArrayOutput
}

type ReplicaSetConditionArray []ReplicaSetConditionInput

func (a ReplicaSetConditionArray) ElementType() reflect.Type {
	return _ReplicaSetConditionArrayType
}

func (a ReplicaSetConditionArray) ToReplicaSetConditionArrayOutput() ReplicaSetConditionArrayOutput {
	return pulumi.ToOutput(a).(ReplicaSetConditionArrayOutput)
}

func (a ReplicaSetConditionArray) ToReplicaSetConditionArrayOutputWithContext(ctx context.Context) ReplicaSetConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ReplicaSetConditionArrayOutput)
}

type ReplicaSetConditionArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ReplicaSetConditionArrayOutput{}) }

func (ReplicaSetConditionArrayOutput) ElementType() reflect.Type {
	return _ReplicaSetConditionArrayType
}

func (o ReplicaSetConditionArrayOutput) Index(i pulumi.IntInput) ReplicaSetConditionOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) ReplicaSetCondition {
		return vs[0].([]ReplicaSetCondition)[vs[1].(int)]
	}).(ReplicaSetConditionOutput)
}
