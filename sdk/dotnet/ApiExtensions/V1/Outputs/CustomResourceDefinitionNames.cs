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
    public sealed class CustomResourceDefinitionNames
    {
        /// <summary>
        /// categories is a list of grouped resources this custom resource belongs to (e.g. 'all'). This is published in API discovery documents, and used by clients to support invocations like `kubectl get all`.
        /// </summary>
        public readonly ImmutableArray<string> Categories;
        /// <summary>
        /// kind is the serialized kind of the resource. It is normally CamelCase and singular. Custom resource instances will use this value as the `kind` attribute in API calls.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// listKind is the serialized kind of the list for this resource. Defaults to "`kind`List".
        /// </summary>
        public readonly string ListKind;
        /// <summary>
        /// plural is the plural name of the resource to serve. The custom resources are served under `/apis/&lt;group&gt;/&lt;version&gt;/.../&lt;plural&gt;`. Must match the name of the CustomResourceDefinition (in the form `&lt;names.plural&gt;.&lt;group&gt;`). Must be all lowercase.
        /// </summary>
        public readonly string Plural;
        /// <summary>
        /// shortNames are short names for the resource, exposed in API discovery documents, and used by clients to support invocations like `kubectl get &lt;shortname&gt;`. It must be all lowercase.
        /// </summary>
        public readonly ImmutableArray<string> ShortNames;
        /// <summary>
        /// singular is the singular name of the resource. It must be all lowercase. Defaults to lowercased `kind`.
        /// </summary>
        public readonly string Singular;

        [OutputConstructor]
        private CustomResourceDefinitionNames(
            ImmutableArray<string> categories,

            string kind,

            string listKind,

            string plural,

            ImmutableArray<string> shortNames,

            string singular)
        {
            Categories = categories;
            Kind = kind;
            ListKind = listKind;
            Plural = plural;
            ShortNames = shortNames;
            Singular = singular;
        }
    }
}
