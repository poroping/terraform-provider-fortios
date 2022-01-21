---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallshaper_peripshaper"
description: |-
  Configure per-IP traffic shaper.
---

## fortios_firewallshaper_peripshaper
Configure per-IP traffic shaper.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `bandwidth_unit` - Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps` `mbps` `gbps` .
* `diffserv_forward` - Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable` `disable` .
* `diffserv_reverse` - Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable` `disable` .
* `diffservcode_forward` - Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
* `diffservcode_rev` - Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
* `max_bandwidth` - Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
* `max_concurrent_session` - Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
* `max_concurrent_tcp_session` - Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
* `max_concurrent_udp_session` - Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
* `name` - Traffic shaper name.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewallshaper_peripshaper can be imported using:
```sh
terraform import fortios_firewallshaper_peripshaper.labelname {{mkey}}
```
