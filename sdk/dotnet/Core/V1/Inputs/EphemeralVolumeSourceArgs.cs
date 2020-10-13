// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Core.V1
{

    /// <summary>
    /// Represents an ephemeral volume that is handled by a normal storage driver.
    /// </summary>
    public class EphemeralVolumeSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies a read-only configuration for the volume. Defaults to false (read/write).
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be `&lt;pod name&gt;-&lt;volume name&gt;` where `&lt;volume name&gt;` is the name from the `PodSpec.Volumes` array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).
        /// 
        /// An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.
        /// 
        /// This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.
        /// 
        /// Required, must not be nil.
        /// </summary>
        [Input("volumeClaimTemplate")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.PersistentVolumeClaimTemplateArgs>? VolumeClaimTemplate { get; set; }

        public EphemeralVolumeSourceArgs()
        {
        }
    }
}
