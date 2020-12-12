// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// DeploymentList is a list of Deployments.
type DeploymentList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Items is the list of Deployments.
	Items DeploymentTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata.
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewDeploymentList registers a new resource with the given unique name, arguments, and options.
func NewDeploymentList(ctx *pulumi.Context,
	name string, args *DeploymentListArgs, opts ...pulumi.ResourceOption) (*DeploymentList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("apps/v1")
	args.Kind = pulumi.StringPtr("DeploymentList")
	var resource DeploymentList
	err := ctx.RegisterResource("kubernetes:apps/v1:DeploymentList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeploymentList gets an existing DeploymentList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeploymentList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentListState, opts ...pulumi.ResourceOption) (*DeploymentList, error) {
	var resource DeploymentList
	err := ctx.ReadResource("kubernetes:apps/v1:DeploymentList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeploymentList resources.
type deploymentListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is the list of Deployments.
	Items []DeploymentType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata.
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type DeploymentListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is the list of Deployments.
	Items DeploymentTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata.
	Metadata metav1.ListMetaPtrInput
}

func (DeploymentListState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentListState)(nil)).Elem()
}

type deploymentListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is the list of Deployments.
	Items []DeploymentType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata.
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a DeploymentList resource.
type DeploymentListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is the list of Deployments.
	Items DeploymentTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata.
	Metadata metav1.ListMetaPtrInput
}

func (DeploymentListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentListArgs)(nil)).Elem()
}

type DeploymentListInput interface {
	pulumi.Input

	ToDeploymentListOutput() DeploymentListOutput
	ToDeploymentListOutputWithContext(ctx context.Context) DeploymentListOutput
}

func (*DeploymentList) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentList)(nil))
}

func (i *DeploymentList) ToDeploymentListOutput() DeploymentListOutput {
	return i.ToDeploymentListOutputWithContext(context.Background())
}

func (i *DeploymentList) ToDeploymentListOutputWithContext(ctx context.Context) DeploymentListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentListOutput)
}

func (i *DeploymentList) ToDeploymentListPtrOutput() DeploymentListPtrOutput {
	return i.ToDeploymentListPtrOutputWithContext(context.Background())
}

func (i *DeploymentList) ToDeploymentListPtrOutputWithContext(ctx context.Context) DeploymentListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentListPtrOutput)
}

type DeploymentListPtrInput interface {
	pulumi.Input

	ToDeploymentListPtrOutput() DeploymentListPtrOutput
	ToDeploymentListPtrOutputWithContext(ctx context.Context) DeploymentListPtrOutput
}

type DeploymentListOutput struct {
	*pulumi.OutputState
}

func (DeploymentListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentList)(nil))
}

func (o DeploymentListOutput) ToDeploymentListOutput() DeploymentListOutput {
	return o
}

func (o DeploymentListOutput) ToDeploymentListOutputWithContext(ctx context.Context) DeploymentListOutput {
	return o
}

type DeploymentListPtrOutput struct {
	*pulumi.OutputState
}

func (DeploymentListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentList)(nil))
}

func (o DeploymentListPtrOutput) ToDeploymentListPtrOutput() DeploymentListPtrOutput {
	return o
}

func (o DeploymentListPtrOutput) ToDeploymentListPtrOutputWithContext(ctx context.Context) DeploymentListPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DeploymentListOutput{})
	pulumi.RegisterOutputType(DeploymentListPtrOutput{})
}
