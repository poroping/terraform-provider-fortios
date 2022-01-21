---
subcategory: "FortiGate Ssh-Filter"
layout: "fortios"
page_title: "FortiOS: fortios_sshfilter_profile"
description: |-
  Configure SSH filter profile.
---

## fortios_sshfilter_profile
Configure SSH filter profile.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `block` - SSH blocking options. Valid values: `x11` `shell` `exec` `port-forward` `tun-forward` `sftp` `scp` `unknown` .
* `default_command_log` - Enable/disable logging unmatched shell commands. Valid values: `enable` `disable` .
* `log` - SSH logging options. Valid values: `x11` `shell` `exec` `port-forward` `tun-forward` `sftp` `scp` `unknown` .
* `name` - SSH filter profile name.
* `file_filter` - File filter. The structure of `file_filter` block is documented below.

The `file_filter` block contains:

* `log` - Enable/disable file filter logging. Valid values: `enable` `disable` .
* `scan_archive_contents` - Enable/disable file filter archive contents scan. Valid values: `enable` `disable` .
* `status` - Enable/disable file filter. Valid values: `enable` `disable` .
* `entries` - File filter entries. The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Action taken for matched file. Valid values: `log` `block` .
* `comment` - Comment.
* `direction` - Match files transmitted in the session's originating or reply direction. Valid values: `incoming` `outgoing` `any` .
* `filter` - Add a file filter.
* `password_protected` - Match password-protected files. Valid values: `yes` `any` .
* `file_type` - Select file type. The structure of `file_type` block is documented below.

The `file_type` block contains:

* `name` - File type name. This attribute must reference one of the following datasources: `antivirus.filetype.name` .
* `shell_commands` - SSH command filter. The structure of `shell_commands` block is documented below.

The `shell_commands` block contains:

* `action` - Action to take for SSH shell command matches. Valid values: `block` `allow` .
* `alert` - Enable/disable alert. Valid values: `enable` `disable` .
* `id` - Id.
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `pattern` - SSH shell command pattern.
* `severity` - Log severity. Valid values: `low` `medium` `high` `critical` .
* `type` - Matching type. Valid values: `simple` `regex` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_sshfilter_profile can be imported using:
```sh
terraform import fortios_sshfilter_profile.labelname {{mkey}}
```
