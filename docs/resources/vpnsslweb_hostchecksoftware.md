---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnsslweb_hostchecksoftware"
description: |-
  SSL-VPN host check software.
---

## fortios_vpnsslweb_hostchecksoftware
SSL-VPN host check software.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `guid` - Globally unique ID.
* `name` - Name.
* `os_type` - OS type. Valid values: `windows` `macos` .
* `type` - Type. Valid values: `av` `fw` .
* `version` - Version.
* `check_item_list` - Check item list. The structure of `check_item_list` block is documented below.

The `check_item_list` block contains:

* `action` - Action. Valid values: `require` `deny` .
* `id` - ID (0 - 4294967295).
* `target` - Target.
* `type` - Type. Valid values: `file` `registry` `process` .
* `version` - Version.
* `md5s` - MD5 checksum. The structure of `md5s` block is documented below.

The `md5s` block contains:

* `id` - Hex string of MD5 checksum.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_vpnsslweb_hostchecksoftware can be imported using:
```sh
terraform import fortios_vpnsslweb_hostchecksoftware.labelname {{mkey}}
```
