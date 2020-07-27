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
from ._inputs import *

__all__ = ['Pod']


class Pod(pulumi.CustomResource):
    api_version: pulumi.Output[Optional[str]] = pulumi.output_property("apiVersion")
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    """
    kind: pulumi.Output[Optional[str]] = pulumi.output_property("kind")
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    """
    metadata: pulumi.Output[Optional['_meta.v1.outputs.ObjectMeta']] = pulumi.output_property("metadata")
    """
    Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
    """
    spec: pulumi.Output[Optional['outputs.PodSpec']] = pulumi.output_property("spec")
    """
    Specification of the desired behavior of the pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
    """
    status: pulumi.Output[Optional['outputs.PodStatus']] = pulumi.output_property("status")
    """
    Most recently observed status of the pod. This data may not be up to date. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, api_version: Optional[pulumi.Input[str]] = None, kind: Optional[pulumi.Input[str]] = None, metadata: Optional[pulumi.Input[pulumi.InputType['_meta.v1.ObjectMetaArgs']]] = None, spec: Optional[pulumi.Input[pulumi.InputType['PodSpecArgs']]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Pod is a collection of containers that can run on a host. This resource is created by clients and scheduled onto hosts.

        This resource waits until its status is ready before registering success
        for create/update, and populating output properties from the current state of the resource.
        The following conditions are used to determine whether the resource creation has
        succeeded or failed:

        1. The Pod is scheduled ("PodScheduled"" '.status.condition' is true).
        2. The Pod is initialized ("Initialized" '.status.condition' is true).
        3. The Pod is ready ("Ready" '.status.condition' is true) and the '.status.phase' is
           set to "Running".
           Or (for Jobs): The Pod succeeded ('.status.phase' set to "Succeeded").

        If the Pod has not reached a Ready state after 10 minutes, it will
        time out and mark the resource update as Failed. You can override the default timeout value
        by setting the 'customTimeouts' option on the resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input[pulumi.InputType['_meta.v1.ObjectMetaArgs']] metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input[pulumi.InputType['PodSpecArgs']] spec: Specification of the desired behavior of the pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['api_version'] = 'v1'
            __props__['kind'] = 'Pod'
            __props__['metadata'] = metadata
            __props__['spec'] = spec
            __props__['status'] = None
        super(Pod, __self__).__init__(
            'kubernetes:core/v1:Pod',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None) -> 'Pod':
        """
        Get an existing Pod resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Pod(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

