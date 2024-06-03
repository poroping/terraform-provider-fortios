---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallservice_custom"
description: |-
  Get information on a fortios Configure custom services.
---

# Data Source: fortios_firewallservice_custom
Use this data source to get information on a fortios Configure custom services.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Custom service name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `app_service_type` - Application service type.
* `category` - Service category.
* `check_reset_range` - Configure the type of ICMP error message verification.
* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `fabric_object` - Security Fabric global object setting.
* `fqdn` - Fully qualified domain name.
* `helper` - Helper name.
* `icmpcode` - ICMP code.
* `icmptype` - ICMP type.
* `iprange` - Start and end of the IP range associated with service.
* `name` - Custom service name.
* `protocol` - Protocol type based on IANA numbers.
* `protocol_number` - IP protocol number.
* `proxy` - Enable/disable web proxy service.
* `sctp_portrange` - Multiple SCTP port ranges.
* `session_ttl` - Session TTL (300 - 2764800, 0 = default).
* `tcp_halfclose_timer` - Wait time to close a TCP session waiting for an unanswered FIN packet (1 - 86400 sec, 0 = default).
* `tcp_halfopen_timer` - Wait time to close a TCP session waiting for an unanswered open session packet (1 - 86400 sec, 0 = default).
* `tcp_portrange` - Multiple TCP port ranges.
* `tcp_rst_timer` - Set the length of the TCP CLOSE state in seconds (5 - 300 sec, 0 = default).
* `tcp_timewait_timer` - Set the length of the TCP TIME-WAIT state in seconds (1 - 300 sec, 0 = default).
* `udp_idle_timer` - Number of seconds before an idle UDP connection times out (0 - 86400 sec, 0 = default).
* `udp_portrange` - Multiple UDP port ranges.
* `visibility` - Enable/disable the visibility of the service on the GUI.
* `app_category` - Application category ID.The structure of `app_category` block is documented below.

The `app_category` block contains:

* `id` - Application category id.
* `application` - Application ID.The structure of `application` block is documented below.

The `application` block contains:

* `id` - Application id.
