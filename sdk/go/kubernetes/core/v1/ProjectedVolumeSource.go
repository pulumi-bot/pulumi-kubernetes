// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Represents a projected volume source
type ProjectedVolumeSource struct {
	// Mode bits to use on created files by default. Must be a value between 0 and 0777. Directories
	// within the path are not affected by this setting. This might be in conflict with other options
	// that affect the file mode, like fsGroup, and the result can be other mode bits set.
	DefaultMode *int `pulumi:"defaultMode"`

	// list of volume projections
	Sources []VolumeProjection `pulumi:"sources"`

}

var _ProjectedVolumeSourceType = reflect.TypeOf((*ProjectedVolumeSource)(nil)).Elem()

// ProjectedVolumeSourceInput represents an input type that resolves to a ProjectedVolumeSource.
type ProjectedVolumeSourceInput interface {
	ElementType() reflect.Type

	ToProjectedVolumeSourceOutput() ProjectedVolumeSourceOutput
	ToProjectedVolumeSourceOutputWithContext(ctx context.Context) ProjectedVolumeSourceOutput
}

// ProjectedVolumeSourceArgs is a ProjectedVolumeSourceInput whose fields are all Input types.
type ProjectedVolumeSourceArgs struct {
	// list of volume projections
	Sources VolumeProjectionArrayInput `pulumi:"sources"`

	// Mode bits to use on created files by default. Must be a value between 0 and 0777. Directories
	// within the path are not affected by this setting. This might be in conflict with other options
	// that affect the file mode, like fsGroup, and the result can be other mode bits set.
	DefaultMode pulumi.IntInput `pulumi:"defaultMode"`

}

func (a ProjectedVolumeSourceArgs) ElementType() reflect.Type {
	return _ProjectedVolumeSourceType
}

func (a ProjectedVolumeSourceArgs) ToProjectedVolumeSourceOutput() ProjectedVolumeSourceOutput {
	return pulumi.ToOutput(a).(ProjectedVolumeSourceOutput)
}

func (a ProjectedVolumeSourceArgs) ToProjectedVolumeSourceOutputWithContext(ctx context.Context) ProjectedVolumeSourceOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ProjectedVolumeSourceOutput)
}

// ProjectedVolumeSourceOutput is an output type that resolves to a Input.
type ProjectedVolumeSourceOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ProjectedVolumeSourceOutput{}) }

func (ProjectedVolumeSourceOutput) ElementType() reflect.Type {
	return _ProjectedVolumeSourceType
}

func (o ProjectedVolumeSourceOutput) DefaultMode() pulumi.IntOutput {
	return o.Apply(func(v ProjectedVolumeSource) *int {
		return v.DefaultMode
	}).(pulumi.IntOutput)
}

func (o ProjectedVolumeSourceOutput) Sources() VolumeProjectionArrayOutput {
	return o.Apply(func(v ProjectedVolumeSource) []VolumeProjection {
		return v.Sources
	}).(VolumeProjectionArrayOutput)
}

