---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_apiuser"
description: |-
  Get information on a fortios Configure API users.
---

# Data Source: fortios_system_apiuser
Use this data source to get information on a fortios Configure API users.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) User name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `accprofile` - Admin user access profile.
* `api_key` - Admin user password.
* `comments` - Comment.
* `cors_allow_origin` - Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
* `name` - User name.
* `peer_auth` - Enable/disable peer authentication.
* `peer_group` - Peer group name.
* `schedule` - Schedule name.
* `trusthost` - Trusthost.The structure of `trusthost` block is documented below.

The `trusthost` block contains:

* `id` - ID.
* `ipv4_trusthost` - IPv4 trusted host address.
* `ipv6_trusthost` - IPv6 trusted host address.
* `type` - Trusthost type.
* `vdom` - Virtual domains.The structure of `vdom` block is documented below.

The `vdom` block contains:

* `name` - Virtual domain name.
