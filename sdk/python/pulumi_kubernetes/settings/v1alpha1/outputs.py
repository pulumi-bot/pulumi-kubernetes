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

@pulumi.output_type
class PodPreset(dict):
    """
    PodPreset is a policy resource that defines additional runtime requirements for a Pod.
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
    spec: Optional['outputs.PodPresetSpec'] = pulumi.output_property("spec")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PodPresetSpec(dict):
    """
    PodPresetSpec is a description of a pod preset.
    """
    env: Optional[List['_core.v1.outputs.EnvVar']] = pulumi.output_property("env")
    """
    Env defines the collection of EnvVar to inject into containers.
    """
    env_from: Optional[List['_core.v1.outputs.EnvFromSource']] = pulumi.output_property("envFrom")
    """
    EnvFrom defines the collection of EnvFromSource to inject into containers.
    """
    selector: Optional['_meta.v1.outputs.LabelSelector'] = pulumi.output_property("selector")
    """
    Selector is a label query over a set of resources, in this case pods. Required.
    """
    volume_mounts: Optional[List['_core.v1.outputs.VolumeMount']] = pulumi.output_property("volumeMounts")
    """
    VolumeMounts defines the collection of VolumeMount to inject into containers.
    """
    volumes: Optional[List['_core.v1.outputs.Volume']] = pulumi.output_property("volumes")
    """
    Volumes defines the collection of Volume to inject into the pod.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


