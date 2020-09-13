# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
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
    def __init__(__self__, *,
                 api_version: Optional[str] = None,
                 kind: Optional[str] = None,
                 metadata: Optional['_meta.v1.outputs.ObjectMeta'] = None,
                 spec: Optional['outputs.JobSpec'] = None,
                 status: Optional['outputs.JobStatus'] = None):
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
        :param str api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param str kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param '_meta.v1.ObjectMetaArgs' metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param 'JobSpecArgs' spec: Specification of the desired behavior of a job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        :param 'JobStatusArgs' status: Current status of a job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        if api_version is not None:
            pulumi.set(__self__, "api_version", 'batch/v1')
        if kind is not None:
            pulumi.set(__self__, "kind", 'Job')
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[str]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def metadata(self) -> Optional['_meta.v1.outputs.ObjectMeta']:
        """
        Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def spec(self) -> Optional['outputs.JobSpec']:
        """
        Specification of the desired behavior of a job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter
    def status(self) -> Optional['outputs.JobStatus']:
        """
        Current status of a job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        return pulumi.get(self, "status")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class JobCondition(dict):
    """
    JobCondition describes current state of a job.
    """
    def __init__(__self__, *,
                 status: str,
                 type: str,
                 last_probe_time: Optional[str] = None,
                 last_transition_time: Optional[str] = None,
                 message: Optional[str] = None,
                 reason: Optional[str] = None):
        """
        JobCondition describes current state of a job.
        :param str status: Status of the condition, one of True, False, Unknown.
        :param str type: Type of job condition, Complete or Failed.
        :param str last_probe_time: Last time the condition was checked.
        :param str last_transition_time: Last time the condition transit from one status to another.
        :param str message: Human readable message indicating details about last transition.
        :param str reason: (brief) reason for the condition's last transition.
        """
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "type", type)
        if last_probe_time is not None:
            pulumi.set(__self__, "last_probe_time", last_probe_time)
        if last_transition_time is not None:
            pulumi.set(__self__, "last_transition_time", last_transition_time)
        if message is not None:
            pulumi.set(__self__, "message", message)
        if reason is not None:
            pulumi.set(__self__, "reason", reason)

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of the condition, one of True, False, Unknown.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of job condition, Complete or Failed.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="lastProbeTime")
    def last_probe_time(self) -> Optional[str]:
        """
        Last time the condition was checked.
        """
        return pulumi.get(self, "last_probe_time")

    @property
    @pulumi.getter(name="lastTransitionTime")
    def last_transition_time(self) -> Optional[str]:
        """
        Last time the condition transit from one status to another.
        """
        return pulumi.get(self, "last_transition_time")

    @property
    @pulumi.getter
    def message(self) -> Optional[str]:
        """
        Human readable message indicating details about last transition.
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter
    def reason(self) -> Optional[str]:
        """
        (brief) reason for the condition's last transition.
        """
        return pulumi.get(self, "reason")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class JobSpec(dict):
    """
    JobSpec describes how the job execution will look like.
    """
    def __init__(__self__, *,
                 template: '_core.v1.outputs.PodTemplateSpec',
                 active_deadline_seconds: Optional[int] = None,
                 backoff_limit: Optional[int] = None,
                 completions: Optional[int] = None,
                 manual_selector: Optional[bool] = None,
                 parallelism: Optional[int] = None,
                 selector: Optional['_meta.v1.outputs.LabelSelector'] = None,
                 ttl_seconds_after_finished: Optional[int] = None):
        """
        JobSpec describes how the job execution will look like.
        :param '_core.v1.PodTemplateSpecArgs' template: Describes the pod that will be created when executing a job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
        :param int active_deadline_seconds: Specifies the duration in seconds relative to the startTime that the job may be active before the system tries to terminate it; value must be positive integer
        :param int backoff_limit: Specifies the number of retries before marking this job failed. Defaults to 6
        :param int completions: Specifies the desired number of successfully finished pods the job should be run with.  Setting to nil means that the success of any pod signals the success of all pods, and allows parallelism to have any positive value.  Setting to 1 means that parallelism is limited to 1 and the success of that pod signals the success of the job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
        :param bool manual_selector: manualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template.  When true, the user is responsible for picking unique labels and specifying the selector.  Failure to pick a unique label may cause this and other jobs to not function correctly.  However, You may see `manualSelector=true` in jobs that were created with the old `extensions/v1beta1` API. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#specifying-your-own-pod-selector
        :param int parallelism: Specifies the maximum desired number of pods the job should run at any given time. The actual number of pods running in steady state will be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max parallelism. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
        :param '_meta.v1.LabelSelectorArgs' selector: A label query over pods that should match the pod count. Normally, the system sets this field for you. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
        :param int ttl_seconds_after_finished: ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either Complete or Failed). If this field is set, ttlSecondsAfterFinished after the Job finishes, it is eligible to be automatically deleted. When the Job is being deleted, its lifecycle guarantees (e.g. finalizers) will be honored. If this field is unset, the Job won't be automatically deleted. If this field is set to zero, the Job becomes eligible to be deleted immediately after it finishes. This field is alpha-level and is only honored by servers that enable the TTLAfterFinished feature.
        """
        pulumi.set(__self__, "template", template)
        if active_deadline_seconds is not None:
            pulumi.set(__self__, "active_deadline_seconds", active_deadline_seconds)
        if backoff_limit is not None:
            pulumi.set(__self__, "backoff_limit", backoff_limit)
        if completions is not None:
            pulumi.set(__self__, "completions", completions)
        if manual_selector is not None:
            pulumi.set(__self__, "manual_selector", manual_selector)
        if parallelism is not None:
            pulumi.set(__self__, "parallelism", parallelism)
        if selector is not None:
            pulumi.set(__self__, "selector", selector)
        if ttl_seconds_after_finished is not None:
            pulumi.set(__self__, "ttl_seconds_after_finished", ttl_seconds_after_finished)

    @property
    @pulumi.getter
    def template(self) -> '_core.v1.outputs.PodTemplateSpec':
        """
        Describes the pod that will be created when executing a job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
        """
        return pulumi.get(self, "template")

    @property
    @pulumi.getter(name="activeDeadlineSeconds")
    def active_deadline_seconds(self) -> Optional[int]:
        """
        Specifies the duration in seconds relative to the startTime that the job may be active before the system tries to terminate it; value must be positive integer
        """
        return pulumi.get(self, "active_deadline_seconds")

    @property
    @pulumi.getter(name="backoffLimit")
    def backoff_limit(self) -> Optional[int]:
        """
        Specifies the number of retries before marking this job failed. Defaults to 6
        """
        return pulumi.get(self, "backoff_limit")

    @property
    @pulumi.getter
    def completions(self) -> Optional[int]:
        """
        Specifies the desired number of successfully finished pods the job should be run with.  Setting to nil means that the success of any pod signals the success of all pods, and allows parallelism to have any positive value.  Setting to 1 means that parallelism is limited to 1 and the success of that pod signals the success of the job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
        """
        return pulumi.get(self, "completions")

    @property
    @pulumi.getter(name="manualSelector")
    def manual_selector(self) -> Optional[bool]:
        """
        manualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template.  When true, the user is responsible for picking unique labels and specifying the selector.  Failure to pick a unique label may cause this and other jobs to not function correctly.  However, You may see `manualSelector=true` in jobs that were created with the old `extensions/v1beta1` API. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#specifying-your-own-pod-selector
        """
        return pulumi.get(self, "manual_selector")

    @property
    @pulumi.getter
    def parallelism(self) -> Optional[int]:
        """
        Specifies the maximum desired number of pods the job should run at any given time. The actual number of pods running in steady state will be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max parallelism. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
        """
        return pulumi.get(self, "parallelism")

    @property
    @pulumi.getter
    def selector(self) -> Optional['_meta.v1.outputs.LabelSelector']:
        """
        A label query over pods that should match the pod count. Normally, the system sets this field for you. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
        """
        return pulumi.get(self, "selector")

    @property
    @pulumi.getter(name="ttlSecondsAfterFinished")
    def ttl_seconds_after_finished(self) -> Optional[int]:
        """
        ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either Complete or Failed). If this field is set, ttlSecondsAfterFinished after the Job finishes, it is eligible to be automatically deleted. When the Job is being deleted, its lifecycle guarantees (e.g. finalizers) will be honored. If this field is unset, the Job won't be automatically deleted. If this field is set to zero, the Job becomes eligible to be deleted immediately after it finishes. This field is alpha-level and is only honored by servers that enable the TTLAfterFinished feature.
        """
        return pulumi.get(self, "ttl_seconds_after_finished")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class JobStatus(dict):
    """
    JobStatus represents the current state of a Job.
    """
    def __init__(__self__, *,
                 active: Optional[int] = None,
                 completion_time: Optional[str] = None,
                 conditions: Optional[Sequence['outputs.JobCondition']] = None,
                 failed: Optional[int] = None,
                 start_time: Optional[str] = None,
                 succeeded: Optional[int] = None):
        """
        JobStatus represents the current state of a Job.
        :param int active: The number of actively running pods.
        :param str completion_time: Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
        :param Sequence['JobConditionArgs'] conditions: The latest available observations of an object's current state. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
        :param int failed: The number of pods which reached phase Failed.
        :param str start_time: Represents time when the job was acknowledged by the job controller. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
        :param int succeeded: The number of pods which reached phase Succeeded.
        """
        if active is not None:
            pulumi.set(__self__, "active", active)
        if completion_time is not None:
            pulumi.set(__self__, "completion_time", completion_time)
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if failed is not None:
            pulumi.set(__self__, "failed", failed)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)
        if succeeded is not None:
            pulumi.set(__self__, "succeeded", succeeded)

    @property
    @pulumi.getter
    def active(self) -> Optional[int]:
        """
        The number of actively running pods.
        """
        return pulumi.get(self, "active")

    @property
    @pulumi.getter(name="completionTime")
    def completion_time(self) -> Optional[str]:
        """
        Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
        """
        return pulumi.get(self, "completion_time")

    @property
    @pulumi.getter
    def conditions(self) -> Optional[Sequence['outputs.JobCondition']]:
        """
        The latest available observations of an object's current state. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
        """
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter
    def failed(self) -> Optional[int]:
        """
        The number of pods which reached phase Failed.
        """
        return pulumi.get(self, "failed")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[str]:
        """
        Represents time when the job was acknowledged by the job controller. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def succeeded(self) -> Optional[int]:
        """
        The number of pods which reached phase Succeeded.
        """
        return pulumi.get(self, "succeeded")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


