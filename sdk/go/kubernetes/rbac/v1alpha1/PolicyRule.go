// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// PolicyRule holds information that describes a policy rule, but does not contain information about
// who the rule applies to or which namespace the rule applies to.
type PolicyRule struct {
	// APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are
	// specified, any action requested against one of the enumerated resources in any API group will be
	// allowed.
	ApiGroups []string `pulumi:"apiGroups"`

	// NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but
	// only as the full, final step in the path This name is intentionally different than the internal
	// type so that the DefaultConvert works nicely and because the ordering may be different. Since
	// non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced
	// from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets")
	// or non-resource URL paths (such as "/api"),  but not both.
	NonResourceURLs []string `pulumi:"nonResourceURLs"`

	// ResourceNames is an optional white list of names that the rule applies to.  An empty set means
	// that everything is allowed.
	ResourceNames []string `pulumi:"resourceNames"`

	// Resources is a list of resources this rule applies to.  ResourceAll represents all resources.
	Resources []string `pulumi:"resources"`

	// Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained
	// in this rule.  VerbAll represents all kinds.
	Verbs []string `pulumi:"verbs"`

}

var _PolicyRuleType = reflect.TypeOf((*PolicyRule)(nil)).Elem()

// PolicyRuleInput represents an input type that resolves to a PolicyRule.
type PolicyRuleInput interface {
	ElementType() reflect.Type

	ToPolicyRuleOutput() PolicyRuleOutput
	ToPolicyRuleOutputWithContext(ctx context.Context) PolicyRuleOutput
}

// PolicyRuleArgs is a PolicyRuleInput whose fields are all Input types.
type PolicyRuleArgs struct {
	// Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained
	// in this rule.  VerbAll represents all kinds.
	Verbs pulumi.StringArrayInput `pulumi:"verbs"`

	// APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are
	// specified, any action requested against one of the enumerated resources in any API group will be
	// allowed.
	ApiGroups pulumi.StringArrayInput `pulumi:"apiGroups"`

	// NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but
	// only as the full, final step in the path This name is intentionally different than the internal
	// type so that the DefaultConvert works nicely and because the ordering may be different. Since
	// non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced
	// from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets")
	// or non-resource URL paths (such as "/api"),  but not both.
	NonResourceURLs pulumi.StringArrayInput `pulumi:"nonResourceURLs"`

	// ResourceNames is an optional white list of names that the rule applies to.  An empty set means
	// that everything is allowed.
	ResourceNames pulumi.StringArrayInput `pulumi:"resourceNames"`

	// Resources is a list of resources this rule applies to.  ResourceAll represents all resources.
	Resources pulumi.StringArrayInput `pulumi:"resources"`

}

func (a PolicyRuleArgs) ElementType() reflect.Type {
	return _PolicyRuleType
}

func (a PolicyRuleArgs) ToPolicyRuleOutput() PolicyRuleOutput {
	return pulumi.ToOutput(a).(PolicyRuleOutput)
}

func (a PolicyRuleArgs) ToPolicyRuleOutputWithContext(ctx context.Context) PolicyRuleOutput {
	return pulumi.ToOutputWithContext(ctx, a).(PolicyRuleOutput)
}

// PolicyRuleOutput is an output type that resolves to a Input.
type PolicyRuleOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(PolicyRuleOutput{}) }

func (PolicyRuleOutput) ElementType() reflect.Type {
	return _PolicyRuleType
}

func (o PolicyRuleOutput) ApiGroups() pulumi.StringArrayOutput {
	return o.Apply(func(v PolicyRule) []string {
		return v.ApiGroups
	}).(pulumi.StringArrayOutput)
}

func (o PolicyRuleOutput) NonResourceURLs() pulumi.StringArrayOutput {
	return o.Apply(func(v PolicyRule) []string {
		return v.NonResourceURLs
	}).(pulumi.StringArrayOutput)
}

func (o PolicyRuleOutput) ResourceNames() pulumi.StringArrayOutput {
	return o.Apply(func(v PolicyRule) []string {
		return v.ResourceNames
	}).(pulumi.StringArrayOutput)
}

func (o PolicyRuleOutput) Resources() pulumi.StringArrayOutput {
	return o.Apply(func(v PolicyRule) []string {
		return v.Resources
	}).(pulumi.StringArrayOutput)
}

func (o PolicyRuleOutput) Verbs() pulumi.StringArrayOutput {
	return o.Apply(func(v PolicyRule) []string {
		return v.Verbs
	}).(pulumi.StringArrayOutput)
}

var _PolicyRuleArrayType = reflect.TypeOf((*[]PolicyRule)(nil)).Elem()

type PolicyRuleArrayInput interface {
	ElementType() reflect.Type

	ToPolicyRuleArrayOutput() PolicyRuleArrayOutput
	ToPolicyRuleArrayOutputWithContext(ctx context.Context) PolicyRuleArrayOutput
}

type PolicyRuleArray []PolicyRuleInput

func (a PolicyRuleArray) ElementType() reflect.Type {
	return _PolicyRuleArrayType
}

func (a PolicyRuleArray) ToPolicyRuleArrayOutput() PolicyRuleArrayOutput {
	return pulumi.ToOutput(a).(PolicyRuleArrayOutput)
}

func (a PolicyRuleArray) ToPolicyRuleArrayOutputWithContext(ctx context.Context) PolicyRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(PolicyRuleArrayOutput)
}

type PolicyRuleArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(PolicyRuleArrayOutput{}) }

func (PolicyRuleArrayOutput) ElementType() reflect.Type {
	return _PolicyRuleArrayType
}

func (o PolicyRuleArrayOutput) Index(i pulumi.IntInput) PolicyRuleOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) PolicyRule {
		return vs[0].([]PolicyRule)[vs[1].(int)]
	}).(PolicyRuleOutput)
}
