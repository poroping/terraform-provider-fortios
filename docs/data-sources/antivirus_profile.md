---
subcategory: "FortiGate Antivirus"
layout: "fortios"
page_title: "FortiOS: fortios_antivirus_profile"
description: |-
  Get information on a fortios Configure AntiVirus profiles.
---

# Data Source: fortios_antivirus_profile
Use this data source to get information on a fortios Configure AntiVirus profiles.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `analytics_accept_filetype` - Only submit files matching this DLP file-pattern to FortiSandbox (post-transfer scan only).
* `analytics_bl_filetype` - Only submit files matching this DLP file-pattern to FortiSandbox.
* `analytics_db` - Enable/disable using the FortiSandbox signature database to supplement the AV signature databases.
* `analytics_ignore_filetype` - Do not submit files matching this DLP file-pattern to FortiSandbox (post-transfer scan only).
* `analytics_max_upload` - Maximum size of files that can be uploaded to FortiSandbox.
* `analytics_wl_filetype` - Do not submit files matching this DLP file-pattern to FortiSandbox.
* `av_block_log` - Enable/disable logging for AntiVirus file blocking.
* `av_virus_log` - Enable/disable AntiVirus logging.
* `comment` - Comment.
* `ems_threat_feed` - Enable/disable use of EMS threat feed when performing AntiVirus scan. Analyzes files including the content of archives.
* `extended_log` - Enable/disable extended logging for antivirus.
* `external_blocklist_archive_scan` - Enable/disable external-blocklist archive scanning.
* `external_blocklist_enable_all` - Enable/disable all external blocklists.
* `feature_set` - Flow/proxy feature set.
* `fortiai_error_action` - Action to take if FortiAI encounters an error.
* `fortiai_timeout_action` - Action to take if FortiAI encounters a scan timeout.
* `fortindr_error_action` - Action to take if FortiNDR encounters an error.
* `fortindr_timeout_action` - Action to take if FortiNDR encounters a scan timeout.
* `fortisandbox_error_action` - Action to take if FortiSandbox inline scan encounters an error.
* `fortisandbox_max_upload` - Maximum size of files that can be uploaded to FortiSandbox.
* `fortisandbox_mode` - FortiSandbox scan modes.
* `fortisandbox_timeout_action` - Action to take if FortiSandbox inline scan encounters a scan timeout.
* `ftgd_analytics` - Settings to control which files are uploaded to FortiSandbox.
* `mobile_malware_db` - Enable/disable using the mobile malware signature database.
* `name` - Profile name.
* `outbreak_prevention_archive_scan` - Enable/disable outbreak-prevention archive scanning.
* `replacemsg_group` - Replacement message group customized for this profile.
* `scan_mode` - Configure scan mode (default or legacy).
* `cifs` - Configure CIFS AntiVirus options.The structure of `cifs` block is documented below.

The `cifs` block contains:

* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `av_scan` - Enable AntiVirus scan service.
* `emulator` - Enable/disable the virus emulator.
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives.
* `fortiai` - Enable/disable scanning of files by FortiAI.
* `fortindr` - Enable scanning of files by FortiNDR.
* `fortisandbox` - Enable scanning of files by FortiSandbox.
* `options` - Enable/disable CIFS AntiVirus scanning, monitoring, and quarantine.
* `outbreak_prevention` - Enable virus outbreak prevention service.
* `quarantine` - Enable/disable quarantine for infected files.
* `content_disarm` - AV Content Disarm and Reconstruction settings.The structure of `content_disarm` block is documented below.

The `content_disarm` block contains:

