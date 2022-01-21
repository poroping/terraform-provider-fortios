---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallssh_hostkey"
description: |-
  Get information on a fortios SSH proxy host public keys.
---

# Data Source: fortios_firewallssh_hostkey
Use this data source to get information on a fortios SSH proxy host public keys.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) SSH public key name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `hostname` - Hostname of the SSH server, to match SSH certificate principals. 
* `ip` - IP address of the SSH server.
* `name` - SSH public key name.
* `nid` - Set the nid of the ECDSA key.
* `port` - Port of the SSH server.
* `public_key` - SSH public key.
* `status` - Set the trust status of the public key.
* `type` - Set the type of the public key.
* `usage` - Usage for this public key.
