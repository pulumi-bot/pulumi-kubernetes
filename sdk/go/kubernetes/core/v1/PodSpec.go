// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// PodSpec is a description of a pod.
type PodSpec struct {
	// Optional duration in seconds the pod may be active on the node relative to StartTime before the
	// system will actively try to mark it failed and kill associated containers. Value must be a
	// positive integer.
	ActiveDeadlineSeconds *int `pulumi:"activeDeadlineSeconds"`

	// If specified, the pod's scheduling constraints
	Affinity *Affinity `pulumi:"affinity"`

	// AutomountServiceAccountToken indicates whether a service account token should be automatically
	// mounted.
	AutomountServiceAccountToken *bool `pulumi:"automountServiceAccountToken"`

	// List of containers belonging to the pod. Containers cannot currently be added or removed. There
	// must be at least one container in a Pod. Cannot be updated.
	Containers []Container `pulumi:"containers"`

	// Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated
	// DNS configuration based on DNSPolicy.
	DnsConfig *PodDNSConfig `pulumi:"dnsConfig"`

	// Set DNS policy for the pod. Defaults to "ClusterFirst". Valid values are
	// 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in
	// DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along
	// with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.
	DnsPolicy *string `pulumi:"dnsPolicy"`

	// EnableServiceLinks indicates whether information about services should be injected into pod's
	// environment variables, matching the syntax of Docker links. Optional: Defaults to true.
	EnableServiceLinks *bool `pulumi:"enableServiceLinks"`

	// List of ephemeral containers run in this pod. Ephemeral containers may be run in an existing pod
	// to perform user-initiated actions such as debugging. This list cannot be specified when creating
	// a pod, and it cannot be modified by updating the pod spec. In order to add an ephemeral
	// container to an existing pod, use the pod's ephemeralcontainers subresource. This field is
	// alpha-level and is only honored by servers that enable the EphemeralContainers feature.
	EphemeralContainers []EphemeralContainer `pulumi:"ephemeralContainers"`

	// HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts file
	// if specified. This is only valid for non-hostNetwork pods.
	HostAliases []HostAlias `pulumi:"hostAliases"`

	// Use the host's ipc namespace. Optional: Default to false.
	HostIPC *bool `pulumi:"hostIPC"`

	// Host networking requested for this pod. Use the host's network namespace. If this option is set,
	// the ports that will be used must be specified. Default to false.
	HostNetwork *bool `pulumi:"hostNetwork"`

	// Use the host's pid namespace. Optional: Default to false.
	HostPID *bool `pulumi:"hostPID"`

	// Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a
	// system-defined value.
	Hostname *string `pulumi:"hostname"`

	// ImagePullSecrets is an optional list of references to secrets in the same namespace to use for
	// pulling any of the images used by this PodSpec. If specified, these secrets will be passed to
	// individual puller implementations for them to use. For example, in the case of docker, only
	// DockerConfig type secrets are honored. More info:
	// https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod
	ImagePullSecrets []LocalObjectReference `pulumi:"imagePullSecrets"`

	// List of initialization containers belonging to the pod. Init containers are executed in order
	// prior to containers being started. If any init container fails, the pod is considered to have
	// failed and is handled according to its restartPolicy. The name for an init container or normal
	// container must be unique among all containers. Init containers may not have Lifecycle actions,
	// Readiness probes, Liveness probes, or Startup probes. The resourceRequirements of an init
	// container are taken into account during scheduling by finding the highest request/limit for each
	// resource type, and then using the max of of that value or the sum of the normal containers.
	// Limits are applied to init containers in a similar fashion. Init containers cannot currently be
	// added or removed. Cannot be updated. More info:
	// https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
	InitContainers []Container `pulumi:"initContainers"`

	// NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the
	// scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.
	NodeName *string `pulumi:"nodeName"`

	// NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must
	// match a node's labels for the pod to be scheduled on that node. More info:
	// https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	NodeSelector map[string]string `pulumi:"nodeSelector"`

	// Overhead represents the resource overhead associated with running a pod for a given
	// RuntimeClass. This field will be autopopulated at admission time by the RuntimeClass admission
	// controller. If the RuntimeClass admission controller is enabled, overhead must not be set in Pod
	// create requests. The RuntimeClass admission controller will reject Pod create requests which
	// have the overhead already set. If RuntimeClass is configured and selected in the PodSpec,
	// Overhead will be set to the value defined in the corresponding RuntimeClass, otherwise it will
	// remain unset and treated as zero. More info:
	// https://git.k8s.io/enhancements/keps/sig-node/20190226-pod-overhead.md This field is alpha-level
	// as of Kubernetes v1.16, and is only honored by servers that enable the PodOverhead feature.
	Overhead map[string]string `pulumi:"overhead"`

	// PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never,
	// PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is alpha-level and
	// is only honored by servers that enable the NonPreemptingPriority feature.
	PreemptionPolicy *string `pulumi:"preemptionPolicy"`

	// The priority value. Various system components use this field to find the priority of the pod.
	// When Priority Admission Controller is enabled, it prevents users from setting this field. The
	// admission controller populates this field from PriorityClassName. The higher the value, the
	// higher the priority.
	Priority *int `pulumi:"priority"`

	// If specified, indicates the pod's priority. "system-node-critical" and "system-cluster-critical"
	// are two special keywords which indicate the highest priorities with the former being the highest
	// priority. Any other name must be defined by creating a PriorityClass object with that name. If
	// not specified, the pod priority will be default or zero if there is no default.
	PriorityClassName *string `pulumi:"priorityClassName"`

	// If specified, all readiness gates will be evaluated for pod readiness. A pod is ready when all
	// its containers are ready AND all conditions specified in the readiness gates have status equal
	// to "True" More info: https://git.k8s.io/enhancements/keps/sig-network/0007-pod-ready%2B%2B.md
	ReadinessGates []PodReadinessGate `pulumi:"readinessGates"`

	// Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to
	// Always. More info:
	// https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	RestartPolicy *string `pulumi:"restartPolicy"`

	// RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used
	// to run this pod.  If no RuntimeClass resource matches the named class, the pod will not be run.
	// If unset or empty, the "legacy" RuntimeClass will be used, which is an implicit class with an
	// empty definition that uses the default runtime handler. More info:
	// https://git.k8s.io/enhancements/keps/sig-node/runtime-class.md This is a beta feature as of
	// Kubernetes v1.14.
	RuntimeClassName *string `pulumi:"runtimeClassName"`

	// If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will
	// be dispatched by default scheduler.
	SchedulerName *string `pulumi:"schedulerName"`

	// SecurityContext holds pod-level security attributes and common container settings. Optional:
	// Defaults to empty.  See type description for default values of each field.
	SecurityContext *PodSecurityContext `pulumi:"securityContext"`

	// DeprecatedServiceAccount is a depreciated alias for ServiceAccountName. Deprecated: Use
	// serviceAccountName instead.
	ServiceAccount *string `pulumi:"serviceAccount"`

	// ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info:
	// https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	ServiceAccountName *string `pulumi:"serviceAccountName"`

	// Share a single process namespace between all of the containers in a pod. When this is set
	// containers will be able to view and signal processes from other containers in the same pod, and
	// the first process in each container will not be assigned PID 1. HostPID and
	// ShareProcessNamespace cannot both be set. Optional: Default to false. This field is beta-level
	// and may be disabled with the PodShareProcessNamespace feature.
	ShareProcessNamespace *bool `pulumi:"shareProcessNamespace"`

	// If specified, the fully qualified Pod hostname will be "<hostname>.<subdomain>.<pod
	// namespace>.svc.<cluster domain>". If not specified, the pod will not have a domainname at all.
	Subdomain *string `pulumi:"subdomain"`

	// Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete
	// request. Value must be non-negative integer. The value zero indicates delete immediately. If
	// this value is nil, the default grace period will be used instead. The grace period is the
	// duration in seconds after the processes running in the pod are sent a termination signal and the
	// time when the processes are forcibly halted with a kill signal. Set this value longer than the
	// expected cleanup time for your process. Defaults to 30 seconds.
	TerminationGracePeriodSeconds *int `pulumi:"terminationGracePeriodSeconds"`

	// If specified, the pod's tolerations.
	Tolerations []Toleration `pulumi:"tolerations"`

	// TopologySpreadConstraints describes how a group of pods ought to spread across topology domains.
	// Scheduler will schedule pods in a way which abides by the constraints. This field is alpha-level
	// and is only honored by clusters that enables the EvenPodsSpread feature. All
	// topologySpreadConstraints are ANDed.
	TopologySpreadConstraints []TopologySpreadConstraint `pulumi:"topologySpreadConstraints"`

	// List of volumes that can be mounted by containers belonging to the pod. More info:
	// https://kubernetes.io/docs/concepts/storage/volumes
	Volumes []Volume `pulumi:"volumes"`

}

