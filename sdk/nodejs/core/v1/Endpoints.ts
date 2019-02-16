// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../types/input";
import * as outputApi from "../../types/output";
import * as rxjs from "rxjs";
import * as operators from "rxjs/operators"

    /**
     * Endpoints is a collection of endpoints that implement the actual service. Example:
     *   Name: "mysvc",
     *   Subsets: [
     *     {
     *       Addresses: [{"ip": "10.10.1.1"}, {"ip": "10.10.2.2"}],
     *       Ports: [{"name": "a", "port": 8675}, {"name": "b", "port": 309}]
     *     },
     *     {
     *       Addresses: [{"ip": "10.10.3.3"}],
     *       Ports: [{"name": "a", "port": 93}, {"name": "b", "port": 76}]
     *     },
     *  ]
     */
    export class Endpoints extends pulumi.CustomResource {
      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"v1">;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"Endpoints">;

      /**
       * Standard object's metadata. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
       */
      public readonly metadata: pulumi.Output<outputApi.meta.v1.ObjectMeta>;

      /**
       * The set of all endpoints is the union of all subsets. Addresses are placed into subsets
       * according to the IPs they share. A single address with multiple ports, some of which are
       * ready and some of which are not (because they come from different containers) will result
       * in the address being displayed in different subsets for the different ports. No address
       * will appear in both Addresses and NotReadyAddresses in the same subset. Sets of addresses
       * and ports that comprise a service.
       */
      public readonly subsets: pulumi.Output<outputApi.core.v1.EndpointSubset[]>;

      /**
       * Get the state of an existing `Endpoints` resource, as identified by `id`.
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
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Endpoints {
          return new Endpoints(name, undefined, { ...opts, id: id });
      }

      public getInputs(): inputApi.core.v1.Endpoints { return this.__inputs; }
      private readonly __inputs: inputApi.core.v1.Endpoints;

      public static list(): rxjs.Observable<outputApi.core.v1.Endpoints> {
        return rxjs.from(
          pulumi.runtime
            .invoke("pulumi:pulumi:readStackResourceOutputs", { stackName: pulumi.runtime.getStack() })
            .then(o => Object.keys(o.outputs).map(k => o.outputs[k]))
        ).pipe(
          operators.mergeAll(),
          operators.filter(outputApi.core.v1.isEndpoints)
        );
      }

      /**
       * Create a core.v1.Endpoints resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputApi.core.v1.Endpoints, opts?: pulumi.CustomResourceOptions) {
          let inputs: pulumi.Inputs = {};
          inputs["apiVersion"] = "v1";
          inputs["kind"] = "Endpoints";
          inputs["metadata"] = args && args.metadata || undefined;
          inputs["subsets"] = args && args.subsets || undefined;
          super("kubernetes:core/v1:Endpoints", name, inputs, opts);
          this.__inputs = <any>args;
      }
    }