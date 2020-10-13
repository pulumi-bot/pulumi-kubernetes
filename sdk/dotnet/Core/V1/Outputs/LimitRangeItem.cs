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
    public sealed class LimitRangeItem
    {
        /// <summary>
        /// Default resource requirement limit value by resource name if resource limit is omitted.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Default;
        /// <summary>
        /// DefaultRequest is the default resource requirement request value by resource name if resource request is omitted.
        /// </summary>
        public readonly ImmutableDictionary<string, string> DefaultRequest;
        /// <summary>
        /// Max usage constraints on this kind by resource name.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Max;
        /// <summary>
        /// MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> MaxLimitRequestRatio;
        /// <summary>
        /// Min usage constraints on this kind by resource name.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Min;
        /// <summary>
        /// Type of resource that this limit applies to.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private LimitRangeItem(
            ImmutableDictionary<string, string> @default,

            ImmutableDictionary<string, string> defaultRequest,

            ImmutableDictionary<string, string> max,

            ImmutableDictionary<string, string> maxLimitRequestRatio,

            ImmutableDictionary<string, string> min,

            string type)
        {
            Default = @default;
            DefaultRequest = defaultRequest;
            Max = max;
            MaxLimitRequestRatio = maxLimitRequestRatio;
            Min = min;
            Type = type;
        }
    }
}
