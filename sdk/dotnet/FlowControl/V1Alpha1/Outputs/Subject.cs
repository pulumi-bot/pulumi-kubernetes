// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Runtime.Serialization;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Alpha1
{

    [OutputType]
    public sealed class Subject
    {
        public readonly Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Alpha1.GroupSubject Group;
        /// <summary>
        /// Required
        /// </summary>
        public readonly string Kind;
        public readonly Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Alpha1.ServiceAccountSubject ServiceAccount;
        public readonly Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Alpha1.UserSubject User;

        [OutputConstructor]
        private Subject(
            Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Alpha1.GroupSubject group,

            string kind,

            Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Alpha1.ServiceAccountSubject serviceAccount,

            Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Alpha1.UserSubject user)
        {
            Group = group;
            Kind = kind;
            ServiceAccount = serviceAccount;
            User = user;
        }
    }
}
