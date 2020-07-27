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

__all__ = [
    'CSIDriverArgs',
    'CSIDriverSpecArgs',
    'CSINodeArgs',
    'CSINodeDriverArgs',
    'CSINodeSpecArgs',
    'StorageClassArgs',
    'VolumeAttachmentArgs',
    'VolumeAttachmentSourceArgs',
    'VolumeAttachmentSpecArgs',
    'VolumeAttachmentStatusArgs',
    'VolumeErrorArgs',
    'VolumeNodeResourcesArgs',
]

@pulumi.input_type
class CSIDriverArgs:
    spec: pulumi.Input['CSIDriverSpecArgs'] = pulumi.input_property("spec")
    """
    Specification of the CSI Driver.
    """
    api_version: Optional[pulumi.Input[str]] = pulumi.input_property("apiVersion")
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    """
    kind: Optional[pulumi.Input[str]] = pulumi.input_property("kind")
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    """
    metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = pulumi.input_property("metadata")
    """
    Standard object metadata. metadata.Name indicates the name of the CSI driver that this object refers to; it MUST be the same name returned by the CSI GetPluginName() call for that driver. The driver name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), dots (.), and alphanumerics between. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, spec: pulumi.Input['CSIDriverSpecArgs'], api_version: Optional[pulumi.Input[str]] = None, kind: Optional[pulumi.Input[str]] = None, metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None) -> None:
        """
        CSIDriver captures information about a Container Storage Interface (CSI) volume driver deployed on the cluster. CSI drivers do not need to create the CSIDriver object directly. Instead they may use the cluster-driver-registrar sidecar container. When deployed with a CSI driver it automatically creates a CSIDriver object representing the driver. Kubernetes attach detach controller uses this object to determine whether attach is required. Kubelet uses this object to determine whether pod information needs to be passed on mount. CSIDriver objects are non-namespaced.
        :param pulumi.Input['CSIDriverSpecArgs'] spec: Specification of the CSI Driver.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object metadata. metadata.Name indicates the name of the CSI driver that this object refers to; it MUST be the same name returned by the CSI GetPluginName() call for that driver. The driver name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), dots (.), and alphanumerics between. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        __self__.spec = spec
        __self__.api_version = 'storage.k8s.io/v1beta1'
        __self__.kind = 'CSIDriver'
        __self__.metadata = metadata

