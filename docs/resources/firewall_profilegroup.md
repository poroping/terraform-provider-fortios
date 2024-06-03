---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_profilegroup"
description: |-
  Configure profile groups.
---

## fortios_firewall_profilegroup
Configure profile groups.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `application_list` - Name of an existing Application list. This attribute must reference one of the following datasources: `application.list.name` .
* `av_profile` - Name of an existing Antivirus profile. This attribute must reference one of the following datasources: `antivirus.profile.name` .
* `cifs_profile` - Name of an existing CIFS profile. This attribute must reference one of the following datasources: `cifs.profile.name` .
* `dlp_profile` - Name of an existing DLP profile. This attribute must reference one of the following datasources: `dlp.profile.name` .
* `dlp_sensor` - Name of an existing DLP sensor. This attribute must reference one of the following datasources: `dlp.sensor.name` .
* `dnsfilter_profile` - Name of an existing DNS filter profile. This attribute must reference one of the following datasources: `dnsfilter.profile.name` .
* `emailfilter_profile` - Name of an existing email filter profile. This attribute must reference one of the following datasources: `emailfilter.profile.name` .
* `file_filter_profile` - Name of an existing file-filter profile. This attribute must reference one of the following datasources: `file-filter.profile.name` .
* `icap_profile` - Name of an existing ICAP profile. This attribute must reference one of the following datasources: `icap.profile.name` .
* `ips_sensor` - Name of an existing IPS sensor. This attribute must reference one of the following datasources: `ips.sensor.name` .
* `ips_voip_filter` - Name of an existing VoIP (ips) profile. This attribute must reference one of the following datasources: `voip.profile.name` .
* `name` - Profile group name.
* `profile_protocol_options` - Name of an existing Protocol options profile. This attribute must reference one of the following datasources: `firewall.profile-protocol-options.name` .
* `sctp_filter_profile` - Name of an existing SCTP filter profile. This attribute must reference one of the following datasources: `sctp-filter.profile.name` .
* `ssh_filter_profile` - Name of an existing SSH filter profile. This attribute must reference one of the following datasources: `ssh-filter.profile.name` .
* `ssl_ssh_profile` - Name of an existing SSL SSH profile. This attribute must reference one of the following datasources: `firewall.ssl-ssh-profile.name` .
* `videofilter_profile` - Name of an existing VideoFilter profile. This attribute must reference one of the following datasources: `videofilter.profile.name` .
* `voip_profile` - Name of an existing VoIP (voipd) profile. This attribute must reference one of the following datasources: `voip.profile.name` .
* `waf_profile` - Name of an existing Web application firewall profile. This attribute must reference one of the following datasources: `waf.profile.name` .
* `webfilter_profile` - Name of an existing Web filter profile. This attribute must reference one of the following datasources: `webfilter.profile.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_profilegroup can be imported using:
```sh
terraform import fortios_firewall_profilegroup.labelname {{mkey}}
```
