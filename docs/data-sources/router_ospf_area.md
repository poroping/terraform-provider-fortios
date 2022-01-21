---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf_area"
description: |-
  Get information on a fortios OSPF area configuration.
---

# Data Source: fortios_router_ospf_area
Use this data source to get information on a fortios OSPF area configuration.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Area entry IP address.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `authentication` - Authentication type.
* `comments` - Comment.
* `default_cost` - Summary default cost of stub or NSSA area.
* `id` - Area entry IP address.
* `nssa_default_information_originate` - Redistribute, advertise, or do not originate Type-7 default route into NSSA area.
* `nssa_default_information_originate_metric` - OSPF default metric.
* `nssa_default_information_originate_metric_type` - OSPF metric type for default routes.
* `nssa_redistribution` - Enable/disable redistribute into NSSA area.
* `nssa_translator_role` - NSSA translator role type.
* `shortcut` - Enable/disable shortcut option.
* `stub_type` - Stub summary setting.
* `type` - Area type setting.
* `filter_list` - OSPF area filter-list configuration.The structure of `filter_list` block is documented below.

The `filter_list` block contains:

* `direction` - Direction.
* `id` - Filter list entry ID.
* `list` - Access-list or prefix-list name.
* `range` - OSPF area range configuration.The structure of `range` block is documented below.

The `range` block contains:

* `advertise` - Enable/disable advertise status.
* `id` - Range entry ID.
* `prefix` - Prefix.
* `substitute` - Substitute prefix.
* `substitute_status` - Enable/disable substitute status.
* `virtual_link` - OSPF virtual link configuration.The structure of `virtual_link` block is documented below.

The `virtual_link` block contains:

* `authentication` - Authentication type.
* `authentication_key` - Authentication key.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `keychain` - Message-digest key-chain name.
* `md5_keychain` - Authentication MD5 key-chain name.
* `name` - Virtual link entry name.
* `peer` - Peer IP.
* `retransmit_interval` - Retransmit interval.
* `transmit_delay` - Transmit delay.
* `md5_keys` - MD5 key.The structure of `md5_keys` block is documented below.

The `md5_keys` block contains:

* `id` - Key ID (1 - 255).
* `key_string` - Password for the key.
