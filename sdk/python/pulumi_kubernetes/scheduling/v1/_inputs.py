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
class PriorityClassArgs:
    value: pulumi.Input[float] = pulumi.input_property("value")
    """
    The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
    """
    api_version: Optional[pulumi.Input[str]] = pulumi.input_property("apiVersion")
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    """
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")
    """
    description is an arbitrary string that usually provides guidelines on when this priority class should be used.
    """
    global_default: Optional[pulumi.Input[bool]] = pulumi.input_property("globalDefault")
    """
    globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
    """
    kind: Optional[pulumi.Input[str]] = pulumi.input_property("kind")
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    """
    metadata: Optional[pulumi.Input['ObjectMetaArgs']] = pulumi.input_property("metadata")
    """
    Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
    """
    preemption_policy: Optional[pulumi.Input[str]] = pulumi.input_property("preemptionPolicy")
    """
    PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is alpha-level and is only honored by servers that enable the NonPreemptingPriority feature.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, value: pulumi.Input[float], api_version: Optional[pulumi.Input[str]] = None, description: Optional[pulumi.Input[str]] = None, global_default: Optional[pulumi.Input[bool]] = None, kind: Optional[pulumi.Input[str]] = None, metadata: Optional[pulumi.Input['ObjectMetaArgs']] = None, preemption_policy: Optional[pulumi.Input[str]] = None) -> None:
        """
        PriorityClass defines mapping from a priority class name to the priority integer value. The value can be any valid integer.
        :param pulumi.Input[float] value: The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] description: description is an arbitrary string that usually provides guidelines on when this priority class should be used.
        :param pulumi.Input[bool] global_default: globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['ObjectMetaArgs'] metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input[str] preemption_policy: PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is alpha-level and is only honored by servers that enable the NonPreemptingPriority feature.
        """
        __self__.value = value
        __self__.api_version = 'scheduling.k8s.io/v1'
        __self__.description = description
        __self__.global_default = global_default
        __self__.kind = 'PriorityClass'
        __self__.metadata = metadata
        __self__.preemption_policy = preemption_policy

