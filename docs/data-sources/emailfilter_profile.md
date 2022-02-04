---
subcategory: "FortiGate Emailfilter"
layout: "fortios"
page_title: "FortiOS: fortios_emailfilter_profile"
description: |-
  Get information on a fortios Configure Email Filter profiles.
---

# Data Source: fortios_emailfilter_profile
Use this data source to get information on a fortios Configure Email Filter profiles.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `external` - Enable/disable external Email inspection.
* `feature_set` - Flow/proxy feature set.
* `name` - Profile name.
* `options` - Options.
* `replacemsg_group` - Replacement message group.
* `spam_bal_table` - Anti-spam block/allow list table ID.
* `spam_bwl_table` - Anti-spam black/white list table ID.
* `spam_bword_table` - Anti-spam banned word table ID.
* `spam_bword_threshold` - Spam banned word threshold.
* `spam_filtering` - Enable/disable spam filtering.
* `spam_iptrust_table` - Anti-spam IP trust table ID.
* `spam_log` - Enable/disable spam logging for email filtering.
* `spam_log_fortiguard_response` - Enable/disable logging FortiGuard spam response.
* `spam_mheader_table` - Anti-spam MIME header table ID.
* `spam_rbl_table` - Anti-spam DNSBL table ID.
* `file_filter` - File filter.The structure of `file_filter` block is documented below.

The `file_filter` block contains:

* `log` - Enable/disable file filter logging.
* `scan_archive_contents` - Enable/disable file filter archive contents scan.
* `status` - Enable/disable file filter.
* `entries` - File filter entries.The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Action taken for matched file.
* `comment` - Comment.
* `filter` - Add a file filter.
* `password_protected` - Match password-protected files.
* `protocol` - Protocols to apply with.
* `file_type` - Select file type.The structure of `file_type` block is documented below.

The `file_type` block contains:

* `name` - File type name.
* `gmail` - Gmail.The structure of `gmail` block is documented below.

The `gmail` block contains:

* `log` - Enable/disable logging.
* `log_all` - Enable/disable logging of all email traffic.
* `imap` - IMAP.The structure of `imap` block is documented below.

The `imap` block contains:

* `action` - Action for spam email.
* `log` - Enable/disable logging.
* `log_all` - Enable/disable logging of all email traffic.
* `tag_msg` - Subject text or header added to spam email.
* `tag_type` - Tag subject or header for spam email.
* `mapi` - MAPI.The structure of `mapi` block is documented below.

The `mapi` block contains:

* `action` - Action for spam email.
* `log` - Enable/disable logging.
* `log_all` - Enable/disable logging of all email traffic.
* `msn_hotmail` - MSN Hotmail.The structure of `msn_hotmail` block is documented below.

The `msn_hotmail` block contains:

* `log` - Enable/disable logging.
* `log_all` - Enable/disable logging of all email traffic.
* `other_webmails` - Other supported webmails.The structure of `other_webmails` block is documented below.

The `other_webmails` block contains:

* `log_all` - Enable/disable logging of all email traffic.
* `pop3` - POP3.The structure of `pop3` block is documented below.

The `pop3` block contains:

* `action` - Action for spam email.
* `log` - Enable/disable logging.
* `log_all` - Enable/disable logging of all email traffic.
* `tag_msg` - Subject text or header added to spam email.
* `tag_type` - Tag subject or header for spam email.
* `smtp` - SMTP.The structure of `smtp` block is documented below.

The `smtp` block contains:

* `action` - Action for spam email.
* `hdrip` - Enable/disable SMTP email header IP checks for spamfsip, spamrbl, and spambal filters.
* `local_override` - Enable/disable local filter to override SMTP remote check result.
* `log` - Enable/disable logging.
* `log_all` - Enable/disable logging of all email traffic.
* `tag_msg` - Subject text or header added to spam email.
* `tag_type` - Tag subject or header for spam email.
* `yahoo_mail` - Yahoo! Mail.The structure of `yahoo_mail` block is documented below.

The `yahoo_mail` block contains:

* `log` - Enable/disable logging.
* `log_all` - Enable/disable logging of all email traffic.
