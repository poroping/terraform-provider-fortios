---
subcategory: "FortiGate Dlp"
layout: "fortios"
page_title: "FortiOS: fortios_dlp_sensor"
description: |-
  Configure DLP sensors.
---

## fortios_dlp_sensor
Configure DLP sensors.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comment` - Comment.
* `dlp_log` - Enable/disable DLP logging. Valid values: `enable` `disable` .
* `extended_log` - Enable/disable extended logging for data leak prevention. Valid values: `enable` `disable` .
* `feature_set` - Flow/proxy feature set. Valid values: `flow` `proxy` .
* `full_archive_proto` - Protocols to always content archive. Valid values: `smtp` `pop3` `imap` `http-get` `http-post` `ftp` `nntp` `mapi` `ssh` `cifs` .
* `nac_quar_log` - Enable/disable NAC quarantine logging. Valid values: `enable` `disable` .
* `name` - Name of the DLP sensor.
* `options` - Configure DLP options. Valid values: .
* `replacemsg_group` - Replacement message group used by this DLP sensor. This attribute must reference one of the following datasources: `system.replacemsg-group.name` .
* `summary_proto` - Protocols to always log summary. Valid values: `smtp` `pop3` `imap` `http-get` `http-post` `ftp` `nntp` `mapi` `ssh` `cifs` .
* `filter` - Set up DLP filters for this sensor. The structure of `filter` block is documented below.

The `filter` block contains:

* `action` - Action to take with content that this DLP sensor matches. Valid values: `allow` `log-only` `block` `quarantine-ip` .
* `archive` - Enable/disable DLP archiving. Valid values: `disable` `enable` .
* `company_identifier` - Enter a company identifier watermark to match. Only watermarks that your company has placed on the files are matched.
* `expiry` - Quarantine duration in days, hours, minutes (format = dddhhmm).
* `file_size` - Match files this size or larger (0 - 4294967295 kbytes).
* `file_type` - Select the number of a DLP file pattern table to match. This attribute must reference one of the following datasources: `dlp.filepattern.id` .
* `filter_by` - Select the type of content to match. Valid values: `credit-card` `ssn` `regexp` `file-type` `file-size` `fingerprint` `watermark` `encrypted` .
* `id` - ID.
* `match_percentage` - Percentage of fingerprints in the fingerprint databases designated with the selected sensitivity to match.
* `name` - Filter name.
* `proto` - Check messages or files over one or more of these protocols. Valid values: `smtp` `pop3` `imap` `http-get` `http-post` `ftp` `nntp` `mapi` `ssh` `cifs` .
* `regexp` - Enter a regular expression to match (max. 255 characters).
* `severity` - Select the severity or threat level that matches this filter. Valid values: `info` `low` `medium` `high` `critical` .
* `type` - Select whether to check the content of messages (an email message) or files (downloaded files or email attachments). Valid values: `file` `message` .
* `sensitivity` - Select a DLP file pattern sensitivity to match. The structure of `sensitivity` block is documented below.

The `sensitivity` block contains:

* `name` - Select a DLP sensitivity. This attribute must reference one of the following datasources: `dlp.sensitivity.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_dlp_sensor can be imported using:
```sh
terraform import fortios_dlp_sensor.labelname {{mkey}}
```