var _PodSpecType = reflect.TypeOf((*PodSpec)(nil)).Elem()

// PodSpecInput represents an input type that resolves to a PodSpec.
type PodSpecInput interface {
	ElementType() reflect.Type

	ToPodSpecOutput() PodSpecOutput
	ToPodSpecOutputWithContext(ctx context.Context) PodSpecOutput
}

// PodSpecArgs is a PodSpecInput whose fields are all Input types.
type PodSpecArgs struct {
	// List of containers belonging to the pod. Containers cannot currently be added or removed. There
	// must be at least one container in a Pod. Cannot be updated.
	Containers ContainerArrayInput `pulumi:"containers"`

	// Optional duration in seconds the pod may be active on the node relative to StartTime before the
	// system will actively try to mark it failed and kill associated containers. Value must be a
	// positive integer.
	ActiveDeadlineSeconds pulumi.IntInput `pulumi:"activeDeadlineSeconds"`

	// If specified, the pod's scheduling constraints
	Affinity AffinityInput `pulumi:"affinity"`

	// AutomountServiceAccountToken indicates whether a service account token should be automatically
	// mounted.
	AutomountServiceAccountToken pulumi.BoolInput `pulumi:"automountServiceAccountToken"`

	// Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated
	// DNS configuration based on DNSPolicy.
	DnsConfig PodDNSConfigInput `pulumi:"dnsConfig"`

	// Set DNS policy for the pod. Defaults to "ClusterFirst". Valid values are
	// 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in
	// DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along
	// with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.
	DnsPolicy pulumi.StringInput `pulumi:"dnsPolicy"`

	// EnableServiceLinks indicates whether information about services should be injected into pod's
	// environment variables, matching the syntax of Docker links. Optional: Defaults to true.
	EnableServiceLinks pulumi.BoolInput `pulumi:"enableServiceLinks"`

	// List of ephemeral containers run in this pod. Ephemeral containers may be run in an existing pod
	// to perform user-initiated actions such as debugging. This list cannot be specified when creating
	// a pod, and it cannot be modified by updating the pod spec. In order to add an ephemeral
	// container to an existing pod, use the pod's ephemeralcontainers subresource. This field is
	// alpha-level and is only honored by servers that enable the EphemeralContainers feature.
	EphemeralContainers EphemeralContainerArrayInput `pulumi:"ephemeralContainers"`

	// HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts file
	// if specified. This is only valid for non-hostNetwork pods.
	HostAliases HostAliasArrayInput `pulumi:"hostAliases"`

	// Use the host's ipc namespace. Optional: Default to false.
	HostIPC pulumi.BoolInput `pulumi:"hostIPC"`

	// Host networking requested for this pod. Use the host's network namespace. If this option is set,
	// the ports that will be used must be specified. Default to false.
	HostNetwork pulumi.BoolInput `pulumi:"hostNetwork"`

	// Use the host's pid namespace. Optional: Default to false.
	HostPID pulumi.BoolInput `pulumi:"hostPID"`

	// Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a
	// system-defined value.
	Hostname pulumi.StringInput `pulumi:"hostname"`

	// ImagePullSecrets is an optional list of references to secrets in the same namespace to use for
	// pulling any of the images used by this PodSpec. If specified, these secrets will be passed to
	// individual puller implementations for them to use. For example, in the case of docker, only
	// DockerConfig type secrets are honored. More info:
	// https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod
	ImagePullSecrets LocalObjectReferenceArrayInput `pulumi:"imagePullSecrets"`

	// List of initialization containers belonging to the pod. Init containers are executed in order
	// prior to containers being started. If any init container fails, the pod is considered to have
	// failed and is handled according to its restartPolicy. The name for an init container or normal
	// container must be unique among all containers. Init containers may not have Lifecycle actions,
	// Readiness probes, Liveness probes, or Startup probes. The resourceRequirements of an init
	// container are taken into account during scheduling by finding the highest request/limit for each
	// resource type, and then using the max of of that value or the sum of the normal containers.
	// Limits are applied to init containers in a similar fashion. Init containers cannot currently be
	// added or removed. Cannot be updated. More info:
	// https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
	InitContainers ContainerArrayInput `pulumi:"initContainers"`

	// NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the
	// scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.
	NodeName pulumi.StringInput `pulumi:"nodeName"`

	// NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must
	// match a node's labels for the pod to be scheduled on that node. More info:
	// https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	NodeSelector pulumi.StringMapInput `pulumi:"nodeSelector"`

	// Overhead represents the resource overhead associated with running a pod for a given
	// RuntimeClass. This field will be autopopulated at admission time by the RuntimeClass admission
	// controller. If the RuntimeClass admission controller is enabled, overhead must not be set in Pod
	// create requests. The RuntimeClass admission controller will reject Pod create requests which
	// have the overhead already set. If RuntimeClass is configured and selected in the PodSpec,
	// Overhead will be set to the value defined in the corresponding RuntimeClass, otherwise it will
	// remain unset and treated as zero. More info:
	// https://git.k8s.io/enhancements/keps/sig-node/20190226-pod-overhead.md This field is alpha-level
	// as of Kubernetes v1.16, and is only honored by servers that enable the PodOverhead feature.
	Overhead pulumi.StringMapInput `pulumi:"overhead"`

	// PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never,
	// PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is alpha-level and
	// is only honored by servers that enable the NonPreemptingPriority feature.
	PreemptionPolicy pulumi.StringInput `pulumi:"preemptionPolicy"`

	// The priority value. Various system components use this field to find the priority of the pod.
	// When Priority Admission Controller is enabled, it prevents users from setting this field. The
	// admission controller populates this field from PriorityClassName. The higher the value, the
	// higher the priority.
	Priority pulumi.IntInput `pulumi:"priority"`

	// If specified, indicates the pod's priority. "system-node-critical" and "system-cluster-critical"
	// are two special keywords which indicate the highest priorities with the former being the highest
	// priority. Any other name must be defined by creating a PriorityClass object with that name. If
	// not specified, the pod priority will be default or zero if there is no default.
	PriorityClassName pulumi.StringInput `pulumi:"priorityClassName"`

	// If specified, all readiness gates will be evaluated for pod readiness. A pod is ready when all
	// its containers are ready AND all conditions specified in the readiness gates have status equal
	// to "True" More info: https://git.k8s.io/enhancements/keps/sig-network/0007-pod-ready%2B%2B.md
	ReadinessGates PodReadinessGateArrayInput `pulumi:"readinessGates"`

	// Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to
	// Always. More info:
	// https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	RestartPolicy pulumi.StringInput `pulumi:"restartPolicy"`

	// RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used
	// to run this pod.  If no RuntimeClass resource matches the named class, the pod will not be run.
	// If unset or empty, the "legacy" RuntimeClass will be used, which is an implicit class with an
	// empty definition that uses the default runtime handler. More info:
	// https://git.k8s.io/enhancements/keps/sig-node/runtime-class.md This is a beta feature as of
	// Kubernetes v1.14.
	RuntimeClassName pulumi.StringInput `pulumi:"runtimeClassName"`

	// If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will
	// be dispatched by default scheduler.
	SchedulerName pulumi.StringInput `pulumi:"schedulerName"`

	// SecurityContext holds pod-level security attributes and common container settings. Optional:
	// Defaults to empty.  See type description for default values of each field.
	SecurityContext PodSecurityContextInput `pulumi:"securityContext"`

	// DeprecatedServiceAccount is a depreciated alias for ServiceAccountName. Deprecated: Use
	// serviceAccountName instead.
	ServiceAccount pulumi.StringInput `pulumi:"serviceAccount"`

	// ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info:
	// https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	ServiceAccountName pulumi.StringInput `pulumi:"serviceAccountName"`

	// Share a single process namespace between all of the containers in a pod. When this is set
	// containers will be able to view and signal processes from other containers in the same pod, and
	// the first process in each container will not be assigned PID 1. HostPID and
	// ShareProcessNamespace cannot both be set. Optional: Default to false. This field is beta-level
	// and may be disabled with the PodShareProcessNamespace feature.
	ShareProcessNamespace pulumi.BoolInput `pulumi:"shareProcessNamespace"`

	// If specified, the fully qualified Pod hostname will be "<hostname>.<subdomain>.<pod
	// namespace>.svc.<cluster domain>". If not specified, the pod will not have a domainname at all.
	Subdomain pulumi.StringInput `pulumi:"subdomain"`

	// Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete
	// request. Value must be non-negative integer. The value zero indicates delete immediately. If
	// this value is nil, the default grace period will be used instead. The grace period is the
	// duration in seconds after the processes running in the pod are sent a termination signal and the
	// time when the processes are forcibly halted with a kill signal. Set this value longer than the
	// expected cleanup time for your process. Defaults to 30 seconds.
	TerminationGracePeriodSeconds pulumi.IntInput `pulumi:"terminationGracePeriodSeconds"`

	// If specified, the pod's tolerations.
	Tolerations TolerationArrayInput `pulumi:"tolerations"`

	// TopologySpreadConstraints describes how a group of pods ought to spread across topology domains.
	// Scheduler will schedule pods in a way which abides by the constraints. This field is alpha-level
	// and is only honored by clusters that enables the EvenPodsSpread feature. All
	// topologySpreadConstraints are ANDed.
	TopologySpreadConstraints TopologySpreadConstraintArrayInput `pulumi:"topologySpreadConstraints"`

	// List of volumes that can be mounted by containers belonging to the pod. More info:
	// https://kubernetes.io/docs/concepts/storage/volumes
	Volumes VolumeArrayInput `pulumi:"volumes"`

}

