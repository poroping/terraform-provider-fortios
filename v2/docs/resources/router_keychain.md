---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_keychain"
description: |-
  Configure key-chain.
---

## fortios_router_keychain
Configure key-chain.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `name` - Key-chain name.
* `key` - Configuration method to edit key settings. The structure of `key` block is documented below.

The `key` block contains:

* `accept_lifetime` - Lifetime of received authentication key (format: hh:mm:ss day month year).
* `algorithm` - Cryptographic algorithm. Valid values: `md5` `hmac-sha1` `hmac-sha256` `hmac-sha384` `hmac-sha512` .
* `id` - Key ID (0 - 2147483647).
* `key_string` - Password for the key (max. = 64 characters).
* `send_lifetime` - Lifetime of sent authentication key (format: hh:mm:ss day month year).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_keychain can be imported using:
```sh
terraform import fortios_router_keychain.labelname {{mkey}}
```
