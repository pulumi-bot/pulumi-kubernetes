# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from ... import _utilities, _tables

@pulumi.output_type
class TokenReviewSpec(dict):
    """
    TokenReviewSpec is a description of the token authentication request.
    """
    audiences: Optional[List[str]] = pulumi.output_property("audiences")
    """
    Audiences is a list of the identifiers that the resource server presented with the token identifies as. Audience-aware token authenticators will verify that the token was intended for at least one of the audiences in this list. If no audiences are provided, the audience will default to the audience of the Kubernetes apiserver.
    """
    token: Optional[str] = pulumi.output_property("token")
    """
    Token is the opaque bearer token.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TokenReviewStatus(dict):
    """
    TokenReviewStatus is the result of the token authentication request.
    """
    audiences: Optional[List[str]] = pulumi.output_property("audiences")
    """
    Audiences are audience identifiers chosen by the authenticator that are compatible with both the TokenReview and token. An identifier is any identifier in the intersection of the TokenReviewSpec audiences and the token's audiences. A client of the TokenReview API that sets the spec.audiences field should validate that a compatible audience identifier is returned in the status.audiences field to ensure that the TokenReview server is audience aware. If a TokenReview returns an empty status.audience field where status.authenticated is "true", the token is valid against the audience of the Kubernetes API server.
    """
    authenticated: Optional[bool] = pulumi.output_property("authenticated")
    """
    Authenticated indicates that the token was associated with a known user.
    """
    error: Optional[str] = pulumi.output_property("error")
    """
    Error indicates that the token couldn't be checked
    """
    user: Optional['outputs.UserInfo'] = pulumi.output_property("user")
    """
    User is the UserInfo associated with the provided token.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class UserInfo(dict):
    """
    UserInfo holds the information about the user needed to implement the user.Info interface.
    """
    extra: Optional[Dict[str, List[str]]] = pulumi.output_property("extra")
    """
    Any additional information provided by the authenticator.
    """
    groups: Optional[List[str]] = pulumi.output_property("groups")
    """
    The names of groups this user is a part of.
    """
    uid: Optional[str] = pulumi.output_property("uid")
    """
    A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.
    """
    username: Optional[str] = pulumi.output_property("username")
    """
    The name that uniquely identifies this user among all active users.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


