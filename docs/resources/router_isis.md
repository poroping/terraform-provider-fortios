---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_isis"
description: |-
  Configure IS-IS.
---

## fortios_router_isis
Configure IS-IS.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `adjacency_check` - Enable/disable adjacency check. Valid values: `enable` `disable` .
* `adjacency_check6` - Enable/disable IPv6 adjacency check. Valid values: `enable` `disable` .
* `adv_passive_only` - Enable/disable IS-IS advertisement of passive interfaces only. Valid values: `enable` `disable` .
* `adv_passive_only6` - Enable/disable IPv6 IS-IS advertisement of passive interfaces only. Valid values: `enable` `disable` .
* `auth_keychain_l1` - Authentication key-chain for level 1 PDUs. This attribute must reference one of the following datasources: `router.key-chain.name` .
* `auth_keychain_l2` - Authentication key-chain for level 2 PDUs. This attribute must reference one of the following datasources: `router.key-chain.name` .
* `auth_mode_l1` - Level 1 authentication mode. Valid values: `password` `md5` .
* `auth_mode_l2` - Level 2 authentication mode. Valid values: `password` `md5` .
* `auth_password_l1` - Authentication password for level 1 PDUs.
* `auth_password_l2` - Authentication password for level 2 PDUs.
* `auth_sendonly_l1` - Enable/disable level 1 authentication send-only. Valid values: `enable` `disable` .
* `auth_sendonly_l2` - Enable/disable level 2 authentication send-only. Valid values: `enable` `disable` .
* `default_originate` - Enable/disable distribution of default route information. Valid values: `enable` `disable` .
* `default_originate6` - Enable/disable distribution of default IPv6 route information. Valid values: `enable` `disable` .
* `dynamic_hostname` - Enable/disable dynamic hostname. Valid values: `enable` `disable` .
* `ignore_lsp_errors` - Enable/disable ignoring of LSP errors with bad checksums. Valid values: `enable` `disable` .
* `is_type` - IS type. Valid values: `level-1-2` `level-1` `level-2-only` .
* `lsp_gen_interval_l1` - Minimum interval for level 1 LSP regenerating.
* `lsp_gen_interval_l2` - Minimum interval for level 2 LSP regenerating.
* `lsp_refresh_interval` - LSP refresh time in seconds.
* `max_lsp_lifetime` - Maximum LSP lifetime in seconds.
* `metric_style` - Use old-style (ISO 10589) or new-style packet formats. Valid values: `narrow` `wide` `transition` `narrow-transition` `narrow-transition-l1` `narrow-transition-l2` `wide-l1` `wide-l2` `wide-transition` `wide-transition-l1` `wide-transition-l2` `transition-l1` `transition-l2` .
* `overload_bit` - Enable/disable signal other routers not to use us in SPF. Valid values: `enable` `disable` .
* `overload_bit_on_startup` - Overload-bit only temporarily after reboot.
* `overload_bit_suppress` - Suppress overload-bit for the specific prefixes. Valid values: `external` `interlevel` .
* `redistribute_l1` - Enable/disable redistribution of level 1 routes into level 2. Valid values: `enable` `disable` .
* `redistribute_l1_list` - Access-list for route redistribution from l1 to l2. This attribute must reference one of the following datasources: `router.access-list.name` .
* `redistribute_l2` - Enable/disable redistribution of level 2 routes into level 1. Valid values: `enable` `disable` .
* `redistribute_l2_list` - Access-list for route redistribution from l2 to l1. This attribute must reference one of the following datasources: `router.access-list.name` .
* `redistribute6_l1` - Enable/disable redistribution of level 1 IPv6 routes into level 2. Valid values: `enable` `disable` .
* `redistribute6_l1_list` - Access-list for IPv6 route redistribution from l1 to l2. This attribute must reference one of the following datasources: `router.access-list6.name` .
* `redistribute6_l2` - Enable/disable redistribution of level 2 IPv6 routes into level 1. Valid values: `enable` `disable` .
* `redistribute6_l2_list` - Access-list for IPv6 route redistribution from l2 to l1. This attribute must reference one of the following datasources: `router.access-list6.name` .
* `spf_interval_exp_l1` - Level 1 SPF calculation delay.
* `spf_interval_exp_l2` - Level 2 SPF calculation delay.
* `isis_interface` - IS-IS interface configuration. The structure of `isis_interface` block is documented below.

The `isis_interface` block contains:

