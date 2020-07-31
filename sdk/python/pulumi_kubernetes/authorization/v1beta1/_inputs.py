# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'NonResourceAttributesArgs',
    'ResourceAttributesArgs',
    'SelfSubjectAccessReviewSpecArgs',
    'SelfSubjectRulesReviewSpecArgs',
    'SubjectAccessReviewSpecArgs',
]

@pulumi.input_type
class NonResourceAttributesArgs:
    path: Optional[pulumi.Input[str]] = pulumi.input_property("path")
    """
    Path is the URL path of the request
    """
    verb: Optional[pulumi.Input[str]] = pulumi.input_property("verb")
    """
    Verb is the standard HTTP verb
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, path: Optional[pulumi.Input[str]] = None, verb: Optional[pulumi.Input[str]] = None) -> None:
        """
        NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface
        :param pulumi.Input[str] path: Path is the URL path of the request
        :param pulumi.Input[str] verb: Verb is the standard HTTP verb
        """
        __self__.path = path
        __self__.verb = verb

@pulumi.input_type
class ResourceAttributesArgs:
    group: Optional[pulumi.Input[str]] = pulumi.input_property("group")
    """
    Group is the API Group of the Resource.  "*" means all.
    """
    name: Optional[pulumi.Input[str]] = pulumi.input_property("name")
    """
    Name is the name of the resource being requested for a "get" or deleted for a "delete". "" (empty) means all.
    """
    namespace: Optional[pulumi.Input[str]] = pulumi.input_property("namespace")
    """
    Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces "" (empty) is defaulted for LocalSubjectAccessReviews "" (empty) is empty for cluster-scoped resources "" (empty) means "all" for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview
    """
    resource: Optional[pulumi.Input[str]] = pulumi.input_property("resource")
    """
    Resource is one of the existing resource types.  "*" means all.
    """
    subresource: Optional[pulumi.Input[str]] = pulumi.input_property("subresource")
    """
    Subresource is one of the existing resource types.  "" means none.
    """
    verb: Optional[pulumi.Input[str]] = pulumi.input_property("verb")
    """
    Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  "*" means all.
    """
    version: Optional[pulumi.Input[str]] = pulumi.input_property("version")
    """
    Version is the API Version of the Resource.  "*" means all.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, group: Optional[pulumi.Input[str]] = None, name: Optional[pulumi.Input[str]] = None, namespace: Optional[pulumi.Input[str]] = None, resource: Optional[pulumi.Input[str]] = None, subresource: Optional[pulumi.Input[str]] = None, verb: Optional[pulumi.Input[str]] = None, version: Optional[pulumi.Input[str]] = None) -> None:
        """
        ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface
        :param pulumi.Input[str] group: Group is the API Group of the Resource.  "*" means all.
        :param pulumi.Input[str] name: Name is the name of the resource being requested for a "get" or deleted for a "delete". "" (empty) means all.
        :param pulumi.Input[str] namespace: Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces "" (empty) is defaulted for LocalSubjectAccessReviews "" (empty) is empty for cluster-scoped resources "" (empty) means "all" for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview
        :param pulumi.Input[str] resource: Resource is one of the existing resource types.  "*" means all.
        :param pulumi.Input[str] subresource: Subresource is one of the existing resource types.  "" means none.
        :param pulumi.Input[str] verb: Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  "*" means all.
        :param pulumi.Input[str] version: Version is the API Version of the Resource.  "*" means all.
        """
        __self__.group = group
        __self__.name = name
        __self__.namespace = namespace
        __self__.resource = resource
        __self__.subresource = subresource
        __self__.verb = verb
        __self__.version = version

