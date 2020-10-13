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
    /// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
    /// </summary>
    public class AzureFilePersistentVolumeSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// the name of secret that contains Azure Storage Account Name and Key
        /// </summary>
        [Input("secretName", required: true)]
        public Input<string> SecretName { get; set; } = null!;

        /// <summary>
        /// the namespace of the secret that contains Azure Storage Account Name and Key default is the same as the Pod
        /// </summary>
        [Input("secretNamespace")]
        public Input<string>? SecretNamespace { get; set; }

        /// <summary>
        /// Share Name
        /// </summary>
        [Input("shareName", required: true)]
        public Input<string> ShareName { get; set; } = null!;

        public AzureFilePersistentVolumeSourceArgs()
        {
        }
    }
}
