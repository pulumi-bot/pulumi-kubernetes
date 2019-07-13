// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { core } from "../..";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import { getVersion } from "../../version";

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

      
      public readonly metadata: pulumi.Output<outputs.meta.v1.ObjectMeta>;

      /**
       * Spec describes how the user wants the resources to appear
       */
      public readonly spec: pulumi.Output<outputs.apiextensions.v1beta1.CustomResourceDefinitionSpec>;

      /**
       * Status indicates the actual state of the CustomResourceDefinition
       */
      public readonly status: pulumi.Output<outputs.apiextensions.v1beta1.CustomResourceDefinitionStatus>;

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

      /** @internal */
      private static readonly __pulumiType = "kubernetes:apiextensions.k8s.io/v1beta1:CustomResourceDefinition";

      /**
       * Returns true if the given object is an instance of CustomResourceDefinition.  This is designed to work even
       * when multiple copies of the Pulumi SDK have been loaded into the same process.
       */
      public static isInstance(obj: any): obj is CustomResourceDefinition {
          if (obj === undefined || obj === null) {
              return false;
          }

          return obj["__pulumiType"] === CustomResourceDefinition.__pulumiType;
      }

      /**
       * Create a apiextensions.v1beta1.CustomResourceDefinition resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputs.apiextensions.v1beta1.CustomResourceDefinition, opts?: pulumi.CustomResourceOptions) {
          const props: pulumi.Inputs = {};
          props["spec"] = args && args.spec || undefined;

          props["apiVersion"] = "apiextensions.k8s.io/v1beta1";
          props["kind"] = "CustomResourceDefinition";
          props["metadata"] = args && args.metadata || undefined;

          props["status"] = undefined;

          if (!opts) {
              opts = {};
          }

          if (!opts.version) {
              opts.version = getVersion();
          }
          super(CustomResourceDefinition.__pulumiType, name, props, opts);
      }
    }
