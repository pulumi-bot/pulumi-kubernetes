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
    public sealed class PodDNSConfigOption
    {
        /// <summary>
        /// Required.
        /// </summary>
        public readonly string Name;
        public readonly string Value;

        [OutputConstructor]
        private PodDNSConfigOption(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}
