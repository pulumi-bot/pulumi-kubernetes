// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// MutatingWebhook describes an admission webhook and the resources and operations it applies to.
type MutatingWebhook struct {
	// AdmissionReviewVersions is an ordered list of preferred `AdmissionReview` versions the Webhook
	// expects. API server will try to use first version in the list which it supports. If none of the
	// versions specified in this list supported by API server, validation will fail for this object.
	// If a persisted webhook configuration specifies allowed versions and does not include any
	// versions known to the API Server, calls to the webhook will fail and be subject to the failure
	// policy. Default to `['v1beta1']`.
	AdmissionReviewVersions []string `pulumi:"admissionReviewVersions"`

	// ClientConfig defines how to communicate with the hook. Required
	ClientConfig WebhookClientConfig `pulumi:"clientConfig"`

	// FailurePolicy defines how unrecognized errors from the admission endpoint are handled - allowed
	// values are Ignore or Fail. Defaults to Ignore.
	FailurePolicy *string `pulumi:"failurePolicy"`

	// matchPolicy defines how the "rules" list is used to match incoming requests. Allowed values are
	// "Exact" or "Equivalent".
	// 
	// - Exact: match a request only if it exactly matches a specified rule. For example, if
	// deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, but "rules" only
	// included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`, a request to
	// apps/v1beta1 or extensions/v1beta1 would not be sent to the webhook.
	// 
	// - Equivalent: match a request if modifies a resource listed in rules, even via another API group
	// or version. For example, if deployments can be modified via apps/v1, apps/v1beta1, and
	// extensions/v1beta1, and "rules" only included `apiGroups:["apps"], apiVersions:["v1"],
	// resources: ["deployments"]`, a request to apps/v1beta1 or extensions/v1beta1 would be converted
	// to apps/v1 and sent to the webhook.
	// 
	// Defaults to "Exact"
	MatchPolicy *string `pulumi:"matchPolicy"`

	// The name of the admission webhook. Name should be fully qualified, e.g.,
	// imagepolicy.kubernetes.io, where "imagepolicy" is the name of the webhook, and kubernetes.io is
	// the name of the organization. Required.
	Name string `pulumi:"name"`

	// NamespaceSelector decides whether to run the webhook on an object based on whether the namespace
	// for that object matches the selector. If the object itself is a namespace, the matching is
	// performed on object.metadata.labels. If the object is another cluster scoped resource, it never
	// skips the webhook.
	// 
	// For example, to run the webhook on any objects whose namespace is not associated with "runlevel"
	// of "0" or "1";  you will set the selector as follows: "namespaceSelector": {
	//   "matchExpressions": [
	//     {
	//       "key": "runlevel",
	//       "operator": "NotIn",
	//       "values": [
	//         "0",
	//         "1"
	//       ]
	//     }
	//   ]
	// }
	// 
	// If instead you want to only run the webhook on any objects whose namespace is associated with
	// the "environment" of "prod" or "staging"; you will set the selector as follows:
	// "namespaceSelector": {
	//   "matchExpressions": [
	//     {
	//       "key": "environment",
	//       "operator": "In",
	//       "values": [
	//         "prod",
	//         "staging"
	//       ]
	//     }
	//   ]
	// }
	// 
	// See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more examples
	// of label selectors.
	// 
	// Default to the empty LabelSelector, which matches everything.
	NamespaceSelector *metaV1.LabelSelector `pulumi:"namespaceSelector"`

	// ObjectSelector decides whether to run the webhook based on if the object has matching labels.
	// objectSelector is evaluated against both the oldObject and newObject that would be sent to the
	// webhook, and is considered to match if either object matches the selector. A null object
	// (oldObject in the case of create, or newObject in the case of delete) or an object that cannot
	// have labels (like a DeploymentRollback or a PodProxyOptions object) is not considered to match.
	// Use the object selector only if the webhook is opt-in, because end users may skip the admission
	// webhook by setting the labels. Default to the empty LabelSelector, which matches everything.
	ObjectSelector *metaV1.LabelSelector `pulumi:"objectSelector"`

	// reinvocationPolicy indicates whether this webhook should be called multiple times as part of a
	// single admission evaluation. Allowed values are "Never" and "IfNeeded".
	// 
	// Never: the webhook will not be called more than once in a single admission evaluation.
	// 
	// IfNeeded: the webhook will be called at least one additional time as part of the admission
	// evaluation if the object being admitted is modified by other admission plugins after the initial
	// webhook call. Webhooks that specify this option *must* be idempotent, able to process objects
	// they previously admitted. Note: * the number of additional invocations is not guaranteed to be
	// exactly one. * if additional invocations result in further modifications to the object, webhooks
	// are not guaranteed to be invoked again. * webhooks that use this option may be reordered to
	// minimize the number of additional invocations. * to validate an object after all mutations are
	// guaranteed complete, use a validating admission webhook instead.
	// 
	// Defaults to "Never".
	ReinvocationPolicy *string `pulumi:"reinvocationPolicy"`

	// Rules describes what operations on what resources/subresources the webhook cares about. The
	// webhook cares about an operation if it matches _any_ Rule. However, in order to prevent
	// ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks from putting the cluster in a state
	// which cannot be recovered from without completely disabling the plugin,
	// ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks are never called on admission requests
	// for ValidatingWebhookConfiguration and MutatingWebhookConfiguration objects.
	Rules []RuleWithOperations `pulumi:"rules"`

	// SideEffects states whether this webhookk has side effects. Acceptable values are: Unknown, None,
	// Some, NoneOnDryRun Webhooks with side effects MUST implement a reconciliation system, since a
	// request may be rejected by a future step in the admission change and the side effects therefore
	// need to be undone. Requests with the dryRun attribute will be auto-rejected if they match a
	// webhook with sideEffects == Unknown or Some. Defaults to Unknown.
	SideEffects *string `pulumi:"sideEffects"`

	// TimeoutSeconds specifies the timeout for this webhook. After the timeout passes, the webhook
	// call will be ignored or the API call will fail based on the failure policy. The timeout value
	// must be between 1 and 30 seconds. Default to 30 seconds.
	TimeoutSeconds *int `pulumi:"timeoutSeconds"`

}

