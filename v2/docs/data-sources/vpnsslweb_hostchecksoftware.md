---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnsslweb_hostchecksoftware"
description: |-
  Get information on a fortios SSL-VPN host check software.
---

# Data Source: fortios_vpnsslweb_hostchecksoftware
Use this data source to get information on a fortios SSL-VPN host check software.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `guid` - Globally unique ID.
* `name` - Name.
* `os_type` - OS type.
* `type` - Type.
* `version` - Version.
* `check_item_list` - Check item list.The structure of `check_item_list` block is documented below.

The `check_item_list` block contains:

* `action` - Action.
* `id` - ID (0 - 4294967295).
* `target` - Target.
* `type` - Type.
* `version` - Version.
* `md5s` - MD5 checksum.The structure of `md5s` block is documented below.

The `md5s` block contains:

* `id` - Hex string of MD5 checksum.
