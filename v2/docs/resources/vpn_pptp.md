---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpn_pptp"
description: |-
  Configure PPTP.
---

## fortios_vpn_pptp
Configure PPTP.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `eip` - End IP.
* `ip_mode` - IP assignment mode for PPTP client. Valid values: `range` `usrgrp` .
* `local_ip` - Local IP to be used for peer's remote IP.
* `sip` - Start IP.
* `status` - Enable/disable FortiGate as a PPTP gateway. Valid values: `enable` `disable` .
* `usrgrp` - User group. This attribute must reference one of the following datasources: `user.group.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_vpn_pptp can be imported using:
```sh
terraform import fortios_vpn_pptp.labelname {{mkey}}
```
