// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// PodSecurityPolicySpec defines the policy enforced.
type PodSecurityPolicySpec struct {
	// allowPrivilegeEscalation determines if a pod can request to allow privilege escalation. If
	// unspecified, defaults to true.
	AllowPrivilegeEscalation *bool `pulumi:"allowPrivilegeEscalation"`

	// AllowedCSIDrivers is a whitelist of inline CSI drivers that must be explicitly set to be
	// embedded within a pod spec. An empty value indicates that any CSI driver can be used for inline
	// ephemeral volumes. This is an alpha field, and is only honored if the API server enables the
	// CSIInlineVolume feature gate.
	AllowedCSIDrivers []AllowedCSIDriver `pulumi:"allowedCSIDrivers"`

	// allowedCapabilities is a list of capabilities that can be requested to add to the container.
	// Capabilities in this field may be added at the pod author's discretion. You must not list a
	// capability in both allowedCapabilities and requiredDropCapabilities.
	AllowedCapabilities []string `pulumi:"allowedCapabilities"`

	// allowedFlexVolumes is a whitelist of allowed Flexvolumes.  Empty or nil indicates that all
	// Flexvolumes may be used.  This parameter is effective only when the usage of the Flexvolumes is
	// allowed in the "volumes" field.
	AllowedFlexVolumes []AllowedFlexVolume `pulumi:"allowedFlexVolumes"`

	// allowedHostPaths is a white list of allowed host paths. Empty indicates that all host paths may
	// be used.
	AllowedHostPaths []AllowedHostPath `pulumi:"allowedHostPaths"`

	// AllowedProcMountTypes is a whitelist of allowed ProcMountTypes. Empty or nil indicates that only
	// the DefaultProcMountType may be used. This requires the ProcMountType feature flag to be
	// enabled.
	AllowedProcMountTypes []string `pulumi:"allowedProcMountTypes"`

	// allowedUnsafeSysctls is a list of explicitly allowed unsafe sysctls, defaults to none. Each
	// entry is either a plain sysctl name or ends in "*" in which case it is considered as a prefix of
	// allowed sysctls. Single * means all unsafe sysctls are allowed. Kubelet has to whitelist all
	// allowed unsafe sysctls explicitly to avoid rejection.
	// 
	// Examples: e.g. "foo/*" allows "foo/bar", "foo/baz", etc. e.g. "foo.*" allows "foo.bar",
	// "foo.baz", etc.
	AllowedUnsafeSysctls []string `pulumi:"allowedUnsafeSysctls"`

	// defaultAddCapabilities is the default set of capabilities that will be added to the container
	// unless the pod spec specifically drops the capability.  You may not list a capability in both
	// defaultAddCapabilities and requiredDropCapabilities. Capabilities added here are implicitly
	// allowed, and need not be included in the allowedCapabilities list.
	DefaultAddCapabilities []string `pulumi:"defaultAddCapabilities"`

	// defaultAllowPrivilegeEscalation controls the default setting for whether a process can gain more
	// privileges than its parent process.
	DefaultAllowPrivilegeEscalation *bool `pulumi:"defaultAllowPrivilegeEscalation"`

	// forbiddenSysctls is a list of explicitly forbidden sysctls, defaults to none. Each entry is
	// either a plain sysctl name or ends in "*" in which case it is considered as a prefix of
	// forbidden sysctls. Single * means all sysctls are forbidden.
	// 
	// Examples: e.g. "foo/*" forbids "foo/bar", "foo/baz", etc. e.g. "foo.*" forbids "foo.bar",
	// "foo.baz", etc.
	ForbiddenSysctls []string `pulumi:"forbiddenSysctls"`

	// fsGroup is the strategy that will dictate what fs group is used by the SecurityContext.
	FsGroup FSGroupStrategyOptions `pulumi:"fsGroup"`

	// hostIPC determines if the policy allows the use of HostIPC in the pod spec.
	HostIPC *bool `pulumi:"hostIPC"`

	// hostNetwork determines if the policy allows the use of HostNetwork in the pod spec.
	HostNetwork *bool `pulumi:"hostNetwork"`

	// hostPID determines if the policy allows the use of HostPID in the pod spec.
	HostPID *bool `pulumi:"hostPID"`

	// hostPorts determines which host port ranges are allowed to be exposed.
	HostPorts []HostPortRange `pulumi:"hostPorts"`

	// privileged determines if a pod can request to be run as privileged.
	Privileged *bool `pulumi:"privileged"`

	// readOnlyRootFilesystem when set to true will force containers to run with a read only root file
	// system.  If the container specifically requests to run with a non-read only root file system the
	// PSP should deny the pod. If set to false the container may run with a read only root file system
	// if it wishes but it will not be forced to.
	ReadOnlyRootFilesystem *bool `pulumi:"readOnlyRootFilesystem"`

	// requiredDropCapabilities are the capabilities that will be dropped from the container.  These
	// are required to be dropped and cannot be added.
	RequiredDropCapabilities []string `pulumi:"requiredDropCapabilities"`

	// RunAsGroup is the strategy that will dictate the allowable RunAsGroup values that may be set. If
	// this field is omitted, the pod's RunAsGroup can take any value. This field requires the
	// RunAsGroup feature gate to be enabled.
	RunAsGroup *RunAsGroupStrategyOptions `pulumi:"runAsGroup"`

	// runAsUser is the strategy that will dictate the allowable RunAsUser values that may be set.
	RunAsUser RunAsUserStrategyOptions `pulumi:"runAsUser"`

	// runtimeClass is the strategy that will dictate the allowable RuntimeClasses for a pod. If this
	// field is omitted, the pod's runtimeClassName field is unrestricted. Enforcement of this field
	// depends on the RuntimeClass feature gate being enabled.
	RuntimeClass *RuntimeClassStrategyOptions `pulumi:"runtimeClass"`

	// seLinux is the strategy that will dictate the allowable labels that may be set.
	SeLinux SELinuxStrategyOptions `pulumi:"seLinux"`

	// supplementalGroups is the strategy that will dictate what supplemental groups are used by the
	// SecurityContext.
	SupplementalGroups SupplementalGroupsStrategyOptions `pulumi:"supplementalGroups"`

	// volumes is a white list of allowed volume plugins. Empty indicates that no volumes may be used.
	// To allow all volumes you may use '*'.
	Volumes []string `pulumi:"volumes"`

}

