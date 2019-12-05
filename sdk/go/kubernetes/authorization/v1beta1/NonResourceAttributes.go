// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// NonResourceAttributes includes the authorization attributes available for non-resource requests
// to the Authorizer interface
type NonResourceAttributes struct {
	// Path is the URL path of the request
	Path *string `pulumi:"path"`

	// Verb is the standard HTTP verb
	Verb *string `pulumi:"verb"`

}

var _NonResourceAttributesType = reflect.TypeOf((*NonResourceAttributes)(nil)).Elem()

// NonResourceAttributesInput represents an input type that resolves to a NonResourceAttributes.
type NonResourceAttributesInput interface {
	ElementType() reflect.Type

	ToNonResourceAttributesOutput() NonResourceAttributesOutput
	ToNonResourceAttributesOutputWithContext(ctx context.Context) NonResourceAttributesOutput
}

// NonResourceAttributesArgs is a NonResourceAttributesInput whose fields are all Input types.
type NonResourceAttributesArgs struct {
	// Path is the URL path of the request
	Path pulumi.StringInput `pulumi:"path"`

	// Verb is the standard HTTP verb
	Verb pulumi.StringInput `pulumi:"verb"`

}

func (a NonResourceAttributesArgs) ElementType() reflect.Type {
	return _NonResourceAttributesType
}

func (a NonResourceAttributesArgs) ToNonResourceAttributesOutput() NonResourceAttributesOutput {
	return pulumi.ToOutput(a).(NonResourceAttributesOutput)
}

func (a NonResourceAttributesArgs) ToNonResourceAttributesOutputWithContext(ctx context.Context) NonResourceAttributesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(NonResourceAttributesOutput)
}

// NonResourceAttributesOutput is an output type that resolves to a Input.
type NonResourceAttributesOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(NonResourceAttributesOutput{}) }

func (NonResourceAttributesOutput) ElementType() reflect.Type {
	return _NonResourceAttributesType
}

func (o NonResourceAttributesOutput) Path() pulumi.StringOutput {
	return o.Apply(func(v NonResourceAttributes) *string {
		return v.Path
	}).(pulumi.StringOutput)
}

func (o NonResourceAttributesOutput) Verb() pulumi.StringOutput {
	return o.Apply(func(v NonResourceAttributes) *string {
		return v.Verb
	}).(pulumi.StringOutput)
}