@pulumi.input_type
class CSIDriverSpecArgs:
    attach_required: Optional[pulumi.Input[bool]] = pulumi.input_property("attachRequired")
    """
    attachRequired indicates this CSI volume driver requires an attach operation (because it implements the CSI ControllerPublishVolume() method), and that the Kubernetes attach detach controller should call the attach volume interface which checks the volumeattachment status and waits until the volume is attached before proceeding to mounting. The CSI external-attacher coordinates with CSI volume driver and updates the volumeattachment status when the attach operation is complete. If the CSIDriverRegistry feature gate is enabled and the value is specified to false, the attach operation will be skipped. Otherwise the attach operation will be called.
    """
    pod_info_on_mount: Optional[pulumi.Input[bool]] = pulumi.input_property("podInfoOnMount")
    """
    If set to true, podInfoOnMount indicates this CSI volume driver requires additional pod information (like podName, podUID, etc.) during mount operations. If set to false, pod information will not be passed on mount. Default is false. The CSI driver specifies podInfoOnMount as part of driver deployment. If true, Kubelet will pass pod information as VolumeContext in the CSI NodePublishVolume() calls. The CSI driver is responsible for parsing and validating the information passed in as VolumeContext. The following VolumeConext will be passed if podInfoOnMount is set to true. This list might grow, but the prefix will be used. "csi.storage.k8s.io/pod.name": pod.Name "csi.storage.k8s.io/pod.namespace": pod.Namespace "csi.storage.k8s.io/pod.uid": string(pod.UID) "csi.storage.k8s.io/ephemeral": "true" iff the volume is an ephemeral inline volume
                                    defined by a CSIVolumeSource, otherwise "false"

    "csi.storage.k8s.io/ephemeral" is a new feature in Kubernetes 1.16. It is only required for drivers which support both the "Persistent" and "Ephemeral" VolumeLifecycleMode. Other drivers can leave pod info disabled and/or ignore this field. As Kubernetes 1.15 doesn't support this field, drivers can only support one mode when deployed on such a cluster and the deployment determines which mode that is, for example via a command line parameter of the driver.
    """
    volume_lifecycle_modes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("volumeLifecycleModes")
    """
    VolumeLifecycleModes defines what kind of volumes this CSI volume driver supports. The default if the list is empty is "Persistent", which is the usage defined by the CSI specification and implemented in Kubernetes via the usual PV/PVC mechanism. The other mode is "Ephemeral". In this mode, volumes are defined inline inside the pod spec with CSIVolumeSource and their lifecycle is tied to the lifecycle of that pod. A driver has to be aware of this because it is only going to get a NodePublishVolume call for such a volume. For more information about implementing this mode, see https://kubernetes-csi.github.io/docs/ephemeral-local-volumes.html A driver can support one or more of these modes and more modes may be added in the future.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, attach_required: Optional[pulumi.Input[bool]] = None, pod_info_on_mount: Optional[pulumi.Input[bool]] = None, volume_lifecycle_modes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> None:
        """
        CSIDriverSpec is the specification of a CSIDriver.
        :param pulumi.Input[bool] attach_required: attachRequired indicates this CSI volume driver requires an attach operation (because it implements the CSI ControllerPublishVolume() method), and that the Kubernetes attach detach controller should call the attach volume interface which checks the volumeattachment status and waits until the volume is attached before proceeding to mounting. The CSI external-attacher coordinates with CSI volume driver and updates the volumeattachment status when the attach operation is complete. If the CSIDriverRegistry feature gate is enabled and the value is specified to false, the attach operation will be skipped. Otherwise the attach operation will be called.
        :param pulumi.Input[bool] pod_info_on_mount: If set to true, podInfoOnMount indicates this CSI volume driver requires additional pod information (like podName, podUID, etc.) during mount operations. If set to false, pod information will not be passed on mount. Default is false. The CSI driver specifies podInfoOnMount as part of driver deployment. If true, Kubelet will pass pod information as VolumeContext in the CSI NodePublishVolume() calls. The CSI driver is responsible for parsing and validating the information passed in as VolumeContext. The following VolumeConext will be passed if podInfoOnMount is set to true. This list might grow, but the prefix will be used. "csi.storage.k8s.io/pod.name": pod.Name "csi.storage.k8s.io/pod.namespace": pod.Namespace "csi.storage.k8s.io/pod.uid": string(pod.UID) "csi.storage.k8s.io/ephemeral": "true" iff the volume is an ephemeral inline volume
                                               defined by a CSIVolumeSource, otherwise "false"
               
               "csi.storage.k8s.io/ephemeral" is a new feature in Kubernetes 1.16. It is only required for drivers which support both the "Persistent" and "Ephemeral" VolumeLifecycleMode. Other drivers can leave pod info disabled and/or ignore this field. As Kubernetes 1.15 doesn't support this field, drivers can only support one mode when deployed on such a cluster and the deployment determines which mode that is, for example via a command line parameter of the driver.
        :param pulumi.Input[List[pulumi.Input[str]]] volume_lifecycle_modes: VolumeLifecycleModes defines what kind of volumes this CSI volume driver supports. The default if the list is empty is "Persistent", which is the usage defined by the CSI specification and implemented in Kubernetes via the usual PV/PVC mechanism. The other mode is "Ephemeral". In this mode, volumes are defined inline inside the pod spec with CSIVolumeSource and their lifecycle is tied to the lifecycle of that pod. A driver has to be aware of this because it is only going to get a NodePublishVolume call for such a volume. For more information about implementing this mode, see https://kubernetes-csi.github.io/docs/ephemeral-local-volumes.html A driver can support one or more of these modes and more modes may be added in the future.
        """
        __self__.attach_required = attach_required
        __self__.pod_info_on_mount = pod_info_on_mount
        __self__.volume_lifecycle_modes = volume_lifecycle_modes

@pulumi.input_type
class CSINodeArgs:
    spec: pulumi.Input['CSINodeSpecArgs'] = pulumi.input_property("spec")
    """
    spec is the specification of CSINode
    """
    api_version: Optional[pulumi.Input[str]] = pulumi.input_property("apiVersion")
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    """
    kind: Optional[pulumi.Input[str]] = pulumi.input_property("kind")
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    """
    metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = pulumi.input_property("metadata")
    """
    metadata.name must be the Kubernetes node name.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, spec: pulumi.Input['CSINodeSpecArgs'], api_version: Optional[pulumi.Input[str]] = None, kind: Optional[pulumi.Input[str]] = None, metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None) -> None:
        """
        CSINode holds information about all CSI drivers installed on a node. CSI drivers do not need to create the CSINode object directly. As long as they use the node-driver-registrar sidecar container, the kubelet will automatically populate the CSINode object for the CSI driver as part of kubelet plugin registration. CSINode has the same name as a node. If the object is missing, it means either there are no CSI Drivers available on the node, or the Kubelet version is low enough that it doesn't create this object. CSINode has an OwnerReference that points to the corresponding node object.
        :param pulumi.Input['CSINodeSpecArgs'] spec: spec is the specification of CSINode
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: metadata.name must be the Kubernetes node name.
        """
        __self__.spec = spec
        __self__.api_version = 'storage.k8s.io/v1beta1'
        __self__.kind = 'CSINode'
        __self__.metadata = metadata

