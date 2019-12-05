// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// VolumeAttachment captures the intent to attach or detach the specified volume to/from the
// specified node.
// 
// VolumeAttachment objects are non-namespaced.
type VolumeAttachment struct {
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

	// Standard object metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaOutput `pulumi:"metadata"`

	// Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.
	Spec VolumeAttachmentSpecOutput `pulumi:"spec"`

	// Status of the VolumeAttachment request. Populated by the entity completing the attach or detach
	// operation, i.e. the external-attacher.
	Status VolumeAttachmentStatusOutput `pulumi:"status"`

}

// VolumeAttachmentArgs is the set of arguments needed to create a VolumeAttachment resource.
type VolumeAttachmentArgs struct {
	// Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.
	Spec VolumeAttachmentSpecInput `pulumi:"spec"`

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

	// Standard object metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

}

// NewVolumeAttachment creates a VolumeAttachment resource with the given unique name, arguments, and options.
func NewVolumeAttachment(ctx *pulumi.Context, name string, args *VolumeAttachmentArgs, opts ...pulumi.ResourceOption) (*VolumeAttachment, error) {
	inputs := map[string]pulumi.Input{}
	if args == nil || args.Spec == nil {
		return nil, errors.New("missing required argument 'Spec'")
	}
	if args != nil {
		args.ApiVersion = pulumi.String("storage.k8s.io/v1beta1")
		args.Kind = pulumi.String("VolumeAttachment")
		inputs["spec"] = args.Spec.ToVolumeAttachmentSpecOutput()
		if i := args.ApiVersion; i != nil {
			inputs["apiVersion"] = i.ToStringOutput()
		}
		if i := args.Kind; i != nil {
			inputs["kind"] = i.ToStringOutput()
		}
		if i := args.Metadata; i != nil {
			inputs["metadata"] = i.ToObjectMetaOutput()
		}
	}
	var resource VolumeAttachment
	err := ctx.RegisterResource("kubernetes:storage.k8s.io/v1beta1:VolumeAttachment", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolumeAttachment gets an existing VolumeAttachment resource's state with the given name and ID.
func GetVolumeAttachment(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*VolumeAttachment, error) {
	var resource VolumeAttachment
	err := ctx.ReadResource("kubernetes:storage.k8s.io/v1beta1:VolumeAttachment", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// VolumeAttachment captures the intent to attach or detach the specified volume to/from the
// specified node.
// 
// VolumeAttachment objects are non-namespaced.
type VolumeAttachmentProperty struct {
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

	// Standard object metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metaV1.ObjectMeta `pulumi:"metadata"`

	// Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.
	Spec VolumeAttachmentSpec `pulumi:"spec"`

	// Status of the VolumeAttachment request. Populated by the entity completing the attach or detach
	// operation, i.e. the external-attacher.
	Status VolumeAttachmentStatus `pulumi:"status"`

}

var _VolumeAttachmentPropertyType = reflect.TypeOf((*VolumeAttachmentProperty)(nil)).Elem()

// VolumeAttachmentPropertyInput represents an input type that resolves to a VolumeAttachmentProperty.
type VolumeAttachmentPropertyInput interface {
	ElementType() reflect.Type

	ToVolumeAttachmentPropertyOutput() VolumeAttachmentPropertyOutput
	ToVolumeAttachmentPropertyOutputWithContext(ctx context.Context) VolumeAttachmentPropertyOutput
}

// VolumeAttachmentPropertyArgs is a VolumeAttachmentPropertyInput whose fields are all Input types.
type VolumeAttachmentPropertyArgs struct {
	// Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.
	Spec VolumeAttachmentSpecInput `pulumi:"spec"`

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

	// Standard object metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

}

func (a VolumeAttachmentPropertyArgs) ElementType() reflect.Type {
	return _VolumeAttachmentPropertyType
}

func (a VolumeAttachmentPropertyArgs) ToVolumeAttachmentPropertyOutput() VolumeAttachmentPropertyOutput {
	return pulumi.ToOutput(a).(VolumeAttachmentPropertyOutput)
}

func (a VolumeAttachmentPropertyArgs) ToVolumeAttachmentPropertyOutputWithContext(ctx context.Context) VolumeAttachmentPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, a).(VolumeAttachmentPropertyOutput)
}

// VolumeAttachmentPropertyOutput is an output type that resolves to a Input.
type VolumeAttachmentPropertyOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(VolumeAttachmentPropertyOutput{}) }

func (VolumeAttachmentPropertyOutput) ElementType() reflect.Type {
	return _VolumeAttachmentPropertyType
}

func (o VolumeAttachmentPropertyOutput) ApiVersion() pulumi.StringOutput {
	return o.Apply(func(v VolumeAttachmentProperty) *string {
		return v.ApiVersion
	}).(pulumi.StringOutput)
}

func (o VolumeAttachmentPropertyOutput) Kind() pulumi.StringOutput {
	return o.Apply(func(v VolumeAttachmentProperty) *string {
		return v.Kind
	}).(pulumi.StringOutput)
}

func (o VolumeAttachmentPropertyOutput) Metadata() metaV1.ObjectMetaOutput {
	return o.Apply(func(v VolumeAttachmentProperty) *metaV1.ObjectMeta {
		return v.Metadata
	}).(metaV1.ObjectMetaOutput)
}

func (o VolumeAttachmentPropertyOutput) Spec() VolumeAttachmentSpecOutput {
	return o.Apply(func(v VolumeAttachmentProperty) VolumeAttachmentSpec {
		return v.Spec
	}).(VolumeAttachmentSpecOutput)
}

func (o VolumeAttachmentPropertyOutput) Status() VolumeAttachmentStatusOutput {
	return o.Apply(func(v VolumeAttachmentProperty) VolumeAttachmentStatus {
		return v.Status
	}).(VolumeAttachmentStatusOutput)
}

var _VolumeAttachmentPropertyArrayType = reflect.TypeOf((*[]VolumeAttachmentProperty)(nil)).Elem()

type VolumeAttachmentPropertyArrayInput interface {
	ElementType() reflect.Type

	ToVolumeAttachmentPropertyArrayOutput() VolumeAttachmentPropertyArrayOutput
	ToVolumeAttachmentPropertyArrayOutputWithContext(ctx context.Context) VolumeAttachmentPropertyArrayOutput
}

type VolumeAttachmentPropertyArray []VolumeAttachmentPropertyInput

func (a VolumeAttachmentPropertyArray) ElementType() reflect.Type {
	return _VolumeAttachmentPropertyArrayType
}

func (a VolumeAttachmentPropertyArray) ToVolumeAttachmentPropertyArrayOutput() VolumeAttachmentPropertyArrayOutput {
	return pulumi.ToOutput(a).(VolumeAttachmentPropertyArrayOutput)
}

func (a VolumeAttachmentPropertyArray) ToVolumeAttachmentPropertyArrayOutputWithContext(ctx context.Context) VolumeAttachmentPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(VolumeAttachmentPropertyArrayOutput)
}

type VolumeAttachmentPropertyArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(VolumeAttachmentPropertyArrayOutput{}) }

func (VolumeAttachmentPropertyArrayOutput) ElementType() reflect.Type {
	return _VolumeAttachmentPropertyArrayType
}

func (o VolumeAttachmentPropertyArrayOutput) Index(i pulumi.IntInput) VolumeAttachmentPropertyOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) VolumeAttachmentProperty {
		return vs[0].([]VolumeAttachmentProperty)[vs[1].(int)]
	}).(VolumeAttachmentPropertyOutput)
}
