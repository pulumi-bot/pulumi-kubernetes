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
    public sealed class DaemonSetUpdateStrategy
    {
        /// <summary>
        /// Rolling update config params. Present only if type = "RollingUpdate".
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1.RollingUpdateDaemonSet RollingUpdate;
        /// <summary>
        /// Type of daemon set update. Can be "RollingUpdate" or "OnDelete". Default is OnDelete.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private DaemonSetUpdateStrategy(
            Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1.RollingUpdateDaemonSet rollingUpdate,

            string type)
        {
            RollingUpdate = rollingUpdate;
            Type = type;
        }
    }
}
