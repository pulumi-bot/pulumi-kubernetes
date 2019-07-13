# *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
from typing import Optional

import pulumi
import pulumi.runtime
from pulumi import Input, ResourceOptions

from ... import tables, version


class PriorityClass(pulumi.CustomResource):
    """
    DEPRECATED - This group version of PriorityClass is deprecated by
    scheduling.k8s.io/v1/PriorityClass. PriorityClass defines mapping from a priority class name to
    the priority integer value. The value can be any valid integer.
    """

    apiVersion: pulumi.Output[str]
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should
    convert recognized schemas to the latest internal value, and may reject unrecognized values.
    More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
    """

    kind: pulumi.Output[str]
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer
    this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
    info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
    """

    description: pulumi.Output[str];
    """
    description is an arbitrary string that usually provides guidelines on when this priority class
    should be used.
    """

    global_default: pulumi.Output[bool];
    """
    globalDefault specifies whether this PriorityClass should be considered as the default priority
    for pods that do not have any priority class. Only one PriorityClass can be marked as
    `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault`
    field set to true, the smallest value of such global default PriorityClasses will be used as the
    default priority.
    """

    metadata: pulumi.Output[dict];
    """
    Standard object's metadata. More info:
    https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
    """

    preemption_policy: pulumi.Output[str];
    """
    PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never,
    PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is alpha-level and
    is only honored by servers that enable the NonPreemptingPriority feature.
    """

    value: pulumi.Output[int];
    """
    The value of this priority class. This is the actual priority that pods receive when they have
    the name of this class in their pod spec.
    """

    def __init__(self, resource_name, opts=None, value=None, description=None, global_default=None, metadata=None, preemption_policy=None, __name__=None, __opts__=None):
        """
        Create a PriorityClass resource with the given unique name, arguments, and options.

        :param str resource_name: The _unique_ name of the resource.
        :param pulumi.ResourceOptions opts: A bag of options that control this resource's behavior.
        :param pulumi.Input[int] value: The value of this priority class. This is the actual priority that pods receive when
               they have the name of this class in their pod spec.
        :param pulumi.Input[str] description: description is an arbitrary string that usually provides guidelines on when this
               priority class should be used.
        :param pulumi.Input[bool] global_default: globalDefault specifies whether this PriorityClass should be considered as the
               default priority for pods that do not have any priority class. Only one PriorityClass
               can be marked as `globalDefault`. However, if more than one PriorityClasses exists
               with their `globalDefault` field set to true, the smallest value of such global
               default PriorityClasses will be used as the default priority.
        :param pulumi.Input[dict] metadata: Standard object's metadata. More info:
               https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
        :param pulumi.Input[str] preemption_policy: PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never,
               PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is
               alpha-level and is only honored by servers that enable the NonPreemptingPriority
               feature.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['apiVersion'] = 'scheduling.k8s.io/v1alpha1'
        __props__['kind'] = 'PriorityClass'
        if value is None:
            raise TypeError('Missing required property value')
        __props__['value'] = value
        __props__['description'] = description
        __props__['globalDefault'] = global_default
        __props__['metadata'] = metadata
        __props__['preemptionPolicy'] = preemption_policy

        __props__['status'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = version.get_version()

        super(PriorityClass, self).__init__(
            "kubernetes:scheduling.k8s.io/v1alpha1:PriorityClass",
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(name: str, id: Input[str], opts: Optional[ResourceOptions] = None):
        opts = ResourceOptions(id=id) if opts is None else opts.merge(ResourceOptions(id=id))
        return PriorityClass(name, opts)

    def translate_output_property(self, prop: str) -> str:
        return tables._CASING_FORWARD_TABLE.get(prop) or prop

    def translate_input_property(self, prop: str) -> str:
        return tables._CASING_BACKWARD_TABLE.get(prop) or prop
