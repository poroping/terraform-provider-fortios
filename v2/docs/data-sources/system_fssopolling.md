---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fssopolling"
description: |-
  Get information on a fortios Configure Fortinet Single Sign On (FSSO) server.
---

# Data Source: fortios_system_fssopolling
Use this data source to get information on a fortios Configure Fortinet Single Sign On (FSSO) server.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `auth_password` - Password to connect to FSSO Agent.
* `authentication` - Enable/disable FSSO Agent Authentication.
* `listening_port` - Listening port to accept clients (1 - 65535).
* `status` - Enable/disable FSSO Polling Mode.
