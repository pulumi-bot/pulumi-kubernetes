# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from ... import _utilities, _tables
from ... import batch as _batch
from ... import core as _core
from ... import meta as _meta

@pulumi.input_type
class CronJobArgs:
    api_version: Optional[pulumi.Input[str]] = pulumi.input_property("apiVersion")
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    """
    kind: Optional[pulumi.Input[str]] = pulumi.input_property("kind")
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    """
    metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = pulumi.input_property("metadata")
    """
    Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
    """
    spec: Optional[pulumi.Input['CronJobSpecArgs']] = pulumi.input_property("spec")
    """
    Specification of the desired behavior of a cron job, including the schedule. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
    """
    status: Optional[pulumi.Input['CronJobStatusArgs']] = pulumi.input_property("status")
    """
    Current status of a cron job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, api_version: Optional[pulumi.Input[str]] = None, kind: Optional[pulumi.Input[str]] = None, metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None, spec: Optional[pulumi.Input['CronJobSpecArgs']] = None, status: Optional[pulumi.Input['CronJobStatusArgs']] = None) -> None:
        """
        CronJob represents the configuration of a single cron job.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input['CronJobSpecArgs'] spec: Specification of the desired behavior of a cron job, including the schedule. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        :param pulumi.Input['CronJobStatusArgs'] status: Current status of a cron job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        __self__.api_version = 'batch/v1beta1'
        __self__.kind = 'CronJob'
        __self__.metadata = metadata
        __self__.spec = spec
        __self__.status = status

@pulumi.input_type
class CronJobSpecArgs:
    job_template: pulumi.Input['JobTemplateSpecArgs'] = pulumi.input_property("jobTemplate")
    """
    Specifies the job that will be created when executing a CronJob.
    """
    schedule: pulumi.Input[str] = pulumi.input_property("schedule")
    """
    The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
    """
    concurrency_policy: Optional[pulumi.Input[str]] = pulumi.input_property("concurrencyPolicy")
    """
    Specifies how to treat concurrent executions of a Job. Valid values are: - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one
    """
    failed_jobs_history_limit: Optional[pulumi.Input[float]] = pulumi.input_property("failedJobsHistoryLimit")
    """
    The number of failed finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
    """
    starting_deadline_seconds: Optional[pulumi.Input[float]] = pulumi.input_property("startingDeadlineSeconds")
    """
    Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.
    """
    successful_jobs_history_limit: Optional[pulumi.Input[float]] = pulumi.input_property("successfulJobsHistoryLimit")
    """
    The number of successful finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified. Defaults to 3.
    """
    suspend: Optional[pulumi.Input[bool]] = pulumi.input_property("suspend")
    """
    This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.  Defaults to false.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, job_template: pulumi.Input['JobTemplateSpecArgs'], schedule: pulumi.Input[str], concurrency_policy: Optional[pulumi.Input[str]] = None, failed_jobs_history_limit: Optional[pulumi.Input[float]] = None, starting_deadline_seconds: Optional[pulumi.Input[float]] = None, successful_jobs_history_limit: Optional[pulumi.Input[float]] = None, suspend: Optional[pulumi.Input[bool]] = None) -> None:
        """
        CronJobSpec describes how the job execution will look like and when it will actually run.
        :param pulumi.Input['JobTemplateSpecArgs'] job_template: Specifies the job that will be created when executing a CronJob.
        :param pulumi.Input[str] schedule: The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
        :param pulumi.Input[str] concurrency_policy: Specifies how to treat concurrent executions of a Job. Valid values are: - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one
        :param pulumi.Input[float] failed_jobs_history_limit: The number of failed finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
        :param pulumi.Input[float] starting_deadline_seconds: Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.
        :param pulumi.Input[float] successful_jobs_history_limit: The number of successful finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified. Defaults to 3.
        :param pulumi.Input[bool] suspend: This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.  Defaults to false.
        """
        __self__.job_template = job_template
        __self__.schedule = schedule
        __self__.concurrency_policy = concurrency_policy
        __self__.failed_jobs_history_limit = failed_jobs_history_limit
        __self__.starting_deadline_seconds = starting_deadline_seconds
        __self__.successful_jobs_history_limit = successful_jobs_history_limit
        __self__.suspend = suspend

@pulumi.input_type
class CronJobStatusArgs:
    active: Optional[pulumi.Input[List[pulumi.Input['_core.v1.ObjectReferenceArgs']]]] = pulumi.input_property("active")
    """
    A list of pointers to currently running jobs.
    """
    last_schedule_time: Optional[pulumi.Input[str]] = pulumi.input_property("lastScheduleTime")
    """
    Information when was the last time the job was successfully scheduled.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, active: Optional[pulumi.Input[List[pulumi.Input['_core.v1.ObjectReferenceArgs']]]] = None, last_schedule_time: Optional[pulumi.Input[str]] = None) -> None:
        """
        CronJobStatus represents the current state of a cron job.
        :param pulumi.Input[List[pulumi.Input['_core.v1.ObjectReferenceArgs']]] active: A list of pointers to currently running jobs.
        :param pulumi.Input[str] last_schedule_time: Information when was the last time the job was successfully scheduled.
        """
        __self__.active = active
        __self__.last_schedule_time = last_schedule_time

@pulumi.input_type
class JobTemplateSpecArgs:
    metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = pulumi.input_property("metadata")
    """
    Standard object's metadata of the jobs created from this template. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
    """
    spec: Optional[pulumi.Input['_batch.v1.JobSpecArgs']] = pulumi.input_property("spec")
    """
    Specification of the desired behavior of the job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None, spec: Optional[pulumi.Input['_batch.v1.JobSpecArgs']] = None) -> None:
        """
        JobTemplateSpec describes the data a Job should have when created from a template
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object's metadata of the jobs created from this template. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input['_batch.v1.JobSpecArgs'] spec: Specification of the desired behavior of the job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        __self__.metadata = metadata
        __self__.spec = spec

