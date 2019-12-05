// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// RollingUpdateStatefulSetStrategy is used to communicate parameter for
// RollingUpdateStatefulSetStrategyType.
type RollingUpdateStatefulSetStrategy struct {
	// Partition indicates the ordinal at which the StatefulSet should be partitioned. Default value is
	// 0.
	Partition *int `pulumi:"partition"`

}

var _RollingUpdateStatefulSetStrategyType = reflect.TypeOf((*RollingUpdateStatefulSetStrategy)(nil)).Elem()

// RollingUpdateStatefulSetStrategyInput represents an input type that resolves to a RollingUpdateStatefulSetStrategy.
type RollingUpdateStatefulSetStrategyInput interface {
	ElementType() reflect.Type

	ToRollingUpdateStatefulSetStrategyOutput() RollingUpdateStatefulSetStrategyOutput
	ToRollingUpdateStatefulSetStrategyOutputWithContext(ctx context.Context) RollingUpdateStatefulSetStrategyOutput
}

// RollingUpdateStatefulSetStrategyArgs is a RollingUpdateStatefulSetStrategyInput whose fields are all Input types.
type RollingUpdateStatefulSetStrategyArgs struct {
	// Partition indicates the ordinal at which the StatefulSet should be partitioned. Default value is
	// 0.
	Partition pulumi.IntInput `pulumi:"partition"`

}

func (a RollingUpdateStatefulSetStrategyArgs) ElementType() reflect.Type {
	return _RollingUpdateStatefulSetStrategyType
}

func (a RollingUpdateStatefulSetStrategyArgs) ToRollingUpdateStatefulSetStrategyOutput() RollingUpdateStatefulSetStrategyOutput {
	return pulumi.ToOutput(a).(RollingUpdateStatefulSetStrategyOutput)
}

func (a RollingUpdateStatefulSetStrategyArgs) ToRollingUpdateStatefulSetStrategyOutputWithContext(ctx context.Context) RollingUpdateStatefulSetStrategyOutput {
	return pulumi.ToOutputWithContext(ctx, a).(RollingUpdateStatefulSetStrategyOutput)
}

// RollingUpdateStatefulSetStrategyOutput is an output type that resolves to a Input.
type RollingUpdateStatefulSetStrategyOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(RollingUpdateStatefulSetStrategyOutput{}) }

func (RollingUpdateStatefulSetStrategyOutput) ElementType() reflect.Type {
	return _RollingUpdateStatefulSetStrategyType
}

func (o RollingUpdateStatefulSetStrategyOutput) Partition() pulumi.IntOutput {
	return o.Apply(func(v RollingUpdateStatefulSetStrategy) *int {
		return v.Partition
	}).(pulumi.IntOutput)
}

