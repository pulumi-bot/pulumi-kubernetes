// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Rbac.V1
{

    [OutputType]
    public sealed class AggregationRule
    {
        /// <summary>
        /// ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole's permissions will be added
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelector> ClusterRoleSelectors;

        [OutputConstructor]
        private AggregationRule(ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelector> clusterRoleSelectors)
        {
            ClusterRoleSelectors = clusterRoleSelectors;
        }
    }
}
