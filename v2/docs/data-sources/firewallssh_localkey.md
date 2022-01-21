---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallssh_localkey"
description: |-
  Get information on a fortios SSH proxy local keys.
---

# Data Source: fortios_firewallssh_localkey
Use this data source to get information on a fortios SSH proxy local keys.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) SSH proxy local key name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - SSH proxy local key name.
* `password` - Password for SSH private key.
* `private_key` - SSH proxy private key, encrypted with a password.
* `public_key` - SSH proxy public key.
* `source` - SSH proxy local key source type.
