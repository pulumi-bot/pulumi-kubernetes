// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Alpha1
{

    [OutputType]
    public sealed class PriorityLevelConfigurationCondition
    {
        /// <summary>
        /// `lastTransitionTime` is the last time the condition transitioned from one status to another.
        /// </summary>
        public readonly string LastTransitionTime;
        /// <summary>
        /// `message` is a human-readable message indicating details about last transition.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// `reason` is a unique, one-word, CamelCase reason for the condition's last transition.
        /// </summary>
        public readonly string Reason;
        /// <summary>
        /// `status` is the status of the condition. Can be True, False, Unknown. Required.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// `type` is the type of the condition. Required.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private PriorityLevelConfigurationCondition(
            string lastTransitionTime,

            string message,

            string reason,

            string status,

            string type)
        {
            LastTransitionTime = lastTransitionTime;
            Message = message;
            Reason = reason;
            Status = status;
            Type = type;
        }
    }
}
