// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// CustomResourceDefinitionSpec describes how a user wants their resource to appear
type CustomResourceDefinitionSpec struct {
	// conversion defines conversion settings for the CRD.
	Conversion *CustomResourceConversion `pulumi:"conversion"`

	// group is the API group of the defined custom resource. The custom resources are served under
	// `/apis/<group>/...`. Must match the name of the CustomResourceDefinition (in the form
	// `<names.plural>.<group>`).
	Group string `pulumi:"group"`

	// names specify the resource and kind names for the custom resource.
	Names CustomResourceDefinitionNames `pulumi:"names"`

	// preserveUnknownFields indicates that object fields which are not specified in the OpenAPI schema
	// should be preserved when persisting to storage. apiVersion, kind, metadata and known fields
	// inside metadata are always preserved. This field is deprecated in favor of setting
	// `x-preserve-unknown-fields` to true in `spec.versions[*].schema.openAPIV3Schema`. See
	// https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/#pruning-versus-preserving-unknown-fields
	// for details.
	PreserveUnknownFields *bool `pulumi:"preserveUnknownFields"`

	// scope indicates whether the defined custom resource is cluster- or namespace-scoped. Allowed
	// values are `Cluster` and `Namespaced`. Default is `Namespaced`.
	Scope string `pulumi:"scope"`

	// versions is the list of all API versions of the defined custom resource. Version names are used
	// to compute the order in which served versions are listed in API discovery. If the version string
	// is "kube-like", it will sort above non "kube-like" version strings, which are ordered
	// lexicographically. "Kube-like" versions start with a "v", then are followed by a number (the
	// major version), then optionally the string "alpha" or "beta" and another number (the minor
	// version). These are sorted first by GA > beta > alpha (where GA is a version with no suffix such
	// as beta or alpha), and then by comparing major version, then minor version. An example sorted
	// list of versions: v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10.
	Versions []CustomResourceDefinitionVersion `pulumi:"versions"`

}

var _CustomResourceDefinitionSpecType = reflect.TypeOf((*CustomResourceDefinitionSpec)(nil)).Elem()

// CustomResourceDefinitionSpecInput represents an input type that resolves to a CustomResourceDefinitionSpec.
type CustomResourceDefinitionSpecInput interface {
	ElementType() reflect.Type

	ToCustomResourceDefinitionSpecOutput() CustomResourceDefinitionSpecOutput
	ToCustomResourceDefinitionSpecOutputWithContext(ctx context.Context) CustomResourceDefinitionSpecOutput
}

// CustomResourceDefinitionSpecArgs is a CustomResourceDefinitionSpecInput whose fields are all Input types.
type CustomResourceDefinitionSpecArgs struct {
	// group is the API group of the defined custom resource. The custom resources are served under
	// `/apis/<group>/...`. Must match the name of the CustomResourceDefinition (in the form
	// `<names.plural>.<group>`).
	Group pulumi.StringInput `pulumi:"group"`

	// names specify the resource and kind names for the custom resource.
	Names CustomResourceDefinitionNamesInput `pulumi:"names"`

	// scope indicates whether the defined custom resource is cluster- or namespace-scoped. Allowed
	// values are `Cluster` and `Namespaced`. Default is `Namespaced`.
	Scope pulumi.StringInput `pulumi:"scope"`

	// versions is the list of all API versions of the defined custom resource. Version names are used
	// to compute the order in which served versions are listed in API discovery. If the version string
	// is "kube-like", it will sort above non "kube-like" version strings, which are ordered
	// lexicographically. "Kube-like" versions start with a "v", then are followed by a number (the
	// major version), then optionally the string "alpha" or "beta" and another number (the minor
	// version). These are sorted first by GA > beta > alpha (where GA is a version with no suffix such
	// as beta or alpha), and then by comparing major version, then minor version. An example sorted
	// list of versions: v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10.
	Versions CustomResourceDefinitionVersionArrayInput `pulumi:"versions"`

	// conversion defines conversion settings for the CRD.
	Conversion CustomResourceConversionInput `pulumi:"conversion"`

	// preserveUnknownFields indicates that object fields which are not specified in the OpenAPI schema
	// should be preserved when persisting to storage. apiVersion, kind, metadata and known fields
	// inside metadata are always preserved. This field is deprecated in favor of setting
	// `x-preserve-unknown-fields` to true in `spec.versions[*].schema.openAPIV3Schema`. See
	// https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/#pruning-versus-preserving-unknown-fields
	// for details.
	PreserveUnknownFields pulumi.BoolInput `pulumi:"preserveUnknownFields"`

}

func (a CustomResourceDefinitionSpecArgs) ElementType() reflect.Type {
	return _CustomResourceDefinitionSpecType
}

func (a CustomResourceDefinitionSpecArgs) ToCustomResourceDefinitionSpecOutput() CustomResourceDefinitionSpecOutput {
	return pulumi.ToOutput(a).(CustomResourceDefinitionSpecOutput)
}

func (a CustomResourceDefinitionSpecArgs) ToCustomResourceDefinitionSpecOutputWithContext(ctx context.Context) CustomResourceDefinitionSpecOutput {
	return pulumi.ToOutputWithContext(ctx, a).(CustomResourceDefinitionSpecOutput)
}

// CustomResourceDefinitionSpecOutput is an output type that resolves to a Input.
type CustomResourceDefinitionSpecOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(CustomResourceDefinitionSpecOutput{}) }

func (CustomResourceDefinitionSpecOutput) ElementType() reflect.Type {
	return _CustomResourceDefinitionSpecType
}

func (o CustomResourceDefinitionSpecOutput) Conversion() CustomResourceConversionOutput {
	return o.Apply(func(v CustomResourceDefinitionSpec) *CustomResourceConversion {
		return v.Conversion
	}).(CustomResourceConversionOutput)
}

func (o CustomResourceDefinitionSpecOutput) Group() pulumi.StringOutput {
	return o.Apply(func(v CustomResourceDefinitionSpec) string {
		return v.Group
	}).(pulumi.StringOutput)
}

func (o CustomResourceDefinitionSpecOutput) Names() CustomResourceDefinitionNamesOutput {
	return o.Apply(func(v CustomResourceDefinitionSpec) CustomResourceDefinitionNames {
		return v.Names
	}).(CustomResourceDefinitionNamesOutput)
}

func (o CustomResourceDefinitionSpecOutput) PreserveUnknownFields() pulumi.BoolOutput {
	return o.Apply(func(v CustomResourceDefinitionSpec) *bool {
		return v.PreserveUnknownFields
	}).(pulumi.BoolOutput)
}

func (o CustomResourceDefinitionSpecOutput) Scope() pulumi.StringOutput {
	return o.Apply(func(v CustomResourceDefinitionSpec) string {
		return v.Scope
	}).(pulumi.StringOutput)
}

func (o CustomResourceDefinitionSpecOutput) Versions() CustomResourceDefinitionVersionArrayOutput {
	return o.Apply(func(v CustomResourceDefinitionSpec) []CustomResourceDefinitionVersion {
		return v.Versions
	}).(CustomResourceDefinitionVersionArrayOutput)
}

