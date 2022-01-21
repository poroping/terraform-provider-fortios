---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_fec"
description: |-
  Get information on a fortios Configure Forward Error Correction (FEC) mapping profiles.
---

# Data Source: fortios_vpnipsec_fec
Use this data source to get information on a fortios Configure Forward Error Correction (FEC) mapping profiles.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - Profile name.
* `mappings` - FEC redundancy mapping table.The structure of `mappings` block is documented below.

The `mappings` block contains:

* `bandwidth_bi_threshold` - Apply FEC parameters when available bi-bandwidth is >= threshold (kbps, 0 means no threshold).
* `bandwidth_down_threshold` - Apply FEC parameters when available down bandwidth is >= threshold (kbps, 0 means no threshold).
* `bandwidth_up_threshold` - Apply FEC parameters when available up bandwidth is >= threshold (kbps, 0 means no threshold).
* `base` - Number of base FEC packets (1 - 20).
* `latency_threshold` - Apply FEC parameters when latency is <= threshold (0 means no threshold).
* `packet_loss_threshold` - Apply FEC parameters when packet loss is >= threshold (0 - 100, 0 means no threshold).
* `redundant` - Number of redundant FEC packets (1 - 5).
* `seqno` - Sequence number (1 - 64).
