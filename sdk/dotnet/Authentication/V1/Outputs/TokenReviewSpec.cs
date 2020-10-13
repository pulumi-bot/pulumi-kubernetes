// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Authentication.V1
{

    [OutputType]
    public sealed class TokenReviewSpec
    {
        /// <summary>
        /// Audiences is a list of the identifiers that the resource server presented with the token identifies as. Audience-aware token authenticators will verify that the token was intended for at least one of the audiences in this list. If no audiences are provided, the audience will default to the audience of the Kubernetes apiserver.
        /// </summary>
        public readonly ImmutableArray<string> Audiences;
        /// <summary>
        /// Token is the opaque bearer token.
        /// </summary>
        public readonly string Token;

        [OutputConstructor]
        private TokenReviewSpec(
            ImmutableArray<string> audiences,

            string token)
        {
            Audiences = audiences;
            Token = token;
        }
    }
}
