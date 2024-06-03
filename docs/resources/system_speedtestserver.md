---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_speedtestserver"
description: |-
  Configure speed test server list.
---

## fortios_system_speedtestserver
Configure speed test server list.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `name` - Speed test server name.
* `timestamp` - Speed test server timestamp.
* `host` - Hosts of the server. The structure of `host` block is documented below.

The `host` block contains:

* `distance` - Speed test host distance.
* `id` - Server host ID.
* `ip` - Server host IPv4 address.
* `latitude` - Speed test host latitude.
* `longitude` - Speed test host longitude.
* `password` - Speed test host password.
* `port` - Server host port number to communicate with client.
* `user` - Speed test host user name.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_speedtestserver can be imported using:
```sh
terraform import fortios_system_speedtestserver.labelname {{mkey}}
```
