// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * SubjectAccessReview checks whether or not a user or group can perform an action.
 */
export class SubjectAccessReview extends pulumi.CustomResource {
    /**
     * Get an existing SubjectAccessReview resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SubjectAccessReview {
        return new SubjectAccessReview(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:authorization.k8s.io/v1:SubjectAccessReview';

    /**
     * Returns true if the given object is an instance of SubjectAccessReview.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SubjectAccessReview {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SubjectAccessReview.__pulumiType;
    }

    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<"authorization.k8s.io/v1">;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<"SubjectAccessReview">;
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ObjectMeta>;
    /**
     * Spec holds information about the request being evaluated
     */
    public readonly spec!: pulumi.Output<outputs.authorization.v1.SubjectAccessReviewSpec>;
    /**
     * Status is filled in by the server and indicates whether the request is allowed or not
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.authorization.v1.SubjectAccessReviewStatus>;

    /**
     * Create a SubjectAccessReview resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SubjectAccessReviewArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.spec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'spec'");
            }
            inputs["apiVersion"] = "authorization.k8s.io/v1";
            inputs["kind"] = "SubjectAccessReview";
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["spec"] = args ? args.spec : undefined;
            inputs["status"] = undefined /*out*/;
        } else {
            inputs["apiVersion"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["metadata"] = undefined /*out*/;
            inputs["spec"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "kubernetes:authorization.k8s.io/v1beta1:SubjectAccessReview" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(SubjectAccessReview.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SubjectAccessReview resource.
 */
export interface SubjectAccessReviewArgs {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    readonly apiVersion?: pulumi.Input<"authorization.k8s.io/v1">;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    readonly kind?: pulumi.Input<"SubjectAccessReview">;
    readonly metadata?: pulumi.Input<inputs.meta.v1.ObjectMeta>;
    /**
     * Spec holds information about the request being evaluated
     */
    readonly spec: pulumi.Input<inputs.authorization.v1.SubjectAccessReviewSpec>;
}
