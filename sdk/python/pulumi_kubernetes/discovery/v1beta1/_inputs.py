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

@pulumi.input_type
class EndpointArgs:
    addresses: pulumi.Input[List[pulumi.Input[str]]] = pulumi.input_property("addresses")
    """
    addresses of this endpoint. The contents of this field are interpreted according to the corresponding EndpointSlice addressType field. Consumers must handle different types of addresses in the context of their own capabilities. This must contain at least one address but no more than 100.
    """
    conditions: Optional[pulumi.Input['EndpointConditionsArgs']] = pulumi.input_property("conditions")
    """
    conditions contains information about the current status of the endpoint.
    """
    hostname: Optional[pulumi.Input[str]] = pulumi.input_property("hostname")
    """
    hostname of this endpoint. This field may be used by consumers of endpoints to distinguish endpoints from each other (e.g. in DNS names). Multiple endpoints which use the same hostname should be considered fungible (e.g. multiple A values in DNS). Must pass DNS Label (RFC 1123) validation.
    """
    target_ref: Optional[pulumi.Input['_core.v1.ObjectReferenceArgs']] = pulumi.input_property("targetRef")
    """
    targetRef is a reference to a Kubernetes object that represents this endpoint.
    """
    topology: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = pulumi.input_property("topology")
    """
    topology contains arbitrary topology information associated with the endpoint. These key/value pairs must conform with the label format. https://kubernetes.io/docs/concepts/overview/working-with-objects/labels Topology may include a maximum of 16 key/value pairs. This includes, but is not limited to the following well known keys: * kubernetes.io/hostname: the value indicates the hostname of the node
      where the endpoint is located. This should match the corresponding
      node label.
    * topology.kubernetes.io/zone: the value indicates the zone where the
      endpoint is located. This should match the corresponding node label.
    * topology.kubernetes.io/region: the value indicates the region where the
      endpoint is located. This should match the corresponding node label.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, addresses: pulumi.Input[List[pulumi.Input[str]]], conditions: Optional[pulumi.Input['EndpointConditionsArgs']] = None, hostname: Optional[pulumi.Input[str]] = None, target_ref: Optional[pulumi.Input['_core.v1.ObjectReferenceArgs']] = None, topology: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None) -> None:
        """
        Endpoint represents a single logical "backend" implementing a service.
        :param pulumi.Input[List[pulumi.Input[str]]] addresses: addresses of this endpoint. The contents of this field are interpreted according to the corresponding EndpointSlice addressType field. Consumers must handle different types of addresses in the context of their own capabilities. This must contain at least one address but no more than 100.
        :param pulumi.Input['EndpointConditionsArgs'] conditions: conditions contains information about the current status of the endpoint.
        :param pulumi.Input[str] hostname: hostname of this endpoint. This field may be used by consumers of endpoints to distinguish endpoints from each other (e.g. in DNS names). Multiple endpoints which use the same hostname should be considered fungible (e.g. multiple A values in DNS). Must pass DNS Label (RFC 1123) validation.
        :param pulumi.Input['_core.v1.ObjectReferenceArgs'] target_ref: targetRef is a reference to a Kubernetes object that represents this endpoint.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] topology: topology contains arbitrary topology information associated with the endpoint. These key/value pairs must conform with the label format. https://kubernetes.io/docs/concepts/overview/working-with-objects/labels Topology may include a maximum of 16 key/value pairs. This includes, but is not limited to the following well known keys: * kubernetes.io/hostname: the value indicates the hostname of the node
                 where the endpoint is located. This should match the corresponding
                 node label.
               * topology.kubernetes.io/zone: the value indicates the zone where the
                 endpoint is located. This should match the corresponding node label.
               * topology.kubernetes.io/region: the value indicates the region where the
                 endpoint is located. This should match the corresponding node label.
        """
        __self__.addresses = addresses
        __self__.conditions = conditions
        __self__.hostname = hostname
        __self__.target_ref = target_ref
        __self__.topology = topology

@pulumi.input_type
class EndpointConditionsArgs:
    ready: Optional[pulumi.Input[bool]] = pulumi.input_property("ready")
    """
    ready indicates that this endpoint is prepared to receive traffic, according to whatever system is managing the endpoint. A nil value indicates an unknown state. In most cases consumers should interpret this unknown state as ready.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, ready: Optional[pulumi.Input[bool]] = None) -> None:
        """
        EndpointConditions represents the current condition of an endpoint.
        :param pulumi.Input[bool] ready: ready indicates that this endpoint is prepared to receive traffic, according to whatever system is managing the endpoint. A nil value indicates an unknown state. In most cases consumers should interpret this unknown state as ready.
        """
        __self__.ready = ready