var _MutatingWebhookType = reflect.TypeOf((*MutatingWebhook)(nil)).Elem()

// MutatingWebhookInput represents an input type that resolves to a MutatingWebhook.
type MutatingWebhookInput interface {
	ElementType() reflect.Type

	ToMutatingWebhookOutput() MutatingWebhookOutput
	ToMutatingWebhookOutputWithContext(ctx context.Context) MutatingWebhookOutput
}

// MutatingWebhookArgs is a MutatingWebhookInput whose fields are all Input types.
type MutatingWebhookArgs struct {
	// ClientConfig defines how to communicate with the hook. Required
	ClientConfig WebhookClientConfigInput `pulumi:"clientConfig"`

	// The name of the admission webhook. Name should be fully qualified, e.g.,
	// imagepolicy.kubernetes.io, where "imagepolicy" is the name of the webhook, and kubernetes.io is
	// the name of the organization. Required.
	Name pulumi.StringInput `pulumi:"name"`

	// AdmissionReviewVersions is an ordered list of preferred `AdmissionReview` versions the Webhook
	// expects. API server will try to use first version in the list which it supports. If none of the
	// versions specified in this list supported by API server, validation will fail for this object.
	// If a persisted webhook configuration specifies allowed versions and does not include any
	// versions known to the API Server, calls to the webhook will fail and be subject to the failure
	// policy. Default to `['v1beta1']`.
	AdmissionReviewVersions pulumi.StringArrayInput `pulumi:"admissionReviewVersions"`

	// FailurePolicy defines how unrecognized errors from the admission endpoint are handled - allowed
	// values are Ignore or Fail. Defaults to Ignore.
	FailurePolicy pulumi.StringInput `pulumi:"failurePolicy"`

	// matchPolicy defines how the "rules" list is used to match incoming requests. Allowed values are
	// "Exact" or "Equivalent".
	// 
	// - Exact: match a request only if it exactly matches a specified rule. For example, if
	// deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, but "rules" only
	// included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`, a request to
	// apps/v1beta1 or extensions/v1beta1 would not be sent to the webhook.
	// 
	// - Equivalent: match a request if modifies a resource listed in rules, even via another API group
	// or version. For example, if deployments can be modified via apps/v1, apps/v1beta1, and
	// extensions/v1beta1, and "rules" only included `apiGroups:["apps"], apiVersions:["v1"],
	// resources: ["deployments"]`, a request to apps/v1beta1 or extensions/v1beta1 would be converted
	// to apps/v1 and sent to the webhook.
	// 
	// Defaults to "Exact"
	MatchPolicy pulumi.StringInput `pulumi:"matchPolicy"`

	// NamespaceSelector decides whether to run the webhook on an object based on whether the namespace
	// for that object matches the selector. If the object itself is a namespace, the matching is
	// performed on object.metadata.labels. If the object is another cluster scoped resource, it never
	// skips the webhook.
	// 
	// For example, to run the webhook on any objects whose namespace is not associated with "runlevel"
	// of "0" or "1";  you will set the selector as follows: "namespaceSelector": {
	//   "matchExpressions": [
	//     {
	//       "key": "runlevel",
	//       "operator": "NotIn",
	//       "values": [
	//         "0",
	//         "1"
	//       ]
	//     }
	//   ]
	// }
	// 
	// If instead you want to only run the webhook on any objects whose namespace is associated with
	// the "environment" of "prod" or "staging"; you will set the selector as follows:
	// "namespaceSelector": {
	//   "matchExpressions": [
	//     {
	//       "key": "environment",
	//       "operator": "In",
	//       "values": [
	//         "prod",
	//         "staging"
	//       ]
	//     }
	//   ]
	// }
	// 
	// See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more examples
	// of label selectors.
	// 
	// Default to the empty LabelSelector, which matches everything.
	NamespaceSelector metaV1.LabelSelectorInput `pulumi:"namespaceSelector"`

	// ObjectSelector decides whether to run the webhook based on if the object has matching labels.
	// objectSelector is evaluated against both the oldObject and newObject that would be sent to the
	// webhook, and is considered to match if either object matches the selector. A null object
	// (oldObject in the case of create, or newObject in the case of delete) or an object that cannot
	// have labels (like a DeploymentRollback or a PodProxyOptions object) is not considered to match.
	// Use the object selector only if the webhook is opt-in, because end users may skip the admission
	// webhook by setting the labels. Default to the empty LabelSelector, which matches everything.
	ObjectSelector metaV1.LabelSelectorInput `pulumi:"objectSelector"`

	// reinvocationPolicy indicates whether this webhook should be called multiple times as part of a
	// single admission evaluation. Allowed values are "Never" and "IfNeeded".
	// 
	// Never: the webhook will not be called more than once in a single admission evaluation.
	// 
	// IfNeeded: the webhook will be called at least one additional time as part of the admission
	// evaluation if the object being admitted is modified by other admission plugins after the initial
	// webhook call. Webhooks that specify this option *must* be idempotent, able to process objects
	// they previously admitted. Note: * the number of additional invocations is not guaranteed to be
	// exactly one. * if additional invocations result in further modifications to the object, webhooks
	// are not guaranteed to be invoked again. * webhooks that use this option may be reordered to
	// minimize the number of additional invocations. * to validate an object after all mutations are
	// guaranteed complete, use a validating admission webhook instead.
	// 
	// Defaults to "Never".
	ReinvocationPolicy pulumi.StringInput `pulumi:"reinvocationPolicy"`

	// Rules describes what operations on what resources/subresources the webhook cares about. The
	// webhook cares about an operation if it matches _any_ Rule. However, in order to prevent
	// ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks from putting the cluster in a state
	// which cannot be recovered from without completely disabling the plugin,
	// ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks are never called on admission requests
	// for ValidatingWebhookConfiguration and MutatingWebhookConfiguration objects.
	Rules RuleWithOperationsArrayInput `pulumi:"rules"`

	// SideEffects states whether this webhookk has side effects. Acceptable values are: Unknown, None,
	// Some, NoneOnDryRun Webhooks with side effects MUST implement a reconciliation system, since a
	// request may be rejected by a future step in the admission change and the side effects therefore
	// need to be undone. Requests with the dryRun attribute will be auto-rejected if they match a
	// webhook with sideEffects == Unknown or Some. Defaults to Unknown.
	SideEffects pulumi.StringInput `pulumi:"sideEffects"`

	// TimeoutSeconds specifies the timeout for this webhook. After the timeout passes, the webhook
	// call will be ignored or the API call will fail based on the failure policy. The timeout value
	// must be between 1 and 30 seconds. Default to 30 seconds.
	TimeoutSeconds pulumi.IntInput `pulumi:"timeoutSeconds"`

}

