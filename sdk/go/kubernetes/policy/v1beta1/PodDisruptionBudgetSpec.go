// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// PodDisruptionBudgetSpec is a description of a PodDisruptionBudget.
type PodDisruptionBudgetSpec struct {
	// An eviction is allowed if at most "maxUnavailable" pods selected by "selector" are unavailable
	// after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all
	// voluntary evictions by specifying 0. This is a mutually exclusive setting with "minAvailable".
	MaxUnavailable interface{} `pulumi:"maxUnavailable"`

	// An eviction is allowed if at least "minAvailable" pods selected by "selector" will still be
	// available after the eviction, i.e. even in the absence of the evicted pod.  So for example you
	// can prevent all voluntary evictions by specifying "100%".
	MinAvailable interface{} `pulumi:"minAvailable"`

	// Label query over pods whose evictions are managed by the disruption budget.
	Selector *metaV1.LabelSelector `pulumi:"selector"`

}

var _PodDisruptionBudgetSpecType = reflect.TypeOf((*PodDisruptionBudgetSpec)(nil)).Elem()

// PodDisruptionBudgetSpecInput represents an input type that resolves to a PodDisruptionBudgetSpec.
type PodDisruptionBudgetSpecInput interface {
	ElementType() reflect.Type

	ToPodDisruptionBudgetSpecOutput() PodDisruptionBudgetSpecOutput
	ToPodDisruptionBudgetSpecOutputWithContext(ctx context.Context) PodDisruptionBudgetSpecOutput
}

// PodDisruptionBudgetSpecArgs is a PodDisruptionBudgetSpecInput whose fields are all Input types.
type PodDisruptionBudgetSpecArgs struct {
	// An eviction is allowed if at most "maxUnavailable" pods selected by "selector" are unavailable
	// after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all
	// voluntary evictions by specifying 0. This is a mutually exclusive setting with "minAvailable".
	MaxUnavailable pulumi.Input `pulumi:"maxUnavailable"`

	// An eviction is allowed if at least "minAvailable" pods selected by "selector" will still be
	// available after the eviction, i.e. even in the absence of the evicted pod.  So for example you
	// can prevent all voluntary evictions by specifying "100%".
	MinAvailable pulumi.Input `pulumi:"minAvailable"`

	// Label query over pods whose evictions are managed by the disruption budget.
	Selector metaV1.LabelSelectorInput `pulumi:"selector"`

}

func (a PodDisruptionBudgetSpecArgs) ElementType() reflect.Type {
	return _PodDisruptionBudgetSpecType
}

func (a PodDisruptionBudgetSpecArgs) ToPodDisruptionBudgetSpecOutput() PodDisruptionBudgetSpecOutput {
	return pulumi.ToOutput(a).(PodDisruptionBudgetSpecOutput)
}

func (a PodDisruptionBudgetSpecArgs) ToPodDisruptionBudgetSpecOutputWithContext(ctx context.Context) PodDisruptionBudgetSpecOutput {
	return pulumi.ToOutputWithContext(ctx, a).(PodDisruptionBudgetSpecOutput)
}

// PodDisruptionBudgetSpecOutput is an output type that resolves to a Input.
type PodDisruptionBudgetSpecOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(PodDisruptionBudgetSpecOutput{}) }

func (PodDisruptionBudgetSpecOutput) ElementType() reflect.Type {
	return _PodDisruptionBudgetSpecType
}

func (o PodDisruptionBudgetSpecOutput) MaxUnavailable() pulumi.AnyOutput {
	return o.Apply(func(v PodDisruptionBudgetSpec) interface{} {
		return v.MaxUnavailable
	}).(pulumi.AnyOutput)
}

func (o PodDisruptionBudgetSpecOutput) MinAvailable() pulumi.AnyOutput {
	return o.Apply(func(v PodDisruptionBudgetSpec) interface{} {
		return v.MinAvailable
	}).(pulumi.AnyOutput)
}

func (o PodDisruptionBudgetSpecOutput) Selector() metaV1.LabelSelectorOutput {
	return o.Apply(func(v PodDisruptionBudgetSpec) *metaV1.LabelSelector {
		return v.Selector
	}).(metaV1.LabelSelectorOutput)
}

