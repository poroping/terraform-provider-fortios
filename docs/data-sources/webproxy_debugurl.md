---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_debugurl"
description: |-
  Get information on a fortios Configure debug URL addresses.
---

# Data Source: fortios_webproxy_debugurl
Use this data source to get information on a fortios Configure debug URL addresses.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Debug URL name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `exact` - Enable/disable matching the exact path.
* `name` - Debug URL name.
* `status` - Enable/disable this URL exemption.
* `url_pattern` - URL exemption pattern.
