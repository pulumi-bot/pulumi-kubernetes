// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Core.V1
{

    /// <summary>
    /// LoadBalancerIngress represents the status of a load-balancer ingress point: traffic intended for the service should be sent to an ingress point.
    /// </summary>
    public class LoadBalancerIngressArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        public LoadBalancerIngressArgs()
        {
        }
    }
}
