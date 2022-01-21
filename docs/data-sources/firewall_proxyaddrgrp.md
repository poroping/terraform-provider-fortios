---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_proxyaddrgrp"
description: |-
  Get information on a fortios Configure web proxy address group.
---

# Data Source: fortios_firewall_proxyaddrgrp
Use this data source to get information on a fortios Configure web proxy address group.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Address group name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
* `comment` - Optional comments.
* `name` - Address group name.
* `type` - Source or destination address group type.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable visibility of the object in the GUI.
* `member` - Members of address group.The structure of `member` block is documented below.

The `member` block contains:

* `name` - Address name.
* `tagging` - Config object tagging.The structure of `tagging` block is documented below.

The `tagging` block contains:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.
