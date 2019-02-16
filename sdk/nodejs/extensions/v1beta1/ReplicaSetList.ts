// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../types/input";
import * as outputApi from "../../types/output";
import * as rxjs from "rxjs";
import * as operators from "rxjs/operators"

    /**
     * ReplicaSetList is a collection of ReplicaSets.
     */
    export class ReplicaSetList extends pulumi.CustomResource {
      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"extensions/v1beta1">;

      /**
       * List of ReplicaSets. More info:
       * https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
       */
      public readonly items: pulumi.Output<outputApi.extensions.v1beta1.ReplicaSet[]>;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"ReplicaSetList">;

      /**
       * Standard list metadata. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
       */
      public readonly metadata: pulumi.Output<outputApi.meta.v1.ListMeta>;

      /**
       * Get the state of an existing `ReplicaSetList` resource, as identified by `id`.
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
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ReplicaSetList {
          return new ReplicaSetList(name, undefined, { ...opts, id: id });
      }

      public getInputs(): inputApi.extensions.v1beta1.ReplicaSetList { return this.__inputs; }
      private readonly __inputs: inputApi.extensions.v1beta1.ReplicaSetList;

      public static list(): rxjs.Observable<outputApi.extensions.v1beta1.ReplicaSetList> {
        return rxjs.from(
          pulumi.runtime
            .invoke("pulumi:pulumi:readStackResourceOutputs", { stackName: pulumi.runtime.getStack() })
            .then(o => Object.keys(o.outputs).map(k => o.outputs[k]))
        ).pipe(
          operators.mergeAll(),
          operators.filter(outputApi.extensions.v1beta1.isReplicaSetList)
        );
      }

      /**
       * Create a extensions.v1beta1.ReplicaSetList resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputApi.extensions.v1beta1.ReplicaSetList, opts?: pulumi.CustomResourceOptions) {
          let inputs: pulumi.Inputs = {};
          inputs["apiVersion"] = "extensions/v1beta1";
          inputs["items"] = args && args.items || undefined;
          inputs["kind"] = "ReplicaSetList";
          inputs["metadata"] = args && args.metadata || undefined;
          super("kubernetes:extensions/v1beta1:ReplicaSetList", name, inputs, opts);
          this.__inputs = <any>args;
      }
    }
