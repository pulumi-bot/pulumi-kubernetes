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
    public sealed class TopologySelectorLabelRequirement
    {
        /// <summary>
        /// The label key that the selector applies to.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// An array of string values. One value must match the label to be selected. Each entry in Values is ORed.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private TopologySelectorLabelRequirement(
            string key,

            ImmutableArray<string> values)
        {
            Key = key;
            Values = values;
        }
    }
}
