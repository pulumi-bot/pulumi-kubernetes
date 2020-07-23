# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from ... import _utilities, _tables
from ... import meta as _meta

@pulumi.input_type
class AggregationRuleArgs:
    cluster_role_selectors: Optional[pulumi.Input[List[pulumi.Input['_meta.v1.LabelSelectorArgs']]]] = pulumi.input_property("clusterRoleSelectors")
    """
    ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole's permissions will be added
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, cluster_role_selectors: Optional[pulumi.Input[List[pulumi.Input['_meta.v1.LabelSelectorArgs']]]] = None) -> None:
        """
        AggregationRule describes how to locate ClusterRoles to aggregate into the ClusterRole
        :param pulumi.Input[List[pulumi.Input['_meta.v1.LabelSelectorArgs']]] cluster_role_selectors: ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole's permissions will be added
        """
        __self__.cluster_role_selectors = cluster_role_selectors

@pulumi.input_type
class ClusterRoleArgs:
    aggregation_rule: Optional[pulumi.Input['AggregationRuleArgs']] = pulumi.input_property("aggregationRule")
    """
    AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
    """
    api_version: Optional[pulumi.Input[str]] = pulumi.input_property("apiVersion")
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    """
    kind: Optional[pulumi.Input[str]] = pulumi.input_property("kind")
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    """
    metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = pulumi.input_property("metadata")
    """
    Standard object's metadata.
    """
    rules: Optional[pulumi.Input[List[pulumi.Input['PolicyRuleArgs']]]] = pulumi.input_property("rules")
    """
    Rules holds all the PolicyRules for this ClusterRole
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, aggregation_rule: Optional[pulumi.Input['AggregationRuleArgs']] = None, api_version: Optional[pulumi.Input[str]] = None, kind: Optional[pulumi.Input[str]] = None, metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None, rules: Optional[pulumi.Input[List[pulumi.Input['PolicyRuleArgs']]]] = None) -> None:
        """
        ClusterRole is a cluster level, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding or ClusterRoleBinding.
        :param pulumi.Input['AggregationRuleArgs'] aggregation_rule: AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object's metadata.
        :param pulumi.Input[List[pulumi.Input['PolicyRuleArgs']]] rules: Rules holds all the PolicyRules for this ClusterRole
        """
        __self__.aggregation_rule = aggregation_rule
        __self__.api_version = 'rbac.authorization.k8s.io/v1'
        __self__.kind = 'ClusterRole'
        __self__.metadata = metadata
        __self__.rules = rules

@pulumi.input_type
class ClusterRoleBindingArgs:
    role_ref: pulumi.Input['RoleRefArgs'] = pulumi.input_property("roleRef")
    """
    RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
    """
    api_version: Optional[pulumi.Input[str]] = pulumi.input_property("apiVersion")
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    """
    kind: Optional[pulumi.Input[str]] = pulumi.input_property("kind")
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    """
    metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = pulumi.input_property("metadata")
    """
    Standard object's metadata.
    """
    subjects: Optional[pulumi.Input[List[pulumi.Input['SubjectArgs']]]] = pulumi.input_property("subjects")
    """
    Subjects holds references to the objects the role applies to.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, role_ref: pulumi.Input['RoleRefArgs'], api_version: Optional[pulumi.Input[str]] = None, kind: Optional[pulumi.Input[str]] = None, metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None, subjects: Optional[pulumi.Input[List[pulumi.Input['SubjectArgs']]]] = None) -> None:
        """
        ClusterRoleBinding references a ClusterRole, but not contain it.  It can reference a ClusterRole in the global namespace, and adds who information via Subject.
        :param pulumi.Input['RoleRefArgs'] role_ref: RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object's metadata.
        :param pulumi.Input[List[pulumi.Input['SubjectArgs']]] subjects: Subjects holds references to the objects the role applies to.
        """
        __self__.role_ref = role_ref
        __self__.api_version = 'rbac.authorization.k8s.io/v1'
        __self__.kind = 'ClusterRoleBinding'
        __self__.metadata = metadata
        __self__.subjects = subjects

@pulumi.input_type
class PolicyRuleArgs:
    verbs: pulumi.Input[List[pulumi.Input[str]]] = pulumi.input_property("verbs")
    """
    Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule.  VerbAll represents all kinds.
    """
    api_groups: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("apiGroups")
    """
    APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.
    """
    non_resource_ur_ls: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("nonResourceURLs")
    """
    NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"),  but not both.
    """
    resource_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("resourceNames")
    """
    ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.
    """
    resources: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("resources")
    """
    Resources is a list of resources this rule applies to.  ResourceAll represents all resources.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, verbs: pulumi.Input[List[pulumi.Input[str]]], api_groups: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, non_resource_ur_ls: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, resource_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, resources: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> None:
        """
        PolicyRule holds information that describes a policy rule, but does not contain information about who the rule applies to or which namespace the rule applies to.
        :param pulumi.Input[List[pulumi.Input[str]]] verbs: Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule.  VerbAll represents all kinds.
        :param pulumi.Input[List[pulumi.Input[str]]] api_groups: APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.
        :param pulumi.Input[List[pulumi.Input[str]]] non_resource_ur_ls: NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"),  but not both.
        :param pulumi.Input[List[pulumi.Input[str]]] resource_names: ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.
        :param pulumi.Input[List[pulumi.Input[str]]] resources: Resources is a list of resources this rule applies to.  ResourceAll represents all resources.
        """
        __self__.verbs = verbs
        __self__.api_groups = api_groups
        __self__.non_resource_ur_ls = non_resource_ur_ls
        __self__.resource_names = resource_names
        __self__.resources = resources

