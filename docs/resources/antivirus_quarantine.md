---
subcategory: "FortiGate Antivirus"
layout: "fortios"
page_title: "FortiOS: fortios_antivirus_quarantine"
description: |-
  Configure quarantine options.
---

## fortios_antivirus_quarantine
Configure quarantine options.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `agelimit` - Age limit for quarantined files (0 - 479 hours, 0 means forever).
* `destination` - Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them. Valid values: `NULL` `disk` `FortiAnalyzer` .
* `drop_blocked` - Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap` `smtp` `pop3` `http` `ftp` `nntp` `imaps` `smtps` `pop3s` `ftps` `mapi` `cifs` `ssh` .
* `drop_heuristic` - Do not quarantine files detected by heuristics found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap` `smtp` `pop3` `http` `ftp` `nntp` `imaps` `smtps` `pop3s` `https` `ftps` `mapi` `cifs` `ssh` .
* `drop_infected` - Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap` `smtp` `pop3` `http` `ftp` `nntp` `imaps` `smtps` `pop3s` `https` `ftps` `mapi` `cifs` `ssh` .
* `drop_machine_learning` - Do not quarantine files detected by machine learning found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined. Valid values: `imap` `smtp` `pop3` `http` `ftp` `nntp` `imaps` `smtps` `pop3s` `https` `ftps` `mapi` `cifs` `ssh` .
* `lowspace` - Select the method for handling additional files when running low on disk space. Valid values: `drop-new` `ovrw-old` .
* `maxfilesize` - Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).
* `quarantine_quota` - The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).
* `store_blocked` - Quarantine blocked files found in sessions using the selected protocols. Valid values: `imap` `smtp` `pop3` `http` `ftp` `nntp` `imaps` `smtps` `pop3s` `ftps` `mapi` `cifs` `ssh` .
* `store_heuristic` - Quarantine files detected by heuristics found in sessions using the selected protocols. Valid values: `imap` `smtp` `pop3` `http` `ftp` `nntp` `imaps` `smtps` `pop3s` `https` `ftps` `mapi` `cifs` `ssh` .
* `store_infected` - Quarantine infected files found in sessions using the selected protocols. Valid values: `imap` `smtp` `pop3` `http` `ftp` `nntp` `imaps` `smtps` `pop3s` `https` `ftps` `mapi` `cifs` `ssh` .
* `store_machine_learning` - Quarantine files detected by machine learning found in sessions using the selected protocols. Valid values: `imap` `smtp` `pop3` `http` `ftp` `nntp` `imaps` `smtps` `pop3s` `https` `ftps` `mapi` `cifs` `ssh` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_antivirus_quarantine can be imported using:
```sh
terraform import fortios_antivirus_quarantine.labelname {{mkey}}
```
