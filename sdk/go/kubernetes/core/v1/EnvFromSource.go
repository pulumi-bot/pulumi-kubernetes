// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// EnvFromSource represents the source of a set of ConfigMaps
type EnvFromSource struct {
	// The ConfigMap to select from
	ConfigMapRef *ConfigMapEnvSource `pulumi:"configMapRef"`

	// An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
	Prefix *string `pulumi:"prefix"`

	// The Secret to select from
	SecretRef *SecretEnvSource `pulumi:"secretRef"`

}

var _EnvFromSourceType = reflect.TypeOf((*EnvFromSource)(nil)).Elem()

// EnvFromSourceInput represents an input type that resolves to a EnvFromSource.
type EnvFromSourceInput interface {
	ElementType() reflect.Type

	ToEnvFromSourceOutput() EnvFromSourceOutput
	ToEnvFromSourceOutputWithContext(ctx context.Context) EnvFromSourceOutput
}

// EnvFromSourceArgs is a EnvFromSourceInput whose fields are all Input types.
type EnvFromSourceArgs struct {
	// The ConfigMap to select from
	ConfigMapRef ConfigMapEnvSourceInput `pulumi:"configMapRef"`

	// An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
	Prefix pulumi.StringInput `pulumi:"prefix"`

	// The Secret to select from
	SecretRef SecretEnvSourceInput `pulumi:"secretRef"`

}

func (a EnvFromSourceArgs) ElementType() reflect.Type {
	return _EnvFromSourceType
}

func (a EnvFromSourceArgs) ToEnvFromSourceOutput() EnvFromSourceOutput {
	return pulumi.ToOutput(a).(EnvFromSourceOutput)
}

func (a EnvFromSourceArgs) ToEnvFromSourceOutputWithContext(ctx context.Context) EnvFromSourceOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EnvFromSourceOutput)
}

// EnvFromSourceOutput is an output type that resolves to a Input.
type EnvFromSourceOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(EnvFromSourceOutput{}) }

func (EnvFromSourceOutput) ElementType() reflect.Type {
	return _EnvFromSourceType
}

func (o EnvFromSourceOutput) ConfigMapRef() ConfigMapEnvSourceOutput {
	return o.Apply(func(v EnvFromSource) *ConfigMapEnvSource {
		return v.ConfigMapRef
	}).(ConfigMapEnvSourceOutput)
}

func (o EnvFromSourceOutput) Prefix() pulumi.StringOutput {
	return o.Apply(func(v EnvFromSource) *string {
		return v.Prefix
	}).(pulumi.StringOutput)
}

func (o EnvFromSourceOutput) SecretRef() SecretEnvSourceOutput {
	return o.Apply(func(v EnvFromSource) *SecretEnvSource {
		return v.SecretRef
	}).(SecretEnvSourceOutput)
}

var _EnvFromSourceArrayType = reflect.TypeOf((*[]EnvFromSource)(nil)).Elem()

type EnvFromSourceArrayInput interface {
	ElementType() reflect.Type

	ToEnvFromSourceArrayOutput() EnvFromSourceArrayOutput
	ToEnvFromSourceArrayOutputWithContext(ctx context.Context) EnvFromSourceArrayOutput
}

type EnvFromSourceArray []EnvFromSourceInput

func (a EnvFromSourceArray) ElementType() reflect.Type {
	return _EnvFromSourceArrayType
}

func (a EnvFromSourceArray) ToEnvFromSourceArrayOutput() EnvFromSourceArrayOutput {
	return pulumi.ToOutput(a).(EnvFromSourceArrayOutput)
}

func (a EnvFromSourceArray) ToEnvFromSourceArrayOutputWithContext(ctx context.Context) EnvFromSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EnvFromSourceArrayOutput)
}

type EnvFromSourceArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(EnvFromSourceArrayOutput{}) }

func (EnvFromSourceArrayOutput) ElementType() reflect.Type {
	return _EnvFromSourceArrayType
}

func (o EnvFromSourceArrayOutput) Index(i pulumi.IntInput) EnvFromSourceOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) EnvFromSource {
		return vs[0].([]EnvFromSource)[vs[1].(int)]
	}).(EnvFromSourceOutput)
}
