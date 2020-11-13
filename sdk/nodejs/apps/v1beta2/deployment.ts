// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Deployment enables declarative updates for Pods and ReplicaSets.
 *
 * This resource waits until its status is ready before registering success
 * for create/update, and populating output properties from the current state of the resource.
 * The following conditions are used to determine whether the resource creation has
 * succeeded or failed:
 *
 * 1. The Deployment has begun to be updated by the Deployment controller. If the current
 *    generation of the Deployment is > 1, then this means that the current generation must
 *    be different from the generation reported by the last outputs.
 * 2. There exists a ReplicaSet whose revision is equal to the current revision of the
 *    Deployment.
 * 3. The Deployment's '.status.conditions' has a status of type 'Available' whose 'status'
 *    member is set to 'True'.
 * 4. If the Deployment has generation > 1, then '.status.conditions' has a status of type
 *    'Progressing', whose 'status' member is set to 'True', and whose 'reason' is
 *    'NewReplicaSetAvailable'. For generation <= 1, this status field does not exist,
 *    because it doesn't do a rollout (i.e., it simply creates the Deployment and
 *    corresponding ReplicaSet), and therefore there is no rollout to mark as 'Progressing'.
 *
 * If the Deployment has not reached a Ready state after 10 minutes, it will
 * time out and mark the resource update as Failed. You can override the default timeout value
 * by setting the 'customTimeouts' option on the resource.
 *
 * @deprecated apps/v1beta2/Deployment is deprecated by apps/v1/Deployment and not supported by Kubernetes v1.16+ clusters.
 */
export class Deployment extends pulumi.CustomResource {
    /**
     * Get an existing Deployment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Deployment {
        return new Deployment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:apps/v1beta2:Deployment';

    /**
     * Returns true if the given object is an instance of Deployment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Deployment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Deployment.__pulumiType;
    }

    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<"apps/v1beta2">;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<"Deployment">;
    /**
     * Standard object metadata.
     */
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ObjectMeta>;
    /**
     * Specification of the desired behavior of the Deployment.
     */
    public readonly spec!: pulumi.Output<outputs.apps.v1beta2.DeploymentSpec>;
    /**
     * Most recently observed status of the Deployment.
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.apps.v1beta2.DeploymentStatus>;

    /**
     * Create a Deployment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated apps/v1beta2/Deployment is deprecated by apps/v1/Deployment and not supported by Kubernetes v1.16+ clusters. */
    constructor(name: string, args?: DeploymentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            inputs["apiVersion"] = "apps/v1beta2";
            inputs["kind"] = "Deployment";
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
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "kubernetes:apps/v1:Deployment" }, { type: "kubernetes:apps/v1beta1:Deployment" }, { type: "kubernetes:extensions/v1beta1:Deployment" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Deployment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Deployment resource.
 */
export interface DeploymentArgs {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    readonly apiVersion?: pulumi.Input<"apps/v1beta2">;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    readonly kind?: pulumi.Input<"Deployment">;
    /**
     * Standard object metadata.
     */
    readonly metadata?: pulumi.Input<inputs.meta.v1.ObjectMeta>;
    /**
     * Specification of the desired behavior of the Deployment.
     */
    readonly spec?: pulumi.Input<inputs.apps.v1beta2.DeploymentSpec>;
}
