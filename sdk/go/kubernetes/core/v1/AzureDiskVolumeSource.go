// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
type AzureDiskVolumeSource struct {
	// Host Caching mode: None, Read Only, Read Write.
	CachingMode *string `pulumi:"cachingMode"`

	// The Name of the data disk in the blob storage
	DiskName string `pulumi:"diskName"`

	// The URI the data disk in the blob storage
	DiskURI string `pulumi:"diskURI"`

	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex.
	// "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	FsType *string `pulumi:"fsType"`

	// Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per
	// storage account  Managed: azure managed data disk (only in managed availability set). defaults
	// to shared
	Kind *string `pulumi:"kind"`

	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `pulumi:"readOnly"`

}

var _AzureDiskVolumeSourceType = reflect.TypeOf((*AzureDiskVolumeSource)(nil)).Elem()

// AzureDiskVolumeSourceInput represents an input type that resolves to a AzureDiskVolumeSource.
type AzureDiskVolumeSourceInput interface {
	ElementType() reflect.Type

	ToAzureDiskVolumeSourceOutput() AzureDiskVolumeSourceOutput
	ToAzureDiskVolumeSourceOutputWithContext(ctx context.Context) AzureDiskVolumeSourceOutput
}

// AzureDiskVolumeSourceArgs is a AzureDiskVolumeSourceInput whose fields are all Input types.
type AzureDiskVolumeSourceArgs struct {
	// The Name of the data disk in the blob storage
	DiskName pulumi.StringInput `pulumi:"diskName"`

	// The URI the data disk in the blob storage
	DiskURI pulumi.StringInput `pulumi:"diskURI"`

	// Host Caching mode: None, Read Only, Read Write.
	CachingMode pulumi.StringInput `pulumi:"cachingMode"`

	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex.
	// "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	FsType pulumi.StringInput `pulumi:"fsType"`

	// Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per
	// storage account  Managed: azure managed data disk (only in managed availability set). defaults
	// to shared
	Kind pulumi.StringInput `pulumi:"kind"`

	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly pulumi.BoolInput `pulumi:"readOnly"`

}

func (a AzureDiskVolumeSourceArgs) ElementType() reflect.Type {
	return _AzureDiskVolumeSourceType
}

func (a AzureDiskVolumeSourceArgs) ToAzureDiskVolumeSourceOutput() AzureDiskVolumeSourceOutput {
	return pulumi.ToOutput(a).(AzureDiskVolumeSourceOutput)
}

func (a AzureDiskVolumeSourceArgs) ToAzureDiskVolumeSourceOutputWithContext(ctx context.Context) AzureDiskVolumeSourceOutput {
	return pulumi.ToOutputWithContext(ctx, a).(AzureDiskVolumeSourceOutput)
}

// AzureDiskVolumeSourceOutput is an output type that resolves to a Input.
type AzureDiskVolumeSourceOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(AzureDiskVolumeSourceOutput{}) }

func (AzureDiskVolumeSourceOutput) ElementType() reflect.Type {
	return _AzureDiskVolumeSourceType
}

func (o AzureDiskVolumeSourceOutput) CachingMode() pulumi.StringOutput {
	return o.Apply(func(v AzureDiskVolumeSource) *string {
		return v.CachingMode
	}).(pulumi.StringOutput)
}

func (o AzureDiskVolumeSourceOutput) DiskName() pulumi.StringOutput {
	return o.Apply(func(v AzureDiskVolumeSource) string {
		return v.DiskName
	}).(pulumi.StringOutput)
}

func (o AzureDiskVolumeSourceOutput) DiskURI() pulumi.StringOutput {
	return o.Apply(func(v AzureDiskVolumeSource) string {
		return v.DiskURI
	}).(pulumi.StringOutput)
}

func (o AzureDiskVolumeSourceOutput) FsType() pulumi.StringOutput {
	return o.Apply(func(v AzureDiskVolumeSource) *string {
		return v.FsType
	}).(pulumi.StringOutput)
}

func (o AzureDiskVolumeSourceOutput) Kind() pulumi.StringOutput {
	return o.Apply(func(v AzureDiskVolumeSource) *string {
		return v.Kind
	}).(pulumi.StringOutput)
}

func (o AzureDiskVolumeSourceOutput) ReadOnly() pulumi.BoolOutput {
	return o.Apply(func(v AzureDiskVolumeSource) *bool {
		return v.ReadOnly
	}).(pulumi.BoolOutput)
}

