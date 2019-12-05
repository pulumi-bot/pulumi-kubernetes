// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// IngressSpec describes the Ingress the user wishes to exist.
type IngressSpec struct {
	// A default backend capable of servicing requests that don't match any rule. At least one of
	// 'backend' or 'rules' must be specified. This field is optional to allow the loadbalancer
	// controller or defaulting logic to specify a global default.
	Backend *IngressBackend `pulumi:"backend"`

	// A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all
	// traffic is sent to the default backend.
	Rules []IngressRule `pulumi:"rules"`

	// TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple
	// members of this list specify different hosts, they will be multiplexed on the same port
	// according to the hostname specified through the SNI TLS extension, if the ingress controller
	// fulfilling the ingress supports SNI.
	Tls []IngressTLS `pulumi:"tls"`

}

var _IngressSpecType = reflect.TypeOf((*IngressSpec)(nil)).Elem()

// IngressSpecInput represents an input type that resolves to a IngressSpec.
type IngressSpecInput interface {
	ElementType() reflect.Type

	ToIngressSpecOutput() IngressSpecOutput
	ToIngressSpecOutputWithContext(ctx context.Context) IngressSpecOutput
}

// IngressSpecArgs is a IngressSpecInput whose fields are all Input types.
type IngressSpecArgs struct {
	// A default backend capable of servicing requests that don't match any rule. At least one of
	// 'backend' or 'rules' must be specified. This field is optional to allow the loadbalancer
	// controller or defaulting logic to specify a global default.
	Backend IngressBackendInput `pulumi:"backend"`

	// A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all
	// traffic is sent to the default backend.
	Rules IngressRuleArrayInput `pulumi:"rules"`

	// TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple
	// members of this list specify different hosts, they will be multiplexed on the same port
	// according to the hostname specified through the SNI TLS extension, if the ingress controller
	// fulfilling the ingress supports SNI.
	Tls IngressTLSArrayInput `pulumi:"tls"`

}

func (a IngressSpecArgs) ElementType() reflect.Type {
	return _IngressSpecType
}

func (a IngressSpecArgs) ToIngressSpecOutput() IngressSpecOutput {
	return pulumi.ToOutput(a).(IngressSpecOutput)
}

func (a IngressSpecArgs) ToIngressSpecOutputWithContext(ctx context.Context) IngressSpecOutput {
	return pulumi.ToOutputWithContext(ctx, a).(IngressSpecOutput)
}

// IngressSpecOutput is an output type that resolves to a Input.
type IngressSpecOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(IngressSpecOutput{}) }

func (IngressSpecOutput) ElementType() reflect.Type {
	return _IngressSpecType
}

func (o IngressSpecOutput) Backend() IngressBackendOutput {
	return o.Apply(func(v IngressSpec) *IngressBackend {
		return v.Backend
	}).(IngressBackendOutput)
}

func (o IngressSpecOutput) Rules() IngressRuleArrayOutput {
	return o.Apply(func(v IngressSpec) []IngressRule {
		return v.Rules
	}).(IngressRuleArrayOutput)
}

func (o IngressSpecOutput) Tls() IngressTLSArrayOutput {
	return o.Apply(func(v IngressSpec) []IngressTLS {
		return v.Tls
	}).(IngressTLSArrayOutput)
}

