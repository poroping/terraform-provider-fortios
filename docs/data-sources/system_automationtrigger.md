---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_automationtrigger"
description: |-
  Get information on a fortios Trigger for automation stitches.
---

# Data Source: fortios_system_automationtrigger
Use this data source to get information on a fortios Trigger for automation stitches.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `description` - Description.
* `event_type` - Event type.
* `fabric_event_name` - Fabric connector event handler name.
* `fabric_event_severity` - Fabric connector event severity.
* `faz_event_name` - FortiAnalyzer event handler name.
* `faz_event_severity` - FortiAnalyzer event severity.
* `faz_event_tags` - FortiAnalyzer event tags.
* `ioc_level` - IOC threat level.
* `license_type` - License type.
* `name` - Name.
* `report_type` - Security Rating report.
* `serial` - Fabric connector serial number.
* `trigger_day` - Day within a month to trigger.
* `trigger_frequency` - Scheduled trigger frequency (default = daily).
* `trigger_hour` - Hour of the day on which to trigger (0 - 23, default = 1).
* `trigger_minute` - Minute of the hour on which to trigger (0 - 59, default = 0).
* `trigger_type` - Trigger type.
* `trigger_weekday` - Day of week for trigger.
* `fields` - Customized trigger field settings.The structure of `fields` block is documented below.

The `fields` block contains:

* `id` - Entry ID.
* `name` - Name.
* `value` - Value.
* `logid` - Log IDs to trigger event.The structure of `logid` block is documented below.

The `logid` block contains:

* `id` - Log ID.
