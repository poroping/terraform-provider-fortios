---
subcategory: "FortiGate Ips"
layout: "fortios"
page_title: "FortiOS: fortios_ips_global"
description: |-
  Configure IPS global parameter.
---

## fortios_ips_global
Configure IPS global parameter.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `anomaly_mode` - Global blocking mode for rate-based anomalies. Valid values: `periodical` `continuous` .
* `cp_accel_mode` - IPS Pattern matching acceleration/offloading to CPx processors. Valid values: `none` `basic` `advanced` .
* `database` - Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks. Valid values: `regular` `extended` .
* `deep_app_insp_db_limit` - Limit on number of entries in deep application inspection database (1 - 2147483647, use recommended setting = 0).
* `deep_app_insp_timeout` - Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).
* `engine_count` - Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.
* `exclude_signatures` - Excluded signatures. Valid values: `none` `industrial` .
* `fail_open` - Enable to allow traffic if the IPS buffer is full. Default is disable and IPS traffic is blocked when the IPS buffer is full. Valid values: `enable` `disable` .
* `intelligent_mode` - Enable/disable IPS adaptive scanning (intelligent mode). Intelligent mode optimizes the scanning method for the type of traffic. Valid values: `enable` `disable` .
* `ips_reserve_cpu` - Enable/disable IPS daemon's use of CPUs other than CPU 0. Valid values: `disable` `enable` .
* `ngfw_max_scan_range` - NGFW policy-mode app detection threshold.
* `np_accel_mode` - Acceleration mode for IPS processing by NPx processors. Valid values: `none` `basic` .
* `packet_log_queue_depth` - Packet/pcap log queue depth per IPS engine.
* `session_limit_mode` - Method of counting concurrent sessions used by session limit anomalies. Choose between greater accuracy (accurate) or improved performance (heuristics). Valid values: `accurate` `heuristic` .
* `skype_client_public_ipaddr` - Public IP addresses of your network that receive Skype sessions. Helps identify Skype sessions. Separate IP addresses with commas.
* `socket_size` - IPS socket buffer size. Max and default value depend on available memory. Can be changed to tune performance.
* `sync_session_ttl` - Enable/disable use of kernel session TTL for IPS sessions. Valid values: `enable` `disable` .
* `traffic_submit` - Enable/disable submitting attack data found by this FortiGate to FortiGuard. Valid values: `enable` `disable` .
* `tls_active_probe` - TLS active probe configuration. The structure of `tls_active_probe` block is documented below.

The `tls_active_probe` block contains:

* `interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `source_ip` - Source IP address used for TLS active probe.
* `source_ip6` - Source IPv6 address used for TLS active probe.
* `vdom` - Virtual domain name for TLS active probe. This attribute must reference one of the following datasources: `system.vdom.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_ips_global can be imported using:
```sh
terraform import fortios_ips_global.labelname {{mkey}}
```
