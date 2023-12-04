// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Example
{
    [ExampleResourceType("example:index:HttpApi")]
    public partial class HttpApi : global::Pulumi.ComponentResource
    {
        [Output("authorizers")]
        public Output<ImmutableArray<string>> Authorizers { get; private set; } = null!;


        /// <summary>
        /// Create a HttpApi resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HttpApi(string name, HttpApiArgs? args = null, ComponentResourceOptions? options = null)
            : base("example:index:HttpApi", name, args ?? new HttpApiArgs(), MakeResourceOptions(options, ""), remote: true)
        {
        }

        private static ComponentResourceOptions MakeResourceOptions(ComponentResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new ComponentResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = ComponentResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class HttpApiArgs : global::Pulumi.ResourceArgs
    {
        [Input("authorizers")]
        private Dictionary<string, Inputs.HttpAuthorizerArgs>? _authorizers;
        public Dictionary<string, Inputs.HttpAuthorizerArgs> Authorizers
        {
            get => _authorizers ?? (_authorizers = new Dictionary<string, Inputs.HttpAuthorizerArgs>());
            set => _authorizers = value;
        }

        public HttpApiArgs()
        {
        }
        public static new HttpApiArgs Empty => new HttpApiArgs();
    }
}
