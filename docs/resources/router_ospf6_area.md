---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf6_area"
description: |-
  OSPF6 area configuration.
---

## fortios_router_ospf6_area
OSPF6 area configuration.

~> This resource is configuring a child table of the parent resource: `fortios_router_ospf6`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.


## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `authentication` - Authentication mode. Valid values: `none` `ah` `esp` .
* `default_cost` - Summary default cost of stub or NSSA area.
* `id` - Area entry IP address.
* `ipsec_auth_alg` - Authentication algorithm. Valid values: `md5` `sha1` `sha256` `sha384` `sha512` .
* `ipsec_enc_alg` - Encryption algorithm. Valid values: `null` `des` `3des` `aes128` `aes192` `aes256` .
* `key_rollover_interval` - Key roll-over interval.
* `nssa_default_information_originate` - Enable/disable originate type 7 default into NSSA area. Valid values: `enable` `disable` .
* `nssa_default_information_originate_metric` - OSPFv3 default metric.
* `nssa_default_information_originate_metric_type` - OSPFv3 metric type for default routes. Valid values: `1` `2` .
* `nssa_redistribution` - Enable/disable redistribute into NSSA area. Valid values: `enable` `disable` .
* `nssa_translator_role` - NSSA translator role type. Valid values: `candidate` `never` `always` .
* `stub_type` - Stub summary setting. Valid values: `no-summary` `summary` .
* `type` - Area type setting. Valid values: `regular` `nssa` `stub` .
* `ipsec_keys` - IPsec authentication and encryption keys. The structure of `ipsec_keys` block is documented below.

The `ipsec_keys` block contains:

* `auth_key` - Authentication key.
* `enc_key` - Encryption key.
* `spi` - Security Parameters Index.
* `range` - OSPF6 area range configuration. The structure of `range` block is documented below.

The `range` block contains:

* `advertise` - Enable/disable advertise status. Valid values: `disable` `enable` .
* `id` - Range entry ID.
* `prefix6` - IPv6 prefix.
* `virtual_link` - OSPF6 virtual link configuration. The structure of `virtual_link` block is documented below.

The `virtual_link` block contains:

* `authentication` - Authentication mode. Valid values: `none` `ah` `esp` `area` .
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `ipsec_auth_alg` - Authentication algorithm. Valid values: `md5` `sha1` `sha256` `sha384` `sha512` .
* `ipsec_enc_alg` - Encryption algorithm. Valid values: `null` `des` `3des` `aes128` `aes192` `aes256` .
* `key_rollover_interval` - Key roll-over interval.
* `name` - Virtual link entry name.
* `peer` - A.B.C.D, peer router ID.
* `retransmit_interval` - Retransmit interval.
* `transmit_delay` - Transmit delay.
* `ipsec_keys` - IPsec authentication and encryption keys. The structure of `ipsec_keys` block is documented below.

The `ipsec_keys` block contains:

* `auth_key` - Authentication key.
* `enc_key` - Encryption key.
* `spi` - Security Parameters Index.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_area can be imported using:
```sh
terraform import fortios_router_area.labelname {{mkey}}
```
