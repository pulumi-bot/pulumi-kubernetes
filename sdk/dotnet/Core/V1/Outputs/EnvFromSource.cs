// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Core.V1
{

    [OutputType]
    public sealed class EnvFromSource
    {
        /// <summary>
        /// The ConfigMap to select from
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.ConfigMapEnvSource ConfigMapRef;
        /// <summary>
        /// An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
        /// </summary>
        public readonly string Prefix;
        /// <summary>
        /// The Secret to select from
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.SecretEnvSource SecretRef;

        [OutputConstructor]
        private EnvFromSource(
            Pulumi.Kubernetes.Types.Outputs.Core.V1.ConfigMapEnvSource configMapRef,

            string prefix,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.SecretEnvSource secretRef)
        {
            ConfigMapRef = configMapRef;
            Prefix = prefix;
            SecretRef = secretRef;
        }
    }
}
