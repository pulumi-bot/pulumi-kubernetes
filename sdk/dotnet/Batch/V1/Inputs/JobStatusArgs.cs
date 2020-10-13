// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Batch.V1
{

    /// <summary>
    /// JobStatus represents the current state of a Job.
    /// </summary>
    public class JobStatusArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of actively running pods.
        /// </summary>
        [Input("active")]
        public Input<int>? Active { get; set; }

        /// <summary>
        /// Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
        /// </summary>
        [Input("completionTime")]
        public Input<string>? CompletionTime { get; set; }

        [Input("conditions")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Batch.V1.JobConditionArgs>? _conditions;

        /// <summary>
        /// The latest available observations of an object's current state. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Batch.V1.JobConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Pulumi.Kubernetes.Types.Inputs.Batch.V1.JobConditionArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// The number of pods which reached phase Failed.
        /// </summary>
        [Input("failed")]
        public Input<int>? Failed { get; set; }

        /// <summary>
        /// Represents time when the job was acknowledged by the job controller. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// The number of pods which reached phase Succeeded.
        /// </summary>
        [Input("succeeded")]
        public Input<int>? Succeeded { get; set; }

        public JobStatusArgs()
        {
        }
    }
}
