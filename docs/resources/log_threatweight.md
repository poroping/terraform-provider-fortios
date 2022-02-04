---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_log_threatweight"
description: |-
  Configure threat weight settings.
---

## fortios_log_threatweight
Configure threat weight settings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `blocked_connection` - Threat weight score for blocked connections. Valid values: `disable` `low` `medium` `high` `critical` .
* `botnet_connection_detected` - Threat weight score for detected botnet connections. Valid values: `disable` `low` `medium` `high` `critical` .
* `failed_connection` - Threat weight score for failed connections. Valid values: `disable` `low` `medium` `high` `critical` .
* `status` - Enable/disable the threat weight feature. Valid values: `enable` `disable` .
* `url_block_detected` - Threat weight score for URL blocking. Valid values: `disable` `low` `medium` `high` `critical` .
* `application` - Application-control threat weight settings. The structure of `application` block is documented below.

The `application` block contains:

* `category` - Application category.
* `id` - Entry ID.
* `level` - Threat weight score for Application events. Valid values: `disable` `low` `medium` `high` `critical` .
* `geolocation` - Geolocation-based threat weight settings. The structure of `geolocation` block is documented below.

The `geolocation` block contains:

* `country` - Country code.
* `id` - Entry ID.
* `level` - Threat weight score for Geolocation-based events. Valid values: `disable` `low` `medium` `high` `critical` .
* `ips` - IPS threat weight settings. The structure of `ips` block is documented below.

The `ips` block contains:

* `critical_severity` - Threat weight score for IPS critical severity events. Valid values: `disable` `low` `medium` `high` `critical` .
* `high_severity` - Threat weight score for IPS high severity events. Valid values: `disable` `low` `medium` `high` `critical` .
* `info_severity` - Threat weight score for IPS info severity events. Valid values: `disable` `low` `medium` `high` `critical` .
* `low_severity` - Threat weight score for IPS low severity events. Valid values: `disable` `low` `medium` `high` `critical` .
* `medium_severity` - Threat weight score for IPS medium severity events. Valid values: `disable` `low` `medium` `high` `critical` .
* `level` - Score mapping for threat weight levels. The structure of `level` block is documented below.

The `level` block contains:

* `critical` - Critical level score value (1 - 100).
* `high` - High level score value (1 - 100).
* `low` - Low level score value (1 - 100).
* `medium` - Medium level score value (1 - 100).
* `malware` - Anti-virus malware threat weight settings. The structure of `malware` block is documented below.

The `malware` block contains:

* `command_blocked` - Threat weight score for blocked command detected. Valid values: `disable` `low` `medium` `high` `critical` .
* `content_disarm` - Threat weight score for virus (content disarm) detected. Valid values: `disable` `low` `medium` `high` `critical` .
* `ems_threat_feed` - Threat weight score for virus (EMS threat feed) detected. Valid values: `disable` `low` `medium` `high` `critical` .
* `file_blocked` - Threat weight score for blocked file detected. Valid values: `disable` `low` `medium` `high` `critical` .
* `fortiai` - Threat weight score for FortiAI-detected virus. Valid values: `disable` `low` `medium` `high` `critical` .
* `fsa_high_risk` - Threat weight score for FortiSandbox high risk malware detected. Valid values: `disable` `low` `medium` `high` `critical` .
* `fsa_malicious` - Threat weight score for FortiSandbox malicious malware detected. Valid values: `disable` `low` `medium` `high` `critical` .
* `fsa_medium_risk` - Threat weight score for FortiSandbox medium risk malware detected. Valid values: `disable` `low` `medium` `high` `critical` .
* `malware_list` - Threat weight score for virus (malware list) detected. Valid values: `disable` `low` `medium` `high` `critical` .
* `mimefragmented` - Threat weight score for mimefragmented detected. Valid values: `disable` `low` `medium` `high` `critical` .
* `oversized` - Threat weight score for oversized file detected. Valid values: `disable` `low` `medium` `high` `critical` .
* `switch_proto` - Threat weight score for switch proto detected. Valid values: `disable` `low` `medium` `high` `critical` .
* `virus_file_type_executable` - Threat weight score for virus (file type executable) detected. Valid values: `disable` `low` `medium` `high` `critical` .
* `virus_infected` - Threat weight score for virus (infected) detected. Valid values: `disable` `low` `medium` `high` `critical` .
* `virus_outbreak_prevention` - Threat weight score for virus (outbreak prevention) event. Valid values: `disable` `low` `medium` `high` `critical` .
* `virus_scan_error` - Threat weight score for virus (scan error) detected. Valid values: `disable` `low` `medium` `high` `critical` .
* `web` - Web filtering threat weight settings. The structure of `web` block is documented below.

The `web` block contains:

* `category` - Threat weight score for web category filtering matches.
* `id` - Entry ID.
* `level` - Threat weight score for web category filtering matches. Valid values: `disable` `low` `medium` `high` `critical` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_log_threatweight can be imported using:
```sh
terraform import fortios_log_threatweight.labelname {{mkey}}
```
