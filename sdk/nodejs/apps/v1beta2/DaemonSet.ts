// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../types/input";
import * as outputApi from "../../types/output";
import * as rxjs from "rxjs";
import * as operators from "rxjs/operators"

    /**
     * DEPRECATED - This group version of DaemonSet is deprecated by apps/v1/DaemonSet. See the
     * release notes for more information. DaemonSet represents the configuration of a daemon set.
     */
    export class DaemonSet extends pulumi.CustomResource {
      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"apps/v1beta2">;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"DaemonSet">;

      /**
       * Standard object's metadata. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
       */
      public readonly metadata: pulumi.Output<outputApi.meta.v1.ObjectMeta>;

      /**
       * The desired behavior of this daemon set. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
       */
      public readonly spec: pulumi.Output<outputApi.apps.v1beta2.DaemonSetSpec>;

      /**
       * The current status of this daemon set. This data may be out of date by some window of time.
       * Populated by the system. Read-only. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
       */
      public readonly status: pulumi.Output<outputApi.apps.v1beta2.DaemonSetStatus>;

      /**
       * Get the state of an existing `DaemonSet` resource, as identified by `id`.
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
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DaemonSet {
          return new DaemonSet(name, undefined, { ...opts, id: id });
      }

      public getInputs(): inputApi.apps.v1beta2.DaemonSet { return this.__inputs; }
      private readonly __inputs: inputApi.apps.v1beta2.DaemonSet;

      public static list(): rxjs.Observable<outputApi.apps.v1beta2.DaemonSet> {
        return rxjs.from(
          pulumi.runtime
            .invoke("pulumi:pulumi:readStackResourceOutputs", { stackName: pulumi.runtime.getStack() })
            .then(o => Object.keys(o.outputs).map(k => o.outputs[k]))
        ).pipe(
          operators.mergeAll(),
          operators.filter(outputApi.apps.v1beta2.isDaemonSet)
        );
      }

      /**
       * Create a apps.v1beta2.DaemonSet resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputApi.apps.v1beta2.DaemonSet, opts?: pulumi.CustomResourceOptions) {
          let inputs: pulumi.Inputs = {};
          inputs["apiVersion"] = "apps/v1beta2";
          inputs["kind"] = "DaemonSet";
          inputs["metadata"] = args && args.metadata || undefined;
          inputs["spec"] = args && args.spec || undefined;
          inputs["status"] = args && args.status || undefined;
          super("kubernetes:apps/v1beta2:DaemonSet", name, inputs, opts);
          this.__inputs = <any>args;
      }
    }