* `auth_keychain_l1` - Authentication key-chain for level 1 PDUs. This attribute must reference one of the following datasources: `router.key-chain.name` .
* `auth_keychain_l2` - Authentication key-chain for level 2 PDUs. This attribute must reference one of the following datasources: `router.key-chain.name` .
* `auth_mode_l1` - Level 1 authentication mode. Valid values: `md5` `password` .
* `auth_mode_l2` - Level 2 authentication mode. Valid values: `md5` `password` .
* `auth_password_l1` - Authentication password for level 1 PDUs.
* `auth_password_l2` - Authentication password for level 2 PDUs.
* `auth_send_only_l1` - Enable/disable authentication send-only for level 1 PDUs. Valid values: `enable` `disable` .
* `auth_send_only_l2` - Enable/disable authentication send-only for level 2 PDUs. Valid values: `enable` `disable` .
* `circuit_type` - IS-IS interface's circuit type. Valid values: `level-1-2` `level-1` `level-2` .
* `csnp_interval_l1` - Level 1 CSNP interval.
* `csnp_interval_l2` - Level 2 CSNP interval.
* `hello_interval_l1` - Level 1 hello interval.
* `hello_interval_l2` - Level 2 hello interval.
* `hello_multiplier_l1` - Level 1 multiplier for Hello holding time.
* `hello_multiplier_l2` - Level 2 multiplier for Hello holding time.
* `hello_padding` - Enable/disable padding to IS-IS hello packets. Valid values: `enable` `disable` .
* `lsp_interval` - LSP transmission interval (milliseconds).
* `lsp_retransmit_interval` - LSP retransmission interval (sec).
* `mesh_group` - Enable/disable IS-IS mesh group. Valid values: `enable` `disable` .
* `mesh_group_id` - Mesh group ID <0-4294967295>, 0: mesh-group blocked.
* `metric_l1` - Level 1 metric for interface.
* `metric_l2` - Level 2 metric for interface.
* `name` - IS-IS interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `network_type` - IS-IS interface's network type. Valid values: `broadcast` `point-to-point` `loopback` .
* `priority_l1` - Level 1 priority.
* `priority_l2` - Level 2 priority.
* `status` - Enable/disable interface for IS-IS. Valid values: `enable` `disable` .
* `status6` - Enable/disable IPv6 interface for IS-IS. Valid values: `enable` `disable` .
* `wide_metric_l1` - Level 1 wide metric for interface.
* `wide_metric_l2` - Level 2 wide metric for interface.
* `isis_net` - IS-IS net configuration. The structure of `isis_net` block is documented below.

The `isis_net` block contains:

* `id` - ISIS network ID.
* `net` - IS-IS networks (format = xx.xxxx.  .xxxx.xx.).
* `redistribute` - IS-IS redistribute protocols. The structure of `redistribute` block is documented below.

The `redistribute` block contains:

* `level` - Level. Valid values: `level-1-2` `level-1` `level-2` .
* `metric` - Metric.
* `metric_type` - Metric type. Valid values: `external` `internal` .
* `protocol` - Protocol name.
* `routemap` - Route map name. This attribute must reference one of the following datasources: `router.route-map.name` .
* `status` - Status. Valid values: `enable` `disable` .
* `redistribute6` - IS-IS IPv6 redistribution for routing protocols. The structure of `redistribute6` block is documented below.

The `redistribute6` block contains:

* `level` - Level. Valid values: `level-1-2` `level-1` `level-2` .
* `metric` - Metric.
* `metric_type` - Metric type. Valid values: `external` `internal` .
* `protocol` - Protocol name.
* `routemap` - Route map name. This attribute must reference one of the following datasources: `router.route-map.name` .
* `status` - Enable/disable redistribution. Valid values: `enable` `disable` .
* `summary_address` - IS-IS summary addresses. The structure of `summary_address` block is documented below.

The `summary_address` block contains:

* `id` - Summary address entry ID.
* `level` - Level. Valid values: `level-1-2` `level-1` `level-2` .
* `prefix` - Prefix.
* `summary_address6` - IS-IS IPv6 summary address. The structure of `summary_address6` block is documented below.

The `summary_address6` block contains:

* `id` - Prefix entry ID.
* `level` - Level. Valid values: `level-1-2` `level-1` `level-2` .
* `prefix6` - IPv6 prefix.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_isis can be imported using:
```sh
terraform import fortios_router_isis.labelname {{mkey}}
```
