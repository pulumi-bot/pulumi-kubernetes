# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .AuditSink import *
from .AuditSinkList import *
from ._inputs import *
from . import outputs

import pulumi

class Module(pulumi.runtime.ResourceModule):
    def version(self):
        return None

    def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
        if typ == "kubernetes:auditregistration.k8s.io/v1alpha1:AuditSink":
            return AuditSink(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "kubernetes:auditregistration.k8s.io/v1alpha1:AuditSinkList":
            return AuditSinkList(name, pulumi.ResourceOptions(urn=urn))
        else:
            raise Exception(f"unknown resource type {typ}")


_module_instance = Module()
pulumi.runtime.register_resource_module("kubernetes", "auditregistration.k8s.io/v1alpha1", _module_instance)
