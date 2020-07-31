# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'BoundObjectReferenceArgs',
    'TokenRequestSpecArgs',
    'TokenReviewSpecArgs',
]

@pulumi.input_type
class BoundObjectReferenceArgs:
    api_version: Optional[pulumi.Input[str]] = pulumi.input_property("apiVersion")
    """
    API version of the referent.
    """
    kind: Optional[pulumi.Input[str]] = pulumi.input_property("kind")
    """
    Kind of the referent. Valid kinds are 'Pod' and 'Secret'.
    """
    name: Optional[pulumi.Input[str]] = pulumi.input_property("name")
    """
    Name of the referent.
    """
    uid: Optional[pulumi.Input[str]] = pulumi.input_property("uid")
    """
    UID of the referent.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, api_version: Optional[pulumi.Input[str]] = None, kind: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, uid: Optional[pulumi.Input[str]] = None) -> None:
        """
        BoundObjectReference is a reference to an object that a token is bound to.
        :param pulumi.Input[str] api_version: API version of the referent.
        :param pulumi.Input[str] kind: Kind of the referent. Valid kinds are 'Pod' and 'Secret'.
        :param pulumi.Input[str] name: Name of the referent.
        :param pulumi.Input[str] uid: UID of the referent.
        """
        __self__.api_version = api_version
        __self__.kind = kind
        __self__.name = name
        __self__.uid = uid

@pulumi.input_type
class TokenRequestSpecArgs:
    audiences: pulumi.Input[List[pulumi.Input[str]]] = pulumi.input_property("audiences")
    """
    Audiences are the intendend audiences of the token. A recipient of a token must identitfy themself with an identifier in the list of audiences of the token, and otherwise should reject the token. A token issued for multiple audiences may be used to authenticate against any of the audiences listed but implies a high degree of trust between the target audiences.
    """
    bound_object_ref: Optional[pulumi.Input['BoundObjectReferenceArgs']] = pulumi.input_property("boundObjectRef")
    """
    BoundObjectRef is a reference to an object that the token will be bound to. The token will only be valid for as long as the bound object exists. NOTE: The API server's TokenReview endpoint will validate the BoundObjectRef, but other audiences may not. Keep ExpirationSeconds small if you want prompt revocation.
    """
    expiration_seconds: Optional[pulumi.Input[float]] = pulumi.input_property("expirationSeconds")
    """
    ExpirationSeconds is the requested duration of validity of the request. The token issuer may return a token with a different validity duration so a client needs to check the 'expiration' field in a response.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, audiences: pulumi.Input[List[pulumi.Input[str]]], bound_object_ref: Optional[pulumi.Input['BoundObjectReferenceArgs']] = None, expiration_seconds: Optional[pulumi.Input[float]] = None) -> None:
        """
        TokenRequestSpec contains client provided parameters of a token request.
        :param pulumi.Input[List[pulumi.Input[str]]] audiences: Audiences are the intendend audiences of the token. A recipient of a token must identitfy themself with an identifier in the list of audiences of the token, and otherwise should reject the token. A token issued for multiple audiences may be used to authenticate against any of the audiences listed but implies a high degree of trust between the target audiences.
        :param pulumi.Input['BoundObjectReferenceArgs'] bound_object_ref: BoundObjectRef is a reference to an object that the token will be bound to. The token will only be valid for as long as the bound object exists. NOTE: The API server's TokenReview endpoint will validate the BoundObjectRef, but other audiences may not. Keep ExpirationSeconds small if you want prompt revocation.
        :param pulumi.Input[float] expiration_seconds: ExpirationSeconds is the requested duration of validity of the request. The token issuer may return a token with a different validity duration so a client needs to check the 'expiration' field in a response.
        """
        __self__.audiences = audiences
        __self__.bound_object_ref = bound_object_ref
        __self__.expiration_seconds = expiration_seconds

@pulumi.input_type
class TokenReviewSpecArgs:
    audiences: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("audiences")
    """
    Audiences is a list of the identifiers that the resource server presented with the token identifies as. Audience-aware token authenticators will verify that the token was intended for at least one of the audiences in this list. If no audiences are provided, the audience will default to the audience of the Kubernetes apiserver.
    """
    token: Optional[pulumi.Input[str]] = pulumi.input_property("token")
    """
    Token is the opaque bearer token.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, audiences: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, token: Optional[pulumi.Input[str]] = None) -> None:
        """
        TokenReviewSpec is a description of the token authentication request.
        :param pulumi.Input[List[pulumi.Input[str]]] audiences: Audiences is a list of the identifiers that the resource server presented with the token identifies as. Audience-aware token authenticators will verify that the token was intended for at least one of the audiences in this list. If no audiences are provided, the audience will default to the audience of the Kubernetes apiserver.
        :param pulumi.Input[str] token: Token is the opaque bearer token.
        """
        __self__.audiences = audiences
        __self__.token = token

