---
subcategory: "FortiGate Report"
layout: "fortios"
page_title: "FortiOS: fortios_report_setting"
description: |-
  Report setting configuration.
---

## fortios_report_setting
Report setting configuration.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `fortiview` - Enable/disable historical FortiView. Valid values: `enable` `disable` .
* `pdf_report` - Enable/disable PDF report. Valid values: `enable` `disable` .
* `report_source` - Report log source. Valid values: `forward-traffic` `sniffer-traffic` `local-deny-traffic` .
* `top_n` - Number of items to populate (1000 - 20000).
* `web_browsing_threshold` - Web browsing time calculation threshold (3 - 15 min).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_report_setting can be imported using:
```sh
terraform import fortios_report_setting.labelname {{mkey}}
```
