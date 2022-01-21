---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_securityexemptlist"
description: |-
  Get information on a fortios Configure security exemption list.
---

# Data Source: fortios_user_securityexemptlist
Use this data source to get information on a fortios Configure security exemption list.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name of the exempt list.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `description` - Description.
* `name` - Name of the exempt list.
* `rule` - Configure rules for exempting users from captive portal authentication.The structure of `rule` block is documented below.

The `rule` block contains:

* `id` - ID.
* `dstaddr` - Destination addresses or address groups.The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address or group name.
* `service` - Destination services.The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service name.
* `srcaddr` - Source addresses or address groups.The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address or group name.
