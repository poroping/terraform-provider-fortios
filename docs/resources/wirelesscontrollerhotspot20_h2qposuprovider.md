---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_h2qposuprovider"
description: |-
  Configure online sign up (OSU) provider list.
---

## fortios_wirelesscontrollerhotspot20_h2qposuprovider
Configure online sign up (OSU) provider list.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `icon` - OSU provider icon. This attribute must reference one of the following datasources: `wireless-controller.hotspot20.icon.name` .
* `name` - OSU provider ID.
* `osu_method` - OSU method list. Valid values: `oma-dm` `soap-xml-spp` `reserved` .
* `osu_nai` - OSU NAI.
* `server_uri` - Server URI.
* `friendly_name` - OSU provider friendly name. The structure of `friendly_name` block is documented below.

The `friendly_name` block contains:

* `friendly_name` - OSU provider friendly name.
* `index` - OSU provider friendly name index.
* `lang` - Language code.
* `service_description` - OSU service name. The structure of `service_description` block is documented below.

The `service_description` block contains:

* `lang` - Language code.
* `service_description` - Service description.
* `service_id` - OSU service ID.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontrollerhotspot20_h2qposuprovider can be imported using:
```sh
terraform import fortios_wirelesscontrollerhotspot20_h2qposuprovider.labelname {{mkey}}
```
