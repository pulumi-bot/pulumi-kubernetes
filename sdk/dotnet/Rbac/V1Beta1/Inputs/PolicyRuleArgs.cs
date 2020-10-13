// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Rbac.V1Beta1
{

    /// <summary>
    /// PolicyRule holds information that describes a policy rule, but does not contain information about who the rule applies to or which namespace the rule applies to.
    /// </summary>
    public class PolicyRuleArgs : Pulumi.ResourceArgs
    {
        [Input("apiGroups")]
        private InputList<string>? _apiGroups;

        /// <summary>
        /// APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.
        /// </summary>
        public InputList<string> ApiGroups
        {
            get => _apiGroups ?? (_apiGroups = new InputList<string>());
            set => _apiGroups = value;
        }

        [Input("nonResourceURLs")]
        private InputList<string>? _nonResourceURLs;

        /// <summary>
        /// NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"),  but not both.
        /// </summary>
        public InputList<string> NonResourceURLs
        {
            get => _nonResourceURLs ?? (_nonResourceURLs = new InputList<string>());
            set => _nonResourceURLs = value;
        }

        [Input("resourceNames")]
        private InputList<string>? _resourceNames;

        /// <summary>
        /// ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.
        /// </summary>
        public InputList<string> ResourceNames
        {
            get => _resourceNames ?? (_resourceNames = new InputList<string>());
            set => _resourceNames = value;
        }

        [Input("resources")]
        private InputList<string>? _resources;

        /// <summary>
        /// Resources is a list of resources this rule applies to.  '*' represents all resources in the specified apiGroups. '*/foo' represents the subresource 'foo' for all resources in the specified apiGroups.
        /// </summary>
        public InputList<string> Resources
        {
            get => _resources ?? (_resources = new InputList<string>());
            set => _resources = value;
        }

        [Input("verbs", required: true)]
        private InputList<string>? _verbs;

        /// <summary>
        /// Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule.  VerbAll represents all kinds.
        /// </summary>
        public InputList<string> Verbs
        {
            get => _verbs ?? (_verbs = new InputList<string>());
            set => _verbs = value;
        }

        public PolicyRuleArgs()
        {
        }
    }
}
