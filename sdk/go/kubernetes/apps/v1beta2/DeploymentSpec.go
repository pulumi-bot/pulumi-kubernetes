// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	coreV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/core/v1"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// DeploymentSpec is the specification of the desired behavior of the Deployment.
type DeploymentSpec struct {
	// Minimum number of seconds for which a newly created pod should be ready without any of its
	// container crashing, for it to be considered available. Defaults to 0 (pod will be considered
	// available as soon as it is ready)
	MinReadySeconds *int `pulumi:"minReadySeconds"`

	// Indicates that the deployment is paused.
	Paused *bool `pulumi:"paused"`

	// The maximum time in seconds for a deployment to make progress before it is considered to be
	// failed. The deployment controller will continue to process failed deployments and a condition
	// with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that
	// progress will not be estimated during the time a deployment is paused. Defaults to 600s.
	ProgressDeadlineSeconds *int `pulumi:"progressDeadlineSeconds"`

	// Number of desired pods. This is a pointer to distinguish between explicit zero and not
	// specified. Defaults to 1.
	Replicas *int `pulumi:"replicas"`

	// The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish
	// between explicit zero and not specified. Defaults to 10.
	RevisionHistoryLimit *int `pulumi:"revisionHistoryLimit"`

	// Label selector for pods. Existing ReplicaSets whose pods are selected by this will be the ones
	// affected by this deployment. It must match the pod template's labels.
	Selector metaV1.LabelSelector `pulumi:"selector"`

	// The deployment strategy to use to replace existing pods with new ones.
	Strategy *DeploymentStrategy `pulumi:"strategy"`

	// Template describes the pods that will be created.
	Template coreV1.PodTemplateSpec `pulumi:"template"`

}

var _DeploymentSpecType = reflect.TypeOf((*DeploymentSpec)(nil)).Elem()

// DeploymentSpecInput represents an input type that resolves to a DeploymentSpec.
type DeploymentSpecInput interface {
	ElementType() reflect.Type

	ToDeploymentSpecOutput() DeploymentSpecOutput
	ToDeploymentSpecOutputWithContext(ctx context.Context) DeploymentSpecOutput
}

// DeploymentSpecArgs is a DeploymentSpecInput whose fields are all Input types.
type DeploymentSpecArgs struct {
	// Label selector for pods. Existing ReplicaSets whose pods are selected by this will be the ones
	// affected by this deployment. It must match the pod template's labels.
	Selector metaV1.LabelSelectorInput `pulumi:"selector"`

	// Template describes the pods that will be created.
	Template coreV1.PodTemplateSpecInput `pulumi:"template"`

	// Minimum number of seconds for which a newly created pod should be ready without any of its
	// container crashing, for it to be considered available. Defaults to 0 (pod will be considered
	// available as soon as it is ready)
	MinReadySeconds pulumi.IntInput `pulumi:"minReadySeconds"`

	// Indicates that the deployment is paused.
	Paused pulumi.BoolInput `pulumi:"paused"`

	// The maximum time in seconds for a deployment to make progress before it is considered to be
	// failed. The deployment controller will continue to process failed deployments and a condition
	// with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that
	// progress will not be estimated during the time a deployment is paused. Defaults to 600s.
	ProgressDeadlineSeconds pulumi.IntInput `pulumi:"progressDeadlineSeconds"`

	// Number of desired pods. This is a pointer to distinguish between explicit zero and not
	// specified. Defaults to 1.
	Replicas pulumi.IntInput `pulumi:"replicas"`

	// The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish
	// between explicit zero and not specified. Defaults to 10.
	RevisionHistoryLimit pulumi.IntInput `pulumi:"revisionHistoryLimit"`

	// The deployment strategy to use to replace existing pods with new ones.
	Strategy DeploymentStrategyInput `pulumi:"strategy"`

}

func (a DeploymentSpecArgs) ElementType() reflect.Type {
	return _DeploymentSpecType
}

func (a DeploymentSpecArgs) ToDeploymentSpecOutput() DeploymentSpecOutput {
	return pulumi.ToOutput(a).(DeploymentSpecOutput)
}

func (a DeploymentSpecArgs) ToDeploymentSpecOutputWithContext(ctx context.Context) DeploymentSpecOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DeploymentSpecOutput)
}

// DeploymentSpecOutput is an output type that resolves to a Input.
type DeploymentSpecOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(DeploymentSpecOutput{}) }

func (DeploymentSpecOutput) ElementType() reflect.Type {
	return _DeploymentSpecType
}

func (o DeploymentSpecOutput) MinReadySeconds() pulumi.IntOutput {
	return o.Apply(func(v DeploymentSpec) *int {
		return v.MinReadySeconds
	}).(pulumi.IntOutput)
}

func (o DeploymentSpecOutput) Paused() pulumi.BoolOutput {
	return o.Apply(func(v DeploymentSpec) *bool {
		return v.Paused
	}).(pulumi.BoolOutput)
}

func (o DeploymentSpecOutput) ProgressDeadlineSeconds() pulumi.IntOutput {
	return o.Apply(func(v DeploymentSpec) *int {
		return v.ProgressDeadlineSeconds
	}).(pulumi.IntOutput)
}

func (o DeploymentSpecOutput) Replicas() pulumi.IntOutput {
	return o.Apply(func(v DeploymentSpec) *int {
		return v.Replicas
	}).(pulumi.IntOutput)
}

func (o DeploymentSpecOutput) RevisionHistoryLimit() pulumi.IntOutput {
	return o.Apply(func(v DeploymentSpec) *int {
		return v.RevisionHistoryLimit
	}).(pulumi.IntOutput)
}

func (o DeploymentSpecOutput) Selector() metaV1.LabelSelectorOutput {
	return o.Apply(func(v DeploymentSpec) metaV1.LabelSelector {
		return v.Selector
	}).(metaV1.LabelSelectorOutput)
}

func (o DeploymentSpecOutput) Strategy() DeploymentStrategyOutput {
	return o.Apply(func(v DeploymentSpec) *DeploymentStrategy {
		return v.Strategy
	}).(DeploymentStrategyOutput)
}

func (o DeploymentSpecOutput) Template() coreV1.PodTemplateSpecOutput {
	return o.Apply(func(v DeploymentSpec) coreV1.PodTemplateSpec {
		return v.Template
	}).(coreV1.PodTemplateSpecOutput)
}