func (a PodSpecArgs) ElementType() reflect.Type {
	return _PodSpecType
}

func (a PodSpecArgs) ToPodSpecOutput() PodSpecOutput {
	return pulumi.ToOutput(a).(PodSpecOutput)
}

func (a PodSpecArgs) ToPodSpecOutputWithContext(ctx context.Context) PodSpecOutput {
	return pulumi.ToOutputWithContext(ctx, a).(PodSpecOutput)
}

// PodSpecOutput is an output type that resolves to a Input.
type PodSpecOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(PodSpecOutput{}) }

func (PodSpecOutput) ElementType() reflect.Type {
	return _PodSpecType
}

func (o PodSpecOutput) ActiveDeadlineSeconds() pulumi.IntOutput {
	return o.Apply(func(v PodSpec) *int {
		return v.ActiveDeadlineSeconds
	}).(pulumi.IntOutput)
}

func (o PodSpecOutput) Affinity() AffinityOutput {
	return o.Apply(func(v PodSpec) *Affinity {
		return v.Affinity
	}).(AffinityOutput)
}

func (o PodSpecOutput) AutomountServiceAccountToken() pulumi.BoolOutput {
	return o.Apply(func(v PodSpec) *bool {
		return v.AutomountServiceAccountToken
	}).(pulumi.BoolOutput)
}

