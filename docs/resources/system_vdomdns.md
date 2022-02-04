---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomdns"
description: |-
  Configure DNS servers for a non-management VDOM.
---

## fortios_system_vdomdns
Configure DNS servers for a non-management VDOM.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `alt_primary` - Alternate primary DNS server. This is not used as a failover DNS server.
* `alt_secondary` - Alternate secondary DNS server. This is not used as a failover DNS server.
* `dns_over_tls` - Enable/disable/enforce DNS over TLS. Valid values: `disable` `enable` `enforce` .
* `interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `ip6_primary` - Primary IPv6 DNS server IP address for the VDOM.
* `ip6_secondary` - Secondary IPv6 DNS server IP address for the VDOM.
* `primary` - Primary DNS server IP address for the VDOM.
* `protocol` - DNS transport protocols. Valid values: `cleartext` `dot` `doh` .
* `secondary` - Secondary DNS server IP address for the VDOM.
* `server_select_method` - Specify how configured servers are prioritized. Valid values: `least-rtt` `failover` .
* `source_ip` - Source IP for communications with the DNS server.
* `ssl_certificate` - Name of local certificate for SSL connections. This attribute must reference one of the following datasources: `certificate.local.name` .
* `vdom_dns` - Enable/disable configuring DNS servers for the current VDOM. Valid values: `enable` `disable` .
* `server_hostname` - DNS server host name list. The structure of `server_hostname` block is documented below.

The `server_hostname` block contains:

* `hostname` - DNS server host name list separated by space (maximum 4 domains).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_vdomdns can be imported using:
```sh
terraform import fortios_system_vdomdns.labelname {{mkey}}
```
