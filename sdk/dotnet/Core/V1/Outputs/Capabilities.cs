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
    public sealed class Capabilities
    {
        /// <summary>
        /// Added capabilities
        /// </summary>
        public readonly ImmutableArray<string> Add;
        /// <summary>
        /// Removed capabilities
        /// </summary>
        public readonly ImmutableArray<string> Drop;

        [OutputConstructor]
        private Capabilities(
            ImmutableArray<string> add,

            ImmutableArray<string> drop)
        {
            Add = add;
            Drop = drop;
        }
    }
}
