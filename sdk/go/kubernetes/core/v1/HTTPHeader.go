// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// HTTPHeader describes a custom header to be used in HTTP probes
type HTTPHeader struct {
	// The header field name
	Name string `pulumi:"name"`

	// The header field value
	Value string `pulumi:"value"`

}

var _HTTPHeaderType = reflect.TypeOf((*HTTPHeader)(nil)).Elem()

// HTTPHeaderInput represents an input type that resolves to a HTTPHeader.
type HTTPHeaderInput interface {
	ElementType() reflect.Type

	ToHTTPHeaderOutput() HTTPHeaderOutput
	ToHTTPHeaderOutputWithContext(ctx context.Context) HTTPHeaderOutput
}

// HTTPHeaderArgs is a HTTPHeaderInput whose fields are all Input types.
type HTTPHeaderArgs struct {
	// The header field name
	Name pulumi.StringInput `pulumi:"name"`

	// The header field value
	Value pulumi.StringInput `pulumi:"value"`

}

func (a HTTPHeaderArgs) ElementType() reflect.Type {
	return _HTTPHeaderType
}

func (a HTTPHeaderArgs) ToHTTPHeaderOutput() HTTPHeaderOutput {
	return pulumi.ToOutput(a).(HTTPHeaderOutput)
}

func (a HTTPHeaderArgs) ToHTTPHeaderOutputWithContext(ctx context.Context) HTTPHeaderOutput {
	return pulumi.ToOutputWithContext(ctx, a).(HTTPHeaderOutput)
}

// HTTPHeaderOutput is an output type that resolves to a Input.
type HTTPHeaderOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(HTTPHeaderOutput{}) }

func (HTTPHeaderOutput) ElementType() reflect.Type {
	return _HTTPHeaderType
}

func (o HTTPHeaderOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v HTTPHeader) string {
		return v.Name
	}).(pulumi.StringOutput)
}

func (o HTTPHeaderOutput) Value() pulumi.StringOutput {
	return o.Apply(func(v HTTPHeader) string {
		return v.Value
	}).(pulumi.StringOutput)
}

var _HTTPHeaderArrayType = reflect.TypeOf((*[]HTTPHeader)(nil)).Elem()

type HTTPHeaderArrayInput interface {
	ElementType() reflect.Type

	ToHTTPHeaderArrayOutput() HTTPHeaderArrayOutput
	ToHTTPHeaderArrayOutputWithContext(ctx context.Context) HTTPHeaderArrayOutput
}

type HTTPHeaderArray []HTTPHeaderInput

func (a HTTPHeaderArray) ElementType() reflect.Type {
	return _HTTPHeaderArrayType
}

func (a HTTPHeaderArray) ToHTTPHeaderArrayOutput() HTTPHeaderArrayOutput {
	return pulumi.ToOutput(a).(HTTPHeaderArrayOutput)
}

func (a HTTPHeaderArray) ToHTTPHeaderArrayOutputWithContext(ctx context.Context) HTTPHeaderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(HTTPHeaderArrayOutput)
}

type HTTPHeaderArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(HTTPHeaderArrayOutput{}) }

func (HTTPHeaderArrayOutput) ElementType() reflect.Type {
	return _HTTPHeaderArrayType
}

func (o HTTPHeaderArrayOutput) Index(i pulumi.IntInput) HTTPHeaderOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) HTTPHeader {
		return vs[0].([]HTTPHeader)[vs[1].(int)]
	}).(HTTPHeaderOutput)
}
