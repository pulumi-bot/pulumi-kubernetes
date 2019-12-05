// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes
// support ownership management and SELinux relabeling.
type ISCSIVolumeSource struct {
	// whether support iSCSI Discovery CHAP authentication
	ChapAuthDiscovery *bool `pulumi:"chapAuthDiscovery"`

	// whether support iSCSI Session CHAP authentication
	ChapAuthSession *bool `pulumi:"chapAuthSession"`

	// Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is
	// supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to
	// be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
	FsType *string `pulumi:"fsType"`

	// Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously,
	// new iSCSI interface <target portal>:<volume name> will be created for the connection.
	InitiatorName *string `pulumi:"initiatorName"`

	// Target iSCSI Qualified Name.
	Iqn string `pulumi:"iqn"`

	// iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).
	IscsiInterface *string `pulumi:"iscsiInterface"`

	// iSCSI Target Lun number.
	Lun int `pulumi:"lun"`

	// iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than
	// default (typically TCP ports 860 and 3260).
	Portals []string `pulumi:"portals"`

	// ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.
	ReadOnly *bool `pulumi:"readOnly"`

	// CHAP Secret for iSCSI target and initiator authentication
	SecretRef *LocalObjectReference `pulumi:"secretRef"`

	// iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than
	// default (typically TCP ports 860 and 3260).
	TargetPortal string `pulumi:"targetPortal"`

}

var _ISCSIVolumeSourceType = reflect.TypeOf((*ISCSIVolumeSource)(nil)).Elem()

// ISCSIVolumeSourceInput represents an input type that resolves to a ISCSIVolumeSource.
type ISCSIVolumeSourceInput interface {
	ElementType() reflect.Type

	ToISCSIVolumeSourceOutput() ISCSIVolumeSourceOutput
	ToISCSIVolumeSourceOutputWithContext(ctx context.Context) ISCSIVolumeSourceOutput
}

// ISCSIVolumeSourceArgs is a ISCSIVolumeSourceInput whose fields are all Input types.
type ISCSIVolumeSourceArgs struct {
	// Target iSCSI Qualified Name.
	Iqn pulumi.StringInput `pulumi:"iqn"`

	// iSCSI Target Lun number.
	Lun pulumi.IntInput `pulumi:"lun"`

	// iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than
	// default (typically TCP ports 860 and 3260).
	TargetPortal pulumi.StringInput `pulumi:"targetPortal"`

	// whether support iSCSI Discovery CHAP authentication
	ChapAuthDiscovery pulumi.BoolInput `pulumi:"chapAuthDiscovery"`

	// whether support iSCSI Session CHAP authentication
	ChapAuthSession pulumi.BoolInput `pulumi:"chapAuthSession"`

	// Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is
	// supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to
	// be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
	FsType pulumi.StringInput `pulumi:"fsType"`

	// Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously,
	// new iSCSI interface <target portal>:<volume name> will be created for the connection.
	InitiatorName pulumi.StringInput `pulumi:"initiatorName"`

	// iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).
	IscsiInterface pulumi.StringInput `pulumi:"iscsiInterface"`

	// iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than
	// default (typically TCP ports 860 and 3260).
	Portals pulumi.StringArrayInput `pulumi:"portals"`

	// ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.
	ReadOnly pulumi.BoolInput `pulumi:"readOnly"`

	// CHAP Secret for iSCSI target and initiator authentication
	SecretRef LocalObjectReferenceInput `pulumi:"secretRef"`

}

func (a ISCSIVolumeSourceArgs) ElementType() reflect.Type {
	return _ISCSIVolumeSourceType
}

func (a ISCSIVolumeSourceArgs) ToISCSIVolumeSourceOutput() ISCSIVolumeSourceOutput {
	return pulumi.ToOutput(a).(ISCSIVolumeSourceOutput)
}

func (a ISCSIVolumeSourceArgs) ToISCSIVolumeSourceOutputWithContext(ctx context.Context) ISCSIVolumeSourceOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ISCSIVolumeSourceOutput)
}

// ISCSIVolumeSourceOutput is an output type that resolves to a Input.
type ISCSIVolumeSourceOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ISCSIVolumeSourceOutput{}) }

func (ISCSIVolumeSourceOutput) ElementType() reflect.Type {
	return _ISCSIVolumeSourceType
}

func (o ISCSIVolumeSourceOutput) ChapAuthDiscovery() pulumi.BoolOutput {
	return o.Apply(func(v ISCSIVolumeSource) *bool {
		return v.ChapAuthDiscovery
	}).(pulumi.BoolOutput)
}

func (o ISCSIVolumeSourceOutput) ChapAuthSession() pulumi.BoolOutput {
	return o.Apply(func(v ISCSIVolumeSource) *bool {
		return v.ChapAuthSession
	}).(pulumi.BoolOutput)
}

func (o ISCSIVolumeSourceOutput) FsType() pulumi.StringOutput {
	return o.Apply(func(v ISCSIVolumeSource) *string {
		return v.FsType
	}).(pulumi.StringOutput)
}

func (o ISCSIVolumeSourceOutput) InitiatorName() pulumi.StringOutput {
	return o.Apply(func(v ISCSIVolumeSource) *string {
		return v.InitiatorName
	}).(pulumi.StringOutput)
}

func (o ISCSIVolumeSourceOutput) Iqn() pulumi.StringOutput {
	return o.Apply(func(v ISCSIVolumeSource) string {
		return v.Iqn
	}).(pulumi.StringOutput)
}

func (o ISCSIVolumeSourceOutput) IscsiInterface() pulumi.StringOutput {
	return o.Apply(func(v ISCSIVolumeSource) *string {
		return v.IscsiInterface
	}).(pulumi.StringOutput)
}

func (o ISCSIVolumeSourceOutput) Lun() pulumi.IntOutput {
	return o.Apply(func(v ISCSIVolumeSource) int {
		return v.Lun
	}).(pulumi.IntOutput)
}

func (o ISCSIVolumeSourceOutput) Portals() pulumi.StringArrayOutput {
	return o.Apply(func(v ISCSIVolumeSource) []string {
		return v.Portals
	}).(pulumi.StringArrayOutput)
}

func (o ISCSIVolumeSourceOutput) ReadOnly() pulumi.BoolOutput {
	return o.Apply(func(v ISCSIVolumeSource) *bool {
		return v.ReadOnly
	}).(pulumi.BoolOutput)
}

func (o ISCSIVolumeSourceOutput) SecretRef() LocalObjectReferenceOutput {
	return o.Apply(func(v ISCSIVolumeSource) *LocalObjectReference {
		return v.SecretRef
	}).(LocalObjectReferenceOutput)
}

func (o ISCSIVolumeSourceOutput) TargetPortal() pulumi.StringOutput {
	return o.Apply(func(v ISCSIVolumeSource) string {
		return v.TargetPortal
	}).(pulumi.StringOutput)
}

