# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ... import core as _core
from ... import meta as _meta

__all__ = ['StorageClass']


class StorageClass(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_volume_expansion: Optional[pulumi.Input[bool]] = None,
                 allowed_topologies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_core.v1.TopologySelectorTermArgs']]]]] = None,
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[pulumi.InputType['_meta.v1.ObjectMetaArgs']]] = None,
                 mount_options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 provisioner: Optional[pulumi.Input[str]] = None,
                 reclaim_policy: Optional[pulumi.Input[str]] = None,
                 volume_binding_mode: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.

        StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_volume_expansion: AllowVolumeExpansion shows whether the storage class allow volume expand
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_core.v1.TopologySelectorTermArgs']]]] allowed_topologies: Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input[pulumi.InputType['_meta.v1.ObjectMetaArgs']] metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input[Sequence[pulumi.Input[str]]] mount_options: Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: Parameters holds the parameters for the provisioner that should create volumes of this storage class.
        :param pulumi.Input[str] provisioner: Provisioner indicates the type of the provisioner.
        :param pulumi.Input[str] reclaim_policy: Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
        :param pulumi.Input[str] volume_binding_mode: VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
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

            __props__['allow_volume_expansion'] = allow_volume_expansion
            __props__['allowed_topologies'] = allowed_topologies
            __props__['api_version'] = 'storage.k8s.io/v1'
            __props__['kind'] = 'StorageClass'
            __props__['metadata'] = metadata
            __props__['mount_options'] = mount_options
            __props__['parameters'] = parameters
            if provisioner is None:
                raise TypeError("Missing required property 'provisioner'")
            __props__['provisioner'] = provisioner
            __props__['reclaim_policy'] = reclaim_policy
            __props__['volume_binding_mode'] = volume_binding_mode
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="kubernetes:storage.k8s.io/v1beta1:StorageClass")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(StorageClass, __self__).__init__(
            'kubernetes:storage.k8s.io/v1:StorageClass',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'StorageClass':
        """
        Get an existing StorageClass resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return StorageClass(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowVolumeExpansion")
    def allow_volume_expansion(self) -> pulumi.Output[Optional[bool]]:
        """
        AllowVolumeExpansion shows whether the storage class allow volume expand
        """
        return pulumi.get(self, "allow_volume_expansion")

    @property
    @pulumi.getter(name="allowedTopologies")
    def allowed_topologies(self) -> pulumi.Output[Optional[Sequence['_core.v1.outputs.TopologySelectorTerm']]]:
        """
        Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
        """
        return pulumi.get(self, "allowed_topologies")

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> pulumi.Output[Optional[str]]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional[str]]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional['_meta.v1.outputs.ObjectMeta']]:
        """
        Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="mountOptions")
    def mount_options(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
        """
        return pulumi.get(self, "mount_options")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Parameters holds the parameters for the provisioner that should create volumes of this storage class.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def provisioner(self) -> pulumi.Output[str]:
        """
        Provisioner indicates the type of the provisioner.
        """
        return pulumi.get(self, "provisioner")

    @property
    @pulumi.getter(name="reclaimPolicy")
    def reclaim_policy(self) -> pulumi.Output[Optional[str]]:
        """
        Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
        """
        return pulumi.get(self, "reclaim_policy")

    @property
    @pulumi.getter(name="volumeBindingMode")
    def volume_binding_mode(self) -> pulumi.Output[Optional[str]]:
        """
        VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
        """
        return pulumi.get(self, "volume_binding_mode")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