var _PodSecurityPolicySpecType = reflect.TypeOf((*PodSecurityPolicySpec)(nil)).Elem()

// PodSecurityPolicySpecInput represents an input type that resolves to a PodSecurityPolicySpec.
type PodSecurityPolicySpecInput interface {
	ElementType() reflect.Type

	ToPodSecurityPolicySpecOutput() PodSecurityPolicySpecOutput
	ToPodSecurityPolicySpecOutputWithContext(ctx context.Context) PodSecurityPolicySpecOutput
}

// PodSecurityPolicySpecArgs is a PodSecurityPolicySpecInput whose fields are all Input types.
type PodSecurityPolicySpecArgs struct {
	// fsGroup is the strategy that will dictate what fs group is used by the SecurityContext.
	FsGroup FSGroupStrategyOptionsInput `pulumi:"fsGroup"`

	// runAsUser is the strategy that will dictate the allowable RunAsUser values that may be set.
	RunAsUser RunAsUserStrategyOptionsInput `pulumi:"runAsUser"`

	// seLinux is the strategy that will dictate the allowable labels that may be set.
	SeLinux SELinuxStrategyOptionsInput `pulumi:"seLinux"`

	// supplementalGroups is the strategy that will dictate what supplemental groups are used by the
	// SecurityContext.
	SupplementalGroups SupplementalGroupsStrategyOptionsInput `pulumi:"supplementalGroups"`

	// allowPrivilegeEscalation determines if a pod can request to allow privilege escalation. If
	// unspecified, defaults to true.
	AllowPrivilegeEscalation pulumi.BoolInput `pulumi:"allowPrivilegeEscalation"`

	// AllowedCSIDrivers is a whitelist of inline CSI drivers that must be explicitly set to be
	// embedded within a pod spec. An empty value indicates that any CSI driver can be used for inline
	// ephemeral volumes. This is an alpha field, and is only honored if the API server enables the
	// CSIInlineVolume feature gate.
	AllowedCSIDrivers AllowedCSIDriverArrayInput `pulumi:"allowedCSIDrivers"`

	// allowedCapabilities is a list of capabilities that can be requested to add to the container.
	// Capabilities in this field may be added at the pod author's discretion. You must not list a
	// capability in both allowedCapabilities and requiredDropCapabilities.
	AllowedCapabilities pulumi.StringArrayInput `pulumi:"allowedCapabilities"`

	// allowedFlexVolumes is a whitelist of allowed Flexvolumes.  Empty or nil indicates that all
	// Flexvolumes may be used.  This parameter is effective only when the usage of the Flexvolumes is
	// allowed in the "volumes" field.
	AllowedFlexVolumes AllowedFlexVolumeArrayInput `pulumi:"allowedFlexVolumes"`

	// allowedHostPaths is a white list of allowed host paths. Empty indicates that all host paths may
	// be used.
	AllowedHostPaths AllowedHostPathArrayInput `pulumi:"allowedHostPaths"`

	// AllowedProcMountTypes is a whitelist of allowed ProcMountTypes. Empty or nil indicates that only
	// the DefaultProcMountType may be used. This requires the ProcMountType feature flag to be
	// enabled.
	AllowedProcMountTypes pulumi.StringArrayInput `pulumi:"allowedProcMountTypes"`

	// allowedUnsafeSysctls is a list of explicitly allowed unsafe sysctls, defaults to none. Each
	// entry is either a plain sysctl name or ends in "*" in which case it is considered as a prefix of
	// allowed sysctls. Single * means all unsafe sysctls are allowed. Kubelet has to whitelist all
	// allowed unsafe sysctls explicitly to avoid rejection.
	// 
	// Examples: e.g. "foo/*" allows "foo/bar", "foo/baz", etc. e.g. "foo.*" allows "foo.bar",
	// "foo.baz", etc.
	AllowedUnsafeSysctls pulumi.StringArrayInput `pulumi:"allowedUnsafeSysctls"`

	// defaultAddCapabilities is the default set of capabilities that will be added to the container
	// unless the pod spec specifically drops the capability.  You may not list a capability in both
	// defaultAddCapabilities and requiredDropCapabilities. Capabilities added here are implicitly
	// allowed, and need not be included in the allowedCapabilities list.
	DefaultAddCapabilities pulumi.StringArrayInput `pulumi:"defaultAddCapabilities"`

	// defaultAllowPrivilegeEscalation controls the default setting for whether a process can gain more
	// privileges than its parent process.
	DefaultAllowPrivilegeEscalation pulumi.BoolInput `pulumi:"defaultAllowPrivilegeEscalation"`

	// forbiddenSysctls is a list of explicitly forbidden sysctls, defaults to none. Each entry is
	// either a plain sysctl name or ends in "*" in which case it is considered as a prefix of
	// forbidden sysctls. Single * means all sysctls are forbidden.
	// 
	// Examples: e.g. "foo/*" forbids "foo/bar", "foo/baz", etc. e.g. "foo.*" forbids "foo.bar",
	// "foo.baz", etc.
	ForbiddenSysctls pulumi.StringArrayInput `pulumi:"forbiddenSysctls"`

	// hostIPC determines if the policy allows the use of HostIPC in the pod spec.
	HostIPC pulumi.BoolInput `pulumi:"hostIPC"`

	// hostNetwork determines if the policy allows the use of HostNetwork in the pod spec.
	HostNetwork pulumi.BoolInput `pulumi:"hostNetwork"`

	// hostPID determines if the policy allows the use of HostPID in the pod spec.
	HostPID pulumi.BoolInput `pulumi:"hostPID"`

	// hostPorts determines which host port ranges are allowed to be exposed.
	HostPorts HostPortRangeArrayInput `pulumi:"hostPorts"`

	// privileged determines if a pod can request to be run as privileged.
	Privileged pulumi.BoolInput `pulumi:"privileged"`

	// readOnlyRootFilesystem when set to true will force containers to run with a read only root file
	// system.  If the container specifically requests to run with a non-read only root file system the
	// PSP should deny the pod. If set to false the container may run with a read only root file system
	// if it wishes but it will not be forced to.
	ReadOnlyRootFilesystem pulumi.BoolInput `pulumi:"readOnlyRootFilesystem"`

	// requiredDropCapabilities are the capabilities that will be dropped from the container.  These
	// are required to be dropped and cannot be added.
	RequiredDropCapabilities pulumi.StringArrayInput `pulumi:"requiredDropCapabilities"`

	// RunAsGroup is the strategy that will dictate the allowable RunAsGroup values that may be set. If
	// this field is omitted, the pod's RunAsGroup can take any value. This field requires the
	// RunAsGroup feature gate to be enabled.
	RunAsGroup RunAsGroupStrategyOptionsInput `pulumi:"runAsGroup"`

	// runtimeClass is the strategy that will dictate the allowable RuntimeClasses for a pod. If this
	// field is omitted, the pod's runtimeClassName field is unrestricted. Enforcement of this field
	// depends on the RuntimeClass feature gate being enabled.
	RuntimeClass RuntimeClassStrategyOptionsInput `pulumi:"runtimeClass"`

	// volumes is a white list of allowed volume plugins. Empty indicates that no volumes may be used.
	// To allow all volumes you may use '*'.
	Volumes pulumi.StringArrayInput `pulumi:"volumes"`

}

