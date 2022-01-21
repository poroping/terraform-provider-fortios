---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_h2qposuprovider"
description: |-
  Get information on a fortios Configure online sign up (OSU) provider list.
---

# Data Source: fortios_wirelesscontrollerhotspot20_h2qposuprovider
Use this data source to get information on a fortios Configure online sign up (OSU) provider list.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) OSU provider ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `icon` - OSU provider icon.
* `name` - OSU provider ID.
* `osu_method` - OSU method list.
* `osu_nai` - OSU NAI.
* `server_uri` - Server URI.
* `friendly_name` - OSU provider friendly name.The structure of `friendly_name` block is documented below.

The `friendly_name` block contains:

* `friendly_name` - OSU provider friendly name.
* `index` - OSU provider friendly name index.
* `lang` - Language code.
* `service_description` - OSU service name.The structure of `service_description` block is documented below.

The `service_description` block contains:

* `lang` - Language code.
* `service_description` - Service description.
* `service_id` - OSU service ID.
