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

@pulumi.input_type
class CrossVersionObjectReferenceArgs:
    kind: pulumi.Input[str] = pulumi.input_property("kind")
    """
    Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds"
    """
    name: pulumi.Input[str] = pulumi.input_property("name")
    """
    Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names
    """
    api_version: Optional[pulumi.Input[str]] = pulumi.input_property("apiVersion")
    """
    API version of the referent
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, kind: pulumi.Input[str], name: pulumi.Input[str], api_version: Optional[pulumi.Input[str]] = None) -> None:
        """
        CrossVersionObjectReference contains enough information to let you identify the referred resource.
        :param pulumi.Input[str] kind: Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds"
        :param pulumi.Input[str] name: Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names
        :param pulumi.Input[str] api_version: API version of the referent
        """
        __self__.kind = kind
        __self__.name = name
        __self__.api_version = api_version

@pulumi.input_type
class HorizontalPodAutoscalerArgs:
    api_version: Optional[pulumi.Input[str]] = pulumi.input_property("apiVersion")
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    """
    kind: Optional[pulumi.Input[str]] = pulumi.input_property("kind")
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    """
    metadata: Optional[pulumi.Input['ObjectMetaArgs']] = pulumi.input_property("metadata")
    """
    Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
    """
    spec: Optional[pulumi.Input['HorizontalPodAutoscalerSpecArgs']] = pulumi.input_property("spec")
    """
    behaviour of autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
    """
    status: Optional[pulumi.Input['HorizontalPodAutoscalerStatusArgs']] = pulumi.input_property("status")
    """
    current information about the autoscaler.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, api_version: Optional[pulumi.Input[str]] = None, kind: Optional[pulumi.Input[str]] = None, metadata: Optional[pulumi.Input['ObjectMetaArgs']] = None, spec: Optional[pulumi.Input['HorizontalPodAutoscalerSpecArgs']] = None, status: Optional[pulumi.Input['HorizontalPodAutoscalerStatusArgs']] = None) -> None:
        """
        configuration of a horizontal pod autoscaler.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['ObjectMetaArgs'] metadata: Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input['HorizontalPodAutoscalerSpecArgs'] spec: behaviour of autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
        :param pulumi.Input['HorizontalPodAutoscalerStatusArgs'] status: current information about the autoscaler.
        """
        __self__.api_version = 'autoscaling/v1'
        __self__.kind = 'HorizontalPodAutoscaler'
        __self__.metadata = metadata
        __self__.spec = spec
        __self__.status = status

@pulumi.input_type
class HorizontalPodAutoscalerSpecArgs:
    max_replicas: pulumi.Input[float] = pulumi.input_property("maxReplicas")
    """
    upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
    """
    scale_target_ref: pulumi.Input['CrossVersionObjectReferenceArgs'] = pulumi.input_property("scaleTargetRef")
    """
    reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.
    """
    min_replicas: Optional[pulumi.Input[float]] = pulumi.input_property("minReplicas")
    """
    minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
    """
    target_cpu_utilization_percentage: Optional[pulumi.Input[float]] = pulumi.input_property("targetCPUUtilizationPercentage")
    """
    target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, max_replicas: pulumi.Input[float], scale_target_ref: pulumi.Input['CrossVersionObjectReferenceArgs'], min_replicas: Optional[pulumi.Input[float]] = None, target_cpu_utilization_percentage: Optional[pulumi.Input[float]] = None) -> None:
        """
        specification of a horizontal pod autoscaler.
        :param pulumi.Input[float] max_replicas: upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
        :param pulumi.Input['CrossVersionObjectReferenceArgs'] scale_target_ref: reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.
        :param pulumi.Input[float] min_replicas: minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
        :param pulumi.Input[float] target_cpu_utilization_percentage: target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
        """
        __self__.max_replicas = max_replicas
        __self__.scale_target_ref = scale_target_ref
        __self__.min_replicas = min_replicas
        __self__.target_cpu_utilization_percentage = target_cpu_utilization_percentage

@pulumi.input_type
class HorizontalPodAutoscalerStatusArgs:
    current_replicas: pulumi.Input[float] = pulumi.input_property("currentReplicas")
    """
    current number of replicas of pods managed by this autoscaler.
    """
    desired_replicas: pulumi.Input[float] = pulumi.input_property("desiredReplicas")
    """
    desired number of replicas of pods managed by this autoscaler.
    """
    current_cpu_utilization_percentage: Optional[pulumi.Input[float]] = pulumi.input_property("currentCPUUtilizationPercentage")
    """
    current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.
    """
    last_scale_time: Optional[pulumi.Input[str]] = pulumi.input_property("lastScaleTime")
    """
    last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed.
    """
    observed_generation: Optional[pulumi.Input[float]] = pulumi.input_property("observedGeneration")
    """
    most recent generation observed by this autoscaler.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, current_replicas: pulumi.Input[float], desired_replicas: pulumi.Input[float], current_cpu_utilization_percentage: Optional[pulumi.Input[float]] = None, last_scale_time: Optional[pulumi.Input[str]] = None, observed_generation: Optional[pulumi.Input[float]] = None) -> None:
        """
        current status of a horizontal pod autoscaler
        :param pulumi.Input[float] current_replicas: current number of replicas of pods managed by this autoscaler.
        :param pulumi.Input[float] desired_replicas: desired number of replicas of pods managed by this autoscaler.
        :param pulumi.Input[float] current_cpu_utilization_percentage: current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.
        :param pulumi.Input[str] last_scale_time: last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed.
        :param pulumi.Input[float] observed_generation: most recent generation observed by this autoscaler.
        """
        __self__.current_replicas = current_replicas
        __self__.desired_replicas = desired_replicas
        __self__.current_cpu_utilization_percentage = current_cpu_utilization_percentage
        __self__.last_scale_time = last_scale_time
        __self__.observed_generation = observed_generation

