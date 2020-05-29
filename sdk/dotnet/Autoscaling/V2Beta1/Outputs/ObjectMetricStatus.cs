// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta1
{

    [OutputType]
    public sealed class ObjectMetricStatus
    {
        /// <summary>
        /// averageValue is the current value of the average of the metric across all relevant pods (as a quantity)
        /// </summary>
        public readonly string AverageValue;
        /// <summary>
        /// currentValue is the current value of the metric (as a quantity).
        /// </summary>
        public readonly string CurrentValue;
        /// <summary>
        /// metricName is the name of the metric in question.
        /// </summary>
        public readonly string MetricName;
        /// <summary>
        /// selector is the string-encoded form of a standard kubernetes label selector for the given metric When set in the ObjectMetricSource, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelector Selector;
        /// <summary>
        /// target is the described Kubernetes object.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta1.CrossVersionObjectReference Target;

        [OutputConstructor]
        private ObjectMetricStatus(
            string averageValue,

            string currentValue,

            string metricName,

            Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelector selector,

            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta1.CrossVersionObjectReference target)
        {
            AverageValue = averageValue;
            CurrentValue = currentValue;
            MetricName = metricName;
            Selector = selector;
            Target = target;
        }
    }
}
