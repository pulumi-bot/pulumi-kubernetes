// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2
{

    /// <summary>
    /// HorizontalPodAutoscalerStatus describes the current status of a horizontal pod autoscaler.
    /// </summary>
    public class HorizontalPodAutoscalerStatusArgs : Pulumi.ResourceArgs
    {
        [Input("conditions", required: true)]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2.HorizontalPodAutoscalerConditionArgs>? _conditions;

        /// <summary>
        /// conditions is the set of conditions required for this autoscaler to scale its target, and indicates whether or not those conditions are met.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2.HorizontalPodAutoscalerConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2.HorizontalPodAutoscalerConditionArgs>());
            set => _conditions = value;
        }

        [Input("currentMetrics")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2.MetricStatusArgs>? _currentMetrics;

        /// <summary>
        /// currentMetrics is the last read state of the metrics used by this autoscaler.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2.MetricStatusArgs> CurrentMetrics
        {
            get => _currentMetrics ?? (_currentMetrics = new InputList<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2.MetricStatusArgs>());
            set => _currentMetrics = value;
        }

        /// <summary>
        /// currentReplicas is current number of replicas of pods managed by this autoscaler, as last seen by the autoscaler.
        /// </summary>
        [Input("currentReplicas", required: true)]
        public Input<int> CurrentReplicas { get; set; } = null!;

        /// <summary>
        /// desiredReplicas is the desired number of replicas of pods managed by this autoscaler, as last calculated by the autoscaler.
        /// </summary>
        [Input("desiredReplicas", required: true)]
        public Input<int> DesiredReplicas { get; set; } = null!;

        /// <summary>
        /// lastScaleTime is the last time the HorizontalPodAutoscaler scaled the number of pods, used by the autoscaler to control how often the number of pods is changed.
        /// </summary>
        [Input("lastScaleTime")]
        public Input<string>? LastScaleTime { get; set; }

        /// <summary>
        /// observedGeneration is the most recent generation observed by this autoscaler.
        /// </summary>
        [Input("observedGeneration")]
        public Input<int>? ObservedGeneration { get; set; }

        public HorizontalPodAutoscalerStatusArgs()
        {
        }
    }
}
