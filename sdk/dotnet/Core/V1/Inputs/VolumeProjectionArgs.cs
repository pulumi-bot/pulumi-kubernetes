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
    /// Projection that may be projected along with other supported volume types
    /// </summary>
    public class VolumeProjectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// information about the configMap data to project
        /// </summary>
        [Input("configMap")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.ConfigMapProjectionArgs>? ConfigMap { get; set; }

        /// <summary>
        /// information about the downwardAPI data to project
        /// </summary>
        [Input("downwardAPI")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.DownwardAPIProjectionArgs>? DownwardAPI { get; set; }

        /// <summary>
        /// information about the secret data to project
        /// </summary>
        [Input("secret")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.SecretProjectionArgs>? Secret { get; set; }

        /// <summary>
        /// information about the serviceAccountToken data to project
        /// </summary>
        [Input("serviceAccountToken")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.ServiceAccountTokenProjectionArgs>? ServiceAccountToken { get; set; }

        public VolumeProjectionArgs()
        {
        }
    }
}
