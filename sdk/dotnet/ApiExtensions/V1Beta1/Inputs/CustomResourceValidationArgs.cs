// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.ApiExtensions.V1Beta1
{

    /// <summary>
    /// CustomResourceValidation is a list of validation methods for CustomResources.
    /// </summary>
    public class CustomResourceValidationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// openAPIV3Schema is the OpenAPI v3 schema to use for validation and pruning.
        /// </summary>
        [Input("openAPIV3Schema")]
        public Input<Pulumi.Kubernetes.Types.Inputs.ApiExtensions.V1Beta1.JSONSchemaPropsArgs>? OpenAPIV3Schema { get; set; }

        public CustomResourceValidationArgs()
        {
        }
    }
}
