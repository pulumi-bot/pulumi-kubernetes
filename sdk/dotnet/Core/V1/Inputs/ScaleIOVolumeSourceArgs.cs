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
    /// ScaleIOVolumeSource represents a persistent ScaleIO volume
    /// </summary>
    public class ScaleIOVolumeSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Default is "xfs".
        /// </summary>
        [Input("fsType")]
        public Input<string>? FsType { get; set; }

        /// <summary>
        /// The host address of the ScaleIO API Gateway.
        /// </summary>
        [Input("gateway", required: true)]
        public Input<string> Gateway { get; set; } = null!;

        /// <summary>
        /// The name of the ScaleIO Protection Domain for the configured storage.
        /// </summary>
        [Input("protectionDomain")]
        public Input<string>? ProtectionDomain { get; set; }

        /// <summary>
        /// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.
        /// </summary>
        [Input("secretRef", required: true)]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.LocalObjectReferenceArgs> SecretRef { get; set; } = null!;

        /// <summary>
        /// Flag to enable/disable SSL communication with Gateway, default false
        /// </summary>
        [Input("sslEnabled")]
        public Input<bool>? SslEnabled { get; set; }

        /// <summary>
        /// Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.
        /// </summary>
        [Input("storageMode")]
        public Input<string>? StorageMode { get; set; }

        /// <summary>
        /// The ScaleIO Storage Pool associated with the protection domain.
        /// </summary>
        [Input("storagePool")]
        public Input<string>? StoragePool { get; set; }

        /// <summary>
        /// The name of the storage system as configured in ScaleIO.
        /// </summary>
        [Input("system", required: true)]
        public Input<string> System { get; set; } = null!;

        /// <summary>
        /// The name of a volume already created in the ScaleIO system that is associated with this volume source.
        /// </summary>
        [Input("volumeName")]
        public Input<string>? VolumeName { get; set; }

        public ScaleIOVolumeSourceArgs()
        {
        }
    }
}
