---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_smsserver"
description: |-
  Configure SMS server for sending SMS messages to support user authentication.
---

## fortios_system_smsserver
Configure SMS server for sending SMS messages to support user authentication.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `mail_server` - Email-to-SMS server domain name.
* `name` - Name of SMS server.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_smsserver can be imported using:
```sh
terraform import fortios_system_smsserver.labelname {{mkey}}
```
