// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.AuditRegistraion.V1Alpha1
{

    [OutputType]
    public sealed class Webhook
    {
        /// <summary>
        /// ClientConfig holds the connection parameters for the webhook required
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.AuditRegistraion.V1Alpha1.WebhookClientConfig ClientConfig;
        /// <summary>
        /// Throttle holds the options for throttling the webhook
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.AuditRegistraion.V1Alpha1.WebhookThrottleConfig Throttle;

        [OutputConstructor]
        private Webhook(
            Pulumi.Kubernetes.Types.Outputs.AuditRegistraion.V1Alpha1.WebhookClientConfig clientConfig,

            Pulumi.Kubernetes.Types.Outputs.AuditRegistraion.V1Alpha1.WebhookThrottleConfig throttle)
        {
            ClientConfig = clientConfig;
            Throttle = throttle;
        }
    }
}
