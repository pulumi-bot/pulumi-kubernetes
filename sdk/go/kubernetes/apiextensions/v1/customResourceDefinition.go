// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// CustomResourceDefinition represents a resource that should be exposed on the API server.  Its name MUST be in the format <.spec.name>.<.spec.group>.
type CustomResourceDefinition struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// spec describes how the user wants the resources to appear
	Spec CustomResourceDefinitionSpecOutput `pulumi:"spec"`
	// status indicates the actual state of the CustomResourceDefinition
	Status CustomResourceDefinitionStatusPtrOutput `pulumi:"status"`
}

// NewCustomResourceDefinition registers a new resource with the given unique name, arguments, and options.
func NewCustomResourceDefinition(ctx *pulumi.Context,
	name string, args *CustomResourceDefinitionArgs, opts ...pulumi.ResourceOption) (*CustomResourceDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	args.ApiVersion = pulumi.StringPtr("apiextensions.k8s.io/v1")
	args.Kind = pulumi.StringPtr("CustomResourceDefinition")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:apiextensions.k8s.io/v1beta1:CustomResourceDefinition"),
		},
	})
	opts = append(opts, aliases)
	var resource CustomResourceDefinition
	err := ctx.RegisterResource("kubernetes:apiextensions.k8s.io/v1:CustomResourceDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomResourceDefinition gets an existing CustomResourceDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomResourceDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomResourceDefinitionState, opts ...pulumi.ResourceOption) (*CustomResourceDefinition, error) {
	var resource CustomResourceDefinition
	err := ctx.ReadResource("kubernetes:apiextensions.k8s.io/v1:CustomResourceDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomResourceDefinition resources.
type customResourceDefinitionState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// spec describes how the user wants the resources to appear
	Spec *CustomResourceDefinitionSpec `pulumi:"spec"`
	// status indicates the actual state of the CustomResourceDefinition
	Status *CustomResourceDefinitionStatus `pulumi:"status"`
}

type CustomResourceDefinitionState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// spec describes how the user wants the resources to appear
	Spec CustomResourceDefinitionSpecPtrInput
	// status indicates the actual state of the CustomResourceDefinition
	Status CustomResourceDefinitionStatusPtrInput
}

func (CustomResourceDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*customResourceDefinitionState)(nil)).Elem()
}

type customResourceDefinitionArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// spec describes how the user wants the resources to appear
	Spec CustomResourceDefinitionSpec `pulumi:"spec"`
}

// The set of arguments for constructing a CustomResourceDefinition resource.
type CustomResourceDefinitionArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// spec describes how the user wants the resources to appear
	Spec CustomResourceDefinitionSpecInput
}

func (CustomResourceDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customResourceDefinitionArgs)(nil)).Elem()
}

type CustomResourceDefinitionInput interface {
	pulumi.Input

	ToCustomResourceDefinitionOutput() CustomResourceDefinitionOutput
	ToCustomResourceDefinitionOutputWithContext(ctx context.Context) CustomResourceDefinitionOutput
}

func (CustomResourceDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomResourceDefinition)(nil)).Elem()
}

func (i CustomResourceDefinition) ToCustomResourceDefinitionOutput() CustomResourceDefinitionOutput {
	return i.ToCustomResourceDefinitionOutputWithContext(context.Background())
}

func (i CustomResourceDefinition) ToCustomResourceDefinitionOutputWithContext(ctx context.Context) CustomResourceDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomResourceDefinitionOutput)
}

type CustomResourceDefinitionOutput struct {
	*pulumi.OutputState
}

func (CustomResourceDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomResourceDefinitionOutput)(nil)).Elem()
}

func (o CustomResourceDefinitionOutput) ToCustomResourceDefinitionOutput() CustomResourceDefinitionOutput {
	return o
}

func (o CustomResourceDefinitionOutput) ToCustomResourceDefinitionOutputWithContext(ctx context.Context) CustomResourceDefinitionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CustomResourceDefinitionOutput{})
}
