# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .Ingress import *
from .IngressClass import *
from .IngressClassList import *
from .IngressList import *
from ._inputs import *
from . import outputs

import pulumi

class Module(pulumi.runtime.ResourceModule):
    def version(self):
        return None

    def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
        if typ == "kubernetes:networking.k8s.io/v1beta1:Ingress":
            return Ingress(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "kubernetes:networking.k8s.io/v1beta1:IngressClass":
            return IngressClass(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "kubernetes:networking.k8s.io/v1beta1:IngressClassList":
            return IngressClassList(name, pulumi.ResourceOptions(urn=urn))
        elif typ == "kubernetes:networking.k8s.io/v1beta1:IngressList":
            return IngressList(name, pulumi.ResourceOptions(urn=urn))
        else:
            raise Exception(f"unknown resource type {typ}")


_module_instance = Module()
pulumi.runtime.register_resource_module("kubernetes", "networking.k8s.io/v1beta1", _module_instance)
