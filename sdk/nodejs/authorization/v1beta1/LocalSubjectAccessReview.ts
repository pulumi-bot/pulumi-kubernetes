// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../types/input";
import * as outputApi from "../../types/output";
import * as rxjs from "rxjs";
import * as operators from "rxjs/operators"

    /**
     * LocalSubjectAccessReview checks whether or not a user or group can perform an action in a
     * given namespace. Having a namespace scoped resource makes it much easier to grant namespace
     * scoped policy that includes permissions checking.
     */
    export class LocalSubjectAccessReview extends pulumi.CustomResource {
      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"authorization.k8s.io/v1beta1">;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"LocalSubjectAccessReview">;

      
      public readonly metadata: pulumi.Output<outputApi.meta.v1.ObjectMeta>;

      /**
       * Spec holds information about the request being evaluated.  spec.namespace must be equal to
       * the namespace you made the request against.  If empty, it is defaulted.
       */
      public readonly spec: pulumi.Output<outputApi.authorization.v1beta1.SubjectAccessReviewSpec>;

      /**
       * Status is filled in by the server and indicates whether the request is allowed or not
       */
      public readonly status: pulumi.Output<outputApi.authorization.v1beta1.SubjectAccessReviewStatus>;

      /**
       * Get the state of an existing `LocalSubjectAccessReview` resource, as identified by `id`.
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
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LocalSubjectAccessReview {
          return new LocalSubjectAccessReview(name, undefined, { ...opts, id: id });
      }

      public getInputs(): inputApi.authorization.v1beta1.LocalSubjectAccessReview { return this.__inputs; }
      private readonly __inputs: inputApi.authorization.v1beta1.LocalSubjectAccessReview;

      public static list(): rxjs.Observable<outputApi.authorization.v1beta1.LocalSubjectAccessReview> {
        return rxjs.from(
          pulumi.runtime
            .invoke("pulumi:pulumi:readStackResourceOutputs", { stackName: pulumi.runtime.getStack() })
            .then(o => Object.keys(o.outputs).map(k => o.outputs[k]))
        ).pipe(
          operators.mergeAll(),
          operators.filter(outputApi.authorization.v1beta1.isLocalSubjectAccessReview)
        );
      }

      /**
       * Create a authorization.v1beta1.LocalSubjectAccessReview resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputApi.authorization.v1beta1.LocalSubjectAccessReview, opts?: pulumi.CustomResourceOptions) {
          let inputs: pulumi.Inputs = {};
          inputs["apiVersion"] = "authorization.k8s.io/v1beta1";
          inputs["kind"] = "LocalSubjectAccessReview";
          inputs["metadata"] = args && args.metadata || undefined;
          inputs["spec"] = args && args.spec || undefined;
          inputs["status"] = args && args.status || undefined;
          super("kubernetes:authorization.k8s.io/v1beta1:LocalSubjectAccessReview", name, inputs, opts);
          this.__inputs = <any>args;
      }
    }
