---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logtacacsaccounting2_filter"
description: |-
  Settings for TACACS+ accounting events filter.
---

## fortios_logtacacsaccounting2_filter
Settings for TACACS+ accounting events filter.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `cli_cmd_audit` - Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable` `disable` .
* `config_change_audit` - Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable` `disable` .
* `login_audit` - Enable/disable TACACS+ accounting for login events audit. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_logtacacsaccounting2_filter can be imported using:
```sh
terraform import fortios_logtacacsaccounting2_filter.labelname {{mkey}}
```
