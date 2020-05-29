// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Core.V1
{

    /// <summary>
    /// Represents a Photon Controller persistent disk resource.
    /// </summary>
    public class PhotonPersistentDiskVolumeSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
        /// </summary>
        [Input("fsType")]
        public Input<string>? FsType { get; set; }

        /// <summary>
        /// ID that identifies Photon Controller persistent disk
        /// </summary>
        [Input("pdID", required: true)]
        public Input<string> PdID { get; set; } = null!;

        public PhotonPersistentDiskVolumeSourceArgs()
        {
        }
    }
}
