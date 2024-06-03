---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_sniffer"
description: |-
  Configure sniffer.
---

## fortios_firewall_sniffer
Configure sniffer.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `application_list` - Name of an existing application list. This attribute must reference one of the following datasources: `application.list.name` .
* `application_list_status` - Enable/disable application control profile. Valid values: `enable` `disable` .
* `av_profile` - Name of an existing antivirus profile. This attribute must reference one of the following datasources: `antivirus.profile.name` .
* `av_profile_status` - Enable/disable antivirus profile. Valid values: `enable` `disable` .
* `dlp_profile` - Name of an existing DLP profile. This attribute must reference one of the following datasources: `dlp.profile.name` .
* `dlp_profile_status` - Enable/disable DLP profile. Valid values: `enable` `disable` .
* `dlp_sensor` - Name of an existing DLP sensor. This attribute must reference one of the following datasources: `dlp.sensor.name` .
* `dlp_sensor_status` - Enable/disable DLP sensor. Valid values: `enable` `disable` .
* `dsri` - Enable/disable DSRI. Valid values: `enable` `disable` .
* `emailfilter_profile` - Name of an existing email filter profile. This attribute must reference one of the following datasources: `emailfilter.profile.name` .
* `emailfilter_profile_status` - Enable/disable emailfilter. Valid values: `enable` `disable` .
* `file_filter_profile` - Name of an existing file-filter profile. This attribute must reference one of the following datasources: `file-filter.profile.name` .
* `file_filter_profile_status` - Enable/disable file filter. Valid values: `enable` `disable` .
* `host` - Hosts to filter for in sniffer traffic (Format examples: 1.1.1.1, 2.2.2.0/24, 3.3.3.3/255.255.255.0, 4.4.4.0-4.4.4.240).
* `id` - Sniffer ID (0 - 9999).
* `interface` - Interface name that traffic sniffing will take place on. This attribute must reference one of the following datasources: `system.interface.name` .
* `ip_threatfeed_status` - Enable/disable IP threat feed. Valid values: `enable` `disable` .
* `ips_dos_status` - Enable/disable IPS DoS anomaly detection. Valid values: `enable` `disable` .
* `ips_sensor` - Name of an existing IPS sensor. This attribute must reference one of the following datasources: `ips.sensor.name` .
* `ips_sensor_status` - Enable/disable IPS sensor. Valid values: `enable` `disable` .
* `ipv6` - Enable/disable sniffing IPv6 packets. Valid values: `enable` `disable` .
* `logtraffic` - Either log all sessions, only sessions that have a security profile applied, or disable all logging for this policy. Valid values: `all` `utm` `disable` .
* `max_packet_count` - Maximum packet count (1 - 1000000, default = 4000).
* `non_ip` - Enable/disable sniffing non-IP packets. Valid values: `enable` `disable` .
* `port` - Ports to sniff (Format examples: 10, :20, 30:40, 50-, 100-200).
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `status` - Enable/disable the active status of the sniffer. Valid values: `enable` `disable` .
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `vlan` - List of VLANs to sniff.
* `webfilter_profile` - Name of an existing web filter profile. This attribute must reference one of the following datasources: `webfilter.profile.name` .
* `webfilter_profile_status` - Enable/disable web filter profile. Valid values: `enable` `disable` .
* `anomaly` - Configuration method to edit Denial of Service (DoS) anomaly settings. The structure of `anomaly` block is documented below.

The `anomaly` block contains:

* `action` - Action taken when the threshold is reached. Valid values: `pass` `block` `proxy` .
* `log` - Enable/disable anomaly logging. Valid values: `enable` `disable` .
* `name` - Anomaly name.
* `quarantine` - Quarantine method. Valid values: `none` `attacker` .
* `quarantine_expiry` - Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.
* `quarantine_log` - Enable/disable quarantine logging. Valid values: `disable` `enable` .
* `status` - Enable/disable this anomaly. Valid values: `disable` `enable` .
* `synproxy_tcp_mss` - Determine TCP maximum segment size (MSS) value for packets replied by syn proxy module. Valid values: `0` `256` `512` `1024` `1300` `1360` `1460` `1500` .
* `synproxy_tcp_sack` - enable/disable TCP selective acknowledage (SACK) for packets replied by syn proxy module. Valid values: `enable` `disable` .
* `synproxy_tcp_timestamp` - enable/disable TCP timestamp option for packets replied by syn proxy module. Valid values: `enable` `disable` .
* `synproxy_tcp_window` - Determine TCP Window size for packets replied by syn proxy module. Valid values: `4096` `8192` `16384` `32768` .
* `synproxy_tcp_windowscale` - Determine TCP window scale option value for packets replied by syn proxy module. Valid values: `0` `1` `2` `3` `4` `5` `6` `7` `8` `9` `10` `11` `12` `13` `14` .
* `synproxy_tos` - Determine TCP differentiated services code point value (type of service). Valid values: `0` `10` `12` `14` `18` `20` `22` `26` `28` `30` `34` `36` `38` `40` `46` `255` .
* `synproxy_ttl` - Determine Time to live (TTL) value for packets replied by syn proxy module. Valid values: `32` `64` `128` `255` .
* `threshold` - Anomaly threshold. Number of detected instances (packets per second or concurrent session number) that triggers the anomaly action.
* `thresholddefault` - Number of detected instances (packets per second or concurrent session number) which triggers action (1 - 2147483647, default = 1000). Note that each anomaly has a different threshold value assigned to it.
* `ip_threatfeed` - Name of an existing IP threat feed. The structure of `ip_threatfeed` block is documented below.

The `ip_threatfeed` block contains:

* `name` - Threat feed name. This attribute must reference one of the following datasources: `system.external-resource.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_sniffer can be imported using:
```sh
terraform import fortios_firewall_sniffer.labelname {{mkey}}
```
