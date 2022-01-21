---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallssh_setting"
description: |-
  SSH proxy settings.
---

## fortios_firewallssh_setting
SSH proxy settings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `caname` - CA certificate used by SSH Inspection. This attribute must reference one of the following datasources: `firewall.ssh.local-ca.name` .
* `host_trusted_checking` - Enable/disable host trusted checking. Valid values: `enable` `disable` .
* `hostkey_dsa1024` - DSA certificate used by SSH proxy. This attribute must reference one of the following datasources: `firewall.ssh.local-key.name` .
* `hostkey_ecdsa256` - ECDSA nid256 certificate used by SSH proxy. This attribute must reference one of the following datasources: `firewall.ssh.local-key.name` .
* `hostkey_ecdsa384` - ECDSA nid384 certificate used by SSH proxy. This attribute must reference one of the following datasources: `firewall.ssh.local-key.name` .
* `hostkey_ecdsa521` - ECDSA nid384 certificate used by SSH proxy. This attribute must reference one of the following datasources: `firewall.ssh.local-key.name` .
* `hostkey_ed25519` - ED25519 hostkey used by SSH proxy. This attribute must reference one of the following datasources: `firewall.ssh.local-key.name` .
* `hostkey_rsa2048` - RSA certificate used by SSH proxy. This attribute must reference one of the following datasources: `firewall.ssh.local-key.name` .
* `untrusted_caname` - Untrusted CA certificate used by SSH Inspection. This attribute must reference one of the following datasources: `firewall.ssh.local-ca.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewallssh_setting can be imported using:
```sh
terraform import fortios_firewallssh_setting.labelname {{mkey}}
```
