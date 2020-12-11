// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// CronJobList is a collection of cron jobs.
type CronJobList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// items is the list of CronJobs.
	Items CronJobTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewCronJobList registers a new resource with the given unique name, arguments, and options.
func NewCronJobList(ctx *pulumi.Context,
	name string, args *CronJobListArgs, opts ...pulumi.ResourceOption) (*CronJobList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("batch/v1beta1")
	args.Kind = pulumi.StringPtr("CronJobList")
	var resource CronJobList
	err := ctx.RegisterResource("kubernetes:batch/v1beta1:CronJobList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCronJobList gets an existing CronJobList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCronJobList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CronJobListState, opts ...pulumi.ResourceOption) (*CronJobList, error) {
	var resource CronJobList
	err := ctx.ReadResource("kubernetes:batch/v1beta1:CronJobList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CronJobList resources.
type cronJobListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is the list of CronJobs.
	Items []CronJobType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type CronJobListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is the list of CronJobs.
	Items CronJobTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (CronJobListState) ElementType() reflect.Type {
	return reflect.TypeOf((*cronJobListState)(nil)).Elem()
}

type cronJobListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is the list of CronJobs.
	Items []CronJobType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a CronJobList resource.
type CronJobListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is the list of CronJobs.
	Items CronJobTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (CronJobListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cronJobListArgs)(nil)).Elem()
}

type CronJobListInput interface {
	pulumi.Input

	ToCronJobListOutput() CronJobListOutput
	ToCronJobListOutputWithContext(ctx context.Context) CronJobListOutput
}

type CronJobListPtrInput interface {
	pulumi.Input

	ToCronJobListPtrOutput() CronJobListPtrOutput
	ToCronJobListPtrOutputWithContext(ctx context.Context) CronJobListPtrOutput
}

func (CronJobList) ElementType() reflect.Type {
	return reflect.TypeOf((*CronJobList)(nil)).Elem()
}

func (i CronJobList) ToCronJobListOutput() CronJobListOutput {
	return i.ToCronJobListOutputWithContext(context.Background())
}

func (i CronJobList) ToCronJobListOutputWithContext(ctx context.Context) CronJobListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobListOutput)
}

func (i CronJobList) ToCronJobListPtrOutput() CronJobListPtrOutput {
	return i.ToCronJobListPtrOutputWithContext(context.Background())
}

func (i CronJobList) ToCronJobListPtrOutputWithContext(ctx context.Context) CronJobListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobListPtrOutput)
}

type CronJobListOutput struct {
	*pulumi.OutputState
}

func (CronJobListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CronJobListOutput)(nil)).Elem()
}

func (o CronJobListOutput) ToCronJobListOutput() CronJobListOutput {
	return o
}

func (o CronJobListOutput) ToCronJobListOutputWithContext(ctx context.Context) CronJobListOutput {
	return o
}

type CronJobListPtrOutput struct {
	*pulumi.OutputState
}

func (CronJobListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CronJobList)(nil)).Elem()
}

func (o CronJobListPtrOutput) ToCronJobListPtrOutput() CronJobListPtrOutput {
	return o
}

func (o CronJobListPtrOutput) ToCronJobListPtrOutputWithContext(ctx context.Context) CronJobListPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CronJobListOutput{})
	pulumi.RegisterOutputType(CronJobListPtrOutput{})
}
