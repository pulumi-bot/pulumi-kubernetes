// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Adapts a ConfigMap into a projected volume.
// 
// The contents of the target ConfigMap's Data field will be presented in a projected volume as
// files using the keys in the Data field as the file names, unless the items element is populated
// with specific mappings of keys to paths. Note that this is identical to a configmap volume source
// without the default mode.
type ConfigMapProjection struct {
	// If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be
	// projected into the volume as a file whose name is the key and content is the value. If
	// specified, the listed keys will be projected into the specified paths, and unlisted keys will
	// not be present. If a key is specified which is not present in the ConfigMap, the volume setup
	// will error unless it is marked optional. Paths must be relative and may not contain the '..'
	// path or start with '..'.
	Items []KeyToPath `pulumi:"items"`

	// Name of the referent. More info:
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `pulumi:"name"`

	// Specify whether the ConfigMap or its keys must be defined
	Optional *bool `pulumi:"optional"`

}

var _ConfigMapProjectionType = reflect.TypeOf((*ConfigMapProjection)(nil)).Elem()

// ConfigMapProjectionInput represents an input type that resolves to a ConfigMapProjection.
type ConfigMapProjectionInput interface {
	ElementType() reflect.Type

	ToConfigMapProjectionOutput() ConfigMapProjectionOutput
	ToConfigMapProjectionOutputWithContext(ctx context.Context) ConfigMapProjectionOutput
}

// ConfigMapProjectionArgs is a ConfigMapProjectionInput whose fields are all Input types.
type ConfigMapProjectionArgs struct {
	// If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be
	// projected into the volume as a file whose name is the key and content is the value. If
	// specified, the listed keys will be projected into the specified paths, and unlisted keys will
	// not be present. If a key is specified which is not present in the ConfigMap, the volume setup
	// will error unless it is marked optional. Paths must be relative and may not contain the '..'
	// path or start with '..'.
	Items KeyToPathArrayInput `pulumi:"items"`

	// Name of the referent. More info:
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name pulumi.StringInput `pulumi:"name"`

	// Specify whether the ConfigMap or its keys must be defined
	Optional pulumi.BoolInput `pulumi:"optional"`

}

func (a ConfigMapProjectionArgs) ElementType() reflect.Type {
	return _ConfigMapProjectionType
}

func (a ConfigMapProjectionArgs) ToConfigMapProjectionOutput() ConfigMapProjectionOutput {
	return pulumi.ToOutput(a).(ConfigMapProjectionOutput)
}

func (a ConfigMapProjectionArgs) ToConfigMapProjectionOutputWithContext(ctx context.Context) ConfigMapProjectionOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ConfigMapProjectionOutput)
}

// ConfigMapProjectionOutput is an output type that resolves to a Input.
type ConfigMapProjectionOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ConfigMapProjectionOutput{}) }

func (ConfigMapProjectionOutput) ElementType() reflect.Type {
	return _ConfigMapProjectionType
}

func (o ConfigMapProjectionOutput) Items() KeyToPathArrayOutput {
	return o.Apply(func(v ConfigMapProjection) []KeyToPath {
		return v.Items
	}).(KeyToPathArrayOutput)
}

func (o ConfigMapProjectionOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v ConfigMapProjection) *string {
		return v.Name
	}).(pulumi.StringOutput)
}

func (o ConfigMapProjectionOutput) Optional() pulumi.BoolOutput {
	return o.Apply(func(v ConfigMapProjection) *bool {
		return v.Optional
	}).(pulumi.BoolOutput)
}

