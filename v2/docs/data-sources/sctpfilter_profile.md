---
subcategory: "FortiGate Sctp-Filter"
layout: "fortios"
page_title: "FortiOS: fortios_sctpfilter_profile"
description: |-
  Get information on a fortios Configure SCTP filter profiles.
---

# Data Source: fortios_sctpfilter_profile
Use this data source to get information on a fortios Configure SCTP filter profiles.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `name` - Profile name.
* `ppid_filters` - PPID filters list.The structure of `ppid_filters` block is documented below.

The `ppid_filters` block contains:

* `action` - Action taken when PPID is matched.
* `comment` - Comment.
* `id` - ID.
* `ppid` - Payload protocol identifier.
