---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf_area"
description: |-
  OSPF area configuration.
---

## fortios_router_ospf_area
OSPF area configuration.

~> This resource is configuring a child table of the parent resource: `fortios_router_ospf`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.


## Example Usage

```hcl
# resource "fortios_router_ospf_area" "example" {
#   fosid = "0.0.0.0"
# }

### This resource may not work due to an API issue
### https://github.com/poroping/terraform-provider-fortios/issues/7
### WORKAROUND:

resource "fortios_router_ospf" "example" {
  area {
    id = "0.0.0.0"
  }

  lifecycle {
    ignore_changes = [
      "redistribute",
      "network",
      "neighbor",
      "ospf_interface",
      "summary_address",
      "distribute_list"
    ]
  }
}
```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `authentication` - Authentication type. Valid values: `none` `text` `message-digest` .
* `comments` - Comment.
* `default_cost` - Summary default cost of stub or NSSA area.
* `id` - Area entry IP address.
* `nssa_default_information_originate` - Redistribute, advertise, or do not originate Type-7 default route into NSSA area. Valid values: `enable` `always` `disable` .
* `nssa_default_information_originate_metric` - OSPF default metric.
* `nssa_default_information_originate_metric_type` - OSPF metric type for default routes. Valid values: `1` `2` .
* `nssa_redistribution` - Enable/disable redistribute into NSSA area. Valid values: `enable` `disable` .
* `nssa_translator_role` - NSSA translator role type. Valid values: `candidate` `never` `always` .
* `shortcut` - Enable/disable shortcut option. Valid values: `disable` `enable` `default` .
* `stub_type` - Stub summary setting. Valid values: `no-summary` `summary` .
* `type` - Area type setting. Valid values: `regular` `nssa` `stub` .
* `filter_list` - OSPF area filter-list configuration. The structure of `filter_list` block is documented below.

The `filter_list` block contains:

* `direction` - Direction. Valid values: `in` `out` .
* `id` - Filter list entry ID.
* `list` - Access-list or prefix-list name. This attribute must reference one of the following datasources: `router.access-list.name` `router.prefix-list.name` .
* `range` - OSPF area range configuration. The structure of `range` block is documented below.

The `range` block contains:

* `advertise` - Enable/disable advertise status. Valid values: `disable` `enable` .
* `id` - Range entry ID.
* `prefix` - Prefix.
* `substitute` - Substitute prefix.
* `substitute_status` - Enable/disable substitute status. Valid values: `enable` `disable` .
* `virtual_link` - OSPF virtual link configuration. The structure of `virtual_link` block is documented below.

The `virtual_link` block contains:

* `authentication` - Authentication type. Valid values: `none` `text` `message-digest` .
* `authentication_key` - Authentication key.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `keychain` - Message-digest key-chain name. This attribute must reference one of the following datasources: `router.key-chain.name` .
* `md5_keychain` - Authentication MD5 key-chain name. This attribute must reference one of the following datasources: `router.key-chain.name` .
* `name` - Virtual link entry name.
* `peer` - Peer IP.
* `retransmit_interval` - Retransmit interval.
* `transmit_delay` - Transmit delay.
* `md5_keys` - MD5 key. The structure of `md5_keys` block is documented below.

The `md5_keys` block contains:

* `id` - Key ID (1 - 255).
* `key_string` - Password for the key.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_area can be imported using:
```sh
terraform import fortios_router_area.labelname {{mkey}}
```
