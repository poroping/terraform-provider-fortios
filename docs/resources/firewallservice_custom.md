---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallservice_custom"
description: |-
  Configure custom services.
---

## fortios_firewallservice_custom
Configure custom services.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `app_service_type` - Application service type. Valid values: `disable` `app-id` `app-category` .
* `category` - Service category. This attribute must reference one of the following datasources: `firewall.service.category.name` .
* `check_reset_range` - Configure the type of ICMP error message verification. Valid values: `disable` `strict` `default` .
* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `fabric_object` - Security Fabric global object setting. Valid values: `enable` `disable` .
* `fqdn` - Fully qualified domain name.
* `helper` - Helper name. Valid values: `auto` `disable` `ftp` `tftp` `ras` `h323` `tns` `mms` `sip` `pptp` `rtsp` `dns-udp` `dns-tcp` `pmap` `rsh` `dcerpc` `mgcp` .
* `icmpcode` - ICMP code.
* `icmptype` - ICMP type.
* `iprange` - Start and end of the IP range associated with service.
* `name` - Custom service name.
* `protocol` - Protocol type based on IANA numbers. Valid values: `TCP/UDP/SCTP` `ICMP` `ICMP6` `IP` `HTTP` `FTP` `CONNECT` `SOCKS-TCP` `SOCKS-UDP` `ALL` .
* `protocol_number` - IP protocol number.
* `proxy` - Enable/disable web proxy service. Valid values: `enable` `disable` .
* `sctp_portrange` - Multiple SCTP port ranges.
* `session_ttl` - Session TTL (300 - 2764800, 0 = default).
* `tcp_halfclose_timer` - Wait time to close a TCP session waiting for an unanswered FIN packet (1 - 86400 sec, 0 = default).
* `tcp_halfopen_timer` - Wait time to close a TCP session waiting for an unanswered open session packet (1 - 86400 sec, 0 = default).
* `tcp_portrange` - Multiple TCP port ranges.
* `tcp_rst_timer` - Set the length of the TCP CLOSE state in seconds (5 - 300 sec, 0 = default).
* `tcp_timewait_timer` - Set the length of the TCP TIME-WAIT state in seconds (1 - 300 sec, 0 = default).
* `udp_idle_timer` - Number of seconds before an idle UDP connection times out (0 - 86400 sec, 0 = default).
* `udp_portrange` - Multiple UDP port ranges.
* `visibility` - Enable/disable the visibility of the service on the GUI. Valid values: `enable` `disable` .
* `app_category` - Application category ID. The structure of `app_category` block is documented below.

The `app_category` block contains:

* `id` - Application category id.
* `application` - Application ID. The structure of `application` block is documented below.

The `application` block contains:

* `id` - Application id.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewallservice_custom can be imported using:
```sh
terraform import fortios_firewallservice_custom.labelname {{mkey}}
```
