---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomproperty"
description: |-
  Configure VDOM property.
---

## fortios_system_vdomproperty
Configure VDOM property.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `custom_service` - Maximum guaranteed number of firewall custom services.
* `description` - Description.
* `dialup_tunnel` - Maximum guaranteed number of dial-up tunnels.
* `firewall_address` - Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
* `firewall_addrgrp` - Maximum guaranteed number of firewall address groups (IPv4, IPv6).
* `firewall_policy` - Maximum guaranteed number of firewall policies (policy, DoS-policy4, DoS-policy6, multicast).
* `ipsec_phase1` - Maximum guaranteed number of VPN IPsec phase 1 tunnels.
* `ipsec_phase1_interface` - Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
* `ipsec_phase2` - Maximum guaranteed number of VPN IPsec phase 2 tunnels.
* `ipsec_phase2_interface` - Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
* `log_disk_quota` - Log disk quota in MiB (range depends on how much disk space is available).
* `name` - VDOM name. This attribute must reference one of the following datasources: `system.vdom.name` .
* `onetime_schedule` - Maximum guaranteed number of firewall one-time schedules.
* `proxy` - Maximum guaranteed number of concurrent proxy users.
* `recurring_schedule` - Maximum guaranteed number of firewall recurring schedules.
* `service_group` - Maximum guaranteed number of firewall service groups.
* `session` - Maximum guaranteed number of sessions.
* `snmp_index` - Permanent SNMP Index of the virtual domain (1 - 2147483647).
* `sslvpn` - Maximum guaranteed number of SSL-VPNs.
* `user` - Maximum guaranteed number of local users.
* `user_group` - Maximum guaranteed number of user groups.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_vdomproperty can be imported using:
```sh
terraform import fortios_system_vdomproperty.labelname {{mkey}}
```
