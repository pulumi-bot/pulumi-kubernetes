# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .ControllerRevision import *
from .ControllerRevisionList import *
from .DaemonSet import *
from .DaemonSetList import *
from .Deployment import *
from .DeploymentList import *
from .ReplicaSet import *
from .ReplicaSetList import *
from .StatefulSet import *
from .StatefulSetList import *
from ._inputs import *
from . import outputs

import pulumi

class Module(pulumi.runtime.ResourceModule):
    def version(self):
        return None

    def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
        if typ == "kubernetes:apps/v1:ControllerRevision":
            return ControllerRevision(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "kubernetes:apps/v1:ControllerRevisionList":
            return ControllerRevisionList(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "kubernetes:apps/v1:DaemonSet":
            return DaemonSet(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "kubernetes:apps/v1:DaemonSetList":
            return DaemonSetList(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "kubernetes:apps/v1:Deployment":
            return Deployment(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "kubernetes:apps/v1:DeploymentList":
            return DeploymentList(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "kubernetes:apps/v1:ReplicaSet":
            return ReplicaSet(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "kubernetes:apps/v1:ReplicaSetList":
            return ReplicaSetList(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "kubernetes:apps/v1:StatefulSet":
            return StatefulSet(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "kubernetes:apps/v1:StatefulSetList":
            return StatefulSetList(name, pulumi.ResourceOptions(urn=urn))
        else:
            raise Exception(f"unknown resource type {typ}")


_module_instance = Module()
pulumi.runtime.register_resource_module("kubernetes", "apps/v1", _module_instance)
