// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// VolumeAttachmentStatus is the status of a VolumeAttachment request.
type VolumeAttachmentStatus struct {
	// The last error encountered during attach operation, if any. This field must only be set by the
	// entity completing the attach operation, i.e. the external-attacher.
	AttachError *VolumeError `pulumi:"attachError"`

	// Indicates the volume is successfully attached. This field must only be set by the entity
	// completing the attach operation, i.e. the external-attacher.
	Attached bool `pulumi:"attached"`

	// Upon successful attach, this field is populated with any information returned by the attach
	// operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only
	// be set by the entity completing the attach operation, i.e. the external-attacher.
	AttachmentMetadata map[string]string `pulumi:"attachmentMetadata"`

	// The last error encountered during detach operation, if any. This field must only be set by the
	// entity completing the detach operation, i.e. the external-attacher.
	DetachError *VolumeError `pulumi:"detachError"`

}

var _VolumeAttachmentStatusType = reflect.TypeOf((*VolumeAttachmentStatus)(nil)).Elem()

// VolumeAttachmentStatusInput represents an input type that resolves to a VolumeAttachmentStatus.
type VolumeAttachmentStatusInput interface {
	ElementType() reflect.Type

	ToVolumeAttachmentStatusOutput() VolumeAttachmentStatusOutput
	ToVolumeAttachmentStatusOutputWithContext(ctx context.Context) VolumeAttachmentStatusOutput
}

// VolumeAttachmentStatusArgs is a VolumeAttachmentStatusInput whose fields are all Input types.
type VolumeAttachmentStatusArgs struct {
	// Indicates the volume is successfully attached. This field must only be set by the entity
	// completing the attach operation, i.e. the external-attacher.
	Attached pulumi.BoolInput `pulumi:"attached"`

	// The last error encountered during attach operation, if any. This field must only be set by the
	// entity completing the attach operation, i.e. the external-attacher.
	AttachError VolumeErrorInput `pulumi:"attachError"`

	// Upon successful attach, this field is populated with any information returned by the attach
	// operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only
	// be set by the entity completing the attach operation, i.e. the external-attacher.
	AttachmentMetadata pulumi.StringMapInput `pulumi:"attachmentMetadata"`

	// The last error encountered during detach operation, if any. This field must only be set by the
	// entity completing the detach operation, i.e. the external-attacher.
	DetachError VolumeErrorInput `pulumi:"detachError"`

}

func (a VolumeAttachmentStatusArgs) ElementType() reflect.Type {
	return _VolumeAttachmentStatusType
}

func (a VolumeAttachmentStatusArgs) ToVolumeAttachmentStatusOutput() VolumeAttachmentStatusOutput {
	return pulumi.ToOutput(a).(VolumeAttachmentStatusOutput)
}

func (a VolumeAttachmentStatusArgs) ToVolumeAttachmentStatusOutputWithContext(ctx context.Context) VolumeAttachmentStatusOutput {
	return pulumi.ToOutputWithContext(ctx, a).(VolumeAttachmentStatusOutput)
}

// VolumeAttachmentStatusOutput is an output type that resolves to a Input.
type VolumeAttachmentStatusOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(VolumeAttachmentStatusOutput{}) }

func (VolumeAttachmentStatusOutput) ElementType() reflect.Type {
	return _VolumeAttachmentStatusType
}

func (o VolumeAttachmentStatusOutput) AttachError() VolumeErrorOutput {
	return o.Apply(func(v VolumeAttachmentStatus) *VolumeError {
		return v.AttachError
	}).(VolumeErrorOutput)
}

func (o VolumeAttachmentStatusOutput) Attached() pulumi.BoolOutput {
	return o.Apply(func(v VolumeAttachmentStatus) bool {
		return v.Attached
	}).(pulumi.BoolOutput)
}

func (o VolumeAttachmentStatusOutput) AttachmentMetadata() pulumi.StringMapOutput {
	return o.Apply(func(v VolumeAttachmentStatus) map[string]string {
		return v.AttachmentMetadata
	}).(pulumi.StringMapOutput)
}

func (o VolumeAttachmentStatusOutput) DetachError() VolumeErrorOutput {
	return o.Apply(func(v VolumeAttachmentStatus) *VolumeError {
		return v.DetachError
	}).(VolumeErrorOutput)
}

