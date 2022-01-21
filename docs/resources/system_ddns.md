---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ddns"
description: |-
  Configure DDNS.
---

## fortios_system_ddns
Configure DDNS.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `ddnsid` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `addr_type` - Address type of interface address in DDNS update. Valid values: `ipv4` `ipv6` .
* `bound_ip` - Bound IP address.
* `clear_text` - Enable/disable use of clear text connections. Valid values: `disable` `enable` .
* `ddns_auth` - Enable/disable TSIG authentication for your DDNS server. Valid values: `disable` `tsig` .
* `ddns_domain` - Your fully qualified domain name (for example, yourname.DDNS.com).
* `ddns_key` - DDNS update key (base 64 encoding).
* `ddns_keyname` - DDNS update key name.
* `ddns_password` - DDNS password.
* `ddns_server` - Select a DDNS service provider. Valid values: `dyndns.org` `dyns.net` `tzo.com` `vavic.com` `dipdns.net` `now.net.cn` `dhs.org` `easydns.com` `genericDDNS` `FortiGuardDDNS` `noip.com` .
* `ddns_server_ip` - Generic DDNS server IP.
* `ddns_sn` - DDNS Serial Number.
* `ddns_ttl` - Time-to-live for DDNS packets.
* `ddns_username` - DDNS user name.
* `ddns_zone` - Zone of your domain name (for example, DDNS.com).
* `ddnsid` - DDNS ID.
* `server_type` - Address type of the DDNS server. Valid values: `ipv4` `ipv6` .
* `ssl_certificate` - Name of local certificate for SSL connections. This attribute must reference one of the following datasources: `certificate.local.name` .
* `update_interval` - DDNS update interval (60 - 2592000 sec, default = 300).
* `use_public_ip` - Enable/disable use of public IP address. Valid values: `disable` `enable` .
* `ddns_server_addr` - Generic DDNS server IP/FQDN list. The structure of `ddns_server_addr` block is documented below.

The `ddns_server_addr` block contains:

* `addr` - IP address or FQDN of the server.
* `monitor_interface` - Monitored interface. The structure of `monitor_interface` block is documented below.

The `monitor_interface` block contains:

* `interface_name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_ddns can be imported using:
```sh
terraform import fortios_system_ddns.labelname {{mkey}}
```