@pulumi.input_type
class CSINodeDriverArgs:
    name: pulumi.Input[str] = pulumi.input_property("name")
    """
    This is the name of the CSI driver that this object refers to. This MUST be the same name returned by the CSI GetPluginName() call for that driver.
    """
    node_id: pulumi.Input[str] = pulumi.input_property("nodeID")
    """
    nodeID of the node from the driver point of view. This field enables Kubernetes to communicate with storage systems that do not share the same nomenclature for nodes. For example, Kubernetes may refer to a given node as "node1", but the storage system may refer to the same node as "nodeA". When Kubernetes issues a command to the storage system to attach a volume to a specific node, it can use this field to refer to the node name using the ID that the storage system will understand, e.g. "nodeA" instead of "node1". This field is required.
    """
    allocatable: Optional[pulumi.Input['VolumeNodeResourcesArgs']] = pulumi.input_property("allocatable")
    """
    allocatable represents the volume resources of a node that are available for scheduling.
    """
    topology_keys: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("topologyKeys")
    """
    topologyKeys is the list of keys supported by the driver. When a driver is initialized on a cluster, it provides a set of topology keys that it understands (e.g. "company.com/zone", "company.com/region"). When a driver is initialized on a node, it provides the same topology keys along with values. Kubelet will expose these topology keys as labels on its own node object. When Kubernetes does topology aware provisioning, it can use this list to determine which labels it should retrieve from the node object and pass back to the driver. It is possible for different nodes to use different topology keys. This can be empty if driver does not support topology.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, name: pulumi.Input[str], node_id: pulumi.Input[str], allocatable: Optional[pulumi.Input['VolumeNodeResourcesArgs']] = None, topology_keys: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> None:
        """
        CSINodeDriver holds information about the specification of one CSI driver installed on a node
        :param pulumi.Input[str] name: This is the name of the CSI driver that this object refers to. This MUST be the same name returned by the CSI GetPluginName() call for that driver.
        :param pulumi.Input[str] node_id: nodeID of the node from the driver point of view. This field enables Kubernetes to communicate with storage systems that do not share the same nomenclature for nodes. For example, Kubernetes may refer to a given node as "node1", but the storage system may refer to the same node as "nodeA". When Kubernetes issues a command to the storage system to attach a volume to a specific node, it can use this field to refer to the node name using the ID that the storage system will understand, e.g. "nodeA" instead of "node1". This field is required.
        :param pulumi.Input['VolumeNodeResourcesArgs'] allocatable: allocatable represents the volume resources of a node that are available for scheduling.
        :param pulumi.Input[List[pulumi.Input[str]]] topology_keys: topologyKeys is the list of keys supported by the driver. When a driver is initialized on a cluster, it provides a set of topology keys that it understands (e.g. "company.com/zone", "company.com/region"). When a driver is initialized on a node, it provides the same topology keys along with values. Kubelet will expose these topology keys as labels on its own node object. When Kubernetes does topology aware provisioning, it can use this list to determine which labels it should retrieve from the node object and pass back to the driver. It is possible for different nodes to use different topology keys. This can be empty if driver does not support topology.
        """
        __self__.name = name
        __self__.node_id = node_id
        __self__.allocatable = allocatable
        __self__.topology_keys = topology_keys

@pulumi.input_type
class CSINodeSpecArgs:
    drivers: pulumi.Input[List[pulumi.Input['CSINodeDriverArgs']]] = pulumi.input_property("drivers")
    """
    drivers is a list of information of all CSI Drivers existing on a node. If all drivers in the list are uninstalled, this can become empty.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, drivers: pulumi.Input[List[pulumi.Input['CSINodeDriverArgs']]]) -> None:
        """
        CSINodeSpec holds information about the specification of all CSI drivers installed on a node
        :param pulumi.Input[List[pulumi.Input['CSINodeDriverArgs']]] drivers: drivers is a list of information of all CSI Drivers existing on a node. If all drivers in the list are uninstalled, this can become empty.
        """
        __self__.drivers = drivers

