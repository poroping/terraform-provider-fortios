---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_accessproxysshclientcert"
description: |-
  Configure Access Proxy SSH client certificate.
---

## fortios_firewall_accessproxysshclientcert
Configure Access Proxy SSH client certificate.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `auth_ca` - Name of the SSH server public key authentication CA. This attribute must reference one of the following datasources: `firewall.ssh.local-ca.name` .
* `name` - SSH client certificate name.
* `permit_agent_forwarding` - Enable/disable appending permit-agent-forwarding certificate extension. Valid values: `enable` `disable` .
* `permit_port_forwarding` - Enable/disable appending permit-port-forwarding certificate extension. Valid values: `enable` `disable` .
* `permit_pty` - Enable/disable appending permit-pty certificate extension. Valid values: `enable` `disable` .
* `permit_user_rc` - Enable/disable appending permit-user-rc certificate extension. Valid values: `enable` `disable` .
* `permit_x11_forwarding` - Enable/disable appending permit-x11-forwarding certificate extension. Valid values: `enable` `disable` .
* `source_address` - Enable/disable appending source-address certificate critical option. This option ensure certificate only accepted from FortiGate source address. Valid values: `enable` `disable` .
* `cert_extension` - Configure certificate extension for user certificate. The structure of `cert_extension` block is documented below.

The `cert_extension` block contains:

* `critical` - Critical option. Valid values: `no` `yes` .
* `data` - Data of certificate extension.
* `name` - Name of certificate extension.
* `type` - Type of certificate extension. Valid values: `fixed` `user` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_accessproxysshclientcert can be imported using:
```sh
terraform import fortios_firewall_accessproxysshclientcert.labelname {{mkey}}
```
