// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// Pod is a collection of containers that can run on a host. This resource is created by clients and
// scheduled onto hosts.
// 
// This resource waits until its status is ready before registering success
// for create/update, and populating output properties from the current state of the resource.
// The following conditions are used to determine whether the resource creation has
// succeeded or failed:
// 
// 1. The Pod is scheduled ("PodScheduled"" '.status.condition' is true).
// 2. The Pod is initialized ("Initialized" '.status.condition' is true).
// 3. The Pod is ready ("Ready" '.status.condition' is true) and the '.status.phase' is
//    set to "Running".
// Or (for Jobs): The Pod succeeded ('.status.phase' set to "Succeeded").
// 
// If the Pod has not reached a Ready state after 10 minutes, it will
// time out and mark the resource update as Failed. You can override the default timeout value
// by setting the 'customTimeouts' option on the resource.
type Pod struct {
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

	// Specification of the desired behavior of the pod. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec PodSpecOutput `pulumi:"spec"`

	// Most recently observed status of the pod. This data may not be up to date. Populated by the
	// system. Read-only. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status PodStatusOutput `pulumi:"status"`

}

// PodArgs is the set of arguments needed to create a Pod resource.
type PodArgs struct {
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

	// Specification of the desired behavior of the pod. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec PodSpecInput `pulumi:"spec"`

}

// NewPod creates a Pod resource with the given unique name, arguments, and options.
func NewPod(ctx *pulumi.Context, name string, args *PodArgs, opts ...pulumi.ResourceOption) (*Pod, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		args.ApiVersion = pulumi.String("v1")
		args.Kind = pulumi.String("Pod")
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
			inputs["spec"] = i.ToPodSpecOutput()
		}
	}
	var resource Pod
	err := ctx.RegisterResource("kubernetes:core/v1:Pod", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPod gets an existing Pod resource's state with the given name and ID.
func GetPod(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*Pod, error) {
	var resource Pod
	err := ctx.ReadResource("kubernetes:core/v1:Pod", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Pod is a collection of containers that can run on a host. This resource is created by clients and
// scheduled onto hosts.
type PodProperty struct {
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

	// Specification of the desired behavior of the pod. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *PodSpec `pulumi:"spec"`

	// Most recently observed status of the pod. This data may not be up to date. Populated by the
	// system. Read-only. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status PodStatus `pulumi:"status"`

}

var _PodPropertyType = reflect.TypeOf((*PodProperty)(nil)).Elem()

// PodPropertyInput represents an input type that resolves to a PodProperty.
type PodPropertyInput interface {
	ElementType() reflect.Type

	ToPodPropertyOutput() PodPropertyOutput
	ToPodPropertyOutputWithContext(ctx context.Context) PodPropertyOutput
}

// PodPropertyArgs is a PodPropertyInput whose fields are all Input types.
type PodPropertyArgs struct {
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

	// Specification of the desired behavior of the pod. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec PodSpecInput `pulumi:"spec"`

}

func (a PodPropertyArgs) ElementType() reflect.Type {
	return _PodPropertyType
}

func (a PodPropertyArgs) ToPodPropertyOutput() PodPropertyOutput {
	return pulumi.ToOutput(a).(PodPropertyOutput)
}

func (a PodPropertyArgs) ToPodPropertyOutputWithContext(ctx context.Context) PodPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, a).(PodPropertyOutput)
}

// PodPropertyOutput is an output type that resolves to a Input.
type PodPropertyOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(PodPropertyOutput{}) }

func (PodPropertyOutput) ElementType() reflect.Type {
	return _PodPropertyType
}

func (o PodPropertyOutput) ApiVersion() pulumi.StringOutput {
	return o.Apply(func(v PodProperty) *string {
		return v.ApiVersion
	}).(pulumi.StringOutput)
}

func (o PodPropertyOutput) Kind() pulumi.StringOutput {
	return o.Apply(func(v PodProperty) *string {
		return v.Kind
	}).(pulumi.StringOutput)
}

func (o PodPropertyOutput) Metadata() metaV1.ObjectMetaOutput {
	return o.Apply(func(v PodProperty) *metaV1.ObjectMeta {
		return v.Metadata
	}).(metaV1.ObjectMetaOutput)
}

func (o PodPropertyOutput) Spec() PodSpecOutput {
	return o.Apply(func(v PodProperty) *PodSpec {
		return v.Spec
	}).(PodSpecOutput)
}

func (o PodPropertyOutput) Status() PodStatusOutput {
	return o.Apply(func(v PodProperty) PodStatus {
		return v.Status
	}).(PodStatusOutput)
}

var _PodPropertyArrayType = reflect.TypeOf((*[]PodProperty)(nil)).Elem()

type PodPropertyArrayInput interface {
	ElementType() reflect.Type

	ToPodPropertyArrayOutput() PodPropertyArrayOutput
	ToPodPropertyArrayOutputWithContext(ctx context.Context) PodPropertyArrayOutput
}

type PodPropertyArray []PodPropertyInput

func (a PodPropertyArray) ElementType() reflect.Type {
	return _PodPropertyArrayType
}

func (a PodPropertyArray) ToPodPropertyArrayOutput() PodPropertyArrayOutput {
	return pulumi.ToOutput(a).(PodPropertyArrayOutput)
}

func (a PodPropertyArray) ToPodPropertyArrayOutputWithContext(ctx context.Context) PodPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(PodPropertyArrayOutput)
}

type PodPropertyArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(PodPropertyArrayOutput{}) }

func (PodPropertyArrayOutput) ElementType() reflect.Type {
	return _PodPropertyArrayType
}

func (o PodPropertyArrayOutput) Index(i pulumi.IntInput) PodPropertyOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) PodProperty {
		return vs[0].([]PodProperty)[vs[1].(int)]
	}).(PodPropertyOutput)
}