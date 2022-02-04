---
subcategory: "FortiGate Emailfilter"
layout: "fortios"
page_title: "FortiOS: fortios_emailfilter_profile"
description: |-
  Configure Email Filter profiles.
---

## fortios_emailfilter_profile
Configure Email Filter profiles.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `comment` - Comment.
* `external` - Enable/disable external Email inspection. Valid values: `enable` `disable` .
* `feature_set` - Flow/proxy feature set. Valid values: `flow` `proxy` .
* `name` - Profile name.
* `options` - Options. Valid values: `bannedword` `spambal` `spamfsip` `spamfssubmit` `spamfschksum` `spamfsurl` `spamhelodns` `spamraddrdns` `spamrbl` `spamhdrcheck` `spamfsphish` .
* `replacemsg_group` - Replacement message group. This attribute must reference one of the following datasources: `system.replacemsg-group.name` .
* `spam_bal_table` - Anti-spam block/allow list table ID. This attribute must reference one of the following datasources: `emailfilter.block-allow-list.id` .
* `spam_bwl_table` - Anti-spam black/white list table ID. This attribute must reference one of the following datasources: `emailfilter.bwl.id` .
* `spam_bword_table` - Anti-spam banned word table ID. This attribute must reference one of the following datasources: `emailfilter.bword.id` .
* `spam_bword_threshold` - Spam banned word threshold.
* `spam_filtering` - Enable/disable spam filtering. Valid values: `enable` `disable` .
* `spam_iptrust_table` - Anti-spam IP trust table ID. This attribute must reference one of the following datasources: `emailfilter.iptrust.id` .
* `spam_log` - Enable/disable spam logging for email filtering. Valid values: `disable` `enable` .
* `spam_log_fortiguard_response` - Enable/disable logging FortiGuard spam response. Valid values: `disable` `enable` .
* `spam_mheader_table` - Anti-spam MIME header table ID. This attribute must reference one of the following datasources: `emailfilter.mheader.id` .
* `spam_rbl_table` - Anti-spam DNSBL table ID. This attribute must reference one of the following datasources: `emailfilter.dnsbl.id` .
* `file_filter` - File filter. The structure of `file_filter` block is documented below.

The `file_filter` block contains:

* `log` - Enable/disable file filter logging. Valid values: `enable` `disable` .
* `scan_archive_contents` - Enable/disable file filter archive contents scan. Valid values: `enable` `disable` .
* `status` - Enable/disable file filter. Valid values: `enable` `disable` .
* `entries` - File filter entries. The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Action taken for matched file. Valid values: `log` `block` .
* `comment` - Comment.
* `filter` - Add a file filter.
* `password_protected` - Match password-protected files. Valid values: `yes` `any` .
* `protocol` - Protocols to apply with. Valid values: `smtp` `imap` `pop3` .
* `file_type` - Select file type. The structure of `file_type` block is documented below.

The `file_type` block contains:

* `name` - File type name. This attribute must reference one of the following datasources: `antivirus.filetype.name` .
* `gmail` - Gmail. The structure of `gmail` block is documented below.

The `gmail` block contains:

* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable` `enable` .
* `imap` - IMAP. The structure of `imap` block is documented below.

The `imap` block contains:

* `action` - Action for spam email. Valid values: `pass` `tag` .
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable` `enable` .
* `tag_msg` - Subject text or header added to spam email.
* `tag_type` - Tag subject or header for spam email. Valid values: `subject` `header` `spaminfo` .
* `mapi` - MAPI. The structure of `mapi` block is documented below.

The `mapi` block contains:

* `action` - Action for spam email. Valid values: `pass` `discard` .
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable` `enable` .
* `msn_hotmail` - MSN Hotmail. The structure of `msn_hotmail` block is documented below.

The `msn_hotmail` block contains:

* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable` `enable` .
* `other_webmails` - Other supported webmails. The structure of `other_webmails` block is documented below.

The `other_webmails` block contains:

* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable` `enable` .
* `pop3` - POP3. The structure of `pop3` block is documented below.

The `pop3` block contains:

* `action` - Action for spam email. Valid values: `pass` `tag` .
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable` `enable` .
* `tag_msg` - Subject text or header added to spam email.
* `tag_type` - Tag subject or header for spam email. Valid values: `subject` `header` `spaminfo` .
* `smtp` - SMTP. The structure of `smtp` block is documented below.

The `smtp` block contains:

* `action` - Action for spam email. Valid values: `pass` `tag` `discard` .
* `hdrip` - Enable/disable SMTP email header IP checks for spamfsip, spamrbl, and spambal filters. Valid values: `disable` `enable` .
* `local_override` - Enable/disable local filter to override SMTP remote check result. Valid values: `disable` `enable` .
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable` `enable` .
* `tag_msg` - Subject text or header added to spam email.
* `tag_type` - Tag subject or header for spam email. Valid values: `subject` `header` `spaminfo` .
* `yahoo_mail` - Yahoo! Mail. The structure of `yahoo_mail` block is documented below.

The `yahoo_mail` block contains:

* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable` `enable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_emailfilter_profile can be imported using:
```sh
terraform import fortios_emailfilter_profile.labelname {{mkey}}
```
