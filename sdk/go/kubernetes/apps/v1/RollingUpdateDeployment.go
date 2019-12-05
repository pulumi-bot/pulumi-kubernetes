// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Spec to control the desired behavior of rolling update.
type RollingUpdateDeployment struct {
	// The maximum number of pods that can be scheduled above the desired number of pods. Value can be
	// an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if
	// MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to
	// 25%. Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when the
	// rolling update starts, such that the total number of old and new pods do not exceed 130% of
	// desired pods. Once old pods have been killed, new ReplicaSet can be scaled up further, ensuring
	// that total number of pods running at any time during the update is at most 130% of desired pods.
	MaxSurge interface{} `pulumi:"maxSurge"`

	// The maximum number of pods that can be unavailable during the update. Value can be an absolute
	// number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from
	// percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when
	// this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods immediately
	// when the rolling update starts. Once new pods are ready, old ReplicaSet can be scaled down
	// further, followed by scaling up the new ReplicaSet, ensuring that the total number of pods
	// available at all times during the update is at least 70% of desired pods.
	MaxUnavailable interface{} `pulumi:"maxUnavailable"`

}

var _RollingUpdateDeploymentType = reflect.TypeOf((*RollingUpdateDeployment)(nil)).Elem()

// RollingUpdateDeploymentInput represents an input type that resolves to a RollingUpdateDeployment.
type RollingUpdateDeploymentInput interface {
	ElementType() reflect.Type

	ToRollingUpdateDeploymentOutput() RollingUpdateDeploymentOutput
	ToRollingUpdateDeploymentOutputWithContext(ctx context.Context) RollingUpdateDeploymentOutput
}

// RollingUpdateDeploymentArgs is a RollingUpdateDeploymentInput whose fields are all Input types.
type RollingUpdateDeploymentArgs struct {
	// The maximum number of pods that can be scheduled above the desired number of pods. Value can be
	// an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if
	// MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to
	// 25%. Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when the
	// rolling update starts, such that the total number of old and new pods do not exceed 130% of
	// desired pods. Once old pods have been killed, new ReplicaSet can be scaled up further, ensuring
	// that total number of pods running at any time during the update is at most 130% of desired pods.
	MaxSurge pulumi.Input `pulumi:"maxSurge"`

	// The maximum number of pods that can be unavailable during the update. Value can be an absolute
	// number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from
	// percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when
	// this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods immediately
	// when the rolling update starts. Once new pods are ready, old ReplicaSet can be scaled down
	// further, followed by scaling up the new ReplicaSet, ensuring that the total number of pods
	// available at all times during the update is at least 70% of desired pods.
	MaxUnavailable pulumi.Input `pulumi:"maxUnavailable"`

}

func (a RollingUpdateDeploymentArgs) ElementType() reflect.Type {
	return _RollingUpdateDeploymentType
}

func (a RollingUpdateDeploymentArgs) ToRollingUpdateDeploymentOutput() RollingUpdateDeploymentOutput {
	return pulumi.ToOutput(a).(RollingUpdateDeploymentOutput)
}

func (a RollingUpdateDeploymentArgs) ToRollingUpdateDeploymentOutputWithContext(ctx context.Context) RollingUpdateDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, a).(RollingUpdateDeploymentOutput)
}

// RollingUpdateDeploymentOutput is an output type that resolves to a Input.
type RollingUpdateDeploymentOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(RollingUpdateDeploymentOutput{}) }

func (RollingUpdateDeploymentOutput) ElementType() reflect.Type {
	return _RollingUpdateDeploymentType
}

func (o RollingUpdateDeploymentOutput) MaxSurge() pulumi.AnyOutput {
	return o.Apply(func(v RollingUpdateDeployment) interface{} {
		return v.MaxSurge
	}).(pulumi.AnyOutput)
}

func (o RollingUpdateDeploymentOutput) MaxUnavailable() pulumi.AnyOutput {
	return o.Apply(func(v RollingUpdateDeployment) interface{} {
		return v.MaxUnavailable
	}).(pulumi.AnyOutput)
}

