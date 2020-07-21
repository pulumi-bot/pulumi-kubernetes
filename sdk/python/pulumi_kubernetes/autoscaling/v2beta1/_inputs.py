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
class ExternalMetricSourceArgs:
    metric_name: pulumi.Input[str] = pulumi.input_property("metricName")
    """
    metricName is the name of the metric in question.
    """
    metric_selector: Optional[pulumi.Input['LabelSelectorArgs']] = pulumi.input_property("metricSelector")
    """
    metricSelector is used to identify a specific time series within a given metric.
    """
    target_average_value: Optional[pulumi.Input[str]] = pulumi.input_property("targetAverageValue")
    """
    targetAverageValue is the target per-pod value of global metric (as a quantity). Mutually exclusive with TargetValue.
    """
    target_value: Optional[pulumi.Input[str]] = pulumi.input_property("targetValue")
    """
    targetValue is the target value of the metric (as a quantity). Mutually exclusive with TargetAverageValue.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, metric_name: pulumi.Input[str], metric_selector: Optional[pulumi.Input['LabelSelectorArgs']] = None, target_average_value: Optional[pulumi.Input[str]] = None, target_value: Optional[pulumi.Input[str]] = None) -> None:
        """
        ExternalMetricSource indicates how to scale on a metric not associated with any Kubernetes object (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster). Exactly one "target" type should be set.
        :param pulumi.Input[str] metric_name: metricName is the name of the metric in question.
        :param pulumi.Input['LabelSelectorArgs'] metric_selector: metricSelector is used to identify a specific time series within a given metric.
        :param pulumi.Input[str] target_average_value: targetAverageValue is the target per-pod value of global metric (as a quantity). Mutually exclusive with TargetValue.
        :param pulumi.Input[str] target_value: targetValue is the target value of the metric (as a quantity). Mutually exclusive with TargetAverageValue.
        """
        __self__.metric_name = metric_name
        __self__.metric_selector = metric_selector
        __self__.target_average_value = target_average_value
        __self__.target_value = target_value

@pulumi.input_type
class ExternalMetricStatusArgs:
    current_value: pulumi.Input[str] = pulumi.input_property("currentValue")
    """
    currentValue is the current value of the metric (as a quantity)
    """
    metric_name: pulumi.Input[str] = pulumi.input_property("metricName")
    """
    metricName is the name of a metric used for autoscaling in metric system.
    """
    current_average_value: Optional[pulumi.Input[str]] = pulumi.input_property("currentAverageValue")
    """
    currentAverageValue is the current value of metric averaged over autoscaled pods.
    """
    metric_selector: Optional[pulumi.Input['LabelSelectorArgs']] = pulumi.input_property("metricSelector")
    """
    metricSelector is used to identify a specific time series within a given metric.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, current_value: pulumi.Input[str], metric_name: pulumi.Input[str], current_average_value: Optional[pulumi.Input[str]] = None, metric_selector: Optional[pulumi.Input['LabelSelectorArgs']] = None) -> None:
        """
        ExternalMetricStatus indicates the current value of a global metric not associated with any Kubernetes object.
        :param pulumi.Input[str] current_value: currentValue is the current value of the metric (as a quantity)
        :param pulumi.Input[str] metric_name: metricName is the name of a metric used for autoscaling in metric system.
        :param pulumi.Input[str] current_average_value: currentAverageValue is the current value of metric averaged over autoscaled pods.
        :param pulumi.Input['LabelSelectorArgs'] metric_selector: metricSelector is used to identify a specific time series within a given metric.
        """
        __self__.current_value = current_value
        __self__.metric_name = metric_name
        __self__.current_average_value = current_average_value
        __self__.metric_selector = metric_selector

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
    metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
    """
    spec: Optional[pulumi.Input['HorizontalPodAutoscalerSpecArgs']] = pulumi.input_property("spec")
    """
    spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
    """
    status: Optional[pulumi.Input['HorizontalPodAutoscalerStatusArgs']] = pulumi.input_property("status")
    """
    status is the current information about the autoscaler.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, api_version: Optional[pulumi.Input[str]] = None, kind: Optional[pulumi.Input[str]] = None, metadata: Optional[pulumi.Input['ObjectMetaArgs']] = None, spec: Optional[pulumi.Input['HorizontalPodAutoscalerSpecArgs']] = None, status: Optional[pulumi.Input['HorizontalPodAutoscalerStatusArgs']] = None) -> None:
        """
        HorizontalPodAutoscaler is the configuration for a horizontal pod autoscaler, which automatically manages the replica count of any resource implementing the scale subresource based on the metrics specified.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['ObjectMetaArgs'] metadata: metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input['HorizontalPodAutoscalerSpecArgs'] spec: spec is the specification for the behaviour of the autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
        :param pulumi.Input['HorizontalPodAutoscalerStatusArgs'] status: status is the current information about the autoscaler.
        """
        __self__.api_version = 'autoscaling/v2beta1'
        __self__.kind = 'HorizontalPodAutoscaler'
        __self__.metadata = metadata
        __self__.spec = spec
        __self__.status = status

