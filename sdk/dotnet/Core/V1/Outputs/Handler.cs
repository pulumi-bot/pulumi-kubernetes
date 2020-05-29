// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Core.V1
{

    [OutputType]
    public sealed class Handler
    {
        /// <summary>
        /// One and only one of the following should be specified. Exec specifies the action to take.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.ExecAction Exec;
        /// <summary>
        /// HTTPGet specifies the http request to perform.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.HTTPGetAction HttpGet;
        /// <summary>
        /// TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.TCPSocketAction TcpSocket;

        [OutputConstructor]
        private Handler(
            Pulumi.Kubernetes.Types.Outputs.Core.V1.ExecAction exec,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.HTTPGetAction httpGet,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.TCPSocketAction tcpSocket)
        {
            Exec = exec;
            HttpGet = httpGet;
            TcpSocket = tcpSocket;
        }
    }
}
