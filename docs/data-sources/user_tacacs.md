---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_tacacs"
description: |-
  Get information on a fortios Configure TACACS+ server entries.
---

# Data Source: fortios_user_tacacs
Use this data source to get information on a fortios Configure TACACS+ server entries.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) TACACS+ server entry name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `authen_type` - Allowed authentication protocols/methods.
* `authorization` - Enable/disable TACACS+ authorization.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `key` - Key to access the primary server.
* `name` - TACACS+ server entry name.
* `port` - Port number of the TACACS+ server.
* `secondary_key` - Key to access the secondary server.
* `secondary_server` - Secondary TACACS+ server CN domain name or IP address.
* `server` - Primary TACACS+ server CN domain name or IP address.
* `source_ip` - source IP for communications to TACACS+ server.
* `tertiary_key` - Key to access the tertiary server.
* `tertiary_server` - Tertiary TACACS+ server CN domain name or IP address.
