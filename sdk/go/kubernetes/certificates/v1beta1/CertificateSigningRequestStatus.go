// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// 
type CertificateSigningRequestStatus struct {
	// If request was approved, the controller will place the issued certificate here.
	Certificate *string `pulumi:"certificate"`

	// Conditions applied to the request, such as approval or denial.
	Conditions []CertificateSigningRequestCondition `pulumi:"conditions"`

}

var _CertificateSigningRequestStatusType = reflect.TypeOf((*CertificateSigningRequestStatus)(nil)).Elem()

// CertificateSigningRequestStatusInput represents an input type that resolves to a CertificateSigningRequestStatus.
type CertificateSigningRequestStatusInput interface {
	ElementType() reflect.Type

	ToCertificateSigningRequestStatusOutput() CertificateSigningRequestStatusOutput
	ToCertificateSigningRequestStatusOutputWithContext(ctx context.Context) CertificateSigningRequestStatusOutput
}

// CertificateSigningRequestStatusArgs is a CertificateSigningRequestStatusInput whose fields are all Input types.
type CertificateSigningRequestStatusArgs struct {
	// If request was approved, the controller will place the issued certificate here.
	Certificate pulumi.StringInput `pulumi:"certificate"`

	// Conditions applied to the request, such as approval or denial.
	Conditions CertificateSigningRequestConditionArrayInput `pulumi:"conditions"`

}

func (a CertificateSigningRequestStatusArgs) ElementType() reflect.Type {
	return _CertificateSigningRequestStatusType
}

func (a CertificateSigningRequestStatusArgs) ToCertificateSigningRequestStatusOutput() CertificateSigningRequestStatusOutput {
	return pulumi.ToOutput(a).(CertificateSigningRequestStatusOutput)
}

func (a CertificateSigningRequestStatusArgs) ToCertificateSigningRequestStatusOutputWithContext(ctx context.Context) CertificateSigningRequestStatusOutput {
	return pulumi.ToOutputWithContext(ctx, a).(CertificateSigningRequestStatusOutput)
}

// CertificateSigningRequestStatusOutput is an output type that resolves to a Input.
type CertificateSigningRequestStatusOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(CertificateSigningRequestStatusOutput{}) }

func (CertificateSigningRequestStatusOutput) ElementType() reflect.Type {
	return _CertificateSigningRequestStatusType
}

func (o CertificateSigningRequestStatusOutput) Certificate() pulumi.StringOutput {
	return o.Apply(func(v CertificateSigningRequestStatus) *string {
		return v.Certificate
	}).(pulumi.StringOutput)
}

func (o CertificateSigningRequestStatusOutput) Conditions() CertificateSigningRequestConditionArrayOutput {
	return o.Apply(func(v CertificateSigningRequestStatus) []CertificateSigningRequestCondition {
		return v.Conditions
	}).(CertificateSigningRequestConditionArrayOutput)
}

