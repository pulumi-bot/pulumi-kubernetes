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
    public sealed class ExternalMetricStatus
    {
        /// <summary>
        /// current contains the current value for the given metric
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2.MetricValueStatus Current;
        /// <summary>
        /// metric identifies the target metric by name and selector
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2.MetricIdentifier Metric;

        [OutputConstructor]
        private ExternalMetricStatus(
            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2.MetricValueStatus current,

            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2.MetricIdentifier metric)
        {
            Current = current;
            Metric = metric;
        }
    }
}
