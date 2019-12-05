// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// DEPRECATED - apps/v1beta2/DaemonSet is not supported by Kubernetes 1.16+ clusters. Use
// apps/v1/DaemonSet instead.
// 
// DaemonSet represents the configuration of a daemon set.
type DaemonSet struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaOutput `pulumi:"metadata"`

	// The desired behavior of this daemon set. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec DaemonSetSpecOutput `pulumi:"spec"`

	// The current status of this daemon set. This data may be out of date by some window of time.
	// Populated by the system. Read-only. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status DaemonSetStatusOutput `pulumi:"status"`

}

// DaemonSetArgs is the set of arguments needed to create a DaemonSet resource.
type DaemonSetArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringInput `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringInput `pulumi:"kind"`

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// The desired behavior of this daemon set. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec DaemonSetSpecInput `pulumi:"spec"`

}

// NewDaemonSet creates a DaemonSet resource with the given unique name, arguments, and options.
func NewDaemonSet(ctx *pulumi.Context, name string, args *DaemonSetArgs, opts ...pulumi.ResourceOption) (*DaemonSet, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		args.ApiVersion = pulumi.String("apps/v1beta2")
		args.Kind = pulumi.String("DaemonSet")
		if i := args.ApiVersion; i != nil {
			inputs["apiVersion"] = i.ToStringOutput()
		}
		if i := args.Kind; i != nil {
			inputs["kind"] = i.ToStringOutput()
		}
		if i := args.Metadata; i != nil {
			inputs["metadata"] = i.ToObjectMetaOutput()
		}
		if i := args.Spec; i != nil {
			inputs["spec"] = i.ToDaemonSetSpecOutput()
		}
	}
	var resource DaemonSet
	err := ctx.RegisterResource("kubernetes:apps/v1beta2:DaemonSet", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDaemonSet gets an existing DaemonSet resource's state with the given name and ID.
func GetDaemonSet(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*DaemonSet, error) {
	var resource DaemonSet
	err := ctx.ReadResource("kubernetes:apps/v1beta2:DaemonSet", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// DEPRECATED - apps/v1beta2/DaemonSet is not supported by Kubernetes 1.16+ clusters. Use
// apps/v1/DaemonSet instead.
// 
// DaemonSet represents the configuration of a daemon set.
type DaemonSetProperty struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metaV1.ObjectMeta `pulumi:"metadata"`

	// The desired behavior of this daemon set. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *DaemonSetSpec `pulumi:"spec"`

	// The current status of this daemon set. This data may be out of date by some window of time.
	// Populated by the system. Read-only. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status DaemonSetStatus `pulumi:"status"`

}

var _DaemonSetPropertyType = reflect.TypeOf((*DaemonSetProperty)(nil)).Elem()

// DaemonSetPropertyInput represents an input type that resolves to a DaemonSetProperty.
type DaemonSetPropertyInput interface {
	ElementType() reflect.Type

	ToDaemonSetPropertyOutput() DaemonSetPropertyOutput
	ToDaemonSetPropertyOutputWithContext(ctx context.Context) DaemonSetPropertyOutput
}

// DaemonSetPropertyArgs is a DaemonSetPropertyInput whose fields are all Input types.
type DaemonSetPropertyArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringInput `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringInput `pulumi:"kind"`

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// The desired behavior of this daemon set. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec DaemonSetSpecInput `pulumi:"spec"`

}

func (a DaemonSetPropertyArgs) ElementType() reflect.Type {
	return _DaemonSetPropertyType
}

func (a DaemonSetPropertyArgs) ToDaemonSetPropertyOutput() DaemonSetPropertyOutput {
	return pulumi.ToOutput(a).(DaemonSetPropertyOutput)
}

func (a DaemonSetPropertyArgs) ToDaemonSetPropertyOutputWithContext(ctx context.Context) DaemonSetPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DaemonSetPropertyOutput)
}

// DaemonSetPropertyOutput is an output type that resolves to a Input.
type DaemonSetPropertyOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(DaemonSetPropertyOutput{}) }

func (DaemonSetPropertyOutput) ElementType() reflect.Type {
	return _DaemonSetPropertyType
}

func (o DaemonSetPropertyOutput) ApiVersion() pulumi.StringOutput {
	return o.Apply(func(v DaemonSetProperty) *string {
		return v.ApiVersion
	}).(pulumi.StringOutput)
}

func (o DaemonSetPropertyOutput) Kind() pulumi.StringOutput {
	return o.Apply(func(v DaemonSetProperty) *string {
		return v.Kind
	}).(pulumi.StringOutput)
}

func (o DaemonSetPropertyOutput) Metadata() metaV1.ObjectMetaOutput {
	return o.Apply(func(v DaemonSetProperty) *metaV1.ObjectMeta {
		return v.Metadata
	}).(metaV1.ObjectMetaOutput)
}

func (o DaemonSetPropertyOutput) Spec() DaemonSetSpecOutput {
	return o.Apply(func(v DaemonSetProperty) *DaemonSetSpec {
		return v.Spec
	}).(DaemonSetSpecOutput)
}

func (o DaemonSetPropertyOutput) Status() DaemonSetStatusOutput {
	return o.Apply(func(v DaemonSetProperty) DaemonSetStatus {
		return v.Status
	}).(DaemonSetStatusOutput)
}

var _DaemonSetPropertyArrayType = reflect.TypeOf((*[]DaemonSetProperty)(nil)).Elem()

type DaemonSetPropertyArrayInput interface {
	ElementType() reflect.Type

	ToDaemonSetPropertyArrayOutput() DaemonSetPropertyArrayOutput
	ToDaemonSetPropertyArrayOutputWithContext(ctx context.Context) DaemonSetPropertyArrayOutput
}

type DaemonSetPropertyArray []DaemonSetPropertyInput

func (a DaemonSetPropertyArray) ElementType() reflect.Type {
	return _DaemonSetPropertyArrayType
}

func (a DaemonSetPropertyArray) ToDaemonSetPropertyArrayOutput() DaemonSetPropertyArrayOutput {
	return pulumi.ToOutput(a).(DaemonSetPropertyArrayOutput)
}

func (a DaemonSetPropertyArray) ToDaemonSetPropertyArrayOutputWithContext(ctx context.Context) DaemonSetPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DaemonSetPropertyArrayOutput)
}

type DaemonSetPropertyArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(DaemonSetPropertyArrayOutput{}) }

func (DaemonSetPropertyArrayOutput) ElementType() reflect.Type {
	return _DaemonSetPropertyArrayType
}

func (o DaemonSetPropertyArrayOutput) Index(i pulumi.IntInput) DaemonSetPropertyOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) DaemonSetProperty {
		return vs[0].([]DaemonSetProperty)[vs[1].(int)]
	}).(DaemonSetPropertyOutput)
}