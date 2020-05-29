// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1
{

    [OutputType]
    public sealed class NetworkPolicySpec
    {
        /// <summary>
        /// List of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic matches at least one egress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy limits all outgoing traffic (and serves solely to ensure that the pods it selects are isolated by default). This field is beta-level in 1.8
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1.NetworkPolicyEgressRule> Egress;
        /// <summary>
        /// List of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies selecting the pod OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy does not allow any traffic (and serves solely to ensure that the pods it selects are isolated by default).
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1.NetworkPolicyIngressRule> Ingress;
        /// <summary>
        /// Selects the pods to which this NetworkPolicy object applies.  The array of ingress rules is applied to any pods selected by this field. Multiple network policies can select the same set of pods.  In this case, the ingress rules for each are combined additively. This field is NOT optional and follows standard label selector semantics. An empty podSelector matches all pods in this namespace.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelector PodSelector;
        /// <summary>
        /// List of rule types that the NetworkPolicy relates to. Valid options are "Ingress", "Egress", or "Ingress,Egress". If this field is not specified, it will default based on the existence of Ingress or Egress rules; policies that contain an Egress section are assumed to affect Egress, and all policies (whether or not they contain an Ingress section) are assumed to affect Ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ "Egress" ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include "Egress" (since such a policy would not include an Egress section and would otherwise default to just [ "Ingress" ]). This field is beta-level in 1.8
        /// </summary>
        public readonly ImmutableArray<string> PolicyTypes;

        [OutputConstructor]
        private NetworkPolicySpec(
            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1.NetworkPolicyEgressRule> egress,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1.NetworkPolicyIngressRule> ingress,

            Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelector podSelector,

            ImmutableArray<string> policyTypes)
        {
            Egress = egress;
            Ingress = ingress;
            PodSelector = podSelector;
            PolicyTypes = policyTypes;
        }
    }
}
