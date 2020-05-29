// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Storage.V1
{

    [OutputType]
    public sealed class VolumeAttachmentSource
    {
        /// <summary>
        /// inlineVolumeSpec contains all the information necessary to attach a persistent volume defined by a pod's inline VolumeSource. This field is populated only for the CSIMigration feature. It contains translated fields from a pod's inline VolumeSource to a PersistentVolumeSpec. This field is alpha-level and is only honored by servers that enabled the CSIMigration feature.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.PersistentVolumeSpec InlineVolumeSpec;
        /// <summary>
        /// Name of the persistent volume to attach.
        /// </summary>
        public readonly string PersistentVolumeName;

        [OutputConstructor]
        private VolumeAttachmentSource(
            Pulumi.Kubernetes.Types.Outputs.Core.V1.PersistentVolumeSpec inlineVolumeSpec,

            string persistentVolumeName)
        {
            InlineVolumeSpec = inlineVolumeSpec;
            PersistentVolumeName = persistentVolumeName;
        }
    }
}
