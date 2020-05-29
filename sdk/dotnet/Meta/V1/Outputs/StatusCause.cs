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
    public sealed class StatusCause
    {
        /// <summary>
        /// The field of the resource that has caused this error, as named by its JSON serialization. May include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may appear more than once in an array of causes due to fields having multiple errors. Optional.
        /// 
        /// Examples:
        ///   "name" - the field "name" on the current resource
        ///   "items[0].name" - the field "name" on the first array entry in "items"
        /// </summary>
        public readonly string Field;
        /// <summary>
        /// A human-readable description of the cause of the error.  This field may be presented as-is to a reader.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// A machine-readable description of the cause of the error. If this value is empty there is no information available.
        /// </summary>
        public readonly string Reason;

        [OutputConstructor]
        private StatusCause(
            string field,

            string message,

            string reason)
        {
            Field = field;
            Message = message;
            Reason = reason;
        }
    }
}
