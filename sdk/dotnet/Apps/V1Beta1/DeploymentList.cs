// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Apps.V1Beta1
{
    /// <summary>
    /// DeploymentList is a list of Deployments.
    /// </summary>
    [KubernetesResourceType("kubernetes:apps/v1beta1:DeploymentList")]
    public partial class DeploymentList : KubernetesResource
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// Items is the list of Deployments.
        /// </summary>
        [Output("items")]
        public Output<ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta1.Deployment>> Items { get; private set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Standard list metadata.
        /// </summary>
        [Output("metadata")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Meta.V1.ListMeta> Metadata { get; private set; } = null!;


        /// <summary>
        /// Create a DeploymentList resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeploymentList(string name, Pulumi.Kubernetes.Types.Inputs.Apps.V1Beta1.DeploymentListArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:apps/v1beta1:DeploymentList", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }
        internal DeploymentList(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:apps/v1beta1:DeploymentList", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private DeploymentList(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes:apps/v1beta1:DeploymentList", name, null, MakeResourceOptions(options, id))
        {
        }

        private static Pulumi.Kubernetes.Types.Inputs.Apps.V1Beta1.DeploymentListArgs? MakeArgs(Pulumi.Kubernetes.Types.Inputs.Apps.V1Beta1.DeploymentListArgs? args)
        {
            args ??= new Pulumi.Kubernetes.Types.Inputs.Apps.V1Beta1.DeploymentListArgs();
            args.ApiVersion = "apps/v1beta1";
            args.Kind = "DeploymentList";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DeploymentList resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeploymentList Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DeploymentList(name, id, options);
        }
    }
}
namespace Pulumi.Kubernetes.Types.Inputs.Apps.V1Beta1
{

    public class DeploymentListArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        [Input("items", required: true)]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Apps.V1Beta1.DeploymentArgs>? _items;

        /// <summary>
        /// Items is the list of Deployments.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Apps.V1Beta1.DeploymentArgs> Items
        {
            get => _items ?? (_items = new InputList<Pulumi.Kubernetes.Types.Inputs.Apps.V1Beta1.DeploymentArgs>());
            set => _items = value;
        }

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Standard list metadata.
        /// </summary>
        [Input("metadata")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ListMetaArgs>? Metadata { get; set; }

        public DeploymentListArgs()
        {
        }
    }
}
