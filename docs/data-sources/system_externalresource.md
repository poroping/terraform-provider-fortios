---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_externalresource"
description: |-
  Get information on a fortios Configure external resource.
---

# Data Source: fortios_system_externalresource
Use this data source to get information on a fortios Configure external resource.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) External resource name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `category` - User resource category.
* `comments` - Comment.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `name` - External resource name.
* `password` - HTTP basic authentication password.
* `refresh_rate` - Time interval to refresh external resource (1 - 43200 min, default = 5 min).
* `resource` - URI of external resource.
* `server_identity_check` - Certificate verification option.
* `source_ip` - Source IPv4 address used to communicate with server.
* `status` - Enable/disable user resource.
* `type` - User resource type.
* `update_method` - External resource update method.
* `user_agent` - HTTP User-Agent header (default = 'curl/7.58.0').
* `username` - HTTP basic authentication user name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
