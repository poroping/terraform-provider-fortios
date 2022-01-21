---
subcategory: "FortiGate Extender-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_extendercontroller_extenderprofile"
description: |-
  FortiExtender extender profile configuration.
---

## fortios_extendercontroller_extenderprofile
FortiExtender extender profile configuration.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `allowaccess` - Control management access to the managed extender. Separate entries with a space. Valid values: `ping` `telnet` `http` `https` `ssh` `snmp` .
* `bandwidth_limit` - FortiExtender LAN extension bandwidth limit (Mbps).
* `enforce_bandwidth` - Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable` `disable` .
* `extension` - Extension option. Valid values: `wan-extension` `lan-extension` .
* `id` - id
* `login_password` - Set the managed extender's administrator password.
* `login_password_change` - Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes` `default` `no` .
* `model` - Model. Valid values: `FX201E` `FX211E` `FX200F` `FXA11F` `FXE11F` `FXA21F` `FXE21F` `FXA22F` `FXE22F` `FX212F` `FX311F` `FX312F` `FX511F` `FVG21F` `FVA21F` `FVG22F` `FVA22F` `FX04DA` .
* `name` - FortiExtender profile name
* `cellular` - FortiExtender cellular configuration. The structure of `cellular` block is documented below.

The `cellular` block contains:

* `controller_report` - FortiExtender controller report configuration. The structure of `controller_report` block is documented below.

The `controller_report` block contains:

* `interval` - Controller report interval.
* `signal_threshold` - Controller report signal threshold.
* `status` - FortiExtender controller report status. Valid values: `disable` `enable` .
* `dataplan` - Dataplan names. The structure of `dataplan` block is documented below.

The `dataplan` block contains:

* `name` - Dataplan name. This attribute must reference one of the following datasources: `extender-controller.dataplan.name` .
* `modem1` - Configuration options for modem 1. The structure of `modem1` block is documented below.

The `modem1` block contains:

* `conn_status` - Connection status.
* `default_sim` - Default SIM selection. Valid values: `sim1` `sim2` `carrier` `cost` .
* `gps` - FortiExtender GPS enable/disable. Valid values: `disable` `enable` .
* `preferred_carrier` - Preferred carrier.
* `redundant_intf` - Redundant interface.
* `redundant_mode` - FortiExtender mode. Valid values: `disable` `enable` .
* `sim1_pin` - SIM #1 PIN status. Valid values: `disable` `enable` .
* `sim1_pin_code` - SIM #1 PIN password.
* `sim2_pin` - SIM #2 PIN status. Valid values: `disable` `enable` .
* `sim2_pin_code` - SIM #2 PIN password.
* `auto_switch` - FortiExtender auto switch configuration. The structure of `auto_switch` block is documented below.

The `auto_switch` block contains:

* `dataplan` - Automatically switch based on data usage. Valid values: `disable` `enable` .
* `disconnect` - Auto switch by disconnect. Valid values: `disable` `enable` .
* `disconnect_period` - Automatically switch based on disconnect period.
* `disconnect_threshold` - Automatically switch based on disconnect threshold.
* `signal` - Automatically switch based on signal strength. Valid values: `disable` `enable` .
* `switch_back` - Auto switch with switch back multi-options. Valid values: `time` `timer` .
* `switch_back_time` - Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).
* `switch_back_timer` - Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).
* `modem2` - Configuration options for modem 2. The structure of `modem2` block is documented below.

The `modem2` block contains:

* `conn_status` - Connection status.
* `default_sim` - Default SIM selection. Valid values: `sim1` `sim2` `carrier` `cost` .
* `gps` - FortiExtender GPS enable/disable. Valid values: `disable` `enable` .
* `preferred_carrier` - Preferred carrier.
* `redundant_intf` - Redundant interface.
* `redundant_mode` - FortiExtender mode. Valid values: `disable` `enable` .
* `sim1_pin` - SIM #1 PIN status. Valid values: `disable` `enable` .
* `sim1_pin_code` - SIM #1 PIN password.
* `sim2_pin` - SIM #2 PIN status. Valid values: `disable` `enable` .
* `sim2_pin_code` - SIM #2 PIN password.
* `auto_switch` - FortiExtender auto switch configuration. The structure of `auto_switch` block is documented below.

The `auto_switch` block contains:

* `dataplan` - Automatically switch based on data usage. Valid values: `disable` `enable` .
* `disconnect` - Auto switch by disconnect. Valid values: `disable` `enable` .
* `disconnect_period` - Automatically switch based on disconnect period.
* `disconnect_threshold` - Automatically switch based on disconnect threshold.
* `signal` - Automatically switch based on signal strength. Valid values: `disable` `enable` .
* `switch_back` - Auto switch with switch back multi-options. Valid values: `time` `timer` .
* `switch_back_time` - Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).
* `switch_back_timer` - Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).
* `sms_notification` - FortiExtender cellular SMS notification configuration. The structure of `sms_notification` block is documented below.

The `sms_notification` block contains:

* `status` - FortiExtender SMS notification status. Valid values: `disable` `enable` .
* `alert` - SMS alert list. The structure of `alert` block is documented below.

The `alert` block contains:

* `data_exhausted` - Display string when data exhausted.
* `fgt_backup_mode_switch` - Display string when FortiGate backup mode switched.
* `low_signal_strength` - Display string when signal strength is low.
* `mode_switch` - Display string when mode is switched.
* `os_image_fallback` - Display string when falling back to a previous OS image.
* `session_disconnect` - Display string when session disconnected.
* `system_reboot` - Display string when system rebooted.
* `receiver` - SMS notification receiver list. The structure of `receiver` block is documented below.

The `receiver` block contains:

* `alert` - Alert multi-options. Valid values: `system-reboot` `data-exhausted` `session-disconnect` `low-signal-strength` `mode-switch` `os-image-fallback` `fgt-backup-mode-switch` .
* `name` - FortiExtender SMS notification receiver name.
* `phone_number` - Receiver phone number.  Format: [+][country code][area code][local phone number].  For example: +16501234567.
* `status` - SMS notification receiver status. Valid values: `disable` `enable` .
* `lan_extension` - FortiExtender lan extension configuration. The structure of `lan_extension` block is documented below.

The `lan_extension` block contains:

* `backhaul_interface` - IPsec phase1 interface. This attribute must reference one of the following datasources: `system.interface.name` .
* `backhaul_ip` - IPsec phase1 IPv4/FQDN. Used to specify the external IP/FQDN when the FortiGate unit is behind a NAT device.
* `ipsec_tunnel` - IPsec tunnel name.
* `link_loadbalance` - LAN extension link load balance strategy. Valid values: `activebackup` `loadbalance` .
* `backhaul` - LAN extension backhaul tunnel configuration. The structure of `backhaul` block is documented below.

The `backhaul` block contains:

* `name` - FortiExtender LAN extension backhaul name
* `port` - FortiExtender uplink port. Valid values: `wan` `lte1` `lte2` `port1` `port2` `port3` `port4` `port5` `sfp` .
* `role` - FortiExtender uplink port. Valid values: `primary` `secondary` .
* `weight` - WRR weight parameter

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_extendercontroller_extenderprofile can be imported using:
```sh
terraform import fortios_extendercontroller_extenderprofile.labelname {{mkey}}
```
