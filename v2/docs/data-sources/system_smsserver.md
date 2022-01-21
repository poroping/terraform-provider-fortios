---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_smsserver"
description: |-
  Get information on a fortios Configure SMS server for sending SMS messages to support user authentication.
---

# Data Source: fortios_system_smsserver
Use this data source to get information on a fortios Configure SMS server for sending SMS messages to support user authentication.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name of SMS server.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `mail_server` - Email-to-SMS server domain name.
* `name` - Name of SMS server.
