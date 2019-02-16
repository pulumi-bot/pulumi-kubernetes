// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../types/input";
import * as outputApi from "../../types/output";
import * as rxjs from "rxjs";
import * as operators from "rxjs/operators"

    /**
     * PodSecurityPolicy governs the ability to make requests that affect the Security Context that
     * will be applied to a pod and container.
     */
    export class PodSecurityPolicy extends pulumi.CustomResource {
      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"policy/v1beta1">;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"PodSecurityPolicy">;

      /**
       * Standard object's metadata. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
       */
      public readonly metadata: pulumi.Output<outputApi.meta.v1.ObjectMeta>;

      /**
       * spec defines the policy enforced.
       */
      public readonly spec: pulumi.Output<outputApi.policy.v1beta1.PodSecurityPolicySpec>;

      /**
       * Get the state of an existing `PodSecurityPolicy` resource, as identified by `id`.
       * Typically this ID  is of the form <namespace>/<name>; if <namespace> is omitted, then (per
       * Kubernetes convention) the ID becomes default/<name>.
       *
       * Pulumi will keep track of this resource using `name` as the Pulumi ID.
       *
       * @param name _Unique_ name used to register this resource with Pulumi.
       * @param id An ID for the Kubernetes resource to retrieve. Takes the form
       *  <namespace>/<name> or <name>.
       * @param opts Uniquely specifies a CustomResource to select.
       */
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PodSecurityPolicy {
          return new PodSecurityPolicy(name, undefined, { ...opts, id: id });
      }

      public getInputs(): inputApi.policy.v1beta1.PodSecurityPolicy { return this.__inputs; }
      private readonly __inputs: inputApi.policy.v1beta1.PodSecurityPolicy;

      public static list(): rxjs.Observable<outputApi.policy.v1beta1.PodSecurityPolicy> {
        return rxjs.from(
          pulumi.runtime
            .invoke("pulumi:pulumi:readStackResourceOutputs", { stackName: pulumi.runtime.getStack() })
            .then(o => Object.keys(o.outputs).map(k => o.outputs[k]))
        ).pipe(
          operators.mergeAll(),
          operators.filter(outputApi.policy.v1beta1.isPodSecurityPolicy)
        );
      }

      /**
       * Create a policy.v1beta1.PodSecurityPolicy resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputApi.policy.v1beta1.PodSecurityPolicy, opts?: pulumi.CustomResourceOptions) {
          let inputs: pulumi.Inputs = {};
          inputs["apiVersion"] = "policy/v1beta1";
          inputs["kind"] = "PodSecurityPolicy";
          inputs["metadata"] = args && args.metadata || undefined;
          inputs["spec"] = args && args.spec || undefined;
          super("kubernetes:policy/v1beta1:PodSecurityPolicy", name, inputs, opts);
          this.__inputs = <any>args;
      }
    }