func (a PodSecurityPolicySpecArgs) ElementType() reflect.Type {
	return _PodSecurityPolicySpecType
}

func (a PodSecurityPolicySpecArgs) ToPodSecurityPolicySpecOutput() PodSecurityPolicySpecOutput {
	return pulumi.ToOutput(a).(PodSecurityPolicySpecOutput)
}

func (a PodSecurityPolicySpecArgs) ToPodSecurityPolicySpecOutputWithContext(ctx context.Context) PodSecurityPolicySpecOutput {
	return pulumi.ToOutputWithContext(ctx, a).(PodSecurityPolicySpecOutput)
}

// PodSecurityPolicySpecOutput is an output type that resolves to a Input.
type PodSecurityPolicySpecOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(PodSecurityPolicySpecOutput{}) }

func (PodSecurityPolicySpecOutput) ElementType() reflect.Type {
	return _PodSecurityPolicySpecType
}

func (o PodSecurityPolicySpecOutput) AllowPrivilegeEscalation() pulumi.BoolOutput {
	return o.Apply(func(v PodSecurityPolicySpec) *bool {
		return v.AllowPrivilegeEscalation
	}).(pulumi.BoolOutput)
}

func (o PodSecurityPolicySpecOutput) AllowedCSIDrivers() AllowedCSIDriverArrayOutput {
	return o.Apply(func(v PodSecurityPolicySpec) []AllowedCSIDriver {
		return v.AllowedCSIDrivers
	}).(AllowedCSIDriverArrayOutput)
}

