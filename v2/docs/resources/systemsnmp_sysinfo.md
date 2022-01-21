---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemsnmp_sysinfo"
description: |-
  SNMP system info configuration.
---

## fortios_systemsnmp_sysinfo
SNMP system info configuration.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `contact_info` - Contact information.
* `description` - System description.
* `engine_id` - Local SNMP engineID string (maximum 27 characters).
* `engine_id_type` - Local SNMP engineID type (text/hex/mac). Valid values: `text` `hex` `mac` .
* `location` - System location.
* `status` - Enable/disable SNMP. Valid values: `enable` `disable` .
* `trap_high_cpu_threshold` - CPU usage when trap is sent.
* `trap_log_full_threshold` - Log disk usage when trap is sent.
* `trap_low_memory_threshold` - Memory usage when trap is sent.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_systemsnmp_sysinfo can be imported using:
```sh
terraform import fortios_systemsnmp_sysinfo.labelname {{mkey}}
```
