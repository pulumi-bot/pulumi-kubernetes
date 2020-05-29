// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Meta.V1
{

    [OutputType]
    public sealed class ManagedFieldsEntry
    {
        /// <summary>
        /// APIVersion defines the version of this resource that this field set applies to. The format is "group/version" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted.
        /// </summary>
        public readonly string ApiVersion;
        /// <summary>
        /// FieldsType is the discriminator for the different fields format and version. There is currently only one possible value: "FieldsV1"
        /// </summary>
        public readonly string FieldsType;
        /// <summary>
        /// FieldsV1 holds the first JSON version format as described in the "FieldsV1" type.
        /// </summary>
        public readonly System.Text.Json.JsonElement FieldsV1;
        /// <summary>
        /// Manager is an identifier of the workflow managing these fields.
        /// </summary>
        public readonly string Manager;
        /// <summary>
        /// Operation is the type of operation which lead to this ManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'.
        /// </summary>
        public readonly string Operation;
        /// <summary>
        /// Time is timestamp of when these fields were set. It should always be empty if Operation is 'Apply'
        /// </summary>
        public readonly string Time;

        [OutputConstructor]
        private ManagedFieldsEntry(
            string apiVersion,

            string fieldsType,

            System.Text.Json.JsonElement fieldsV1,

            string manager,

            string operation,

            string time)
        {
            ApiVersion = apiVersion;
            FieldsType = fieldsType;
            FieldsV1 = fieldsV1;
            Manager = manager;
            Operation = operation;
            Time = time;
        }
    }
}
