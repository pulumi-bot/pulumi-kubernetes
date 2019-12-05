// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ScaleSpec describes the attributes of a scale subresource
type ScaleSpec struct {
	// desired number of instances for the scaled object.
	Replicas *int `pulumi:"replicas"`

}

var _ScaleSpecType = reflect.TypeOf((*ScaleSpec)(nil)).Elem()

// ScaleSpecInput represents an input type that resolves to a ScaleSpec.
type ScaleSpecInput interface {
	ElementType() reflect.Type

	ToScaleSpecOutput() ScaleSpecOutput
	ToScaleSpecOutputWithContext(ctx context.Context) ScaleSpecOutput
}

// ScaleSpecArgs is a ScaleSpecInput whose fields are all Input types.
type ScaleSpecArgs struct {
	// desired number of instances for the scaled object.
	Replicas pulumi.IntInput `pulumi:"replicas"`

}

func (a ScaleSpecArgs) ElementType() reflect.Type {
	return _ScaleSpecType
}

func (a ScaleSpecArgs) ToScaleSpecOutput() ScaleSpecOutput {
	return pulumi.ToOutput(a).(ScaleSpecOutput)
}

func (a ScaleSpecArgs) ToScaleSpecOutputWithContext(ctx context.Context) ScaleSpecOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ScaleSpecOutput)
}

// ScaleSpecOutput is an output type that resolves to a Input.
type ScaleSpecOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ScaleSpecOutput{}) }

func (ScaleSpecOutput) ElementType() reflect.Type {
	return _ScaleSpecType
}

func (o ScaleSpecOutput) Replicas() pulumi.IntOutput {
	return o.Apply(func(v ScaleSpec) *int {
		return v.Replicas
	}).(pulumi.IntOutput)
}

