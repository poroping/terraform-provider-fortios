---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_resourcelimits"
description: |-
  Configure resource limits.
---

## fortios_system_resourcelimits
Configure resource limits.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `custom_service` - Maximum number of firewall custom services.
* `dialup_tunnel` - Maximum number of dial-up tunnels.
* `firewall_address` - Maximum number of firewall addresses (IPv4, IPv6, multicast).
* `firewall_addrgrp` - Maximum number of firewall address groups (IPv4, IPv6).
* `firewall_policy` - Maximum number of firewall policies (policy, DoS-policy4, DoS-policy6, multicast).
* `ipsec_phase1` - Maximum number of VPN IPsec phase1 tunnels.
* `ipsec_phase1_interface` - Maximum number of VPN IPsec phase1 interface tunnels.
* `ipsec_phase2` - Maximum number of VPN IPsec phase2 tunnels.
* `ipsec_phase2_interface` - Maximum number of VPN IPsec phase2 interface tunnels.
* `log_disk_quota` - Log disk quota in megabytes (MB).
* `onetime_schedule` - Maximum number of firewall one-time schedules.
* `proxy` - Maximum number of concurrent proxy users.
* `recurring_schedule` - Maximum number of firewall recurring schedules.
* `service_group` - Maximum number of firewall service groups.
* `session` - Maximum number of sessions.
* `sslvpn` - Maximum number of SSL-VPN.
* `user` - Maximum number of local users.
* `user_group` - Maximum number of user groups.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_resourcelimits can be imported using:
```sh
terraform import fortios_system_resourcelimits.labelname {{mkey}}
```
