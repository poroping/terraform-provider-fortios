---
subcategory: "FortiGate Cifs"
layout: "fortios"
page_title: "FortiOS: fortios_cifs_domaincontroller"
description: |-
  Get information on a fortios Define known domain controller servers.
---

# Data Source: fortios_cifs_domaincontroller
Use this data source to get information on a fortios Define known domain controller servers.


## Example Usage

```hcl

```

## Argument Reference

* `server_name` - (Required) Name of the server to connect to.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `domain_name` - Fully qualified domain name (FQDN). E.g. 'EXAMPLE.COM'.
* `ip` - IPv4 server address.
* `ip6` - IPv6 server address.
* `password` - Password for specified username.
* `port` - Port number of service. Port number 0 indicates automatic discovery.
* `server_name` - Name of the server to connect to.
* `username` - User name to sign in with. Must have proper permissions for service.
