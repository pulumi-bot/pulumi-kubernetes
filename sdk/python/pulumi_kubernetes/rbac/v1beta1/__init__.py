# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .ClusterRole import *
from .ClusterRoleBinding import *
from .ClusterRoleBindingList import *
from .ClusterRoleList import *
from .Role import *
from .RoleBinding import *
from .RoleBindingList import *
from .RoleList import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from ... import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRole":
                return ClusterRole(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRoleBinding":
                return ClusterRoleBinding(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRoleBindingList":
                return ClusterRoleBindingList(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRoleList":
                return ClusterRoleList(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:rbac.authorization.k8s.io/v1beta1:Role":
                return Role(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:rbac.authorization.k8s.io/v1beta1:RoleBinding":
                return RoleBinding(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:rbac.authorization.k8s.io/v1beta1:RoleBindingList":
                return RoleBindingList(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:rbac.authorization.k8s.io/v1beta1:RoleList":
                return RoleList(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("kubernetes", "rbac.authorization.k8s.io/v1beta1", _module_instance)

_register_module()
