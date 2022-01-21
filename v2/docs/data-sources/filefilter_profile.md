---
subcategory: "FortiGate File-Filter"
layout: "fortios"
page_title: "FortiOS: fortios_filefilter_profile"
description: |-
  Get information on a fortios Configure file-filter profiles.
---

# Data Source: fortios_filefilter_profile
Use this data source to get information on a fortios Configure file-filter profiles.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `extended_log` - Enable/disable file-filter extended logging.
* `feature_set` - Flow/proxy feature set.
* `log` - Enable/disable file-filter logging.
* `name` - Profile name.
* `replacemsg_group` - Replacement message group
* `scan_archive_contents` - Enable/disable archive contents scan.
* `rules` - File filter rules.The structure of `rules` block is documented below.

The `rules` block contains:

* `action` - Action taken for matched file.
* `comment` - Comment.
* `direction` - Traffic direction. (HTTP, FTP, SSH, CIFS only)
* `name` - File-filter rule name.
* `password_protected` - Match password-protected files.
* `protocol` - Protocols to apply rule to.
* `file_type` - Select file type.The structure of `file_type` block is documented below.

The `file_type` block contains:

* `name` - File type name.
