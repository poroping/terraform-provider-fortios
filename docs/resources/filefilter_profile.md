---
subcategory: "FortiGate File-Filter"
layout: "fortios"
page_title: "FortiOS: fortios_filefilter_profile"
description: |-
  Configure file-filter profiles.
---

## fortios_filefilter_profile
Configure file-filter profiles.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comment` - Comment.
* `extended_log` - Enable/disable file-filter extended logging. Valid values: `disable` `enable` .
* `feature_set` - Flow/proxy feature set. Valid values: `flow` `proxy` .
* `log` - Enable/disable file-filter logging. Valid values: `disable` `enable` .
* `name` - Profile name.
* `replacemsg_group` - Replacement message group. This attribute must reference one of the following datasources: `system.replacemsg-group.name` .
* `scan_archive_contents` - Enable/disable archive contents scan. Valid values: `disable` `enable` .
* `rules` - File filter rules. The structure of `rules` block is documented below.

The `rules` block contains:

* `action` - Action taken for matched file. Valid values: `log-only` `block` .
* `comment` - Comment.
* `direction` - Traffic direction (HTTP, FTP, SSH, CIFS only). Valid values: `incoming` `outgoing` `any` .
* `name` - File-filter rule name.
* `password_protected` - Match password-protected files. Valid values: `yes` `any` .
* `protocol` - Protocols to apply rule to. Valid values: `http` `ftp` `smtp` `imap` `pop3` `mapi` `cifs` `ssh` .
* `file_type` - Select file type. The structure of `file_type` block is documented below.

The `file_type` block contains:

* `name` - File type name. This attribute must reference one of the following datasources: `antivirus.filetype.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_filefilter_profile can be imported using:
```sh
terraform import fortios_filefilter_profile.labelname {{mkey}}
```
