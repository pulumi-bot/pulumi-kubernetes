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
    public sealed class LimitRangeSpec
    {
        /// <summary>
        /// Limits is the list of LimitRangeItem objects that are enforced.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.LimitRangeItem> Limits;

        [OutputConstructor]
        private LimitRangeSpec(ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.LimitRangeItem> limits)
        {
            Limits = limits;
        }
    }
}
