// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ServerAddressByClientCIDR helps the client to determine the server address that they should use,
// depending on the clientCIDR that they match.
type ServerAddressByClientCIDR struct {
	// The CIDR with which clients can match their IP to figure out the server address that they should
	// use.
	ClientCIDR string `pulumi:"clientCIDR"`

	// Address of this server, suitable for a client that matches the above CIDR. This can be a
	// hostname, hostname:port, IP or IP:port.
	ServerAddress string `pulumi:"serverAddress"`

}

var _ServerAddressByClientCIDRType = reflect.TypeOf((*ServerAddressByClientCIDR)(nil)).Elem()

// ServerAddressByClientCIDRInput represents an input type that resolves to a ServerAddressByClientCIDR.
type ServerAddressByClientCIDRInput interface {
	ElementType() reflect.Type

	ToServerAddressByClientCIDROutput() ServerAddressByClientCIDROutput
	ToServerAddressByClientCIDROutputWithContext(ctx context.Context) ServerAddressByClientCIDROutput
}

// ServerAddressByClientCIDRArgs is a ServerAddressByClientCIDRInput whose fields are all Input types.
type ServerAddressByClientCIDRArgs struct {
	// The CIDR with which clients can match their IP to figure out the server address that they should
	// use.
	ClientCIDR pulumi.StringInput `pulumi:"clientCIDR"`

	// Address of this server, suitable for a client that matches the above CIDR. This can be a
	// hostname, hostname:port, IP or IP:port.
	ServerAddress pulumi.StringInput `pulumi:"serverAddress"`

}

func (a ServerAddressByClientCIDRArgs) ElementType() reflect.Type {
	return _ServerAddressByClientCIDRType
}

func (a ServerAddressByClientCIDRArgs) ToServerAddressByClientCIDROutput() ServerAddressByClientCIDROutput {
	return pulumi.ToOutput(a).(ServerAddressByClientCIDROutput)
}

func (a ServerAddressByClientCIDRArgs) ToServerAddressByClientCIDROutputWithContext(ctx context.Context) ServerAddressByClientCIDROutput {
	return pulumi.ToOutputWithContext(ctx, a).(ServerAddressByClientCIDROutput)
}

// ServerAddressByClientCIDROutput is an output type that resolves to a Input.
type ServerAddressByClientCIDROutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ServerAddressByClientCIDROutput{}) }

func (ServerAddressByClientCIDROutput) ElementType() reflect.Type {
	return _ServerAddressByClientCIDRType
}

func (o ServerAddressByClientCIDROutput) ClientCIDR() pulumi.StringOutput {
	return o.Apply(func(v ServerAddressByClientCIDR) string {
		return v.ClientCIDR
	}).(pulumi.StringOutput)
}

func (o ServerAddressByClientCIDROutput) ServerAddress() pulumi.StringOutput {
	return o.Apply(func(v ServerAddressByClientCIDR) string {
		return v.ServerAddress
	}).(pulumi.StringOutput)
}

var _ServerAddressByClientCIDRArrayType = reflect.TypeOf((*[]ServerAddressByClientCIDR)(nil)).Elem()

type ServerAddressByClientCIDRArrayInput interface {
	ElementType() reflect.Type

	ToServerAddressByClientCIDRArrayOutput() ServerAddressByClientCIDRArrayOutput
	ToServerAddressByClientCIDRArrayOutputWithContext(ctx context.Context) ServerAddressByClientCIDRArrayOutput
}

type ServerAddressByClientCIDRArray []ServerAddressByClientCIDRInput

func (a ServerAddressByClientCIDRArray) ElementType() reflect.Type {
	return _ServerAddressByClientCIDRArrayType
}

func (a ServerAddressByClientCIDRArray) ToServerAddressByClientCIDRArrayOutput() ServerAddressByClientCIDRArrayOutput {
	return pulumi.ToOutput(a).(ServerAddressByClientCIDRArrayOutput)
}

func (a ServerAddressByClientCIDRArray) ToServerAddressByClientCIDRArrayOutputWithContext(ctx context.Context) ServerAddressByClientCIDRArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ServerAddressByClientCIDRArrayOutput)
}

type ServerAddressByClientCIDRArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ServerAddressByClientCIDRArrayOutput{}) }

func (ServerAddressByClientCIDRArrayOutput) ElementType() reflect.Type {
	return _ServerAddressByClientCIDRArrayType
}

func (o ServerAddressByClientCIDRArrayOutput) Index(i pulumi.IntInput) ServerAddressByClientCIDROutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) ServerAddressByClientCIDR {
		return vs[0].([]ServerAddressByClientCIDR)[vs[1].(int)]
	}).(ServerAddressByClientCIDROutput)
}
