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
    /// PortworxVolumeSource represents a Portworx volume resource.
    /// </summary>
    public class PortworxVolumeSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs". Implicitly inferred to be "ext4" if unspecified.
        /// </summary>
        [Input("fsType")]
        public Input<string>? FsType { get; set; }

        /// <summary>
        /// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// VolumeID uniquely identifies a Portworx volume
        /// </summary>
        [Input("volumeID", required: true)]
        public Input<string> VolumeID { get; set; } = null!;

        public PortworxVolumeSourceArgs()
        {
        }
    }
}
