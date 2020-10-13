// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta2
{

    [OutputType]
    public sealed class DaemonSetSpec
    {
        /// <summary>
        /// The minimum number of seconds for which a newly created DaemonSet pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready).
        /// </summary>
        public readonly int MinReadySeconds;
        /// <summary>
        /// The number of old history to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
        /// </summary>
        public readonly int RevisionHistoryLimit;
        /// <summary>
        /// A label query over pods that are managed by the daemon set. Must match in order to be controlled. It must match the pod template's labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelector Selector;
        /// <summary>
        /// An object that describes the pod that will be created. The DaemonSet will create exactly one copy of this pod on every node that matches the template's node selector (or on every node if no node selector is specified). More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.PodTemplateSpec Template;
        /// <summary>
        /// An update strategy to replace existing DaemonSet pods with new pods.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta2.DaemonSetUpdateStrategy UpdateStrategy;

        [OutputConstructor]
        private DaemonSetSpec(
            int minReadySeconds,

            int revisionHistoryLimit,

            Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelector selector,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.PodTemplateSpec template,

            Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta2.DaemonSetUpdateStrategy updateStrategy)
        {
            MinReadySeconds = minReadySeconds;
            RevisionHistoryLimit = revisionHistoryLimit;
            Selector = selector;
            Template = template;
            UpdateStrategy = updateStrategy;
        }
    }
}
