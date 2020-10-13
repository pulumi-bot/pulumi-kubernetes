// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1
{

    [OutputType]
    public sealed class CustomResourceDefinitionVersion
    {
        /// <summary>
        /// additionalPrinterColumns specifies additional columns returned in Table output. See https://kubernetes.io/docs/reference/using-api/api-concepts/#receiving-resources-as-tables for details. If no columns are specified, a single column displaying the age of the custom resource is used.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1.CustomResourceColumnDefinition> AdditionalPrinterColumns;
        /// <summary>
        /// deprecated indicates this version of the custom resource API is deprecated. When set to true, API requests to this version receive a warning header in the server response. Defaults to false.
        /// </summary>
        public readonly bool Deprecated;
        /// <summary>
        /// deprecationWarning overrides the default warning returned to API clients. May only be set when `deprecated` is true. The default warning indicates this version is deprecated and recommends use of the newest served version of equal or greater stability, if one exists.
        /// </summary>
        public readonly string DeprecationWarning;
        /// <summary>
        /// name is the version name, e.g. “v1”, “v2beta1”, etc. The custom resources are served under this version at `/apis/&lt;group&gt;/&lt;version&gt;/...` if `served` is true.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// schema describes the schema used for validation, pruning, and defaulting of this version of the custom resource.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1.CustomResourceValidation Schema;
        /// <summary>
        /// served is a flag enabling/disabling this version from being served via REST APIs
        /// </summary>
        public readonly bool Served;
        /// <summary>
        /// storage indicates this version should be used when persisting custom resources to storage. There must be exactly one version with storage=true.
        /// </summary>
        public readonly bool Storage;
        /// <summary>
        /// subresources specify what subresources this version of the defined custom resource have.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1.CustomResourceSubresources Subresources;

        [OutputConstructor]
        private CustomResourceDefinitionVersion(
            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1.CustomResourceColumnDefinition> additionalPrinterColumns,

            bool deprecated,

            string deprecationWarning,

            string name,

            Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1.CustomResourceValidation schema,

            bool served,

            bool storage,

            Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1.CustomResourceSubresources subresources)
        {
            AdditionalPrinterColumns = additionalPrinterColumns;
            Deprecated = deprecated;
            DeprecationWarning = deprecationWarning;
            Name = name;
            Schema = schema;
            Served = served;
            Storage = storage;
            Subresources = subresources;
        }
    }
}