func (o PodSpecOutput) Containers() ContainerArrayOutput {
	return o.Apply(func(v PodSpec) []Container {
		return v.Containers
	}).(ContainerArrayOutput)
}

func (o PodSpecOutput) DnsConfig() PodDNSConfigOutput {
	return o.Apply(func(v PodSpec) *PodDNSConfig {
		return v.DnsConfig
	}).(PodDNSConfigOutput)
}

func (o PodSpecOutput) DnsPolicy() pulumi.StringOutput {
	return o.Apply(func(v PodSpec) *string {
		return v.DnsPolicy
	}).(pulumi.StringOutput)
}

func (o PodSpecOutput) EnableServiceLinks() pulumi.BoolOutput {
	return o.Apply(func(v PodSpec) *bool {
		return v.EnableServiceLinks
	}).(pulumi.BoolOutput)
}

func (o PodSpecOutput) EphemeralContainers() EphemeralContainerArrayOutput {
	return o.Apply(func(v PodSpec) []EphemeralContainer {
		return v.EphemeralContainers
	}).(EphemeralContainerArrayOutput)
}

func (o PodSpecOutput) HostAliases() HostAliasArrayOutput {
	return o.Apply(func(v PodSpec) []HostAlias {
		return v.HostAliases
	}).(HostAliasArrayOutput)
}