func (a MutatingWebhookArgs) ElementType() reflect.Type {
	return _MutatingWebhookType
}

func (a MutatingWebhookArgs) ToMutatingWebhookOutput() MutatingWebhookOutput {
	return pulumi.ToOutput(a).(MutatingWebhookOutput)
}

func (a MutatingWebhookArgs) ToMutatingWebhookOutputWithContext(ctx context.Context) MutatingWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, a).(MutatingWebhookOutput)
}

// MutatingWebhookOutput is an output type that resolves to a Input.
type MutatingWebhookOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(MutatingWebhookOutput{}) }

func (MutatingWebhookOutput) ElementType() reflect.Type {
	return _MutatingWebhookType
}

func (o MutatingWebhookOutput) AdmissionReviewVersions() pulumi.StringArrayOutput {
	return o.Apply(func(v MutatingWebhook) []string {
		return v.AdmissionReviewVersions
	}).(pulumi.StringArrayOutput)
}

func (o MutatingWebhookOutput) ClientConfig() WebhookClientConfigOutput {
	return o.Apply(func(v MutatingWebhook) WebhookClientConfig {
		return v.ClientConfig
	}).(WebhookClientConfigOutput)
}

func (o MutatingWebhookOutput) FailurePolicy() pulumi.StringOutput {
	return o.Apply(func(v MutatingWebhook) *string {
		return v.FailurePolicy
	}).(pulumi.StringOutput)
}

