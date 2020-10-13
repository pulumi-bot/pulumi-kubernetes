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
    public sealed class PodTemplateSpec
    {
        /// <summary>
        /// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta Metadata;
        /// <summary>
        /// Specification of the desired behavior of the pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.PodSpec Spec;

        [OutputConstructor]
        private PodTemplateSpec(
            Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta metadata,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.PodSpec spec)
        {
            Metadata = metadata;
            Spec = spec;
        }
    }
}
