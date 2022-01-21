---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_acme"
description: |-
  Get information on a fortios Configure ACME client.
---

# Data Source: fortios_system_acme
Use this data source to get information on a fortios Configure ACME client.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `accounts` - ACME accounts list.The structure of `accounts` block is documented below.

The `accounts` block contains:

* `ca_url` - Account ca_url.
* `email` - Account email.
* `id` - Account id.
* `privatekey` - Account Private Key.
* `status` - Account status.
* `url` - Account url.
* `interface` - Interface(s) on which the ACME client will listen for challenges.The structure of `interface` block is documented below.

The `interface` block contains:

* `interface_name` - Interface name.
