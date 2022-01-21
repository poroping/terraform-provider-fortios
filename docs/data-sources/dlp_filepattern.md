---
subcategory: "FortiGate Dlp"
layout: "fortios"
page_title: "FortiOS: fortios_dlp_filepattern"
description: |-
  Get information on a fortios Configure file patterns used by DLP blocking.
---

# Data Source: fortios_dlp_filepattern
Use this data source to get information on a fortios Configure file patterns used by DLP blocking.


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
* `name` - Name of table containing the file pattern list.
* `entries` - Configure file patterns used by DLP blocking.The structure of `entries` block is documented below.

The `entries` block contains:

* `file_type` - Select a file type.
* `filter_type` - Filter by file name pattern or by file type.
* `pattern` - Add a file name pattern.
