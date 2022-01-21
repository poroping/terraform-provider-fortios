---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_debugurl"
description: |-
  Configure debug URL addresses.
---

## fortios_webproxy_debugurl
Configure debug URL addresses.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `exact` - Enable/disable matching the exact path. Valid values: `enable` `disable` .
* `name` - Debug URL name.
* `status` - Enable/disable this URL exemption. Valid values: `enable` `disable` .
* `url_pattern` - URL exemption pattern.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_webproxy_debugurl can be imported using:
```sh
terraform import fortios_webproxy_debugurl.labelname {{mkey}}
```
