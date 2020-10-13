// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2
{

    [OutputType]
    public sealed class HorizontalPodAutoscalerStatus
    {
        /// <summary>
        /// conditions is the set of conditions required for this autoscaler to scale its target, and indicates whether or not those conditions are met.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2.HorizontalPodAutoscalerCondition> Conditions;
        /// <summary>
        /// currentMetrics is the last read state of the metrics used by this autoscaler.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2.MetricStatus> CurrentMetrics;
        /// <summary>
        /// currentReplicas is current number of replicas of pods managed by this autoscaler, as last seen by the autoscaler.
        /// </summary>
        public readonly int CurrentReplicas;
        /// <summary>
        /// desiredReplicas is the desired number of replicas of pods managed by this autoscaler, as last calculated by the autoscaler.
        /// </summary>
        public readonly int DesiredReplicas;
        /// <summary>
        /// lastScaleTime is the last time the HorizontalPodAutoscaler scaled the number of pods, used by the autoscaler to control how often the number of pods is changed.
        /// </summary>
        public readonly string LastScaleTime;
        /// <summary>
        /// observedGeneration is the most recent generation observed by this autoscaler.
        /// </summary>
        public readonly int ObservedGeneration;

        [OutputConstructor]
        private HorizontalPodAutoscalerStatus(
            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2.HorizontalPodAutoscalerCondition> conditions,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2.MetricStatus> currentMetrics,

            int currentReplicas,

            int desiredReplicas,

            string lastScaleTime,

            int observedGeneration)
        {
            Conditions = conditions;
            CurrentMetrics = currentMetrics;
            CurrentReplicas = currentReplicas;
            DesiredReplicas = desiredReplicas;
            LastScaleTime = lastScaleTime;
            ObservedGeneration = observedGeneration;
        }
    }
}
