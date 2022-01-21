---
subcategory: "FortiGate Report"
layout: "fortios"
page_title: "FortiOS: fortios_report_setting"
description: |-
  Get information on a fortios Report setting configuration.
---

# Data Source: fortios_report_setting
Use this data source to get information on a fortios Report setting configuration.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `fortiview` - Enable/disable historical FortiView.
* `pdf_report` - Enable/disable PDF report.
* `report_source` - Report log source.
* `top_n` - Number of items to populate (1000 - 20000).
* `web_browsing_threshold` - Web browsing time calculation threshold (3 - 15 min).
