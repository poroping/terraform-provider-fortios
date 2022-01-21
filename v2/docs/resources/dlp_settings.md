---
subcategory: "FortiGate Dlp"
layout: "fortios"
page_title: "FortiOS: fortios_dlp_settings"
description: |-
  Designate logical storage for DLP fingerprint database.
---

## fortios_dlp_settings
Designate logical storage for DLP fingerprint database.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `cache_mem_percent` - Maximum percentage of available memory allocated to caching (1 - 15%).
* `chunk_size` - Maximum fingerprint chunk size.  **Changing will flush the entire database**.
* `db_mode` - Behaviour when the maximum size is reached. Valid values: `stop-adding` `remove-modified-then-oldest` `remove-oldest` .
* `size` - Maximum total size of files within the storage (MB).
* `storage_device` - Storage device name. This attribute must reference one of the following datasources: `system.storage.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_dlp_settings can be imported using:
```sh
terraform import fortios_dlp_settings.labelname {{mkey}}
```
