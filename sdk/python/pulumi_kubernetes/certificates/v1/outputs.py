# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities, _tables
from . import outputs
from ... import meta as _meta

__all__ = [
    'CertificateSigningRequest',
    'CertificateSigningRequestCondition',
    'CertificateSigningRequestSpec',
    'CertificateSigningRequestStatus',
]

@pulumi.output_type
class CertificateSigningRequest(dict):
    """
    CertificateSigningRequest objects provide a mechanism to obtain x509 certificates by submitting a certificate signing request, and having it asynchronously approved and issued.

    Kubelets use this API to obtain:
     1. client certificates to authenticate to kube-apiserver (with the "kubernetes.io/kube-apiserver-client-kubelet" signerName).
     2. serving certificates for TLS endpoints kube-apiserver can connect to securely (with the "kubernetes.io/kubelet-serving" signerName).

    This API can be used to request client certificates to authenticate to kube-apiserver (with the "kubernetes.io/kube-apiserver-client" signerName), or to obtain certificates from custom non-Kubernetes signers.
    """
    def __init__(__self__, *,
                 spec: 'outputs.CertificateSigningRequestSpec',
                 api_version: Optional[str] = None,
                 kind: Optional[str] = None,
                 metadata: Optional['_meta.v1.outputs.ObjectMeta'] = None,
                 status: Optional['outputs.CertificateSigningRequestStatus'] = None):
        """
        CertificateSigningRequest objects provide a mechanism to obtain x509 certificates by submitting a certificate signing request, and having it asynchronously approved and issued.

        Kubelets use this API to obtain:
         1. client certificates to authenticate to kube-apiserver (with the "kubernetes.io/kube-apiserver-client-kubelet" signerName).
         2. serving certificates for TLS endpoints kube-apiserver can connect to securely (with the "kubernetes.io/kubelet-serving" signerName).

        This API can be used to request client certificates to authenticate to kube-apiserver (with the "kubernetes.io/kube-apiserver-client" signerName), or to obtain certificates from custom non-Kubernetes signers.
        :param 'CertificateSigningRequestSpecArgs' spec: spec contains the certificate request, and is immutable after creation. Only the request, signerName, and usages fields can be set on creation. Other fields are derived by Kubernetes and cannot be modified by users.
        :param str api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param str kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param 'CertificateSigningRequestStatusArgs' status: status contains information about whether the request is approved or denied, and the certificate issued by the signer, or the failure condition indicating signer failure.
        """
        pulumi.set(__self__, "spec", spec)
        if api_version is not None:
            pulumi.set(__self__, "api_version", 'certificates.k8s.io/v1')
        if kind is not None:
            pulumi.set(__self__, "kind", 'CertificateSigningRequest')
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def spec(self) -> 'outputs.CertificateSigningRequestSpec':
        """
        spec contains the certificate request, and is immutable after creation. Only the request, signerName, and usages fields can be set on creation. Other fields are derived by Kubernetes and cannot be modified by users.
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[str]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def metadata(self) -> Optional['_meta.v1.outputs.ObjectMeta']:
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def status(self) -> Optional['outputs.CertificateSigningRequestStatus']:
        """
        status contains information about whether the request is approved or denied, and the certificate issued by the signer, or the failure condition indicating signer failure.
        """
        return pulumi.get(self, "status")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CertificateSigningRequestCondition(dict):
    """
    CertificateSigningRequestCondition describes a condition of a CertificateSigningRequest object
    """
    def __init__(__self__, *,
                 status: str,
                 type: str,
                 last_transition_time: Optional[str] = None,
                 last_update_time: Optional[str] = None,
                 message: Optional[str] = None,
                 reason: Optional[str] = None):
        """
        CertificateSigningRequestCondition describes a condition of a CertificateSigningRequest object
        :param str status: status of the condition, one of True, False, Unknown. Approved, Denied, and Failed conditions may not be "False" or "Unknown".
        :param str type: type of the condition. Known conditions are "Approved", "Denied", and "Failed".
               
               An "Approved" condition is added via the /approval subresource, indicating the request was approved and should be issued by the signer.
               
               A "Denied" condition is added via the /approval subresource, indicating the request was denied and should not be issued by the signer.
               
               A "Failed" condition is added via the /status subresource, indicating the signer failed to issue the certificate.
               
               Approved and Denied conditions are mutually exclusive. Approved, Denied, and Failed conditions cannot be removed once added.
               
               Only one condition of a given type is allowed.
        :param str last_transition_time: lastTransitionTime is the time the condition last transitioned from one status to another. If unset, when a new condition type is added or an existing condition's status is changed, the server defaults this to the current time.
        :param str last_update_time: lastUpdateTime is the time of the last update to this condition
        :param str message: message contains a human readable message with details about the request state
        :param str reason: reason indicates a brief reason for the request state
        """
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "type", type)
        if last_transition_time is not None:
            pulumi.set(__self__, "last_transition_time", last_transition_time)
        if last_update_time is not None:
            pulumi.set(__self__, "last_update_time", last_update_time)
        if message is not None:
            pulumi.set(__self__, "message", message)
        if reason is not None:
            pulumi.set(__self__, "reason", reason)

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        status of the condition, one of True, False, Unknown. Approved, Denied, and Failed conditions may not be "False" or "Unknown".
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        type of the condition. Known conditions are "Approved", "Denied", and "Failed".

        An "Approved" condition is added via the /approval subresource, indicating the request was approved and should be issued by the signer.

        A "Denied" condition is added via the /approval subresource, indicating the request was denied and should not be issued by the signer.

        A "Failed" condition is added via the /status subresource, indicating the signer failed to issue the certificate.

        Approved and Denied conditions are mutually exclusive. Approved, Denied, and Failed conditions cannot be removed once added.

        Only one condition of a given type is allowed.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="lastTransitionTime")
    def last_transition_time(self) -> Optional[str]:
        """
        lastTransitionTime is the time the condition last transitioned from one status to another. If unset, when a new condition type is added or an existing condition's status is changed, the server defaults this to the current time.
        """
        return pulumi.get(self, "last_transition_time")

    @property
    @pulumi.getter(name="lastUpdateTime")
    def last_update_time(self) -> Optional[str]:
        """
        lastUpdateTime is the time of the last update to this condition
        """
        return pulumi.get(self, "last_update_time")

    @property
    @pulumi.getter
    def message(self) -> Optional[str]:
        """
        message contains a human readable message with details about the request state
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter
    def reason(self) -> Optional[str]:
        """
        reason indicates a brief reason for the request state
        """
        return pulumi.get(self, "reason")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CertificateSigningRequestSpec(dict):
    """
    CertificateSigningRequestSpec contains the certificate request.
    """
    def __init__(__self__, *,
                 request: str,
                 signer_name: str,
                 extra: Optional[Mapping[str, Sequence[str]]] = None,
                 groups: Optional[Sequence[str]] = None,
                 uid: Optional[str] = None,
                 usages: Optional[Sequence[str]] = None,
                 username: Optional[str] = None):
        """
        CertificateSigningRequestSpec contains the certificate request.
        :param str request: request contains an x509 certificate signing request encoded in a "CERTIFICATE REQUEST" PEM block. When serialized as JSON or YAML, the data is additionally base64-encoded.
        :param str signer_name: signerName indicates the requested signer, and is a qualified name.
               
               List/watch requests for CertificateSigningRequests can filter on this field using a "spec.signerName=NAME" fieldSelector.
               
               Well-known Kubernetes signers are:
                1. "kubernetes.io/kube-apiserver-client": issues client certificates that can be used to authenticate to kube-apiserver.
                 Requests for this signer are never auto-approved by kube-controller-manager, can be issued by the "csrsigning" controller in kube-controller-manager.
                2. "kubernetes.io/kube-apiserver-client-kubelet": issues client certificates that kubelets use to authenticate to kube-apiserver.
                 Requests for this signer can be auto-approved by the "csrapproving" controller in kube-controller-manager, and can be issued by the "csrsigning" controller in kube-controller-manager.
                3. "kubernetes.io/kubelet-serving" issues serving certificates that kubelets use to serve TLS endpoints, which kube-apiserver can connect to securely.
                 Requests for this signer are never auto-approved by kube-controller-manager, and can be issued by the "csrsigning" controller in kube-controller-manager.
               
               More details are available at https://k8s.io/docs/reference/access-authn-authz/certificate-signing-requests/#kubernetes-signers
               
               Custom signerNames can also be specified. The signer defines:
                1. Trust distribution: how trust (CA bundles) are distributed.
                2. Permitted subjects: and behavior when a disallowed subject is requested.
                3. Required, permitted, or forbidden x509 extensions in the request (including whether subjectAltNames are allowed, which types, restrictions on allowed values) and behavior when a disallowed extension is requested.
                4. Required, permitted, or forbidden key usages / extended key usages.
                5. Expiration/certificate lifetime: whether it is fixed by the signer, configurable by the admin.
                6. Whether or not requests for CA certificates are allowed.
        :param Mapping[str, Sequence[str]] extra: extra contains extra attributes of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
        :param Sequence[str] groups: groups contains group membership of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
        :param str uid: uid contains the uid of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
        :param Sequence[str] usages: usages specifies a set of key usages requested in the issued certificate.
               
               Requests for TLS client certificates typically request: "digital signature", "key encipherment", "client auth".
               
               Requests for TLS serving certificates typically request: "key encipherment", "digital signature", "server auth".
               
               Valid values are:
                "signing", "digital signature", "content commitment",
                "key encipherment", "key agreement", "data encipherment",
                "cert sign", "crl sign", "encipher only", "decipher only", "any",
                "server auth", "client auth",
                "code signing", "email protection", "s/mime",
                "ipsec end system", "ipsec tunnel", "ipsec user",
                "timestamping", "ocsp signing", "microsoft sgc", "netscape sgc"
        :param str username: username contains the name of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
        """
        pulumi.set(__self__, "request", request)
        pulumi.set(__self__, "signer_name", signer_name)
        if extra is not None:
            pulumi.set(__self__, "extra", extra)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)
        if uid is not None:
            pulumi.set(__self__, "uid", uid)
        if usages is not None:
            pulumi.set(__self__, "usages", usages)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def request(self) -> str:
        """
        request contains an x509 certificate signing request encoded in a "CERTIFICATE REQUEST" PEM block. When serialized as JSON or YAML, the data is additionally base64-encoded.
        """
        return pulumi.get(self, "request")

    @property
    @pulumi.getter(name="signerName")
    def signer_name(self) -> str:
        """
        signerName indicates the requested signer, and is a qualified name.

        List/watch requests for CertificateSigningRequests can filter on this field using a "spec.signerName=NAME" fieldSelector.

        Well-known Kubernetes signers are:
         1. "kubernetes.io/kube-apiserver-client": issues client certificates that can be used to authenticate to kube-apiserver.
          Requests for this signer are never auto-approved by kube-controller-manager, can be issued by the "csrsigning" controller in kube-controller-manager.
         2. "kubernetes.io/kube-apiserver-client-kubelet": issues client certificates that kubelets use to authenticate to kube-apiserver.
          Requests for this signer can be auto-approved by the "csrapproving" controller in kube-controller-manager, and can be issued by the "csrsigning" controller in kube-controller-manager.
         3. "kubernetes.io/kubelet-serving" issues serving certificates that kubelets use to serve TLS endpoints, which kube-apiserver can connect to securely.
          Requests for this signer are never auto-approved by kube-controller-manager, and can be issued by the "csrsigning" controller in kube-controller-manager.

        More details are available at https://k8s.io/docs/reference/access-authn-authz/certificate-signing-requests/#kubernetes-signers

        Custom signerNames can also be specified. The signer defines:
         1. Trust distribution: how trust (CA bundles) are distributed.
         2. Permitted subjects: and behavior when a disallowed subject is requested.
         3. Required, permitted, or forbidden x509 extensions in the request (including whether subjectAltNames are allowed, which types, restrictions on allowed values) and behavior when a disallowed extension is requested.
         4. Required, permitted, or forbidden key usages / extended key usages.
         5. Expiration/certificate lifetime: whether it is fixed by the signer, configurable by the admin.
         6. Whether or not requests for CA certificates are allowed.
        """
        return pulumi.get(self, "signer_name")

    @property
    @pulumi.getter
    def extra(self) -> Optional[Mapping[str, Sequence[str]]]:
        """
        extra contains extra attributes of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
        """
        return pulumi.get(self, "extra")

    @property
    @pulumi.getter
    def groups(self) -> Optional[Sequence[str]]:
        """
        groups contains group membership of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter
    def uid(self) -> Optional[str]:
        """
        uid contains the uid of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter
    def usages(self) -> Optional[Sequence[str]]:
        """
        usages specifies a set of key usages requested in the issued certificate.

        Requests for TLS client certificates typically request: "digital signature", "key encipherment", "client auth".

        Requests for TLS serving certificates typically request: "key encipherment", "digital signature", "server auth".

        Valid values are:
         "signing", "digital signature", "content commitment",
         "key encipherment", "key agreement", "data encipherment",
         "cert sign", "crl sign", "encipher only", "decipher only", "any",
         "server auth", "client auth",
         "code signing", "email protection", "s/mime",
         "ipsec end system", "ipsec tunnel", "ipsec user",
         "timestamping", "ocsp signing", "microsoft sgc", "netscape sgc"
        """
        return pulumi.get(self, "usages")

    @property
    @pulumi.getter
    def username(self) -> Optional[str]:
        """
        username contains the name of the user that created the CertificateSigningRequest. Populated by the API server on creation and immutable.
        """
        return pulumi.get(self, "username")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CertificateSigningRequestStatus(dict):
    """
    CertificateSigningRequestStatus contains conditions used to indicate approved/denied/failed status of the request, and the issued certificate.
    """
    def __init__(__self__, *,
                 certificate: Optional[str] = None,
                 conditions: Optional[Sequence['outputs.CertificateSigningRequestCondition']] = None):
        """
        CertificateSigningRequestStatus contains conditions used to indicate approved/denied/failed status of the request, and the issued certificate.
        :param str certificate: certificate is populated with an issued certificate by the signer after an Approved condition is present. This field is set via the /status subresource. Once populated, this field is immutable.
               
               If the certificate signing request is denied, a condition of type "Denied" is added and this field remains empty. If the signer cannot issue the certificate, a condition of type "Failed" is added and this field remains empty.
               
               Validation requirements:
                1. certificate must contain one or more PEM blocks.
                2. All PEM blocks must have the "CERTIFICATE" label, contain no headers, and the encoded data
                 must be a BER-encoded ASN.1 Certificate structure as described in section 4 of RFC5280.
                3. Non-PEM content may appear before or after the "CERTIFICATE" PEM blocks and is unvalidated,
                 to allow for explanatory text as described in section 5.2 of RFC7468.
               
               If more than one PEM block is present, and the definition of the requested spec.signerName does not indicate otherwise, the first block is the issued certificate, and subsequent blocks should be treated as intermediate certificates and presented in TLS handshakes.
               
               The certificate is encoded in PEM format.
               
               When serialized as JSON or YAML, the data is additionally base64-encoded, so it consists of:
               
                   base64(
                   -----BEGIN CERTIFICATE-----
                   ...
                   -----END CERTIFICATE-----
                   )
        :param Sequence['CertificateSigningRequestConditionArgs'] conditions: conditions applied to the request. Known conditions are "Approved", "Denied", and "Failed".
        """
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[str]:
        """
        certificate is populated with an issued certificate by the signer after an Approved condition is present. This field is set via the /status subresource. Once populated, this field is immutable.

        If the certificate signing request is denied, a condition of type "Denied" is added and this field remains empty. If the signer cannot issue the certificate, a condition of type "Failed" is added and this field remains empty.

        Validation requirements:
         1. certificate must contain one or more PEM blocks.
         2. All PEM blocks must have the "CERTIFICATE" label, contain no headers, and the encoded data
          must be a BER-encoded ASN.1 Certificate structure as described in section 4 of RFC5280.
         3. Non-PEM content may appear before or after the "CERTIFICATE" PEM blocks and is unvalidated,
          to allow for explanatory text as described in section 5.2 of RFC7468.

        If more than one PEM block is present, and the definition of the requested spec.signerName does not indicate otherwise, the first block is the issued certificate, and subsequent blocks should be treated as intermediate certificates and presented in TLS handshakes.

        The certificate is encoded in PEM format.

        When serialized as JSON or YAML, the data is additionally base64-encoded, so it consists of:

            base64(
            -----BEGIN CERTIFICATE-----
            ...
            -----END CERTIFICATE-----
            )
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter
    def conditions(self) -> Optional[Sequence['outputs.CertificateSigningRequestCondition']]:
        """
        conditions applied to the request. Known conditions are "Approved", "Denied", and "Failed".
        """
        return pulumi.get(self, "conditions")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


