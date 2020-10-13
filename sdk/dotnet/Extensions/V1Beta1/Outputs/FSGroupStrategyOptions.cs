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
    public sealed class FSGroupStrategyOptions
    {
        /// <summary>
        /// ranges are the allowed ranges of fs groups.  If you would like to force a single fs group then supply a single range with the same start and end. Required for MustRunAs.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1.IDRange> Ranges;
        /// <summary>
        /// rule is the strategy that will dictate what FSGroup is used in the SecurityContext.
        /// </summary>
        public readonly string Rule;

        [OutputConstructor]
        private FSGroupStrategyOptions(
            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1.IDRange> ranges,

            string rule)
        {
            Ranges = ranges;
            Rule = rule;
        }
    }
}
