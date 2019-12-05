// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// NetworkPolicyEgressRule describes a particular set of traffic that is allowed out of pods matched
// by a NetworkPolicySpec's podSelector. The traffic must match both ports and to. This type is
// beta-level in 1.8
type NetworkPolicyEgressRule struct {
	// List of destination ports for outgoing traffic. Each item in this list is combined using a
	// logical OR. If this field is empty or missing, this rule matches all ports (traffic not
	// restricted by port). If this field is present and contains at least one item, then this rule
	// allows traffic only if the traffic matches at least one port in the list.
	Ports []NetworkPolicyPort `pulumi:"ports"`

	// List of destinations for outgoing traffic of pods selected for this rule. Items in this list are
	// combined using a logical OR operation. If this field is empty or missing, this rule matches all
	// destinations (traffic not restricted by destination). If this field is present and contains at
	// least one item, this rule allows traffic only if the traffic matches at least one item in the to
	// list.
	To []NetworkPolicyPeer `pulumi:"to"`

}

var _NetworkPolicyEgressRuleType = reflect.TypeOf((*NetworkPolicyEgressRule)(nil)).Elem()

// NetworkPolicyEgressRuleInput represents an input type that resolves to a NetworkPolicyEgressRule.
type NetworkPolicyEgressRuleInput interface {
	ElementType() reflect.Type

	ToNetworkPolicyEgressRuleOutput() NetworkPolicyEgressRuleOutput
	ToNetworkPolicyEgressRuleOutputWithContext(ctx context.Context) NetworkPolicyEgressRuleOutput
}

// NetworkPolicyEgressRuleArgs is a NetworkPolicyEgressRuleInput whose fields are all Input types.
type NetworkPolicyEgressRuleArgs struct {
	// List of destination ports for outgoing traffic. Each item in this list is combined using a
	// logical OR. If this field is empty or missing, this rule matches all ports (traffic not
	// restricted by port). If this field is present and contains at least one item, then this rule
	// allows traffic only if the traffic matches at least one port in the list.
	Ports NetworkPolicyPortArrayInput `pulumi:"ports"`

	// List of destinations for outgoing traffic of pods selected for this rule. Items in this list are
	// combined using a logical OR operation. If this field is empty or missing, this rule matches all
	// destinations (traffic not restricted by destination). If this field is present and contains at
	// least one item, this rule allows traffic only if the traffic matches at least one item in the to
	// list.
	To NetworkPolicyPeerArrayInput `pulumi:"to"`

}

func (a NetworkPolicyEgressRuleArgs) ElementType() reflect.Type {
	return _NetworkPolicyEgressRuleType
}

func (a NetworkPolicyEgressRuleArgs) ToNetworkPolicyEgressRuleOutput() NetworkPolicyEgressRuleOutput {
	return pulumi.ToOutput(a).(NetworkPolicyEgressRuleOutput)
}

func (a NetworkPolicyEgressRuleArgs) ToNetworkPolicyEgressRuleOutputWithContext(ctx context.Context) NetworkPolicyEgressRuleOutput {
	return pulumi.ToOutputWithContext(ctx, a).(NetworkPolicyEgressRuleOutput)
}

// NetworkPolicyEgressRuleOutput is an output type that resolves to a Input.
type NetworkPolicyEgressRuleOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(NetworkPolicyEgressRuleOutput{}) }

func (NetworkPolicyEgressRuleOutput) ElementType() reflect.Type {
	return _NetworkPolicyEgressRuleType
}

func (o NetworkPolicyEgressRuleOutput) Ports() NetworkPolicyPortArrayOutput {
	return o.Apply(func(v NetworkPolicyEgressRule) []NetworkPolicyPort {
		return v.Ports
	}).(NetworkPolicyPortArrayOutput)
}

func (o NetworkPolicyEgressRuleOutput) To() NetworkPolicyPeerArrayOutput {
	return o.Apply(func(v NetworkPolicyEgressRule) []NetworkPolicyPeer {
		return v.To
	}).(NetworkPolicyPeerArrayOutput)
}

var _NetworkPolicyEgressRuleArrayType = reflect.TypeOf((*[]NetworkPolicyEgressRule)(nil)).Elem()

type NetworkPolicyEgressRuleArrayInput interface {
	ElementType() reflect.Type

	ToNetworkPolicyEgressRuleArrayOutput() NetworkPolicyEgressRuleArrayOutput
	ToNetworkPolicyEgressRuleArrayOutputWithContext(ctx context.Context) NetworkPolicyEgressRuleArrayOutput
}

type NetworkPolicyEgressRuleArray []NetworkPolicyEgressRuleInput

func (a NetworkPolicyEgressRuleArray) ElementType() reflect.Type {
	return _NetworkPolicyEgressRuleArrayType
}

func (a NetworkPolicyEgressRuleArray) ToNetworkPolicyEgressRuleArrayOutput() NetworkPolicyEgressRuleArrayOutput {
	return pulumi.ToOutput(a).(NetworkPolicyEgressRuleArrayOutput)
}

func (a NetworkPolicyEgressRuleArray) ToNetworkPolicyEgressRuleArrayOutputWithContext(ctx context.Context) NetworkPolicyEgressRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(NetworkPolicyEgressRuleArrayOutput)
}

type NetworkPolicyEgressRuleArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(NetworkPolicyEgressRuleArrayOutput{}) }

func (NetworkPolicyEgressRuleArrayOutput) ElementType() reflect.Type {
	return _NetworkPolicyEgressRuleArrayType
}

func (o NetworkPolicyEgressRuleArrayOutput) Index(i pulumi.IntInput) NetworkPolicyEgressRuleOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) NetworkPolicyEgressRule {
		return vs[0].([]NetworkPolicyEgressRule)[vs[1].(int)]
	}).(NetworkPolicyEgressRuleOutput)
}