func (o PodSecurityPolicySpecOutput) AllowedCapabilities() pulumi.StringArrayOutput {
	return o.Apply(func(v PodSecurityPolicySpec) []string {
		return v.AllowedCapabilities
	}).(pulumi.StringArrayOutput)
}

func (o PodSecurityPolicySpecOutput) AllowedFlexVolumes() AllowedFlexVolumeArrayOutput {
	return o.Apply(func(v PodSecurityPolicySpec) []AllowedFlexVolume {
		return v.AllowedFlexVolumes
	}).(AllowedFlexVolumeArrayOutput)
}

func (o PodSecurityPolicySpecOutput) AllowedHostPaths() AllowedHostPathArrayOutput {
	return o.Apply(func(v PodSecurityPolicySpec) []AllowedHostPath {
		return v.AllowedHostPaths
	}).(AllowedHostPathArrayOutput)
}

func (o PodSecurityPolicySpecOutput) AllowedProcMountTypes() pulumi.StringArrayOutput {
	return o.Apply(func(v PodSecurityPolicySpec) []string {
		return v.AllowedProcMountTypes
	}).(pulumi.StringArrayOutput)
}

func (o PodSecurityPolicySpecOutput) AllowedUnsafeSysctls() pulumi.StringArrayOutput {
	return o.Apply(func(v PodSecurityPolicySpec) []string {
		return v.AllowedUnsafeSysctls
	}).(pulumi.StringArrayOutput)
}

func (o PodSecurityPolicySpecOutput) DefaultAddCapabilities() pulumi.StringArrayOutput {
	return o.Apply(func(v PodSecurityPolicySpec) []string {
		return v.DefaultAddCapabilities
	}).(pulumi.StringArrayOutput)
}

func (o PodSecurityPolicySpecOutput) DefaultAllowPrivilegeEscalation() pulumi.BoolOutput {
	return o.Apply(func(v PodSecurityPolicySpec) *bool {
		return v.DefaultAllowPrivilegeEscalation
	}).(pulumi.BoolOutput)
}

