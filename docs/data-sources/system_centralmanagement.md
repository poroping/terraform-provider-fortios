---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_centralmanagement"
description: |-
  Get information on a fortios Configure central management.
---

# Data Source: fortios_system_centralmanagement
Use this data source to get information on a fortios Configure central management.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `allow_monitor` - Enable/disable allowing the central management server to remotely monitor this FortiGate unit.
* `allow_push_configuration` - Enable/disable allowing the central management server to push configuration changes to this FortiGate.
* `allow_push_firmware` - Enable/disable allowing the central management server to push firmware updates to this FortiGate.
* `allow_remote_firmware_upgrade` - Enable/disable remotely upgrading the firmware on this FortiGate from the central management server.
* `ca_cert` - CA certificate to be used by FGFM protocol.
* `enc_algorithm` - Encryption strength for communications between the FortiGate and central management.
* `fmg` - IP address or FQDN of the FortiManager.
* `fmg_source_ip` - IPv4 source address that this FortiGate uses when communicating with FortiManager.
* `fmg_source_ip6` - IPv6 source address that this FortiGate uses when communicating with FortiManager.
* `fmg_update_port` - Port used to communicate with FortiManager that is acting as a FortiGuard update server.
* `fortigate_cloud_sso_default_profile` - Override access profile.
* `include_default_servers` - Enable/disable inclusion of public FortiGuard servers in the override server list.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `local_cert` - Certificate to be used by FGFM protocol.
* `mode` - Central management mode.
* `schedule_config_restore` - Enable/disable allowing the central management server to restore the configuration of this FortiGate.
* `schedule_script_restore` - Enable/disable allowing the central management server to restore the scripts stored on this FortiGate.
* `serial_number` - Serial number.
* `type` - Central management type.
* `vdom` - Virtual domain (VDOM) name to use when communicating with FortiManager.
* `server_list` - Additional severs that the FortiGate can use for updates (for AV, IPS, updates) and ratings (for web filter and antispam ratings) servers.The structure of `server_list` block is documented below.

The `server_list` block contains:

* `addr_type` - Indicate whether the FortiGate communicates with the override server using an IPv4 address, an IPv6 address or a FQDN.
* `fqdn` - FQDN address of override server.
* `id` - ID.
* `server_address` - IPv4 address of override server.
* `server_address6` - IPv6 address of override server.
* `server_type` - FortiGuard service type.
