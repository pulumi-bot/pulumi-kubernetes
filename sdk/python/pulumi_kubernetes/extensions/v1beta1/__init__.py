# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .DaemonSet import *
from .DaemonSetList import *
from .Deployment import *
from .DeploymentList import *
from .Ingress import *
from .IngressList import *
from .NetworkPolicy import *
from .NetworkPolicyList import *
from .PodSecurityPolicy import *
from .PodSecurityPolicyList import *
from .ReplicaSet import *
from .ReplicaSetList import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi

    class Module(pulumi.runtime.ResourceModule):
        def version(self):
            return None

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "kubernetes:extensions/v1beta1:DaemonSet":
                return DaemonSet(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:extensions/v1beta1:DaemonSetList":
                return DaemonSetList(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:extensions/v1beta1:Deployment":
                return Deployment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:extensions/v1beta1:DeploymentList":
                return DeploymentList(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:extensions/v1beta1:Ingress":
                return Ingress(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:extensions/v1beta1:IngressList":
                return IngressList(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:extensions/v1beta1:NetworkPolicy":
                return NetworkPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:extensions/v1beta1:NetworkPolicyList":
                return NetworkPolicyList(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:extensions/v1beta1:PodSecurityPolicy":
                return PodSecurityPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:extensions/v1beta1:PodSecurityPolicyList":
                return PodSecurityPolicyList(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:extensions/v1beta1:ReplicaSet":
                return ReplicaSet(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:extensions/v1beta1:ReplicaSetList":
                return ReplicaSetList(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("kubernetes", "extensions/v1beta1", _module_instance)

_register_module()
