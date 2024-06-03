---
subcategory: "FortiGate Ips"
layout: "fortios"
page_title: "FortiOS: fortios_ips_rule"
description: |-
  Get information on a fortios Configure IPS rules.
---

# Data Source: fortios_ips_rule
Use this data source to get information on a fortios Configure IPS rules.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Rule name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `action` - Action.
* `application` - Vulnerable applications.
* `date` - Date.
* `group` - Group.
* `location` - Vulnerable location.
* `log` - Enable/disable logging.
* `log_packet` - Enable/disable packet logging.
* `name` - Rule name.
* `os` - Vulnerable operation systems.
* `rev` - Revision.
* `rule_id` - Rule ID.
* `service` - Vulnerable service.
* `severity` - Severity.
* `metadata` - Meta data.The structure of `metadata` block is documented below.

The `metadata` block contains:

* `id` - ID.
* `metaid` - Meta ID.
* `valueid` - Value ID.
