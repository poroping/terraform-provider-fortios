---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_centralmanagement"
description: |-
  Configure central management.
---

## fortios_system_centralmanagement
Configure central management.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `allow_monitor` - Enable/disable allowing the central management server to remotely monitor this FortiGate Valid values: `enable` `disable` .
* `allow_push_configuration` - Enable/disable allowing the central management server to push configuration changes to this FortiGate. Valid values: `enable` `disable` .
* `allow_push_firmware` - Enable/disable allowing the central management server to push firmware updates to this FortiGate. Valid values: `enable` `disable` .
* `allow_remote_firmware_upgrade` - Enable/disable remotely upgrading the firmware on this FortiGate from the central management server. Valid values: `enable` `disable` .
* `ca_cert` - CA certificate to be used by FGFM protocol.
* `enc_algorithm` - Encryption strength for communications between the FortiGate and central management. Valid values: `default` `high` `low` .
* `fmg` - IP address or FQDN of the FortiManager.
* `fmg_source_ip` - IPv4 source address that this FortiGate uses when communicating with FortiManager.
* `fmg_source_ip6` - IPv6 source address that this FortiGate uses when communicating with FortiManager.
* `fmg_update_port` - Port used to communicate with FortiManager that is acting as a FortiGuard update server. Valid values: `8890` `443` .
* `include_default_servers` - Enable/disable inclusion of public FortiGuard servers in the override server list. Valid values: `enable` `disable` .
* `interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `local_cert` - Certificate to be used by FGFM protocol.
* `mode` - Central management mode. Valid values: `normal` `backup` .
* `schedule_config_restore` - Enable/disable allowing the central management server to restore the configuration of this FortiGate. Valid values: `enable` `disable` .
* `schedule_script_restore` - Enable/disable allowing the central management server to restore the scripts stored on this FortiGate. Valid values: `enable` `disable` .
* `serial_number` - Serial number.
* `type` - Central management type. Valid values: `fortimanager` `fortiguard` `none` .
* `vdom` - Virtual domain (VDOM) name to use when communicating with FortiManager. This attribute must reference one of the following datasources: `system.vdom.name` .
* `server_list` - Additional severs that the FortiGate can use for updates (for AV, IPS, updates) and ratings (for web filter and antispam ratings) servers. The structure of `server_list` block is documented below.

The `server_list` block contains:

* `addr_type` - Indicate whether the FortiGate communicates with the override server using an IPv4 address, an IPv6 address or a FQDN. Valid values: `ipv4` `ipv6` `fqdn` .
* `fqdn` - FQDN address of override server.
* `id` - ID.
* `server_address` - IPv4 address of override server.
* `server_address6` - IPv6 address of override server.
* `server_type` - FortiGuard service type. Valid values: `update` `rating` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_centralmanagement can be imported using:
```sh
terraform import fortios_system_centralmanagement.labelname {{mkey}}
```
