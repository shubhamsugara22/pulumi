# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._inputs import *

__all__ = ['HttpApiArgs', 'HttpApi']

@pulumi.input_type
class HttpApiArgs:
    def __init__(__self__, *,
                 authorizers: Optional[Mapping[str, 'HttpAuthorizerArgs']] = None):
        """
        The set of arguments for constructing a HttpApi resource.
        """
        if authorizers is not None:
            pulumi.set(__self__, "authorizers", authorizers)

    @property
    @pulumi.getter
    def authorizers(self) -> Optional[Mapping[str, 'HttpAuthorizerArgs']]:
        return pulumi.get(self, "authorizers")

    @authorizers.setter
    def authorizers(self, value: Optional[Mapping[str, 'HttpAuthorizerArgs']]):
        pulumi.set(self, "authorizers", value)


class HttpApi(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorizers: Optional[Mapping[str, pulumi.InputType['HttpAuthorizerArgs']]] = None,
                 __props__=None):
        """
        Create a HttpApi resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[HttpApiArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a HttpApi resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param HttpApiArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HttpApiArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorizers: Optional[Mapping[str, pulumi.InputType['HttpAuthorizerArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HttpApiArgs.__new__(HttpApiArgs)

            __props__.__dict__["authorizers"] = authorizers
        super(HttpApi, __self__).__init__(
            'example:index:HttpApi',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def authorizers(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "authorizers")

