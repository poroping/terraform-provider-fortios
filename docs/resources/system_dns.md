---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_dns"
description: |-
  Configure DNS.
---

## fortios_system_dns
Configure DNS.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `alt_primary` - Alternate primary DNS server. This is not used as a failover DNS server.
* `alt_secondary` - Alternate secondary DNS server. This is not used as a failover DNS server.
* `cache_notfound_responses` - Enable/disable response from the DNS server when a record is not in cache. Valid values: `disable` `enable` .
* `dns_cache_limit` - Maximum number of records in the DNS cache.
* `dns_cache_ttl` - Duration in seconds that the DNS cache retains information.
* `dns_over_tls` - Enable/disable/enforce DNS over TLS. Valid values: `disable` `enable` `enforce` .
* `interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `ip6_primary` - Primary DNS server IPv6 address.
* `ip6_secondary` - Secondary DNS server IPv6 address.
* `log` - Local DNS log setting. Valid values: `disable` `error` `all` .
* `primary` - Primary DNS server IP address.
* `protocol` - DNS transport protocols. Valid values: `cleartext` `dot` `doh` .
* `retry` - Number of times to retry (0 - 5).
* `secondary` - Secondary DNS server IP address.
* `server_select_method` - Specify how configured servers are prioritized. Valid values: `least-rtt` `failover` .
* `source_ip` - IP address used by the DNS server as its source IP.
* `ssl_certificate` - Name of local certificate for SSL connections. This attribute must reference one of the following datasources: `certificate.local.name` .
* `timeout` - DNS query timeout interval in seconds (1 - 10).
* `domain` - Search suffix list for hostname lookup. The structure of `domain` block is documented below.

The `domain` block contains:

* `domain` - DNS search domain list separated by space (maximum 8 domains).
* `server_hostname` - DNS server host name list. The structure of `server_hostname` block is documented below.

The `server_hostname` block contains:

* `hostname` - DNS server host name list separated by space (maximum 4 domains).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_dns can be imported using:
```sh
terraform import fortios_system_dns.labelname {{mkey}}
```
