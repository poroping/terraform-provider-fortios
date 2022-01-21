---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpn_l2tp"
description: |-
  Configure L2TP.
---

## fortios_vpn_l2tp
Configure L2TP.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `compress` - Enable/disable data compression. Valid values: `enable` `disable` .
* `eip` - End IP.
* `enforce_ipsec` - Enable/disable IPsec enforcement. Valid values: `enable` `disable` .
* `hello_interval` - L2TP hello message interval in seconds (0 - 3600 sec, default = 60).
* `lcp_echo_interval` - Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
* `lcp_max_echo_fails` - Maximum number of missed LCP echo messages before disconnect.
* `sip` - Start IP.
* `status` - Enable/disable FortiGate as a L2TP gateway. Valid values: `enable` `disable` .
* `usrgrp` - User group. This attribute must reference one of the following datasources: `user.group.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_vpn_l2tp can be imported using:
```sh
terraform import fortios_vpn_l2tp.labelname {{mkey}}
```
