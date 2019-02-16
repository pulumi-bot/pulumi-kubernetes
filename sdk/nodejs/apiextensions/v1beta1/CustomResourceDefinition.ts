// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../types/input";
import * as outputApi from "../../types/output";
import * as rxjs from "rxjs";
import * as operators from "rxjs/operators"

    /**
     * CustomResourceDefinition represents a resource that should be exposed on the API server.  Its
     * name MUST be in the format <.spec.name>.<.spec.group>.
     */
    export class CustomResourceDefinition extends pulumi.CustomResource {
      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"apiextensions.k8s.io/v1beta1">;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"CustomResourceDefinition">;

      
      public readonly metadata: pulumi.Output<outputApi.meta.v1.ObjectMeta>;

      /**
       * Spec describes how the user wants the resources to appear
       */
      public readonly spec: pulumi.Output<outputApi.apiextensions.v1beta1.CustomResourceDefinitionSpec>;

      /**
       * Status indicates the actual state of the CustomResourceDefinition
       */
      public readonly status: pulumi.Output<outputApi.apiextensions.v1beta1.CustomResourceDefinitionStatus>;

      /**
       * Get the state of an existing `CustomResourceDefinition` resource, as identified by `id`.
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
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CustomResourceDefinition {
          return new CustomResourceDefinition(name, undefined, { ...opts, id: id });
      }

      public getInputs(): inputApi.apiextensions.v1beta1.CustomResourceDefinition { return this.__inputs; }
      private readonly __inputs: inputApi.apiextensions.v1beta1.CustomResourceDefinition;

      public static list(): rxjs.Observable<outputApi.apiextensions.v1beta1.CustomResourceDefinition> {
        return rxjs.from(
          pulumi.runtime
            .invoke("pulumi:pulumi:readStackResourceOutputs", { stackName: pulumi.runtime.getStack() })
            .then(o => Object.keys(o.outputs).map(k => o.outputs[k]))
        ).pipe(
          operators.mergeAll(),
          operators.filter(outputApi.apiextensions.v1beta1.isCustomResourceDefinition)
        );
      }

      /**
       * Create a apiextensions.v1beta1.CustomResourceDefinition resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputApi.apiextensions.v1beta1.CustomResourceDefinition, opts?: pulumi.CustomResourceOptions) {
          let inputs: pulumi.Inputs = {};
          inputs["apiVersion"] = "apiextensions.k8s.io/v1beta1";
          inputs["kind"] = "CustomResourceDefinition";
          inputs["metadata"] = args && args.metadata || undefined;
          inputs["spec"] = args && args.spec || undefined;
          inputs["status"] = args && args.status || undefined;
          super("kubernetes:apiextensions.k8s.io/v1beta1:CustomResourceDefinition", name, inputs, opts);
          this.__inputs = <any>args;
      }
    }
