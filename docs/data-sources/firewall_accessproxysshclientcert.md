---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_accessproxysshclientcert"
description: |-
  Get information on a fortios Configure Access Proxy SSH client certificate.
---

# Data Source: fortios_firewall_accessproxysshclientcert
Use this data source to get information on a fortios Configure Access Proxy SSH client certificate.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) SSH client certificate name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `auth_ca` - Name of the SSH server public key authentication CA.
* `name` - SSH client certificate name.
* `permit_agent_forwarding` - Enable/disable appending permit-agent-forwarding certificate extension.
* `permit_port_forwarding` - Enable/disable appending permit-port-forwarding certificate extension.
* `permit_pty` - Enable/disable appending permit-pty certificate extension.
* `permit_user_rc` - Enable/disable appending permit-user-rc certificate extension.
* `permit_x11_forwarding` - Enable/disable appending permit-x11-forwarding certificate extension.
* `source_address` - Enable/disable appending source-address certificate critical option. This option ensure certificate only accepted from FortiGate source address.
* `cert_extension` - Configure certificate extension for user certificate.The structure of `cert_extension` block is documented below.

The `cert_extension` block contains:

* `critical` - Critical option.
* `data` - Data of certificate extension.
* `name` - Name of certificate extension.
* `type` - Type of certificate extension.
