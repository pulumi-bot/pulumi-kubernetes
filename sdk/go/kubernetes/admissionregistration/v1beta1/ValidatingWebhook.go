// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// ValidatingWebhook describes an admission webhook and the resources and operations it applies to.
type ValidatingWebhook struct {
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
	// See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels for more examples
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

var _ValidatingWebhookType = reflect.TypeOf((*ValidatingWebhook)(nil)).Elem()

// ValidatingWebhookInput represents an input type that resolves to a ValidatingWebhook.
type ValidatingWebhookInput interface {
	ElementType() reflect.Type

	ToValidatingWebhookOutput() ValidatingWebhookOutput
	ToValidatingWebhookOutputWithContext(ctx context.Context) ValidatingWebhookOutput
}

// ValidatingWebhookArgs is a ValidatingWebhookInput whose fields are all Input types.
type ValidatingWebhookArgs struct {
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
	// See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels for more examples
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

func (a ValidatingWebhookArgs) ElementType() reflect.Type {
	return _ValidatingWebhookType
}

func (a ValidatingWebhookArgs) ToValidatingWebhookOutput() ValidatingWebhookOutput {
	return pulumi.ToOutput(a).(ValidatingWebhookOutput)
}

func (a ValidatingWebhookArgs) ToValidatingWebhookOutputWithContext(ctx context.Context) ValidatingWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ValidatingWebhookOutput)
}

// ValidatingWebhookOutput is an output type that resolves to a Input.
type ValidatingWebhookOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ValidatingWebhookOutput{}) }

func (ValidatingWebhookOutput) ElementType() reflect.Type {
	return _ValidatingWebhookType
}

func (o ValidatingWebhookOutput) AdmissionReviewVersions() pulumi.StringArrayOutput {
	return o.Apply(func(v ValidatingWebhook) []string {
		return v.AdmissionReviewVersions
	}).(pulumi.StringArrayOutput)
}

func (o ValidatingWebhookOutput) ClientConfig() WebhookClientConfigOutput {
	return o.Apply(func(v ValidatingWebhook) WebhookClientConfig {
		return v.ClientConfig
	}).(WebhookClientConfigOutput)
}

func (o ValidatingWebhookOutput) FailurePolicy() pulumi.StringOutput {
	return o.Apply(func(v ValidatingWebhook) *string {
		return v.FailurePolicy
	}).(pulumi.StringOutput)
}

func (o ValidatingWebhookOutput) MatchPolicy() pulumi.StringOutput {
	return o.Apply(func(v ValidatingWebhook) *string {
		return v.MatchPolicy
	}).(pulumi.StringOutput)
}

func (o ValidatingWebhookOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v ValidatingWebhook) string {
		return v.Name
	}).(pulumi.StringOutput)
}

func (o ValidatingWebhookOutput) NamespaceSelector() metaV1.LabelSelectorOutput {
	return o.Apply(func(v ValidatingWebhook) *metaV1.LabelSelector {
		return v.NamespaceSelector
	}).(metaV1.LabelSelectorOutput)
}

func (o ValidatingWebhookOutput) ObjectSelector() metaV1.LabelSelectorOutput {
	return o.Apply(func(v ValidatingWebhook) *metaV1.LabelSelector {
		return v.ObjectSelector
	}).(metaV1.LabelSelectorOutput)
}

func (o ValidatingWebhookOutput) Rules() RuleWithOperationsArrayOutput {
	return o.Apply(func(v ValidatingWebhook) []RuleWithOperations {
		return v.Rules
	}).(RuleWithOperationsArrayOutput)
}

func (o ValidatingWebhookOutput) SideEffects() pulumi.StringOutput {
	return o.Apply(func(v ValidatingWebhook) *string {
		return v.SideEffects
	}).(pulumi.StringOutput)
}

func (o ValidatingWebhookOutput) TimeoutSeconds() pulumi.IntOutput {
	return o.Apply(func(v ValidatingWebhook) *int {
		return v.TimeoutSeconds
	}).(pulumi.IntOutput)
}

var _ValidatingWebhookArrayType = reflect.TypeOf((*[]ValidatingWebhook)(nil)).Elem()

type ValidatingWebhookArrayInput interface {
	ElementType() reflect.Type

	ToValidatingWebhookArrayOutput() ValidatingWebhookArrayOutput
	ToValidatingWebhookArrayOutputWithContext(ctx context.Context) ValidatingWebhookArrayOutput
}

type ValidatingWebhookArray []ValidatingWebhookInput

func (a ValidatingWebhookArray) ElementType() reflect.Type {
	return _ValidatingWebhookArrayType
}

func (a ValidatingWebhookArray) ToValidatingWebhookArrayOutput() ValidatingWebhookArrayOutput {
	return pulumi.ToOutput(a).(ValidatingWebhookArrayOutput)
}

func (a ValidatingWebhookArray) ToValidatingWebhookArrayOutputWithContext(ctx context.Context) ValidatingWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ValidatingWebhookArrayOutput)
}

type ValidatingWebhookArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ValidatingWebhookArrayOutput{}) }

func (ValidatingWebhookArrayOutput) ElementType() reflect.Type {
	return _ValidatingWebhookArrayType
}

func (o ValidatingWebhookArrayOutput) Index(i pulumi.IntInput) ValidatingWebhookOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) ValidatingWebhook {
		return vs[0].([]ValidatingWebhook)[vs[1].(int)]
	}).(ValidatingWebhookOutput)
}
