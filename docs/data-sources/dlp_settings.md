---
subcategory: "FortiGate Dlp"
layout: "fortios"
page_title: "FortiOS: fortios_dlp_settings"
description: |-
  Get information on a fortios Designate logical storage for DLP fingerprint database.
---

# Data Source: fortios_dlp_settings
Use this data source to get information on a fortios Designate logical storage for DLP fingerprint database.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `cache_mem_percent` - Maximum percentage of available memory allocated to caching (1 - 15).
* `chunk_size` - Maximum fingerprint chunk size. Caution, changing this setting will flush the entire database.
* `db_mode` - Behavior when the maximum size is reached.
* `size` - Maximum total size of files within the storage (MB).
* `storage_device` - Storage device name.
