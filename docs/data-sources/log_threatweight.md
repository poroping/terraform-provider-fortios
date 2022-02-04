---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_log_threatweight"
description: |-
  Get information on a fortios Configure threat weight settings.
---

# Data Source: fortios_log_threatweight
Use this data source to get information on a fortios Configure threat weight settings.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `blocked_connection` - Threat weight score for blocked connections.
* `botnet_connection_detected` - Threat weight score for detected botnet connections.
* `failed_connection` - Threat weight score for failed connections.
* `status` - Enable/disable the threat weight feature.
* `url_block_detected` - Threat weight score for URL blocking.
* `application` - Application-control threat weight settings.The structure of `application` block is documented below.

The `application` block contains:

* `category` - Application category.
* `id` - Entry ID.
* `level` - Threat weight score for Application events.
* `geolocation` - Geolocation-based threat weight settings.The structure of `geolocation` block is documented below.

The `geolocation` block contains:

* `country` - Country code.
* `id` - Entry ID.
* `level` - Threat weight score for Geolocation-based events.
* `ips` - IPS threat weight settings.The structure of `ips` block is documented below.

The `ips` block contains:

* `critical_severity` - Threat weight score for IPS critical severity events.
* `high_severity` - Threat weight score for IPS high severity events.
* `info_severity` - Threat weight score for IPS info severity events.
* `low_severity` - Threat weight score for IPS low severity events.
* `medium_severity` - Threat weight score for IPS medium severity events.
* `level` - Score mapping for threat weight levels.The structure of `level` block is documented below.

The `level` block contains:

* `critical` - Critical level score value (1 - 100).
* `high` - High level score value (1 - 100).
* `low` - Low level score value (1 - 100).
* `medium` - Medium level score value (1 - 100).
* `malware` - Anti-virus malware threat weight settings.The structure of `malware` block is documented below.

The `malware` block contains:

* `command_blocked` - Threat weight score for blocked command detected.
* `content_disarm` - Threat weight score for virus (content disarm) detected.
* `ems_threat_feed` - Threat weight score for virus (EMS threat feed) detected.
* `file_blocked` - Threat weight score for blocked file detected.
* `fortiai` - Threat weight score for FortiAI-detected virus.
* `fsa_high_risk` - Threat weight score for FortiSandbox high risk malware detected.
* `fsa_malicious` - Threat weight score for FortiSandbox malicious malware detected.
* `fsa_medium_risk` - Threat weight score for FortiSandbox medium risk malware detected.
* `malware_list` - Threat weight score for virus (malware list) detected.
* `mimefragmented` - Threat weight score for mimefragmented detected.
* `oversized` - Threat weight score for oversized file detected.
* `switch_proto` - Threat weight score for switch proto detected.
* `virus_file_type_executable` - Threat weight score for virus (file type executable) detected.
* `virus_infected` - Threat weight score for virus (infected) detected.
* `virus_outbreak_prevention` - Threat weight score for virus (outbreak prevention) event.
* `virus_scan_error` - Threat weight score for virus (scan error) detected.
* `web` - Web filtering threat weight settings.The structure of `web` block is documented below.

The `web` block contains:

* `category` - Threat weight score for web category filtering matches.
* `id` - Entry ID.
* `level` - Threat weight score for web category filtering matches.
