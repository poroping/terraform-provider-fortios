---
subcategory: "FortiGate Webfilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_profile"
description: |-
  Get information on a fortios Configure Web filter profiles.
---

# Data Source: fortios_webfilter_profile
Use this data source to get information on a fortios Configure Web filter profiles.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Optional comments.
* `extended_log` - Enable/disable extended logging for web filtering.
* `feature_set` - Flow/proxy feature set.
* `https_replacemsg` - Enable replacement messages for HTTPS.
* `log_all_url` - Enable/disable logging all URLs visited.
* `name` - Profile name.
* `options` - Options.
* `ovrd_perm` - Permitted override types.
* `post_action` - Action taken for HTTP POST traffic.
* `replacemsg_group` - Replacement message group.
* `web_antiphishing_log` - Enable/disable logging of AntiPhishing checks.
* `web_content_log` - Enable/disable logging logging blocked web content.
* `web_extended_all_action_log` - Enable/disable extended any filter action logging for web filtering.
* `web_filter_activex_log` - Enable/disable logging ActiveX.
* `web_filter_applet_log` - Enable/disable logging Java applets.
* `web_filter_command_block_log` - Enable/disable logging blocked commands.
* `web_filter_cookie_log` - Enable/disable logging cookie filtering.
* `web_filter_cookie_removal_log` - Enable/disable logging blocked cookies.
* `web_filter_js_log` - Enable/disable logging Java scripts.
* `web_filter_jscript_log` - Enable/disable logging JScripts.
* `web_filter_referer_log` - Enable/disable logging referrers.
* `web_filter_unknown_log` - Enable/disable logging unknown scripts.
* `web_filter_vbs_log` - Enable/disable logging VBS scripts.
* `web_ftgd_err_log` - Enable/disable logging rating errors.
* `web_ftgd_quota_usage` - Enable/disable logging daily quota usage.
* `web_invalid_domain_log` - Enable/disable logging invalid domain names.
* `web_url_log` - Enable/disable logging URL filtering.
* `wisp` - Enable/disable web proxy WISP.
* `wisp_algorithm` - WISP server selection algorithm.
* `youtube_channel_status` - YouTube channel filter status.
* `antiphish` - AntiPhishing profile.The structure of `antiphish` block is documented below.

The `antiphish` block contains:

* `authentication` - Authentication methods.
* `check_basic_auth` - Enable/disable checking of HTTP Basic Auth field for known credentials.
* `check_uri` - Enable/disable checking of GET URI parameters for known credentials.
* `check_username_only` - Enable/disable username only matching of credentials. Action will be taken for valid usernames regardless of password validity.
* `default_action` - Action to be taken when there is no matching rule.
* `domain_controller` - Domain for which to verify received credentials against.
* `ldap` - LDAP server for which to verify received credentials against.
* `max_body_len` - Maximum size of a POST body to check for credentials.
* `status` - Toggle AntiPhishing functionality.
* `custom_patterns` - Custom username and password regex patterns.The structure of `custom_patterns` block is documented below.

The `custom_patterns` block contains:

* `category` - Category that the pattern matches.
* `pattern` - Target pattern.
* `type` - Pattern will be treated either as a regex pattern or literal string.
* `inspection_entries` - AntiPhishing entries.The structure of `inspection_entries` block is documented below.

The `inspection_entries` block contains:

* `action` - Action to be taken upon an AntiPhishing match.
* `fortiguard_category` - FortiGuard category to match.
* `name` - Inspection target name.
* `file_filter` - File filter.The structure of `file_filter` block is documented below.

The `file_filter` block contains:

* `log` - Enable/disable file filter logging.
* `scan_archive_contents` - Enable/disable file filter archive contents scan.
* `status` - Enable/disable file filter.
* `entries` - File filter entries.The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Action taken for matched file.
* `comment` - Comment.
* `direction` - Match files transmitted in the session's originating or reply direction.
* `filter` - Add a file filter.
* `password_protected` - Match password-protected files.
* `protocol` - Protocols to apply with.
* `file_type` - Select file type.The structure of `file_type` block is documented below.

