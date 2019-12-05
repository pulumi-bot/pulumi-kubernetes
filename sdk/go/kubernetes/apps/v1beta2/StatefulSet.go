// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// DEPRECATED - apps/v1beta2/StatefulSet is not supported by Kubernetes 1.16+ clusters. Use
// apps/v1/StatefulSet instead.
// 
// StatefulSet represents a set of pods with consistent identities. Identities are defined as:
//  - Network: A single stable DNS and hostname.
//  - Storage: As many VolumeClaims as requested.
// The StatefulSet guarantees that a given network identity will always map to the same storage
// identity.
// 
// This resource waits until its status is ready before registering success
// for create/update, and populating output properties from the current state of the resource.
// The following conditions are used to determine whether the resource creation has
// succeeded or failed:
// 
// 1. The value of 'spec.replicas' matches '.status.replicas', '.status.currentReplicas',
//    and '.status.readyReplicas'.
// 2. The value of '.status.updateRevision' matches '.status.currentRevision'.
// 
// If the StatefulSet has not reached a Ready state after 10 minutes, it will
// time out and mark the resource update as Failed. You can override the default timeout value
// by setting the 'customTimeouts' option on the resource.
type StatefulSet struct {
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

	
	Metadata metaV1.ObjectMetaOutput `pulumi:"metadata"`

	// Spec defines the desired identities of pods in this set.
	Spec StatefulSetSpecOutput `pulumi:"spec"`

	// Status is the current status of Pods in this StatefulSet. This data may be out of date by some
	// window of time.
	Status StatefulSetStatusOutput `pulumi:"status"`

}

// StatefulSetArgs is the set of arguments needed to create a StatefulSet resource.
type StatefulSetArgs struct {
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

	
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// Spec defines the desired identities of pods in this set.
	Spec StatefulSetSpecInput `pulumi:"spec"`

}

// NewStatefulSet creates a StatefulSet resource with the given unique name, arguments, and options.
func NewStatefulSet(ctx *pulumi.Context, name string, args *StatefulSetArgs, opts ...pulumi.ResourceOption) (*StatefulSet, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		args.ApiVersion = pulumi.String("apps/v1beta2")
		args.Kind = pulumi.String("StatefulSet")
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
			inputs["spec"] = i.ToStatefulSetSpecOutput()
		}
	}
	var resource StatefulSet
	err := ctx.RegisterResource("kubernetes:apps/v1beta2:StatefulSet", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStatefulSet gets an existing StatefulSet resource's state with the given name and ID.
func GetStatefulSet(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*StatefulSet, error) {
	var resource StatefulSet
	err := ctx.ReadResource("kubernetes:apps/v1beta2:StatefulSet", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// DEPRECATED - apps/v1beta2/StatefulSet is not supported by Kubernetes 1.16+ clusters. Use
// apps/v1/StatefulSet instead.
// 
// StatefulSet represents a set of pods with consistent identities. Identities are defined as:
//  - Network: A single stable DNS and hostname.
//  - Storage: As many VolumeClaims as requested.
// The StatefulSet guarantees that a given network identity will always map to the same storage
// identity.
type StatefulSetProperty struct {
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

	
	Metadata *metaV1.ObjectMeta `pulumi:"metadata"`

	// Spec defines the desired identities of pods in this set.
	Spec *StatefulSetSpec `pulumi:"spec"`

	// Status is the current status of Pods in this StatefulSet. This data may be out of date by some
	// window of time.
	Status StatefulSetStatus `pulumi:"status"`

}

var _StatefulSetPropertyType = reflect.TypeOf((*StatefulSetProperty)(nil)).Elem()

// StatefulSetPropertyInput represents an input type that resolves to a StatefulSetProperty.
type StatefulSetPropertyInput interface {
	ElementType() reflect.Type

	ToStatefulSetPropertyOutput() StatefulSetPropertyOutput
	ToStatefulSetPropertyOutputWithContext(ctx context.Context) StatefulSetPropertyOutput
}

// StatefulSetPropertyArgs is a StatefulSetPropertyInput whose fields are all Input types.
type StatefulSetPropertyArgs struct {
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

	
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// Spec defines the desired identities of pods in this set.
	Spec StatefulSetSpecInput `pulumi:"spec"`

}

func (a StatefulSetPropertyArgs) ElementType() reflect.Type {
	return _StatefulSetPropertyType
}

func (a StatefulSetPropertyArgs) ToStatefulSetPropertyOutput() StatefulSetPropertyOutput {
	return pulumi.ToOutput(a).(StatefulSetPropertyOutput)
}

func (a StatefulSetPropertyArgs) ToStatefulSetPropertyOutputWithContext(ctx context.Context) StatefulSetPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, a).(StatefulSetPropertyOutput)
}

// StatefulSetPropertyOutput is an output type that resolves to a Input.
type StatefulSetPropertyOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(StatefulSetPropertyOutput{}) }

