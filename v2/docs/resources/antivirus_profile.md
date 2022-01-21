---
subcategory: "FortiGate Antivirus"
layout: "fortios"
page_title: "FortiOS: fortios_antivirus_profile"
description: |-
  Configure AntiVirus profiles.
---

## fortios_antivirus_profile
Configure AntiVirus profiles.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `analytics_accept_filetype` - Only submit files matching this DLP file-pattern to FortiSandbox. This attribute must reference one of the following datasources: `dlp.filepattern.id` .
* `analytics_bl_filetype` - Only submit files matching this DLP file-pattern to FortiSandbox. This attribute must reference one of the following datasources: `dlp.filepattern.id` .
* `analytics_db` - Enable/disable using the FortiSandbox signature database to supplement the AV signature databases. Valid values: `disable` `enable` .
* `analytics_ignore_filetype` - Do not submit files matching this DLP file-pattern to FortiSandbox. This attribute must reference one of the following datasources: `dlp.filepattern.id` .
* `analytics_max_upload` - Maximum size of files that can be uploaded to FortiSandbox.
* `analytics_wl_filetype` - Do not submit files matching this DLP file-pattern to FortiSandbox. This attribute must reference one of the following datasources: `dlp.filepattern.id` .
* `av_block_log` - Enable/disable logging for AntiVirus file blocking. Valid values: `enable` `disable` .
* `av_virus_log` - Enable/disable AntiVirus logging. Valid values: `enable` `disable` .
* `comment` - Comment.
* `ems_threat_feed` - Enable/disable use of EMS threat feed when performing AntiVirus scan. Analyzes files including the content of archives. Valid values: `disable` `enable` .
* `extended_log` - Enable/disable extended logging for antivirus. Valid values: `enable` `disable` .
* `external_blocklist_archive_scan` - Enable/disable external-blocklist archive scanning. Valid values: `disable` `enable` .
* `external_blocklist_enable_all` - Enable/disable all external blocklists. Valid values: `disable` `enable` .
* `feature_set` - Flow/proxy feature set. Valid values: `flow` `proxy` .
* `fortiai_error_action` - Action to take if FortiAI encounters an error. Valid values: `log-only` `block` `ignore` .
* `fortiai_timeout_action` - Action to take if FortiAI encounters a scan timeout. Valid values: `log-only` `block` `ignore` .
* `ftgd_analytics` - Settings to control which files are uploaded to FortiSandbox. Valid values: `disable` `suspicious` `everything` .
* `mobile_malware_db` - Enable/disable using the mobile malware signature database. Valid values: `disable` `enable` .
* `name` - Profile name.
* `outbreak_prevention_archive_scan` - Enable/disable outbreak-prevention archive scanning. Valid values: `disable` `enable` .
* `replacemsg_group` - Replacement message group customized for this profile. This attribute must reference one of the following datasources: `system.replacemsg-group.name` .
* `scan_mode` - Choose between default scan mode and legacy scan mode.  Valid values: `default` `legacy` .
* `cifs` - Configure CIFS AntiVirus options. The structure of `cifs` block is documented below.

The `cifs` block contains:

