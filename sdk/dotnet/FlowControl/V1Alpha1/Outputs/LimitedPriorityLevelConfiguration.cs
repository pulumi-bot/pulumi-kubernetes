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
    public sealed class LimitedPriorityLevelConfiguration
    {
        /// <summary>
        /// `assuredConcurrencyShares` (ACS) configures the execution limit, which is a limit on the number of requests of this priority level that may be exeucting at a given time.  ACS must be a positive number. The server's concurrency limit (SCL) is divided among the concurrency-controlled priority levels in proportion to their assured concurrency shares. This produces the assured concurrency value (ACV) --- the number of requests that may be executing at a time --- for each such priority level:
        /// 
        ///             ACV(l) = ceil( SCL * ACS(l) / ( sum[priority levels k] ACS(k) ) )
        /// 
        /// bigger numbers of ACS mean more reserved concurrent requests (at the expense of every other PL). This field has a default value of 30.
        /// </summary>
        public readonly int AssuredConcurrencyShares;
        /// <summary>
        /// `limitResponse` indicates what to do with requests that can not be executed right now
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Alpha1.LimitResponse LimitResponse;

        [OutputConstructor]
        private LimitedPriorityLevelConfiguration(
            int assuredConcurrencyShares,

            Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Alpha1.LimitResponse limitResponse)
        {
            AssuredConcurrencyShares = assuredConcurrencyShares;
            LimitResponse = limitResponse;
        }
    }
}
