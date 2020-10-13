// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1Beta1
{

    [OutputType]
    public sealed class CustomResourceValidation
    {
        /// <summary>
        /// openAPIV3Schema is the OpenAPI v3 schema to use for validation and pruning.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1Beta1.JSONSchemaProps OpenAPIV3Schema;

        [OutputConstructor]
        private CustomResourceValidation(Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1Beta1.JSONSchemaProps openAPIV3Schema)
        {
            OpenAPIV3Schema = openAPIV3Schema;
        }
    }
}
