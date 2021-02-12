// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * EndpointSlice represents a subset of the endpoints that implement a service. For a given service there may be multiple EndpointSlice objects, selected by labels, which must be joined to produce the full set of endpoints.
 */
export class EndpointSlice extends pulumi.CustomResource {
    /**
     * Get an existing EndpointSlice resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EndpointSlice {
        return new EndpointSlice(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:discovery.k8s.io/v1beta1:EndpointSlice';

    /**
     * Returns true if the given object is an instance of EndpointSlice.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EndpointSlice {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EndpointSlice.__pulumiType;
    }

    /**
     * addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.
     */
    public readonly addressType!: pulumi.Output<string>;
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<"discovery.k8s.io/v1beta1">;
    /**
     * endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.
     */
    public readonly endpoints!: pulumi.Output<outputs.discovery.v1beta1.Endpoint[]>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<"EndpointSlice">;
    /**
     * Standard object's metadata.
     */
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ObjectMeta>;
    /**
     * ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates "all ports". Each slice may include a maximum of 100 ports.
     */
    public readonly ports!: pulumi.Output<outputs.discovery.v1beta1.EndpointPort[]>;

    /**
     * Create a EndpointSlice resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EndpointSliceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.addressType === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'addressType'");
            }
            if ((!args || args.endpoints === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'endpoints'");
            }
            inputs["addressType"] = args ? args.addressType : undefined;
            inputs["apiVersion"] = "discovery.k8s.io/v1beta1";
            inputs["endpoints"] = args ? args.endpoints : undefined;
            inputs["kind"] = "EndpointSlice";
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["ports"] = args ? args.ports : undefined;
        } else {
            inputs["addressType"] = undefined /*out*/;
            inputs["apiVersion"] = undefined /*out*/;
            inputs["endpoints"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["metadata"] = undefined /*out*/;
            inputs["ports"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EndpointSlice.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a EndpointSlice resource.
 */
export interface EndpointSliceArgs {
    /**
     * addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.
     */
    readonly addressType: pulumi.Input<string>;
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    readonly apiVersion?: pulumi.Input<"discovery.k8s.io/v1beta1" | undefined>;
    /**
     * endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.
     */
    readonly endpoints: pulumi.Input<pulumi.Input<inputs.discovery.v1beta1.Endpoint>[]>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    readonly kind?: pulumi.Input<"EndpointSlice" | undefined>;
    /**
     * Standard object's metadata.
     */
    readonly metadata?: pulumi.Input<inputs.meta.v1.ObjectMeta | undefined>;
    /**
     * ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates "all ports". Each slice may include a maximum of 100 ports.
     */
    readonly ports?: pulumi.Input<pulumi.Input<inputs.discovery.v1beta1.EndpointPort>[] | undefined>;
}
