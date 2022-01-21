---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_acme"
description: |-
  Configure ACME client.
---

## fortios_system_acme
Configure ACME client.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `accounts` - ACME accounts list. The structure of `accounts` block is documented below.

The `accounts` block contains:

* `ca_url` - Account ca_url.
* `email` - Account email.
* `id` - Account id.
* `privatekey` - Account Private Key.
* `status` - Account status.
* `url` - Account url.
* `interface` - Interface(s) on which the ACME client will listen for challenges. The structure of `interface` block is documented below.

The `interface` block contains:

* `interface_name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_acme can be imported using:
```sh
terraform import fortios_system_acme.labelname {{mkey}}
```
