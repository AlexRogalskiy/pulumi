// *** WARNING: this file was generated by . ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Codegentest
{
    public static class GetClientConfig
    {
        /// <summary>
        /// Use this function to access the current configuration of the native Azure provider.
        /// </summary>
        public static Task<GetClientConfigResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClientConfigResult>("azure-native:codegentest:getClientConfig", InvokeArgs.Empty, options.WithVersion());
    }


    [OutputType]
    public sealed class GetClientConfigResult
    {
        /// <summary>
        /// Azure Client ID (Application Object ID).
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// Azure Object ID of the current user or service principal.
        /// </summary>
        public readonly string ObjectId;
        /// <summary>
        /// Azure Subscription ID
        /// </summary>
        public readonly string SubscriptionId;
        /// <summary>
        /// Azure Tenant ID
        /// </summary>
        public readonly string TenantId;

        [OutputConstructor]
        private GetClientConfigResult(
            string clientId,

            string objectId,

            string subscriptionId,

            string tenantId)
        {
            ClientId = clientId;
            ObjectId = objectId;
            SubscriptionId = subscriptionId;
            TenantId = tenantId;
        }
    }
}
