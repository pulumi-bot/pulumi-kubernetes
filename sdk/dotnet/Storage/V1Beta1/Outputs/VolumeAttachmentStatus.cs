// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Storage.V1Beta1
{

    [OutputType]
    public sealed class VolumeAttachmentStatus
    {
        /// <summary>
        /// The last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Storage.V1Beta1.VolumeError AttachError;
        /// <summary>
        /// Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        /// </summary>
        public readonly bool Attached;
        /// <summary>
        /// Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        /// </summary>
        public readonly ImmutableDictionary<string, string> AttachmentMetadata;
        /// <summary>
        /// The last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Storage.V1Beta1.VolumeError DetachError;

        [OutputConstructor]
        private VolumeAttachmentStatus(
            Pulumi.Kubernetes.Types.Outputs.Storage.V1Beta1.VolumeError attachError,

            bool attached,

            ImmutableDictionary<string, string> attachmentMetadata,

            Pulumi.Kubernetes.Types.Outputs.Storage.V1Beta1.VolumeError detachError)
        {
            AttachError = attachError;
            Attached = attached;
            AttachmentMetadata = attachmentMetadata;
            DetachError = detachError;
        }
    }
}
