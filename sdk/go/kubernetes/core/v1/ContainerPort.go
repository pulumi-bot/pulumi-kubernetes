// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ContainerPort represents a network port in a single container.
type ContainerPort struct {
	// Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x <
	// 65536.
	ContainerPort int `pulumi:"containerPort"`

	// What host IP to bind the external port to.
	HostIP *string `pulumi:"hostIP"`

	// Number of port to expose on the host. If specified, this must be a valid port number, 0 < x <
	// 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need
	// this.
	HostPort *int `pulumi:"hostPort"`

	// If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod
	// must have a unique name. Name for the port that can be referred to by services.
	Name *string `pulumi:"name"`

	// Protocol for port. Must be UDP, TCP, or SCTP. Defaults to "TCP".
	Protocol *string `pulumi:"protocol"`

}

var _ContainerPortType = reflect.TypeOf((*ContainerPort)(nil)).Elem()

// ContainerPortInput represents an input type that resolves to a ContainerPort.
type ContainerPortInput interface {
	ElementType() reflect.Type

	ToContainerPortOutput() ContainerPortOutput
	ToContainerPortOutputWithContext(ctx context.Context) ContainerPortOutput
}

// ContainerPortArgs is a ContainerPortInput whose fields are all Input types.
type ContainerPortArgs struct {
	// Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x <
	// 65536.
	ContainerPort pulumi.IntInput `pulumi:"containerPort"`

	// What host IP to bind the external port to.
	HostIP pulumi.StringInput `pulumi:"hostIP"`

	// Number of port to expose on the host. If specified, this must be a valid port number, 0 < x <
	// 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need
	// this.
	HostPort pulumi.IntInput `pulumi:"hostPort"`

	// If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod
	// must have a unique name. Name for the port that can be referred to by services.
	Name pulumi.StringInput `pulumi:"name"`

	// Protocol for port. Must be UDP, TCP, or SCTP. Defaults to "TCP".
	Protocol pulumi.StringInput `pulumi:"protocol"`

}

func (a ContainerPortArgs) ElementType() reflect.Type {
	return _ContainerPortType
}

func (a ContainerPortArgs) ToContainerPortOutput() ContainerPortOutput {
	return pulumi.ToOutput(a).(ContainerPortOutput)
}

func (a ContainerPortArgs) ToContainerPortOutputWithContext(ctx context.Context) ContainerPortOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ContainerPortOutput)
}

// ContainerPortOutput is an output type that resolves to a Input.
type ContainerPortOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ContainerPortOutput{}) }

func (ContainerPortOutput) ElementType() reflect.Type {
	return _ContainerPortType
}

func (o ContainerPortOutput) ContainerPort() pulumi.IntOutput {
	return o.Apply(func(v ContainerPort) int {
		return v.ContainerPort
	}).(pulumi.IntOutput)
}

func (o ContainerPortOutput) HostIP() pulumi.StringOutput {
	return o.Apply(func(v ContainerPort) *string {
		return v.HostIP
	}).(pulumi.StringOutput)
}

func (o ContainerPortOutput) HostPort() pulumi.IntOutput {
	return o.Apply(func(v ContainerPort) *int {
		return v.HostPort
	}).(pulumi.IntOutput)
}

func (o ContainerPortOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v ContainerPort) *string {
		return v.Name
	}).(pulumi.StringOutput)
}

func (o ContainerPortOutput) Protocol() pulumi.StringOutput {
	return o.Apply(func(v ContainerPort) *string {
		return v.Protocol
	}).(pulumi.StringOutput)
}

var _ContainerPortArrayType = reflect.TypeOf((*[]ContainerPort)(nil)).Elem()

type ContainerPortArrayInput interface {
	ElementType() reflect.Type

	ToContainerPortArrayOutput() ContainerPortArrayOutput
	ToContainerPortArrayOutputWithContext(ctx context.Context) ContainerPortArrayOutput
}

type ContainerPortArray []ContainerPortInput

func (a ContainerPortArray) ElementType() reflect.Type {
	return _ContainerPortArrayType
}

func (a ContainerPortArray) ToContainerPortArrayOutput() ContainerPortArrayOutput {
	return pulumi.ToOutput(a).(ContainerPortArrayOutput)
}

func (a ContainerPortArray) ToContainerPortArrayOutputWithContext(ctx context.Context) ContainerPortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ContainerPortArrayOutput)
}

type ContainerPortArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ContainerPortArrayOutput{}) }

func (ContainerPortArrayOutput) ElementType() reflect.Type {
	return _ContainerPortArrayType
}

func (o ContainerPortArrayOutput) Index(i pulumi.IntInput) ContainerPortOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) ContainerPort {
		return vs[0].([]ContainerPort)[vs[1].(int)]
	}).(ContainerPortOutput)
}
