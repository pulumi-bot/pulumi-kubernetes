// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Authorization.V1Beta1
{

    /// <summary>
    /// NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface
    /// </summary>
    public class NonResourceAttributesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path is the URL path of the request
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Verb is the standard HTTP verb
        /// </summary>
        [Input("verb")]
        public Input<string>? Verb { get; set; }

        public NonResourceAttributesArgs()
        {
        }
    }
}
