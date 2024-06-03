---
subcategory: "FortiGate Dlp"
layout: "fortios"
page_title: "FortiOS: fortios_dlp_fpdocsource"
description: |-
  Create a DLP fingerprint database by allowing the FortiGate to access a file server containing files from which to create fingerprints.
---

## fortios_dlp_fpdocsource
Create a DLP fingerprint database by allowing the FortiGate to access a file server containing files from which to create fingerprints.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `date` - Day of the month on which to scan the server (1 - 31).
* `file_path` - Path on the server to the fingerprint files (max 119 characters).
* `file_pattern` - Files matching this pattern on the server are fingerprinted. Optionally use the * and ? wildcards.
* `keep_modified` - Enable so that when a file is changed on the server the FortiGate keeps the old fingerprint and adds a new fingerprint to the database. Valid values: `enable` `disable` .
* `name` - Name of the DLP fingerprint database.
* `password` - Password required to log into the file server.
* `period` - Frequency for which the FortiGate checks the server for new or changed files. Valid values: `none` `daily` `weekly` `monthly` .
* `remove_deleted` - Enable to keep the fingerprint database up to date when a file is deleted from the server. Valid values: `enable` `disable` .
* `scan_on_creation` - Enable to keep the fingerprint database up to date when a file is added or changed on the server. Valid values: `enable` `disable` .
* `scan_subdirectories` - Enable/disable scanning subdirectories to find files to create fingerprints from. Valid values: `enable` `disable` .
* `sensitivity` - Select a sensitivity or threat level for matches with this fingerprint database. Add sensitivities using sensitivity.
* `server` - IPv4 or IPv6 address of the server.
* `server_type` - Protocol used to communicate with the file server. Currently only Samba (SMB) servers are supported. Valid values: `samba` .
* `tod_hour` - Hour of the day on which to scan the server (0 - 23, default = 1).
* `tod_min` - Minute of the hour on which to scan the server (0 - 59).
* `username` - User name required to log into the file server.
* `vdom` - Select the VDOM that can communicate with the file server. Valid values: `mgmt` `current` .
* `weekday` - Day of the week on which to scan the server. Valid values: `sunday` `monday` `tuesday` `wednesday` `thursday` `friday` `saturday` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_dlp_fpdocsource can be imported using:
```sh
terraform import fortios_dlp_fpdocsource.labelname {{mkey}}
```
