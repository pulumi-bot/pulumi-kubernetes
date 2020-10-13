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
    /// CustomResourceDefinitionNames indicates the names to serve this CustomResourceDefinition
    /// </summary>
    public class CustomResourceDefinitionNamesArgs : Pulumi.ResourceArgs
    {
        [Input("categories")]
        private InputList<string>? _categories;

        /// <summary>
        /// categories is a list of grouped resources this custom resource belongs to (e.g. 'all'). This is published in API discovery documents, and used by clients to support invocations like `kubectl get all`.
        /// </summary>
        public InputList<string> Categories
        {
            get => _categories ?? (_categories = new InputList<string>());
            set => _categories = value;
        }

        /// <summary>
        /// kind is the serialized kind of the resource. It is normally CamelCase and singular. Custom resource instances will use this value as the `kind` attribute in API calls.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// listKind is the serialized kind of the list for this resource. Defaults to "`kind`List".
        /// </summary>
        [Input("listKind")]
        public Input<string>? ListKind { get; set; }

        /// <summary>
        /// plural is the plural name of the resource to serve. The custom resources are served under `/apis/&lt;group&gt;/&lt;version&gt;/.../&lt;plural&gt;`. Must match the name of the CustomResourceDefinition (in the form `&lt;names.plural&gt;.&lt;group&gt;`). Must be all lowercase.
        /// </summary>
        [Input("plural", required: true)]
        public Input<string> Plural { get; set; } = null!;

        [Input("shortNames")]
        private InputList<string>? _shortNames;

        /// <summary>
        /// shortNames are short names for the resource, exposed in API discovery documents, and used by clients to support invocations like `kubectl get &lt;shortname&gt;`. It must be all lowercase.
        /// </summary>
        public InputList<string> ShortNames
        {
            get => _shortNames ?? (_shortNames = new InputList<string>());
            set => _shortNames = value;
        }

        /// <summary>
        /// singular is the singular name of the resource. It must be all lowercase. Defaults to lowercased `kind`.
        /// </summary>
        [Input("singular")]
        public Input<string>? Singular { get; set; }

        public CustomResourceDefinitionNamesArgs()
        {
        }
    }
}
