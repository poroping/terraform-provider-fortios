---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpn_l2tp"
description: |-
  Get information on a fortios Configure L2TP.
---

# Data Source: fortios_vpn_l2tp
Use this data source to get information on a fortios Configure L2TP.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `compress` - Enable/disable data compression.
* `eip` - End IP.
* `enforce_ipsec` - Enable/disable IPsec enforcement.
* `hello_interval` - L2TP hello message interval in seconds (0 - 3600 sec, default = 60).
* `lcp_echo_interval` - Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
* `lcp_max_echo_fails` - Maximum number of missed LCP echo messages before disconnect.
* `sip` - Start IP.
* `status` - Enable/disable FortiGate as a L2TP gateway.
* `usrgrp` - User group.