func (StatefulSetPropertyOutput) ElementType() reflect.Type {
	return _StatefulSetPropertyType
}

func (o StatefulSetPropertyOutput) ApiVersion() pulumi.StringOutput {
	return o.Apply(func(v StatefulSetProperty) *string {
		return v.ApiVersion
	}).(pulumi.StringOutput)
}

func (o StatefulSetPropertyOutput) Kind() pulumi.StringOutput {
	return o.Apply(func(v StatefulSetProperty) *string {
		return v.Kind
	}).(pulumi.StringOutput)
}

func (o StatefulSetPropertyOutput) Metadata() metaV1.ObjectMetaOutput {
	return o.Apply(func(v StatefulSetProperty) *metaV1.ObjectMeta {
		return v.Metadata
	}).(metaV1.ObjectMetaOutput)
}

func (o StatefulSetPropertyOutput) Spec() StatefulSetSpecOutput {
	return o.Apply(func(v StatefulSetProperty) *StatefulSetSpec {
		return v.Spec
	}).(StatefulSetSpecOutput)
}

func (o StatefulSetPropertyOutput) Status() StatefulSetStatusOutput {
	return o.Apply(func(v StatefulSetProperty) StatefulSetStatus {
		return v.Status
	}).(StatefulSetStatusOutput)
}

var _StatefulSetPropertyArrayType = reflect.TypeOf((*[]StatefulSetProperty)(nil)).Elem()

type StatefulSetPropertyArrayInput interface {
	ElementType() reflect.Type

	ToStatefulSetPropertyArrayOutput() StatefulSetPropertyArrayOutput
	ToStatefulSetPropertyArrayOutputWithContext(ctx context.Context) StatefulSetPropertyArrayOutput
}

type StatefulSetPropertyArray []StatefulSetPropertyInput

func (a StatefulSetPropertyArray) ElementType() reflect.Type {
	return _StatefulSetPropertyArrayType
}

func (a StatefulSetPropertyArray) ToStatefulSetPropertyArrayOutput() StatefulSetPropertyArrayOutput {
	return pulumi.ToOutput(a).(StatefulSetPropertyArrayOutput)
}

func (a StatefulSetPropertyArray) ToStatefulSetPropertyArrayOutputWithContext(ctx context.Context) StatefulSetPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(StatefulSetPropertyArrayOutput)
}

type StatefulSetPropertyArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(StatefulSetPropertyArrayOutput{}) }

func (StatefulSetPropertyArrayOutput) ElementType() reflect.Type {
	return _StatefulSetPropertyArrayType
}

func (o StatefulSetPropertyArrayOutput) Index(i pulumi.IntInput) StatefulSetPropertyOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) StatefulSetProperty {
		return vs[0].([]StatefulSetProperty)[vs[1].(int)]
	}).(StatefulSetPropertyOutput)
}
