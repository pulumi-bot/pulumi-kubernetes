// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// CustomResourceConversion describes how to convert different versions of a CR.
type CustomResourceConversion struct {
	// strategy specifies how custom resources are converted between versions. Allowed values are: -
	// `None`: The converter only change the apiVersion and would not touch any other field in the
	// custom resource. - `Webhook`: API Server will call to an external webhook to do the conversion.
	// Additional information
	//   is needed for this option. This requires spec.preserveUnknownFields to be false, and
	// spec.conversion.webhook to be set.
	Strategy string `pulumi:"strategy"`

	// webhook describes how to call the conversion webhook. Required when `strategy` is set to
	// `Webhook`.
	Webhook *WebhookConversion `pulumi:"webhook"`

}

var _CustomResourceConversionType = reflect.TypeOf((*CustomResourceConversion)(nil)).Elem()

// CustomResourceConversionInput represents an input type that resolves to a CustomResourceConversion.
type CustomResourceConversionInput interface {
	ElementType() reflect.Type

	ToCustomResourceConversionOutput() CustomResourceConversionOutput
	ToCustomResourceConversionOutputWithContext(ctx context.Context) CustomResourceConversionOutput
}

// CustomResourceConversionArgs is a CustomResourceConversionInput whose fields are all Input types.
type CustomResourceConversionArgs struct {
	// strategy specifies how custom resources are converted between versions. Allowed values are: -
	// `None`: The converter only change the apiVersion and would not touch any other field in the
	// custom resource. - `Webhook`: API Server will call to an external webhook to do the conversion.
	// Additional information
	//   is needed for this option. This requires spec.preserveUnknownFields to be false, and
	// spec.conversion.webhook to be set.
	Strategy pulumi.StringInput `pulumi:"strategy"`

	// webhook describes how to call the conversion webhook. Required when `strategy` is set to
	// `Webhook`.
	Webhook WebhookConversionInput `pulumi:"webhook"`

}

func (a CustomResourceConversionArgs) ElementType() reflect.Type {
	return _CustomResourceConversionType
}

func (a CustomResourceConversionArgs) ToCustomResourceConversionOutput() CustomResourceConversionOutput {
	return pulumi.ToOutput(a).(CustomResourceConversionOutput)
}

func (a CustomResourceConversionArgs) ToCustomResourceConversionOutputWithContext(ctx context.Context) CustomResourceConversionOutput {
	return pulumi.ToOutputWithContext(ctx, a).(CustomResourceConversionOutput)
}

// CustomResourceConversionOutput is an output type that resolves to a Input.
type CustomResourceConversionOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(CustomResourceConversionOutput{}) }

func (CustomResourceConversionOutput) ElementType() reflect.Type {
	return _CustomResourceConversionType
}

func (o CustomResourceConversionOutput) Strategy() pulumi.StringOutput {
	return o.Apply(func(v CustomResourceConversion) string {
		return v.Strategy
	}).(pulumi.StringOutput)
}

func (o CustomResourceConversionOutput) Webhook() WebhookConversionOutput {
	return o.Apply(func(v CustomResourceConversion) *WebhookConversion {
		return v.Webhook
	}).(WebhookConversionOutput)
}

