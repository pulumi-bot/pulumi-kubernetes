# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .EndpointSlice import *
from .EndpointSliceList import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi

    class Module(pulumi.runtime.ResourceModule):
        def version(self):
            return None

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "kubernetes:discovery.k8s.io/v1beta1:EndpointSlice":
                return EndpointSlice(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:discovery.k8s.io/v1beta1:EndpointSliceList":
                return EndpointSliceList(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("kubernetes", "discovery.k8s.io/v1beta1", _module_instance)

_register_module()