func (o PodSpecOutput) HostIPC() pulumi.BoolOutput {
	return o.Apply(func(v PodSpec) *bool {
		return v.HostIPC
	}).(pulumi.BoolOutput)
}

func (o PodSpecOutput) HostNetwork() pulumi.BoolOutput {
	return o.Apply(func(v PodSpec) *bool {
		return v.HostNetwork
	}).(pulumi.BoolOutput)
}

func (o PodSpecOutput) HostPID() pulumi.BoolOutput {
	return o.Apply(func(v PodSpec) *bool {
		return v.HostPID
	}).(pulumi.BoolOutput)
}

func (o PodSpecOutput) Hostname() pulumi.StringOutput {
	return o.Apply(func(v PodSpec) *string {
		return v.Hostname
	}).(pulumi.StringOutput)
}

func (o PodSpecOutput) ImagePullSecrets() LocalObjectReferenceArrayOutput {
	return o.Apply(func(v PodSpec) []LocalObjectReference {
		return v.ImagePullSecrets
	}).(LocalObjectReferenceArrayOutput)
}

func (o PodSpecOutput) InitContainers() ContainerArrayOutput {
	return o.Apply(func(v PodSpec) []Container {
		return v.InitContainers
	}).(ContainerArrayOutput)
}

func (o PodSpecOutput) NodeName() pulumi.StringOutput {
	return o.Apply(func(v PodSpec) *string {
		return v.NodeName
	}).(pulumi.StringOutput)
}

