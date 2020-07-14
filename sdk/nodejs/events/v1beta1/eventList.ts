// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * EventList is a list of Event objects.
 */
export class EventList extends pulumi.CustomResource {
    /**
     * Get an existing EventList resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EventList {
        return new EventList(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:events.k8s.io/v1beta1:EventList';

    /**
     * Returns true if the given object is an instance of EventList.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventList {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventList.__pulumiType;
    }

    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<"events.k8s.io/v1beta1">;
    /**
     * Items is a list of schema objects.
     */
    public readonly items!: pulumi.Output<outputs.events.v1beta1.Event[]>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<"EventList">;
    /**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ListMeta>;

    /**
     * Create a EventList resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EventListArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventListArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as EventListArgs | undefined;
            if (!args || args.items === undefined) {
                throw new Error("Missing required property 'items'");
            }
            inputs["apiVersion"] = "events.k8s.io/v1beta1";
            inputs["items"] = args ? args.items : undefined;
            inputs["kind"] = "EventList";
            inputs["metadata"] = args ? args.metadata : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EventList.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a EventList resource.
 */
export interface EventListArgs {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    readonly apiVersion?: pulumi.Input<"events.k8s.io/v1beta1">;
    /**
     * Items is a list of schema objects.
     */
    readonly items: pulumi.Input<pulumi.Input<inputs.events.v1beta1.Event>[]>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    readonly kind?: pulumi.Input<"EventList">;
    /**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    readonly metadata?: pulumi.Input<inputs.meta.v1.ListMeta>;
}