@pulumi.input_type
class StorageClassArgs:
    provisioner: pulumi.Input[str] = pulumi.input_property("provisioner")
    """
    Provisioner indicates the type of the provisioner.
    """
    allow_volume_expansion: Optional[pulumi.Input[bool]] = pulumi.input_property("allowVolumeExpansion")
    """
    AllowVolumeExpansion shows whether the storage class allow volume expand
    """
    allowed_topologies: Optional[pulumi.Input[List[pulumi.Input['_core.v1.TopologySelectorTermArgs']]]] = pulumi.input_property("allowedTopologies")
    """
    Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
    """
    api_version: Optional[pulumi.Input[str]] = pulumi.input_property("apiVersion")
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    """
    kind: Optional[pulumi.Input[str]] = pulumi.input_property("kind")
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    """
    metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = pulumi.input_property("metadata")
    """
    Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
    """
    mount_options: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("mountOptions")
    """
    Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
    """
    parameters: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = pulumi.input_property("parameters")
    """
    Parameters holds the parameters for the provisioner that should create volumes of this storage class.
    """
    reclaim_policy: Optional[pulumi.Input[str]] = pulumi.input_property("reclaimPolicy")
    """
    Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
    """
    volume_binding_mode: Optional[pulumi.Input[str]] = pulumi.input_property("volumeBindingMode")
    """
    VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, provisioner: pulumi.Input[str], allow_volume_expansion: Optional[pulumi.Input[bool]] = None, allowed_topologies: Optional[pulumi.Input[List[pulumi.Input['_core.v1.TopologySelectorTermArgs']]]] = None, api_version: Optional[pulumi.Input[str]] = None, kind: Optional[pulumi.Input[str]] = None, metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None, mount_options: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, parameters: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, reclaim_policy: Optional[pulumi.Input[str]] = None, volume_binding_mode: Optional[pulumi.Input[str]] = None) -> None:
        """
        StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.

        StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.
        :param pulumi.Input[str] provisioner: Provisioner indicates the type of the provisioner.
        :param pulumi.Input[bool] allow_volume_expansion: AllowVolumeExpansion shows whether the storage class allow volume expand
        :param pulumi.Input[List[pulumi.Input['_core.v1.TopologySelectorTermArgs']]] allowed_topologies: Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input[List[pulumi.Input[str]]] mount_options: Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] parameters: Parameters holds the parameters for the provisioner that should create volumes of this storage class.
        :param pulumi.Input[str] reclaim_policy: Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.
        :param pulumi.Input[str] volume_binding_mode: VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
        """
        __self__.provisioner = provisioner
        __self__.allow_volume_expansion = allow_volume_expansion
        __self__.allowed_topologies = allowed_topologies
        __self__.api_version = 'storage.k8s.io/v1beta1'
        __self__.kind = 'StorageClass'
        __self__.metadata = metadata
        __self__.mount_options = mount_options
        __self__.parameters = parameters
        __self__.reclaim_policy = reclaim_policy
        __self__.volume_binding_mode = volume_binding_mode

@pulumi.input_type
class VolumeAttachmentArgs:
    spec: pulumi.Input['VolumeAttachmentSpecArgs'] = pulumi.input_property("spec")
    """
    Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.
    """
    api_version: Optional[pulumi.Input[str]] = pulumi.input_property("apiVersion")
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    """
    kind: Optional[pulumi.Input[str]] = pulumi.input_property("kind")
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    """
    metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = pulumi.input_property("metadata")
    """
    Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
    """
    status: Optional[pulumi.Input['VolumeAttachmentStatusArgs']] = pulumi.input_property("status")
    """
    Status of the VolumeAttachment request. Populated by the entity completing the attach or detach operation, i.e. the external-attacher.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, spec: pulumi.Input['VolumeAttachmentSpecArgs'], api_version: Optional[pulumi.Input[str]] = None, kind: Optional[pulumi.Input[str]] = None, metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None, status: Optional[pulumi.Input['VolumeAttachmentStatusArgs']] = None) -> None:
        """
        VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node.

        VolumeAttachment objects are non-namespaced.
        :param pulumi.Input['VolumeAttachmentSpecArgs'] spec: Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input['VolumeAttachmentStatusArgs'] status: Status of the VolumeAttachment request. Populated by the entity completing the attach or detach operation, i.e. the external-attacher.
        """
        __self__.spec = spec
        __self__.api_version = 'storage.k8s.io/v1beta1'
        __self__.kind = 'VolumeAttachment'
        __self__.metadata = metadata
        __self__.status = status

