// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./runtimeClass";
export * from "./runtimeClassList";

// Import resources to register:
import { RuntimeClass } from "./runtimeClass";
import { RuntimeClassList } from "./runtimeClassList";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "kubernetes:node.k8s.io/v1alpha1:RuntimeClass":
                return new RuntimeClass(name, <any>undefined, { urn })
            case "kubernetes:node.k8s.io/v1alpha1:RuntimeClassList":
                return new RuntimeClassList(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("kubernetes", "node.k8s.io/v1alpha1", _module)
