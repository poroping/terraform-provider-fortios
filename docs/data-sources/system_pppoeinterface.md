---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_pppoeinterface"
description: |-
  Get information on a fortios Configure the PPPoE interfaces.
---

# Data Source: fortios_system_pppoeinterface
Use this data source to get information on a fortios Configure the PPPoE interfaces.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name of the PPPoE interface.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `ac_name` - PPPoE AC name.
* `auth_type` - PPP authentication type to use.
* `device` - Name for the physical interface.
* `dial_on_demand` - Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface.
* `disc_retry_timeout` - PPPoE discovery init timeout value in (0-4294967295 sec).
* `idle_timeout` - PPPoE auto disconnect after idle timeout (0-4294967295 sec).
* `ipunnumbered` - PPPoE unnumbered IP.
* `ipv6` - Enable/disable IPv6 Control Protocol (IPv6CP).
* `lcp_echo_interval` - Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
* `lcp_max_echo_fails` - Maximum missed LCP echo messages before disconnect.
* `name` - Name of the PPPoE interface.
* `padt_retry_timeout` - PPPoE terminate timeout value in (0-4294967295 sec).
* `password` - Enter the password.
* `pppoe_unnumbered_negotiate` - Enable/disable PPPoE unnumbered negotiation.
* `service_name` - PPPoE service name.
* `username` - User name.
