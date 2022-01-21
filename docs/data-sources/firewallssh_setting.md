---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallssh_setting"
description: |-
  Get information on a fortios SSH proxy settings.
---

# Data Source: fortios_firewallssh_setting
Use this data source to get information on a fortios SSH proxy settings.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `caname` - CA certificate used by SSH Inspection.
* `host_trusted_checking` - Enable/disable host trusted checking.
* `hostkey_dsa1024` - DSA certificate used by SSH proxy.
* `hostkey_ecdsa256` - ECDSA nid256 certificate used by SSH proxy.
* `hostkey_ecdsa384` - ECDSA nid384 certificate used by SSH proxy.
* `hostkey_ecdsa521` - ECDSA nid384 certificate used by SSH proxy.
* `hostkey_ed25519` - ED25519 hostkey used by SSH proxy.
* `hostkey_rsa2048` - RSA certificate used by SSH proxy.
* `untrusted_caname` - Untrusted CA certificate used by SSH Inspection.
