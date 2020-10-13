// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Scheduling.V1
{

    [OutputType]
    public sealed class PriorityClass
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        public readonly string ApiVersion;
        /// <summary>
        /// description is an arbitrary string that usually provides guidelines on when this priority class should be used.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
        /// </summary>
        public readonly bool GlobalDefault;
        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta Metadata;
        /// <summary>
        /// PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is beta-level, gated by the NonPreemptingPriority feature-gate.
        /// </summary>
        public readonly string PreemptionPolicy;
        /// <summary>
        /// The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
        /// </summary>
        public readonly int Value;

        [OutputConstructor]
        private PriorityClass(
            string apiVersion,

            string description,

            bool globalDefault,

            string kind,

            Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta metadata,

            string preemptionPolicy,

            int value)
        {
            ApiVersion = apiVersion;
            Description = description;
            GlobalDefault = globalDefault;
            Kind = kind;
            Metadata = metadata;
            PreemptionPolicy = preemptionPolicy;
            Value = value;
        }
    }
}
