---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_apiuser"
description: |-
  Configure API users.
---

## fortios_system_apiuser
Configure API users.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `accprofile` - Admin user access profile. This attribute must reference one of the following datasources: `system.accprofile.name` .
* `api_key` - Admin user password.
* `comments` - Comment.
* `cors_allow_origin` - Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
* `name` - User name.
* `peer_auth` - Enable/disable peer authentication. Valid values: `enable` `disable` .
* `peer_group` - Peer group name.
* `schedule` - Schedule name.
* `trusthost` - Trusthost. The structure of `trusthost` block is documented below.

The `trusthost` block contains:

* `id` - ID.
* `ipv4_trusthost` - IPv4 trusted host address.
* `ipv6_trusthost` - IPv6 trusted host address.
* `type` - Trusthost type. Valid values: `ipv4-trusthost` `ipv6-trusthost` .
* `vdom` - Virtual domains. The structure of `vdom` block is documented below.

The `vdom` block contains:

* `name` - Virtual domain name. This attribute must reference one of the following datasources: `system.vdom.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_apiuser can be imported using:
```sh
terraform import fortios_system_apiuser.labelname {{mkey}}
```
