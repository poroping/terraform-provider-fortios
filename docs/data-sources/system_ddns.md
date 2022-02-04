---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ddns"
description: |-
  Get information on a fortios Configure DDNS.
---

# Data Source: fortios_system_ddns
Use this data source to get information on a fortios Configure DDNS.


## Example Usage

```hcl

```

## Argument Reference

* `ddnsid` - (Required) DDNS ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `addr_type` - Address type of interface address in DDNS update.
* `bound_ip` - Bound IP address.
* `clear_text` - Enable/disable use of clear text connections.
* `ddns_auth` - Enable/disable TSIG authentication for your DDNS server.
* `ddns_domain` - Your fully qualified domain name. For example, yourname.ddns.com.
* `ddns_key` - DDNS update key (base 64 encoding).
* `ddns_keyname` - DDNS update key name.
* `ddns_password` - DDNS password.
* `ddns_server` - Select a DDNS service provider.
* `ddns_server_ip` - Generic DDNS server IP.
* `ddns_sn` - DDNS Serial Number.
* `ddns_ttl` - Time-to-live for DDNS packets.
* `ddns_username` - DDNS user name.
* `ddns_zone` - Zone of your domain name (for example, DDNS.com).
* `ddnsid` - DDNS ID.
* `server_type` - Address type of the DDNS server.
* `ssl_certificate` - Name of local certificate for SSL connections.
* `update_interval` - DDNS update interval (60 - 2592000 sec, 0 means default).
* `use_public_ip` - Enable/disable use of public IP address.
* `ddns_server_addr` - Generic DDNS server IP/FQDN list.The structure of `ddns_server_addr` block is documented below.

The `ddns_server_addr` block contains:

* `addr` - IP address or FQDN of the server.
* `monitor_interface` - Monitored interface.The structure of `monitor_interface` block is documented below.

The `monitor_interface` block contains:

* `interface_name` - Interface name.
