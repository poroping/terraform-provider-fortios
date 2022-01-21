---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_securityexemptlist"
description: |-
  Configure security exemption list.
---

## fortios_user_securityexemptlist
Configure security exemption list.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `description` - Description.
* `name` - Name of the exempt list.
* `rule` - Configure rules for exempting users from captive portal authentication. The structure of `rule` block is documented below.

The `rule` block contains:

* `id` - ID.
* `dstaddr` - Destination addresses or address groups. The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address or group name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `service` - Destination services. The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service name. This attribute must reference one of the following datasources: `firewall.service.custom.name` `firewall.service.group.name` .
* `srcaddr` - Source addresses or address groups. The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address or group name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_securityexemptlist can be imported using:
```sh
terraform import fortios_user_securityexemptlist.labelname {{mkey}}
```