The `file_type` block contains:

* `name` - File type name.
* `ftgd_wf` - FortiGuard Web Filter settings.The structure of `ftgd_wf` block is documented below.

The `ftgd_wf` block contains:

* `exempt_quota` - Do not stop quota for these categories.
* `max_quota_timeout` - Maximum FortiGuard quota used by single page view in seconds (excludes streams).
* `options` - Options for FortiGuard Web Filter.
* `ovrd` - Allow web filter profile overrides.
* `rate_crl_urls` - Enable/disable rating CRL by URL.
* `rate_css_urls` - Enable/disable rating CSS by URL.
* `rate_image_urls` - Enable/disable rating images by URL.
* `rate_javascript_urls` - Enable/disable rating JavaScript by URL.
* `filters` - FortiGuard filters.The structure of `filters` block is documented below.

The `filters` block contains:

* `action` - Action to take for matches.
* `category` - Categories and groups the filter examines.
* `id` - ID number.
* `log` - Enable/disable logging.
* `override_replacemsg` - Override replacement message.
* `warn_duration` - Duration of warnings.
* `warning_duration_type` - Re-display warning after closing browser or after a timeout.
* `warning_prompt` - Warning prompts in each category or each domain.
* `auth_usr_grp` - Groups with permission to authenticate.The structure of `auth_usr_grp` block is documented below.

The `auth_usr_grp` block contains:

* `name` - User group name.
* `quota` - FortiGuard traffic quota settings.The structure of `quota` block is documented below.

The `quota` block contains:

* `category` - FortiGuard categories to apply quota to (category action must be set to monitor).
* `duration` - Duration of quota.
* `id` - ID number.
* `override_replacemsg` - Override replacement message.
* `type` - Quota type.
* `unit` - Traffic quota unit of measurement.
* `value` - Traffic quota value.
* `override` - Web Filter override settings.The structure of `override` block is documented below.

The `override` block contains:

* `ovrd_cookie` - Allow/deny browser-based (cookie) overrides.
* `ovrd_dur` - Override duration.
* `ovrd_dur_mode` - Override duration mode.
* `ovrd_scope` - Override scope.
* `profile_attribute` - Profile attribute to retrieve from the RADIUS server.
* `profile_type` - Override profile type.
* `ovrd_user_group` - User groups with permission to use the override.The structure of `ovrd_user_group` block is documented below.

The `ovrd_user_group` block contains:

* `name` - User group name.
* `profile` - Web filter profile with permission to create overrides.The structure of `profile` block is documented below.

The `profile` block contains:

* `name` - Web profile.
* `web` - Web content filtering settings.The structure of `web` block is documented below.

The `web` block contains:

* `allowlist` - FortiGuard allowlist settings.
* `blacklist` - Enable/disable automatic addition of URLs detected by FortiSandbox to blacklist.
* `blocklist` - Enable/disable automatic addition of URLs detected by FortiSandbox to blocklist.
* `bword_table` - Banned word table ID.
* `bword_threshold` - Banned word score threshold.
* `content_header_list` - Content header list.
* `log_search` - Enable/disable logging all search phrases.
* `safe_search` - Safe search type.
* `urlfilter_table` - URL filter table ID.
* `vimeo_restrict` - Set Vimeo-restrict ("7" = don't show mature content, "134" = don't show unrated and mature content). A value of cookie "content_rating".
* `whitelist` - FortiGuard whitelist settings.
* `youtube_restrict` - YouTube EDU filter level.
* `keyword_match` - Search keywords to log when match is found.The structure of `keyword_match` block is documented below.

The `keyword_match` block contains:

* `pattern` - Pattern/keyword to search for.
* `wisp_servers` - WISP servers.The structure of `wisp_servers` block is documented below.

The `wisp_servers` block contains:

* `name` - Server name.
* `youtube_channel_filter` - YouTube channel filter.The structure of `youtube_channel_filter` block is documented below.

The `youtube_channel_filter` block contains:

* `channel_id` - YouTube channel ID to be filtered.
* `comment` - Comment.
* `id` - ID.
