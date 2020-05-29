// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta1
{

    [OutputType]
    public sealed class StatefulSetStatus
    {
        /// <summary>
        /// collisionCount is the count of hash collisions for the StatefulSet. The StatefulSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.
        /// </summary>
        public readonly int CollisionCount;
        /// <summary>
        /// Represents the latest available observations of a statefulset's current state.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta1.StatefulSetCondition> Conditions;
        /// <summary>
        /// currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by currentRevision.
        /// </summary>
        public readonly int CurrentReplicas;
        /// <summary>
        /// currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [0,currentReplicas).
        /// </summary>
        public readonly string CurrentRevision;
        /// <summary>
        /// observedGeneration is the most recent generation observed for this StatefulSet. It corresponds to the StatefulSet's generation, which is updated on mutation by the API Server.
        /// </summary>
        public readonly int ObservedGeneration;
        /// <summary>
        /// readyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition.
        /// </summary>
        public readonly int ReadyReplicas;
        /// <summary>
        /// replicas is the number of Pods created by the StatefulSet controller.
        /// </summary>
        public readonly int Replicas;
        /// <summary>
        /// updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-updatedReplicas,replicas)
        /// </summary>
        public readonly string UpdateRevision;
        /// <summary>
        /// updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by updateRevision.
        /// </summary>
        public readonly int UpdatedReplicas;

        [OutputConstructor]
        private StatefulSetStatus(
            int collisionCount,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta1.StatefulSetCondition> conditions,

            int currentReplicas,

            string currentRevision,

            int observedGeneration,

            int readyReplicas,

            int replicas,

            string updateRevision,

            int updatedReplicas)
        {
            CollisionCount = collisionCount;
            Conditions = conditions;
            CurrentReplicas = currentReplicas;
            CurrentRevision = currentRevision;
            ObservedGeneration = observedGeneration;
            ReadyReplicas = readyReplicas;
            Replicas = replicas;
            UpdateRevision = updateRevision;
            UpdatedReplicas = updatedReplicas;
        }
    }
}
