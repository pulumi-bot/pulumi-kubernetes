// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// A scope selector represents the AND of the selectors represented by the scoped-resource selector
// requirements.
type ScopeSelector struct {
	// A list of scope selector requirements by scope of the resources.
	MatchExpressions []ScopedResourceSelectorRequirement `pulumi:"matchExpressions"`

}

var _ScopeSelectorType = reflect.TypeOf((*ScopeSelector)(nil)).Elem()

// ScopeSelectorInput represents an input type that resolves to a ScopeSelector.
type ScopeSelectorInput interface {
	ElementType() reflect.Type

	ToScopeSelectorOutput() ScopeSelectorOutput
	ToScopeSelectorOutputWithContext(ctx context.Context) ScopeSelectorOutput
}

// ScopeSelectorArgs is a ScopeSelectorInput whose fields are all Input types.
type ScopeSelectorArgs struct {
	// A list of scope selector requirements by scope of the resources.
	MatchExpressions ScopedResourceSelectorRequirementArrayInput `pulumi:"matchExpressions"`

}

func (a ScopeSelectorArgs) ElementType() reflect.Type {
	return _ScopeSelectorType
}

func (a ScopeSelectorArgs) ToScopeSelectorOutput() ScopeSelectorOutput {
	return pulumi.ToOutput(a).(ScopeSelectorOutput)
}

func (a ScopeSelectorArgs) ToScopeSelectorOutputWithContext(ctx context.Context) ScopeSelectorOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ScopeSelectorOutput)
}

// ScopeSelectorOutput is an output type that resolves to a Input.
type ScopeSelectorOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ScopeSelectorOutput{}) }

func (ScopeSelectorOutput) ElementType() reflect.Type {
	return _ScopeSelectorType
}

func (o ScopeSelectorOutput) MatchExpressions() ScopedResourceSelectorRequirementArrayOutput {
	return o.Apply(func(v ScopeSelector) []ScopedResourceSelectorRequirement {
		return v.MatchExpressions
	}).(ScopedResourceSelectorRequirementArrayOutput)
}

