// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Storage.V1
{
    /// <summary>
    /// VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node.
    /// 
    /// VolumeAttachment objects are non-namespaced.
    /// </summary>
    [KubernetesResourceType("kubernetes:storage.k8s.io/v1:VolumeAttachment")]
    public partial class VolumeAttachment : KubernetesResource
    {
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
        /// Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Output("metadata")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta> Metadata { get; private set; } = null!;

        /// <summary>
        /// Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.
        /// </summary>
        [Output("spec")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Storage.V1.VolumeAttachmentSpec> Spec { get; private set; } = null!;

        /// <summary>
        /// Status of the VolumeAttachment request. Populated by the entity completing the attach or detach operation, i.e. the external-attacher.
        /// </summary>
        [Output("status")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Storage.V1.VolumeAttachmentStatus> Status { get; private set; } = null!;


        /// <summary>
        /// Create a VolumeAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VolumeAttachment(string name, Pulumi.Kubernetes.Types.Inputs.Storage.V1.VolumeAttachmentArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:storage.k8s.io/v1:VolumeAttachment", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }
        internal VolumeAttachment(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:storage.k8s.io/v1:VolumeAttachment", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private VolumeAttachment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes:storage.k8s.io/v1:VolumeAttachment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static Pulumi.Kubernetes.Types.Inputs.Storage.V1.VolumeAttachmentArgs? MakeArgs(Pulumi.Kubernetes.Types.Inputs.Storage.V1.VolumeAttachmentArgs? args)
        {
            args ??= new Pulumi.Kubernetes.Types.Inputs.Storage.V1.VolumeAttachmentArgs();
            args.ApiVersion = "storage.k8s.io/v1";
            args.Kind = "VolumeAttachment";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "kubernetes:storage.k8s.io/v1alpha1:VolumeAttachment"},
                    new Pulumi.Alias { Type = "kubernetes:storage.k8s.io/v1beta1:VolumeAttachment"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VolumeAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VolumeAttachment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VolumeAttachment(name, id, options);
        }
    }
}
namespace Pulumi.Kubernetes.Types.Inputs.Storage.V1
{

    public class VolumeAttachmentArgs : Pulumi.ResourceArgs
    {
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
        /// Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Input("metadata")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs>? Metadata { get; set; }

        /// <summary>
        /// Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.
        /// </summary>
        [Input("spec", required: true)]
        public Input<Pulumi.Kubernetes.Types.Inputs.Storage.V1.VolumeAttachmentSpecArgs> Spec { get; set; } = null!;

        public VolumeAttachmentArgs()
        {
        }
    }
}
