---
subcategory: "FortiGate Dlp"
layout: "fortios"
page_title: "FortiOS: fortios_dlp_sensor"
description: |-
  Get information on a fortios Configure DLP sensors.
---

# Data Source: fortios_dlp_sensor
Use this data source to get information on a fortios Configure DLP sensors.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name of the DLP sensor.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `dlp_log` - Enable/disable DLP logging.
* `extended_log` - Enable/disable extended logging for data leak prevention.
* `feature_set` - Flow/proxy feature set.
* `full_archive_proto` - Protocols to always content archive.
* `nac_quar_log` - Enable/disable NAC quarantine logging.
* `name` - Name of the DLP sensor.
* `options` - Configure DLP options.
* `replacemsg_group` - Replacement message group used by this DLP sensor.
* `summary_proto` - Protocols to always log summary.
* `filter` - Set up DLP filters for this sensor.The structure of `filter` block is documented below.

The `filter` block contains:

* `action` - Action to take with content that this DLP sensor matches.
* `archive` - Enable/disable DLP archiving.
* `company_identifier` - Enter a company identifier watermark to match. Only watermarks that your company has placed on the files are matched.
* `expiry` - Quarantine duration in days, hours, minutes (format = dddhhmm).
* `file_size` - Match files this size or larger (0 - 4294967295 kbytes).
* `file_type` - Select the number of a DLP file pattern table to match.
* `filter_by` - Select the type of content to match.
* `id` - ID.
* `match_percentage` - Percentage of fingerprints in the fingerprint databases designated with the selected sensitivity to match.
* `name` - Filter name.
* `proto` - Check messages or files over one or more of these protocols.
* `regexp` - Enter a regular expression to match (max. 255 characters).
* `severity` - Select the severity or threat level that matches this filter.
* `type` - Select whether to check the content of messages (an email message) or files (downloaded files or email attachments).
* `sensitivity` - Select a DLP file pattern sensitivity to match.The structure of `sensitivity` block is documented below.

The `sensitivity` block contains:

* `name` - Select a DLP sensitivity.