@pulumi.input_type
class HorizontalPodAutoscalerConditionArgs:
    status: pulumi.Input[str] = pulumi.input_property("status")
    """
    status is the status of the condition (True, False, Unknown)
    """
    type: pulumi.Input[str] = pulumi.input_property("type")
    """
    type describes the current condition
    """
    last_transition_time: Optional[pulumi.Input[str]] = pulumi.input_property("lastTransitionTime")
    """
    lastTransitionTime is the last time the condition transitioned from one status to another
    """
    message: Optional[pulumi.Input[str]] = pulumi.input_property("message")
    """
    message is a human-readable explanation containing details about the transition
    """
    reason: Optional[pulumi.Input[str]] = pulumi.input_property("reason")
    """
    reason is the reason for the condition's last transition.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, status: pulumi.Input[str], type: pulumi.Input[str], last_transition_time: Optional[pulumi.Input[str]] = None, message: Optional[pulumi.Input[str]] = None, reason: Optional[pulumi.Input[str]] = None) -> None:
        """
        HorizontalPodAutoscalerCondition describes the state of a HorizontalPodAutoscaler at a certain point.
        :param pulumi.Input[str] status: status is the status of the condition (True, False, Unknown)
        :param pulumi.Input[str] type: type describes the current condition
        :param pulumi.Input[str] last_transition_time: lastTransitionTime is the last time the condition transitioned from one status to another
        :param pulumi.Input[str] message: message is a human-readable explanation containing details about the transition
        :param pulumi.Input[str] reason: reason is the reason for the condition's last transition.
        """
        __self__.status = status
        __self__.type = type
        __self__.last_transition_time = last_transition_time
        __self__.message = message
        __self__.reason = reason

@pulumi.input_type
class HorizontalPodAutoscalerSpecArgs:
    max_replicas: pulumi.Input[float] = pulumi.input_property("maxReplicas")
    """
    maxReplicas is the upper limit for the number of replicas to which the autoscaler can scale up. It cannot be less that minReplicas.
    """
    scale_target_ref: pulumi.Input['CrossVersionObjectReferenceArgs'] = pulumi.input_property("scaleTargetRef")
    """
    scaleTargetRef points to the target resource to scale, and is used to the pods for which metrics should be collected, as well as to actually change the replica count.
    """
    metrics: Optional[pulumi.Input[List[pulumi.Input['MetricSpecArgs']]]] = pulumi.input_property("metrics")
    """
    metrics contains the specifications for which to use to calculate the desired replica count (the maximum replica count across all metrics will be used).  The desired replica count is calculated multiplying the ratio between the target value and the current value by the current number of pods.  Ergo, metrics used must decrease as the pod count is increased, and vice-versa.  See the individual metric source types for more information about how each type of metric must respond.
    """
    min_replicas: Optional[pulumi.Input[float]] = pulumi.input_property("minReplicas")
    """
    minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, max_replicas: pulumi.Input[float], scale_target_ref: pulumi.Input['CrossVersionObjectReferenceArgs'], metrics: Optional[pulumi.Input[List[pulumi.Input['MetricSpecArgs']]]] = None, min_replicas: Optional[pulumi.Input[float]] = None) -> None:
        """
        HorizontalPodAutoscalerSpec describes the desired functionality of the HorizontalPodAutoscaler.
        :param pulumi.Input[float] max_replicas: maxReplicas is the upper limit for the number of replicas to which the autoscaler can scale up. It cannot be less that minReplicas.
        :param pulumi.Input['CrossVersionObjectReferenceArgs'] scale_target_ref: scaleTargetRef points to the target resource to scale, and is used to the pods for which metrics should be collected, as well as to actually change the replica count.
        :param pulumi.Input[List[pulumi.Input['MetricSpecArgs']]] metrics: metrics contains the specifications for which to use to calculate the desired replica count (the maximum replica count across all metrics will be used).  The desired replica count is calculated multiplying the ratio between the target value and the current value by the current number of pods.  Ergo, metrics used must decrease as the pod count is increased, and vice-versa.  See the individual metric source types for more information about how each type of metric must respond.
        :param pulumi.Input[float] min_replicas: minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
        """
        __self__.max_replicas = max_replicas
        __self__.scale_target_ref = scale_target_ref
        __self__.metrics = metrics
        __self__.min_replicas = min_replicas

