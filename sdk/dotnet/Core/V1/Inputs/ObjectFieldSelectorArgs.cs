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
    /// ObjectFieldSelector selects an APIVersioned field of an object.
    /// </summary>
    public class ObjectFieldSelectorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Version of the schema the FieldPath is written in terms of, defaults to "v1".
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        /// <summary>
        /// Path of the field to select in the specified API version.
        /// </summary>
        [Input("fieldPath", required: true)]
        public Input<string> FieldPath { get; set; } = null!;

        public ObjectFieldSelectorArgs()
        {
        }
    }
}