* `cover_page` - Enable/disable inserting a cover page into the disarmed document.
* `detect_only` - Enable/disable only detect disarmable files, do not alter content.
* `error_action` - Action to be taken if CDR engine encounters an unrecoverable error.
* `office_action` - Enable/disable stripping of PowerPoint action events in Microsoft Office documents.
* `office_dde` - Enable/disable stripping of Dynamic Data Exchange events in Microsoft Office documents.
* `office_embed` - Enable/disable stripping of embedded objects in Microsoft Office documents.
* `office_hylink` - Enable/disable stripping of hyperlinks in Microsoft Office documents.
* `office_linked` - Enable/disable stripping of linked objects in Microsoft Office documents.
* `office_macro` - Enable/disable stripping of macros in Microsoft Office documents.
* `original_file_destination` - Destination to send original file if active content is removed.
* `pdf_act_form` - Enable/disable stripping of PDF document actions that submit data to other targets.
* `pdf_act_gotor` - Enable/disable stripping of PDF document actions that access other PDF documents.
* `pdf_act_java` - Enable/disable stripping of PDF document actions that execute JavaScript code.
* `pdf_act_launch` - Enable/disable stripping of PDF document actions that launch other applications.
* `pdf_act_movie` - Enable/disable stripping of PDF document actions that play a movie.
* `pdf_act_sound` - Enable/disable stripping of PDF document actions that play a sound.
* `pdf_embedfile` - Enable/disable stripping of embedded files in PDF documents.
* `pdf_hyperlink` - Enable/disable stripping of hyperlinks from PDF documents.
* `pdf_javacode` - Enable/disable stripping of JavaScript code in PDF documents.
* `external_blocklist` - One or more external malware block lists.The structure of `external_blocklist` block is documented below.

The `external_blocklist` block contains:

* `name` - External blocklist.
* `ftp` - Configure FTP AntiVirus options.The structure of `ftp` block is documented below.

The `ftp` block contains:

* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `av_scan` - Enable AntiVirus scan service.
* `emulator` - Enable/disable the virus emulator.
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives.
* `fortiai` - Enable/disable scanning of files by FortiAI.
* `fortindr` - Enable scanning of files by FortiNDR.
* `fortisandbox` - Enable scanning of files by FortiSandbox.
* `options` - Enable/disable FTP AntiVirus scanning, monitoring, and quarantine.
* `outbreak_prevention` - Enable virus outbreak prevention service.
* `quarantine` - Enable/disable quarantine for infected files.
* `http` - Configure HTTP AntiVirus options.The structure of `http` block is documented below.

The `http` block contains:

* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `av_scan` - Enable AntiVirus scan service.
* `content_disarm` - Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan.
* `emulator` - Enable/disable the virus emulator.
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives.
* `fortiai` - Enable/disable scanning of files by FortiAI.
* `fortindr` - Enable scanning of files by FortiNDR.
* `fortisandbox` - Enable scanning of files by FortiSandbox.
* `options` - Enable/disable HTTP AntiVirus scanning, monitoring, and quarantine.
* `outbreak_prevention` - Enable virus outbreak prevention service.
* `quarantine` - Enable/disable quarantine for infected files.
* `unknown_content_encoding` - Configure the action the FortiGate unit will take on unknown content-encoding.
* `imap` - Configure IMAP AntiVirus options.The structure of `imap` block is documented below.

The `imap` block contains:

* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `av_scan` - Enable AntiVirus scan service.
* `content_disarm` - Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan.
* `emulator` - Enable/disable the virus emulator.
* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring.
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives.
* `fortiai` - Enable/disable scanning of files by FortiAI.
* `fortindr` - Enable scanning of files by FortiNDR.
* `fortisandbox` - Enable scanning of files by FortiSandbox.
* `options` - Enable/disable IMAP AntiVirus scanning, monitoring, and quarantine.
* `outbreak_prevention` - Enable virus outbreak prevention service.
* `quarantine` - Enable/disable quarantine for infected files.
* `mapi` - Configure MAPI AntiVirus options.The structure of `mapi` block is documented below.

The `mapi` block contains:

* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `av_scan` - Enable AntiVirus scan service.
* `emulator` - Enable/disable the virus emulator.
* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring.
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives.
* `fortiai` - Enable/disable scanning of files by FortiAI.
* `fortindr` - Enable scanning of files by FortiNDR.
* `fortisandbox` - Enable scanning of files by FortiSandbox.
* `options` - Enable/disable MAPI AntiVirus scanning, monitoring, and quarantine.
* `outbreak_prevention` - Enable virus outbreak prevention service.
* `quarantine` - Enable/disable quarantine for infected files.
* `nac_quar` - Configure AntiVirus quarantine settings.The structure of `nac_quar` block is documented below.

The `nac_quar` block contains:

* `expiry` - Duration of quarantine.
* `infected` - Enable/Disable quarantining infected hosts to the banned user list.
* `log` - Enable/disable AntiVirus quarantine logging.
* `nntp` - Configure NNTP AntiVirus options.The structure of `nntp` block is documented below.

The `nntp` block contains:

* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `av_scan` - Enable AntiVirus scan service.
* `emulator` - Enable/disable the virus emulator.
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives.
* `fortiai` - Enable/disable scanning of files by FortiAI.
* `fortindr` - Enable scanning of files by FortiNDR.
* `fortisandbox` - Enable scanning of files by FortiSandbox.
* `options` - Enable/disable NNTP AntiVirus scanning, monitoring, and quarantine.
* `outbreak_prevention` - Enable virus outbreak prevention service.
* `quarantine` - Enable/disable quarantine for infected files.
* `outbreak_prevention` - Configure Virus Outbreak Prevention settings.The structure of `outbreak_prevention` block is documented below.

The `outbreak_prevention` block contains:

* `external_blocklist` - Enable/disable external malware blocklist.
* `ftgd_service` - Enable/disable FortiGuard Virus outbreak prevention service.
* `pop3` - Configure POP3 AntiVirus options.The structure of `pop3` block is documented below.

The `pop3` block contains:

* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `av_scan` - Enable AntiVirus scan service.
* `content_disarm` - Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan.
* `emulator` - Enable/disable the virus emulator.
* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring.
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives.
* `fortiai` - Enable/disable scanning of files by FortiAI.
* `fortindr` - Enable scanning of files by FortiNDR.
* `fortisandbox` - Enable scanning of files by FortiSandbox.
* `options` - Enable/disable POP3 AntiVirus scanning, monitoring, and quarantine.
* `outbreak_prevention` - Enable virus outbreak prevention service.
* `quarantine` - Enable/disable quarantine for infected files.
* `smtp` - Configure SMTP AntiVirus options.The structure of `smtp` block is documented below.

The `smtp` block contains:

* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `av_scan` - Enable AntiVirus scan service.
* `content_disarm` - Enable/disable Content Disarm and Reconstruction when performing AntiVirus scan.
* `emulator` - Enable/disable the virus emulator.
* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring.
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives.
* `fortiai` - Enable/disable scanning of files by FortiAI.
* `fortindr` - Enable scanning of files by FortiNDR.
* `fortisandbox` - Enable scanning of files by FortiSandbox.
* `options` - Enable/disable SMTP AntiVirus scanning, monitoring, and quarantine.
* `outbreak_prevention` - Enable virus outbreak prevention service.
* `quarantine` - Enable/disable quarantine for infected files.
* `ssh` - Configure SFTP and SCP AntiVirus options.The structure of `ssh` block is documented below.

The `ssh` block contains:

* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `av_scan` - Enable AntiVirus scan service.
* `emulator` - Enable/disable the virus emulator.
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives.
* `fortiai` - Enable/disable scanning of files by FortiAI.
* `fortindr` - Enable scanning of files by FortiNDR.
* `fortisandbox` - Enable scanning of files by FortiSandbox.
* `options` - Enable/disable SFTP and SCP AntiVirus scanning, monitoring, and quarantine.
* `outbreak_prevention` - Enable virus outbreak prevention service.
* `quarantine` - Enable/disable quarantine for infected files.
