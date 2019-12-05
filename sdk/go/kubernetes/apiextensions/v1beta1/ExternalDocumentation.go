// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ExternalDocumentation allows referencing an external resource for extended documentation.
type ExternalDocumentation struct {
	
	Description *string `pulumi:"description"`

	
	Url *string `pulumi:"url"`

}

var _ExternalDocumentationType = reflect.TypeOf((*ExternalDocumentation)(nil)).Elem()

// ExternalDocumentationInput represents an input type that resolves to a ExternalDocumentation.
type ExternalDocumentationInput interface {
	ElementType() reflect.Type

	ToExternalDocumentationOutput() ExternalDocumentationOutput
	ToExternalDocumentationOutputWithContext(ctx context.Context) ExternalDocumentationOutput
}

// ExternalDocumentationArgs is a ExternalDocumentationInput whose fields are all Input types.
type ExternalDocumentationArgs struct {
	
	Description pulumi.StringInput `pulumi:"description"`

	
	Url pulumi.StringInput `pulumi:"url"`

}

func (a ExternalDocumentationArgs) ElementType() reflect.Type {
	return _ExternalDocumentationType
}

func (a ExternalDocumentationArgs) ToExternalDocumentationOutput() ExternalDocumentationOutput {
	return pulumi.ToOutput(a).(ExternalDocumentationOutput)
}

func (a ExternalDocumentationArgs) ToExternalDocumentationOutputWithContext(ctx context.Context) ExternalDocumentationOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ExternalDocumentationOutput)
}

// ExternalDocumentationOutput is an output type that resolves to a Input.
type ExternalDocumentationOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ExternalDocumentationOutput{}) }

func (ExternalDocumentationOutput) ElementType() reflect.Type {
	return _ExternalDocumentationType
}

func (o ExternalDocumentationOutput) Description() pulumi.StringOutput {
	return o.Apply(func(v ExternalDocumentation) *string {
		return v.Description
	}).(pulumi.StringOutput)
}

func (o ExternalDocumentationOutput) Url() pulumi.StringOutput {
	return o.Apply(func(v ExternalDocumentation) *string {
		return v.Url
	}).(pulumi.StringOutput)
}

