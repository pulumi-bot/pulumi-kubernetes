# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from ... import _utilities, _tables
from ._inputs import *
from . import outputs

@pulumi.output_type
class CronJob(dict):
    """
    CronJob represents the configuration of a single cron job.
    """
    api_version: Optional[str] = pulumi.output_property("apiVersion")
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    """
    kind: Optional[str] = pulumi.output_property("kind")
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    """
    metadata: Optional['outputs.ObjectMeta'] = pulumi.output_property("metadata")
    """
    Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
    """
    spec: Optional['outputs.CronJobSpec'] = pulumi.output_property("spec")
    """
    Specification of the desired behavior of a cron job, including the schedule. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
    """
    status: Optional['outputs.CronJobStatus'] = pulumi.output_property("status")
    """
    Current status of a cron job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CronJobSpec(dict):
    """
    CronJobSpec describes how the job execution will look like and when it will actually run.
    """
    concurrency_policy: Optional[str] = pulumi.output_property("concurrencyPolicy")
    """
    Specifies how to treat concurrent executions of a Job. Valid values are: - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one
    """
    failed_jobs_history_limit: Optional[float] = pulumi.output_property("failedJobsHistoryLimit")
    """
    The number of failed finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
    """
    job_template: 'outputs.JobTemplateSpec' = pulumi.output_property("jobTemplate")
    """
    Specifies the job that will be created when executing a CronJob.
    """
    schedule: str = pulumi.output_property("schedule")
    """
    The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
    """
    starting_deadline_seconds: Optional[float] = pulumi.output_property("startingDeadlineSeconds")
    """
    Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.
    """
    successful_jobs_history_limit: Optional[float] = pulumi.output_property("successfulJobsHistoryLimit")
    """
    The number of successful finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
    """
    suspend: Optional[bool] = pulumi.output_property("suspend")
    """
    This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.  Defaults to false.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CronJobStatus(dict):
    """
    CronJobStatus represents the current state of a cron job.
    """
    active: Optional[List['outputs.ObjectReference']] = pulumi.output_property("active")
    """
    A list of pointers to currently running jobs.
    """
    last_schedule_time: Optional[str] = pulumi.output_property("lastScheduleTime")
    """
    Information when was the last time the job was successfully scheduled.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class JobTemplateSpec(dict):
    """
    JobTemplateSpec describes the data a Job should have when created from a template
    """
    metadata: Optional['outputs.ObjectMeta'] = pulumi.output_property("metadata")
    """
    Standard object's metadata of the jobs created from this template. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
    """
    spec: Optional['outputs.JobSpec'] = pulumi.output_property("spec")
    """
    Specification of the desired behavior of the job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


