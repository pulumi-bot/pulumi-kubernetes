// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Core.V1
{

    [OutputType]
    public sealed class NodeSelectorTerm
    {
        /// <summary>
        /// A list of node selector requirements by node's labels.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeSelectorRequirement> MatchExpressions;
        /// <summary>
        /// A list of node selector requirements by node's fields.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeSelectorRequirement> MatchFields;

        [OutputConstructor]
        private NodeSelectorTerm(
            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeSelectorRequirement> matchExpressions,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeSelectorRequirement> matchFields)
        {
            MatchExpressions = matchExpressions;
            MatchFields = matchFields;
        }
    }
}
