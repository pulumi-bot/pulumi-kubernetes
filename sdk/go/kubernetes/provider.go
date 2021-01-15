// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The provider type for the kubernetes package.
type Provider struct {
	pulumi.ProviderResourceState
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if args.EnableDryRun == nil {
		args.EnableDryRun = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "PULUMI_K8S_ENABLE_DRY_RUN").(bool))
	}
	if args.Kubeconfig == nil {
		args.Kubeconfig = pulumi.StringPtr(getEnvOrDefault("", nil, "KUBECONFIG").(string))
	}
	if args.SuppressDeprecationWarnings == nil {
		args.SuppressDeprecationWarnings = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "PULUMI_K8S_SUPPRESS_DEPRECATION_WARNINGS").(bool))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:kubernetes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// If present, the name of the kubeconfig cluster to use.
	Cluster *string `pulumi:"cluster"`
	// If present, the name of the kubeconfig context to use.
	Context *string `pulumi:"context"`
	// BETA FEATURE - If present and set to true, enable server-side diff calculations.
	// This feature is in developer preview, and is disabled by default.
	//
	// This config can be specified in the following ways, using this precedence:
	// 1. This `enableDryRun` parameter.
	// 2. The `PULUMI_K8S_ENABLE_DRY_RUN` environment variable.
	EnableDryRun *bool `pulumi:"enableDryRun"`
	// The contents of a kubeconfig file or the path to a kubeconfig file. If this is set, this config will be used instead of $KUBECONFIG.
	Kubeconfig *string `pulumi:"kubeconfig"`
	// If present, the default namespace to use. This flag is ignored for cluster-scoped resources.
	//
	// A namespace can be specified in multiple places, and the precedence is as follows:
	// 1. `.metadata.namespace` set on the resource.
	// 2. This `namespace` parameter.
	// 3. `namespace` set for the active context in the kubeconfig.
	Namespace *string `pulumi:"namespace"`
	// BETA FEATURE - If present, render resource manifests to this directory. In this mode, resources will not
	// be created on a Kubernetes cluster, but the rendered manifests will be kept in sync with changes
	// to the Pulumi program. This feature is in developer preview, and is disabled by default.
	//
	// Note that some computed Outputs such as status fields will not be populated
	// since the resources are not created on a Kubernetes cluster. These Output values will remain undefined,
	// and may result in an error if they are referenced by other resources. Also note that any secret values
	// used in these resources will be rendered in plaintext to the resulting YAML.
	RenderYamlToDirectory *string `pulumi:"renderYamlToDirectory"`
	// If present and set to true, suppress apiVersion deprecation warnings from the CLI.
	//
	// This config can be specified in the following ways, using this precedence:
	// 1. This `suppressDeprecationWarnings` parameter.
	// 2. The `PULUMI_K8S_SUPPRESS_DEPRECATION_WARNINGS` environment variable.
	SuppressDeprecationWarnings *bool `pulumi:"suppressDeprecationWarnings"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// If present, the name of the kubeconfig cluster to use.
	Cluster pulumi.StringPtrInput
	// If present, the name of the kubeconfig context to use.
	Context pulumi.StringPtrInput
	// BETA FEATURE - If present and set to true, enable server-side diff calculations.
	// This feature is in developer preview, and is disabled by default.
	//
	// This config can be specified in the following ways, using this precedence:
	// 1. This `enableDryRun` parameter.
	// 2. The `PULUMI_K8S_ENABLE_DRY_RUN` environment variable.
	EnableDryRun pulumi.BoolPtrInput
	// The contents of a kubeconfig file or the path to a kubeconfig file. If this is set, this config will be used instead of $KUBECONFIG.
	Kubeconfig pulumi.StringPtrInput
	// If present, the default namespace to use. This flag is ignored for cluster-scoped resources.
	//
	// A namespace can be specified in multiple places, and the precedence is as follows:
	// 1. `.metadata.namespace` set on the resource.
	// 2. This `namespace` parameter.
	// 3. `namespace` set for the active context in the kubeconfig.
	Namespace pulumi.StringPtrInput
	// BETA FEATURE - If present, render resource manifests to this directory. In this mode, resources will not
	// be created on a Kubernetes cluster, but the rendered manifests will be kept in sync with changes
	// to the Pulumi program. This feature is in developer preview, and is disabled by default.
	//
	// Note that some computed Outputs such as status fields will not be populated
	// since the resources are not created on a Kubernetes cluster. These Output values will remain undefined,
	// and may result in an error if they are referenced by other resources. Also note that any secret values
	// used in these resources will be rendered in plaintext to the resulting YAML.
	RenderYamlToDirectory pulumi.StringPtrInput
	// If present and set to true, suppress apiVersion deprecation warnings from the CLI.
	//
	// This config can be specified in the following ways, using this precedence:
	// 1. This `suppressDeprecationWarnings` parameter.
	// 2. The `PULUMI_K8S_SUPPRESS_DEPRECATION_WARNINGS` environment variable.
	SuppressDeprecationWarnings pulumi.BoolPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil))
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

func (i *Provider) ToProviderPtrOutput() ProviderPtrOutput {
	return i.ToProviderPtrOutputWithContext(context.Background())
}

func (i *Provider) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderPtrOutput)
}

type ProviderPtrInput interface {
	pulumi.Input

	ToProviderPtrOutput() ProviderPtrOutput
	ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput
}

type providerPtrType ProviderArgs

func (*providerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil))
}

func (i *providerPtrType) ToProviderPtrOutput() ProviderPtrOutput {
	return i.ToProviderPtrOutputWithContext(context.Background())
}

func (i *providerPtrType) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderPtrOutput)
}

type ProviderOutput struct {
	*pulumi.OutputState
}

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil))
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderPtrOutput() ProviderPtrOutput {
	return o.ToProviderPtrOutputWithContext(context.Background())
}

func (o ProviderOutput) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return o.ApplyT(func(v Provider) *Provider {
		return &v
	}).(ProviderPtrOutput)
}

type ProviderPtrOutput struct {
	*pulumi.OutputState
}

func (ProviderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil))
}

func (o ProviderPtrOutput) ToProviderPtrOutput() ProviderPtrOutput {
	return o
}

func (o ProviderPtrOutput) ToProviderPtrOutputWithContext(ctx context.Context) ProviderPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProviderOutput{})
	pulumi.RegisterOutputType(ProviderPtrOutput{})
}
