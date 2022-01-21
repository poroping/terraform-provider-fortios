---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fortisandbox"
description: |-
  Configure FortiSandbox.
---

## fortios_system_fortisandbox
Configure FortiSandbox.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `email` - Notifier email address.
* `enc_algorithm` - Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default` `high` `low` .
* `forticloud` - Enable/disable FortiSandbox Cloud. Valid values: `enable` `disable` .
* `interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `server` - Server address of the remote FortiSandbox.
* `source_ip` - Source IP address for communications to FortiSandbox.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default` `SSLv3` `TLSv1` `TLSv1-1` `TLSv1-2` .
* `status` - Enable/disable FortiSandbox. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_fortisandbox can be imported using:
```sh
terraform import fortios_system_fortisandbox.labelname {{mkey}}
```
