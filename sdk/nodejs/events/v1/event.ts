// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
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
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Event {
        return new Event(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:events.k8s.io/v1:Event';

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
     * action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field can have at most 128 characters.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<"events.k8s.io/v1">;
    /**
     * deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.
     */
    public readonly deprecatedCount!: pulumi.Output<number>;
    /**
     * deprecatedFirstTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
     */
    public readonly deprecatedFirstTimestamp!: pulumi.Output<string>;
    /**
     * deprecatedLastTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
     */
    public readonly deprecatedLastTimestamp!: pulumi.Output<string>;
    /**
     * deprecatedSource is the deprecated field assuring backward compatibility with core.v1 Event type.
     */
    public readonly deprecatedSource!: pulumi.Output<outputs.core.v1.EventSource>;
    /**
     * eventTime is the time when this Event was first observed. It is required.
     */
    public readonly eventTime!: pulumi.Output<string>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<"Event">;
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ObjectMeta>;
    /**
     * note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
     */
    public readonly note!: pulumi.Output<string>;
    /**
     * reason is why the action was taken. It is human-readable. This field can have at most 128 characters.
     */
    public readonly reason!: pulumi.Output<string>;
    /**
     * regarding contains the object this Event is about. In most cases it's an Object reporting controller implements, e.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
     */
    public readonly regarding!: pulumi.Output<outputs.core.v1.ObjectReference>;
    /**
     * related is the optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
     */
    public readonly related!: pulumi.Output<outputs.core.v1.ObjectReference>;
    /**
     * reportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`. This field cannot be empty for new Events.
     */
    public readonly reportingController!: pulumi.Output<string>;
    /**
     * reportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`. This field cannot be empty for new Events and it can have at most 128 characters.
     */
    public readonly reportingInstance!: pulumi.Output<string>;
    /**
     * series is data about the Event series this event represents or nil if it's a singleton Event.
     */
    public readonly series!: pulumi.Output<outputs.events.v1.EventSeries>;
    /**
     * type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Event resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EventArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.eventTime === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'eventTime'");
            }
            inputs["action"] = args ? args.action : undefined;
            inputs["apiVersion"] = "events.k8s.io/v1";
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
        } else {
            inputs["action"] = undefined /*out*/;
            inputs["apiVersion"] = undefined /*out*/;
            inputs["deprecatedCount"] = undefined /*out*/;
            inputs["deprecatedFirstTimestamp"] = undefined /*out*/;
            inputs["deprecatedLastTimestamp"] = undefined /*out*/;
            inputs["deprecatedSource"] = undefined /*out*/;
            inputs["eventTime"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["metadata"] = undefined /*out*/;
            inputs["note"] = undefined /*out*/;
            inputs["reason"] = undefined /*out*/;
            inputs["regarding"] = undefined /*out*/;
            inputs["related"] = undefined /*out*/;
            inputs["reportingController"] = undefined /*out*/;
            inputs["reportingInstance"] = undefined /*out*/;
            inputs["series"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "kubernetes:core/v1:Event" }, { type: "kubernetes:events.k8s.io/v1beta1:Event" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Event.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Event resource.
 */
export interface EventArgs {
    /**
     * action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field can have at most 128 characters.
     */
    readonly action?: pulumi.Input<string>;
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    readonly apiVersion?: pulumi.Input<"events.k8s.io/v1">;
    /**
     * deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.
     */
    readonly deprecatedCount?: pulumi.Input<number>;
    /**
     * deprecatedFirstTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
     */
    readonly deprecatedFirstTimestamp?: pulumi.Input<string>;
    /**
     * deprecatedLastTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
     */
    readonly deprecatedLastTimestamp?: pulumi.Input<string>;
    /**
     * deprecatedSource is the deprecated field assuring backward compatibility with core.v1 Event type.
     */
    readonly deprecatedSource?: pulumi.Input<inputs.core.v1.EventSource>;
    /**
     * eventTime is the time when this Event was first observed. It is required.
     */
    readonly eventTime: pulumi.Input<string>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    readonly kind?: pulumi.Input<"Event">;
    readonly metadata?: pulumi.Input<inputs.meta.v1.ObjectMeta>;
    /**
     * note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
     */
    readonly note?: pulumi.Input<string>;
    /**
     * reason is why the action was taken. It is human-readable. This field can have at most 128 characters.
     */
    readonly reason?: pulumi.Input<string>;
    /**
     * regarding contains the object this Event is about. In most cases it's an Object reporting controller implements, e.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
     */
    readonly regarding?: pulumi.Input<inputs.core.v1.ObjectReference>;
    /**
     * related is the optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
     */
    readonly related?: pulumi.Input<inputs.core.v1.ObjectReference>;
    /**
     * reportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`. This field cannot be empty for new Events.
     */
    readonly reportingController?: pulumi.Input<string>;
    /**
     * reportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`. This field cannot be empty for new Events and it can have at most 128 characters.
     */
    readonly reportingInstance?: pulumi.Input<string>;
    /**
     * series is data about the Event series this event represents or nil if it's a singleton Event.
     */
    readonly series?: pulumi.Input<inputs.events.v1.EventSeries>;
    /**
     * type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable.
     */
    readonly type?: pulumi.Input<string>;
}
