---
subcategory: "FortiGate Application"
layout: "fortios"
page_title: "FortiOS: fortios_application_list"
description: |-
  Get information on a fortios Configure application control lists.
---

# Data Source: fortios_application_list
Use this data source to get information on a fortios Configure application control lists.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) List name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `app_replacemsg` - Enable/disable replacement messages for blocked applications.
* `comment` - Comments.
* `control_default_network_services` - Enable/disable enforcement of protocols over selected ports.
* `deep_app_inspection` - Enable/disable deep application inspection.
* `enforce_default_app_port` - Enable/disable default application port enforcement for allowed applications.
* `extended_log` - Enable/disable extended logging.
* `force_inclusion_ssl_di_sigs` - Enable/disable forced inclusion of SSL deep inspection signatures.
* `name` - List name.
* `options` - Basic application protocol signatures allowed by default.
* `other_application_action` - Action for other applications.
* `other_application_log` - Enable/disable logging for other applications.
* `p2p_black_list` - P2P applications to be black listed.
* `p2p_block_list` - P2P applications to be block listed.
* `replacemsg_group` - Replacement message group.
* `unknown_application_action` - Pass or block traffic from unknown applications.
* `unknown_application_log` - Enable/disable logging for unknown applications.
* `default_network_services` - Default network service entries.The structure of `default_network_services` block is documented below.

The `default_network_services` block contains:

* `id` - Entry ID.
* `port` - Port number.
* `services` - Network protocols.
* `violation_action` - Action for protocols not in the allowlist for selected port.
* `entries` - Application list entries.The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Pass or block traffic, or reset connection for traffic from this application.
* `behavior` - Application behavior filter.
* `id` - Entry ID.
* `log` - Enable/disable logging for this application list.
* `log_packet` - Enable/disable packet logging.
* `per_ip_shaper` - Per-IP traffic shaper.
* `popularity` - Application popularity filter (1 - 5, from least to most popular).
* `protocols` - Application protocol filter.
* `quarantine` - Quarantine method.
* `quarantine_expiry` - Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.
* `quarantine_log` - Enable/disable quarantine logging.
* `rate_count` - Count of the rate.
* `rate_duration` - Duration (sec) of the rate.
* `rate_mode` - Rate limit mode.
* `rate_track` - Track the packet protocol field.
* `session_ttl` - Session TTL (0 = default).
* `shaper` - Traffic shaper.
* `shaper_reverse` - Reverse traffic shaper.
* `technology` - Application technology filter.
* `vendor` - Application vendor filter.
* `application` - ID of allowed applications.The structure of `application` block is documented below.

The `application` block contains:

* `id` - Application IDs.
* `category` - Category ID list.The structure of `category` block is documented below.

The `category` block contains:

* `id` - Application category ID.
* `exclusion` - ID of excluded applications.The structure of `exclusion` block is documented below.

The `exclusion` block contains:

* `id` - Excluded application IDs.
* `parameters` - Application parameters.The structure of `parameters` block is documented below.

The `parameters` block contains:

* `id` - Parameter tuple ID.
* `value` - Parameter value.
* `members` - Parameter tuple members.The structure of `members` block is documented below.

The `members` block contains:

* `id` - Parameter.
* `name` - Parameter name.
* `value` - Parameter value.
* `risk` - Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical).The structure of `risk` block is documented below.

The `risk` block contains:

* `level` - Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical).
* `sub_category` - Application Sub-category ID list.The structure of `sub_category` block is documented below.

The `sub_category` block contains:

* `id` - Application sub-category ID.
