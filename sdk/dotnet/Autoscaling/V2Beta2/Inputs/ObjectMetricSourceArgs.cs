// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2
{

    /// <summary>
    /// ObjectMetricSource indicates how to scale on a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
    /// </summary>
    public class ObjectMetricSourceArgs : Pulumi.ResourceArgs
    {
        [Input("describedObject", required: true)]
        public Input<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2.CrossVersionObjectReferenceArgs> DescribedObject { get; set; } = null!;

        /// <summary>
        /// metric identifies the target metric by name and selector
        /// </summary>
        [Input("metric", required: true)]
        public Input<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2.MetricIdentifierArgs> Metric { get; set; } = null!;

        /// <summary>
        /// target specifies the target value for the given metric
        /// </summary>
        [Input("target", required: true)]
        public Input<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2.MetricTargetArgs> Target { get; set; } = null!;

        public ObjectMetricSourceArgs()
        {
        }
    }
}
