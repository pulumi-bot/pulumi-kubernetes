// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Networking.V1Beta1
{

    [OutputType]
    public sealed class IngressSpec
    {
        /// <summary>
        /// A default backend capable of servicing requests that don't match any rule. At least one of 'backend' or 'rules' must be specified. This field is optional to allow the loadbalancer controller or defaulting logic to specify a global default.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Networking.V1Beta1.IngressBackend Backend;
        /// <summary>
        /// IngressClassName is the name of the IngressClass cluster resource. The associated IngressClass defines which controller will implement the resource. This replaces the deprecated `kubernetes.io/ingress.class` annotation. For backwards compatibility, when that annotation is set, it must be given precedence over this field. The controller may emit a warning if the field and annotation have different values. Implementations of this API should ignore Ingresses without a class specified. An IngressClass resource may be marked as default, which can be used to set a default value for this field. For more information, refer to the IngressClass documentation.
        /// </summary>
        public readonly string IngressClassName;
        /// <summary>
        /// A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Networking.V1Beta1.IngressRule> Rules;
        /// <summary>
        /// TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Networking.V1Beta1.IngressTLS> Tls;

        [OutputConstructor]
        private IngressSpec(
            Pulumi.Kubernetes.Types.Outputs.Networking.V1Beta1.IngressBackend backend,

            string ingressClassName,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Networking.V1Beta1.IngressRule> rules,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Networking.V1Beta1.IngressTLS> tls)
        {
            Backend = backend;
            IngressClassName = ingressClassName;
            Rules = rules;
            Tls = tls;
        }
    }
}
