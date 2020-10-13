// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1
{

    [OutputType]
    public sealed class RuntimeClassStrategyOptions
    {
        /// <summary>
        /// allowedRuntimeClassNames is a whitelist of RuntimeClass names that may be specified on a pod. A value of "*" means that any RuntimeClass name is allowed, and must be the only item in the list. An empty list requires the RuntimeClassName field to be unset.
        /// </summary>
        public readonly ImmutableArray<string> AllowedRuntimeClassNames;
        /// <summary>
        /// defaultRuntimeClassName is the default RuntimeClassName to set on the pod. The default MUST be allowed by the allowedRuntimeClassNames list. A value of nil does not mutate the Pod.
        /// </summary>
        public readonly string DefaultRuntimeClassName;

        [OutputConstructor]
        private RuntimeClassStrategyOptions(
            ImmutableArray<string> allowedRuntimeClassNames,

            string defaultRuntimeClassName)
        {
            AllowedRuntimeClassNames = allowedRuntimeClassNames;
            DefaultRuntimeClassName = defaultRuntimeClassName;
        }
    }
}
