// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Selects a key from a ConfigMap.
type ConfigMapKeySelector struct {
	// The key to select.
	Key string `pulumi:"key"`

	// Name of the referent. More info:
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `pulumi:"name"`

	// Specify whether the ConfigMap or its key must be defined
	Optional *bool `pulumi:"optional"`

}

var _ConfigMapKeySelectorType = reflect.TypeOf((*ConfigMapKeySelector)(nil)).Elem()

// ConfigMapKeySelectorInput represents an input type that resolves to a ConfigMapKeySelector.
type ConfigMapKeySelectorInput interface {
	ElementType() reflect.Type

	ToConfigMapKeySelectorOutput() ConfigMapKeySelectorOutput
	ToConfigMapKeySelectorOutputWithContext(ctx context.Context) ConfigMapKeySelectorOutput
}

// ConfigMapKeySelectorArgs is a ConfigMapKeySelectorInput whose fields are all Input types.
type ConfigMapKeySelectorArgs struct {
	// The key to select.
	Key pulumi.StringInput `pulumi:"key"`

	// Name of the referent. More info:
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name pulumi.StringInput `pulumi:"name"`

	// Specify whether the ConfigMap or its key must be defined
	Optional pulumi.BoolInput `pulumi:"optional"`

}

func (a ConfigMapKeySelectorArgs) ElementType() reflect.Type {
	return _ConfigMapKeySelectorType
}

func (a ConfigMapKeySelectorArgs) ToConfigMapKeySelectorOutput() ConfigMapKeySelectorOutput {
	return pulumi.ToOutput(a).(ConfigMapKeySelectorOutput)
}

func (a ConfigMapKeySelectorArgs) ToConfigMapKeySelectorOutputWithContext(ctx context.Context) ConfigMapKeySelectorOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ConfigMapKeySelectorOutput)
}

// ConfigMapKeySelectorOutput is an output type that resolves to a Input.
type ConfigMapKeySelectorOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ConfigMapKeySelectorOutput{}) }

func (ConfigMapKeySelectorOutput) ElementType() reflect.Type {
	return _ConfigMapKeySelectorType
}

func (o ConfigMapKeySelectorOutput) Key() pulumi.StringOutput {
	return o.Apply(func(v ConfigMapKeySelector) string {
		return v.Key
	}).(pulumi.StringOutput)
}

func (o ConfigMapKeySelectorOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v ConfigMapKeySelector) *string {
		return v.Name
	}).(pulumi.StringOutput)
}

func (o ConfigMapKeySelectorOutput) Optional() pulumi.BoolOutput {
	return o.Apply(func(v ConfigMapKeySelector) *bool {
		return v.Optional
	}).(pulumi.BoolOutput)
}

