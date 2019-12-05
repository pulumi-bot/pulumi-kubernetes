// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// FSGroupStrategyOptions defines the strategy type and options used to create the strategy.
type FSGroupStrategyOptions struct {
	// ranges are the allowed ranges of fs groups.  If you would like to force a single fs group then
	// supply a single range with the same start and end. Required for MustRunAs.
	Ranges []IDRange `pulumi:"ranges"`

	// rule is the strategy that will dictate what FSGroup is used in the SecurityContext.
	Rule *string `pulumi:"rule"`

}

var _FSGroupStrategyOptionsType = reflect.TypeOf((*FSGroupStrategyOptions)(nil)).Elem()

// FSGroupStrategyOptionsInput represents an input type that resolves to a FSGroupStrategyOptions.
type FSGroupStrategyOptionsInput interface {
	ElementType() reflect.Type

	ToFSGroupStrategyOptionsOutput() FSGroupStrategyOptionsOutput
	ToFSGroupStrategyOptionsOutputWithContext(ctx context.Context) FSGroupStrategyOptionsOutput
}

// FSGroupStrategyOptionsArgs is a FSGroupStrategyOptionsInput whose fields are all Input types.
type FSGroupStrategyOptionsArgs struct {
	// ranges are the allowed ranges of fs groups.  If you would like to force a single fs group then
	// supply a single range with the same start and end. Required for MustRunAs.
	Ranges IDRangeArrayInput `pulumi:"ranges"`

	// rule is the strategy that will dictate what FSGroup is used in the SecurityContext.
	Rule pulumi.StringInput `pulumi:"rule"`

}

func (a FSGroupStrategyOptionsArgs) ElementType() reflect.Type {
	return _FSGroupStrategyOptionsType
}

func (a FSGroupStrategyOptionsArgs) ToFSGroupStrategyOptionsOutput() FSGroupStrategyOptionsOutput {
	return pulumi.ToOutput(a).(FSGroupStrategyOptionsOutput)
}

func (a FSGroupStrategyOptionsArgs) ToFSGroupStrategyOptionsOutputWithContext(ctx context.Context) FSGroupStrategyOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(FSGroupStrategyOptionsOutput)
}

// FSGroupStrategyOptionsOutput is an output type that resolves to a Input.
type FSGroupStrategyOptionsOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(FSGroupStrategyOptionsOutput{}) }

func (FSGroupStrategyOptionsOutput) ElementType() reflect.Type {
	return _FSGroupStrategyOptionsType
}

func (o FSGroupStrategyOptionsOutput) Ranges() IDRangeArrayOutput {
	return o.Apply(func(v FSGroupStrategyOptions) []IDRange {
		return v.Ranges
	}).(IDRangeArrayOutput)
}

func (o FSGroupStrategyOptionsOutput) Rule() pulumi.StringOutput {
	return o.Apply(func(v FSGroupStrategyOptions) *string {
		return v.Rule
	}).(pulumi.StringOutput)
}

