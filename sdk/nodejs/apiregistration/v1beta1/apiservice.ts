// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * APIService represents a server for a particular GroupVersion. Name must be "version.group".
 */
export class APIService extends pulumi.CustomResource {
    /**
     * Get an existing APIService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): APIService {
        return new APIService(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:apiregistration.k8s.io/v1beta1:APIService';

    /**
     * Returns true if the given object is an instance of APIService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is APIService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === APIService.__pulumiType;
    }

    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<"apiregistration.k8s.io/v1beta1">;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<"APIService">;
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ObjectMeta>;
    /**
     * Spec contains information for locating and communicating with a server
     */
    public readonly spec!: pulumi.Output<outputs.apiregistration.v1beta1.APIServiceSpec>;
    /**
     * Status contains derived information about an API server
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.apiregistration.v1beta1.APIServiceStatus>;

    /**
     * Create a APIService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: APIServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        inputs["apiVersion"] = "apiregistration.k8s.io/v1beta1";
        inputs["kind"] = "APIService";
        inputs["metadata"] = args ? args.metadata : undefined;
        inputs["spec"] = args ? args.spec : undefined;
        inputs["status"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "kubernetes:apiregistration.k8s.io/v1:APIService" }, { type: "kubernetes:apiregistration/v1:APIService" }, { type: "kubernetes:apiregistration/v1beta1:APIService" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(APIService.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a APIService resource.
 */
export interface APIServiceArgs {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    readonly apiVersion?: pulumi.Input<"apiregistration.k8s.io/v1beta1">;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    readonly kind?: pulumi.Input<"APIService">;
    readonly metadata?: pulumi.Input<inputs.meta.v1.ObjectMeta>;
    /**
     * Spec contains information for locating and communicating with a server
     */
    readonly spec?: pulumi.Input<inputs.apiregistration.v1beta1.APIServiceSpec>;
}