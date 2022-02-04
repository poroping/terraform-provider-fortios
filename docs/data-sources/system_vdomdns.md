---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomdns"
description: |-
  Get information on a fortios Configure DNS servers for a non-management VDOM.
---

# Data Source: fortios_system_vdomdns
Use this data source to get information on a fortios Configure DNS servers for a non-management VDOM.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `alt_primary` - Alternate primary DNS server. This is not used as a failover DNS server.
* `alt_secondary` - Alternate secondary DNS server. This is not used as a failover DNS server.
* `dns_over_tls` - Enable/disable/enforce DNS over TLS.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `ip6_primary` - Primary IPv6 DNS server IP address for the VDOM.
* `ip6_secondary` - Secondary IPv6 DNS server IP address for the VDOM.
* `primary` - Primary DNS server IP address for the VDOM.
* `protocol` - DNS transport protocols.
* `secondary` - Secondary DNS server IP address for the VDOM.
* `server_select_method` - Specify how configured servers are prioritized.
* `source_ip` - Source IP for communications with the DNS server.
* `ssl_certificate` - Name of local certificate for SSL connections.
* `vdom_dns` - Enable/disable configuring DNS servers for the current VDOM.
* `server_hostname` - DNS server host name list.The structure of `server_hostname` block is documented below.

The `server_hostname` block contains:

* `hostname` - DNS server host name list separated by space (maximum 4 domains).
