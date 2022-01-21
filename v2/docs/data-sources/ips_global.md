---
subcategory: "FortiGate Ips"
layout: "fortios"
page_title: "FortiOS: fortios_ips_global"
description: |-
  Get information on a fortios Configure IPS global parameter.
---

# Data Source: fortios_ips_global
Use this data source to get information on a fortios Configure IPS global parameter.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `anomaly_mode` - Global blocking mode for rate-based anomalies.
* `cp_accel_mode` - IPS Pattern matching acceleration/offloading to CPx processors.
* `database` - Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks.
* `deep_app_insp_db_limit` - Limit on number of entries in deep application inspection database (1 - 2147483647, 0 = use recommended setting)
* `deep_app_insp_timeout` - Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).
* `engine_count` - Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.
* `exclude_signatures` - Excluded signatures.
* `fail_open` - Enable to allow traffic if the IPS buffer is full. Default is disable and IPS traffic is blocked when the IPS buffer is full.
* `intelligent_mode` - Enable/disable IPS adaptive scanning (intelligent mode). Intelligent mode optimizes the scanning method for the type of traffic.
* `ips_reserve_cpu` - Enable/disable IPS daemon's use of CPUs other than CPU 0
* `ngfw_max_scan_range` - NGFW policy-mode app detection threshold.
* `np_accel_mode` - Acceleration mode for IPS processing by NPx processors.
* `packet_log_queue_depth` - Packet/pcap log queue depth per IPS engine.
* `session_limit_mode` - Method of counting concurrent sessions used by session limit anomalies. Choose between greater accuracy (accurate) or improved performance (heuristics).
* `skype_client_public_ipaddr` - Public IP addresses of your network that receive Skype sessions. Helps identify Skype sessions. Separate IP addresses with commas.
* `socket_size` - IPS socket buffer size. Max and default value depend on available memory. Can be changed to tune performance.
* `sync_session_ttl` - Enable/disable use of kernel session TTL for IPS sessions.
* `traffic_submit` - Enable/disable submitting attack data found by this FortiGate to FortiGuard.
* `tls_active_probe` - TLS active probe configuration.The structure of `tls_active_probe` block is documented below.

The `tls_active_probe` block contains:

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `source_ip` - Source IP address used for TLS active probe.
* `source_ip6` - Source IPv6 address used for TLS active probe.
* `vdom` - Virtual domain name for TLS active probe.
