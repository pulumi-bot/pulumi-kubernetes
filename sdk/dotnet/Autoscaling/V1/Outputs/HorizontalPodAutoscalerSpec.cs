// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Autoscaling.V1
{

    [OutputType]
    public sealed class HorizontalPodAutoscalerSpec
    {
        /// <summary>
        /// upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
        /// </summary>
        public readonly int MaxReplicas;
        /// <summary>
        /// minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
        /// </summary>
        public readonly int MinReplicas;
        /// <summary>
        /// reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V1.CrossVersionObjectReference ScaleTargetRef;
        /// <summary>
        /// target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
        /// </summary>
        public readonly int TargetCPUUtilizationPercentage;

        [OutputConstructor]
        private HorizontalPodAutoscalerSpec(
            int maxReplicas,

            int minReplicas,

            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V1.CrossVersionObjectReference scaleTargetRef,

            int targetCPUUtilizationPercentage)
        {
            MaxReplicas = maxReplicas;
            MinReplicas = minReplicas;
            ScaleTargetRef = scaleTargetRef;
            TargetCPUUtilizationPercentage = targetCPUUtilizationPercentage;
        }
    }
}
