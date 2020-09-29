// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enum";
import * as utilities from "../../utilities";

/**
 * list of horizontal pod autoscaler objects.
 */
export class HorizontalPodAutoscalerList extends pulumi.CustomResource {
    /**
     * Get an existing HorizontalPodAutoscalerList resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): HorizontalPodAutoscalerList {
        return new HorizontalPodAutoscalerList(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:autoscaling/v1:HorizontalPodAutoscalerList';

    /**
     * Returns true if the given object is an instance of HorizontalPodAutoscalerList.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HorizontalPodAutoscalerList {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HorizontalPodAutoscalerList.__pulumiType;
    }

    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<"autoscaling/v1">;
    /**
     * list of horizontal pod autoscaler objects.
     */
    public readonly items!: pulumi.Output<outputs.autoscaling.v1.HorizontalPodAutoscaler[]>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<"HorizontalPodAutoscalerList">;
    /**
     * Standard list metadata.
     */
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ListMeta>;

    /**
     * Create a HorizontalPodAutoscalerList resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: HorizontalPodAutoscalerListArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.items === undefined) {
                throw new Error("Missing required property 'items'");
            }
            inputs["apiVersion"] = "autoscaling/v1";
            inputs["items"] = args ? args.items : undefined;
            inputs["kind"] = "HorizontalPodAutoscalerList";
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
        super(HorizontalPodAutoscalerList.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a HorizontalPodAutoscalerList resource.
 */
export interface HorizontalPodAutoscalerListArgs {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    readonly apiVersion?: pulumi.Input<"autoscaling/v1">;
    /**
     * list of horizontal pod autoscaler objects.
     */
    readonly items: pulumi.Input<pulumi.Input<inputs.autoscaling.v1.HorizontalPodAutoscaler>[]>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    readonly kind?: pulumi.Input<"HorizontalPodAutoscalerList">;
    /**
     * Standard list metadata.
     */
    readonly metadata?: pulumi.Input<inputs.meta.v1.ListMeta>;
}
