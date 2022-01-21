---
subcategory: "FortiGate Ssh-Filter"
layout: "fortios"
page_title: "FortiOS: fortios_sshfilter_profile"
description: |-
  Get information on a fortios Configure SSH filter profile.
---

# Data Source: fortios_sshfilter_profile
Use this data source to get information on a fortios Configure SSH filter profile.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) SSH filter profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `block` - SSH blocking options.
* `default_command_log` - Enable/disable logging unmatched shell commands.
* `log` - SSH logging options.
* `name` - SSH filter profile name.
* `file_filter` - File filter.The structure of `file_filter` block is documented below.

The `file_filter` block contains:

* `log` - Enable/disable file filter logging.
* `scan_archive_contents` - Enable/disable file filter archive contents scan.
* `status` - Enable/disable file filter.
* `entries` - File filter entries.The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Action taken for matched file.
* `comment` - Comment.
* `direction` - Match files transmitted in the session's originating or reply direction.
* `filter` - Add a file filter.
* `password_protected` - Match password-protected files.
* `file_type` - Select file type.The structure of `file_type` block is documented below.

The `file_type` block contains:

* `name` - File type name.
* `shell_commands` - SSH command filter.The structure of `shell_commands` block is documented below.

The `shell_commands` block contains:

* `action` - Action to take for SSH shell command matches.
* `alert` - Enable/disable alert.
* `id` - Id.
* `log` - Enable/disable logging.
* `pattern` - SSH shell command pattern.
* `severity` - Log severity.
* `type` - Matching type.
