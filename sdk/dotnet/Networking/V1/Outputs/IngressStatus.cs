// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Networking.V1
{

    [OutputType]
    public sealed class IngressStatus
    {
        /// <summary>
        /// LoadBalancer contains the current status of the load-balancer.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.LoadBalancerStatus LoadBalancer;

        [OutputConstructor]
        private IngressStatus(Pulumi.Kubernetes.Types.Outputs.Core.V1.LoadBalancerStatus loadBalancer)
        {
            LoadBalancer = loadBalancer;
        }
    }
}
