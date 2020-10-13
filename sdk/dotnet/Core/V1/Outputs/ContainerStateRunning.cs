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
    public sealed class ContainerStateRunning
    {
        /// <summary>
        /// Time at which the container was last (re-)started
        /// </summary>
        public readonly string StartedAt;

        [OutputConstructor]
        private ContainerStateRunning(string startedAt)
        {
            StartedAt = startedAt;
        }
    }
}
