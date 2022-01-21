---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_automationstitch"
description: |-
  Get information on a fortios Automation stitches.
---

# Data Source: fortios_system_automationstitch
Use this data source to get information on a fortios Automation stitches.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `description` - Description.
* `name` - Name.
* `status` - Enable/disable this stitch.
* `trigger` - Trigger name.
* `action` - Action names.The structure of `action` block is documented below.

The `action` block contains:

* `name` - Action name.
* `actions` - Configure stitch actions.The structure of `actions` block is documented below.

The `actions` block contains:

* `action` - Action name.
* `delay` - Delay before execution (in seconds).
* `id` - Entry ID.
* `required` - Required in action chain.
* `destination` - Serial number/HA group-name of destination devices.The structure of `destination` block is documented below.

The `destination` block contains:

* `name` - Destination name.
