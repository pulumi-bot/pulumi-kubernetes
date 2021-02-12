// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * AuditSinkList is a list of AuditSink items.
 */
export class AuditSinkList extends pulumi.CustomResource {
    /**
     * Get an existing AuditSinkList resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AuditSinkList {
        return new AuditSinkList(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:auditregistration.k8s.io/v1alpha1:AuditSinkList';

    /**
     * Returns true if the given object is an instance of AuditSinkList.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuditSinkList {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuditSinkList.__pulumiType;
    }

    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<"auditregistration.k8s.io/v1alpha1">;
    /**
     * List of audit configurations.
     */
    public readonly items!: pulumi.Output<outputs.auditregistration.v1alpha1.AuditSink[]>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<"AuditSinkList">;
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ListMeta>;

    /**
     * Create a AuditSinkList resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AuditSinkListArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.items === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'items'");
            }
            inputs["apiVersion"] = "auditregistration.k8s.io/v1alpha1";
            inputs["items"] = args ? args.items : undefined;
            inputs["kind"] = "AuditSinkList";
            inputs["metadata"] = args ? args.metadata : undefined;
        } else {
            inputs["apiVersion"] = undefined /*out*/;
            inputs["items"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["metadata"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AuditSinkList.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AuditSinkList resource.
 */
export interface AuditSinkListArgs {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    readonly apiVersion?: pulumi.Input<"auditregistration.k8s.io/v1alpha1" | undefined>;
    /**
     * List of audit configurations.
     */
    readonly items: pulumi.Input<pulumi.Input<inputs.auditregistration.v1alpha1.AuditSink>[]>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    readonly kind?: pulumi.Input<"AuditSinkList" | undefined>;
    readonly metadata?: pulumi.Input<inputs.meta.v1.ListMeta | undefined>;
}
