// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// CSIDriverSpec is the specification of a CSIDriver.
type CSIDriverSpec struct {
	// attachRequired indicates this CSI volume driver requires an attach operation (because it
	// implements the CSI ControllerPublishVolume() method), and that the Kubernetes attach detach
	// controller should call the attach volume interface which checks the volumeattachment status and
	// waits until the volume is attached before proceeding to mounting. The CSI external-attacher
	// coordinates with CSI volume driver and updates the volumeattachment status when the attach
	// operation is complete. If the CSIDriverRegistry feature gate is enabled and the value is
	// specified to false, the attach operation will be skipped. Otherwise the attach operation will be
	// called.
	AttachRequired *bool `pulumi:"attachRequired"`

	// If set to true, podInfoOnMount indicates this CSI volume driver requires additional pod
	// information (like podName, podUID, etc.) during mount operations. If set to false, pod
	// information will not be passed on mount. Default is false. The CSI driver specifies
	// podInfoOnMount as part of driver deployment. If true, Kubelet will pass pod information as
	// VolumeContext in the CSI NodePublishVolume() calls. The CSI driver is responsible for parsing
	// and validating the information passed in as VolumeContext. The following VolumeConext will be
	// passed if podInfoOnMount is set to true. This list might grow, but the prefix will be used.
	// "csi.storage.k8s.io/pod.name": pod.Name "csi.storage.k8s.io/pod.namespace": pod.Namespace
	// "csi.storage.k8s.io/pod.uid": string(pod.UID) "csi.storage.k8s.io/ephemeral": "true" iff the
	// volume is an ephemeral inline volume
	//                                 defined by a CSIVolumeSource, otherwise "false"
	// 
	// "csi.storage.k8s.io/ephemeral" is a new feature in Kubernetes 1.16. It is only required for
	// drivers which support both the "Persistent" and "Ephemeral" VolumeLifecycleMode. Other drivers
	// can leave pod info disabled and/or ignore this field. As Kubernetes 1.15 doesn't support this
	// field, drivers can only support one mode when deployed on such a cluster and the deployment
	// determines which mode that is, for example via a command line parameter of the driver.
	PodInfoOnMount *bool `pulumi:"podInfoOnMount"`

	// VolumeLifecycleModes defines what kind of volumes this CSI volume driver supports. The default
	// if the list is empty is "Persistent", which is the usage defined by the CSI specification and
	// implemented in Kubernetes via the usual PV/PVC mechanism. The other mode is "Ephemeral". In this
	// mode, volumes are defined inline inside the pod spec with CSIVolumeSource and their lifecycle is
	// tied to the lifecycle of that pod. A driver has to be aware of this because it is only going to
	// get a NodePublishVolume call for such a volume. For more information about implementing this
	// mode, see https://kubernetes-csi.github.io/docs/ephemeral-local-volumes.html A driver can
	// support one or more of these modes and more modes may be added in the future.
	VolumeLifecycleModes []string `pulumi:"volumeLifecycleModes"`

}

var _CSIDriverSpecType = reflect.TypeOf((*CSIDriverSpec)(nil)).Elem()

// CSIDriverSpecInput represents an input type that resolves to a CSIDriverSpec.
type CSIDriverSpecInput interface {
	ElementType() reflect.Type

	ToCSIDriverSpecOutput() CSIDriverSpecOutput
	ToCSIDriverSpecOutputWithContext(ctx context.Context) CSIDriverSpecOutput
}