@pulumi.input_type
class HorizontalPodAutoscalerStatusArgs:
    conditions: pulumi.Input[List[pulumi.Input['HorizontalPodAutoscalerConditionArgs']]] = pulumi.input_property("conditions")
    """
    conditions is the set of conditions required for this autoscaler to scale its target, and indicates whether or not those conditions are met.
    """
    current_replicas: pulumi.Input[float] = pulumi.input_property("currentReplicas")
    """
    currentReplicas is current number of replicas of pods managed by this autoscaler, as last seen by the autoscaler.
    """
    desired_replicas: pulumi.Input[float] = pulumi.input_property("desiredReplicas")
    """
    desiredReplicas is the desired number of replicas of pods managed by this autoscaler, as last calculated by the autoscaler.
    """
    current_metrics: Optional[pulumi.Input[List[pulumi.Input['MetricStatusArgs']]]] = pulumi.input_property("currentMetrics")
    """
    currentMetrics is the last read state of the metrics used by this autoscaler.
    """
    last_scale_time: Optional[pulumi.Input[str]] = pulumi.input_property("lastScaleTime")
    """
    lastScaleTime is the last time the HorizontalPodAutoscaler scaled the number of pods, used by the autoscaler to control how often the number of pods is changed.
    """
    observed_generation: Optional[pulumi.Input[float]] = pulumi.input_property("observedGeneration")
    """
    observedGeneration is the most recent generation observed by this autoscaler.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, conditions: pulumi.Input[List[pulumi.Input['HorizontalPodAutoscalerConditionArgs']]], current_replicas: pulumi.Input[float], desired_replicas: pulumi.Input[float], current_metrics: Optional[pulumi.Input[List[pulumi.Input['MetricStatusArgs']]]] = None, last_scale_time: Optional[pulumi.Input[str]] = None, observed_generation: Optional[pulumi.Input[float]] = None) -> None:
        """
        HorizontalPodAutoscalerStatus describes the current status of a horizontal pod autoscaler.
        :param pulumi.Input[List[pulumi.Input['HorizontalPodAutoscalerConditionArgs']]] conditions: conditions is the set of conditions required for this autoscaler to scale its target, and indicates whether or not those conditions are met.
        :param pulumi.Input[float] current_replicas: currentReplicas is current number of replicas of pods managed by this autoscaler, as last seen by the autoscaler.
        :param pulumi.Input[float] desired_replicas: desiredReplicas is the desired number of replicas of pods managed by this autoscaler, as last calculated by the autoscaler.
        :param pulumi.Input[List[pulumi.Input['MetricStatusArgs']]] current_metrics: currentMetrics is the last read state of the metrics used by this autoscaler.
        :param pulumi.Input[str] last_scale_time: lastScaleTime is the last time the HorizontalPodAutoscaler scaled the number of pods, used by the autoscaler to control how often the number of pods is changed.
        :param pulumi.Input[float] observed_generation: observedGeneration is the most recent generation observed by this autoscaler.
        """
        __self__.conditions = conditions
        __self__.current_replicas = current_replicas
        __self__.desired_replicas = desired_replicas
        __self__.current_metrics = current_metrics
        __self__.last_scale_time = last_scale_time
        __self__.observed_generation = observed_generation

@pulumi.input_type
class MetricSpecArgs:
    type: pulumi.Input[str] = pulumi.input_property("type")
    """
    type is the type of metric source.  It should be one of "Object", "Pods" or "Resource", each mapping to a matching field in the object.
    """
    external: Optional[pulumi.Input['ExternalMetricSourceArgs']] = pulumi.input_property("external")
    """
    external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).
    """
    object: Optional[pulumi.Input['ObjectMetricSourceArgs']] = pulumi.input_property("object")
    """
    object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).
    """
    pods: Optional[pulumi.Input['PodsMetricSourceArgs']] = pulumi.input_property("pods")
    """
    pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.
    """
    resource: Optional[pulumi.Input['ResourceMetricSourceArgs']] = pulumi.input_property("resource")
    """
    resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, type: pulumi.Input[str], external: Optional[pulumi.Input['ExternalMetricSourceArgs']] = None, object: Optional[pulumi.Input['ObjectMetricSourceArgs']] = None, pods: Optional[pulumi.Input['PodsMetricSourceArgs']] = None, resource: Optional[pulumi.Input['ResourceMetricSourceArgs']] = None) -> None:
        """
        MetricSpec specifies how to scale based on a single metric (only `type` and one other matching field should be set at once).
        :param pulumi.Input[str] type: type is the type of metric source.  It should be one of "Object", "Pods" or "Resource", each mapping to a matching field in the object.
        :param pulumi.Input['ExternalMetricSourceArgs'] external: external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).
        :param pulumi.Input['ObjectMetricSourceArgs'] object: object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).
        :param pulumi.Input['PodsMetricSourceArgs'] pods: pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.
        :param pulumi.Input['ResourceMetricSourceArgs'] resource: resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
        """
        __self__.type = type
        __self__.external = external
        __self__.object = object
        __self__.pods = pods
        __self__.resource = resource

