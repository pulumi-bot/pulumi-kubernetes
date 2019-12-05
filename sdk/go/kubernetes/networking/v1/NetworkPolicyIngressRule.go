// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// NetworkPolicyIngressRule describes a particular set of traffic that is allowed to the pods
// matched by a NetworkPolicySpec's podSelector. The traffic must match both ports and from.
type NetworkPolicyIngressRule struct {
	// List of sources which should be able to access the pods selected for this rule. Items in this
	// list are combined using a logical OR operation. If this field is empty or missing, this rule
	// matches all sources (traffic not restricted by source). If this field is present and contains at
	// least one item, this rule allows traffic only if the traffic matches at least one item in the
	// from list.
	From []NetworkPolicyPeer `pulumi:"from"`

	// List of ports which should be made accessible on the pods selected for this rule. Each item in
	// this list is combined using a logical OR. If this field is empty or missing, this rule matches
	// all ports (traffic not restricted by port). If this field is present and contains at least one
	// item, then this rule allows traffic only if the traffic matches at least one port in the list.
	Ports []NetworkPolicyPort `pulumi:"ports"`

}

var _NetworkPolicyIngressRuleType = reflect.TypeOf((*NetworkPolicyIngressRule)(nil)).Elem()

// NetworkPolicyIngressRuleInput represents an input type that resolves to a NetworkPolicyIngressRule.
type NetworkPolicyIngressRuleInput interface {
	ElementType() reflect.Type

	ToNetworkPolicyIngressRuleOutput() NetworkPolicyIngressRuleOutput
	ToNetworkPolicyIngressRuleOutputWithContext(ctx context.Context) NetworkPolicyIngressRuleOutput
}

// NetworkPolicyIngressRuleArgs is a NetworkPolicyIngressRuleInput whose fields are all Input types.
type NetworkPolicyIngressRuleArgs struct {
	// List of sources which should be able to access the pods selected for this rule. Items in this
	// list are combined using a logical OR operation. If this field is empty or missing, this rule
	// matches all sources (traffic not restricted by source). If this field is present and contains at
	// least one item, this rule allows traffic only if the traffic matches at least one item in the
	// from list.
	From NetworkPolicyPeerArrayInput `pulumi:"from"`

	// List of ports which should be made accessible on the pods selected for this rule. Each item in
	// this list is combined using a logical OR. If this field is empty or missing, this rule matches
	// all ports (traffic not restricted by port). If this field is present and contains at least one
	// item, then this rule allows traffic only if the traffic matches at least one port in the list.
	Ports NetworkPolicyPortArrayInput `pulumi:"ports"`

}

func (a NetworkPolicyIngressRuleArgs) ElementType() reflect.Type {
	return _NetworkPolicyIngressRuleType
}

func (a NetworkPolicyIngressRuleArgs) ToNetworkPolicyIngressRuleOutput() NetworkPolicyIngressRuleOutput {
	return pulumi.ToOutput(a).(NetworkPolicyIngressRuleOutput)
}

func (a NetworkPolicyIngressRuleArgs) ToNetworkPolicyIngressRuleOutputWithContext(ctx context.Context) NetworkPolicyIngressRuleOutput {
	return pulumi.ToOutputWithContext(ctx, a).(NetworkPolicyIngressRuleOutput)
}

// NetworkPolicyIngressRuleOutput is an output type that resolves to a Input.
type NetworkPolicyIngressRuleOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(NetworkPolicyIngressRuleOutput{}) }

func (NetworkPolicyIngressRuleOutput) ElementType() reflect.Type {
	return _NetworkPolicyIngressRuleType
}

func (o NetworkPolicyIngressRuleOutput) From() NetworkPolicyPeerArrayOutput {
	return o.Apply(func(v NetworkPolicyIngressRule) []NetworkPolicyPeer {
		return v.From
	}).(NetworkPolicyPeerArrayOutput)
}

func (o NetworkPolicyIngressRuleOutput) Ports() NetworkPolicyPortArrayOutput {
	return o.Apply(func(v NetworkPolicyIngressRule) []NetworkPolicyPort {
		return v.Ports
	}).(NetworkPolicyPortArrayOutput)
}

var _NetworkPolicyIngressRuleArrayType = reflect.TypeOf((*[]NetworkPolicyIngressRule)(nil)).Elem()

type NetworkPolicyIngressRuleArrayInput interface {
	ElementType() reflect.Type

	ToNetworkPolicyIngressRuleArrayOutput() NetworkPolicyIngressRuleArrayOutput
	ToNetworkPolicyIngressRuleArrayOutputWithContext(ctx context.Context) NetworkPolicyIngressRuleArrayOutput
}

type NetworkPolicyIngressRuleArray []NetworkPolicyIngressRuleInput

func (a NetworkPolicyIngressRuleArray) ElementType() reflect.Type {
	return _NetworkPolicyIngressRuleArrayType
}

func (a NetworkPolicyIngressRuleArray) ToNetworkPolicyIngressRuleArrayOutput() NetworkPolicyIngressRuleArrayOutput {
	return pulumi.ToOutput(a).(NetworkPolicyIngressRuleArrayOutput)
}

func (a NetworkPolicyIngressRuleArray) ToNetworkPolicyIngressRuleArrayOutputWithContext(ctx context.Context) NetworkPolicyIngressRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(NetworkPolicyIngressRuleArrayOutput)
}

type NetworkPolicyIngressRuleArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(NetworkPolicyIngressRuleArrayOutput{}) }

func (NetworkPolicyIngressRuleArrayOutput) ElementType() reflect.Type {
	return _NetworkPolicyIngressRuleArrayType
}

func (o NetworkPolicyIngressRuleArrayOutput) Index(i pulumi.IntInput) NetworkPolicyIngressRuleOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) NetworkPolicyIngressRule {
		return vs[0].([]NetworkPolicyIngressRule)[vs[1].(int)]
	}).(NetworkPolicyIngressRuleOutput)
}
