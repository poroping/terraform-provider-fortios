---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_syslogprofile"
description: |-
  Get information on a fortios Configure Wireless Termination Points (WTP) system log server profile.
---

# Data Source: fortios_wirelesscontroller_syslogprofile
Use this data source to get information on a fortios Configure Wireless Termination Points (WTP) system log server profile.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) WTP system log server profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `log_level` - Lowest level of log messages that FortiAP units send to this server (default = information)
* `name` - WTP system log server profile name.
* `server_addr_type` - Syslog server address type (default = IP)
* `server_fqdn` - FQDN of syslog server that FortiAP units send log messages to.
* `server_ip` - IP address of syslog server that FortiAP units send log messages to.
* `server_port` - Port number of syslog server that FortiAP units send log messages to (default = 514).
* `server_status` - Enable/disable FortiAP units to send log messages to a syslog server (default = enable).
