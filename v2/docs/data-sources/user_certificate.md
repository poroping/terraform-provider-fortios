---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_certificate"
description: |-
  Get information on a fortios Configure certificate users.
---

# Data Source: fortios_user_certificate
Use this data source to get information on a fortios Configure certificate users.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) User name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `common_name` - Certificate common name.
* `id` - User ID.
* `issuer` - CA certificate used for client certificate verification.
* `name` - User name.
* `status` - Enable/disable allowing the certificate user to authenticate with the FortiGate unit.
* `type` - Type of certificate authentication method.
