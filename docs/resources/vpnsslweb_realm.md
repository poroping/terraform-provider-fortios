---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnsslweb_realm"
description: |-
  Realm.
---

## fortios_vpnsslweb_realm
Realm.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `url_path` to be defined.

* `login_page` - Replacement HTML for SSL-VPN login page.
* `max_concurrent_user` - Maximum concurrent users (0 - 65535, 0 means unlimited).
* `nas_ip` - IP address used as a NAS-IP to communicate with the RADIUS server.
* `radius_port` - RADIUS service port number (0 - 65535, 0 means user.radius.radius-port).
* `radius_server` - RADIUS server associated with realm. This attribute must reference one of the following datasources: `user.radius.name` .
* `url_path` - URL path to access SSL-VPN login page.
* `virtual_host` - Virtual host name for realm.
* `virtual_host_only` - Enable/disable enforcement of virtual host method for SSL-VPN client access. Valid values: `enable` `disable` .
* `virtual_host_server_cert` - Name of the server certificate to used for this realm. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_vpnsslweb_realm can be imported using:
```sh
terraform import fortios_vpnsslweb_realm.labelname {{mkey}}
```
