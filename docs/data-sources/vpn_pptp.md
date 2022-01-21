---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpn_pptp"
description: |-
  Get information on a fortios Configure PPTP.
---

# Data Source: fortios_vpn_pptp
Use this data source to get information on a fortios Configure PPTP.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `eip` - End IP.
* `ip_mode` - IP assignment mode for PPTP client.
* `local_ip` - Local IP to be used for peer's remote IP.
* `sip` - Start IP.
* `status` - Enable/disable FortiGate as a PPTP gateway.
* `usrgrp` - User group.