@pulumi.input_type
class MetricStatusArgs:
    type: pulumi.Input[str] = pulumi.input_property("type")
    """
    type is the type of metric source.  It will be one of "Object", "Pods" or "Resource", each corresponds to a matching field in the object.
    """
    external: Optional[pulumi.Input['ExternalMetricStatusArgs']] = pulumi.input_property("external")
    """
    external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).
    """
    object: Optional[pulumi.Input['ObjectMetricStatusArgs']] = pulumi.input_property("object")
    """
    object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).
    """
    pods: Optional[pulumi.Input['PodsMetricStatusArgs']] = pulumi.input_property("pods")
    """
    pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.
    """
    resource: Optional[pulumi.Input['ResourceMetricStatusArgs']] = pulumi.input_property("resource")
    """
    resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, type: pulumi.Input[str], external: Optional[pulumi.Input['ExternalMetricStatusArgs']] = None, object: Optional[pulumi.Input['ObjectMetricStatusArgs']] = None, pods: Optional[pulumi.Input['PodsMetricStatusArgs']] = None, resource: Optional[pulumi.Input['ResourceMetricStatusArgs']] = None) -> None:
        """
        MetricStatus describes the last-read state of a single metric.
        :param pulumi.Input[str] type: type is the type of metric source.  It will be one of "Object", "Pods" or "Resource", each corresponds to a matching field in the object.
        :param pulumi.Input['ExternalMetricStatusArgs'] external: external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).
        :param pulumi.Input['ObjectMetricStatusArgs'] object: object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).
        :param pulumi.Input['PodsMetricStatusArgs'] pods: pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.
        :param pulumi.Input['ResourceMetricStatusArgs'] resource: resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
        """
        __self__.type = type
        __self__.external = external
        __self__.object = object
        __self__.pods = pods
        __self__.resource = resource

@pulumi.input_type
class ObjectMetricSourceArgs:
    metric_name: pulumi.Input[str] = pulumi.input_property("metricName")
    """
    metricName is the name of the metric in question.
    """
    target: pulumi.Input['CrossVersionObjectReferenceArgs'] = pulumi.input_property("target")
    """
    target is the described Kubernetes object.
    """
    target_value: pulumi.Input[str] = pulumi.input_property("targetValue")
    """
    targetValue is the target value of the metric (as a quantity).
    """
    average_value: Optional[pulumi.Input[str]] = pulumi.input_property("averageValue")
    """
    averageValue is the target value of the average of the metric across all relevant pods (as a quantity)
    """
    selector: Optional[pulumi.Input['LabelSelectorArgs']] = pulumi.input_property("selector")
    """
    selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping When unset, just the metricName will be used to gather metrics.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, metric_name: pulumi.Input[str], target: pulumi.Input['CrossVersionObjectReferenceArgs'], target_value: pulumi.Input[str], average_value: Optional[pulumi.Input[str]] = None, selector: Optional[pulumi.Input['LabelSelectorArgs']] = None) -> None:
        """
        ObjectMetricSource indicates how to scale on a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
        :param pulumi.Input[str] metric_name: metricName is the name of the metric in question.
        :param pulumi.Input['CrossVersionObjectReferenceArgs'] target: target is the described Kubernetes object.
        :param pulumi.Input[str] target_value: targetValue is the target value of the metric (as a quantity).
        :param pulumi.Input[str] average_value: averageValue is the target value of the average of the metric across all relevant pods (as a quantity)
        :param pulumi.Input['LabelSelectorArgs'] selector: selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping When unset, just the metricName will be used to gather metrics.
        """
        __self__.metric_name = metric_name
        __self__.target = target
        __self__.target_value = target_value
        __self__.average_value = average_value
        __self__.selector = selector

