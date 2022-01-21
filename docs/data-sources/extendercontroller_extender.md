---
subcategory: "FortiGate Extender-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_extendercontroller_extender"
description: |-
  Get information on a fortios Extender controller configuration.
---

# Data Source: fortios_extendercontroller_extender
Use this data source to get information on a fortios Extender controller configuration.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) FortiExtender entry name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `aaa_shared_secret` - AAA shared secret.
* `access_point_name` - Access point name(APN).
* `admin` - FortiExtender Administration (enable or disable).
* `allowaccess` - Control management access to the managed extender. Separate entries with a space.
* `at_dial_script` - Initialization AT commands specific to the MODEM.
* `authorized` - FortiExtender Administration (enable or disable).
* `bandwidth_limit` - FortiExtender LAN extension bandwidth limit (Mbps).
* `billing_start_day` - Billing start day.
* `cdma_aaa_spi` - CDMA AAA SPI.
* `cdma_ha_spi` - CDMA HA SPI.
* `cdma_nai` - NAI for CDMA MODEMS.
* `conn_status` - Connection status.
* `description` - Description.
* `device_id` - device-id
* `dial_mode` - Dial mode (dial-on-demand or always-connect).
* `dial_status` - Dial status.
* `enforce_bandwidth` - Enable/disable enforcement of bandwidth on LAN extension interface.
* `ext_name` - FortiExtender name.
* `extension_type` - Extension type for this FortiExtender.
* `ha_shared_secret` - HA shared secret.
* `id` - FortiExtender serial number.
* `ifname` - FortiExtender interface name.
* `initiated_update` - Allow/disallow network initiated updates to the MODEM.
* `login_password` - Set the managed extender's administrator password.
* `login_password_change` - Change or reset the administrator password of a managed extender (yes, default, or no, default = no).
* `mode` - FortiExtender mode.
* `modem_passwd` - MODEM password.
* `modem_type` - MODEM type (CDMA, GSM/LTE or WIMAX).
* `multi_mode` - MODEM mode of operation(3G,LTE,etc).
* `name` - FortiExtender entry name.
* `override_allowaccess` - Enable to override the extender profile management access configuration.
* `override_enforce_bandwidth` - Enable to override the extender profile enforce-bandwidth setting.
* `override_login_password_change` - Enable to override the extender profile login-password (administrator password) setting.
* `ppp_auth_protocol` - PPP authentication protocol (PAP,CHAP or auto).
* `ppp_echo_request` - Enable/disable PPP echo request.
* `ppp_password` - PPP password.
* `ppp_username` - PPP username.
* `primary_ha` - Primary HA.
* `profile` - FortiExtender profile configuration.
* `quota_limit_mb` - Monthly quota limit (MB).
* `redial` - Number of redials allowed based on failed attempts.
* `redundant_intf` - Redundant interface.
* `roaming` - Enable/disable MODEM roaming.
* `role` - FortiExtender work role(Primary, Secondary, None).
* `secondary_ha` - Secondary HA.
* `sim_pin` - SIM PIN.
* `vdom` - VDOM
* `wimax_auth_protocol` - WiMax authentication protocol(TLS or TTLS).
* `wimax_carrier` - WiMax carrier.
* `wimax_realm` - WiMax realm.
* `controller_report` - FortiExtender controller report configuration.The structure of `controller_report` block is documented below.

The `controller_report` block contains:

* `interval` - Controller report interval.
* `signal_threshold` - Controller report signal threshold.
* `status` - FortiExtender controller report status.
* `modem1` - Configuration options for modem 1.The structure of `modem1` block is documented below.

The `modem1` block contains:

* `conn_status` - Connection status.
* `default_sim` - Default SIM selection.
* `gps` - FortiExtender GPS enable/disable.
* `ifname` - FortiExtender interface name.
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
* `ifname` - FortiExtender interface name.
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
* `wan_extension` - FortiExtender wan extension configuration.The structure of `wan_extension` block is documented below.

The `wan_extension` block contains:

* `modem1_extension` - FortiExtender interface name.
* `modem2_extension` - FortiExtender interface name.
