---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_sniffer"
description: |-
  Get information on a fortios Configure sniffer.
---

# Data Source: fortios_firewall_sniffer
Use this data source to get information on a fortios Configure sniffer.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Sniffer ID (0 - 9999).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `application_list` - Name of an existing application list.
* `application_list_status` - Enable/disable application control profile.
* `av_profile` - Name of an existing antivirus profile.
* `av_profile_status` - Enable/disable antivirus profile.
* `dlp_sensor` - Name of an existing DLP sensor.
* `dlp_sensor_status` - Enable/disable DLP sensor.
* `dsri` - Enable/disable DSRI.
* `emailfilter_profile` - Name of an existing email filter profile.
* `emailfilter_profile_status` - Enable/disable emailfilter.
* `file_filter_profile` - Name of an existing file-filter profile.
* `file_filter_profile_status` - Enable/disable file filter.
* `host` - Hosts to filter for in sniffer traffic (Format examples: 1.1.1.1, 2.2.2.0/24, 3.3.3.3/255.255.255.0, 4.4.4.0-4.4.4.240).
* `id` - Sniffer ID (0 - 9999).
* `interface` - Interface name that traffic sniffing will take place on.
* `ip_threatfeed_status` - Enable/disable IP threat feed.
* `ips_dos_status` - Enable/disable IPS DoS anomaly detection.
* `ips_sensor` - Name of an existing IPS sensor.
* `ips_sensor_status` - Enable/disable IPS sensor.
* `ipv6` - Enable/disable sniffing IPv6 packets.
* `logtraffic` - Either log all sessions, only sessions that have a security profile applied, or disable all logging for this policy.
* `max_packet_count` - Maximum packet count (1 - 1000000, default = 4000).
* `non_ip` - Enable/disable sniffing non-IP packets.
* `port` - Ports to sniff (Format examples: 10, :20, 30:40, 50-, 100-200).
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `status` - Enable/disable the active status of the sniffer.
* `vlan` - List of VLANs to sniff.
* `webfilter_profile` - Name of an existing web filter profile.
* `webfilter_profile_status` - Enable/disable web filter profile.
* `anomaly` - Configuration method to edit Denial of Service (DoS) anomaly settings.The structure of `anomaly` block is documented below.

The `anomaly` block contains:

* `action` - Action taken when the threshold is reached.
* `log` - Enable/disable anomaly logging.
* `name` - Anomaly name.
* `quarantine` - Quarantine method.
* `quarantine_expiry` - Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.
* `quarantine_log` - Enable/disable quarantine logging.
* `status` - Enable/disable this anomaly.
* `threshold` - Anomaly threshold. Number of detected instances per minute that triggers the anomaly action.
* `thresholddefault` - Number of detected instances per minute which triggers action (1 - 2147483647, default = 1000). Note that each anomaly has a different threshold value assigned to it.
* `ip_threatfeed` - Name of an existing IP threat feed.The structure of `ip_threatfeed` block is documented below.

The `ip_threatfeed` block contains:

* `name` - Threat feed name.
