// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// ClusterRole is a cluster level, logical grouping of PolicyRules that can be referenced as a unit
// by a RoleBinding or ClusterRoleBinding.
type ClusterRole struct {
	pulumi.CustomResourceState

	// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole.
	// If AggregationRule is set, then the Rules are controller managed and direct changes to Rules
	// will be stomped by the controller.
	AggregationRule AggregationRuleOutput `pulumi:"aggregationRule"`

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

	// Standard object's metadata.
	Metadata metaV1.ObjectMetaOutput `pulumi:"metadata"`

	// Rules holds all the PolicyRules for this ClusterRole
	Rules PolicyRuleArrayOutput `pulumi:"rules"`

}

// ClusterRoleArgs is the set of arguments needed to create a ClusterRole resource.
type ClusterRoleArgs struct {
	// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole.
	// If AggregationRule is set, then the Rules are controller managed and direct changes to Rules
	// will be stomped by the controller.
	AggregationRule AggregationRuleInput `pulumi:"aggregationRule"`

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

	// Standard object's metadata.
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// Rules holds all the PolicyRules for this ClusterRole
	Rules PolicyRuleArrayInput `pulumi:"rules"`

}

// NewClusterRole creates a ClusterRole resource with the given unique name, arguments, and options.
func NewClusterRole(ctx *pulumi.Context, name string, args *ClusterRoleArgs, opts ...pulumi.ResourceOption) (*ClusterRole, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		args.ApiVersion = pulumi.String("rbac.authorization.k8s.io/v1alpha1")
		args.Kind = pulumi.String("ClusterRole")
		if i := args.AggregationRule; i != nil {
			inputs["aggregationRule"] = i.ToAggregationRuleOutput()
		}
		if i := args.ApiVersion; i != nil {
			inputs["apiVersion"] = i.ToStringOutput()
		}
		if i := args.Kind; i != nil {
			inputs["kind"] = i.ToStringOutput()
		}
		if i := args.Metadata; i != nil {
			inputs["metadata"] = i.ToObjectMetaOutput()
		}
		if i := args.Rules; i != nil {
			inputs["rules"] = i.ToPolicyRuleArrayOutput()
		}
	}
	var resource ClusterRole
	err := ctx.RegisterResource("kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRole", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterRole gets an existing ClusterRole resource's state with the given name and ID.
func GetClusterRole(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*ClusterRole, error) {
	var resource ClusterRole
	err := ctx.ReadResource("kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRole", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// ClusterRole is a cluster level, logical grouping of PolicyRules that can be referenced as a unit
// by a RoleBinding or ClusterRoleBinding.
type ClusterRoleProperty struct {
	// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole.
	// If AggregationRule is set, then the Rules are controller managed and direct changes to Rules
	// will be stomped by the controller.
	AggregationRule *AggregationRule `pulumi:"aggregationRule"`

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

	// Standard object's metadata.
	Metadata *metaV1.ObjectMeta `pulumi:"metadata"`

	// Rules holds all the PolicyRules for this ClusterRole
	Rules []PolicyRule `pulumi:"rules"`

}

var _ClusterRolePropertyType = reflect.TypeOf((*ClusterRoleProperty)(nil)).Elem()

// ClusterRolePropertyInput represents an input type that resolves to a ClusterRoleProperty.
type ClusterRolePropertyInput interface {
	ElementType() reflect.Type

	ToClusterRolePropertyOutput() ClusterRolePropertyOutput
	ToClusterRolePropertyOutputWithContext(ctx context.Context) ClusterRolePropertyOutput
}

// ClusterRolePropertyArgs is a ClusterRolePropertyInput whose fields are all Input types.
type ClusterRolePropertyArgs struct {
	// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole.
	// If AggregationRule is set, then the Rules are controller managed and direct changes to Rules
	// will be stomped by the controller.
	AggregationRule AggregationRuleInput `pulumi:"aggregationRule"`

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

	// Standard object's metadata.
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// Rules holds all the PolicyRules for this ClusterRole
	Rules PolicyRuleArrayInput `pulumi:"rules"`

}

func (a ClusterRolePropertyArgs) ElementType() reflect.Type {
	return _ClusterRolePropertyType
}

func (a ClusterRolePropertyArgs) ToClusterRolePropertyOutput() ClusterRolePropertyOutput {
	return pulumi.ToOutput(a).(ClusterRolePropertyOutput)
}

func (a ClusterRolePropertyArgs) ToClusterRolePropertyOutputWithContext(ctx context.Context) ClusterRolePropertyOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ClusterRolePropertyOutput)
}

// ClusterRolePropertyOutput is an output type that resolves to a Input.
type ClusterRolePropertyOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ClusterRolePropertyOutput{}) }

func (ClusterRolePropertyOutput) ElementType() reflect.Type {
	return _ClusterRolePropertyType
}

func (o ClusterRolePropertyOutput) AggregationRule() AggregationRuleOutput {
	return o.Apply(func(v ClusterRoleProperty) *AggregationRule {
		return v.AggregationRule
	}).(AggregationRuleOutput)
}

func (o ClusterRolePropertyOutput) ApiVersion() pulumi.StringOutput {
	return o.Apply(func(v ClusterRoleProperty) *string {
		return v.ApiVersion
	}).(pulumi.StringOutput)
}

func (o ClusterRolePropertyOutput) Kind() pulumi.StringOutput {
	return o.Apply(func(v ClusterRoleProperty) *string {
		return v.Kind
	}).(pulumi.StringOutput)
}

func (o ClusterRolePropertyOutput) Metadata() metaV1.ObjectMetaOutput {
	return o.Apply(func(v ClusterRoleProperty) *metaV1.ObjectMeta {
		return v.Metadata
	}).(metaV1.ObjectMetaOutput)
}

func (o ClusterRolePropertyOutput) Rules() PolicyRuleArrayOutput {
	return o.Apply(func(v ClusterRoleProperty) []PolicyRule {
		return v.Rules
	}).(PolicyRuleArrayOutput)
}

var _ClusterRolePropertyArrayType = reflect.TypeOf((*[]ClusterRoleProperty)(nil)).Elem()

type ClusterRolePropertyArrayInput interface {
	ElementType() reflect.Type

	ToClusterRolePropertyArrayOutput() ClusterRolePropertyArrayOutput
	ToClusterRolePropertyArrayOutputWithContext(ctx context.Context) ClusterRolePropertyArrayOutput
}

type ClusterRolePropertyArray []ClusterRolePropertyInput

func (a ClusterRolePropertyArray) ElementType() reflect.Type {
	return _ClusterRolePropertyArrayType
}

func (a ClusterRolePropertyArray) ToClusterRolePropertyArrayOutput() ClusterRolePropertyArrayOutput {
	return pulumi.ToOutput(a).(ClusterRolePropertyArrayOutput)
}

func (a ClusterRolePropertyArray) ToClusterRolePropertyArrayOutputWithContext(ctx context.Context) ClusterRolePropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ClusterRolePropertyArrayOutput)
}

type ClusterRolePropertyArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(ClusterRolePropertyArrayOutput{}) }

func (ClusterRolePropertyArrayOutput) ElementType() reflect.Type {
	return _ClusterRolePropertyArrayType
}

func (o ClusterRolePropertyArrayOutput) Index(i pulumi.IntInput) ClusterRolePropertyOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) ClusterRoleProperty {
		return vs[0].([]ClusterRoleProperty)[vs[1].(int)]
	}).(ClusterRolePropertyOutput)
}
