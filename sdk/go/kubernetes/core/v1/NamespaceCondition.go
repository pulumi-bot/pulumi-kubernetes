// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// NamespaceCondition contains details about state of namespace.
type NamespaceCondition struct {
	
	LastTransitionTime *string `pulumi:"lastTransitionTime"`

	
	Message *string `pulumi:"message"`

	
	Reason *string `pulumi:"reason"`

	// Status of the condition, one of True, False, Unknown.
	Status string `pulumi:"status"`

	// Type of namespace controller condition.
	Type string `pulumi:"type"`

}

var _NamespaceConditionType = reflect.TypeOf((*NamespaceCondition)(nil)).Elem()

// NamespaceConditionInput represents an input type that resolves to a NamespaceCondition.
type NamespaceConditionInput interface {
	ElementType() reflect.Type

	ToNamespaceConditionOutput() NamespaceConditionOutput
	ToNamespaceConditionOutputWithContext(ctx context.Context) NamespaceConditionOutput
}

// NamespaceConditionArgs is a NamespaceConditionInput whose fields are all Input types.
type NamespaceConditionArgs struct {
	// Status of the condition, one of True, False, Unknown.
	Status pulumi.StringInput `pulumi:"status"`

	// Type of namespace controller condition.
	Type pulumi.StringInput `pulumi:"type"`

	
	LastTransitionTime pulumi.StringInput `pulumi:"lastTransitionTime"`

	
	Message pulumi.StringInput `pulumi:"message"`

	
	Reason pulumi.StringInput `pulumi:"reason"`

}

func (a NamespaceConditionArgs) ElementType() reflect.Type {
	return _NamespaceConditionType
}

func (a NamespaceConditionArgs) ToNamespaceConditionOutput() NamespaceConditionOutput {
	return pulumi.ToOutput(a).(NamespaceConditionOutput)
}

func (a NamespaceConditionArgs) ToNamespaceConditionOutputWithContext(ctx context.Context) NamespaceConditionOutput {
	return pulumi.ToOutputWithContext(ctx, a).(NamespaceConditionOutput)
}

// NamespaceConditionOutput is an output type that resolves to a Input.
type NamespaceConditionOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(NamespaceConditionOutput{}) }

func (NamespaceConditionOutput) ElementType() reflect.Type {
	return _NamespaceConditionType
}

func (o NamespaceConditionOutput) LastTransitionTime() pulumi.StringOutput {
	return o.Apply(func(v NamespaceCondition) *string {
		return v.LastTransitionTime
	}).(pulumi.StringOutput)
}

func (o NamespaceConditionOutput) Message() pulumi.StringOutput {
	return o.Apply(func(v NamespaceCondition) *string {
		return v.Message
	}).(pulumi.StringOutput)
}

func (o NamespaceConditionOutput) Reason() pulumi.StringOutput {
	return o.Apply(func(v NamespaceCondition) *string {
		return v.Reason
	}).(pulumi.StringOutput)
}

func (o NamespaceConditionOutput) Status() pulumi.StringOutput {
	return o.Apply(func(v NamespaceCondition) string {
		return v.Status
	}).(pulumi.StringOutput)
}

func (o NamespaceConditionOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v NamespaceCondition) string {
		return v.Type
	}).(pulumi.StringOutput)
}

var _NamespaceConditionArrayType = reflect.TypeOf((*[]NamespaceCondition)(nil)).Elem()

type NamespaceConditionArrayInput interface {
	ElementType() reflect.Type

	ToNamespaceConditionArrayOutput() NamespaceConditionArrayOutput
	ToNamespaceConditionArrayOutputWithContext(ctx context.Context) NamespaceConditionArrayOutput
}

type NamespaceConditionArray []NamespaceConditionInput

func (a NamespaceConditionArray) ElementType() reflect.Type {
	return _NamespaceConditionArrayType
}

func (a NamespaceConditionArray) ToNamespaceConditionArrayOutput() NamespaceConditionArrayOutput {
	return pulumi.ToOutput(a).(NamespaceConditionArrayOutput)
}

func (a NamespaceConditionArray) ToNamespaceConditionArrayOutputWithContext(ctx context.Context) NamespaceConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(NamespaceConditionArrayOutput)
}

type NamespaceConditionArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(NamespaceConditionArrayOutput{}) }

func (NamespaceConditionArrayOutput) ElementType() reflect.Type {
	return _NamespaceConditionArrayType
}

func (o NamespaceConditionArrayOutput) Index(i pulumi.IntInput) NamespaceConditionOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) NamespaceCondition {
		return vs[0].([]NamespaceCondition)[vs[1].(int)]
	}).(NamespaceConditionOutput)
}