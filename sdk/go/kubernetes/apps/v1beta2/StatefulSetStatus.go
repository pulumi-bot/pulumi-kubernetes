// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// StatefulSetStatus represents the current state of a StatefulSet.
type StatefulSetStatus struct {
	// collisionCount is the count of hash collisions for the StatefulSet. The StatefulSet controller
	// uses this field as a collision avoidance mechanism when it needs to create the name for the
	// newest ControllerRevision.
	CollisionCount *int `pulumi:"collisionCount"`

	// Represents the latest available observations of a statefulset's current state.
	Conditions []StatefulSetCondition `pulumi:"conditions"`

	// currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet
	// version indicated by currentRevision.
	CurrentReplicas *int `pulumi:"currentReplicas"`

	// currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in
	// the sequence [0,currentReplicas).
	CurrentRevision *string `pulumi:"currentRevision"`

	// observedGeneration is the most recent generation observed for this StatefulSet. It corresponds
	// to the StatefulSet's generation, which is updated on mutation by the API Server.
	ObservedGeneration *int `pulumi:"observedGeneration"`

	// readyReplicas is the number of Pods created by the StatefulSet controller that have a Ready
	// Condition.
	ReadyReplicas *int `pulumi:"readyReplicas"`

	// replicas is the number of Pods created by the StatefulSet controller.
	Replicas int `pulumi:"replicas"`

	// updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in
	// the sequence [replicas-updatedReplicas,replicas)
	UpdateRevision *string `pulumi:"updateRevision"`

	// updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet
	// version indicated by updateRevision.
	UpdatedReplicas *int `pulumi:"updatedReplicas"`

}

var _StatefulSetStatusType = reflect.TypeOf((*StatefulSetStatus)(nil)).Elem()

// StatefulSetStatusInput represents an input type that resolves to a StatefulSetStatus.
type StatefulSetStatusInput interface {
	ElementType() reflect.Type

	ToStatefulSetStatusOutput() StatefulSetStatusOutput
	ToStatefulSetStatusOutputWithContext(ctx context.Context) StatefulSetStatusOutput
}

// StatefulSetStatusArgs is a StatefulSetStatusInput whose fields are all Input types.
type StatefulSetStatusArgs struct {
	// replicas is the number of Pods created by the StatefulSet controller.
	Replicas pulumi.IntInput `pulumi:"replicas"`

	// collisionCount is the count of hash collisions for the StatefulSet. The StatefulSet controller
	// uses this field as a collision avoidance mechanism when it needs to create the name for the
	// newest ControllerRevision.
	CollisionCount pulumi.IntInput `pulumi:"collisionCount"`

	// Represents the latest available observations of a statefulset's current state.
	Conditions StatefulSetConditionArrayInput `pulumi:"conditions"`

	// currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet
	// version indicated by currentRevision.
	CurrentReplicas pulumi.IntInput `pulumi:"currentReplicas"`

	// currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in
	// the sequence [0,currentReplicas).
	CurrentRevision pulumi.StringInput `pulumi:"currentRevision"`

	// observedGeneration is the most recent generation observed for this StatefulSet. It corresponds
	// to the StatefulSet's generation, which is updated on mutation by the API Server.
	ObservedGeneration pulumi.IntInput `pulumi:"observedGeneration"`

	// readyReplicas is the number of Pods created by the StatefulSet controller that have a Ready
	// Condition.
	ReadyReplicas pulumi.IntInput `pulumi:"readyReplicas"`

	// updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in
	// the sequence [replicas-updatedReplicas,replicas)
	UpdateRevision pulumi.StringInput `pulumi:"updateRevision"`

	// updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet
	// version indicated by updateRevision.
	UpdatedReplicas pulumi.IntInput `pulumi:"updatedReplicas"`

}

func (a StatefulSetStatusArgs) ElementType() reflect.Type {
	return _StatefulSetStatusType
}

func (a StatefulSetStatusArgs) ToStatefulSetStatusOutput() StatefulSetStatusOutput {
	return pulumi.ToOutput(a).(StatefulSetStatusOutput)
}

func (a StatefulSetStatusArgs) ToStatefulSetStatusOutputWithContext(ctx context.Context) StatefulSetStatusOutput {
	return pulumi.ToOutputWithContext(ctx, a).(StatefulSetStatusOutput)
}

// StatefulSetStatusOutput is an output type that resolves to a Input.
type StatefulSetStatusOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(StatefulSetStatusOutput{}) }

func (StatefulSetStatusOutput) ElementType() reflect.Type {
	return _StatefulSetStatusType
}

func (o StatefulSetStatusOutput) CollisionCount() pulumi.IntOutput {
	return o.Apply(func(v StatefulSetStatus) *int {
		return v.CollisionCount
	}).(pulumi.IntOutput)
}

func (o StatefulSetStatusOutput) Conditions() StatefulSetConditionArrayOutput {
	return o.Apply(func(v StatefulSetStatus) []StatefulSetCondition {
		return v.Conditions
	}).(StatefulSetConditionArrayOutput)
}

func (o StatefulSetStatusOutput) CurrentReplicas() pulumi.IntOutput {
	return o.Apply(func(v StatefulSetStatus) *int {
		return v.CurrentReplicas
	}).(pulumi.IntOutput)
}

func (o StatefulSetStatusOutput) CurrentRevision() pulumi.StringOutput {
	return o.Apply(func(v StatefulSetStatus) *string {
		return v.CurrentRevision
	}).(pulumi.StringOutput)
}

func (o StatefulSetStatusOutput) ObservedGeneration() pulumi.IntOutput {
	return o.Apply(func(v StatefulSetStatus) *int {
		return v.ObservedGeneration
	}).(pulumi.IntOutput)
}

func (o StatefulSetStatusOutput) ReadyReplicas() pulumi.IntOutput {
	return o.Apply(func(v StatefulSetStatus) *int {
		return v.ReadyReplicas
	}).(pulumi.IntOutput)
}

func (o StatefulSetStatusOutput) Replicas() pulumi.IntOutput {
	return o.Apply(func(v StatefulSetStatus) int {
		return v.Replicas
	}).(pulumi.IntOutput)
}

func (o StatefulSetStatusOutput) UpdateRevision() pulumi.StringOutput {
	return o.Apply(func(v StatefulSetStatus) *string {
		return v.UpdateRevision
	}).(pulumi.StringOutput)
}

func (o StatefulSetStatusOutput) UpdatedReplicas() pulumi.IntOutput {
	return o.Apply(func(v StatefulSetStatus) *int {
		return v.UpdatedReplicas
	}).(pulumi.IntOutput)
}

