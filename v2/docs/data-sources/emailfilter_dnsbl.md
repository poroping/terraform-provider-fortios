---
subcategory: "FortiGate Emailfilter"
layout: "fortios"
page_title: "FortiOS: fortios_emailfilter_dnsbl"
description: |-
  Get information on a fortios Configure AntiSpam DNSBL/ORBL.
---

# Data Source: fortios_emailfilter_dnsbl
Use this data source to get information on a fortios Configure AntiSpam DNSBL/ORBL.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Optional comments.
* `id` - ID.
* `name` - Name of table.
* `entries` - Spam filter DNSBL and ORBL server.The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Reject connection or mark as spam email.
* `id` - DNSBL/ORBL entry ID.
* `server` - DNSBL or ORBL server name.
* `status` - Enable/disable status.
