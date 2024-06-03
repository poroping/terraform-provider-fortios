---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_modem"
description: |-
  Configure MODEM.
---

## fortios_system_modem
Configure MODEM.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `action` - Dial up/stop MODEM. Valid values: `dial` `stop` `none` .
* `altmode` - Enable/disable altmode for installations using PPP in China. Valid values: `enable` `disable` .
* `authtype1` - Allowed authentication types for ISP 1. Valid values: `pap` `chap` `mschap` `mschapv2` .
* `authtype2` - Allowed authentication types for ISP 2. Valid values: `pap` `chap` `mschap` `mschapv2` .
* `authtype3` - Allowed authentication types for ISP 3. Valid values: `pap` `chap` `mschap` `mschapv2` .
* `auto_dial` - Enable/disable auto-dial after a reboot or disconnection. Valid values: `enable` `disable` .
* `connect_timeout` - Connection completion timeout (30 - 255 sec, default = 90).
* `dial_cmd1` - Dial command (this is often an ATD or ATDT command).
* `dial_cmd2` - Dial command (this is often an ATD or ATDT command).
* `dial_cmd3` - Dial command (this is often an ATD or ATDT command).
* `dial_on_demand` - Enable/disable to dial the modem when packets are routed to the modem interface. Valid values: `enable` `disable` .
* `distance` - Distance of learned routes (1 - 255, default = 1).
* `dont_send_cr1` - Do not send CR when connected (ISP1). Valid values: `enable` `disable` .
* `dont_send_cr2` - Do not send CR when connected (ISP2). Valid values: `enable` `disable` .
* `dont_send_cr3` - Do not send CR when connected (ISP3). Valid values: `enable` `disable` .
* `extra_init1` - Extra initialization string to ISP 1.
* `extra_init2` - Extra initialization string to ISP 2.
* `extra_init3` - Extra initialization string to ISP 3.
* `holddown_timer` - Hold down timer in seconds (1 - 60 sec).
* `idle_timer` - MODEM connection idle time (1 - 9999 min, default = 5).
* `interface` - Name of redundant interface. This attribute must reference one of the following datasources: `system.interface.name` .
* `lockdown_lac` - Allow connection only to the specified Location Area Code (LAC).
* `mode` - Set MODEM operation mode to redundant or standalone. Valid values: `standalone` `redundant` .
* `network_init` - AT command to set the Network name/type (AT+COPS=<mode>,[<format>,<oper>[,<AcT>]]).
* `passwd1` - Password to access the specified dialup account.
* `passwd2` - Password to access the specified dialup account.
* `passwd3` - Password to access the specified dialup account.
* `peer_modem1` - Specify peer MODEM type for phone1. Valid values: `generic` `actiontec` `ascend_TNT` .
* `peer_modem2` - Specify peer MODEM type for phone2. Valid values: `generic` `actiontec` `ascend_TNT` .
* `peer_modem3` - Specify peer MODEM type for phone3. Valid values: `generic` `actiontec` `ascend_TNT` .
* `phone1` - Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
* `phone2` - Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
* `phone3` - Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
* `pin_init` - AT command to set the PIN (AT+PIN=<pin>).
* `ppp_echo_request1` - Enable/disable PPP echo-request to ISP 1. Valid values: `enable` `disable` .
* `ppp_echo_request2` - Enable/disable PPP echo-request to ISP 2. Valid values: `enable` `disable` .
* `ppp_echo_request3` - Enable/disable PPP echo-request to ISP 3. Valid values: `enable` `disable` .
* `priority` - Priority of learned routes (0 - 4294967295, default = 0).
* `redial` - Redial limit (1 - 10 attempts, none = redial forever). Valid values: `none` `1` `2` `3` `4` `5` `6` `7` `8` `9` `10` .
* `reset` - Number of dial attempts before resetting modem (0 = never reset).
* `status` - Enable/disable Modem support (equivalent to bringing an interface up or down). Valid values: `enable` `disable` .
* `traffic_check` - Enable/disable traffic-check. Valid values: `enable` `disable` .
* `username1` - User name to access the specified dialup account.
* `username2` - User name to access the specified dialup account.
* `username3` - User name to access the specified dialup account.
* `wireless_port` - Enter wireless port number: 0 for default, 1 for first port, and so on (0 - 4294967295).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_modem can be imported using:
```sh
terraform import fortios_system_modem.labelname {{mkey}}
```
