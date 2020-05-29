// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1
{

    [OutputType]
    public sealed class ReplicaSetStatus
    {
        /// <summary>
        /// The number of available replicas (ready for at least minReadySeconds) for this replica set.
        /// </summary>
        public readonly int AvailableReplicas;
        /// <summary>
        /// Represents the latest available observations of a replica set's current state.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1.ReplicaSetCondition> Conditions;
        /// <summary>
        /// The number of pods that have labels matching the labels of the pod template of the replicaset.
        /// </summary>
        public readonly int FullyLabeledReplicas;
        /// <summary>
        /// ObservedGeneration reflects the generation of the most recently observed ReplicaSet.
        /// </summary>
        public readonly int ObservedGeneration;
        /// <summary>
        /// The number of ready replicas for this replica set.
        /// </summary>
        public readonly int ReadyReplicas;
        /// <summary>
        /// Replicas is the most recently oberved number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
        /// </summary>
        public readonly int Replicas;

        [OutputConstructor]
        private ReplicaSetStatus(
            int availableReplicas,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1.ReplicaSetCondition> conditions,

            int fullyLabeledReplicas,

            int observedGeneration,

            int readyReplicas,

            int replicas)
        {
            AvailableReplicas = availableReplicas;
            Conditions = conditions;
            FullyLabeledReplicas = fullyLabeledReplicas;
            ObservedGeneration = observedGeneration;
            ReadyReplicas = readyReplicas;
            Replicas = replicas;
        }
    }
}
