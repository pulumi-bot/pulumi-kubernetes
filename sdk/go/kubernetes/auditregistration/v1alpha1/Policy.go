// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Policy defines the configuration of how audit events are logged
type Policy struct {
	// The Level that all requests are recorded at. available options: None, Metadata, Request,
	// RequestResponse required
	Level string `pulumi:"level"`

	// Stages is a list of stages for which events are created.
	Stages []string `pulumi:"stages"`

}

var _PolicyType = reflect.TypeOf((*Policy)(nil)).Elem()

// PolicyInput represents an input type that resolves to a Policy.
type PolicyInput interface {
	ElementType() reflect.Type

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

// PolicyArgs is a PolicyInput whose fields are all Input types.
type PolicyArgs struct {
	// The Level that all requests are recorded at. available options: None, Metadata, Request,
	// RequestResponse required
	Level pulumi.StringInput `pulumi:"level"`

	// Stages is a list of stages for which events are created.
	Stages pulumi.StringArrayInput `pulumi:"stages"`

}

func (a PolicyArgs) ElementType() reflect.Type {
	return _PolicyType
}

func (a PolicyArgs) ToPolicyOutput() PolicyOutput {
	return pulumi.ToOutput(a).(PolicyOutput)
}

func (a PolicyArgs) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, a).(PolicyOutput)
}

// PolicyOutput is an output type that resolves to a Input.
type PolicyOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(PolicyOutput{}) }

func (PolicyOutput) ElementType() reflect.Type {
	return _PolicyType
}

func (o PolicyOutput) Level() pulumi.StringOutput {
	return o.Apply(func(v Policy) string {
		return v.Level
	}).(pulumi.StringOutput)
}

func (o PolicyOutput) Stages() pulumi.StringArrayOutput {
	return o.Apply(func(v Policy) []string {
		return v.Stages
	}).(pulumi.StringArrayOutput)
}

