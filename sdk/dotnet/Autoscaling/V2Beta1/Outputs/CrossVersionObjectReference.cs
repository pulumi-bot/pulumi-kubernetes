// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta1
{

    [OutputType]
    public sealed class CrossVersionObjectReference
    {
        /// <summary>
        /// API version of the referent
        /// </summary>
        public readonly string ApiVersion;
        /// <summary>
        /// Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds"
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private CrossVersionObjectReference(
            string apiVersion,

            string kind,

            string name)
        {
            ApiVersion = apiVersion;
            Kind = kind;
            Name = name;
        }
    }
}
