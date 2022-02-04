---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_log_guidisplay"
description: |-
  Get information on a fortios Configure how log messages are displayed on the GUI.
---

# Data Source: fortios_log_guidisplay
Use this data source to get information on a fortios Configure how log messages are displayed on the GUI.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `fortiview_unscanned_apps` - Enable/disable showing unscanned traffic in FortiView application charts.
* `resolve_apps` - Resolve unknown applications on the GUI using Fortinet's remote application database.
* `resolve_hosts` - Enable/disable resolving IP addresses to hostname in log messages on the GUI using reverse DNS lookup.
