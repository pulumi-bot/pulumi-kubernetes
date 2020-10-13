// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Batch.V1
{

    [OutputType]
    public sealed class JobStatus
    {
        /// <summary>
        /// The number of actively running pods.
        /// </summary>
        public readonly int Active;
        /// <summary>
        /// Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
        /// </summary>
        public readonly string CompletionTime;
        /// <summary>
        /// The latest available observations of an object's current state. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Batch.V1.JobCondition> Conditions;
        /// <summary>
        /// The number of pods which reached phase Failed.
        /// </summary>
        public readonly int Failed;
        /// <summary>
        /// Represents time when the job was acknowledged by the job controller. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// The number of pods which reached phase Succeeded.
        /// </summary>
        public readonly int Succeeded;

        [OutputConstructor]
        private JobStatus(
            int active,

            string completionTime,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Batch.V1.JobCondition> conditions,

            int failed,

            string startTime,

            int succeeded)
        {
            Active = active;
            CompletionTime = completionTime;
            Conditions = conditions;
            Failed = failed;
            StartTime = startTime;
            Succeeded = succeeded;
        }
    }
}
