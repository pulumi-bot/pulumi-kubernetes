// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Extensions.V1Beta1
{

    /// <summary>
    /// DEPRECATED.
    /// </summary>
    public class RollbackConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The revision to rollback to. If set to 0, rollback to the last revision.
        /// </summary>
        [Input("revision")]
        public Input<int>? Revision { get; set; }

        public RollbackConfigArgs()
        {
        }
    }
}
