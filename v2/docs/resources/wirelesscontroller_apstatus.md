---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_apstatus"
description: |-
  Configure access point status (rogue | accepted | suppressed).
---

## fortios_wirelesscontroller_apstatus
Configure access point status (rogue | accepted | suppressed).

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.

* `bssid` - Access Point's (AP's) BSSID.
* `id` - AP ID.
* `ssid` - Access Point's (AP's) SSID.
* `status` - Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue` `accepted` `suppressed` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_apstatus can be imported using:
```sh
terraform import fortios_wirelesscontroller_apstatus.labelname {{mkey}}
```