func (o PodSecurityPolicySpecOutput) ForbiddenSysctls() pulumi.StringArrayOutput {
	return o.Apply(func(v PodSecurityPolicySpec) []string {
		return v.ForbiddenSysctls
	}).(pulumi.StringArrayOutput)
}

func (o PodSecurityPolicySpecOutput) FsGroup() FSGroupStrategyOptionsOutput {
	return o.Apply(func(v PodSecurityPolicySpec) FSGroupStrategyOptions {
		return v.FsGroup
	}).(FSGroupStrategyOptionsOutput)
}

func (o PodSecurityPolicySpecOutput) HostIPC() pulumi.BoolOutput {
	return o.Apply(func(v PodSecurityPolicySpec) *bool {
		return v.HostIPC
	}).(pulumi.BoolOutput)
}

func (o PodSecurityPolicySpecOutput) HostNetwork() pulumi.BoolOutput {
	return o.Apply(func(v PodSecurityPolicySpec) *bool {
		return v.HostNetwork
	}).(pulumi.BoolOutput)
}

func (o PodSecurityPolicySpecOutput) HostPID() pulumi.BoolOutput {
	return o.Apply(func(v PodSecurityPolicySpec) *bool {
		return v.HostPID
	}).(pulumi.BoolOutput)
}

func (o PodSecurityPolicySpecOutput) HostPorts() HostPortRangeArrayOutput {
	return o.Apply(func(v PodSecurityPolicySpec) []HostPortRange {
		return v.HostPorts
	}).(HostPortRangeArrayOutput)
}

func (o PodSecurityPolicySpecOutput) Privileged() pulumi.BoolOutput {
	return o.Apply(func(v PodSecurityPolicySpec) *bool {
		return v.Privileged
	}).(pulumi.BoolOutput)
}

func (o PodSecurityPolicySpecOutput) ReadOnlyRootFilesystem() pulumi.BoolOutput {
	return o.Apply(func(v PodSecurityPolicySpec) *bool {
		return v.ReadOnlyRootFilesystem
	}).(pulumi.BoolOutput)
}

func (o PodSecurityPolicySpecOutput) RequiredDropCapabilities() pulumi.StringArrayOutput {
	return o.Apply(func(v PodSecurityPolicySpec) []string {
		return v.RequiredDropCapabilities
	}).(pulumi.StringArrayOutput)
}

func (o PodSecurityPolicySpecOutput) RunAsGroup() RunAsGroupStrategyOptionsOutput {
	return o.Apply(func(v PodSecurityPolicySpec) *RunAsGroupStrategyOptions {
		return v.RunAsGroup
	}).(RunAsGroupStrategyOptionsOutput)
}

func (o PodSecurityPolicySpecOutput) RunAsUser() RunAsUserStrategyOptionsOutput {
	return o.Apply(func(v PodSecurityPolicySpec) RunAsUserStrategyOptions {
		return v.RunAsUser
	}).(RunAsUserStrategyOptionsOutput)
}

func (o PodSecurityPolicySpecOutput) RuntimeClass() RuntimeClassStrategyOptionsOutput {
	return o.Apply(func(v PodSecurityPolicySpec) *RuntimeClassStrategyOptions {
		return v.RuntimeClass
	}).(RuntimeClassStrategyOptionsOutput)
}

func (o PodSecurityPolicySpecOutput) SeLinux() SELinuxStrategyOptionsOutput {
	return o.Apply(func(v PodSecurityPolicySpec) SELinuxStrategyOptions {
		return v.SeLinux
	}).(SELinuxStrategyOptionsOutput)
}

func (o PodSecurityPolicySpecOutput) SupplementalGroups() SupplementalGroupsStrategyOptionsOutput {
	return o.Apply(func(v PodSecurityPolicySpec) SupplementalGroupsStrategyOptions {
		return v.SupplementalGroups
	}).(SupplementalGroupsStrategyOptionsOutput)
}

func (o PodSecurityPolicySpecOutput) Volumes() pulumi.StringArrayOutput {
	return o.Apply(func(v PodSecurityPolicySpec) []string {
		return v.Volumes
	}).(pulumi.StringArrayOutput)
}

