# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .FlowSchema import *
from .FlowSchemaList import *
from .PriorityLevelConfiguration import *
from .PriorityLevelConfigurationList import *
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
            if typ == "kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:FlowSchema":
                return FlowSchema(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:FlowSchemaList":
                return FlowSchemaList(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:PriorityLevelConfiguration":
                return PriorityLevelConfiguration(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:PriorityLevelConfigurationList":
                return PriorityLevelConfigurationList(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("kubernetes", "flowcontrol.apiserver.k8s.io/v1alpha1", _module_instance)

_register_module()
