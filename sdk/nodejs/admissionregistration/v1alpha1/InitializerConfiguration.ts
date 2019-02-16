// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../types/input";
import * as outputApi from "../../types/output";
import * as rxjs from "rxjs";
import * as operators from "rxjs/operators"

    /**
     * InitializerConfiguration describes the configuration of initializers.
     */
    export class InitializerConfiguration extends pulumi.CustomResource {
      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"admissionregistration.k8s.io/v1alpha1">;

      /**
       * Initializers is a list of resources and their default initializers Order-sensitive. When
       * merging multiple InitializerConfigurations, we sort the initializers from different
       * InitializerConfigurations by the name of the InitializerConfigurations; the order of the
       * initializers from the same InitializerConfiguration is preserved.
       */
      public readonly initializers: pulumi.Output<outputApi.admissionregistration.v1alpha1.Initializer[]>;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"InitializerConfiguration">;

      /**
       * Standard object metadata; More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata.
       */
      public readonly metadata: pulumi.Output<outputApi.meta.v1.ObjectMeta>;

      /**
       * Get the state of an existing `InitializerConfiguration` resource, as identified by `id`.
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
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): InitializerConfiguration {
          return new InitializerConfiguration(name, undefined, { ...opts, id: id });
      }

      public getInputs(): inputApi.admissionregistration.v1alpha1.InitializerConfiguration { return this.__inputs; }
      private readonly __inputs: inputApi.admissionregistration.v1alpha1.InitializerConfiguration;

      public static list(): rxjs.Observable<outputApi.admissionregistration.v1alpha1.InitializerConfiguration> {
        return rxjs.from(
          pulumi.runtime
            .invoke("pulumi:pulumi:readStackResourceOutputs", { stackName: pulumi.runtime.getStack() })
            .then(o => Object.keys(o.outputs).map(k => o.outputs[k]))
        ).pipe(
          operators.mergeAll(),
          operators.filter(outputApi.admissionregistration.v1alpha1.isInitializerConfiguration)
        );
      }

      /**
       * Create a admissionregistration.v1alpha1.InitializerConfiguration resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputApi.admissionregistration.v1alpha1.InitializerConfiguration, opts?: pulumi.CustomResourceOptions) {
          let inputs: pulumi.Inputs = {};
          inputs["apiVersion"] = "admissionregistration.k8s.io/v1alpha1";
          inputs["initializers"] = args && args.initializers || undefined;
          inputs["kind"] = "InitializerConfiguration";
          inputs["metadata"] = args && args.metadata || undefined;
          super("kubernetes:admissionregistration.k8s.io/v1alpha1:InitializerConfiguration", name, inputs, opts);
          this.__inputs = <any>args;
      }
    }