@pulumi.input_type
class ObjectMetricStatusArgs:
    current_value: pulumi.Input[str] = pulumi.input_property("currentValue")
    """
    currentValue is the current value of the metric (as a quantity).
    """
    metric_name: pulumi.Input[str] = pulumi.input_property("metricName")
    """
    metricName is the name of the metric in question.
    """
    target: pulumi.Input['CrossVersionObjectReferenceArgs'] = pulumi.input_property("target")
    """
    target is the described Kubernetes object.
    """
    average_value: Optional[pulumi.Input[str]] = pulumi.input_property("averageValue")
    """
    averageValue is the current value of the average of the metric across all relevant pods (as a quantity)
    """
    selector: Optional[pulumi.Input['LabelSelectorArgs']] = pulumi.input_property("selector")
    """
    selector is the string-encoded form of a standard kubernetes label selector for the given metric When set in the ObjectMetricSource, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, current_value: pulumi.Input[str], metric_name: pulumi.Input[str], target: pulumi.Input['CrossVersionObjectReferenceArgs'], average_value: Optional[pulumi.Input[str]] = None, selector: Optional[pulumi.Input['LabelSelectorArgs']] = None) -> None:
        """
        ObjectMetricStatus indicates the current value of a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
        :param pulumi.Input[str] current_value: currentValue is the current value of the metric (as a quantity).
        :param pulumi.Input[str] metric_name: metricName is the name of the metric in question.
        :param pulumi.Input['CrossVersionObjectReferenceArgs'] target: target is the described Kubernetes object.
        :param pulumi.Input[str] average_value: averageValue is the current value of the average of the metric across all relevant pods (as a quantity)
        :param pulumi.Input['LabelSelectorArgs'] selector: selector is the string-encoded form of a standard kubernetes label selector for the given metric When set in the ObjectMetricSource, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.
        """
        __self__.current_value = current_value
        __self__.metric_name = metric_name
        __self__.target = target
        __self__.average_value = average_value
        __self__.selector = selector

@pulumi.input_type
class PodsMetricSourceArgs:
    metric_name: pulumi.Input[str] = pulumi.input_property("metricName")
    """
    metricName is the name of the metric in question
    """
    target_average_value: pulumi.Input[str] = pulumi.input_property("targetAverageValue")
    """
    targetAverageValue is the target value of the average of the metric across all relevant pods (as a quantity)
    """
    selector: Optional[pulumi.Input['LabelSelectorArgs']] = pulumi.input_property("selector")
    """
    selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping When unset, just the metricName will be used to gather metrics.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, metric_name: pulumi.Input[str], target_average_value: pulumi.Input[str], selector: Optional[pulumi.Input['LabelSelectorArgs']] = None) -> None:
        """
        PodsMetricSource indicates how to scale on a metric describing each pod in the current scale target (for example, transactions-processed-per-second). The values will be averaged together before being compared to the target value.
        :param pulumi.Input[str] metric_name: metricName is the name of the metric in question
        :param pulumi.Input[str] target_average_value: targetAverageValue is the target value of the average of the metric across all relevant pods (as a quantity)
        :param pulumi.Input['LabelSelectorArgs'] selector: selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping When unset, just the metricName will be used to gather metrics.
        """
        __self__.metric_name = metric_name
        __self__.target_average_value = target_average_value
        __self__.selector = selector