// CSIDriverSpecArgs is a CSIDriverSpecInput whose fields are all Input types.
type CSIDriverSpecArgs struct {
	// attachRequired indicates this CSI volume driver requires an attach operation (because it
	// implements the CSI ControllerPublishVolume() method), and that the Kubernetes attach detach
	// controller should call the attach volume interface which checks the volumeattachment status and
	// waits until the volume is attached before proceeding to mounting. The CSI external-attacher
	// coordinates with CSI volume driver and updates the volumeattachment status when the attach
	// operation is complete. If the CSIDriverRegistry feature gate is enabled and the value is
	// specified to false, the attach operation will be skipped. Otherwise the attach operation will be
	// called.
	AttachRequired pulumi.BoolInput `pulumi:"attachRequired"`

	// If set to true, podInfoOnMount indicates this CSI volume driver requires additional pod
	// information (like podName, podUID, etc.) during mount operations. If set to false, pod
	// information will not be passed on mount. Default is false. The CSI driver specifies
	// podInfoOnMount as part of driver deployment. If true, Kubelet will pass pod information as
	// VolumeContext in the CSI NodePublishVolume() calls. The CSI driver is responsible for parsing
	// and validating the information passed in as VolumeContext. The following VolumeConext will be
	// passed if podInfoOnMount is set to true. This list might grow, but the prefix will be used.
	// "csi.storage.k8s.io/pod.name": pod.Name "csi.storage.k8s.io/pod.namespace": pod.Namespace
	// "csi.storage.k8s.io/pod.uid": string(pod.UID) "csi.storage.k8s.io/ephemeral": "true" iff the
	// volume is an ephemeral inline volume
	//                                 defined by a CSIVolumeSource, otherwise "false"
	// 
	// "csi.storage.k8s.io/ephemeral" is a new feature in Kubernetes 1.16. It is only required for
	// drivers which support both the "Persistent" and "Ephemeral" VolumeLifecycleMode. Other drivers
	// can leave pod info disabled and/or ignore this field. As Kubernetes 1.15 doesn't support this
	// field, drivers can only support one mode when deployed on such a cluster and the deployment
	// determines which mode that is, for example via a command line parameter of the driver.
	PodInfoOnMount pulumi.BoolInput `pulumi:"podInfoOnMount"`

	// VolumeLifecycleModes defines what kind of volumes this CSI volume driver supports. The default
	// if the list is empty is "Persistent", which is the usage defined by the CSI specification and
	// implemented in Kubernetes via the usual PV/PVC mechanism. The other mode is "Ephemeral". In this
	// mode, volumes are defined inline inside the pod spec with CSIVolumeSource and their lifecycle is
	// tied to the lifecycle of that pod. A driver has to be aware of this because it is only going to
	// get a NodePublishVolume call for such a volume. For more information about implementing this
	// mode, see https://kubernetes-csi.github.io/docs/ephemeral-local-volumes.html A driver can
	// support one or more of these modes and more modes may be added in the future.
	VolumeLifecycleModes pulumi.StringArrayInput `pulumi:"volumeLifecycleModes"`

}

func (a CSIDriverSpecArgs) ElementType() reflect.Type {
	return _CSIDriverSpecType
}

func (a CSIDriverSpecArgs) ToCSIDriverSpecOutput() CSIDriverSpecOutput {
	return pulumi.ToOutput(a).(CSIDriverSpecOutput)
}

func (a CSIDriverSpecArgs) ToCSIDriverSpecOutputWithContext(ctx context.Context) CSIDriverSpecOutput {
	return pulumi.ToOutputWithContext(ctx, a).(CSIDriverSpecOutput)
}

// CSIDriverSpecOutput is an output type that resolves to a Input.
type CSIDriverSpecOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(CSIDriverSpecOutput{}) }

func (CSIDriverSpecOutput) ElementType() reflect.Type {
	return _CSIDriverSpecType
}

func (o CSIDriverSpecOutput) AttachRequired() pulumi.BoolOutput {
	return o.Apply(func(v CSIDriverSpec) *bool {
		return v.AttachRequired
	}).(pulumi.BoolOutput)
}

func (o CSIDriverSpecOutput) PodInfoOnMount() pulumi.BoolOutput {
	return o.Apply(func(v CSIDriverSpec) *bool {
		return v.PodInfoOnMount
	}).(pulumi.BoolOutput)
}

func (o CSIDriverSpecOutput) VolumeLifecycleModes() pulumi.StringArrayOutput {
	return o.Apply(func(v CSIDriverSpec) []string {
		return v.VolumeLifecycleModes
	}).(pulumi.StringArrayOutput)
}
