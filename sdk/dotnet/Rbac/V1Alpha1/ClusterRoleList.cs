// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Rbac.V1Alpha1
{
    /// <summary>
    /// ClusterRoleList is a collection of ClusterRoles. Deprecated in v1.17 in favor of
    /// rbac.authorization.k8s.io/v1 ClusterRoles, and will no longer be served in v1.22.
    /// </summary>
    public partial class ClusterRoleList : KubernetesResource
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers
        /// should convert recognized schemas to the latest internal value, and may reject
        /// unrecognized values. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// Items is a list of ClusterRoles
        /// </summary>
        [Output("items")]
        public Output<ImmutableArray<Types.Outputs.Rbac.V1Alpha1.ClusterRole>> Items { get; private set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers
        /// may infer this from the endpoint the client submits requests to. Cannot be updated. In
        /// CamelCase. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Standard object's metadata.
        /// </summary>
        [Output("metadata")]
        public Output<Types.Outputs.Meta.V1.ListMeta> Metadata { get; private set; } = null!;


        /// <summary>
        /// Create a ClusterRoleList resource with the given unique name, arguments, and options.
        /// </summary>
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterRoleList(string name, Types.Inputs.Rbac.V1Alpha1.ClusterRoleListArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRoleList", name, SetAPIKindAndVersion(args), MakeOptions(options))
        {
        }

        internal ClusterRoleList(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRoleList", name, dictionary, options)
        {
        }

        private static ResourceArgs SetAPIKindAndVersion(Types.Inputs.Rbac.V1Alpha1.ClusterRoleListArgs? args)
        {
            args ??= new Types.Inputs.Rbac.V1Alpha1.ClusterRoleListArgs();
            args.ApiVersion = "rbac.authorization.k8s.io/v1alpha1";
            args.Kind = "ClusterRoleList";
            return args;
        }

        private static CustomResourceOptions? MakeOptions(CustomResourceOptions? options)
        {
            return options;
        }

        /// <summary>
        /// Get an existing ClusterRoleList resource's state with the given name and ID.
        /// </summary>
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterRoleList Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ClusterRoleList(name, default(Types.Inputs.Rbac.V1Alpha1.ClusterRoleListArgs),
                CustomResourceOptions.Merge(options, new CustomResourceOptions {Id = id}));
        }
    }
}
