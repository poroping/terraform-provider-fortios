---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logtacacsaccounting2_setting"
description: |-
  Settings for TACACS+ accounting.
---

## fortios_logtacacsaccounting2_setting
Settings for TACACS+ accounting.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `server` - Address of TACACS+ server.
* `server_key` - Key to access the TACACS+ server.
* `status` - Enable/disable TACACS+ accounting. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_logtacacsaccounting2_setting can be imported using:
```sh
terraform import fortios_logtacacsaccounting2_setting.labelname {{mkey}}
```
