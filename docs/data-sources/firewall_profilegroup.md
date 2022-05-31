---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_profilegroup"
description: |-
  Get information on a fortios Configure profile groups.
---

# Data Source: fortios_firewall_profilegroup
Use this data source to get information on a fortios Configure profile groups.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Profile group name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `application_list` - Name of an existing Application list.
* `av_profile` - Name of an existing Antivirus profile.
* `cifs_profile` - Name of an existing CIFS profile.
* `dlp_profile` - Name of an existing DLP profile.
* `dlp_sensor` - Name of an existing DLP sensor.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `emailfilter_profile` - Name of an existing email filter profile.
* `file_filter_profile` - Name of an existing file-filter profile.
* `icap_profile` - Name of an existing ICAP profile.
* `ips_sensor` - Name of an existing IPS sensor.
* `name` - Profile group name.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `sctp_filter_profile` - Name of an existing SCTP filter profile.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `videofilter_profile` - Name of an existing VideoFilter profile.
* `voip_profile` - Name of an existing VoIP profile.
* `waf_profile` - Name of an existing Web application firewall profile.
* `webfilter_profile` - Name of an existing Web filter profile.