* `archive_block` - Select the archive types to block. Valid values: `encrypted` `corrupted` `partiallycorrupted` `multipart` `nested` `mailbomb` `timeout` `unhandled` .
* `archive_log` - Select the archive types to log. Valid values: `encrypted` `corrupted` `partiallycorrupted` `multipart` `nested` `mailbomb` `timeout` `unhandled` .
* `av_scan` - Enable AntiVirus scan service. Valid values: `disable` `block` `monitor` .
* `emulator` - Enable/disable the virus emulator. Valid values: `enable` `disable` .
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable` `block` `monitor` .
* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable` `block` `monitor` .
* `options` - Enable/disable CIFS AntiVirus scanning, monitoring, and quarantine. Valid values: `scan` `avmonitor` `quarantine` .
* `outbreak_prevention` - Enable virus outbreak prevention service. Valid values: `disable` `block` `monitor` .
* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable` `enable` .
* `content_disarm` - AV Content Disarm and Reconstruction settings. The structure of `content_disarm` block is documented below.

The `content_disarm` block contains:

* `cover_page` - Enable/disable inserting a cover page into the disarmed document. Valid values: `disable` `enable` .
* `detect_only` - Enable/disable only detect disarmable files, do not alter content. Valid values: `disable` `enable` .
* `error_action` - Action to be taken if CDR engine encounters an unrecoverable error. Valid values: `block` `log-only` `ignore` .
* `office_action` - Enable/disable stripping of PowerPoint action events in Microsoft Office documents. Valid values: `disable` `enable` .
* `office_dde` - Enable/disable stripping of Dynamic Data Exchange events in Microsoft Office documents. Valid values: `disable` `enable` .
* `office_embed` - Enable/disable stripping of embedded objects in Microsoft Office documents. Valid values: `disable` `enable` .
* `office_hylink` - Enable/disable stripping of hyperlinks in Microsoft Office documents. Valid values: `disable` `enable` .
* `office_linked` - Enable/disable stripping of linked objects in Microsoft Office documents. Valid values: `disable` `enable` .
* `office_macro` - Enable/disable stripping of macros in Microsoft Office documents. Valid values: `disable` `enable` .
* `original_file_destination` - Destination to send original file if active content is removed. Valid values: `fortisandbox` `quarantine` `discard` .
* `pdf_act_form` - Enable/disable stripping of PDF document actions that submit data to other targets. Valid values: `disable` `enable` .
* `pdf_act_gotor` - Enable/disable stripping of PDF document actions that access other PDF documents. Valid values: `disable` `enable` .
* `pdf_act_java` - Enable/disable stripping of PDF document actions that execute JavaScript code. Valid values: `disable` `enable` .
* `pdf_act_launch` - Enable/disable stripping of PDF document actions that launch other applications. Valid values: `disable` `enable` .
* `pdf_act_movie` - Enable/disable stripping of PDF document actions that play a movie. Valid values: `disable` `enable` .
* `pdf_act_sound` - Enable/disable stripping of PDF document actions that play a sound. Valid values: `disable` `enable` .
* `pdf_embedfile` - Enable/disable stripping of embedded files in PDF documents. Valid values: `disable` `enable` .
* `pdf_hyperlink` - Enable/disable stripping of hyperlinks from PDF documents. Valid values: `disable` `enable` .
* `pdf_javacode` - Enable/disable stripping of JavaScript code in PDF documents. Valid values: `disable` `enable` .
* `external_blocklist` - One or more external malware block lists. The structure of `external_blocklist` block is documented below.

The `external_blocklist` block contains:

* `name` - External blocklist. This attribute must reference one of the following datasources: `system.external-resource.name` .
* `ftp` - Configure FTP AntiVirus options. The structure of `ftp` block is documented below.

The `ftp` block contains:

* `archive_block` - Select the archive types to block. Valid values: `encrypted` `corrupted` `partiallycorrupted` `multipart` `nested` `mailbomb` `timeout` `unhandled` .
* `archive_log` - Select the archive types to log. Valid values: `encrypted` `corrupted` `partiallycorrupted` `multipart` `nested` `mailbomb` `timeout` `unhandled` .
* `av_scan` - Enable AntiVirus scan service. Valid values: `disable` `block` `monitor` .
* `emulator` - Enable/disable the virus emulator. Valid values: `enable` `disable` .
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable` `block` `monitor` .
* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable` `block` `monitor` .
* `options` - Enable/disable FTP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan` `avmonitor` `quarantine` .
* `outbreak_prevention` - Enable virus outbreak prevention service. Valid values: `disable` `block` `monitor` .
* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable` `enable` .
* `http` - Configure HTTP AntiVirus options. The structure of `http` block is documented below.

The `http` block contains:

* `archive_block` - Select the archive types to block. Valid values: `encrypted` `corrupted` `partiallycorrupted` `multipart` `nested` `mailbomb` `timeout` `unhandled` .
* `archive_log` - Select the archive types to log. Valid values: `encrypted` `corrupted` `partiallycorrupted` `multipart` `nested` `mailbomb` `timeout` `unhandled` .
* `av_scan` - Enable AntiVirus scan service. Valid values: `disable` `block` `monitor` .
* `content_disarm` - Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan. Valid values: `disable` `enable` .
* `emulator` - Enable/disable the virus emulator. Valid values: `enable` `disable` .
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable` `block` `monitor` .
* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable` `block` `monitor` .
* `options` - Enable/disable HTTP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan` `avmonitor` `quarantine` .
* `outbreak_prevention` - Enable virus outbreak prevention service. Valid values: `disable` `block` `monitor` .
* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable` `enable` .
* `imap` - Configure IMAP AntiVirus options. The structure of `imap` block is documented below.

The `imap` block contains:

* `archive_block` - Select the archive types to block. Valid values: `encrypted` `corrupted` `partiallycorrupted` `multipart` `nested` `mailbomb` `timeout` `unhandled` .
* `archive_log` - Select the archive types to log. Valid values: `encrypted` `corrupted` `partiallycorrupted` `multipart` `nested` `mailbomb` `timeout` `unhandled` .
* `av_scan` - Enable AntiVirus scan service. Valid values: `disable` `block` `monitor` .
* `content_disarm` - Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan. Valid values: `disable` `enable` .
* `emulator` - Enable/disable the virus emulator. Valid values: `enable` `disable` .
* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default` `virus` .
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable` `block` `monitor` .
* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable` `block` `monitor` .
* `options` - Enable/disable IMAP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan` `avmonitor` `quarantine` .
* `outbreak_prevention` - Enable virus outbreak prevention service. Valid values: `disable` `block` `monitor` .
* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable` `enable` .
* `mapi` - Configure MAPI AntiVirus options. The structure of `mapi` block is documented below.

