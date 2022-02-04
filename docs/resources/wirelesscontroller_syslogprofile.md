---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_syslogprofile"
description: |-
  Configure Wireless Termination Points (WTP) system log server profile.
---

## fortios_wirelesscontroller_syslogprofile
Configure Wireless Termination Points (WTP) system log server profile.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `comment` - Comment.
* `log_level` - Lowest level of log messages that FortiAP units send to this server (default = information). Valid values: `emergency` `alert` `critical` `error` `warning` `notification` `information` `debugging` .
* `name` - WTP system log server profile name.
* `server_addr_type` - Syslog server address type (default = ip). Valid values: `fqdn` `ip` .
* `server_fqdn` - FQDN of syslog server that FortiAP units send log messages to.
* `server_ip` - IP address of syslog server that FortiAP units send log messages to.
* `server_port` - Port number of syslog server that FortiAP units send log messages to (default = 514).
* `server_status` - Enable/disable FortiAP units to send log messages to a syslog server (default = enable). Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_syslogprofile can be imported using:
```sh
terraform import fortios_wirelesscontroller_syslogprofile.labelname {{mkey}}
```
