# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ... import meta as _meta

__all__ = [
    'CrossVersionObjectReference',
    'HorizontalPodAutoscaler',
    'HorizontalPodAutoscalerSpec',
    'HorizontalPodAutoscalerStatus',
]

@pulumi.output_type
class CrossVersionObjectReference(dict):
    """
    CrossVersionObjectReference contains enough information to let you identify the referred resource.
    """
    def __init__(__self__, *,
                 kind: str,
                 name: str,
                 api_version: Optional[str] = None):
        """
        CrossVersionObjectReference contains enough information to let you identify the referred resource.
        :param str kind: Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds"
        :param str name: Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names
        :param str api_version: API version of the referent
        """
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "name", name)
        if api_version is not None:
            pulumi.set(__self__, "api_version", api_version)

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds"
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[str]:
        """
        API version of the referent
        """
        return pulumi.get(self, "api_version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class HorizontalPodAutoscaler(dict):
    """
    configuration of a horizontal pod autoscaler.
    """
    def __init__(__self__, *,
                 api_version: Optional[str] = None,
                 kind: Optional[str] = None,
                 metadata: Optional['_meta.v1.outputs.ObjectMeta'] = None,
                 spec: Optional['outputs.HorizontalPodAutoscalerSpec'] = None,
                 status: Optional['outputs.HorizontalPodAutoscalerStatus'] = None):
        """
        configuration of a horizontal pod autoscaler.
        :param str api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param str kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param '_meta.v1.ObjectMetaArgs' metadata: Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param 'HorizontalPodAutoscalerSpecArgs' spec: behaviour of autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
        :param 'HorizontalPodAutoscalerStatusArgs' status: current information about the autoscaler.
        """
        if api_version is not None:
            pulumi.set(__self__, "api_version", 'autoscaling/v1')
        if kind is not None:
            pulumi.set(__self__, "kind", 'HorizontalPodAutoscaler')
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
        Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def spec(self) -> Optional['outputs.HorizontalPodAutoscalerSpec']:
        """
        behaviour of autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter
    def status(self) -> Optional['outputs.HorizontalPodAutoscalerStatus']:
        """
        current information about the autoscaler.
        """
        return pulumi.get(self, "status")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class HorizontalPodAutoscalerSpec(dict):
    """
    specification of a horizontal pod autoscaler.
    """
    def __init__(__self__, *,
                 max_replicas: int,
                 scale_target_ref: 'outputs.CrossVersionObjectReference',
                 min_replicas: Optional[int] = None,
                 target_cpu_utilization_percentage: Optional[int] = None):
        """
        specification of a horizontal pod autoscaler.
        :param int max_replicas: upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
        :param 'CrossVersionObjectReferenceArgs' scale_target_ref: reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.
        :param int min_replicas: minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
        :param int target_cpu_utilization_percentage: target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
        """
        pulumi.set(__self__, "max_replicas", max_replicas)
        pulumi.set(__self__, "scale_target_ref", scale_target_ref)
        if min_replicas is not None:
            pulumi.set(__self__, "min_replicas", min_replicas)
        if target_cpu_utilization_percentage is not None:
            pulumi.set(__self__, "target_cpu_utilization_percentage", target_cpu_utilization_percentage)

    @property
    @pulumi.getter(name="maxReplicas")
    def max_replicas(self) -> int:
        """
        upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
        """
        return pulumi.get(self, "max_replicas")

    @property
    @pulumi.getter(name="scaleTargetRef")
    def scale_target_ref(self) -> 'outputs.CrossVersionObjectReference':
        """
        reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.
        """
        return pulumi.get(self, "scale_target_ref")

    @property
    @pulumi.getter(name="minReplicas")
    def min_replicas(self) -> Optional[int]:
        """
        minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
        """
        return pulumi.get(self, "min_replicas")

    @property
    @pulumi.getter(name="targetCPUUtilizationPercentage")
    def target_cpu_utilization_percentage(self) -> Optional[int]:
        """
        target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
        """
        return pulumi.get(self, "target_cpu_utilization_percentage")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class HorizontalPodAutoscalerStatus(dict):
    """
    current status of a horizontal pod autoscaler
    """
    def __init__(__self__, *,
                 current_replicas: int,
                 desired_replicas: int,
                 current_cpu_utilization_percentage: Optional[int] = None,
                 last_scale_time: Optional[str] = None,
                 observed_generation: Optional[int] = None):
        """
        current status of a horizontal pod autoscaler
        :param int current_replicas: current number of replicas of pods managed by this autoscaler.
        :param int desired_replicas: desired number of replicas of pods managed by this autoscaler.
        :param int current_cpu_utilization_percentage: current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.
        :param str last_scale_time: last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed.
        :param int observed_generation: most recent generation observed by this autoscaler.
        """
        pulumi.set(__self__, "current_replicas", current_replicas)
        pulumi.set(__self__, "desired_replicas", desired_replicas)
        if current_cpu_utilization_percentage is not None:
            pulumi.set(__self__, "current_cpu_utilization_percentage", current_cpu_utilization_percentage)
        if last_scale_time is not None:
            pulumi.set(__self__, "last_scale_time", last_scale_time)
        if observed_generation is not None:
            pulumi.set(__self__, "observed_generation", observed_generation)

    @property
    @pulumi.getter(name="currentReplicas")
    def current_replicas(self) -> int:
        """
        current number of replicas of pods managed by this autoscaler.
        """
        return pulumi.get(self, "current_replicas")

    @property
    @pulumi.getter(name="desiredReplicas")
    def desired_replicas(self) -> int:
        """
        desired number of replicas of pods managed by this autoscaler.
        """
        return pulumi.get(self, "desired_replicas")

    @property
    @pulumi.getter(name="currentCPUUtilizationPercentage")
    def current_cpu_utilization_percentage(self) -> Optional[int]:
        """
        current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.
        """
        return pulumi.get(self, "current_cpu_utilization_percentage")

    @property
    @pulumi.getter(name="lastScaleTime")
    def last_scale_time(self) -> Optional[str]:
        """
        last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed.
        """
        return pulumi.get(self, "last_scale_time")

    @property
    @pulumi.getter(name="observedGeneration")
    def observed_generation(self) -> Optional[int]:
        """
        most recent generation observed by this autoscaler.
        """
        return pulumi.get(self, "observed_generation")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


