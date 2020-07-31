# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ... import meta as _meta

__all__ = [
    'AggregationRule',
    'ClusterRole',
    'ClusterRoleBinding',
    'PolicyRule',
    'Role',
    'RoleBinding',
    'RoleRef',
    'Subject',
]

@pulumi.output_type
class AggregationRule(dict):
    """
    AggregationRule describes how to locate ClusterRoles to aggregate into the ClusterRole
    """
    cluster_role_selectors: Optional[List['_meta.v1.outputs.LabelSelector']] = pulumi.output_property("clusterRoleSelectors")
    """
    ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole's permissions will be added
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterRole(dict):
    """
    ClusterRole is a cluster level, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding or ClusterRoleBinding.
    """
    aggregation_rule: Optional['outputs.AggregationRule'] = pulumi.output_property("aggregationRule")
    """
    AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
    """
    api_version: Optional[str] = pulumi.output_property("apiVersion")
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    """
    kind: Optional[str] = pulumi.output_property("kind")
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    """
    metadata: Optional['_meta.v1.outputs.ObjectMeta'] = pulumi.output_property("metadata")
    """
    Standard object's metadata.
    """
    rules: Optional[List['outputs.PolicyRule']] = pulumi.output_property("rules")
    """
    Rules holds all the PolicyRules for this ClusterRole
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterRoleBinding(dict):
    """
    ClusterRoleBinding references a ClusterRole, but not contain it.  It can reference a ClusterRole in the global namespace, and adds who information via Subject.
    """
    api_version: Optional[str] = pulumi.output_property("apiVersion")
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    """
    kind: Optional[str] = pulumi.output_property("kind")
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    """
    metadata: Optional['_meta.v1.outputs.ObjectMeta'] = pulumi.output_property("metadata")
    """
    Standard object's metadata.
    """
    role_ref: 'outputs.RoleRef' = pulumi.output_property("roleRef")
    """
    RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
    """
    subjects: Optional[List['outputs.Subject']] = pulumi.output_property("subjects")
    """
    Subjects holds references to the objects the role applies to.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyRule(dict):
    """
    PolicyRule holds information that describes a policy rule, but does not contain information about who the rule applies to or which namespace the rule applies to.
    """
    api_groups: Optional[List[str]] = pulumi.output_property("apiGroups")
    """
    APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.
    """
    non_resource_ur_ls: Optional[List[str]] = pulumi.output_property("nonResourceURLs")
    """
    NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"),  but not both.
    """
    resource_names: Optional[List[str]] = pulumi.output_property("resourceNames")
    """
    ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.
    """
    resources: Optional[List[str]] = pulumi.output_property("resources")
    """
    Resources is a list of resources this rule applies to.  ResourceAll represents all resources.
    """
    verbs: List[str] = pulumi.output_property("verbs")
    """
    Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule.  VerbAll represents all kinds.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class Role(dict):
    """
    Role is a namespaced, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding.
    """
    api_version: Optional[str] = pulumi.output_property("apiVersion")
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    """
    kind: Optional[str] = pulumi.output_property("kind")
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    """
    metadata: Optional['_meta.v1.outputs.ObjectMeta'] = pulumi.output_property("metadata")
    """
    Standard object's metadata.
    """
    rules: Optional[List['outputs.PolicyRule']] = pulumi.output_property("rules")
    """
    Rules holds all the PolicyRules for this Role
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class RoleBinding(dict):
    """
    RoleBinding references a role, but does not contain it.  It can reference a Role in the same namespace or a ClusterRole in the global namespace. It adds who information via Subjects and namespace information by which namespace it exists in.  RoleBindings in a given namespace only have effect in that namespace.
    """
    api_version: Optional[str] = pulumi.output_property("apiVersion")
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
    """
    kind: Optional[str] = pulumi.output_property("kind")
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
    """
    metadata: Optional['_meta.v1.outputs.ObjectMeta'] = pulumi.output_property("metadata")
    """
    Standard object's metadata.
    """
    role_ref: 'outputs.RoleRef' = pulumi.output_property("roleRef")
    """
    RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
    """
    subjects: Optional[List['outputs.Subject']] = pulumi.output_property("subjects")
    """
    Subjects holds references to the objects the role applies to.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class RoleRef(dict):
    """
    RoleRef contains information that points to the role being used
    """
    api_group: str = pulumi.output_property("apiGroup")
    """
    APIGroup is the group for the resource being referenced
    """
    kind: str = pulumi.output_property("kind")
    """
    Kind is the type of resource being referenced
    """
    name: str = pulumi.output_property("name")
    """
    Name is the name of resource being referenced
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class Subject(dict):
    """
    Subject contains a reference to the object or user identities a role binding applies to.  This can either hold a direct API object reference, or a value for non-objects such as user and group names.
    """
    api_group: Optional[str] = pulumi.output_property("apiGroup")
    """
    APIGroup holds the API group of the referenced subject. Defaults to "" for ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io" for User and Group subjects.
    """
    kind: str = pulumi.output_property("kind")
    """
    Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount". If the Authorizer does not recognized the kind value, the Authorizer should report an error.
    """
    name: str = pulumi.output_property("name")
    """
    Name of the object being referenced.
    """
    namespace: Optional[str] = pulumi.output_property("namespace")
    """
    Namespace of the referenced object.  If the object kind is non-namespace, such as "User" or "Group", and this value is not empty the Authorizer should report an error.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


