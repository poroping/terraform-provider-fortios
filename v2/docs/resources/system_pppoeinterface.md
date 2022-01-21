---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_pppoeinterface"
description: |-
  Configure the PPPoE interfaces.
---

## fortios_system_pppoeinterface
Configure the PPPoE interfaces.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `ac_name` - PPPoE AC name.
* `auth_type` - PPP authentication type to use. Valid values: `auto` `pap` `chap` `mschapv1` `mschapv2` .
* `device` - Name for the physical interface. This attribute must reference one of the following datasources: `system.interface.name` .
* `dial_on_demand` - Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface. Valid values: `enable` `disable` .
* `disc_retry_timeout` - PPPoE discovery init timeout value in (0-4294967295 sec).
* `idle_timeout` - PPPoE auto disconnect after idle timeout (0-4294967295 sec).
* `ipunnumbered` - PPPoE unnumbered IP.
* `ipv6` - Enable/disable IPv6 Control Protocol (IPv6CP). Valid values: `enable` `disable` .
* `lcp_echo_interval` - Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
* `lcp_max_echo_fails` - Maximum missed LCP echo messages before disconnect.
* `name` - Name of the PPPoE interface.
* `padt_retry_timeout` - PPPoE terminate timeout value in (0-4294967295 sec).
* `password` - Enter the password.
* `pppoe_unnumbered_negotiate` - Enable/disable PPPoE unnumbered negotiation. Valid values: `enable` `disable` .
* `service_name` - PPPoE service name.
* `username` - User name.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_pppoeinterface can be imported using:
```sh
terraform import fortios_system_pppoeinterface.labelname {{mkey}}
```
