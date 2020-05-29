// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Core.V1
{

    [OutputType]
    public sealed class AzureFilePersistentVolumeSource
    {
        /// <summary>
        /// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
        /// </summary>
        public readonly bool ReadOnly;
        /// <summary>
        /// the name of secret that contains Azure Storage Account Name and Key
        /// </summary>
        public readonly string SecretName;
        /// <summary>
        /// the namespace of the secret that contains Azure Storage Account Name and Key default is the same as the Pod
        /// </summary>
        public readonly string SecretNamespace;
        /// <summary>
        /// Share Name
        /// </summary>
        public readonly string ShareName;

        [OutputConstructor]
        private AzureFilePersistentVolumeSource(
            bool readOnly,

            string secretName,

            string secretNamespace,

            string shareName)
        {
            ReadOnly = readOnly;
            SecretName = secretName;
            SecretNamespace = secretNamespace;
            ShareName = shareName;
        }
    }
}
