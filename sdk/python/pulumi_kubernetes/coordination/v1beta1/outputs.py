# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from ... import _utilities, _tables
from ... import meta as _meta

__all__ = [
    'Lease',
    'LeaseSpec',
]

@pulumi.output_type
class Lease(dict):
    """
    Lease defines a lease concept.
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
    More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
    """
    spec: Optional['outputs.LeaseSpec'] = pulumi.output_property("spec")
    """
    Specification of the Lease. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class LeaseSpec(dict):
    """
    LeaseSpec is a specification of a Lease.
    """
    acquire_time: Optional[str] = pulumi.output_property("acquireTime")
    """
    acquireTime is a time when the current lease was acquired.
    """
    holder_identity: Optional[str] = pulumi.output_property("holderIdentity")
    """
    holderIdentity contains the identity of the holder of a current lease.
    """
    lease_duration_seconds: Optional[float] = pulumi.output_property("leaseDurationSeconds")
    """
    leaseDurationSeconds is a duration that candidates for a lease need to wait to force acquire it. This is measure against time of last observed RenewTime.
    """
    lease_transitions: Optional[float] = pulumi.output_property("leaseTransitions")
    """
    leaseTransitions is the number of transitions of a lease between holders.
    """
    renew_time: Optional[str] = pulumi.output_property("renewTime")
    """
    renewTime is a time when the current holder of a lease has last updated the lease.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


