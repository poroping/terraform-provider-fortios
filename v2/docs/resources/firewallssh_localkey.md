---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallssh_localkey"
description: |-
  SSH proxy local keys.
---

## fortios_firewallssh_localkey
SSH proxy local keys.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `name` - SSH proxy local key name.
* `password` - Password for SSH private key.
* `private_key` - SSH proxy private key, encrypted with a password.
* `public_key` - SSH proxy public key.
* `source` - SSH proxy local key source type. Valid values: `built-in` `user` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewallssh_localkey can be imported using:
```sh
terraform import fortios_firewallssh_localkey.labelname {{mkey}}
```
