---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetserviceappend"
description: |-
  Configure additional port mappings for Internet Services.
---

## fortios_firewall_internetserviceappend
Configure additional port mappings for Internet Services.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4` `ipv6` `both` .
* `append_port` - Appending TCP/UDP/SCTP destination port (1 to 65535).
* `match_port` - Matching TCP/UDP/SCTP destination port (0 to 65535, 0 means any port).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_internetserviceappend can be imported using:
```sh
terraform import fortios_firewall_internetserviceappend.labelname {{mkey}}
```
