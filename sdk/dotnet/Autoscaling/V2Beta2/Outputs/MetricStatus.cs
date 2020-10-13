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
    public sealed class MetricStatus
    {
        /// <summary>
        /// external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2.ExternalMetricStatus External;
        /// <summary>
        /// object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2.ObjectMetricStatus Object;
        /// <summary>
        /// pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2.PodsMetricStatus Pods;
        /// <summary>
        /// resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2.ResourceMetricStatus Resource;
        /// <summary>
        /// type is the type of metric source.  It will be one of "Object", "Pods" or "Resource", each corresponds to a matching field in the object.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private MetricStatus(
            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2.ExternalMetricStatus external,

            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2.ObjectMetricStatus @object,

            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2.PodsMetricStatus pods,

            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2.ResourceMetricStatus resource,

            string type)
        {
            External = external;
            Object = @object;
            Pods = pods;
            Resource = resource;
            Type = type;
        }
    }
}
