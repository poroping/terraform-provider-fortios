---
subcategory: "FortiGate Extender-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_extendercontroller_extender"
description: |-
  Extender controller configuration.
---

## fortios_extendercontroller_extender
Extender controller configuration.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `aaa_shared_secret` - AAA shared secret.
* `access_point_name` - Access point name(APN).
* `admin` - FortiExtender Administration (enable or disable). Valid values: `disable` `discovered` `enable` .
* `allowaccess` - Control management access to the managed extender. Separate entries with a space. Valid values: `ping` `telnet` `http` `https` `ssh` `snmp` .
* `at_dial_script` - Initialization AT commands specific to the MODEM.
* `authorized` - FortiExtender Administration (enable or disable). Valid values: `disable` `enable` .
* `bandwidth_limit` - FortiExtender LAN extension bandwidth limit (Mbps).
* `billing_start_day` - Billing start day.
* `cdma_aaa_spi` - CDMA AAA SPI.
* `cdma_ha_spi` - CDMA HA SPI.
* `cdma_nai` - NAI for CDMA MODEMS.
* `conn_status` - Connection status.
* `description` - Description.
* `device_id` - Device ID.
* `dial_mode` - Dial mode (dial-on-demand or always-connect). Valid values: `dial-on-demand` `always-connect` .
* `dial_status` - Dial status.
* `enforce_bandwidth` - Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable` `disable` .
* `ext_name` - FortiExtender name.
* `extension_type` - Extension type for this FortiExtender. Valid values: `lan-extension` .
* `ha_shared_secret` - HA shared secret.
* `id` - FortiExtender serial number.
* `ifname` - FortiExtender interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `initiated_update` - Allow/disallow network initiated updates to the MODEM. Valid values: `enable` `disable` .
* `login_password` - Set the managed extender's administrator password.
* `login_password_change` - Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes` `default` `no` .
* `mode` - FortiExtender mode. Valid values: `standalone` `redundant` .
* `modem_passwd` - MODEM password.
* `modem_type` - MODEM type (CDMA, GSM/LTE or WIMAX). Valid values: `cdma` `gsm/lte` `wimax` .
* `multi_mode` - MODEM mode of operation(3G,LTE,etc). Valid values: `auto` `auto-3g` `force-lte` `force-3g` `force-2g` .
* `name` - FortiExtender entry name.
* `override_allowaccess` - Enable to override the extender profile management access configuration. Valid values: `enable` `disable` .
* `override_enforce_bandwidth` - Enable to override the extender profile enforce-bandwidth setting. Valid values: `enable` `disable` .
* `override_login_password_change` - Enable to override the extender profile login-password (administrator password) setting. Valid values: `enable` `disable` .
* `ppp_auth_protocol` - PPP authentication protocol (PAP,CHAP or auto). Valid values: `auto` `pap` `chap` .
* `ppp_echo_request` - Enable/disable PPP echo request. Valid values: `enable` `disable` .
* `ppp_password` - PPP password.
* `ppp_username` - PPP username.
* `primary_ha` - Primary HA.
* `profile` - FortiExtender profile configuration. This attribute must reference one of the following datasources: `extender-controller.extender-profile.name` .
* `quota_limit_mb` - Monthly quota limit (MB).
* `redial` - Number of redials allowed based on failed attempts. Valid values: `none` `1` `2` `3` `4` `5` `6` `7` `8` `9` `10` .
* `redundant_intf` - Redundant interface.
* `roaming` - Enable/disable MODEM roaming. Valid values: `enable` `disable` .
* `role` - FortiExtender work role(Primary, Secondary, None). Valid values: `none` `primary` `secondary` .
* `secondary_ha` - Secondary HA.
* `sim_pin` - SIM PIN.
* `vdom` - VDOM.
* `wimax_auth_protocol` - WiMax authentication protocol(TLS or TTLS). Valid values: `tls` `ttls` .
* `wimax_carrier` - WiMax carrier.
* `wimax_realm` - WiMax realm.
* `controller_report` - FortiExtender controller report configuration. The structure of `controller_report` block is documented below.

The `controller_report` block contains:

* `interval` - Controller report interval.
* `signal_threshold` - Controller report signal threshold.
* `status` - FortiExtender controller report status. Valid values: `disable` `enable` .
* `modem1` - Configuration options for modem 1. The structure of `modem1` block is documented below.

The `modem1` block contains:

* `conn_status` - Connection status.
* `default_sim` - Default SIM selection. Valid values: `sim1` `sim2` `carrier` `cost` .
* `gps` - FortiExtender GPS enable/disable. Valid values: `disable` `enable` .
* `ifname` - FortiExtender interface name. This attribute must reference one of the following datasources: `system.interface.name` .
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
* `ifname` - FortiExtender interface name. This attribute must reference one of the following datasources: `system.interface.name` .
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
* `wan_extension` - FortiExtender wan extension configuration. The structure of `wan_extension` block is documented below.

The `wan_extension` block contains:

* `modem1_extension` - FortiExtender interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `modem2_extension` - FortiExtender interface name. This attribute must reference one of the following datasources: `system.interface.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_extendercontroller_extender can be imported using:
```sh
terraform import fortios_extendercontroller_extender.labelname {{mkey}}
```
