// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	coreV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/core/v1"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// JobSpec describes how the job execution will look like.
type JobSpec struct {
	// Specifies the duration in seconds relative to the startTime that the job may be active before
	// the system tries to terminate it; value must be positive integer
	ActiveDeadlineSeconds *int `pulumi:"activeDeadlineSeconds"`

	// Specifies the number of retries before marking this job failed. Defaults to 6
	BackoffLimit *int `pulumi:"backoffLimit"`

	// Specifies the desired number of successfully finished pods the job should be run with.  Setting
	// to nil means that the success of any pod signals the success of all pods, and allows parallelism
	// to have any positive value.  Setting to 1 means that parallelism is limited to 1 and the success
	// of that pod signals the success of the job. More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Completions *int `pulumi:"completions"`

	// manualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset
	// unless you are certain what you are doing. When false or unset, the system pick labels unique to
	// this job and appends those labels to the pod template.  When true, the user is responsible for
	// picking unique labels and specifying the selector.  Failure to pick a unique label may cause
	// this and other jobs to not function correctly.  However, You may see `manualSelector=true` in
	// jobs that were created with the old `extensions/v1beta1` API. More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#specifying-your-own-pod-selector
	ManualSelector *bool `pulumi:"manualSelector"`

	// Specifies the maximum desired number of pods the job should run at any given time. The actual
	// number of pods running in steady state will be less than this number when ((.spec.completions -
	// .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max
	// parallelism. More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Parallelism *int `pulumi:"parallelism"`

	// A label query over pods that should match the pod count. Normally, the system sets this field
	// for you. More info:
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Selector *metaV1.LabelSelector `pulumi:"selector"`

	// Describes the pod that will be created when executing a job. More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Template coreV1.PodTemplateSpec `pulumi:"template"`

	// ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either
	// Complete or Failed). If this field is set, ttlSecondsAfterFinished after the Job finishes, it is
	// eligible to be automatically deleted. When the Job is being deleted, its lifecycle guarantees
	// (e.g. finalizers) will be honored. If this field is unset, the Job won't be automatically
	// deleted. If this field is set to zero, the Job becomes eligible to be deleted immediately after
	// it finishes. This field is alpha-level and is only honored by servers that enable the
	// TTLAfterFinished feature.
	TtlSecondsAfterFinished *int `pulumi:"ttlSecondsAfterFinished"`

}

var _JobSpecType = reflect.TypeOf((*JobSpec)(nil)).Elem()

// JobSpecInput represents an input type that resolves to a JobSpec.
type JobSpecInput interface {
	ElementType() reflect.Type

	ToJobSpecOutput() JobSpecOutput
	ToJobSpecOutputWithContext(ctx context.Context) JobSpecOutput
}

// JobSpecArgs is a JobSpecInput whose fields are all Input types.
type JobSpecArgs struct {
	// Describes the pod that will be created when executing a job. More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Template coreV1.PodTemplateSpecInput `pulumi:"template"`

	// Specifies the duration in seconds relative to the startTime that the job may be active before
	// the system tries to terminate it; value must be positive integer
	ActiveDeadlineSeconds pulumi.IntInput `pulumi:"activeDeadlineSeconds"`

	// Specifies the number of retries before marking this job failed. Defaults to 6
	BackoffLimit pulumi.IntInput `pulumi:"backoffLimit"`

	// Specifies the desired number of successfully finished pods the job should be run with.  Setting
	// to nil means that the success of any pod signals the success of all pods, and allows parallelism
	// to have any positive value.  Setting to 1 means that parallelism is limited to 1 and the success
	// of that pod signals the success of the job. More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Completions pulumi.IntInput `pulumi:"completions"`

	// manualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset
	// unless you are certain what you are doing. When false or unset, the system pick labels unique to
	// this job and appends those labels to the pod template.  When true, the user is responsible for
	// picking unique labels and specifying the selector.  Failure to pick a unique label may cause
	// this and other jobs to not function correctly.  However, You may see `manualSelector=true` in
	// jobs that were created with the old `extensions/v1beta1` API. More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#specifying-your-own-pod-selector
	ManualSelector pulumi.BoolInput `pulumi:"manualSelector"`

	// Specifies the maximum desired number of pods the job should run at any given time. The actual
	// number of pods running in steady state will be less than this number when ((.spec.completions -
	// .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max
	// parallelism. More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Parallelism pulumi.IntInput `pulumi:"parallelism"`

	// A label query over pods that should match the pod count. Normally, the system sets this field
	// for you. More info:
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Selector metaV1.LabelSelectorInput `pulumi:"selector"`

	// ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either
	// Complete or Failed). If this field is set, ttlSecondsAfterFinished after the Job finishes, it is
	// eligible to be automatically deleted. When the Job is being deleted, its lifecycle guarantees
	// (e.g. finalizers) will be honored. If this field is unset, the Job won't be automatically
	// deleted. If this field is set to zero, the Job becomes eligible to be deleted immediately after
	// it finishes. This field is alpha-level and is only honored by servers that enable the
	// TTLAfterFinished feature.
	TtlSecondsAfterFinished pulumi.IntInput `pulumi:"ttlSecondsAfterFinished"`

}

func (a JobSpecArgs) ElementType() reflect.Type {
	return _JobSpecType
}

func (a JobSpecArgs) ToJobSpecOutput() JobSpecOutput {
	return pulumi.ToOutput(a).(JobSpecOutput)
}

func (a JobSpecArgs) ToJobSpecOutputWithContext(ctx context.Context) JobSpecOutput {
	return pulumi.ToOutputWithContext(ctx, a).(JobSpecOutput)
}

// JobSpecOutput is an output type that resolves to a Input.
type JobSpecOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(JobSpecOutput{}) }

func (JobSpecOutput) ElementType() reflect.Type {
	return _JobSpecType
}

func (o JobSpecOutput) ActiveDeadlineSeconds() pulumi.IntOutput {
	return o.Apply(func(v JobSpec) *int {
		return v.ActiveDeadlineSeconds
	}).(pulumi.IntOutput)
}

func (o JobSpecOutput) BackoffLimit() pulumi.IntOutput {
	return o.Apply(func(v JobSpec) *int {
		return v.BackoffLimit
	}).(pulumi.IntOutput)
}

func (o JobSpecOutput) Completions() pulumi.IntOutput {
	return o.Apply(func(v JobSpec) *int {
		return v.Completions
	}).(pulumi.IntOutput)
}

func (o JobSpecOutput) ManualSelector() pulumi.BoolOutput {
	return o.Apply(func(v JobSpec) *bool {
		return v.ManualSelector
	}).(pulumi.BoolOutput)
}

func (o JobSpecOutput) Parallelism() pulumi.IntOutput {
	return o.Apply(func(v JobSpec) *int {
		return v.Parallelism
	}).(pulumi.IntOutput)
}

func (o JobSpecOutput) Selector() metaV1.LabelSelectorOutput {
	return o.Apply(func(v JobSpec) *metaV1.LabelSelector {
		return v.Selector
	}).(metaV1.LabelSelectorOutput)
}

func (o JobSpecOutput) Template() coreV1.PodTemplateSpecOutput {
	return o.Apply(func(v JobSpec) coreV1.PodTemplateSpec {
		return v.Template
	}).(coreV1.PodTemplateSpecOutput)
}

func (o JobSpecOutput) TtlSecondsAfterFinished() pulumi.IntOutput {
	return o.Apply(func(v JobSpec) *int {
		return v.TtlSecondsAfterFinished
	}).(pulumi.IntOutput)
}

