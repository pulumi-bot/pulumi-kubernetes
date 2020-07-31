# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ... import meta as _meta

__all__ = [
    'CrossVersionObjectReference',
    'ExternalMetricSource',
    'ExternalMetricStatus',
    'HorizontalPodAutoscaler',
    'HorizontalPodAutoscalerCondition',
    'HorizontalPodAutoscalerSpec',
    'HorizontalPodAutoscalerStatus',
    'MetricSpec',
    'MetricStatus',
    'ObjectMetricSource',
    'ObjectMetricStatus',
    'PodsMetricSource',
    'PodsMetricStatus',
    'ResourceMetricSource',
    'ResourceMetricStatus',
]

@pulumi.output_type
class CrossVersionObjectReference(dict):
    """
    CrossVersionObjectReference contains enough information to let you identify the referred resource.
    """
    api_version: Optional[str] = pulumi.output_property("apiVersion")
    """
    API version of the referent
    """
    kind: str = pulumi.output_property("kind")
    """
    Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds"
    """
    name: str = pulumi.output_property("name")
    """
    Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ExternalMetricSource(dict):
    """
    ExternalMetricSource indicates how to scale on a metric not associated with any Kubernetes object (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster). Exactly one "target" type should be set.
    """
    metric_name: str = pulumi.output_property("metricName")
    """
    metricName is the name of the metric in question.
    """
    metric_selector: Optional['_meta.v1.outputs.LabelSelector'] = pulumi.output_property("metricSelector")
    """
    metricSelector is used to identify a specific time series within a given metric.
    """
    target_average_value: Optional[str] = pulumi.output_property("targetAverageValue")
    """
    targetAverageValue is the target per-pod value of global metric (as a quantity). Mutually exclusive with TargetValue.
    """
    target_value: Optional[str] = pulumi.output_property("targetValue")
    """
    targetValue is the target value of the metric (as a quantity). Mutually exclusive with TargetAverageValue.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ExternalMetricStatus(dict):
    """
    ExternalMetricStatus indicates the current value of a global metric not associated with any Kubernetes object.
    """
    current_average_value: Optional[str] = pulumi.output_property("currentAverageValue")
    """
    currentAverageValue is the current value of metric averaged over autoscaled pods.
    """
    current_value: str = pulumi.output_property("currentValue")
    """
    currentValue is the current value of the metric (as a quantity)
    """
    metric_name: str = pulumi.output_property("metricName")
    """
    metricName is the name of a metric used for autoscaling in metric system.
    """
    metric_selector: Optional['_meta.v1.outputs.LabelSelector'] = pulumi.output_property("metricSelector")
    """
    metricSelector is used to identify a specific time series within a given metric.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class HorizontalPodAutoscaler(dict):
    """
    HorizontalPodAutoscaler is the configuration for a horizontal pod autoscaler, which automatically manages the replica count of any resource implementing the scale subresource based on the metrics specified.
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
    metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
    """
    spec: Optional['outputs.HorizontalPodAutoscalerSpec'] = pulumi.output_property("spec")
    """
    spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
    """
    status: Optional['outputs.HorizontalPodAutoscalerStatus'] = pulumi.output_property("status")
    """
    status is the current information about the autoscaler.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class HorizontalPodAutoscalerCondition(dict):
    """
    HorizontalPodAutoscalerCondition describes the state of a HorizontalPodAutoscaler at a certain point.
    """
    last_transition_time: Optional[str] = pulumi.output_property("lastTransitionTime")
    """
    lastTransitionTime is the last time the condition transitioned from one status to another
    """
    message: Optional[str] = pulumi.output_property("message")
    """
    message is a human-readable explanation containing details about the transition
    """
    reason: Optional[str] = pulumi.output_property("reason")
    """
    reason is the reason for the condition's last transition.
    """
    status: str = pulumi.output_property("status")
    """
    status is the status of the condition (True, False, Unknown)
    """
    type: str = pulumi.output_property("type")
    """
    type describes the current condition
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class HorizontalPodAutoscalerSpec(dict):
    """
    HorizontalPodAutoscalerSpec describes the desired functionality of the HorizontalPodAutoscaler.
    """
    max_replicas: float = pulumi.output_property("maxReplicas")
    """
    maxReplicas is the upper limit for the number of replicas to which the autoscaler can scale up. It cannot be less that minReplicas.
    """
    metrics: Optional[List['outputs.MetricSpec']] = pulumi.output_property("metrics")
    """
    metrics contains the specifications for which to use to calculate the desired replica count (the maximum replica count across all metrics will be used).  The desired replica count is calculated multiplying the ratio between the target value and the current value by the current number of pods.  Ergo, metrics used must decrease as the pod count is increased, and vice-versa.  See the individual metric source types for more information about how each type of metric must respond.
    """
    min_replicas: Optional[float] = pulumi.output_property("minReplicas")
    """
    minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
    """
    scale_target_ref: 'outputs.CrossVersionObjectReference' = pulumi.output_property("scaleTargetRef")
    """
    scaleTargetRef points to the target resource to scale, and is used to the pods for which metrics should be collected, as well as to actually change the replica count.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class HorizontalPodAutoscalerStatus(dict):
    """
    HorizontalPodAutoscalerStatus describes the current status of a horizontal pod autoscaler.
    """
    conditions: List['outputs.HorizontalPodAutoscalerCondition'] = pulumi.output_property("conditions")
    """
    conditions is the set of conditions required for this autoscaler to scale its target, and indicates whether or not those conditions are met.
    """
    current_metrics: Optional[List['outputs.MetricStatus']] = pulumi.output_property("currentMetrics")
    """
    currentMetrics is the last read state of the metrics used by this autoscaler.
    """
    current_replicas: float = pulumi.output_property("currentReplicas")
    """
    currentReplicas is current number of replicas of pods managed by this autoscaler, as last seen by the autoscaler.
    """
    desired_replicas: float = pulumi.output_property("desiredReplicas")
    """
    desiredReplicas is the desired number of replicas of pods managed by this autoscaler, as last calculated by the autoscaler.
    """
    last_scale_time: Optional[str] = pulumi.output_property("lastScaleTime")
    """
    lastScaleTime is the last time the HorizontalPodAutoscaler scaled the number of pods, used by the autoscaler to control how often the number of pods is changed.
    """
    observed_generation: Optional[float] = pulumi.output_property("observedGeneration")
    """
    observedGeneration is the most recent generation observed by this autoscaler.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MetricSpec(dict):
    """
    MetricSpec specifies how to scale based on a single metric (only `type` and one other matching field should be set at once).
    """
    external: Optional['outputs.ExternalMetricSource'] = pulumi.output_property("external")
    """
    external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).
    """
    object: Optional['outputs.ObjectMetricSource'] = pulumi.output_property("object")
    """
    object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).
    """
    pods: Optional['outputs.PodsMetricSource'] = pulumi.output_property("pods")
    """
    pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.
    """
    resource: Optional['outputs.ResourceMetricSource'] = pulumi.output_property("resource")
    """
    resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
    """
    type: str = pulumi.output_property("type")
    """
    type is the type of metric source.  It should be one of "Object", "Pods" or "Resource", each mapping to a matching field in the object.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MetricStatus(dict):
    """
    MetricStatus describes the last-read state of a single metric.
    """
    external: Optional['outputs.ExternalMetricStatus'] = pulumi.output_property("external")
    """
    external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).
    """
    object: Optional['outputs.ObjectMetricStatus'] = pulumi.output_property("object")
    """
    object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).
    """
    pods: Optional['outputs.PodsMetricStatus'] = pulumi.output_property("pods")
    """
    pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.
    """
    resource: Optional['outputs.ResourceMetricStatus'] = pulumi.output_property("resource")
    """
    resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
    """
    type: str = pulumi.output_property("type")
    """
    type is the type of metric source.  It will be one of "Object", "Pods" or "Resource", each corresponds to a matching field in the object.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ObjectMetricSource(dict):
    """
    ObjectMetricSource indicates how to scale on a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
    """
    average_value: Optional[str] = pulumi.output_property("averageValue")
    """
    averageValue is the target value of the average of the metric across all relevant pods (as a quantity)
    """
    metric_name: str = pulumi.output_property("metricName")
    """
    metricName is the name of the metric in question.
    """
    selector: Optional['_meta.v1.outputs.LabelSelector'] = pulumi.output_property("selector")
    """
    selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping When unset, just the metricName will be used to gather metrics.
    """
    target: 'outputs.CrossVersionObjectReference' = pulumi.output_property("target")
    """
    target is the described Kubernetes object.
    """
    target_value: str = pulumi.output_property("targetValue")
    """
    targetValue is the target value of the metric (as a quantity).
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ObjectMetricStatus(dict):
    """
    ObjectMetricStatus indicates the current value of a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
    """
    average_value: Optional[str] = pulumi.output_property("averageValue")
    """
    averageValue is the current value of the average of the metric across all relevant pods (as a quantity)
    """
    current_value: str = pulumi.output_property("currentValue")
    """
    currentValue is the current value of the metric (as a quantity).
    """
    metric_name: str = pulumi.output_property("metricName")
    """
    metricName is the name of the metric in question.
    """
    selector: Optional['_meta.v1.outputs.LabelSelector'] = pulumi.output_property("selector")
    """
    selector is the string-encoded form of a standard kubernetes label selector for the given metric When set in the ObjectMetricSource, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.
    """
    target: 'outputs.CrossVersionObjectReference' = pulumi.output_property("target")
    """
    target is the described Kubernetes object.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PodsMetricSource(dict):
    """
    PodsMetricSource indicates how to scale on a metric describing each pod in the current scale target (for example, transactions-processed-per-second). The values will be averaged together before being compared to the target value.
    """
    metric_name: str = pulumi.output_property("metricName")
    """
    metricName is the name of the metric in question
    """
    selector: Optional['_meta.v1.outputs.LabelSelector'] = pulumi.output_property("selector")
    """
    selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping When unset, just the metricName will be used to gather metrics.
    """
    target_average_value: str = pulumi.output_property("targetAverageValue")
    """
    targetAverageValue is the target value of the average of the metric across all relevant pods (as a quantity)
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PodsMetricStatus(dict):
    """
    PodsMetricStatus indicates the current value of a metric describing each pod in the current scale target (for example, transactions-processed-per-second).
    """
    current_average_value: str = pulumi.output_property("currentAverageValue")
    """
    currentAverageValue is the current value of the average of the metric across all relevant pods (as a quantity)
    """
    metric_name: str = pulumi.output_property("metricName")
    """
    metricName is the name of the metric in question
    """
    selector: Optional['_meta.v1.outputs.LabelSelector'] = pulumi.output_property("selector")
    """
    selector is the string-encoded form of a standard kubernetes label selector for the given metric When set in the PodsMetricSource, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ResourceMetricSource(dict):
    """
    ResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  The values will be averaged together before being compared to the target.  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.  Only one "target" type should be set.
    """
    name: str = pulumi.output_property("name")
    """
    name is the name of the resource in question.
    """
    target_average_utilization: Optional[float] = pulumi.output_property("targetAverageUtilization")
    """
    targetAverageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.
    """
    target_average_value: Optional[str] = pulumi.output_property("targetAverageValue")
    """
    targetAverageValue is the target value of the average of the resource metric across all relevant pods, as a raw value (instead of as a percentage of the request), similar to the "pods" metric source type.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ResourceMetricStatus(dict):
    """
    ResourceMetricStatus indicates the current value of a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
    """
    current_average_utilization: Optional[float] = pulumi.output_property("currentAverageUtilization")
    """
    currentAverageUtilization is the current value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.  It will only be present if `targetAverageValue` was set in the corresponding metric specification.
    """
    current_average_value: str = pulumi.output_property("currentAverageValue")
    """
    currentAverageValue is the current value of the average of the resource metric across all relevant pods, as a raw value (instead of as a percentage of the request), similar to the "pods" metric source type. It will always be set, regardless of the corresponding metric specification.
    """
    name: str = pulumi.output_property("name")
    """
    name is the name of the resource in question.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


