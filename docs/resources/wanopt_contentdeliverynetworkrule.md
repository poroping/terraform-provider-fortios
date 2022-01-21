---
subcategory: "FortiGate Wanopt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_contentdeliverynetworkrule"
description: |-
  Configure WAN optimization content delivery network rules.
---

## fortios_wanopt_contentdeliverynetworkrule
Configure WAN optimization content delivery network rules.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `category` - Content delivery network rule category. Valid values: `vcache` `youtube` .
* `comment` - Comment about this CDN-rule.
* `name` - Name of table.
* `request_cache_control` - Enable/disable HTTP request cache control. Valid values: `enable` `disable` .
* `response_cache_control` - Enable/disable HTTP response cache control. Valid values: `enable` `disable` .
* `response_expires` - Enable/disable HTTP response cache expires. Valid values: `enable` `disable` .
* `status` - Enable/disable WAN optimization content delivery network rules. Valid values: `enable` `disable` .
* `updateserver` - Enable/disable update server. Valid values: `enable` `disable` .
* `host_domain_name_suffix` - Suffix portion of the fully qualified domain name (eg. fortinet.com in "www.fortinet.com"). The structure of `host_domain_name_suffix` block is documented below.

The `host_domain_name_suffix` block contains:

* `name` - Suffix portion of the fully qualified domain name.
* `rules` - WAN optimization content delivery network rule entries. The structure of `rules` block is documented below.

The `rules` block contains:

* `match_mode` - Match criteria for collecting content ID. Valid values: `all` `any` .
* `name` - WAN optimization content delivery network rule name.
* `skip_rule_mode` - Skip mode when evaluating skip-rules. Valid values: `all` `any` .
* `content_id` - Content ID settings. The structure of `content_id` block is documented below.

The `content_id` block contains:

* `end_direction` - Search direction from end-str match. Valid values: `forward` `backward` .
* `end_skip` - Number of characters in URL to skip after end-str has been matched.
* `end_str` - String from which to end search.
* `range_str` - Name of content ID within the start string and end string.
* `start_direction` - Search direction from start-str match. Valid values: `forward` `backward` .
* `start_skip` - Number of characters in URL to skip after start-str has been matched.
* `start_str` - String from which to start search.
* `target` - Option in HTTP header or URL parameter to match. Valid values: `path` `parameter` `referrer` `youtube-map` `youtube-id` `youku-id` `hls-manifest` `dash-manifest` `hls-fragment` `dash-fragment` .
* `match_entries` - List of entries to match. The structure of `match_entries` block is documented below.

The `match_entries` block contains:

* `id` - Rule ID.
* `target` - Option in HTTP header or URL parameter to match. Valid values: `path` `parameter` `referrer` `youtube-map` `youtube-id` `youku-id` .
* `pattern` - Pattern string for matching target (Referrer or URL pattern, eg. "a", "a*c", "*a*", "a*c*e", and "*"). The structure of `pattern` block is documented below.

The `pattern` block contains:

* `string` - Pattern strings.
* `skip_entries` - List of entries to skip. The structure of `skip_entries` block is documented below.

The `skip_entries` block contains:

* `id` - Rule ID.
* `target` - Option in HTTP header or URL parameter to match. Valid values: `path` `parameter` `referrer` `youtube-map` `youtube-id` `youku-id` .
* `pattern` - Pattern string for matching target (Referrer or URL pattern, eg. "a", "a*c", "*a*", "a*c*e", and "*"). The structure of `pattern` block is documented below.

The `pattern` block contains:

* `string` - Pattern strings.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wanopt_contentdeliverynetworkrule can be imported using:
```sh
terraform import fortios_wanopt_contentdeliverynetworkrule.labelname {{mkey}}
```
