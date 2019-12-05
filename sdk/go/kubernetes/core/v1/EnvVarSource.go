// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// EnvVarSource represents a source for the value of an EnvVar.
type EnvVarSource struct {
	// Selects a key of a ConfigMap.
	ConfigMapKeyRef *ConfigMapKeySelector `pulumi:"configMapKeyRef"`

	// Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels,
	// metadata.annotations, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP.
	FieldRef *ObjectFieldSelector `pulumi:"fieldRef"`

	// Selects a resource of the container: only resources limits and requests (limits.cpu,
	// limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and
	// requests.ephemeral-storage) are currently supported.
	ResourceFieldRef *ResourceFieldSelector `pulumi:"resourceFieldRef"`

	// Selects a key of a secret in the pod's namespace
	SecretKeyRef *SecretKeySelector `pulumi:"secretKeyRef"`

}

var _EnvVarSourceType = reflect.TypeOf((*EnvVarSource)(nil)).Elem()

// EnvVarSourceInput represents an input type that resolves to a EnvVarSource.
type EnvVarSourceInput interface {
	ElementType() reflect.Type

	ToEnvVarSourceOutput() EnvVarSourceOutput
	ToEnvVarSourceOutputWithContext(ctx context.Context) EnvVarSourceOutput
}

// EnvVarSourceArgs is a EnvVarSourceInput whose fields are all Input types.
type EnvVarSourceArgs struct {
	// Selects a key of a ConfigMap.
	ConfigMapKeyRef ConfigMapKeySelectorInput `pulumi:"configMapKeyRef"`

	// Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels,
	// metadata.annotations, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP.
	FieldRef ObjectFieldSelectorInput `pulumi:"fieldRef"`

	// Selects a resource of the container: only resources limits and requests (limits.cpu,
	// limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and
	// requests.ephemeral-storage) are currently supported.
	ResourceFieldRef ResourceFieldSelectorInput `pulumi:"resourceFieldRef"`

	// Selects a key of a secret in the pod's namespace
	SecretKeyRef SecretKeySelectorInput `pulumi:"secretKeyRef"`

}

func (a EnvVarSourceArgs) ElementType() reflect.Type {
	return _EnvVarSourceType
}

func (a EnvVarSourceArgs) ToEnvVarSourceOutput() EnvVarSourceOutput {
	return pulumi.ToOutput(a).(EnvVarSourceOutput)
}

func (a EnvVarSourceArgs) ToEnvVarSourceOutputWithContext(ctx context.Context) EnvVarSourceOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EnvVarSourceOutput)
}

// EnvVarSourceOutput is an output type that resolves to a Input.
type EnvVarSourceOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(EnvVarSourceOutput{}) }

func (EnvVarSourceOutput) ElementType() reflect.Type {
	return _EnvVarSourceType
}

func (o EnvVarSourceOutput) ConfigMapKeyRef() ConfigMapKeySelectorOutput {
	return o.Apply(func(v EnvVarSource) *ConfigMapKeySelector {
		return v.ConfigMapKeyRef
	}).(ConfigMapKeySelectorOutput)
}

func (o EnvVarSourceOutput) FieldRef() ObjectFieldSelectorOutput {
	return o.Apply(func(v EnvVarSource) *ObjectFieldSelector {
		return v.FieldRef
	}).(ObjectFieldSelectorOutput)
}

func (o EnvVarSourceOutput) ResourceFieldRef() ResourceFieldSelectorOutput {
	return o.Apply(func(v EnvVarSource) *ResourceFieldSelector {
		return v.ResourceFieldRef
	}).(ResourceFieldSelectorOutput)
}

func (o EnvVarSourceOutput) SecretKeyRef() SecretKeySelectorOutput {
	return o.Apply(func(v EnvVarSource) *SecretKeySelector {
		return v.SecretKeyRef
	}).(SecretKeySelectorOutput)
}

