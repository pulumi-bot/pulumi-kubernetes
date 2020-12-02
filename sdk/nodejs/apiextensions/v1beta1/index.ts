// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./customResourceDefinition";
export * from "./customResourceDefinitionList";

// Import resources to register:
import { CustomResourceDefinition } from "./customResourceDefinition";
import { CustomResourceDefinitionList } from "./customResourceDefinitionList";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "kubernetes:apiextensions.k8s.io/v1beta1:CustomResourceDefinition":
                return new CustomResourceDefinition(name, <any>undefined, { urn })
            case "kubernetes:apiextensions.k8s.io/v1beta1:CustomResourceDefinitionList":
                return new CustomResourceDefinitionList(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("kubernetes", "apiextensions.k8s.io/v1beta1", _module)