The `mapi` block contains:

* `archive_block` - Select the archive types to block. Valid values: `encrypted` `corrupted` `partiallycorrupted` `multipart` `nested` `mailbomb` `timeout` `unhandled` .
* `archive_log` - Select the archive types to log. Valid values: `encrypted` `corrupted` `partiallycorrupted` `multipart` `nested` `mailbomb` `timeout` `unhandled` .
* `av_scan` - Enable AntiVirus scan service. Valid values: `disable` `block` `monitor` .
* `emulator` - Enable/disable the virus emulator. Valid values: `enable` `disable` .
* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default` `virus` .
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable` `block` `monitor` .
* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable` `block` `monitor` .
* `options` - Enable/disable MAPI AntiVirus scanning, monitoring, and quarantine. Valid values: `scan` `avmonitor` `quarantine` .
* `outbreak_prevention` - Enable virus outbreak prevention service. Valid values: `disable` `block` `monitor` .
* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable` `enable` .
* `nac_quar` - Configure AntiVirus quarantine settings. The structure of `nac_quar` block is documented below.

The `nac_quar` block contains:

* `expiry` - Duration of quarantine.
* `infected` - Enable/Disable quarantining infected hosts to the banned user list. Valid values: `none` `quar-src-ip` .
* `log` - Enable/disable AntiVirus quarantine logging. Valid values: `enable` `disable` .
* `nntp` - Configure NNTP AntiVirus options. The structure of `nntp` block is documented below.

The `nntp` block contains:

* `archive_block` - Select the archive types to block. Valid values: `encrypted` `corrupted` `partiallycorrupted` `multipart` `nested` `mailbomb` `timeout` `unhandled` .
* `archive_log` - Select the archive types to log. Valid values: `encrypted` `corrupted` `partiallycorrupted` `multipart` `nested` `mailbomb` `timeout` `unhandled` .
* `av_scan` - Enable AntiVirus scan service. Valid values: `disable` `block` `monitor` .
* `emulator` - Enable/disable the virus emulator. Valid values: `enable` `disable` .
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable` `block` `monitor` .
* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable` `block` `monitor` .
* `options` - Enable/disable NNTP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan` `avmonitor` `quarantine` .
* `outbreak_prevention` - Enable virus outbreak prevention service. Valid values: `disable` `block` `monitor` .
* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable` `enable` .
* `outbreak_prevention` - Configure Virus Outbreak Prevention settings. The structure of `outbreak_prevention` block is documented below.

