// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// TypedLocalObjectReference contains enough information to let you locate the typed referenced
// object inside the same namespace.
type TypedLocalObjectReference struct {
	// APIGroup is the group for the resource being referenced. If APIGroup is not specified, the
	// specified Kind must be in the core API group. For any other third-party types, APIGroup is
	// required.
	ApiGroup *string `pulumi:"apiGroup"`

	// Kind is the type of resource being referenced
	Kind string `pulumi:"kind"`

	// Name is the name of resource being referenced
	Name string `pulumi:"name"`

}

var _TypedLocalObjectReferenceType = reflect.TypeOf((*TypedLocalObjectReference)(nil)).Elem()

// TypedLocalObjectReferenceInput represents an input type that resolves to a TypedLocalObjectReference.
type TypedLocalObjectReferenceInput interface {
	ElementType() reflect.Type

	ToTypedLocalObjectReferenceOutput() TypedLocalObjectReferenceOutput
	ToTypedLocalObjectReferenceOutputWithContext(ctx context.Context) TypedLocalObjectReferenceOutput
}

// TypedLocalObjectReferenceArgs is a TypedLocalObjectReferenceInput whose fields are all Input types.
type TypedLocalObjectReferenceArgs struct {
	// Kind is the type of resource being referenced
	Kind pulumi.StringInput `pulumi:"kind"`

	// Name is the name of resource being referenced
	Name pulumi.StringInput `pulumi:"name"`

	// APIGroup is the group for the resource being referenced. If APIGroup is not specified, the
	// specified Kind must be in the core API group. For any other third-party types, APIGroup is
	// required.
	ApiGroup pulumi.StringInput `pulumi:"apiGroup"`

}

func (a TypedLocalObjectReferenceArgs) ElementType() reflect.Type {
	return _TypedLocalObjectReferenceType
}

func (a TypedLocalObjectReferenceArgs) ToTypedLocalObjectReferenceOutput() TypedLocalObjectReferenceOutput {
	return pulumi.ToOutput(a).(TypedLocalObjectReferenceOutput)
}

func (a TypedLocalObjectReferenceArgs) ToTypedLocalObjectReferenceOutputWithContext(ctx context.Context) TypedLocalObjectReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TypedLocalObjectReferenceOutput)
}

// TypedLocalObjectReferenceOutput is an output type that resolves to a Input.
type TypedLocalObjectReferenceOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(TypedLocalObjectReferenceOutput{}) }

func (TypedLocalObjectReferenceOutput) ElementType() reflect.Type {
	return _TypedLocalObjectReferenceType
}

func (o TypedLocalObjectReferenceOutput) ApiGroup() pulumi.StringOutput {
	return o.Apply(func(v TypedLocalObjectReference) *string {
		return v.ApiGroup
	}).(pulumi.StringOutput)
}

func (o TypedLocalObjectReferenceOutput) Kind() pulumi.StringOutput {
	return o.Apply(func(v TypedLocalObjectReference) string {
		return v.Kind
	}).(pulumi.StringOutput)
}

func (o TypedLocalObjectReferenceOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v TypedLocalObjectReference) string {
		return v.Name
	}).(pulumi.StringOutput)
}

