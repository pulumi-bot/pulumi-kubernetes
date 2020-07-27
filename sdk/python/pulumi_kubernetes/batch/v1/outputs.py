# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from ... import _utilities, _tables
from ... import core as _core
from ... import meta as _meta

__all__ = [
    'Job',
    'JobCondition',
    'JobSpec',
    'JobStatus',
]

@pulumi.output_type
class Job(dict):
    """
    Job represents the configuration of a single job.

    This resource waits until its status is ready before registering success
    for create/update, and populating output properties from the current state of the resource.
    The following conditions are used to determine whether the resource creation has
    succeeded or failed:

    1. The Job's '.status.startTime' is set, which indicates that the Job has started running.
    2. The Job's '.status.conditions' has a status of type 'Complete', and a 'status' set
       to 'True'.
    3. The Job's '.status.conditions' do not have a status of type 'Failed', with a
    	'status' set to 'True'. If this condition is set, we should fail the Job immediately.

    If the Job has not reached a Ready state after 10 minutes, it will
    time out and mark the resource update as Failed. You can override the default timeout value
    by setting the 'customTimeouts' option on the resource.
    """
    api_version: Optional[str] = pulumi.output_property("apiVersion")
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    """
    kind: Optional[str] = pulumi.output_property("kind")
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    """
    metadata: Optional['_meta.v1.outputs.ObjectMeta'] = pulumi.output_property("metadata")
    """
    Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
    """
    spec: Optional['outputs.JobSpec'] = pulumi.output_property("spec")
    """
    Specification of the desired behavior of a job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
    """
    status: Optional['outputs.JobStatus'] = pulumi.output_property("status")
    """
    Current status of a job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class JobCondition(dict):
    """
    JobCondition describes current state of a job.
    """
    last_probe_time: Optional[str] = pulumi.output_property("lastProbeTime")
    """
    Last time the condition was checked.
    """
    last_transition_time: Optional[str] = pulumi.output_property("lastTransitionTime")
    """
    Last time the condition transit from one status to another.
    """
    message: Optional[str] = pulumi.output_property("message")
    """
    Human readable message indicating details about last transition.
    """
    reason: Optional[str] = pulumi.output_property("reason")
    """
    (brief) reason for the condition's last transition.
    """
    status: str = pulumi.output_property("status")
    """
    Status of the condition, one of True, False, Unknown.
    """
    type: str = pulumi.output_property("type")
    """
    Type of job condition, Complete or Failed.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class JobSpec(dict):
    """
    JobSpec describes how the job execution will look like.
    """
    active_deadline_seconds: Optional[float] = pulumi.output_property("activeDeadlineSeconds")
    """
    Specifies the duration in seconds relative to the startTime that the job may be active before the system tries to terminate it; value must be positive integer
    """
    backoff_limit: Optional[float] = pulumi.output_property("backoffLimit")
    """
    Specifies the number of retries before marking this job failed. Defaults to 6
    """
    completions: Optional[float] = pulumi.output_property("completions")
    """
    Specifies the desired number of successfully finished pods the job should be run with.  Setting to nil means that the success of any pod signals the success of all pods, and allows parallelism to have any positive value.  Setting to 1 means that parallelism is limited to 1 and the success of that pod signals the success of the job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
    """
    manual_selector: Optional[bool] = pulumi.output_property("manualSelector")
    """
    manualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template.  When true, the user is responsible for picking unique labels and specifying the selector.  Failure to pick a unique label may cause this and other jobs to not function correctly.  However, You may see `manualSelector=true` in jobs that were created with the old `extensions/v1beta1` API. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#specifying-your-own-pod-selector
    """
    parallelism: Optional[float] = pulumi.output_property("parallelism")
    """
    Specifies the maximum desired number of pods the job should run at any given time. The actual number of pods running in steady state will be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max parallelism. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
    """
    selector: Optional['_meta.v1.outputs.LabelSelector'] = pulumi.output_property("selector")
    """
    A label query over pods that should match the pod count. Normally, the system sets this field for you. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
    """
    template: '_core.v1.outputs.PodTemplateSpec' = pulumi.output_property("template")
    """
    Describes the pod that will be created when executing a job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
    """
    ttl_seconds_after_finished: Optional[float] = pulumi.output_property("ttlSecondsAfterFinished")
    """
    ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either Complete or Failed). If this field is set, ttlSecondsAfterFinished after the Job finishes, it is eligible to be automatically deleted. When the Job is being deleted, its lifecycle guarantees (e.g. finalizers) will be honored. If this field is unset, the Job won't be automatically deleted. If this field is set to zero, the Job becomes eligible to be deleted immediately after it finishes. This field is alpha-level and is only honored by servers that enable the TTLAfterFinished feature.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class JobStatus(dict):
    """
    JobStatus represents the current state of a Job.
    """
    active: Optional[float] = pulumi.output_property("active")
    """
    The number of actively running pods.
    """
    completion_time: Optional[str] = pulumi.output_property("completionTime")
    """
    Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
    """
    conditions: Optional[List['outputs.JobCondition']] = pulumi.output_property("conditions")
    """
    The latest available observations of an object's current state. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
    """
    failed: Optional[float] = pulumi.output_property("failed")
    """
    The number of pods which reached phase Failed.
    """
    start_time: Optional[str] = pulumi.output_property("startTime")
    """
    Represents time when the job was acknowledged by the job controller. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
    """
    succeeded: Optional[float] = pulumi.output_property("succeeded")
    """
    The number of pods which reached phase Succeeded.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


