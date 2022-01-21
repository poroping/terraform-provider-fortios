---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf6_area"
description: |-
  Get information on a fortios OSPF6 area configuration.
---

# Data Source: fortios_router_ospf6_area
Use this data source to get information on a fortios OSPF6 area configuration.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Area entry IP address.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `authentication` - Authentication mode.
* `default_cost` - Summary default cost of stub or NSSA area.
* `id` - Area entry IP address.
* `ipsec_auth_alg` - Authentication algorithm.
* `ipsec_enc_alg` - Encryption algorithm.
* `key_rollover_interval` - Key roll-over interval.
* `nssa_default_information_originate` - Enable/disable originate type 7 default into NSSA area.
* `nssa_default_information_originate_metric` - OSPFv3 default metric.
* `nssa_default_information_originate_metric_type` - OSPFv3 metric type for default routes.
* `nssa_redistribution` - Enable/disable redistribute into NSSA area.
* `nssa_translator_role` - NSSA translator role type.
* `stub_type` - Stub summary setting.
* `type` - Area type setting.
* `ipsec_keys` - IPsec authentication and encryption keys.The structure of `ipsec_keys` block is documented below.

The `ipsec_keys` block contains:

* `auth_key` - Authentication key.
* `enc_key` - Encryption key.
* `spi` - Security Parameters Index.
* `range` - OSPF6 area range configuration.The structure of `range` block is documented below.

The `range` block contains:

* `advertise` - Enable/disable advertise status.
* `id` - Range entry ID.
* `prefix6` - IPv6 prefix.
* `virtual_link` - OSPF6 virtual link configuration.The structure of `virtual_link` block is documented below.

The `virtual_link` block contains:

* `authentication` - Authentication mode.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `ipsec_auth_alg` - Authentication algorithm.
* `ipsec_enc_alg` - Encryption algorithm.
* `key_rollover_interval` - Key roll-over interval.
* `name` - Virtual link entry name.
* `peer` - A.B.C.D, peer router ID.
* `retransmit_interval` - Retransmit interval.
* `transmit_delay` - Transmit delay.
* `ipsec_keys` - IPsec authentication and encryption keys.The structure of `ipsec_keys` block is documented below.

The `ipsec_keys` block contains:

* `auth_key` - Authentication key.
* `enc_key` - Encryption key.
* `spi` - Security Parameters Index.
