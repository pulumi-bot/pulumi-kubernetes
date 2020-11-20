# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .PodDisruptionBudget import *
from .PodDisruptionBudgetList import *
from .PodSecurityPolicy import *
from .PodSecurityPolicyList import *
from ._inputs import *
from . import outputs

import pulumi

class Module(pulumi.runtime.ResourceModule):
    def version(self):
        return None

    def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
        if typ == "kubernetes:policy/v1beta1:PodDisruptionBudget":
            return PodDisruptionBudget(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "kubernetes:policy/v1beta1:PodDisruptionBudgetList":
            return PodDisruptionBudgetList(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "kubernetes:policy/v1beta1:PodSecurityPolicy":
            return PodSecurityPolicy(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "kubernetes:policy/v1beta1:PodSecurityPolicyList":
            return PodSecurityPolicyList(name, pulumi.ResourceOptions(urn=urn))
        else:
            raise Exception(f"unknown resource type {typ}")


_module_instance = Module()
pulumi.runtime.register_resource_module("kubernetes", "policy/v1beta1", _module_instance)
