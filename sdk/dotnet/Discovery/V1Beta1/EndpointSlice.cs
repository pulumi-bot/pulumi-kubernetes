// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Discovery.V1Beta1
{
    /// <summary>
    /// EndpointSlice represents a subset of the endpoints that implement a service. For a given service there may be multiple EndpointSlice objects, selected by labels, which must be joined to produce the full set of endpoints.
    /// </summary>
    [KubernetesResourceType("kubernetes:discovery.k8s.io/v1beta1:EndpointSlice")]
    public partial class EndpointSlice : KubernetesResource
    {
        /// <summary>
        /// addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.
        /// </summary>
        [Output("addressType")]
        public Output<string> AddressType { get; private set; } = null!;

        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.
        /// </summary>
        [Output("endpoints")]
        public Output<ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Discovery.V1Beta1.Endpoint>> Endpoints { get; private set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Standard object's metadata.
        /// </summary>
        [Output("metadata")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta> Metadata { get; private set; } = null!;

        /// <summary>
        /// ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates "all ports". Each slice may include a maximum of 100 ports.
        /// </summary>
        [Output("ports")]
        public Output<ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Discovery.V1Beta1.EndpointPort>> Ports { get; private set; } = null!;


        /// <summary>
        /// Create a EndpointSlice resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EndpointSlice(string name, Pulumi.Kubernetes.Types.Inputs.Discovery.V1Beta1.EndpointSliceArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:discovery.k8s.io/v1beta1:EndpointSlice", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }
        internal EndpointSlice(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:discovery.k8s.io/v1beta1:EndpointSlice", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private EndpointSlice(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes:discovery.k8s.io/v1beta1:EndpointSlice", name, null, MakeResourceOptions(options, id))
        {
        }

        private static Pulumi.Kubernetes.Types.Inputs.Discovery.V1Beta1.EndpointSliceArgs? MakeArgs(Pulumi.Kubernetes.Types.Inputs.Discovery.V1Beta1.EndpointSliceArgs? args)
        {
            args ??= new Pulumi.Kubernetes.Types.Inputs.Discovery.V1Beta1.EndpointSliceArgs();
            args.ApiVersion = "discovery.k8s.io/v1beta1";
            args.Kind = "EndpointSlice";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EndpointSlice resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EndpointSlice Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EndpointSlice(name, id, options);
        }
    }
}
namespace Pulumi.Kubernetes.Types.Inputs.Discovery.V1Beta1
{

    public class EndpointSliceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.
        /// </summary>
        [Input("addressType", required: true)]
        public Input<string> AddressType { get; set; } = null!;

        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        [Input("endpoints", required: true)]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Discovery.V1Beta1.EndpointArgs>? _endpoints;

        /// <summary>
        /// endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Discovery.V1Beta1.EndpointArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Pulumi.Kubernetes.Types.Inputs.Discovery.V1Beta1.EndpointArgs>());
            set => _endpoints = value;
        }

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Standard object's metadata.
        /// </summary>
        [Input("metadata")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs>? Metadata { get; set; }

        [Input("ports")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Discovery.V1Beta1.EndpointPortArgs>? _ports;

        /// <summary>
        /// ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates "all ports". Each slice may include a maximum of 100 ports.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Discovery.V1Beta1.EndpointPortArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Pulumi.Kubernetes.Types.Inputs.Discovery.V1Beta1.EndpointPortArgs>());
            set => _ports = value;
        }

        public EndpointSliceArgs()
        {
        }
    }
}
