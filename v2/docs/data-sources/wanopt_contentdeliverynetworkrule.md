---
subcategory: "FortiGate Wanopt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_contentdeliverynetworkrule"
description: |-
  Get information on a fortios Configure WAN optimization content delivery network rules.
---

# Data Source: fortios_wanopt_contentdeliverynetworkrule
Use this data source to get information on a fortios Configure WAN optimization content delivery network rules.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name of table.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `category` - Content delivery network rule category.
* `comment` - Comment about this CDN-rule.
* `name` - Name of table.
* `request_cache_control` - Enable/disable HTTP request cache control.
* `response_cache_control` - Enable/disable HTTP response cache control.
* `response_expires` - Enable/disable HTTP response cache expires.
* `status` - Enable/disable WAN optimization content delivery network rules.
* `updateserver` - Enable/disable update server.
* `host_domain_name_suffix` - Suffix portion of the fully qualified domain name (eg. fortinet.com in "www.fortinet.com").The structure of `host_domain_name_suffix` block is documented below.

The `host_domain_name_suffix` block contains:

* `name` - Suffix portion of the fully qualified domain name.
* `rules` - WAN optimization content delivery network rule entries.The structure of `rules` block is documented below.

The `rules` block contains:

* `match_mode` - Match criteria for collecting content ID.
* `name` - WAN optimization content delivery network rule name.
* `skip_rule_mode` - Skip mode when evaluating skip-rules.
* `content_id` - Content ID settings.The structure of `content_id` block is documented below.

The `content_id` block contains:

* `end_direction` - Search direction from end-str match.
* `end_skip` - Number of characters in URL to skip after end-str has been matched.
* `end_str` - String from which to end search.
* `range_str` - Name of content ID within the start string and end string.
* `start_direction` - Search direction from start-str match.
* `start_skip` - Number of characters in URL to skip after start-str has been matched.
* `start_str` - String from which to start search.
* `target` - Option in HTTP header or URL parameter to match.
* `match_entries` - List of entries to match.The structure of `match_entries` block is documented below.

The `match_entries` block contains:

* `id` - Rule ID.
* `target` - Option in HTTP header or URL parameter to match.
* `pattern` - Pattern string for matching target (Referrer or URL pattern, eg. "a", "a*c", "*a*", "a*c*e", and "*").The structure of `pattern` block is documented below.

The `pattern` block contains:

* `string` - Pattern strings.
* `skip_entries` - List of entries to skip.The structure of `skip_entries` block is documented below.

The `skip_entries` block contains:

* `id` - Rule ID.
* `target` - Option in HTTP header or URL parameter to match.
* `pattern` - Pattern string for matching target (Referrer or URL pattern, eg. "a", "a*c", "*a*", "a*c*e", and "*").The structure of `pattern` block is documented below.

The `pattern` block contains:

* `string` - Pattern strings.
