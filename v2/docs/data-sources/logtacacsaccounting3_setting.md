---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logtacacsaccounting3_setting"
description: |-
  Get information on a fortios Settings for TACACS+ accounting.
---

# Data Source: fortios_logtacacsaccounting3_setting
Use this data source to get information on a fortios Settings for TACACS+ accounting.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `server` - Address of TACACS+ server.
* `server_key` - Key to access the TACACS+ server.
* `status` - Enable/disable TACACS+ accounting.
