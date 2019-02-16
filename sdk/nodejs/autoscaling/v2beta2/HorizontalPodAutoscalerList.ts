// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../types/input";
import * as outputApi from "../../types/output";
import * as rxjs from "rxjs";
import * as operators from "rxjs/operators"

    /**
     * HorizontalPodAutoscalerList is a list of horizontal pod autoscaler objects.
     */
    export class HorizontalPodAutoscalerList extends pulumi.CustomResource {
      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"autoscaling/v2beta2">;

      /**
       * items is the list of horizontal pod autoscaler objects.
       */
      public readonly items: pulumi.Output<outputApi.autoscaling.v2beta2.HorizontalPodAutoscaler[]>;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"HorizontalPodAutoscalerList">;

      /**
       * metadata is the standard list metadata.
       */
      public readonly metadata: pulumi.Output<outputApi.meta.v1.ListMeta>;

      /**
       * Get the state of an existing `HorizontalPodAutoscalerList` resource, as identified by `id`.
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
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): HorizontalPodAutoscalerList {
          return new HorizontalPodAutoscalerList(name, undefined, { ...opts, id: id });
      }

      public getInputs(): inputApi.autoscaling.v2beta2.HorizontalPodAutoscalerList { return this.__inputs; }
      private readonly __inputs: inputApi.autoscaling.v2beta2.HorizontalPodAutoscalerList;

      public static list(): rxjs.Observable<outputApi.autoscaling.v2beta2.HorizontalPodAutoscalerList> {
        return rxjs.from(
          pulumi.runtime
            .invoke("pulumi:pulumi:readStackResourceOutputs", { stackName: pulumi.runtime.getStack() })
            .then(o => Object.keys(o.outputs).map(k => o.outputs[k]))
        ).pipe(
          operators.mergeAll(),
          operators.filter(outputApi.autoscaling.v2beta2.isHorizontalPodAutoscalerList)
        );
      }

      /**
       * Create a autoscaling.v2beta2.HorizontalPodAutoscalerList resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputApi.autoscaling.v2beta2.HorizontalPodAutoscalerList, opts?: pulumi.CustomResourceOptions) {
          let inputs: pulumi.Inputs = {};
          inputs["apiVersion"] = "autoscaling/v2beta2";
          inputs["items"] = args && args.items || undefined;
          inputs["kind"] = "HorizontalPodAutoscalerList";
          inputs["metadata"] = args && args.metadata || undefined;
          super("kubernetes:autoscaling/v2beta2:HorizontalPodAutoscalerList", name, inputs, opts);
          this.__inputs = <any>args;
      }
    }
