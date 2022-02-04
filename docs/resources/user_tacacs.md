---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_tacacs"
description: |-
  Configure TACACS+ server entries.
---

## fortios_user_tacacs
Configure TACACS+ server entries.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `authen_type` - Allowed authentication protocols/methods. Valid values: `mschap` `chap` `pap` `ascii` `auto` .
* `authorization` - Enable/disable TACACS+ authorization. Valid values: `enable` `disable` .
* `interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `key` - Key to access the primary server.
* `name` - TACACS+ server entry name.
* `port` - Port number of the TACACS+ server.
* `secondary_key` - Key to access the secondary server.
* `secondary_server` - Secondary TACACS+ server CN domain name or IP address.
* `server` - Primary TACACS+ server CN domain name or IP address.
* `source_ip` - Source IP address for communications to TACACS+ server.
* `tertiary_key` - Key to access the tertiary server.
* `tertiary_server` - Tertiary TACACS+ server CN domain name or IP address.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_tacacs can be imported using:
```sh
terraform import fortios_user_tacacs.labelname {{mkey}}
```
