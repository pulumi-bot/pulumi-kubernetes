// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support
// ownership management and SELinux relabeling.
type RBDPersistentVolumeSource struct {
	// Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is
	// supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to
	// be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
	FsType *string `pulumi:"fsType"`

	// The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Image string `pulumi:"image"`

	// Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info:
	// https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Keyring *string `pulumi:"keyring"`

	// A collection of Ceph monitors. More info:
	// https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Monitors []string `pulumi:"monitors"`

	// The rados pool name. Default is rbd. More info:
	// https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Pool *string `pulumi:"pool"`

	// ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info:
	// https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	ReadOnly *bool `pulumi:"readOnly"`

	// SecretRef is name of the authentication secret for RBDUser. If provided overrides keyring.
	// Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	SecretRef *SecretReference `pulumi:"secretRef"`

	// The rados user name. Default is admin. More info:
	// https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	User *string `pulumi:"user"`

}

var _RBDPersistentVolumeSourceType = reflect.TypeOf((*RBDPersistentVolumeSource)(nil)).Elem()

// RBDPersistentVolumeSourceInput represents an input type that resolves to a RBDPersistentVolumeSource.
type RBDPersistentVolumeSourceInput interface {
	ElementType() reflect.Type

	ToRBDPersistentVolumeSourceOutput() RBDPersistentVolumeSourceOutput
	ToRBDPersistentVolumeSourceOutputWithContext(ctx context.Context) RBDPersistentVolumeSourceOutput
}

// RBDPersistentVolumeSourceArgs is a RBDPersistentVolumeSourceInput whose fields are all Input types.
type RBDPersistentVolumeSourceArgs struct {
	// The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Image pulumi.StringInput `pulumi:"image"`

	// A collection of Ceph monitors. More info:
	// https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Monitors pulumi.StringArrayInput `pulumi:"monitors"`

	// Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is
	// supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to
	// be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
	FsType pulumi.StringInput `pulumi:"fsType"`

	// Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info:
	// https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Keyring pulumi.StringInput `pulumi:"keyring"`

	// The rados pool name. Default is rbd. More info:
	// https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	Pool pulumi.StringInput `pulumi:"pool"`

	// ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info:
	// https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	ReadOnly pulumi.BoolInput `pulumi:"readOnly"`

	// SecretRef is name of the authentication secret for RBDUser. If provided overrides keyring.
	// Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	SecretRef SecretReferenceInput `pulumi:"secretRef"`

	// The rados user name. Default is admin. More info:
	// https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
	User pulumi.StringInput `pulumi:"user"`

}

func (a RBDPersistentVolumeSourceArgs) ElementType() reflect.Type {
	return _RBDPersistentVolumeSourceType
}

func (a RBDPersistentVolumeSourceArgs) ToRBDPersistentVolumeSourceOutput() RBDPersistentVolumeSourceOutput {
	return pulumi.ToOutput(a).(RBDPersistentVolumeSourceOutput)
}

func (a RBDPersistentVolumeSourceArgs) ToRBDPersistentVolumeSourceOutputWithContext(ctx context.Context) RBDPersistentVolumeSourceOutput {
	return pulumi.ToOutputWithContext(ctx, a).(RBDPersistentVolumeSourceOutput)
}

// RBDPersistentVolumeSourceOutput is an output type that resolves to a Input.
type RBDPersistentVolumeSourceOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(RBDPersistentVolumeSourceOutput{}) }

func (RBDPersistentVolumeSourceOutput) ElementType() reflect.Type {
	return _RBDPersistentVolumeSourceType
}

func (o RBDPersistentVolumeSourceOutput) FsType() pulumi.StringOutput {
	return o.Apply(func(v RBDPersistentVolumeSource) *string {
		return v.FsType
	}).(pulumi.StringOutput)
}

func (o RBDPersistentVolumeSourceOutput) Image() pulumi.StringOutput {
	return o.Apply(func(v RBDPersistentVolumeSource) string {
		return v.Image
	}).(pulumi.StringOutput)
}

func (o RBDPersistentVolumeSourceOutput) Keyring() pulumi.StringOutput {
	return o.Apply(func(v RBDPersistentVolumeSource) *string {
		return v.Keyring
	}).(pulumi.StringOutput)
}

func (o RBDPersistentVolumeSourceOutput) Monitors() pulumi.StringArrayOutput {
	return o.Apply(func(v RBDPersistentVolumeSource) []string {
		return v.Monitors
	}).(pulumi.StringArrayOutput)
}

func (o RBDPersistentVolumeSourceOutput) Pool() pulumi.StringOutput {
	return o.Apply(func(v RBDPersistentVolumeSource) *string {
		return v.Pool
	}).(pulumi.StringOutput)
}

func (o RBDPersistentVolumeSourceOutput) ReadOnly() pulumi.BoolOutput {
	return o.Apply(func(v RBDPersistentVolumeSource) *bool {
		return v.ReadOnly
	}).(pulumi.BoolOutput)
}

func (o RBDPersistentVolumeSourceOutput) SecretRef() SecretReferenceOutput {
	return o.Apply(func(v RBDPersistentVolumeSource) *SecretReference {
		return v.SecretRef
	}).(SecretReferenceOutput)
}

func (o RBDPersistentVolumeSourceOutput) User() pulumi.StringOutput {
	return o.Apply(func(v RBDPersistentVolumeSource) *string {
		return v.User
	}).(pulumi.StringOutput)
}

