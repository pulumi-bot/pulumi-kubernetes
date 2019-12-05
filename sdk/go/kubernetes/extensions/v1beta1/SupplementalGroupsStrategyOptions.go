// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// SupplementalGroupsStrategyOptions defines the strategy type and options used to create the
// strategy. Deprecated: use SupplementalGroupsStrategyOptions from policy API Group instead.
type SupplementalGroupsStrategyOptions struct {
	// ranges are the allowed ranges of supplemental groups.  If you would like to force a single
	// supplemental group then supply a single range with the same start and end. Required for
	// MustRunAs.
	Ranges []IDRange `pulumi:"ranges"`

	// rule is the strategy that will dictate what supplemental groups is used in the SecurityContext.
	Rule *string `pulumi:"rule"`

}

var _SupplementalGroupsStrategyOptionsType = reflect.TypeOf((*SupplementalGroupsStrategyOptions)(nil)).Elem()

// SupplementalGroupsStrategyOptionsInput represents an input type that resolves to a SupplementalGroupsStrategyOptions.
type SupplementalGroupsStrategyOptionsInput interface {
	ElementType() reflect.Type

	ToSupplementalGroupsStrategyOptionsOutput() SupplementalGroupsStrategyOptionsOutput
	ToSupplementalGroupsStrategyOptionsOutputWithContext(ctx context.Context) SupplementalGroupsStrategyOptionsOutput
}

// SupplementalGroupsStrategyOptionsArgs is a SupplementalGroupsStrategyOptionsInput whose fields are all Input types.
type SupplementalGroupsStrategyOptionsArgs struct {
	// ranges are the allowed ranges of supplemental groups.  If you would like to force a single
	// supplemental group then supply a single range with the same start and end. Required for
	// MustRunAs.
	Ranges IDRangeArrayInput `pulumi:"ranges"`

	// rule is the strategy that will dictate what supplemental groups is used in the SecurityContext.
	Rule pulumi.StringInput `pulumi:"rule"`

}

func (a SupplementalGroupsStrategyOptionsArgs) ElementType() reflect.Type {
	return _SupplementalGroupsStrategyOptionsType
}

func (a SupplementalGroupsStrategyOptionsArgs) ToSupplementalGroupsStrategyOptionsOutput() SupplementalGroupsStrategyOptionsOutput {
	return pulumi.ToOutput(a).(SupplementalGroupsStrategyOptionsOutput)
}

func (a SupplementalGroupsStrategyOptionsArgs) ToSupplementalGroupsStrategyOptionsOutputWithContext(ctx context.Context) SupplementalGroupsStrategyOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(SupplementalGroupsStrategyOptionsOutput)
}

// SupplementalGroupsStrategyOptionsOutput is an output type that resolves to a Input.
type SupplementalGroupsStrategyOptionsOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(SupplementalGroupsStrategyOptionsOutput{}) }

func (SupplementalGroupsStrategyOptionsOutput) ElementType() reflect.Type {
	return _SupplementalGroupsStrategyOptionsType
}

func (o SupplementalGroupsStrategyOptionsOutput) Ranges() IDRangeArrayOutput {
	return o.Apply(func(v SupplementalGroupsStrategyOptions) []IDRange {
		return v.Ranges
	}).(IDRangeArrayOutput)
}

func (o SupplementalGroupsStrategyOptionsOutput) Rule() pulumi.StringOutput {
	return o.Apply(func(v SupplementalGroupsStrategyOptions) *string {
		return v.Rule
	}).(pulumi.StringOutput)
}

