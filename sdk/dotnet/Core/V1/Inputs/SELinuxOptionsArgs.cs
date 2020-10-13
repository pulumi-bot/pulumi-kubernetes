// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Core.V1
{

    /// <summary>
    /// SELinuxOptions are the labels to be applied to the container
    /// </summary>
    public class SELinuxOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Level is SELinux level label that applies to the container.
        /// </summary>
        [Input("level")]
        public Input<string>? Level { get; set; }

        /// <summary>
        /// Role is a SELinux role label that applies to the container.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Type is a SELinux type label that applies to the container.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// User is a SELinux user label that applies to the container.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public SELinuxOptionsArgs()
        {
        }
    }
}
