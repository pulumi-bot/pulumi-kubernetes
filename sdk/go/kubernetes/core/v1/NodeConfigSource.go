// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// NodeConfigSource specifies a source of node configuration. Exactly one subfield (excluding
// metadata) must be non-nil.
type NodeConfigSource struct {
	// ConfigMap is a reference to a Node's ConfigMap
	ConfigMap *ConfigMapNodeConfigSource `pulumi:"configMap"`

}

var _NodeConfigSourceType = reflect.TypeOf((*NodeConfigSource)(nil)).Elem()

// NodeConfigSourceInput represents an input type that resolves to a NodeConfigSource.
type NodeConfigSourceInput interface {
	ElementType() reflect.Type

	ToNodeConfigSourceOutput() NodeConfigSourceOutput
	ToNodeConfigSourceOutputWithContext(ctx context.Context) NodeConfigSourceOutput
}

// NodeConfigSourceArgs is a NodeConfigSourceInput whose fields are all Input types.
type NodeConfigSourceArgs struct {
	// ConfigMap is a reference to a Node's ConfigMap
	ConfigMap ConfigMapNodeConfigSourceInput `pulumi:"configMap"`

}

func (a NodeConfigSourceArgs) ElementType() reflect.Type {
	return _NodeConfigSourceType
}

func (a NodeConfigSourceArgs) ToNodeConfigSourceOutput() NodeConfigSourceOutput {
	return pulumi.ToOutput(a).(NodeConfigSourceOutput)
}

func (a NodeConfigSourceArgs) ToNodeConfigSourceOutputWithContext(ctx context.Context) NodeConfigSourceOutput {
	return pulumi.ToOutputWithContext(ctx, a).(NodeConfigSourceOutput)
}

// NodeConfigSourceOutput is an output type that resolves to a Input.
type NodeConfigSourceOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(NodeConfigSourceOutput{}) }

func (NodeConfigSourceOutput) ElementType() reflect.Type {
	return _NodeConfigSourceType
}

func (o NodeConfigSourceOutput) ConfigMap() ConfigMapNodeConfigSourceOutput {
	return o.Apply(func(v NodeConfigSource) *ConfigMapNodeConfigSource {
		return v.ConfigMap
	}).(ConfigMapNodeConfigSourceOutput)
}

