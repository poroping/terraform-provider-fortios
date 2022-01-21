---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_iptranslation"
description: |-
  Get information on a fortios Configure firewall IP-translation.
---

# Data Source: fortios_firewall_iptranslation
Use this data source to get information on a fortios Configure firewall IP-translation.


## Example Usage

```hcl

```

## Argument Reference

* `transid` - (Required) IP translation ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `endip` - Final IPv4 address (inclusive) in the range of the addresses to be translated (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
* `map_startip` - Address to be used as the starting point for translation in the range (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
* `startip` - First IPv4 address (inclusive) in the range of the addresses to be translated (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
* `transid` - IP translation ID.
* `type` - IP translation type (option: SCTP).
