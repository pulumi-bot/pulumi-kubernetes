// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Authorization.V1Beta1
{

    [OutputType]
    public sealed class NonResourceAttributes
    {
        /// <summary>
        /// Path is the URL path of the request
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Verb is the standard HTTP verb
        /// </summary>
        public readonly string Verb;

        [OutputConstructor]
        private NonResourceAttributes(
            string path,

            string verb)
        {
            Path = path;
            Verb = verb;
        }
    }
}
