---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallssh_hostkey"
description: |-
  SSH proxy host public keys.
---

## fortios_firewallssh_hostkey
SSH proxy host public keys.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `hostname` - Hostname of the SSH server to match SSH certificate principals.
* `ip` - IP address of the SSH server.
* `name` - SSH public key name.
* `nid` - Set the nid of the ECDSA key. Valid values: `256` `384` `521` .
* `port` - Port of the SSH server.
* `public_key` - SSH public key.
* `status` - Set the trust status of the public key. Valid values: `trusted` `revoked` .
* `type` - Set the type of the public key. Valid values: `RSA` `DSA` `ECDSA` `ED25519` `RSA-CA` `DSA-CA` `ECDSA-CA` `ED25519-CA` .
* `usage` - Usage for this public key. Valid values: `transparent-proxy` `access-proxy` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewallssh_hostkey can be imported using:
```sh
terraform import fortios_firewallssh_hostkey.labelname {{mkey}}
```
