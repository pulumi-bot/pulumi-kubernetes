// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Batch.V1Beta1
{

    /// <summary>
    /// CronJobStatus represents the current state of a cron job.
    /// </summary>
    public class CronJobStatusArgs : Pulumi.ResourceArgs
    {
        [Input("active")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ObjectReferenceArgs>? _active;

        /// <summary>
        /// A list of pointers to currently running jobs.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ObjectReferenceArgs> Active
        {
            get => _active ?? (_active = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ObjectReferenceArgs>());
            set => _active = value;
        }

        /// <summary>
        /// Information when was the last time the job was successfully scheduled.
        /// </summary>
        [Input("lastScheduleTime")]
        public Input<string>? LastScheduleTime { get; set; }

        public CronJobStatusArgs()
        {
        }
    }
}
