---
subcategory: "FortiGate Antivirus"
layout: "fortios"
page_title: "FortiOS: fortios_antivirus_quarantine"
description: |-
  Get information on a fortios Configure quarantine options.
---

# Data Source: fortios_antivirus_quarantine
Use this data source to get information on a fortios Configure quarantine options.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `agelimit` - Age limit for quarantined files (0 - 479 hours, 0 means forever).
* `destination` - Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them.
* `drop_blocked` - Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
* `drop_heuristic` - Do not quarantine files detected by heuristics found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
* `drop_infected` - Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
* `drop_machine_learning` - Do not quarantine files detected by machine learning found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
* `lowspace` - Select the method for handling additional files when running low on disk space.
* `maxfilesize` - Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).
* `quarantine_quota` - The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).
* `store_blocked` - Quarantine blocked files found in sessions using the selected protocols.
* `store_heuristic` - Quarantine files detected by heuristics found in sessions using the selected protocols.
* `store_infected` - Quarantine infected files found in sessions using the selected protocols.
* `store_machine_learning` - Quarantine files detected by machine learning found in sessions using the selected protocols.
