---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vnetunnel"
description: |-
  Configure virtual network enabler tunnel.
---

## fortios_system_vnetunnel
Configure virtual network enabler tunnel.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `auto_asic_offload` - Enable/disable tunnel ASIC offloading. Valid values: `enable` `disable` .
* `bmr_hostname` - BMR hostname.
* `br` - IPv6 address or FQDN of the border relay.
* `http_password` - HTTP authentication password.
* `http_username` - HTTP authentication user name.
* `interface` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `ipv4_address` - Tunnel IPv4 address and netmask.
* `mode` - VNE tunnel mode. Valid values: `map-e` `fixed-ip` `ds-lite` .
* `ssl_certificate` - Name of local certificate for SSL connections. This attribute must reference one of the following datasources: `certificate.local.name` .
* `status` - Enable/disable VNE tunnel. Valid values: `enable` `disable` .
* `update_url` - URL of provisioning server.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_vnetunnel can be imported using:
```sh
terraform import fortios_system_vnetunnel.labelname {{mkey}}
```
