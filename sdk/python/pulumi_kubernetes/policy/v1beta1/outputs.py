# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ... import core as _core
from ... import meta as _meta

__all__ = [
    'AllowedCSIDriver',
    'AllowedFlexVolume',
    'AllowedHostPath',
    'FSGroupStrategyOptions',
    'HostPortRange',
    'IDRange',
    'PodDisruptionBudget',
    'PodDisruptionBudgetSpec',
    'PodDisruptionBudgetStatus',
    'PodSecurityPolicy',
    'PodSecurityPolicySpec',
    'RunAsGroupStrategyOptions',
    'RunAsUserStrategyOptions',
    'RuntimeClassStrategyOptions',
    'SELinuxStrategyOptions',
    'SupplementalGroupsStrategyOptions',
]

@pulumi.output_type
class AllowedCSIDriver(dict):
    """
    AllowedCSIDriver represents a single inline CSI Driver that is allowed to be used.
    """
    name: str = pulumi.output_property("name")
    """
    Name is the registered name of the CSI driver
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AllowedFlexVolume(dict):
    """
    AllowedFlexVolume represents a single Flexvolume that is allowed to be used.
    """
    driver: str = pulumi.output_property("driver")
    """
    driver is the name of the Flexvolume driver.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AllowedHostPath(dict):
    """
    AllowedHostPath defines the host volume conditions that will be enabled by a policy for pods to use. It requires the path prefix to be defined.
    """
    path_prefix: Optional[str] = pulumi.output_property("pathPrefix")
    """
    pathPrefix is the path prefix that the host volume must match. It does not support `*`. Trailing slashes are trimmed when validating the path prefix with a host path.

    Examples: `/foo` would allow `/foo`, `/foo/` and `/foo/bar` `/foo` would not allow `/food` or `/etc/foo`
    """
    read_only: Optional[bool] = pulumi.output_property("readOnly")
    """
    when set to true, will allow host volumes matching the pathPrefix only if all volume mounts are readOnly.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class FSGroupStrategyOptions(dict):
    """
    FSGroupStrategyOptions defines the strategy type and options used to create the strategy.
    """
    ranges: Optional[List['outputs.IDRange']] = pulumi.output_property("ranges")
    """
    ranges are the allowed ranges of fs groups.  If you would like to force a single fs group then supply a single range with the same start and end. Required for MustRunAs.
    """
    rule: Optional[str] = pulumi.output_property("rule")
    """
    rule is the strategy that will dictate what FSGroup is used in the SecurityContext.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class HostPortRange(dict):
    """
    HostPortRange defines a range of host ports that will be enabled by a policy for pods to use.  It requires both the start and end to be defined.
    """
    max: float = pulumi.output_property("max")
    """
    max is the end of the range, inclusive.
    """
    min: float = pulumi.output_property("min")
    """
    min is the start of the range, inclusive.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IDRange(dict):
    """
    IDRange provides a min/max of an allowed range of IDs.
    """
    max: float = pulumi.output_property("max")
    """
    max is the end of the range, inclusive.
    """
    min: float = pulumi.output_property("min")
    """
    min is the start of the range, inclusive.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PodDisruptionBudget(dict):
    """
    PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods
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
    spec: Optional['outputs.PodDisruptionBudgetSpec'] = pulumi.output_property("spec")
    """
    Specification of the desired behavior of the PodDisruptionBudget.
    """
    status: Optional['outputs.PodDisruptionBudgetStatus'] = pulumi.output_property("status")
    """
    Most recently observed status of the PodDisruptionBudget.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PodDisruptionBudgetSpec(dict):
    """
    PodDisruptionBudgetSpec is a description of a PodDisruptionBudget.
    """
    max_unavailable: Optional[Any] = pulumi.output_property("maxUnavailable")
    """
    An eviction is allowed if at most "maxUnavailable" pods selected by "selector" are unavailable after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a mutually exclusive setting with "minAvailable".
    """
    min_available: Optional[Any] = pulumi.output_property("minAvailable")
    """
    An eviction is allowed if at least "minAvailable" pods selected by "selector" will still be available after the eviction, i.e. even in the absence of the evicted pod.  So for example you can prevent all voluntary evictions by specifying "100%".
    """
    selector: Optional['_meta.v1.outputs.LabelSelector'] = pulumi.output_property("selector")
    """
    Label query over pods whose evictions are managed by the disruption budget.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PodDisruptionBudgetStatus(dict):
    """
    PodDisruptionBudgetStatus represents information about the status of a PodDisruptionBudget. Status may trail the actual state of a system.
    """
    current_healthy: float = pulumi.output_property("currentHealthy")
    """
    current number of healthy pods
    """
    desired_healthy: float = pulumi.output_property("desiredHealthy")
    """
    minimum desired number of healthy pods
    """
    disrupted_pods: Optional[Dict[str, str]] = pulumi.output_property("disruptedPods")
    """
    DisruptedPods contains information about pods whose eviction was processed by the API server eviction subresource handler but has not yet been observed by the PodDisruptionBudget controller. A pod will be in this map from the time when the API server processed the eviction request to the time when the pod is seen by PDB controller as having been marked for deletion (or after a timeout). The key in the map is the name of the pod and the value is the time when the API server processed the eviction request. If the deletion didn't occur and a pod is still there it will be removed from the list automatically by PodDisruptionBudget controller after some time. If everything goes smooth this map should be empty for the most of the time. Large number of entries in the map may indicate problems with pod deletions.
    """
    disruptions_allowed: float = pulumi.output_property("disruptionsAllowed")
    """
    Number of pod disruptions that are currently allowed.
    """
    expected_pods: float = pulumi.output_property("expectedPods")
    """
    total number of pods counted by this disruption budget
    """
    observed_generation: Optional[float] = pulumi.output_property("observedGeneration")
    """
    Most recent generation observed when updating this PDB status. DisruptionsAllowed and other status information is valid only if observedGeneration equals to PDB's object generation.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PodSecurityPolicy(dict):
    """
    PodSecurityPolicy governs the ability to make requests that affect the Security Context that will be applied to a pod and container.
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
    Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
    """
    spec: Optional['outputs.PodSecurityPolicySpec'] = pulumi.output_property("spec")
    """
    spec defines the policy enforced.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PodSecurityPolicySpec(dict):
    """
    PodSecurityPolicySpec defines the policy enforced.
    """
    allow_privilege_escalation: Optional[bool] = pulumi.output_property("allowPrivilegeEscalation")
    """
    allowPrivilegeEscalation determines if a pod can request to allow privilege escalation. If unspecified, defaults to true.
    """
    allowed_csi_drivers: Optional[List['outputs.AllowedCSIDriver']] = pulumi.output_property("allowedCSIDrivers")
    """
    AllowedCSIDrivers is a whitelist of inline CSI drivers that must be explicitly set to be embedded within a pod spec. An empty value indicates that any CSI driver can be used for inline ephemeral volumes. This is an alpha field, and is only honored if the API server enables the CSIInlineVolume feature gate.
    """
    allowed_capabilities: Optional[List[str]] = pulumi.output_property("allowedCapabilities")
    """
    allowedCapabilities is a list of capabilities that can be requested to add to the container. Capabilities in this field may be added at the pod author's discretion. You must not list a capability in both allowedCapabilities and requiredDropCapabilities.
    """
    allowed_flex_volumes: Optional[List['outputs.AllowedFlexVolume']] = pulumi.output_property("allowedFlexVolumes")
    """
    allowedFlexVolumes is a whitelist of allowed Flexvolumes.  Empty or nil indicates that all Flexvolumes may be used.  This parameter is effective only when the usage of the Flexvolumes is allowed in the "volumes" field.
    """
    allowed_host_paths: Optional[List['outputs.AllowedHostPath']] = pulumi.output_property("allowedHostPaths")
    """
    allowedHostPaths is a white list of allowed host paths. Empty indicates that all host paths may be used.
    """
    allowed_proc_mount_types: Optional[List[str]] = pulumi.output_property("allowedProcMountTypes")
    """
    AllowedProcMountTypes is a whitelist of allowed ProcMountTypes. Empty or nil indicates that only the DefaultProcMountType may be used. This requires the ProcMountType feature flag to be enabled.
    """
    allowed_unsafe_sysctls: Optional[List[str]] = pulumi.output_property("allowedUnsafeSysctls")
    """
    allowedUnsafeSysctls is a list of explicitly allowed unsafe sysctls, defaults to none. Each entry is either a plain sysctl name or ends in "*" in which case it is considered as a prefix of allowed sysctls. Single * means all unsafe sysctls are allowed. Kubelet has to whitelist all allowed unsafe sysctls explicitly to avoid rejection.

    Examples: e.g. "foo/*" allows "foo/bar", "foo/baz", etc. e.g. "foo.*" allows "foo.bar", "foo.baz", etc.
    """
    default_add_capabilities: Optional[List[str]] = pulumi.output_property("defaultAddCapabilities")
    """
    defaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability.  You may not list a capability in both defaultAddCapabilities and requiredDropCapabilities. Capabilities added here are implicitly allowed, and need not be included in the allowedCapabilities list.
    """
    default_allow_privilege_escalation: Optional[bool] = pulumi.output_property("defaultAllowPrivilegeEscalation")
    """
    defaultAllowPrivilegeEscalation controls the default setting for whether a process can gain more privileges than its parent process.
    """
    forbidden_sysctls: Optional[List[str]] = pulumi.output_property("forbiddenSysctls")
    """
    forbiddenSysctls is a list of explicitly forbidden sysctls, defaults to none. Each entry is either a plain sysctl name or ends in "*" in which case it is considered as a prefix of forbidden sysctls. Single * means all sysctls are forbidden.

    Examples: e.g. "foo/*" forbids "foo/bar", "foo/baz", etc. e.g. "foo.*" forbids "foo.bar", "foo.baz", etc.
    """
    fs_group: 'outputs.FSGroupStrategyOptions' = pulumi.output_property("fsGroup")
    """
    fsGroup is the strategy that will dictate what fs group is used by the SecurityContext.
    """
    host_ipc: Optional[bool] = pulumi.output_property("hostIPC")
    """
    hostIPC determines if the policy allows the use of HostIPC in the pod spec.
    """
    host_network: Optional[bool] = pulumi.output_property("hostNetwork")
    """
    hostNetwork determines if the policy allows the use of HostNetwork in the pod spec.
    """
    host_pid: Optional[bool] = pulumi.output_property("hostPID")
    """
    hostPID determines if the policy allows the use of HostPID in the pod spec.
    """
    host_ports: Optional[List['outputs.HostPortRange']] = pulumi.output_property("hostPorts")
    """
    hostPorts determines which host port ranges are allowed to be exposed.
    """
    privileged: Optional[bool] = pulumi.output_property("privileged")
    """
    privileged determines if a pod can request to be run as privileged.
    """
    read_only_root_filesystem: Optional[bool] = pulumi.output_property("readOnlyRootFilesystem")
    """
    readOnlyRootFilesystem when set to true will force containers to run with a read only root file system.  If the container specifically requests to run with a non-read only root file system the PSP should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to.
    """
    required_drop_capabilities: Optional[List[str]] = pulumi.output_property("requiredDropCapabilities")
    """
    requiredDropCapabilities are the capabilities that will be dropped from the container.  These are required to be dropped and cannot be added.
    """
    run_as_group: Optional['outputs.RunAsGroupStrategyOptions'] = pulumi.output_property("runAsGroup")
    """
    RunAsGroup is the strategy that will dictate the allowable RunAsGroup values that may be set. If this field is omitted, the pod's RunAsGroup can take any value. This field requires the RunAsGroup feature gate to be enabled.
    """
    run_as_user: 'outputs.RunAsUserStrategyOptions' = pulumi.output_property("runAsUser")
    """
    runAsUser is the strategy that will dictate the allowable RunAsUser values that may be set.
    """
    runtime_class: Optional['outputs.RuntimeClassStrategyOptions'] = pulumi.output_property("runtimeClass")
    """
    runtimeClass is the strategy that will dictate the allowable RuntimeClasses for a pod. If this field is omitted, the pod's runtimeClassName field is unrestricted. Enforcement of this field depends on the RuntimeClass feature gate being enabled.
    """
    se_linux: 'outputs.SELinuxStrategyOptions' = pulumi.output_property("seLinux")
    """
    seLinux is the strategy that will dictate the allowable labels that may be set.
    """
    supplemental_groups: 'outputs.SupplementalGroupsStrategyOptions' = pulumi.output_property("supplementalGroups")
    """
    supplementalGroups is the strategy that will dictate what supplemental groups are used by the SecurityContext.
    """
    volumes: Optional[List[str]] = pulumi.output_property("volumes")
    """
    volumes is a white list of allowed volume plugins. Empty indicates that no volumes may be used. To allow all volumes you may use '*'.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class RunAsGroupStrategyOptions(dict):
    """
    RunAsGroupStrategyOptions defines the strategy type and any options used to create the strategy.
    """
    ranges: Optional[List['outputs.IDRange']] = pulumi.output_property("ranges")
    """
    ranges are the allowed ranges of gids that may be used. If you would like to force a single gid then supply a single range with the same start and end. Required for MustRunAs.
    """
    rule: str = pulumi.output_property("rule")
    """
    rule is the strategy that will dictate the allowable RunAsGroup values that may be set.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class RunAsUserStrategyOptions(dict):
    """
    RunAsUserStrategyOptions defines the strategy type and any options used to create the strategy.
    """
    ranges: Optional[List['outputs.IDRange']] = pulumi.output_property("ranges")
    """
    ranges are the allowed ranges of uids that may be used. If you would like to force a single uid then supply a single range with the same start and end. Required for MustRunAs.
    """
    rule: str = pulumi.output_property("rule")
    """
    rule is the strategy that will dictate the allowable RunAsUser values that may be set.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class RuntimeClassStrategyOptions(dict):
    """
    RuntimeClassStrategyOptions define the strategy that will dictate the allowable RuntimeClasses for a pod.
    """
    allowed_runtime_class_names: List[str] = pulumi.output_property("allowedRuntimeClassNames")
    """
    allowedRuntimeClassNames is a whitelist of RuntimeClass names that may be specified on a pod. A value of "*" means that any RuntimeClass name is allowed, and must be the only item in the list. An empty list requires the RuntimeClassName field to be unset.
    """
    default_runtime_class_name: Optional[str] = pulumi.output_property("defaultRuntimeClassName")
    """
    defaultRuntimeClassName is the default RuntimeClassName to set on the pod. The default MUST be allowed by the allowedRuntimeClassNames list. A value of nil does not mutate the Pod.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SELinuxStrategyOptions(dict):
    """
    SELinuxStrategyOptions defines the strategy type and any options used to create the strategy.
    """
    rule: str = pulumi.output_property("rule")
    """
    rule is the strategy that will dictate the allowable labels that may be set.
    """
    se_linux_options: Optional['_core.v1.outputs.SELinuxOptions'] = pulumi.output_property("seLinuxOptions")
    """
    seLinuxOptions required to run as; required for MustRunAs More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SupplementalGroupsStrategyOptions(dict):
    """
    SupplementalGroupsStrategyOptions defines the strategy type and options used to create the strategy.
    """
    ranges: Optional[List['outputs.IDRange']] = pulumi.output_property("ranges")
    """
    ranges are the allowed ranges of supplemental groups.  If you would like to force a single supplemental group then supply a single range with the same start and end. Required for MustRunAs.
    """
    rule: Optional[str] = pulumi.output_property("rule")
    """
    rule is the strategy that will dictate what supplemental groups is used in the SecurityContext.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


