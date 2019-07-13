# *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
from typing import Optional

import pulumi
import pulumi.runtime
from pulumi import Input, ResourceOptions

from ... import tables, version


class StorageClass(pulumi.CustomResource):
    """
    StorageClass describes the parameters for a class of storage for which PersistentVolumes can be
    dynamically provisioned.
    
    StorageClasses are non-namespaced; the name of the storage class according to etcd is in
    ObjectMeta.Name.
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

    allow_volume_expansion: pulumi.Output[bool];
    """
    AllowVolumeExpansion shows whether the storage class allow volume expand
    """

    allowed_topologies: pulumi.Output[list];
    """
    Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin
    defines its own supported topology specifications. An empty TopologySelectorTerm list means
    there is no topology restriction. This field is only honored by servers that enable the
    VolumeScheduling feature.
    """

    metadata: pulumi.Output[dict];
    """
    Standard object's metadata. More info:
    https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
    """

    mount_options: pulumi.Output[list];
    """
    Dynamically provisioned PersistentVolumes of this storage class are created with these
    mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is
    invalid.
    """

    parameters: pulumi.Output[dict];
    """
    Parameters holds the parameters for the provisioner that should create volumes of this storage
    class.
    """

    provisioner: pulumi.Output[str];
    """
    Provisioner indicates the type of the provisioner.
    """

    reclaim_policy: pulumi.Output[str];
    """
    Dynamically provisioned PersistentVolumes of this storage class are created with this
    reclaimPolicy. Defaults to Delete.
    """

    volume_binding_mode: pulumi.Output[str];
    """
    VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When
    unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the
    VolumeScheduling feature.
    """

    def __init__(self, resource_name, opts=None, provisioner=None, allow_volume_expansion=None, allowed_topologies=None, metadata=None, mount_options=None, parameters=None, reclaim_policy=None, volume_binding_mode=None, __name__=None, __opts__=None):
        """
        Create a StorageClass resource with the given unique name, arguments, and options.

        :param str resource_name: The _unique_ name of the resource.
        :param pulumi.ResourceOptions opts: A bag of options that control this resource's behavior.
        :param pulumi.Input[str] provisioner: Provisioner indicates the type of the provisioner.
        :param pulumi.Input[bool] allow_volume_expansion: AllowVolumeExpansion shows whether the storage class allow volume expand
        :param pulumi.Input[list] allowed_topologies: Restrict the node topologies where volumes can be dynamically provisioned. Each
               volume plugin defines its own supported topology specifications. An empty
               TopologySelectorTerm list means there is no topology restriction. This field is only
               honored by servers that enable the VolumeScheduling feature.
        :param pulumi.Input[dict] metadata: Standard object's metadata. More info:
               https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
        :param pulumi.Input[list] mount_options: Dynamically provisioned PersistentVolumes of this storage class are created with
               these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply
               fail if one is invalid.
        :param pulumi.Input[dict] parameters: Parameters holds the parameters for the provisioner that should create volumes of
               this storage class.
        :param pulumi.Input[str] reclaim_policy: Dynamically provisioned PersistentVolumes of this storage class are created with this
               reclaimPolicy. Defaults to Delete.
        :param pulumi.Input[str] volume_binding_mode: VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and
               bound.  When unset, VolumeBindingImmediate is used. This field is only honored by
               servers that enable the VolumeScheduling feature.
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

        __props__['apiVersion'] = 'storage.k8s.io/v1beta1'
        __props__['kind'] = 'StorageClass'
        if provisioner is None:
            raise TypeError('Missing required property provisioner')
        __props__['provisioner'] = provisioner
        __props__['allowVolumeExpansion'] = allow_volume_expansion
        __props__['allowedTopologies'] = allowed_topologies
        __props__['metadata'] = metadata
        __props__['mountOptions'] = mount_options
        __props__['parameters'] = parameters
        __props__['reclaimPolicy'] = reclaim_policy
        __props__['volumeBindingMode'] = volume_binding_mode

        __props__['status'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = version.get_version()

        super(StorageClass, self).__init__(
            "kubernetes:storage.k8s.io/v1beta1:StorageClass",
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(name: str, id: Input[str], opts: Optional[ResourceOptions] = None):
        opts = ResourceOptions(id=id) if opts is None else opts.merge(ResourceOptions(id=id))
        return StorageClass(name, opts)

    def translate_output_property(self, prop: str) -> str:
        return tables._CASING_FORWARD_TABLE.get(prop) or prop

    def translate_input_property(self, prop: str) -> str:
        return tables._CASING_BACKWARD_TABLE.get(prop) or prop
