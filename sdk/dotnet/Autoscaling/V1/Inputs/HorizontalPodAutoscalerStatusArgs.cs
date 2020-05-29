// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Autoscaling.V1
{

    /// <summary>
    /// current status of a horizontal pod autoscaler
    /// </summary>
    public class HorizontalPodAutoscalerStatusArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.
        /// </summary>
        [Input("currentCPUUtilizationPercentage")]
        public Input<int>? CurrentCPUUtilizationPercentage { get; set; }

        /// <summary>
        /// current number of replicas of pods managed by this autoscaler.
        /// </summary>
        [Input("currentReplicas", required: true)]
        public Input<int> CurrentReplicas { get; set; } = null!;

        /// <summary>
        /// desired number of replicas of pods managed by this autoscaler.
        /// </summary>
        [Input("desiredReplicas", required: true)]
        public Input<int> DesiredReplicas { get; set; } = null!;

        /// <summary>
        /// last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed.
        /// </summary>
        [Input("lastScaleTime")]
        public Input<string>? LastScaleTime { get; set; }

        /// <summary>
        /// most recent generation observed by this autoscaler.
        /// </summary>
        [Input("observedGeneration")]
        public Input<int>? ObservedGeneration { get; set; }

        public HorizontalPodAutoscalerStatusArgs()
        {
        }
    }
}
