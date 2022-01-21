---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_fortitoken"
description: |-
  Get information on a fortios Configure FortiToken.
---

# Data Source: fortios_user_fortitoken
Use this data source to get information on a fortios Configure FortiToken.


## Example Usage

```hcl

```

## Argument Reference

* `serial_number` - (Required) Serial number.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `activation_code` - Mobile token user activation-code.
* `activation_expire` - Mobile token user activation-code expire time.
* `comments` - Comment.
* `license` - Mobile token license.
* `os_ver` - Device Mobile Version.
* `reg_id` - Device Reg ID.
* `serial_number` - Serial number.
* `status` - Status
