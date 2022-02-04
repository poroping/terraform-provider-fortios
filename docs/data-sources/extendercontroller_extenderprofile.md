---
subcategory: "FortiGate Extender-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_extendercontroller_extenderprofile"
description: |-
  Get information on a fortios FortiExtender extender profile configuration.
---

# Data Source: fortios_extendercontroller_extenderprofile
Use this data source to get information on a fortios FortiExtender extender profile configuration.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) FortiExtender profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `allowaccess` - Control management access to the managed extender. Separate entries with a space.
* `bandwidth_limit` - FortiExtender LAN extension bandwidth limit (Mbps).
* `enforce_bandwidth` - Enable/disable enforcement of bandwidth on LAN extension interface.
* `extension` - Extension option.
* `id` - ID.
* `login_password` - Set the managed extender's administrator password.
* `login_password_change` - Change or reset the administrator password of a managed extender (yes, default, or no, default = no).
* `model` - Model.
* `name` - FortiExtender profile name.
* `cellular` - FortiExtender cellular configuration.The structure of `cellular` block is documented below.

The `cellular` block contains:

* `controller_report` - FortiExtender controller report configuration.The structure of `controller_report` block is documented below.

The `controller_report` block contains:

* `interval` - Controller report interval.
* `signal_threshold` - Controller report signal threshold.
* `status` - FortiExtender controller report status.
* `dataplan` - Dataplan names.The structure of `dataplan` block is documented below.

The `dataplan` block contains:

* `name` - Dataplan name.
* `modem1` - Configuration options for modem 1.The structure of `modem1` block is documented below.

The `modem1` block contains:

* `conn_status` - Connection status.
* `default_sim` - Default SIM selection.
* `gps` - FortiExtender GPS enable/disable.
* `preferred_carrier` - Preferred carrier.
* `redundant_intf` - Redundant interface.
* `redundant_mode` - FortiExtender mode.
* `sim1_pin` - SIM #1 PIN status.
* `sim1_pin_code` - SIM #1 PIN password.
* `sim2_pin` - SIM #2 PIN status.
* `sim2_pin_code` - SIM #2 PIN password.
* `auto_switch` - FortiExtender auto switch configuration.The structure of `auto_switch` block is documented below.

The `auto_switch` block contains:

* `dataplan` - Automatically switch based on data usage.
* `disconnect` - Auto switch by disconnect.
* `disconnect_period` - Automatically switch based on disconnect period.
* `disconnect_threshold` - Automatically switch based on disconnect threshold.
* `signal` - Automatically switch based on signal strength.
* `switch_back` - Auto switch with switch back multi-options.
* `switch_back_time` - Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).
* `switch_back_timer` - Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).
* `modem2` - Configuration options for modem 2.The structure of `modem2` block is documented below.

The `modem2` block contains:

* `conn_status` - Connection status.
* `default_sim` - Default SIM selection.
* `gps` - FortiExtender GPS enable/disable.
* `preferred_carrier` - Preferred carrier.
* `redundant_intf` - Redundant interface.
* `redundant_mode` - FortiExtender mode.
* `sim1_pin` - SIM #1 PIN status.
* `sim1_pin_code` - SIM #1 PIN password.
* `sim2_pin` - SIM #2 PIN status.
* `sim2_pin_code` - SIM #2 PIN password.
* `auto_switch` - FortiExtender auto switch configuration.The structure of `auto_switch` block is documented below.

The `auto_switch` block contains:

* `dataplan` - Automatically switch based on data usage.
* `disconnect` - Auto switch by disconnect.
* `disconnect_period` - Automatically switch based on disconnect period.
* `disconnect_threshold` - Automatically switch based on disconnect threshold.
* `signal` - Automatically switch based on signal strength.
* `switch_back` - Auto switch with switch back multi-options.
* `switch_back_time` - Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).
* `switch_back_timer` - Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).
* `sms_notification` - FortiExtender cellular SMS notification configuration.The structure of `sms_notification` block is documented below.

The `sms_notification` block contains:

* `status` - FortiExtender SMS notification status.
* `alert` - SMS alert list.The structure of `alert` block is documented below.

The `alert` block contains:

* `data_exhausted` - Display string when data exhausted.
* `fgt_backup_mode_switch` - Display string when FortiGate backup mode switched.
* `low_signal_strength` - Display string when signal strength is low.
* `mode_switch` - Display string when mode is switched.
* `os_image_fallback` - Display string when falling back to a previous OS image.
* `session_disconnect` - Display string when session disconnected.
* `system_reboot` - Display string when system rebooted.
* `receiver` - SMS notification receiver list.The structure of `receiver` block is documented below.

The `receiver` block contains:

* `alert` - Alert multi-options.
* `name` - FortiExtender SMS notification receiver name.
* `phone_number` - Receiver phone number. Format: [+][country code][area code][local phone number]. For example, +16501234567.
* `status` - SMS notification receiver status.
* `lan_extension` - FortiExtender lan extension configuration.The structure of `lan_extension` block is documented below.

The `lan_extension` block contains:

* `backhaul_interface` - IPsec phase1 interface.
* `backhaul_ip` - IPsec phase1 IPv4/FQDN. Used to specify the external IP/FQDN when the FortiGate unit is behind a NAT device.
* `ipsec_tunnel` - IPsec tunnel name.
* `link_loadbalance` - LAN extension link load balance strategy.
* `backhaul` - LAN extension backhaul tunnel configuration.The structure of `backhaul` block is documented below.

The `backhaul` block contains:

* `name` - FortiExtender LAN extension backhaul name.
* `port` - FortiExtender uplink port.
* `role` - FortiExtender uplink port.
* `weight` - WRR weight parameter.