@pulumi.input_type
class VolumeAttachmentSourceArgs:
    inline_volume_spec: Optional[pulumi.Input['_core.v1.PersistentVolumeSpecArgs']] = pulumi.input_property("inlineVolumeSpec")
    """
    inlineVolumeSpec contains all the information necessary to attach a persistent volume defined by a pod's inline VolumeSource. This field is populated only for the CSIMigration feature. It contains translated fields from a pod's inline VolumeSource to a PersistentVolumeSpec. This field is alpha-level and is only honored by servers that enabled the CSIMigration feature.
    """
    persistent_volume_name: Optional[pulumi.Input[str]] = pulumi.input_property("persistentVolumeName")
    """
    Name of the persistent volume to attach.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, inline_volume_spec: Optional[pulumi.Input['_core.v1.PersistentVolumeSpecArgs']] = None, persistent_volume_name: Optional[pulumi.Input[str]] = None) -> None:
        """
        VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
        :param pulumi.Input['_core.v1.PersistentVolumeSpecArgs'] inline_volume_spec: inlineVolumeSpec contains all the information necessary to attach a persistent volume defined by a pod's inline VolumeSource. This field is populated only for the CSIMigration feature. It contains translated fields from a pod's inline VolumeSource to a PersistentVolumeSpec. This field is alpha-level and is only honored by servers that enabled the CSIMigration feature.
        :param pulumi.Input[str] persistent_volume_name: Name of the persistent volume to attach.
        """
        __self__.inline_volume_spec = inline_volume_spec
        __self__.persistent_volume_name = persistent_volume_name

@pulumi.input_type
class VolumeAttachmentSpecArgs:
    attacher: pulumi.Input[str] = pulumi.input_property("attacher")
    """
    Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
    """
    node_name: pulumi.Input[str] = pulumi.input_property("nodeName")
    """
    The node that the volume should be attached to.
    """
    source: pulumi.Input['VolumeAttachmentSourceArgs'] = pulumi.input_property("source")
    """
    Source represents the volume that should be attached.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, attacher: pulumi.Input[str], node_name: pulumi.Input[str], source: pulumi.Input['VolumeAttachmentSourceArgs']) -> None:
        """
        VolumeAttachmentSpec is the specification of a VolumeAttachment request.
        :param pulumi.Input[str] attacher: Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
        :param pulumi.Input[str] node_name: The node that the volume should be attached to.
        :param pulumi.Input['VolumeAttachmentSourceArgs'] source: Source represents the volume that should be attached.
        """
        __self__.attacher = attacher
        __self__.node_name = node_name
        __self__.source = source

