// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta2
{

    [OutputType]
    public sealed class StatefulSetUpdateStrategy
    {
        /// <summary>
        /// RollingUpdate is used to communicate parameters when Type is RollingUpdateStatefulSetStrategyType.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta2.RollingUpdateStatefulSetStrategy RollingUpdate;
        /// <summary>
        /// Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private StatefulSetUpdateStrategy(
            Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta2.RollingUpdateStatefulSetStrategy rollingUpdate,

            string type)
        {
            RollingUpdate = rollingUpdate;
            Type = type;
        }
    }
}
