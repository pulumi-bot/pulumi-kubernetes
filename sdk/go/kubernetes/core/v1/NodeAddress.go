// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// NodeAddress contains information for the node's address.
type NodeAddress struct {
	// The node address.
	Address string `pulumi:"address"`

	// Node address type, one of Hostname, ExternalIP or InternalIP.
	Type string `pulumi:"type"`

}

var _NodeAddressType = reflect.TypeOf((*NodeAddress)(nil)).Elem()

// NodeAddressInput represents an input type that resolves to a NodeAddress.
type NodeAddressInput interface {
	ElementType() reflect.Type

	ToNodeAddressOutput() NodeAddressOutput
	ToNodeAddressOutputWithContext(ctx context.Context) NodeAddressOutput
}

// NodeAddressArgs is a NodeAddressInput whose fields are all Input types.
type NodeAddressArgs struct {
	// The node address.
	Address pulumi.StringInput `pulumi:"address"`

	// Node address type, one of Hostname, ExternalIP or InternalIP.
	Type pulumi.StringInput `pulumi:"type"`

}

func (a NodeAddressArgs) ElementType() reflect.Type {
	return _NodeAddressType
}

func (a NodeAddressArgs) ToNodeAddressOutput() NodeAddressOutput {
	return pulumi.ToOutput(a).(NodeAddressOutput)
}

func (a NodeAddressArgs) ToNodeAddressOutputWithContext(ctx context.Context) NodeAddressOutput {
	return pulumi.ToOutputWithContext(ctx, a).(NodeAddressOutput)
}

// NodeAddressOutput is an output type that resolves to a Input.
type NodeAddressOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(NodeAddressOutput{}) }

func (NodeAddressOutput) ElementType() reflect.Type {
	return _NodeAddressType
}

func (o NodeAddressOutput) Address() pulumi.StringOutput {
	return o.Apply(func(v NodeAddress) string {
		return v.Address
	}).(pulumi.StringOutput)
}

func (o NodeAddressOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v NodeAddress) string {
		return v.Type
	}).(pulumi.StringOutput)
}

var _NodeAddressArrayType = reflect.TypeOf((*[]NodeAddress)(nil)).Elem()

type NodeAddressArrayInput interface {
	ElementType() reflect.Type

	ToNodeAddressArrayOutput() NodeAddressArrayOutput
	ToNodeAddressArrayOutputWithContext(ctx context.Context) NodeAddressArrayOutput
}

type NodeAddressArray []NodeAddressInput

func (a NodeAddressArray) ElementType() reflect.Type {
	return _NodeAddressArrayType
}

func (a NodeAddressArray) ToNodeAddressArrayOutput() NodeAddressArrayOutput {
	return pulumi.ToOutput(a).(NodeAddressArrayOutput)
}

func (a NodeAddressArray) ToNodeAddressArrayOutputWithContext(ctx context.Context) NodeAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(NodeAddressArrayOutput)
}

type NodeAddressArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(NodeAddressArrayOutput{}) }

func (NodeAddressArrayOutput) ElementType() reflect.Type {
	return _NodeAddressArrayType
}

func (o NodeAddressArrayOutput) Index(i pulumi.IntInput) NodeAddressOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) NodeAddress {
		return vs[0].([]NodeAddress)[vs[1].(int)]
	}).(NodeAddressOutput)
}