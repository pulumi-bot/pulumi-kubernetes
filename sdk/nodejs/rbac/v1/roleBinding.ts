// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * RoleBinding references a role, but does not contain it.  It can reference a Role in the same namespace or a ClusterRole in the global namespace. It adds who information via Subjects and namespace information by which namespace it exists in.  RoleBindings in a given namespace only have effect in that namespace.
 */
export class RoleBinding extends pulumi.CustomResource {
    /**
     * Get an existing RoleBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RoleBinding {
        return new RoleBinding(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:rbac.authorization.k8s.io/v1:RoleBinding';

    /**
     * Returns true if the given object is an instance of RoleBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RoleBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RoleBinding.__pulumiType;
    }

    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<"rbac.authorization.k8s.io/v1">;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<"RoleBinding">;
    /**
     * Standard object's metadata.
     */
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ObjectMeta>;
    /**
     * RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
     */
    public readonly roleRef!: pulumi.Output<outputs.rbac.v1.RoleRef>;
    /**
     * Subjects holds references to the objects the role applies to.
     */
    public readonly subjects!: pulumi.Output<outputs.rbac.v1.Subject[]>;

    /**
     * Create a RoleBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RoleBindingArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.roleRef === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'roleRef'");
            }
            inputs["apiVersion"] = "rbac.authorization.k8s.io/v1";
            inputs["kind"] = "RoleBinding";
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["roleRef"] = args ? args.roleRef : undefined;
            inputs["subjects"] = args ? args.subjects : undefined;
        } else {
            inputs["apiVersion"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["metadata"] = undefined /*out*/;
            inputs["roleRef"] = undefined /*out*/;
            inputs["subjects"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "kubernetes:rbac.authorization.k8s.io/v1alpha1:RoleBinding" }, { type: "kubernetes:rbac.authorization.k8s.io/v1beta1:RoleBinding" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(RoleBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RoleBinding resource.
 */
export interface RoleBindingArgs {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    readonly apiVersion?: pulumi.Input<"rbac.authorization.k8s.io/v1" | undefined>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    readonly kind?: pulumi.Input<"RoleBinding" | undefined>;
    /**
     * Standard object's metadata.
     */
    readonly metadata?: pulumi.Input<inputs.meta.v1.ObjectMeta | undefined>;
    /**
     * RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
     */
    readonly roleRef: pulumi.Input<inputs.rbac.v1.RoleRef>;
    /**
     * Subjects holds references to the objects the role applies to.
     */
    readonly subjects?: pulumi.Input<pulumi.Input<inputs.rbac.v1.Subject>[] | undefined>;
}
