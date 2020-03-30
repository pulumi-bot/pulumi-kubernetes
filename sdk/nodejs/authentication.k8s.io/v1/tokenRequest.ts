// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * TokenRequest requests a token for a given service account.
 */
export class TokenRequest extends pulumi.CustomResource {
    /**
     * Get an existing TokenRequest resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TokenRequest {
        return new TokenRequest(name, undefined{ ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:authentication.k8s.io/v1:TokenRequest';

    /**
     * Returns true if the given object is an instance of TokenRequest.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TokenRequest {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TokenRequest.__pulumiType;
    }

    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<string | undefined>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ObjectMeta | undefined>;
    public readonly spec!: pulumi.Output<outputs.authentication.k8s.io.v1.TokenRequestSpec | undefined>;
    public /*out*/ readonly status!: pulumi.Output<outputs.authentication.k8s.io.v1.TokenRequestStatus | undefined>;

    /**
     * Create a TokenRequest resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TokenRequestArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TokenRequestArgs | TokenRequestState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TokenRequestState | undefined;
            inputs["apiVersion"] = state ? state.apiVersion : undefined;
            inputs["kind"] = state ? state.kind : undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["spec"] = state ? state.spec : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as TokenRequestArgs | undefined;
            if (!args || args.spec === undefined) {
                throw new Error("Missing required property 'spec'");
            }
            inputs["apiVersion"] = (args ? args.apiVersion : undefined) || "authentication.k8s.io/v1";
            inputs["kind"] = (args ? args.kind : undefined) || "TokenRequest";
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["spec"] = args ? args.spec : undefined;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TokenRequest.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a TokenRequest resource.
 */
export interface TokenRequestArgs {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    readonly apiVersion?: pulumi.Input<string>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    readonly kind?: pulumi.Input<string>;
    readonly metadata?: pulumi.Input<inputs.meta.v1.ObjectMeta>;
    readonly spec: pulumi.Input<inputs.authentication.k8s.io.v1.TokenRequestSpec>;
}