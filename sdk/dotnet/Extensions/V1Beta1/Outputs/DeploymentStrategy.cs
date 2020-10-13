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
    public sealed class DeploymentStrategy
    {
        /// <summary>
        /// Rolling update config params. Present only if DeploymentStrategyType = RollingUpdate.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1.RollingUpdateDeployment RollingUpdate;
        /// <summary>
        /// Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private DeploymentStrategy(
            Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1.RollingUpdateDeployment rollingUpdate,

            string type)
        {
            RollingUpdate = rollingUpdate;
            Type = type;
        }
    }
}