@pulumi.input_type
class PodsMetricStatusArgs:
    current_average_value: pulumi.Input[str] = pulumi.input_property("currentAverageValue")
    """
    currentAverageValue is the current value of the average of the metric across all relevant pods (as a quantity)
    """
    metric_name: pulumi.Input[str] = pulumi.input_property("metricName")
    """
    metricName is the name of the metric in question
    """
    selector: Optional[pulumi.Input['LabelSelectorArgs']] = pulumi.input_property("selector")
    """
    selector is the string-encoded form of a standard kubernetes label selector for the given metric When set in the PodsMetricSource, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, current_average_value: pulumi.Input[str], metric_name: pulumi.Input[str], selector: Optional[pulumi.Input['LabelSelectorArgs']] = None) -> None:
        """
        PodsMetricStatus indicates the current value of a metric describing each pod in the current scale target (for example, transactions-processed-per-second).
        :param pulumi.Input[str] current_average_value: currentAverageValue is the current value of the average of the metric across all relevant pods (as a quantity)
        :param pulumi.Input[str] metric_name: metricName is the name of the metric in question
        :param pulumi.Input['LabelSelectorArgs'] selector: selector is the string-encoded form of a standard kubernetes label selector for the given metric When set in the PodsMetricSource, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.
        """
        __self__.current_average_value = current_average_value
        __self__.metric_name = metric_name
        __self__.selector = selector

@pulumi.input_type
class ResourceMetricSourceArgs:
    name: pulumi.Input[str] = pulumi.input_property("name")
    """
    name is the name of the resource in question.
    """
    target_average_utilization: Optional[pulumi.Input[float]] = pulumi.input_property("targetAverageUtilization")
    """
    targetAverageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.
    """
    target_average_value: Optional[pulumi.Input[str]] = pulumi.input_property("targetAverageValue")
    """
    targetAverageValue is the target value of the average of the resource metric across all relevant pods, as a raw value (instead of as a percentage of the request), similar to the "pods" metric source type.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, name: pulumi.Input[str], target_average_utilization: Optional[pulumi.Input[float]] = None, target_average_value: Optional[pulumi.Input[str]] = None) -> None:
        """
        ResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  The values will be averaged together before being compared to the target.  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.  Only one "target" type should be set.
        :param pulumi.Input[str] name: name is the name of the resource in question.
        :param pulumi.Input[float] target_average_utilization: targetAverageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.
        :param pulumi.Input[str] target_average_value: targetAverageValue is the target value of the average of the resource metric across all relevant pods, as a raw value (instead of as a percentage of the request), similar to the "pods" metric source type.
        """
        __self__.name = name
        __self__.target_average_utilization = target_average_utilization
        __self__.target_average_value = target_average_value

@pulumi.input_type
class ResourceMetricStatusArgs:
    current_average_value: pulumi.Input[str] = pulumi.input_property("currentAverageValue")
    """
    currentAverageValue is the current value of the average of the resource metric across all relevant pods, as a raw value (instead of as a percentage of the request), similar to the "pods" metric source type. It will always be set, regardless of the corresponding metric specification.
    """
    name: pulumi.Input[str] = pulumi.input_property("name")
    """
    name is the name of the resource in question.
    """
    current_average_utilization: Optional[pulumi.Input[float]] = pulumi.input_property("currentAverageUtilization")
    """
    currentAverageUtilization is the current value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.  It will only be present if `targetAverageValue` was set in the corresponding metric specification.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, current_average_value: pulumi.Input[str], name: pulumi.Input[str], current_average_utilization: Optional[pulumi.Input[float]] = None) -> None:
        """
        ResourceMetricStatus indicates the current value of a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
        :param pulumi.Input[str] current_average_value: currentAverageValue is the current value of the average of the resource metric across all relevant pods, as a raw value (instead of as a percentage of the request), similar to the "pods" metric source type. It will always be set, regardless of the corresponding metric specification.
        :param pulumi.Input[str] name: name is the name of the resource in question.
        :param pulumi.Input[float] current_average_utilization: currentAverageUtilization is the current value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.  It will only be present if `targetAverageValue` was set in the corresponding metric specification.
        """
        __self__.current_average_value = current_average_value
        __self__.name = name
        __self__.current_average_utilization = current_average_utilization