@pulumi.input_type
class RoleArgs:
    api_version: Optional[pulumi.Input[str]] = pulumi.input_property("apiVersion")
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    """
    kind: Optional[pulumi.Input[str]] = pulumi.input_property("kind")
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    """
    metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = pulumi.input_property("metadata")
    """
    Standard object's metadata.
    """
    rules: Optional[pulumi.Input[List[pulumi.Input['PolicyRuleArgs']]]] = pulumi.input_property("rules")
    """
    Rules holds all the PolicyRules for this Role
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, api_version: Optional[pulumi.Input[str]] = None, kind: Optional[pulumi.Input[str]] = None, metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None, rules: Optional[pulumi.Input[List[pulumi.Input['PolicyRuleArgs']]]] = None) -> None:
        """
        Role is a namespaced, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object's metadata.
        :param pulumi.Input[List[pulumi.Input['PolicyRuleArgs']]] rules: Rules holds all the PolicyRules for this Role
        """
        __self__.api_version = 'rbac.authorization.k8s.io/v1'
        __self__.kind = 'Role'
        __self__.metadata = metadata
        __self__.rules = rules

@pulumi.input_type
class RoleBindingArgs:
    role_ref: pulumi.Input['RoleRefArgs'] = pulumi.input_property("roleRef")
    """
    RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
    """
    api_version: Optional[pulumi.Input[str]] = pulumi.input_property("apiVersion")
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    """
    kind: Optional[pulumi.Input[str]] = pulumi.input_property("kind")
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    """
    metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = pulumi.input_property("metadata")
    """
    Standard object's metadata.
    """
    subjects: Optional[pulumi.Input[List[pulumi.Input['SubjectArgs']]]] = pulumi.input_property("subjects")
    """
    Subjects holds references to the objects the role applies to.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, role_ref: pulumi.Input['RoleRefArgs'], api_version: Optional[pulumi.Input[str]] = None, kind: Optional[pulumi.Input[str]] = None, metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None, subjects: Optional[pulumi.Input[List[pulumi.Input['SubjectArgs']]]] = None) -> None:
        """
        RoleBinding references a role, but does not contain it.  It can reference a Role in the same namespace or a ClusterRole in the global namespace. It adds who information via Subjects and namespace information by which namespace it exists in.  RoleBindings in a given namespace only have effect in that namespace.
        :param pulumi.Input['RoleRefArgs'] role_ref: RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object's metadata.
        :param pulumi.Input[List[pulumi.Input['SubjectArgs']]] subjects: Subjects holds references to the objects the role applies to.
        """
        __self__.role_ref = role_ref
        __self__.api_version = 'rbac.authorization.k8s.io/v1'
        __self__.kind = 'RoleBinding'
        __self__.metadata = metadata
        __self__.subjects = subjects

@pulumi.input_type
class RoleRefArgs:
    api_group: pulumi.Input[str] = pulumi.input_property("apiGroup")
    """
    APIGroup is the group for the resource being referenced
    """
    kind: pulumi.Input[str] = pulumi.input_property("kind")
    """
    Kind is the type of resource being referenced
    """
    name: pulumi.Input[str] = pulumi.input_property("name")
    """
    Name is the name of resource being referenced
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, api_group: pulumi.Input[str], kind: pulumi.Input[str], name: pulumi.Input[str]) -> None:
        """
        RoleRef contains information that points to the role being used
        :param pulumi.Input[str] api_group: APIGroup is the group for the resource being referenced
        :param pulumi.Input[str] kind: Kind is the type of resource being referenced
        :param pulumi.Input[str] name: Name is the name of resource being referenced
        """
        __self__.api_group = api_group
        __self__.kind = kind
        __self__.name = name

@pulumi.input_type
class SubjectArgs:
    kind: pulumi.Input[str] = pulumi.input_property("kind")
    """
    Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount". If the Authorizer does not recognized the kind value, the Authorizer should report an error.
    """
    name: pulumi.Input[str] = pulumi.input_property("name")
    """
    Name of the object being referenced.
    """
    api_group: Optional[pulumi.Input[str]] = pulumi.input_property("apiGroup")
    """
    APIGroup holds the API group of the referenced subject. Defaults to "" for ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io" for User and Group subjects.
    """
    namespace: Optional[pulumi.Input[str]] = pulumi.input_property("namespace")
    """
    Namespace of the referenced object.  If the object kind is non-namespace, such as "User" or "Group", and this value is not empty the Authorizer should report an error.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, kind: pulumi.Input[str], name: pulumi.Input[str], api_group: Optional[pulumi.Input[str]] = None, namespace: Optional[pulumi.Input[str]] = None) -> None:
        """
        Subject contains a reference to the object or user identities a role binding applies to.  This can either hold a direct API object reference, or a value for non-objects such as user and group names.
        :param pulumi.Input[str] kind: Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount". If the Authorizer does not recognized the kind value, the Authorizer should report an error.
        :param pulumi.Input[str] name: Name of the object being referenced.
        :param pulumi.Input[str] api_group: APIGroup holds the API group of the referenced subject. Defaults to "" for ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io" for User and Group subjects.
        :param pulumi.Input[str] namespace: Namespace of the referenced object.  If the object kind is non-namespace, such as "User" or "Group", and this value is not empty the Authorizer should report an error.
        """
        __self__.kind = kind
        __self__.name = name
        __self__.api_group = api_group
        __self__.namespace = namespace