@pulumi.input_type
class SelfSubjectAccessReviewSpecArgs:
    non_resource_attributes: Optional[pulumi.Input['NonResourceAttributesArgs']] = pulumi.input_property("nonResourceAttributes")
    """
    NonResourceAttributes describes information for a non-resource access request
    """
    resource_attributes: Optional[pulumi.Input['ResourceAttributesArgs']] = pulumi.input_property("resourceAttributes")
    """
    ResourceAuthorizationAttributes describes information for a resource access request
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, non_resource_attributes: Optional[pulumi.Input['NonResourceAttributesArgs']] = None, resource_attributes: Optional[pulumi.Input['ResourceAttributesArgs']] = None) -> None:
        """
        SelfSubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
        :param pulumi.Input['NonResourceAttributesArgs'] non_resource_attributes: NonResourceAttributes describes information for a non-resource access request
        :param pulumi.Input['ResourceAttributesArgs'] resource_attributes: ResourceAuthorizationAttributes describes information for a resource access request
        """
        __self__.non_resource_attributes = non_resource_attributes
        __self__.resource_attributes = resource_attributes

@pulumi.input_type
class SelfSubjectRulesReviewSpecArgs:
    namespace: Optional[pulumi.Input[str]] = pulumi.input_property("namespace")
    """
    Namespace to evaluate rules for. Required.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, namespace: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] namespace: Namespace to evaluate rules for. Required.
        """
        __self__.namespace = namespace

@pulumi.input_type
class SubjectAccessReviewSpecArgs:
    extra: Optional[pulumi.Input[Dict[str, pulumi.Input[List[pulumi.Input[str]]]]]] = pulumi.input_property("extra")
    """
    Extra corresponds to the user.Info.GetExtra() method from the authenticator.  Since that is input to the authorizer it needs a reflection here.
    """
    group: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("group")
    """
    Groups is the groups you're testing for.
    """
    non_resource_attributes: Optional[pulumi.Input['NonResourceAttributesArgs']] = pulumi.input_property("nonResourceAttributes")
    """
    NonResourceAttributes describes information for a non-resource access request
    """
    resource_attributes: Optional[pulumi.Input['ResourceAttributesArgs']] = pulumi.input_property("resourceAttributes")
    """
    ResourceAuthorizationAttributes describes information for a resource access request
    """
    uid: Optional[pulumi.Input[str]] = pulumi.input_property("uid")
    """
    UID information about the requesting user.
    """
    user: Optional[pulumi.Input[str]] = pulumi.input_property("user")
    """
    User is the user you're testing for. If you specify "User" but not "Group", then is it interpreted as "What if User were not a member of any groups
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, extra: Optional[pulumi.Input[Dict[str, pulumi.Input[List[pulumi.Input[str]]]]]] = None, group: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, non_resource_attributes: Optional[pulumi.Input['NonResourceAttributesArgs']] = None, resource_attributes: Optional[pulumi.Input['ResourceAttributesArgs']] = None, uid: Optional[pulumi.Input[str]] = None, user: Optional[pulumi.Input[str]] = None) -> None:
        """
        SubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
        :param pulumi.Input[Dict[str, pulumi.Input[List[pulumi.Input[str]]]]] extra: Extra corresponds to the user.Info.GetExtra() method from the authenticator.  Since that is input to the authorizer it needs a reflection here.
        :param pulumi.Input[List[pulumi.Input[str]]] group: Groups is the groups you're testing for.
        :param pulumi.Input['NonResourceAttributesArgs'] non_resource_attributes: NonResourceAttributes describes information for a non-resource access request
        :param pulumi.Input['ResourceAttributesArgs'] resource_attributes: ResourceAuthorizationAttributes describes information for a resource access request
        :param pulumi.Input[str] uid: UID information about the requesting user.
        :param pulumi.Input[str] user: User is the user you're testing for. If you specify "User" but not "Group", then is it interpreted as "What if User were not a member of any groups
        """
        __self__.extra = extra
        __self__.group = group
        __self__.non_resource_attributes = non_resource_attributes
        __self__.resource_attributes = resource_attributes
        __self__.uid = uid
        __self__.user = user

