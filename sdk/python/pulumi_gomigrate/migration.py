# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['MigrationArgs', 'Migration']

@pulumi.input_type
class MigrationArgs:
    def __init__(__self__, *,
                 database_url: pulumi.Input[str],
                 prev_version: pulumi.Input[int],
                 source_url: pulumi.Input[str],
                 version: pulumi.Input[int]):
        """
        The set of arguments for constructing a Migration resource.
        :param pulumi.Input[str] database_url: Database URL to run the migrations on
        :param pulumi.Input[int] prev_version: Previous version to migrate on undo
        :param pulumi.Input[str] source_url: Source URL for the migrations
        :param pulumi.Input[int] version: Version to migrate
        """
        pulumi.set(__self__, "database_url", database_url)
        pulumi.set(__self__, "prev_version", prev_version)
        pulumi.set(__self__, "source_url", source_url)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="databaseURL")
    def database_url(self) -> pulumi.Input[str]:
        """
        Database URL to run the migrations on
        """
        return pulumi.get(self, "database_url")

    @database_url.setter
    def database_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_url", value)

    @property
    @pulumi.getter(name="prevVersion")
    def prev_version(self) -> pulumi.Input[int]:
        """
        Previous version to migrate on undo
        """
        return pulumi.get(self, "prev_version")

    @prev_version.setter
    def prev_version(self, value: pulumi.Input[int]):
        pulumi.set(self, "prev_version", value)

    @property
    @pulumi.getter(name="sourceURL")
    def source_url(self) -> pulumi.Input[str]:
        """
        Source URL for the migrations
        """
        return pulumi.get(self, "source_url")

    @source_url.setter
    def source_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_url", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[int]:
        """
        Version to migrate
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[int]):
        pulumi.set(self, "version", value)


class Migration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_url: Optional[pulumi.Input[str]] = None,
                 prev_version: Optional[pulumi.Input[int]] = None,
                 source_url: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a Migration resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_url: Database URL to run the migrations on
        :param pulumi.Input[int] prev_version: Previous version to migrate on undo
        :param pulumi.Input[str] source_url: Source URL for the migrations
        :param pulumi.Input[int] version: Version to migrate
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MigrationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Migration resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param MigrationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MigrationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_url: Optional[pulumi.Input[str]] = None,
                 prev_version: Optional[pulumi.Input[int]] = None,
                 source_url: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MigrationArgs.__new__(MigrationArgs)

            if database_url is None and not opts.urn:
                raise TypeError("Missing required property 'database_url'")
            __props__.__dict__["database_url"] = None if database_url is None else pulumi.Output.secret(database_url)
            if prev_version is None and not opts.urn:
                raise TypeError("Missing required property 'prev_version'")
            __props__.__dict__["prev_version"] = prev_version
            if source_url is None and not opts.urn:
                raise TypeError("Missing required property 'source_url'")
            __props__.__dict__["source_url"] = None if source_url is None else pulumi.Output.secret(source_url)
            if version is None and not opts.urn:
                raise TypeError("Missing required property 'version'")
            __props__.__dict__["version"] = version
            __props__.__dict__["migrated_at"] = None
        super(Migration, __self__).__init__(
            'gomigrate:index:Migration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Migration':
        """
        Get an existing Migration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = MigrationArgs.__new__(MigrationArgs)

        __props__.__dict__["migrated_at"] = None
        return Migration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="migratedAt")
    def migrated_at(self) -> pulumi.Output[str]:
        """
        Date of the migration
        """
        return pulumi.get(self, "migrated_at")