@pulumi.input_type
class EndpointPortArgs:
    app_protocol: Optional[pulumi.Input[str]] = pulumi.input_property("appProtocol")
    """
    The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and http://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol.
    """
    name: Optional[pulumi.Input[str]] = pulumi.input_property("name")
    """
    The name of this port. All ports in an EndpointSlice must have a unique name. If the EndpointSlice is dervied from a Kubernetes service, this corresponds to the Service.ports[].name. Name must either be an empty string or pass DNS_LABEL validation: * must be no more than 63 characters long. * must consist of lower case alphanumeric characters or '-'. * must start and end with an alphanumeric character. Default is empty string.
    """
    port: Optional[pulumi.Input[float]] = pulumi.input_property("port")
    """
    The port number of the endpoint. If this is not specified, ports are not restricted and must be interpreted in the context of the specific consumer.
    """
    protocol: Optional[pulumi.Input[str]] = pulumi.input_property("protocol")
    """
    The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, app_protocol: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, port: Optional[pulumi.Input[float]] = None, protocol: Optional[pulumi.Input[str]] = None) -> None:
        """
        EndpointPort represents a Port used by an EndpointSlice
        :param pulumi.Input[str] app_protocol: The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and http://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol.
        :param pulumi.Input[str] name: The name of this port. All ports in an EndpointSlice must have a unique name. If the EndpointSlice is dervied from a Kubernetes service, this corresponds to the Service.ports[].name. Name must either be an empty string or pass DNS_LABEL validation: * must be no more than 63 characters long. * must consist of lower case alphanumeric characters or '-'. * must start and end with an alphanumeric character. Default is empty string.
        :param pulumi.Input[float] port: The port number of the endpoint. If this is not specified, ports are not restricted and must be interpreted in the context of the specific consumer.
        :param pulumi.Input[str] protocol: The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
        """
        __self__.app_protocol = app_protocol
        __self__.name = name
        __self__.port = port
        __self__.protocol = protocol

@pulumi.input_type
class EndpointSliceArgs:
    address_type: pulumi.Input[str] = pulumi.input_property("addressType")
    """
    addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.
    """
    endpoints: pulumi.Input[List[pulumi.Input['EndpointArgs']]] = pulumi.input_property("endpoints")
    """
    endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.
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
    Standard object's metadata.
    """
    ports: Optional[pulumi.Input[List[pulumi.Input['EndpointPortArgs']]]] = pulumi.input_property("ports")
    """
    ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates "all ports". Each slice may include a maximum of 100 ports.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, address_type: pulumi.Input[str], endpoints: pulumi.Input[List[pulumi.Input['EndpointArgs']]], api_version: Optional[pulumi.Input[str]] = None, kind: Optional[pulumi.Input[str]] = None, metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None, ports: Optional[pulumi.Input[List[pulumi.Input['EndpointPortArgs']]]] = None) -> None:
        """
        EndpointSlice represents a subset of the endpoints that implement a service. For a given service there may be multiple EndpointSlice objects, selected by labels, which must be joined to produce the full set of endpoints.
        :param pulumi.Input[str] address_type: addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.
        :param pulumi.Input[List[pulumi.Input['EndpointArgs']]] endpoints: endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object's metadata.
        :param pulumi.Input[List[pulumi.Input['EndpointPortArgs']]] ports: ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates "all ports". Each slice may include a maximum of 100 ports.
        """
        __self__.address_type = address_type
        __self__.endpoints = endpoints
        __self__.api_version = 'discovery.k8s.io/v1beta1'
        __self__.kind = 'EndpointSlice'
        __self__.metadata = metadata
        __self__.ports = ports

