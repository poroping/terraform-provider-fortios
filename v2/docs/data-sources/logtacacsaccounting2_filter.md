---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logtacacsaccounting2_filter"
description: |-
  Get information on a fortios Settings for TACACS+ accounting events filter.
---

# Data Source: fortios_logtacacsaccounting2_filter
Use this data source to get information on a fortios Settings for TACACS+ accounting events filter.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `cli_cmd_audit` - Enable/disable TACACS+ accounting for CLI commands audit.
* `config_change_audit` - Enable/disable TACACS+ accounting for configuration change events audit.
* `login_audit` - Enable/disable TACACS+ accounting for login events audit.
