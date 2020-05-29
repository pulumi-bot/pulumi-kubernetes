// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1
{

    [OutputType]
    public sealed class AllowedFlexVolume
    {
        /// <summary>
        /// driver is the name of the Flexvolume driver.
        /// </summary>
        public readonly string Driver;

        [OutputConstructor]
        private AllowedFlexVolume(string driver)
        {
            Driver = driver;
        }
    }
}
