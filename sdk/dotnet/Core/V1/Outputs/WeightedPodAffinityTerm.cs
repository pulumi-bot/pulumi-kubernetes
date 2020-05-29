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
    public sealed class WeightedPodAffinityTerm
    {
        /// <summary>
        /// Required. A pod affinity term, associated with the corresponding weight.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.PodAffinityTerm PodAffinityTerm;
        /// <summary>
        /// weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
        /// </summary>
        public readonly int Weight;

        [OutputConstructor]
        private WeightedPodAffinityTerm(
            Pulumi.Kubernetes.Types.Outputs.Core.V1.PodAffinityTerm podAffinityTerm,

            int weight)
        {
            PodAffinityTerm = podAffinityTerm;
            Weight = weight;
        }
    }
}
