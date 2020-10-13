// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.AuditRegistraion.V1Alpha1
{

    [OutputType]
    public sealed class AuditSinkSpec
    {
        /// <summary>
        /// Policy defines the policy for selecting which events should be sent to the webhook required
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.AuditRegistraion.V1Alpha1.Policy Policy;
        /// <summary>
        /// Webhook to send events required
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.AuditRegistraion.V1Alpha1.Webhook Webhook;

        [OutputConstructor]
        private AuditSinkSpec(
            Pulumi.Kubernetes.Types.Outputs.AuditRegistraion.V1Alpha1.Policy policy,

            Pulumi.Kubernetes.Types.Outputs.AuditRegistraion.V1Alpha1.Webhook webhook)
        {
            Policy = policy;
            Webhook = webhook;
        }
    }
}
