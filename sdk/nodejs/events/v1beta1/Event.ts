// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system.
 */
export class Event extends pulumi.CustomResource {
    /**
     * Get an existing Event resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Event {
        return new Event(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:events.k8s.io/v1beta1:Event';

    /**
     * Returns true if the given object is an instance of Event.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Event {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Event.__pulumiType;
    }

    /**
     * What action was taken/failed regarding to the regarding object.
     */
    public readonly action!: pulumi.Output<string | undefined>;
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<string | undefined>;
    /**
     * Deprecated field assuring backward compatibility with core.v1 Event type
     */
    public readonly deprecatedCount!: pulumi.Output<number | undefined>;
    /**
     * Deprecated field assuring backward compatibility with core.v1 Event type
     */
    public readonly deprecatedFirstTimestamp!: pulumi.Output<string | undefined>;
    /**
     * Deprecated field assuring backward compatibility with core.v1 Event type
     */
    public readonly deprecatedLastTimestamp!: pulumi.Output<string | undefined>;
    /**
     * Deprecated field assuring backward compatibility with core.v1 Event type
     */
    public readonly deprecatedSource!: pulumi.Output<outputs.core.v1.EventSource | undefined>;
    /**
     * Required. Time when this Event was first observed.
     */
    public readonly eventTime!: pulumi.Output<string | undefined>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ObjectMeta | undefined>;
    /**
     * Optional. A human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
     */
    public readonly note!: pulumi.Output<string | undefined>;
    /**
     * Why the action was taken.
     */
    public readonly reason!: pulumi.Output<string | undefined>;
    /**
     * The object this Event is about. In most cases it's an Object reporting controller implements. E.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
     */
    public readonly regarding!: pulumi.Output<outputs.core.v1.ObjectReference | undefined>;
    /**
     * Optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
     */
    public readonly related!: pulumi.Output<outputs.core.v1.ObjectReference | undefined>;
    /**
     * Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
     */
    public readonly reportingController!: pulumi.Output<string | undefined>;
    /**
     * ID of the controller instance, e.g. `kubelet-xyzf`.
     */
    public readonly reportingInstance!: pulumi.Output<string | undefined>;
    /**
     * Data about the Event series this event represents or nil if it's a singleton Event.
     */
    public readonly series!: pulumi.Output<outputs.events.v1beta1.EventSeries | undefined>;
    /**
     * Type of this event (Normal, Warning), new types could be added in the future.
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a Event resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.eventTime === undefined) {
                throw new Error("Missing required property 'eventTime'");
            }
        inputs["action"] = args ? args.action : undefined;
        inputs["apiVersion"] = "events.k8s.io/v1beta1";
        inputs["deprecatedCount"] = args ? args.deprecatedCount : undefined;
        inputs["deprecatedFirstTimestamp"] = args ? args.deprecatedFirstTimestamp : undefined;
        inputs["deprecatedLastTimestamp"] = args ? args.deprecatedLastTimestamp : undefined;
        inputs["deprecatedSource"] = args ? args.deprecatedSource : undefined;
        inputs["eventTime"] = args ? args.eventTime : undefined;
        inputs["kind"] = "Event";
        inputs["metadata"] = args ? args.metadata : undefined;
        inputs["note"] = args ? args.note : undefined;
        inputs["reason"] = args ? args.reason : undefined;
        inputs["regarding"] = args ? args.regarding : undefined;
        inputs["related"] = args ? args.related : undefined;
        inputs["reportingController"] = args ? args.reportingController : undefined;
        inputs["reportingInstance"] = args ? args.reportingInstance : undefined;
        inputs["series"] = args ? args.series : undefined;
        inputs["type"] = args ? args.type : undefined;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "kubernetes:core/v1:Event" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Event.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Event resource.
 */
export interface EventArgs {
    /**
     * What action was taken/failed regarding to the regarding object.
     */
    readonly action?: pulumi.Input<string>;
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    readonly apiVersion?: pulumi.Input<string>;
    /**
     * Deprecated field assuring backward compatibility with core.v1 Event type
     */
    readonly deprecatedCount?: pulumi.Input<number>;
    /**
     * Deprecated field assuring backward compatibility with core.v1 Event type
     */
    readonly deprecatedFirstTimestamp?: pulumi.Input<string>;
    /**
     * Deprecated field assuring backward compatibility with core.v1 Event type
     */
    readonly deprecatedLastTimestamp?: pulumi.Input<string>;
    /**
     * Deprecated field assuring backward compatibility with core.v1 Event type
     */
    readonly deprecatedSource?: pulumi.Input<inputs.core.v1.EventSource>;
    /**
     * Required. Time when this Event was first observed.
     */
    readonly eventTime: pulumi.Input<string>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    readonly kind?: pulumi.Input<string>;
    readonly metadata?: pulumi.Input<inputs.meta.v1.ObjectMeta>;
    /**
     * Optional. A human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
     */
    readonly note?: pulumi.Input<string>;
    /**
     * Why the action was taken.
     */
    readonly reason?: pulumi.Input<string>;
    /**
     * The object this Event is about. In most cases it's an Object reporting controller implements. E.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
     */
    readonly regarding?: pulumi.Input<inputs.core.v1.ObjectReference>;
    /**
     * Optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
     */
    readonly related?: pulumi.Input<inputs.core.v1.ObjectReference>;
    /**
     * Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
     */
    readonly reportingController?: pulumi.Input<string>;
    /**
     * ID of the controller instance, e.g. `kubelet-xyzf`.
     */
    readonly reportingInstance?: pulumi.Input<string>;
    /**
     * Data about the Event series this event represents or nil if it's a singleton Event.
     */
    readonly series?: pulumi.Input<inputs.events.v1beta1.EventSeries>;
    /**
     * Type of this event (Normal, Warning), new types could be added in the future.
     */
    readonly type?: pulumi.Input<string>;
}
