---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_externalresource"
description: |-
  Configure external resource.
---

## fortios_system_externalresource
Configure external resource.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `category` - User resource category.
* `comments` - Comment.
* `interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `name` - External resource name.
* `password` - HTTP basic authentication password.
* `refresh_rate` - Time interval to refresh external resource (1 - 43200 min, default = 5 min).
* `resource` - URI of external resource.
* `source_ip` - Source IPv4 address used to communicate with server.
* `status` - Enable/disable user resource. Valid values: `enable` `disable` .
* `type` - User resource type. Valid values: `category` `address` `domain` `malware` .
* `user_agent` - HTTP User-Agent header (default = 'curl/7.58.0').
* `username` - HTTP basic authentication user name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_externalresource can be imported using:
```sh
terraform import fortios_system_externalresource.labelname {{mkey}}
```
