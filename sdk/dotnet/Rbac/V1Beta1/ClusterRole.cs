// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Rbac.V1Beta1
{
    /// <summary>
    /// ClusterRole is a cluster level, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding or ClusterRoleBinding. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 ClusterRole, and will no longer be served in v1.22.
    /// </summary>
    public partial class ClusterRole : KubernetesResource
    {
        /// <summary>
        /// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
        /// </summary>
        [Output("aggregationRule")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Rbac.V1Beta1.AggregationRule> AggregationRule { get; private set; } = null!;

        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Standard object's metadata.
        /// </summary>
        [Output("metadata")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta> Metadata { get; private set; } = null!;

        /// <summary>
        /// Rules holds all the PolicyRules for this ClusterRole
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Rbac.V1Beta1.PolicyRule>> Rules { get; private set; } = null!;


        /// <summary>
        /// Create a ClusterRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterRole(string name, Pulumi.Kubernetes.Types.Inputs.Rbac.V1Beta1.ClusterRoleArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRole", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }
        internal ClusterRole(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRole", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private ClusterRole(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRole", name, null, MakeResourceOptions(options, id))
        {
        }

        private static Pulumi.Kubernetes.Types.Inputs.Rbac.V1Beta1.ClusterRoleArgs? MakeArgs(Pulumi.Kubernetes.Types.Inputs.Rbac.V1Beta1.ClusterRoleArgs? args)
        {
            args ??= new Pulumi.Kubernetes.Types.Inputs.Rbac.V1Beta1.ClusterRoleArgs();
            args.ApiVersion = "rbac.authorization.k8s.io/v1beta1";
            args.Kind = "ClusterRole";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "kubernetes:rbac.authorization.k8s.io/v1:ClusterRole"},
                    new Pulumi.Alias { Type = "kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRole"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ClusterRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterRole Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ClusterRole(name, id, options);
        }
    }
}
namespace Pulumi.Kubernetes.Types.Inputs.Rbac.V1Beta1
{

    public class ClusterRoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
        /// </summary>
        [Input("aggregationRule")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Rbac.V1Beta1.AggregationRuleArgs>? AggregationRule { get; set; }

        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Standard object's metadata.
        /// </summary>
        [Input("metadata")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs>? Metadata { get; set; }

        [Input("rules")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Rbac.V1Beta1.PolicyRuleArgs>? _rules;

        /// <summary>
        /// Rules holds all the PolicyRules for this ClusterRole
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Rbac.V1Beta1.PolicyRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Pulumi.Kubernetes.Types.Inputs.Rbac.V1Beta1.PolicyRuleArgs>());
            set => _rules = value;
        }

        public ClusterRoleArgs()
        {
        }
    }
}
