// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Node.V1Beta1
{
    /// <summary>
    /// RuntimeClass defines a class of container runtime supported in the cluster. The RuntimeClass is used to determine which container runtime is used to run all containers in a pod. RuntimeClasses are (currently) manually defined by a user or cluster provisioner, and referenced in the PodSpec. The Kubelet is responsible for resolving the RuntimeClassName reference before running the pod.  For more details, see https://git.k8s.io/enhancements/keps/sig-node/runtime-class.md
    /// </summary>
    [KubernetesResourceType("kubernetes:node.k8s.io/v1beta1:RuntimeClass")]
    public partial class RuntimeClass : KubernetesResource
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// Handler specifies the underlying runtime and configuration that the CRI implementation will use to handle pods of this class. The possible values are specific to the node &amp; CRI configuration.  It is assumed that all handlers are available on every node, and handlers of the same name are equivalent on every node. For example, a handler called "runc" might specify that the runc OCI runtime (using native Linux containers) will be used to run the containers in a pod. The Handler must be lowercase, conform to the DNS Label (RFC 1123) requirements, and is immutable.
        /// </summary>
        [Output("handler")]
        public Output<string> Handler { get; private set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Output("metadata")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta> Metadata { get; private set; } = null!;

        /// <summary>
        /// Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. For more details, see https://git.k8s.io/enhancements/keps/sig-node/20190226-pod-overhead.md This field is alpha-level as of Kubernetes v1.15, and is only honored by servers that enable the PodOverhead feature.
        /// </summary>
        [Output("overhead")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Node.V1Beta1.Overhead> Overhead { get; private set; } = null!;

        /// <summary>
        /// Scheduling holds the scheduling constraints to ensure that pods running with this RuntimeClass are scheduled to nodes that support it. If scheduling is nil, this RuntimeClass is assumed to be supported by all nodes.
        /// </summary>
        [Output("scheduling")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Node.V1Beta1.Scheduling> Scheduling { get; private set; } = null!;


        /// <summary>
        /// Create a RuntimeClass resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RuntimeClass(string name, Pulumi.Kubernetes.Types.Inputs.Node.V1Beta1.RuntimeClassArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:node.k8s.io/v1beta1:RuntimeClass", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }
        internal RuntimeClass(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:node.k8s.io/v1beta1:RuntimeClass", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private RuntimeClass(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes:node.k8s.io/v1beta1:RuntimeClass", name, null, MakeResourceOptions(options, id))
        {
        }

        private static Pulumi.Kubernetes.Types.Inputs.Node.V1Beta1.RuntimeClassArgs? MakeArgs(Pulumi.Kubernetes.Types.Inputs.Node.V1Beta1.RuntimeClassArgs? args)
        {
            args ??= new Pulumi.Kubernetes.Types.Inputs.Node.V1Beta1.RuntimeClassArgs();
            args.ApiVersion = "node.k8s.io/v1beta1";
            args.Kind = "RuntimeClass";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "kubernetes:node.k8s.io/v1:RuntimeClass"},
                    new Pulumi.Alias { Type = "kubernetes:node.k8s.io/v1alpha1:RuntimeClass"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RuntimeClass resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RuntimeClass Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RuntimeClass(name, id, options);
        }
    }
}
namespace Pulumi.Kubernetes.Types.Inputs.Node.V1Beta1
{

    public class RuntimeClassArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        /// <summary>
        /// Handler specifies the underlying runtime and configuration that the CRI implementation will use to handle pods of this class. The possible values are specific to the node &amp; CRI configuration.  It is assumed that all handlers are available on every node, and handlers of the same name are equivalent on every node. For example, a handler called "runc" might specify that the runc OCI runtime (using native Linux containers) will be used to run the containers in a pod. The Handler must be lowercase, conform to the DNS Label (RFC 1123) requirements, and is immutable.
        /// </summary>
        [Input("handler", required: true)]
        public Input<string> Handler { get; set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Input("metadata")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs>? Metadata { get; set; }

        /// <summary>
        /// Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. For more details, see https://git.k8s.io/enhancements/keps/sig-node/20190226-pod-overhead.md This field is alpha-level as of Kubernetes v1.15, and is only honored by servers that enable the PodOverhead feature.
        /// </summary>
        [Input("overhead")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Node.V1Beta1.OverheadArgs>? Overhead { get; set; }

        /// <summary>
        /// Scheduling holds the scheduling constraints to ensure that pods running with this RuntimeClass are scheduled to nodes that support it. If scheduling is nil, this RuntimeClass is assumed to be supported by all nodes.
        /// </summary>
        [Input("scheduling")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Node.V1Beta1.SchedulingArgs>? Scheduling { get; set; }

        public RuntimeClassArgs()
        {
        }
    }
}
