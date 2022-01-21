---
subcategory: "FortiGate Endpoint-Control"
layout: "fortios"
page_title: "FortiOS: fortios_endpointcontrol_settings"
description: |-
  Get information on a fortios Configure endpoint control settings.
---

# Data Source: fortios_endpointcontrol_settings
Use this data source to get information on a fortios Configure endpoint control settings.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `forticlient_disconnect_unsupported_client` - Enable/disable disconnecting of unsupported FortiClient endpoints.
* `forticlient_keepalive_interval` - Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).
* `forticlient_sys_update_interval` - Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).
* `forticlient_user_avatar` - Enable/disable uploading FortiClient user avatars.
