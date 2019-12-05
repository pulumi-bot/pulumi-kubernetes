// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// DownwardAPIVolumeFile represents information to create the file containing the pod field
type DownwardAPIVolumeFile struct {
	// Required: Selects a field of the pod: only annotations, labels, name and namespace are
	// supported.
	FieldRef *ObjectFieldSelector `pulumi:"fieldRef"`

	// Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified,
	// the volume defaultMode will be used. This might be in conflict with other options that affect
	// the file mode, like fsGroup, and the result can be other mode bits set.
	Mode *int `pulumi:"mode"`

	// Required: Path is  the relative path name of the file to be created. Must not be absolute or
	// contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start
	// with '..'
	Path string `pulumi:"path"`

	// Selects a resource of the container: only resources limits and requests (limits.cpu,
	// limits.memory, requests.cpu and requests.memory) are currently supported.
	ResourceFieldRef *ResourceFieldSelector `pulumi:"resourceFieldRef"`

}

var _DownwardAPIVolumeFileType = reflect.TypeOf((*DownwardAPIVolumeFile)(nil)).Elem()

// DownwardAPIVolumeFileInput represents an input type that resolves to a DownwardAPIVolumeFile.
type DownwardAPIVolumeFileInput interface {
	ElementType() reflect.Type

	ToDownwardAPIVolumeFileOutput() DownwardAPIVolumeFileOutput
	ToDownwardAPIVolumeFileOutputWithContext(ctx context.Context) DownwardAPIVolumeFileOutput
}

// DownwardAPIVolumeFileArgs is a DownwardAPIVolumeFileInput whose fields are all Input types.
type DownwardAPIVolumeFileArgs struct {
	// Required: Path is  the relative path name of the file to be created. Must not be absolute or
	// contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start
	// with '..'
	Path pulumi.StringInput `pulumi:"path"`

	// Required: Selects a field of the pod: only annotations, labels, name and namespace are
	// supported.
	FieldRef ObjectFieldSelectorInput `pulumi:"fieldRef"`

	// Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified,
	// the volume defaultMode will be used. This might be in conflict with other options that affect
	// the file mode, like fsGroup, and the result can be other mode bits set.
	Mode pulumi.IntInput `pulumi:"mode"`

	// Selects a resource of the container: only resources limits and requests (limits.cpu,
	// limits.memory, requests.cpu and requests.memory) are currently supported.
	ResourceFieldRef ResourceFieldSelectorInput `pulumi:"resourceFieldRef"`

}

func (a DownwardAPIVolumeFileArgs) ElementType() reflect.Type {
	return _DownwardAPIVolumeFileType
}

func (a DownwardAPIVolumeFileArgs) ToDownwardAPIVolumeFileOutput() DownwardAPIVolumeFileOutput {
	return pulumi.ToOutput(a).(DownwardAPIVolumeFileOutput)
}

func (a DownwardAPIVolumeFileArgs) ToDownwardAPIVolumeFileOutputWithContext(ctx context.Context) DownwardAPIVolumeFileOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DownwardAPIVolumeFileOutput)
}

// DownwardAPIVolumeFileOutput is an output type that resolves to a Input.
type DownwardAPIVolumeFileOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(DownwardAPIVolumeFileOutput{}) }

func (DownwardAPIVolumeFileOutput) ElementType() reflect.Type {
	return _DownwardAPIVolumeFileType
}

func (o DownwardAPIVolumeFileOutput) FieldRef() ObjectFieldSelectorOutput {
	return o.Apply(func(v DownwardAPIVolumeFile) *ObjectFieldSelector {
		return v.FieldRef
	}).(ObjectFieldSelectorOutput)
}

func (o DownwardAPIVolumeFileOutput) Mode() pulumi.IntOutput {
	return o.Apply(func(v DownwardAPIVolumeFile) *int {
		return v.Mode
	}).(pulumi.IntOutput)
}

func (o DownwardAPIVolumeFileOutput) Path() pulumi.StringOutput {
	return o.Apply(func(v DownwardAPIVolumeFile) string {
		return v.Path
	}).(pulumi.StringOutput)
}

func (o DownwardAPIVolumeFileOutput) ResourceFieldRef() ResourceFieldSelectorOutput {
	return o.Apply(func(v DownwardAPIVolumeFile) *ResourceFieldSelector {
		return v.ResourceFieldRef
	}).(ResourceFieldSelectorOutput)
}

var _DownwardAPIVolumeFileArrayType = reflect.TypeOf((*[]DownwardAPIVolumeFile)(nil)).Elem()

type DownwardAPIVolumeFileArrayInput interface {
	ElementType() reflect.Type

	ToDownwardAPIVolumeFileArrayOutput() DownwardAPIVolumeFileArrayOutput
	ToDownwardAPIVolumeFileArrayOutputWithContext(ctx context.Context) DownwardAPIVolumeFileArrayOutput
}

type DownwardAPIVolumeFileArray []DownwardAPIVolumeFileInput

func (a DownwardAPIVolumeFileArray) ElementType() reflect.Type {
	return _DownwardAPIVolumeFileArrayType
}

func (a DownwardAPIVolumeFileArray) ToDownwardAPIVolumeFileArrayOutput() DownwardAPIVolumeFileArrayOutput {
	return pulumi.ToOutput(a).(DownwardAPIVolumeFileArrayOutput)
}

func (a DownwardAPIVolumeFileArray) ToDownwardAPIVolumeFileArrayOutputWithContext(ctx context.Context) DownwardAPIVolumeFileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DownwardAPIVolumeFileArrayOutput)
}

type DownwardAPIVolumeFileArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(DownwardAPIVolumeFileArrayOutput{}) }

func (DownwardAPIVolumeFileArrayOutput) ElementType() reflect.Type {
	return _DownwardAPIVolumeFileArrayType
}

func (o DownwardAPIVolumeFileArrayOutput) Index(i pulumi.IntInput) DownwardAPIVolumeFileOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) DownwardAPIVolumeFile {
		return vs[0].([]DownwardAPIVolumeFile)[vs[1].(int)]
	}).(DownwardAPIVolumeFileOutput)
}