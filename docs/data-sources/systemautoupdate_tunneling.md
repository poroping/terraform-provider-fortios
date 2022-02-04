---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemautoupdate_tunneling"
description: |-
  Get information on a fortios Configure web proxy tunneling for the FDN.
---

# Data Source: fortios_systemautoupdate_tunneling
Use this data source to get information on a fortios Configure web proxy tunneling for the FDN.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `address` - Web proxy IP address or FQDN.
* `password` - Web proxy password.
* `port` - Web proxy port.
* `status` - Enable/disable web proxy tunneling.
* `username` - Web proxy username.
