// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Alpha1
{

    [OutputType]
    public sealed class UserSubject
    {
        /// <summary>
        /// `name` is the username that matches, or "*" to match all usernames. Required.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private UserSubject(string name)
        {
            Name = name;
        }
    }
}
