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
    public sealed class RollbackConfig
    {
        /// <summary>
        /// The revision to rollback to. If set to 0, rollback to the last revision.
        /// </summary>
        public readonly int Revision;

        [OutputConstructor]
        private RollbackConfig(int revision)
        {
            Revision = revision;
        }
    }
}