func (o MutatingWebhookOutput) MatchPolicy() pulumi.StringOutput {
	return o.Apply(func(v MutatingWebhook) *string {
		return v.MatchPolicy
	}).(pulumi.StringOutput)
}

func (o MutatingWebhookOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v MutatingWebhook) string {
		return v.Name
	}).(pulumi.StringOutput)
}

func (o MutatingWebhookOutput) NamespaceSelector() metaV1.LabelSelectorOutput {
	return o.Apply(func(v MutatingWebhook) *metaV1.LabelSelector {
		return v.NamespaceSelector
	}).(metaV1.LabelSelectorOutput)
}

func (o MutatingWebhookOutput) ObjectSelector() metaV1.LabelSelectorOutput {
	return o.Apply(func(v MutatingWebhook) *metaV1.LabelSelector {
		return v.ObjectSelector
	}).(metaV1.LabelSelectorOutput)
}

func (o MutatingWebhookOutput) ReinvocationPolicy() pulumi.StringOutput {
	return o.Apply(func(v MutatingWebhook) *string {
		return v.ReinvocationPolicy
	}).(pulumi.StringOutput)
}

func (o MutatingWebhookOutput) Rules() RuleWithOperationsArrayOutput {
	return o.Apply(func(v MutatingWebhook) []RuleWithOperations {
		return v.Rules
	}).(RuleWithOperationsArrayOutput)
}

func (o MutatingWebhookOutput) SideEffects() pulumi.StringOutput {
	return o.Apply(func(v MutatingWebhook) *string {
		return v.SideEffects
	}).(pulumi.StringOutput)
}

func (o MutatingWebhookOutput) TimeoutSeconds() pulumi.IntOutput {
	return o.Apply(func(v MutatingWebhook) *int {
		return v.TimeoutSeconds
	}).(pulumi.IntOutput)
}

var _MutatingWebhookArrayType = reflect.TypeOf((*[]MutatingWebhook)(nil)).Elem()

type MutatingWebhookArrayInput interface {
	ElementType() reflect.Type

	ToMutatingWebhookArrayOutput() MutatingWebhookArrayOutput
	ToMutatingWebhookArrayOutputWithContext(ctx context.Context) MutatingWebhookArrayOutput
}

type MutatingWebhookArray []MutatingWebhookInput

func (a MutatingWebhookArray) ElementType() reflect.Type {
	return _MutatingWebhookArrayType
}

func (a MutatingWebhookArray) ToMutatingWebhookArrayOutput() MutatingWebhookArrayOutput {
	return pulumi.ToOutput(a).(MutatingWebhookArrayOutput)
}

func (a MutatingWebhookArray) ToMutatingWebhookArrayOutputWithContext(ctx context.Context) MutatingWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(MutatingWebhookArrayOutput)
}

type MutatingWebhookArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(MutatingWebhookArrayOutput{}) }

func (MutatingWebhookArrayOutput) ElementType() reflect.Type {
	return _MutatingWebhookArrayType
}

func (o MutatingWebhookArrayOutput) Index(i pulumi.IntInput) MutatingWebhookOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) MutatingWebhook {
		return vs[0].([]MutatingWebhook)[vs[1].(int)]
	}).(MutatingWebhookOutput)
}
