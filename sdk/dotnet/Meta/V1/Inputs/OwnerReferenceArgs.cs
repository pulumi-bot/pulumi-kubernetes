// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Meta.V1
{

    /// <summary>
    /// OwnerReference contains enough information to let you identify an owning object. An owning object must be in the same namespace as the dependent, or be cluster-scoped, so there is no namespace field.
    /// </summary>
    public class OwnerReferenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// API version of the referent.
        /// </summary>
        [Input("apiVersion", required: true)]
        public Input<string> ApiVersion { get; set; } = null!;

        /// <summary>
        /// If true, AND if the owner has the "foregroundDeletion" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs "delete" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.
        /// </summary>
        [Input("blockOwnerDeletion")]
        public Input<bool>? BlockOwnerDeletion { get; set; }

        /// <summary>
        /// If true, this reference points to the managing controller.
        /// </summary>
        [Input("controller")]
        public Input<bool>? Controller { get; set; }

        /// <summary>
        /// Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids
        /// </summary>
        [Input("uid", required: true)]
        public Input<string> Uid { get; set; } = null!;

        public OwnerReferenceArgs()
        {
        }
    }
}
