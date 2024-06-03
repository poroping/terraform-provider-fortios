---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_dnsdatabase"
description: |-
  Configure DNS databases.
---

## fortios_system_dnsdatabase
Configure DNS databases.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `allow_transfer` - DNS zone transfer IP address list.
* `authoritative` - Enable/disable authoritative zone. Valid values: `enable` `disable` .
* `contact` - Email address of the administrator for this zone. You can specify only the username, such as admin or the full email address, such as admin@test.com When using only a username, the domain of the email will be this zone.
* `domain` - Domain name.
* `forwarder` - DNS zone forwarder IP address list.
* `ip_master` - IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.
* `ip_primary` - IP address of primary DNS server. Entries in this primary DNS server and imported into the DNS zone.
* `name` - Zone name.
* `primary_name` - Domain name of the default DNS server for this zone.
* `rr_max` - Maximum number of resource records (10 - 65536, 0 means infinite).
* `source_ip` - Source IP for forwarding to DNS server.
* `status` - Enable/disable this DNS zone. Valid values: `enable` `disable` .
* `ttl` - Default time-to-live value for the entries of this DNS zone (0 - 2147483647 sec, default = 86400).
* `type` - Zone type (primary to manage entries directly, secondary to import entries from other zones). Valid values: `primary` `secondary` .
* `view` - Zone view (public to serve public clients, shadow to serve internal clients). Valid values: `shadow` `public` `shadow-ztna` .
* `dns_entry` - DNS entry. The structure of `dns_entry` block is documented below.

The `dns_entry` block contains:

* `canonical_name` - Canonical name of the host.
* `hostname` - Name of the host.
* `id` - DNS entry ID.
* `ip` - IPv4 address of the host.
* `ipv6` - IPv6 address of the host.
* `preference` - DNS entry preference (0 - 65535, highest preference = 0, default = 10).
* `status` - Enable/disable resource record status. Valid values: `enable` `disable` .
* `ttl` - Time-to-live for this entry (0 to 2147483647 sec, default = 0).
* `type` - Resource record type. Valid values: `A` `NS` `CNAME` `MX` `AAAA` `PTR` `PTR_V6` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_dnsdatabase can be imported using:
```sh
terraform import fortios_system_dnsdatabase.labelname {{mkey}}
```
