// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ServiceReference holds a reference to Service.legacy.k8s.io
type ServiceReference struct {
	// `name` is the name of the service. Required
	Name string `pulumi:"name"`

	// `namespace` is the namespace of the service. Required
	Namespace string `pulumi:"namespace"`

	// `path` is an optional URL path which will be sent in any request to this service.
	Path *string `pulumi:"path"`

	// If specified, the port on the service that hosting webhook. Default to 443 for backward
	// compatibility. `port` should be a valid port number (1-65535, inclusive).
	Port *int `pulumi:"port"`

}

var _ServiceReferenceType = reflect.TypeOf((*ServiceReference)(nil)).Elem()

// ServiceReferenceInput represents an input type that resolves to a ServiceReference.
type ServiceReferenceInput interface {
	ElementType() reflect.Type

	ToServiceReferenceOutput() ServiceReferenceOutput
	ToServiceReferenceOutputWithContext(ctx context.Context) ServiceReferenceOutput
}

// ServiceReferenceArgs is a ServiceReferenceInput whose fields are all Input types.
type ServiceReferenceArgs struct {
	// `name` is the name of the service. Required
	Name pulumi.StringInput `pulumi:"name"`

	// `namespace` is the namespace of the service. Required
	Namespace pulumi.StringInput `pulumi:"namespace"`

	// `path` is an optional URL path which will be sent in any request to this service.
	Path pulumi.StringInput `pulumi:"path"`

	// If specified, the port on the service that hosting webhook. Default to 443 for backward
	// compatibility. `port` should be a valid port number (1-65535, inclusive).
	Port pulumi.IntInput `pulumi:"port"`

}

func (a ServiceReferenceArgs) ElementType() reflect.Type {
	return _ServiceReferenceType
}

func (a ServiceReferenceArgs) ToServiceReferenceOutput() ServiceReferenceOutput {
	return pulumi.ToOutput(a).(ServiceReferenceOutput)
}

func (a ServiceReferenceArgs) ToServiceReferenceOutputWithContext(ctx context.Context) ServiceReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ServiceReferenceOutput)
}

// ServiceReferenceOutput is an output type that resolves to a Input.
type ServiceReferenceOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ServiceReferenceOutput{}) }

func (ServiceReferenceOutput) ElementType() reflect.Type {
	return _ServiceReferenceType
}

func (o ServiceReferenceOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v ServiceReference) string {
		return v.Name
	}).(pulumi.StringOutput)
}

func (o ServiceReferenceOutput) Namespace() pulumi.StringOutput {
	return o.Apply(func(v ServiceReference) string {
		return v.Namespace
	}).(pulumi.StringOutput)
}

func (o ServiceReferenceOutput) Path() pulumi.StringOutput {
	return o.Apply(func(v ServiceReference) *string {
		return v.Path
	}).(pulumi.StringOutput)
}

func (o ServiceReferenceOutput) Port() pulumi.IntOutput {
	return o.Apply(func(v ServiceReference) *int {
		return v.Port
	}).(pulumi.IntOutput)
}

