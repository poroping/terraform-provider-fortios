---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnsslweb_realm"
description: |-
  Get information on a fortios Realm.
---

# Data Source: fortios_vpnsslweb_realm
Use this data source to get information on a fortios Realm.


## Example Usage

```hcl

```

## Argument Reference

* `url_path` - (Required) URL path to access SSL-VPN login page.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `login_page` - Replacement HTML for SSL-VPN login page.
* `max_concurrent_user` - Maximum concurrent users (0 - 65535, 0 means unlimited).
* `nas_ip` - IP address used as a NAS-IP to communicate with the RADIUS server.
* `radius_port` - RADIUS service port number (0 - 65535, 0 means user.radius.radius-port).
* `radius_server` - RADIUS server associated with realm.
* `url_path` - URL path to access SSL-VPN login page.
* `virtual_host` - Virtual host name for realm.
* `virtual_host_only` - Enable/disable enforcement of virtual host method for SSL-VPN client access.
* `virtual_host_server_cert` - Name of the server certificate to used for this realm.
