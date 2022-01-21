---
subcategory: "FortiGate Webfilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_ipsurlfiltersetting"
description: |-
  Configure IPS URL filter settings.
---

## fortios_webfilter_ipsurlfiltersetting
Configure IPS URL filter settings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `device` - Interface for this route. This attribute must reference one of the following datasources: `system.interface.name` .
* `distance` - Administrative distance (1 - 255) for this route.
* `gateway` - Gateway IP address for this route.
* `geo_filter` - Filter based on geographical location. Route will NOT be installed if the resolved IP address belongs to the country in the filter.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_webfilter_ipsurlfiltersetting can be imported using:
```sh
terraform import fortios_webfilter_ipsurlfiltersetting.labelname {{mkey}}
```
