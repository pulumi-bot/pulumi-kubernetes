// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// DEPRECATED 1.9 - This group version of NetworkPolicySpec is deprecated by
// networking/v1/NetworkPolicySpec.
type NetworkPolicySpec struct {
	// List of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there
	// are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR
	// if the traffic matches at least one egress rule across all of the NetworkPolicy objects whose
	// podSelector matches the pod. If this field is empty then this NetworkPolicy limits all outgoing
	// traffic (and serves solely to ensure that the pods it selects are isolated by default). This
	// field is beta-level in 1.8
	Egress []NetworkPolicyEgressRule `pulumi:"egress"`

	// List of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there
	// are no NetworkPolicies selecting the pod OR if the traffic source is the pod's local node, OR if
	// the traffic matches at least one ingress rule across all of the NetworkPolicy objects whose
	// podSelector matches the pod. If this field is empty then this NetworkPolicy does not allow any
	// traffic (and serves solely to ensure that the pods it selects are isolated by default).
	Ingress []NetworkPolicyIngressRule `pulumi:"ingress"`

	// Selects the pods to which this NetworkPolicy object applies.  The array of ingress rules is
	// applied to any pods selected by this field. Multiple network policies can select the same set of
	// pods.  In this case, the ingress rules for each are combined additively. This field is NOT
	// optional and follows standard label selector semantics. An empty podSelector matches all pods in
	// this namespace.
	PodSelector metaV1.LabelSelector `pulumi:"podSelector"`

	// List of rule types that the NetworkPolicy relates to. Valid options are "Ingress", "Egress", or
	// "Ingress,Egress". If this field is not specified, it will default based on the existence of
	// Ingress or Egress rules; policies that contain an Egress section are assumed to affect Egress,
	// and all policies (whether or not they contain an Ingress section) are assumed to affect Ingress.
	// If you want to write an egress-only policy, you must explicitly specify policyTypes [ "Egress"
	// ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must
	// specify a policyTypes value that include "Egress" (since such a policy would not include an
	// Egress section and would otherwise default to just [ "Ingress" ]). This field is beta-level in
	// 1.8
	PolicyTypes []string `pulumi:"policyTypes"`

}

var _NetworkPolicySpecType = reflect.TypeOf((*NetworkPolicySpec)(nil)).Elem()

// NetworkPolicySpecInput represents an input type that resolves to a NetworkPolicySpec.
type NetworkPolicySpecInput interface {
	ElementType() reflect.Type

	ToNetworkPolicySpecOutput() NetworkPolicySpecOutput
	ToNetworkPolicySpecOutputWithContext(ctx context.Context) NetworkPolicySpecOutput
}

// NetworkPolicySpecArgs is a NetworkPolicySpecInput whose fields are all Input types.
type NetworkPolicySpecArgs struct {
	// Selects the pods to which this NetworkPolicy object applies.  The array of ingress rules is
	// applied to any pods selected by this field. Multiple network policies can select the same set of
	// pods.  In this case, the ingress rules for each are combined additively. This field is NOT
	// optional and follows standard label selector semantics. An empty podSelector matches all pods in
	// this namespace.
	PodSelector metaV1.LabelSelectorInput `pulumi:"podSelector"`

	// List of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there
	// are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR
	// if the traffic matches at least one egress rule across all of the NetworkPolicy objects whose
	// podSelector matches the pod. If this field is empty then this NetworkPolicy limits all outgoing
	// traffic (and serves solely to ensure that the pods it selects are isolated by default). This
	// field is beta-level in 1.8
	Egress NetworkPolicyEgressRuleArrayInput `pulumi:"egress"`

	// List of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there
	// are no NetworkPolicies selecting the pod OR if the traffic source is the pod's local node, OR if
	// the traffic matches at least one ingress rule across all of the NetworkPolicy objects whose
	// podSelector matches the pod. If this field is empty then this NetworkPolicy does not allow any
	// traffic (and serves solely to ensure that the pods it selects are isolated by default).
	Ingress NetworkPolicyIngressRuleArrayInput `pulumi:"ingress"`

	// List of rule types that the NetworkPolicy relates to. Valid options are "Ingress", "Egress", or
	// "Ingress,Egress". If this field is not specified, it will default based on the existence of
	// Ingress or Egress rules; policies that contain an Egress section are assumed to affect Egress,
	// and all policies (whether or not they contain an Ingress section) are assumed to affect Ingress.
	// If you want to write an egress-only policy, you must explicitly specify policyTypes [ "Egress"
	// ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must
	// specify a policyTypes value that include "Egress" (since such a policy would not include an
	// Egress section and would otherwise default to just [ "Ingress" ]). This field is beta-level in
	// 1.8
	PolicyTypes pulumi.StringArrayInput `pulumi:"policyTypes"`

}

func (a NetworkPolicySpecArgs) ElementType() reflect.Type {
	return _NetworkPolicySpecType
}

func (a NetworkPolicySpecArgs) ToNetworkPolicySpecOutput() NetworkPolicySpecOutput {
	return pulumi.ToOutput(a).(NetworkPolicySpecOutput)
}

func (a NetworkPolicySpecArgs) ToNetworkPolicySpecOutputWithContext(ctx context.Context) NetworkPolicySpecOutput {
	return pulumi.ToOutputWithContext(ctx, a).(NetworkPolicySpecOutput)
}

// NetworkPolicySpecOutput is an output type that resolves to a Input.
type NetworkPolicySpecOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(NetworkPolicySpecOutput{}) }

func (NetworkPolicySpecOutput) ElementType() reflect.Type {
	return _NetworkPolicySpecType
}

func (o NetworkPolicySpecOutput) Egress() NetworkPolicyEgressRuleArrayOutput {
	return o.Apply(func(v NetworkPolicySpec) []NetworkPolicyEgressRule {
		return v.Egress
	}).(NetworkPolicyEgressRuleArrayOutput)
}

func (o NetworkPolicySpecOutput) Ingress() NetworkPolicyIngressRuleArrayOutput {
	return o.Apply(func(v NetworkPolicySpec) []NetworkPolicyIngressRule {
		return v.Ingress
	}).(NetworkPolicyIngressRuleArrayOutput)
}

func (o NetworkPolicySpecOutput) PodSelector() metaV1.LabelSelectorOutput {
	return o.Apply(func(v NetworkPolicySpec) metaV1.LabelSelector {
		return v.PodSelector
	}).(metaV1.LabelSelectorOutput)
}

func (o NetworkPolicySpecOutput) PolicyTypes() pulumi.StringArrayOutput {
	return o.Apply(func(v NetworkPolicySpec) []string {
		return v.PolicyTypes
	}).(pulumi.StringArrayOutput)
}

