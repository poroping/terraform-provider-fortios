---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_automationtrigger"
description: |-
  Trigger for automation stitches.
---

## fortios_system_automationtrigger
Trigger for automation stitches.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `description` - Description.
* `event_type` - Event type. Valid values: `ioc` `event-log` `reboot` `low-memory` `high-cpu` `license-near-expiry` `ha-failover` `config-change` `security-rating-summary` `virus-ips-db-updated` `faz-event` `incoming-webhook` `fabric-event` .
* `fabric_event_name` - Fabric connector event handler name.
* `fabric_event_severity` - Fabric connector event severity.
* `faz_event_name` - FortiAnalyzer event handler name.
* `faz_event_severity` - FortiAnalyzer event severity.
* `faz_event_tags` - FortiAnalyzer event tags.
* `ioc_level` - IOC threat level. Valid values: `medium` `high` .
* `license_type` - License type. Valid values: `forticare-support` `fortiguard-webfilter` `fortiguard-antispam` `fortiguard-antivirus` `fortiguard-ips` `fortiguard-management` `forticloud` `any` .
* `name` - Name.
* `report_type` - Security Rating report. Valid values: `posture` `coverage` `optimization` `any` .
* `serial` - Fabric connector serial number.
* `trigger_day` - Day within a month to trigger.
* `trigger_frequency` - Scheduled trigger frequency (default = daily). Valid values: `hourly` `daily` `weekly` `monthly` .
* `trigger_hour` - Hour of the day on which to trigger (0 - 23, default = 1).
* `trigger_minute` - Minute of the hour on which to trigger (0 - 59, default = 0).
* `trigger_type` - Trigger type. Valid values: `event-based` `scheduled` .
* `trigger_weekday` - Day of week for trigger. Valid values: `sunday` `monday` `tuesday` `wednesday` `thursday` `friday` `saturday` .
* `fields` - Customized trigger field settings. The structure of `fields` block is documented below.

The `fields` block contains:

* `id` - Entry ID.
* `name` - Name.
* `value` - Value.
* `logid` - Log IDs to trigger event. The structure of `logid` block is documented below.

The `logid` block contains:

* `id` - Log ID.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_automationtrigger can be imported using:
```sh
terraform import fortios_system_automationtrigger.labelname {{mkey}}
```
