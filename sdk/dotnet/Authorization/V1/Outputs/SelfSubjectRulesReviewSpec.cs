// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Authorization.V1
{

    [OutputType]
    public sealed class SelfSubjectRulesReviewSpec
    {
        /// <summary>
        /// Namespace to evaluate rules for. Required.
        /// </summary>
        public readonly string Namespace;

        [OutputConstructor]
        private SelfSubjectRulesReviewSpec(string @namespace)
        {
            Namespace = @namespace;
        }
    }
}
