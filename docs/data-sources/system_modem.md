---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_modem"
description: |-
  Get information on a fortios Configure MODEM.
---

# Data Source: fortios_system_modem
Use this data source to get information on a fortios Configure MODEM.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `action` - Dial up/stop MODEM.
* `altmode` - Enable/disable altmode for installations using PPP in China.
* `authtype1` - Allowed authentication types for ISP 1.
* `authtype2` - Allowed authentication types for ISP 2.
* `authtype3` - Allowed authentication types for ISP 3.
* `auto_dial` - Enable/disable auto-dial after a reboot or disconnection.
* `connect_timeout` - Connection completion timeout (30 - 255 sec, default = 90).
* `dial_cmd1` - Dial command (this is often an ATD or ATDT command).
* `dial_cmd2` - Dial command (this is often an ATD or ATDT command).
* `dial_cmd3` - Dial command (this is often an ATD or ATDT command).
* `dial_on_demand` - Enable/disable to dial the modem when packets are routed to the modem interface.
* `distance` - Distance of learned routes (1 - 255, default = 1).
* `dont_send_cr1` - Do not send CR when connected (ISP1).
* `dont_send_cr2` - Do not send CR when connected (ISP2).
* `dont_send_cr3` - Do not send CR when connected (ISP3).
* `extra_init1` - Extra initialization string to ISP 1.
* `extra_init2` - Extra initialization string to ISP 2.
* `extra_init3` - Extra initialization string to ISP 3.
* `holddown_timer` - Hold down timer in seconds (1 - 60 sec).
* `idle_timer` - MODEM connection idle time (1 - 9999 min, default = 5).
* `interface` - Name of redundant interface.
* `lockdown_lac` - Allow connection only to the specified Location Area Code (LAC).
* `mode` - Set MODEM operation mode to redundant or standalone.
* `network_init` - AT command to set the Network name/type (AT+COPS=<mode>,[<format>,<oper>[,<AcT>]]).
* `passwd1` - Password to access the specified dialup account.
* `passwd2` - Password to access the specified dialup account.
* `passwd3` - Password to access the specified dialup account.
* `peer_modem1` - Specify peer MODEM type for phone1.
* `peer_modem2` - Specify peer MODEM type for phone2.
* `peer_modem3` - Specify peer MODEM type for phone3.
* `phone1` - Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
* `phone2` - Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
* `phone3` - Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
* `pin_init` - AT command to set the PIN (AT+PIN=<pin>).
* `ppp_echo_request1` - Enable/disable PPP echo-request to ISP 1.
* `ppp_echo_request2` - Enable/disable PPP echo-request to ISP 2.
* `ppp_echo_request3` - Enable/disable PPP echo-request to ISP 3.
* `priority` - Priority of learned routes (0 - 4294967295, default = 0).
* `redial` - Redial limit (1 - 10 attempts, none = redial forever).
* `reset` - Number of dial attempts before resetting modem (0 = never reset).
* `status` - Enable/disable Modem support (equivalent to bringing an interface up or down).
* `traffic_check` - Enable/disable traffic-check.
* `username1` - User name to access the specified dialup account.
* `username2` - User name to access the specified dialup account.
* `username3` - User name to access the specified dialup account.
* `wireless_port` - Enter wireless port number: 0 for default, 1 for first port, and so on (0 - 4294967295).
