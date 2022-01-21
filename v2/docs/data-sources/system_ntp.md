---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ntp"
description: |-
  Get information on a fortios Configure system NTP information.
---

# Data Source: fortios_system_ntp
Use this data source to get information on a fortios Configure system NTP information.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `authentication` - Enable/disable authentication.
* `key` - Key for authentication.
* `key_id` - Key ID for authentication.
* `key_type` - Key type for authentication (MD5, SHA1).
* `ntpsync` - Enable/disable setting the FortiGate system time by synchronizing with an NTP Server.
* `server_mode` - Enable/disable FortiGate NTP Server Mode. Your FortiGate becomes an NTP server for other devices on your network. The FortiGate relays NTP requests to its configured NTP server.
* `source_ip` - Source IP address for communication to the NTP server.
* `source_ip6` - Source IPv6 address for communication to the NTP server.
* `syncinterval` - NTP synchronization interval (1 - 1440 min).
* `type` - Use the FortiGuard NTP server or any other available NTP Server.
* `interface` - FortiGate interface(s) with NTP server mode enabled. Devices on your network can contact these interfaces for NTP services.The structure of `interface` block is documented below.

The `interface` block contains:

* `interface_name` - Interface name.
* `ntpserver` - Configure the FortiGate to connect to any available third-party NTP server.The structure of `ntpserver` block is documented below.

The `ntpserver` block contains:

* `authentication` - Enable/disable MD5(NTPv3)/SHA1(NTPv4) authentication.
* `id` - NTP server ID.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `key` - Key for MD5(NTPv3)/SHA1(NTPv4) authentication.
* `key_id` - Key ID for authentication.
* `ntpv3` - Enable to use NTPv3 instead of NTPv4.
* `server` - IP address or hostname of the NTP Server.