@pulumi.input_type
class VolumeAttachmentStatusArgs:
    attached: pulumi.Input[bool] = pulumi.input_property("attached")
    """
    Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
    """
    attach_error: Optional[pulumi.Input['VolumeErrorArgs']] = pulumi.input_property("attachError")
    """
    The last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
    """
    attachment_metadata: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = pulumi.input_property("attachmentMetadata")
    """
    Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
    """
    detach_error: Optional[pulumi.Input['VolumeErrorArgs']] = pulumi.input_property("detachError")
    """
    The last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, attached: pulumi.Input[bool], attach_error: Optional[pulumi.Input['VolumeErrorArgs']] = None, attachment_metadata: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, detach_error: Optional[pulumi.Input['VolumeErrorArgs']] = None) -> None:
        """
        VolumeAttachmentStatus is the status of a VolumeAttachment request.
        :param pulumi.Input[bool] attached: Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        :param pulumi.Input['VolumeErrorArgs'] attach_error: The last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] attachment_metadata: Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        :param pulumi.Input['VolumeErrorArgs'] detach_error: The last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.
        """
        __self__.attached = attached
        __self__.attach_error = attach_error
        __self__.attachment_metadata = attachment_metadata
        __self__.detach_error = detach_error

@pulumi.input_type
class VolumeErrorArgs:
    message: Optional[pulumi.Input[str]] = pulumi.input_property("message")
    """
    String detailing the error encountered during Attach or Detach operation. This string may be logged, so it should not contain sensitive information.
    """
    time: Optional[pulumi.Input[str]] = pulumi.input_property("time")
    """
    Time the error was encountered.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, message: Optional[pulumi.Input[str]] = None, time: Optional[pulumi.Input[str]] = None) -> None:
        """
        VolumeError captures an error encountered during a volume operation.
        :param pulumi.Input[str] message: String detailing the error encountered during Attach or Detach operation. This string may be logged, so it should not contain sensitive information.
        :param pulumi.Input[str] time: Time the error was encountered.
        """
        __self__.message = message
        __self__.time = time

@pulumi.input_type
class VolumeNodeResourcesArgs:
    count: Optional[pulumi.Input[float]] = pulumi.input_property("count")
    """
    Maximum number of unique volumes managed by the CSI driver that can be used on a node. A volume that is both attached and mounted on a node is considered to be used once, not twice. The same rule applies for a unique volume that is shared among multiple pods on the same node. If this field is nil, then the supported number of volumes on this node is unbounded.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, count: Optional[pulumi.Input[float]] = None) -> None:
        """
        VolumeNodeResources is a set of resource limits for scheduling of volumes.
        :param pulumi.Input[float] count: Maximum number of unique volumes managed by the CSI driver that can be used on a node. A volume that is both attached and mounted on a node is considered to be used once, not twice. The same rule applies for a unique volume that is shared among multiple pods on the same node. If this field is nil, then the supported number of volumes on this node is unbounded.
        """
        __self__.count = count

