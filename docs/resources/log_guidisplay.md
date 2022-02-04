---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_log_guidisplay"
description: |-
  Configure how log messages are displayed on the GUI.
---

## fortios_log_guidisplay
Configure how log messages are displayed on the GUI.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `fortiview_unscanned_apps` - Enable/disable showing unscanned traffic in FortiView application charts. Valid values: `enable` `disable` .
* `resolve_apps` - Resolve unknown applications on the GUI using Fortinet's remote application database. Valid values: `enable` `disable` .
* `resolve_hosts` - Enable/disable resolving IP addresses to hostname in log messages on the GUI using reverse DNS lookup. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_log_guidisplay can be imported using:
```sh
terraform import fortios_log_guidisplay.labelname {{mkey}}
```
