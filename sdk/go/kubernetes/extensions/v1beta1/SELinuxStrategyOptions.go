// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	coreV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/core/v1"
)

// SELinuxStrategyOptions defines the strategy type and any options used to create the strategy.
// Deprecated: use SELinuxStrategyOptions from policy API Group instead.
type SELinuxStrategyOptions struct {
	// rule is the strategy that will dictate the allowable labels that may be set.
	Rule string `pulumi:"rule"`

	// seLinuxOptions required to run as; required for MustRunAs More info:
	// https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	SeLinuxOptions *coreV1.SELinuxOptions `pulumi:"seLinuxOptions"`

}

var _SELinuxStrategyOptionsType = reflect.TypeOf((*SELinuxStrategyOptions)(nil)).Elem()

// SELinuxStrategyOptionsInput represents an input type that resolves to a SELinuxStrategyOptions.
type SELinuxStrategyOptionsInput interface {
	ElementType() reflect.Type

	ToSELinuxStrategyOptionsOutput() SELinuxStrategyOptionsOutput
	ToSELinuxStrategyOptionsOutputWithContext(ctx context.Context) SELinuxStrategyOptionsOutput
}

// SELinuxStrategyOptionsArgs is a SELinuxStrategyOptionsInput whose fields are all Input types.
type SELinuxStrategyOptionsArgs struct {
	// rule is the strategy that will dictate the allowable labels that may be set.
	Rule pulumi.StringInput `pulumi:"rule"`

	// seLinuxOptions required to run as; required for MustRunAs More info:
	// https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	SeLinuxOptions coreV1.SELinuxOptionsInput `pulumi:"seLinuxOptions"`

}

func (a SELinuxStrategyOptionsArgs) ElementType() reflect.Type {
	return _SELinuxStrategyOptionsType
}

func (a SELinuxStrategyOptionsArgs) ToSELinuxStrategyOptionsOutput() SELinuxStrategyOptionsOutput {
	return pulumi.ToOutput(a).(SELinuxStrategyOptionsOutput)
}

func (a SELinuxStrategyOptionsArgs) ToSELinuxStrategyOptionsOutputWithContext(ctx context.Context) SELinuxStrategyOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(SELinuxStrategyOptionsOutput)
}

// SELinuxStrategyOptionsOutput is an output type that resolves to a Input.
type SELinuxStrategyOptionsOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(SELinuxStrategyOptionsOutput{}) }

func (SELinuxStrategyOptionsOutput) ElementType() reflect.Type {
	return _SELinuxStrategyOptionsType
}

func (o SELinuxStrategyOptionsOutput) Rule() pulumi.StringOutput {
	return o.Apply(func(v SELinuxStrategyOptions) string {
		return v.Rule
	}).(pulumi.StringOutput)
}

func (o SELinuxStrategyOptionsOutput) SeLinuxOptions() coreV1.SELinuxOptionsOutput {
	return o.Apply(func(v SELinuxStrategyOptions) *coreV1.SELinuxOptions {
		return v.SeLinuxOptions
	}).(coreV1.SELinuxOptionsOutput)
}

