---
subcategory: "FortiGate Endpoint-Control"
layout: "fortios"
page_title: "FortiOS: fortios_endpointcontrol_settings"
description: |-
  Configure endpoint control settings.
---

## fortios_endpointcontrol_settings
Configure endpoint control settings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `forticlient_disconnect_unsupported_client` - Enable/disable disconnecting of unsupported FortiClient endpoints. Valid values: `enable` `disable` .
* `forticlient_keepalive_interval` - Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).
* `forticlient_sys_update_interval` - Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).
* `forticlient_user_avatar` - Enable/disable uploading FortiClient user avatars. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_endpointcontrol_settings can be imported using:
```sh
terraform import fortios_endpointcontrol_settings.labelname {{mkey}}
```
