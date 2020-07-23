# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from ... import _utilities, _tables

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

