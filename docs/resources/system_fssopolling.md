---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fssopolling"
description: |-
  Configure Fortinet Single Sign On (FSSO) server.
---

## fortios_system_fssopolling
Configure Fortinet Single Sign On (FSSO) server.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `auth_password` - Password to connect to FSSO Agent.
* `authentication` - Enable/disable FSSO Agent Authentication. Valid values: `enable` `disable` .
* `listening_port` - Listening port to accept clients (1 - 65535).
* `status` - Enable/disable FSSO Polling Mode. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_fssopolling can be imported using:
```sh
terraform import fortios_system_fssopolling.labelname {{mkey}}
```