func (o PodSpecOutput) NodeSelector() pulumi.StringMapOutput {
	return o.Apply(func(v PodSpec) map[string]string {
		return v.NodeSelector
	}).(pulumi.StringMapOutput)
}

func (o PodSpecOutput) Overhead() pulumi.StringMapOutput {
	return o.Apply(func(v PodSpec) map[string]string {
		return v.Overhead
	}).(pulumi.StringMapOutput)
}

func (o PodSpecOutput) PreemptionPolicy() pulumi.StringOutput {
	return o.Apply(func(v PodSpec) *string {
		return v.PreemptionPolicy
	}).(pulumi.StringOutput)
}

func (o PodSpecOutput) Priority() pulumi.IntOutput {
	return o.Apply(func(v PodSpec) *int {
		return v.Priority
	}).(pulumi.IntOutput)
}

func (o PodSpecOutput) PriorityClassName() pulumi.StringOutput {
	return o.Apply(func(v PodSpec) *string {
		return v.PriorityClassName
	}).(pulumi.StringOutput)
}

func (o PodSpecOutput) ReadinessGates() PodReadinessGateArrayOutput {
	return o.Apply(func(v PodSpec) []PodReadinessGate {
		return v.ReadinessGates
	}).(PodReadinessGateArrayOutput)
}

func (o PodSpecOutput) RestartPolicy() pulumi.StringOutput {
	return o.Apply(func(v PodSpec) *string {
		return v.RestartPolicy
	}).(pulumi.StringOutput)
}

func (o PodSpecOutput) RuntimeClassName() pulumi.StringOutput {
	return o.Apply(func(v PodSpec) *string {
		return v.RuntimeClassName
	}).(pulumi.StringOutput)
}

func (o PodSpecOutput) SchedulerName() pulumi.StringOutput {
	return o.Apply(func(v PodSpec) *string {
		return v.SchedulerName
	}).(pulumi.StringOutput)
}

func (o PodSpecOutput) SecurityContext() PodSecurityContextOutput {
	return o.Apply(func(v PodSpec) *PodSecurityContext {
		return v.SecurityContext
	}).(PodSecurityContextOutput)
}

func (o PodSpecOutput) ServiceAccount() pulumi.StringOutput {
	return o.Apply(func(v PodSpec) *string {
		return v.ServiceAccount
	}).(pulumi.StringOutput)
}

func (o PodSpecOutput) ServiceAccountName() pulumi.StringOutput {
	return o.Apply(func(v PodSpec) *string {
		return v.ServiceAccountName
	}).(pulumi.StringOutput)
}

func (o PodSpecOutput) ShareProcessNamespace() pulumi.BoolOutput {
	return o.Apply(func(v PodSpec) *bool {
		return v.ShareProcessNamespace
	}).(pulumi.BoolOutput)
}

func (o PodSpecOutput) Subdomain() pulumi.StringOutput {
	return o.Apply(func(v PodSpec) *string {
		return v.Subdomain
	}).(pulumi.StringOutput)
}

func (o PodSpecOutput) TerminationGracePeriodSeconds() pulumi.IntOutput {
	return o.Apply(func(v PodSpec) *int {
		return v.TerminationGracePeriodSeconds
	}).(pulumi.IntOutput)
}

func (o PodSpecOutput) Tolerations() TolerationArrayOutput {
	return o.Apply(func(v PodSpec) []Toleration {
		return v.Tolerations
	}).(TolerationArrayOutput)
}

func (o PodSpecOutput) TopologySpreadConstraints() TopologySpreadConstraintArrayOutput {
	return o.Apply(func(v PodSpec) []TopologySpreadConstraint {
		return v.TopologySpreadConstraints
	}).(TopologySpreadConstraintArrayOutput)
}

func (o PodSpecOutput) Volumes() VolumeArrayOutput {
	return o.Apply(func(v PodSpec) []Volume {
		return v.Volumes
	}).(VolumeArrayOutput)
}

