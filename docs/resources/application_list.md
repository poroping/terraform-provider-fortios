---
subcategory: "FortiGate Application"
layout: "fortios"
page_title: "FortiOS: fortios_application_list"
description: |-
  Configure application control lists.
---

## fortios_application_list
Configure application control lists.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `app_replacemsg` - Enable/disable replacement messages for blocked applications. Valid values: `disable` `enable` .
* `comment` - comments
* `control_default_network_services` - Enable/disable enforcement of protocols over selected ports. Valid values: `disable` `enable` .
* `deep_app_inspection` - Enable/disable deep application inspection. Valid values: `disable` `enable` .
* `enforce_default_app_port` - Enable/disable default application port enforcement for allowed applications. Valid values: `disable` `enable` .
* `extended_log` - Enable/disable extended logging. Valid values: `enable` `disable` .
* `force_inclusion_ssl_di_sigs` - Enable/disable forced inclusion of SSL deep inspection signatures. Valid values: `disable` `enable` .
* `name` - List name.
* `options` - Basic application protocol signatures allowed by default. Valid values: `allow-dns` `allow-icmp` `allow-http` `allow-ssl` `allow-quic` .
* `other_application_action` - Action for other applications. Valid values: `pass` `block` .
* `other_application_log` - Enable/disable logging for other applications. Valid values: `disable` `enable` .
* `p2p_black_list` - P2P applications to be black listed. Valid values: `skype` `edonkey` `bittorrent` .
* `p2p_block_list` - P2P applications to be blocklisted. Valid values: `skype` `edonkey` `bittorrent` .
* `replacemsg_group` - Replacement message group. This attribute must reference one of the following datasources: `system.replacemsg-group.name` .
* `unknown_application_action` - Pass or block traffic from unknown applications. Valid values: `pass` `block` .
* `unknown_application_log` - Enable/disable logging for unknown applications. Valid values: `disable` `enable` .
* `default_network_services` - Default network service entries. The structure of `default_network_services` block is documented below.

The `default_network_services` block contains:

* `id` - Entry ID.
* `port` - Port number.
* `services` - Network protocols. Valid values: `http` `ssh` `telnet` `ftp` `dns` `smtp` `pop3` `imap` `snmp` `nntp` `https` .
* `violation_action` - Action for protocols not in the allowlist for selected port. Valid values: `pass` `monitor` `block` .
* `entries` - Application list entries. The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Pass or block traffic, or reset connection for traffic from this application. Valid values: `pass` `block` `reset` .
* `behavior` - Application behavior filter.
* `id` - Entry ID.
* `log` - Enable/disable logging for this application list. Valid values: `disable` `enable` .
* `log_packet` - Enable/disable packet logging. Valid values: `disable` `enable` .
* `per_ip_shaper` - Per-IP traffic shaper. This attribute must reference one of the following datasources: `firewall.shaper.per-ip-shaper.name` .
* `popularity` - Application popularity filter (1 - 5, from least to most popular). Valid values: `1` `2` `3` `4` `5` .
* `protocols` - Application protocol filter.
* `quarantine` - Quarantine method. Valid values: `none` `attacker` .
* `quarantine_expiry` - Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.
* `quarantine_log` - Enable/disable quarantine logging. Valid values: `disable` `enable` .
* `rate_count` - Count of the rate.
* `rate_duration` - Duration (sec) of the rate.
* `rate_mode` - Rate limit mode. Valid values: `periodical` `continuous` .
* `rate_track` - Track the packet protocol field. Valid values: `none` `src-ip` `dest-ip` `dhcp-client-mac` `dns-domain` .
* `session_ttl` - Session TTL (0 = default).
* `shaper` - Traffic shaper. This attribute must reference one of the following datasources: `firewall.shaper.traffic-shaper.name` .
* `shaper_reverse` - Reverse traffic shaper. This attribute must reference one of the following datasources: `firewall.shaper.traffic-shaper.name` .
* `technology` - Application technology filter.
* `vendor` - Application vendor filter.
* `application` - ID of allowed applications. The structure of `application` block is documented below.

The `application` block contains:

* `id` - Application IDs.
* `category` - Category ID list. The structure of `category` block is documented below.

The `category` block contains:

* `id` - Application category ID.
* `exclusion` - ID of excluded applications. The structure of `exclusion` block is documented below.

The `exclusion` block contains:

* `id` - Excluded application IDs.
* `parameters` - Application parameters. The structure of `parameters` block is documented below.

The `parameters` block contains:

* `id` - Parameter tuple ID.
* `value` - Parameter value.
* `members` - Parameter tuple members. The structure of `members` block is documented below.

The `members` block contains:

* `id` - Parameter.
* `name` - Parameter name.
* `value` - Parameter value.
* `risk` - Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.

The `risk` block contains:

* `level` - Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical).
* `sub_category` - Application Sub-category ID list. The structure of `sub_category` block is documented below.

The `sub_category` block contains:

* `id` - Application sub-category ID.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_application_list can be imported using:
```sh
terraform import fortios_application_list.labelname {{mkey}}
```
