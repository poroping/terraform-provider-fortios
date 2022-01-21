---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fortisandbox"
description: |-
  Get information on a fortios Configure FortiSandbox.
---

# Data Source: fortios_system_fortisandbox
Use this data source to get information on a fortios Configure FortiSandbox.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `email` - Notifier email address.
* `enc_algorithm` - Configure the level of SSL protection for secure communication with FortiSandbox.
* `forticloud` - Enable/disable FortiSandbox Cloud.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `server` - Server address of the remote FortiSandbox.
* `source_ip` - Source IP address for communications to FortiSandbox.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `status` - Enable/disable FortiSandbox.
