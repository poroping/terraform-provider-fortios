---
subcategory: "FortiGate Cifs"
layout: "fortios"
page_title: "FortiOS: fortios_cifs_domaincontroller"
description: |-
  Define known domain controller servers.
---

## fortios_cifs_domaincontroller
Define known domain controller servers.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `server_name` to be defined.

* `domain_name` - Fully qualified domain name (FQDN). E.g. 'EXAMPLE.COM'.
* `ip` - IPv4 server address.
* `ip6` - IPv6 server address.
* `password` - Password for specified username.
* `port` - Port number of service. Port number 0 indicates automatic discovery.
* `server_name` - Name of the server to connect to.
* `username` - User name to sign in with. Must have proper permissions for service.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_cifs_domaincontroller can be imported using:
```sh
terraform import fortios_cifs_domaincontroller.labelname {{mkey}}
```
