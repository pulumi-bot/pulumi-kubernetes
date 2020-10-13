// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1
{

    [OutputType]
    public sealed class PodSecurityPolicySpec
    {
        /// <summary>
        /// allowPrivilegeEscalation determines if a pod can request to allow privilege escalation. If unspecified, defaults to true.
        /// </summary>
        public readonly bool AllowPrivilegeEscalation;
        /// <summary>
        /// AllowedCSIDrivers is an allowlist of inline CSI drivers that must be explicitly set to be embedded within a pod spec. An empty value indicates that any CSI driver can be used for inline ephemeral volumes. This is a beta field, and is only honored if the API server enables the CSIInlineVolume feature gate.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1.AllowedCSIDriver> AllowedCSIDrivers;
        /// <summary>
        /// allowedCapabilities is a list of capabilities that can be requested to add to the container. Capabilities in this field may be added at the pod author's discretion. You must not list a capability in both allowedCapabilities and requiredDropCapabilities.
        /// </summary>
        public readonly ImmutableArray<string> AllowedCapabilities;
        /// <summary>
        /// allowedFlexVolumes is an allowlist of Flexvolumes.  Empty or nil indicates that all Flexvolumes may be used.  This parameter is effective only when the usage of the Flexvolumes is allowed in the "volumes" field.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1.AllowedFlexVolume> AllowedFlexVolumes;
        /// <summary>
        /// allowedHostPaths is an allowlist of host paths. Empty indicates that all host paths may be used.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1.AllowedHostPath> AllowedHostPaths;
        /// <summary>
        /// AllowedProcMountTypes is an allowlist of allowed ProcMountTypes. Empty or nil indicates that only the DefaultProcMountType may be used. This requires the ProcMountType feature flag to be enabled.
        /// </summary>
        public readonly ImmutableArray<string> AllowedProcMountTypes;
        /// <summary>
        /// allowedUnsafeSysctls is a list of explicitly allowed unsafe sysctls, defaults to none. Each entry is either a plain sysctl name or ends in "*" in which case it is considered as a prefix of allowed sysctls. Single * means all unsafe sysctls are allowed. Kubelet has to allowlist all allowed unsafe sysctls explicitly to avoid rejection.
        /// 
        /// Examples: e.g. "foo/*" allows "foo/bar", "foo/baz", etc. e.g. "foo.*" allows "foo.bar", "foo.baz", etc.
        /// </summary>
        public readonly ImmutableArray<string> AllowedUnsafeSysctls;
        /// <summary>
        /// defaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability.  You may not list a capability in both defaultAddCapabilities and requiredDropCapabilities. Capabilities added here are implicitly allowed, and need not be included in the allowedCapabilities list.
        /// </summary>
        public readonly ImmutableArray<string> DefaultAddCapabilities;
        /// <summary>
        /// defaultAllowPrivilegeEscalation controls the default setting for whether a process can gain more privileges than its parent process.
        /// </summary>
        public readonly bool DefaultAllowPrivilegeEscalation;
        /// <summary>
        /// forbiddenSysctls is a list of explicitly forbidden sysctls, defaults to none. Each entry is either a plain sysctl name or ends in "*" in which case it is considered as a prefix of forbidden sysctls. Single * means all sysctls are forbidden.
        /// 
        /// Examples: e.g. "foo/*" forbids "foo/bar", "foo/baz", etc. e.g. "foo.*" forbids "foo.bar", "foo.baz", etc.
        /// </summary>
        public readonly ImmutableArray<string> ForbiddenSysctls;
        /// <summary>
        /// fsGroup is the strategy that will dictate what fs group is used by the SecurityContext.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1.FSGroupStrategyOptions FsGroup;
        /// <summary>
        /// hostIPC determines if the policy allows the use of HostIPC in the pod spec.
        /// </summary>
        public readonly bool HostIPC;
        /// <summary>
        /// hostNetwork determines if the policy allows the use of HostNetwork in the pod spec.
        /// </summary>
        public readonly bool HostNetwork;
        /// <summary>
        /// hostPID determines if the policy allows the use of HostPID in the pod spec.
        /// </summary>
        public readonly bool HostPID;
        /// <summary>
        /// hostPorts determines which host port ranges are allowed to be exposed.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1.HostPortRange> HostPorts;
        /// <summary>
        /// privileged determines if a pod can request to be run as privileged.
        /// </summary>
        public readonly bool Privileged;
        /// <summary>
        /// readOnlyRootFilesystem when set to true will force containers to run with a read only root file system.  If the container specifically requests to run with a non-read only root file system the PSP should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to.
        /// </summary>
        public readonly bool ReadOnlyRootFilesystem;
        /// <summary>
        /// requiredDropCapabilities are the capabilities that will be dropped from the container.  These are required to be dropped and cannot be added.
        /// </summary>
        public readonly ImmutableArray<string> RequiredDropCapabilities;
        /// <summary>
        /// RunAsGroup is the strategy that will dictate the allowable RunAsGroup values that may be set. If this field is omitted, the pod's RunAsGroup can take any value. This field requires the RunAsGroup feature gate to be enabled.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1.RunAsGroupStrategyOptions RunAsGroup;
        /// <summary>
        /// runAsUser is the strategy that will dictate the allowable RunAsUser values that may be set.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1.RunAsUserStrategyOptions RunAsUser;
        /// <summary>
        /// runtimeClass is the strategy that will dictate the allowable RuntimeClasses for a pod. If this field is omitted, the pod's runtimeClassName field is unrestricted. Enforcement of this field depends on the RuntimeClass feature gate being enabled.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1.RuntimeClassStrategyOptions RuntimeClass;
        /// <summary>
        /// seLinux is the strategy that will dictate the allowable labels that may be set.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1.SELinuxStrategyOptions SeLinux;
        /// <summary>
        /// supplementalGroups is the strategy that will dictate what supplemental groups are used by the SecurityContext.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1.SupplementalGroupsStrategyOptions SupplementalGroups;
        /// <summary>
        /// volumes is an allowlist of volume plugins. Empty indicates that no volumes may be used. To allow all volumes you may use '*'.
        /// </summary>
        public readonly ImmutableArray<string> Volumes;

        [OutputConstructor]
        private PodSecurityPolicySpec(
            bool allowPrivilegeEscalation,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1.AllowedCSIDriver> allowedCSIDrivers,

            ImmutableArray<string> allowedCapabilities,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1.AllowedFlexVolume> allowedFlexVolumes,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1.AllowedHostPath> allowedHostPaths,

            ImmutableArray<string> allowedProcMountTypes,

            ImmutableArray<string> allowedUnsafeSysctls,

            ImmutableArray<string> defaultAddCapabilities,

            bool defaultAllowPrivilegeEscalation,

            ImmutableArray<string> forbiddenSysctls,

            Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1.FSGroupStrategyOptions fsGroup,

            bool hostIPC,

            bool hostNetwork,

            bool hostPID,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1.HostPortRange> hostPorts,

            bool privileged,

            bool readOnlyRootFilesystem,

            ImmutableArray<string> requiredDropCapabilities,

            Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1.RunAsGroupStrategyOptions runAsGroup,

            Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1.RunAsUserStrategyOptions runAsUser,

            Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1.RuntimeClassStrategyOptions runtimeClass,

            Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1.SELinuxStrategyOptions seLinux,

            Pulumi.Kubernetes.Types.Outputs.Policy.V1Beta1.SupplementalGroupsStrategyOptions supplementalGroups,

            ImmutableArray<string> volumes)
        {
            AllowPrivilegeEscalation = allowPrivilegeEscalation;
            AllowedCSIDrivers = allowedCSIDrivers;
            AllowedCapabilities = allowedCapabilities;
            AllowedFlexVolumes = allowedFlexVolumes;
            AllowedHostPaths = allowedHostPaths;
            AllowedProcMountTypes = allowedProcMountTypes;
            AllowedUnsafeSysctls = allowedUnsafeSysctls;
            DefaultAddCapabilities = defaultAddCapabilities;
            DefaultAllowPrivilegeEscalation = defaultAllowPrivilegeEscalation;
            ForbiddenSysctls = forbiddenSysctls;
            FsGroup = fsGroup;
            HostIPC = hostIPC;
            HostNetwork = hostNetwork;
            HostPID = hostPID;
            HostPorts = hostPorts;
            Privileged = privileged;
            ReadOnlyRootFilesystem = readOnlyRootFilesystem;
            RequiredDropCapabilities = requiredDropCapabilities;
            RunAsGroup = runAsGroup;
            RunAsUser = runAsUser;
            RuntimeClass = runtimeClass;
            SeLinux = seLinux;
            SupplementalGroups = supplementalGroups;
            Volumes = volumes;
        }
    }
}
