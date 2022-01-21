---
subcategory: "FortiGate Ips"
layout: "fortios"
page_title: "FortiOS: fortios_ips_settings"
description: |-
  Get information on a fortios Configure IPS VDOM parameter.
---

# Data Source: fortios_ips_settings
Use this data source to get information on a fortios Configure IPS VDOM parameter.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `ips_packet_quota` - Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
* `packet_log_history` - Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
* `packet_log_memory` - Maximum memory can be used by packet log (64 - 8192 kB).
* `packet_log_post_attack` - Number of packets to log after the IPS signature is detected (0 - 255).
