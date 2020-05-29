// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.ApiRegistration.V1
{

    /// <summary>
    /// APIServiceStatus contains derived information about an API server
    /// </summary>
    public class APIServiceStatusArgs : Pulumi.ResourceArgs
    {
        [Input("conditions")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.ApiRegistration.V1.APIServiceConditionArgs>? _conditions;

        /// <summary>
        /// Current service state of apiService.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.ApiRegistration.V1.APIServiceConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Pulumi.Kubernetes.Types.Inputs.ApiRegistration.V1.APIServiceConditionArgs>());
            set => _conditions = value;
        }

        public APIServiceStatusArgs()
        {
        }
    }
}