The `outbreak_prevention` block contains:

* `external_blocklist` - Enable/disable external malware blocklist. Valid values: `disable` `enable` .
* `ftgd_service` - Enable/disable FortiGuard Virus outbreak prevention service. Valid values: `disable` `enable` .
* `pop3` - Configure POP3 AntiVirus options. The structure of `pop3` block is documented below.

The `pop3` block contains:

* `archive_block` - Select the archive types to block. Valid values: `encrypted` `corrupted` `partiallycorrupted` `multipart` `nested` `mailbomb` `timeout` `unhandled` .
* `archive_log` - Select the archive types to log. Valid values: `encrypted` `corrupted` `partiallycorrupted` `multipart` `nested` `mailbomb` `timeout` `unhandled` .
* `av_scan` - Enable AntiVirus scan service. Valid values: `disable` `block` `monitor` .
* `content_disarm` - Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan. Valid values: `disable` `enable` .
* `emulator` - Enable/disable the virus emulator. Valid values: `enable` `disable` .
* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default` `virus` .
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable` `block` `monitor` .
* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable` `block` `monitor` .
* `options` - Enable/disable POP3 AntiVirus scanning, monitoring, and quarantine. Valid values: `scan` `avmonitor` `quarantine` .
* `outbreak_prevention` - Enable virus outbreak prevention service. Valid values: `disable` `block` `monitor` .
* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable` `enable` .
* `smtp` - Configure SMTP AntiVirus options. The structure of `smtp` block is documented below.

The `smtp` block contains:

* `archive_block` - Select the archive types to block. Valid values: `encrypted` `corrupted` `partiallycorrupted` `multipart` `nested` `mailbomb` `timeout` `unhandled` .
* `archive_log` - Select the archive types to log. Valid values: `encrypted` `corrupted` `partiallycorrupted` `multipart` `nested` `mailbomb` `timeout` `unhandled` .
* `av_scan` - Enable AntiVirus scan service. Valid values: `disable` `block` `monitor` .
* `content_disarm` - Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan. Valid values: `disable` `enable` .
* `emulator` - Enable/disable the virus emulator. Valid values: `enable` `disable` .
* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default` `virus` .
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable` `block` `monitor` .
* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable` `block` `monitor` .
* `options` - Enable/disable SMTP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan` `avmonitor` `quarantine` .
* `outbreak_prevention` - Enable virus outbreak prevention service. Valid values: `disable` `block` `monitor` .
* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable` `enable` .
* `ssh` - Configure SFTP and SCP AntiVirus options. The structure of `ssh` block is documented below.

The `ssh` block contains:

* `archive_block` - Select the archive types to block. Valid values: `encrypted` `corrupted` `partiallycorrupted` `multipart` `nested` `mailbomb` `timeout` `unhandled` .
* `archive_log` - Select the archive types to log. Valid values: `encrypted` `corrupted` `partiallycorrupted` `multipart` `nested` `mailbomb` `timeout` `unhandled` .
* `av_scan` - Enable AntiVirus scan service. Valid values: `disable` `block` `monitor` .
* `emulator` - Enable/disable the virus emulator. Valid values: `enable` `disable` .
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable` `block` `monitor` .
* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable` `block` `monitor` .
* `options` - Enable/disable SFTP and SCP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan` `avmonitor` `quarantine` .
* `outbreak_prevention` - Enable virus outbreak prevention service. Valid values: `disable` `block` `monitor` .
* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable` `enable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_antivirus_profile can be imported using:
```sh
terraform import fortios_antivirus_profile.labelname {{mkey}}
```
