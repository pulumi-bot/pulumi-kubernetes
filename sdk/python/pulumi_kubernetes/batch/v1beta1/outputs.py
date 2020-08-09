# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ... import batch as _batch
from ... import core as _core
from ... import meta as _meta

__all__ = [
    'CronJob',
    'CronJobSpec',
    'CronJobStatus',
    'JobTemplateSpec',
]

@pulumi.output_type
class CronJob(dict):
    """
    CronJob represents the configuration of a single cron job.
    """
    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[str]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        ...

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        ...

    @property
    @pulumi.getter
    def metadata(self) -> Optional['_meta.v1.outputs.ObjectMeta']:
        """
        Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        ...

    @property
    @pulumi.getter
    def spec(self) -> Optional['outputs.CronJobSpec']:
        """
        Specification of the desired behavior of a cron job, including the schedule. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        ...

    @property
    @pulumi.getter
    def status(self) -> Optional['outputs.CronJobStatus']:
        """
        Current status of a cron job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CronJobSpec(dict):
    """
    CronJobSpec describes how the job execution will look like and when it will actually run.
    """
    @property
    @pulumi.getter(name="concurrencyPolicy")
    def concurrency_policy(self) -> Optional[str]:
        """
        Specifies how to treat concurrent executions of a Job. Valid values are: - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one
        """
        ...

    @property
    @pulumi.getter(name="failedJobsHistoryLimit")
    def failed_jobs_history_limit(self) -> Optional[float]:
        """
        The number of failed finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
        """
        ...

    @property
    @pulumi.getter(name="jobTemplate")
    def job_template(self) -> 'outputs.JobTemplateSpec':
        """
        Specifies the job that will be created when executing a CronJob.
        """
        ...

    @property
    @pulumi.getter
    def schedule(self) -> str:
        """
        The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
        """
        ...

    @property
    @pulumi.getter(name="startingDeadlineSeconds")
    def starting_deadline_seconds(self) -> Optional[float]:
        """
        Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.
        """
        ...

    @property
    @pulumi.getter(name="successfulJobsHistoryLimit")
    def successful_jobs_history_limit(self) -> Optional[float]:
        """
        The number of successful finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified. Defaults to 3.
        """
        ...

    @property
    @pulumi.getter
    def suspend(self) -> Optional[bool]:
        """
        This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.  Defaults to false.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CronJobStatus(dict):
    """
    CronJobStatus represents the current state of a cron job.
    """
    @property
    @pulumi.getter
    def active(self) -> Optional[List['_core.v1.outputs.ObjectReference']]:
        """
        A list of pointers to currently running jobs.
        """
        ...

    @property
    @pulumi.getter(name="lastScheduleTime")
    def last_schedule_time(self) -> Optional[str]:
        """
        Information when was the last time the job was successfully scheduled.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class JobTemplateSpec(dict):
    """
    JobTemplateSpec describes the data a Job should have when created from a template
    """
    @property
    @pulumi.getter
    def metadata(self) -> Optional['_meta.v1.outputs.ObjectMeta']:
        """
        Standard object's metadata of the jobs created from this template. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        ...

    @property
    @pulumi.getter
    def spec(self) -> Optional['_batch.v1.outputs.JobSpec']:
        """
        Specification of the desired behavior of the job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


