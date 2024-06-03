---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_accprofile"
description: |-
  Configure access profiles for system administrators.
---

## fortios_system_accprofile
Configure access profiles for system administrators.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `admintimeout` - Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
* `admintimeout_override` - Enable/disable overriding the global administrator idle timeout. Valid values: `enable` `disable` .
* `authgrp` - Administrator access to Users and Devices. Valid values: `none` `read` `read-write` .
* `comments` - Comment.
* `ftviewgrp` - FortiView. Valid values: `none` `read` `read-write` .
* `fwgrp` - Administrator access to the Firewall configuration. Valid values: `none` `read` `read-write` `custom` .
* `loggrp` - Administrator access to Logging and Reporting including viewing log messages. Valid values: `none` `read` `read-write` `custom` .
* `name` - Profile name.
* `netgrp` - Network Configuration. Valid values: `none` `read` `read-write` `custom` .
* `scope` - Scope of admin access: global or specific VDOM(s). Valid values: `vdom` `global` .
* `secfabgrp` - Security Fabric. Valid values: `none` `read` `read-write` .
* `sysgrp` - System Configuration. Valid values: `none` `read` `read-write` `custom` .
* `system_diagnostics` - Enable/disable permission to run system diagnostic commands. Valid values: `enable` `disable` .
* `system_execute_ssh` - Enable/disable permission to execute SSH commands. Valid values: `enable` `disable` .
* `system_execute_telnet` - Enable/disable permission to execute TELNET commands. Valid values: `enable` `disable` .
* `utmgrp` - Administrator access to Security Profiles. Valid values: `none` `read` `read-write` `custom` .
* `vpngrp` - Administrator access to IPsec, SSL, PPTP, and L2TP VPN. Valid values: `none` `read` `read-write` .
* `wanoptgrp` - Administrator access to WAN Opt & Cache. Valid values: `none` `read` `read-write` .
* `wifi` - Administrator access to the WiFi controller and Switch controller. Valid values: `none` `read` `read-write` .
* `fwgrp_permission` - Custom firewall permission. The structure of `fwgrp_permission` block is documented below.

The `fwgrp_permission` block contains:

* `address` - Address Configuration. Valid values: `none` `read` `read-write` .
* `others` - Other Firewall Configuration. Valid values: `none` `read` `read-write` .
* `policy` - Policy Configuration. Valid values: `none` `read` `read-write` .
* `schedule` - Schedule Configuration. Valid values: `none` `read` `read-write` .
* `service` - Service Configuration. Valid values: `none` `read` `read-write` .
* `loggrp_permission` - Custom Log & Report permission. The structure of `loggrp_permission` block is documented below.

The `loggrp_permission` block contains:

* `config` - Log & Report configuration. Valid values: `none` `read` `read-write` .
* `data_access` - Log & Report Data Access. Valid values: `none` `read` `read-write` .
* `report_access` - Log & Report Report Access. Valid values: `none` `read` `read-write` .
* `threat_weight` - Log & Report Threat Weight. Valid values: `none` `read` `read-write` .
* `netgrp_permission` - Custom network permission. The structure of `netgrp_permission` block is documented below.

The `netgrp_permission` block contains:

* `cfg` - Network Configuration. Valid values: `none` `read` `read-write` .
* `packet_capture` - Packet Capture Configuration. Valid values: `none` `read` `read-write` .
* `route_cfg` - Router Configuration. Valid values: `none` `read` `read-write` .
* `sysgrp_permission` - Custom system permission. The structure of `sysgrp_permission` block is documented below.

The `sysgrp_permission` block contains:

* `admin` - Administrator Users. Valid values: `none` `read` `read-write` .
* `cfg` - System Configuration. Valid values: `none` `read` `read-write` .
* `mnt` - Maintenance. Valid values: `none` `read` `read-write` .
* `upd` - FortiGuard Updates. Valid values: `none` `read` `read-write` .
* `utmgrp_permission` - Custom Security Profile permissions. The structure of `utmgrp_permission` block is documented below.

The `utmgrp_permission` block contains:

* `antivirus` - Antivirus profiles and settings. Valid values: `none` `read` `read-write` .
* `application_control` - Application Control profiles and settings. Valid values: `none` `read` `read-write` .
* `data_leak_prevention` - DLP profiles and settings. Valid values: `none` `read` `read-write` .
* `data_loss_prevention` - DLP profiles and settings. Valid values: `none` `read` `read-write` .
* `dnsfilter` - DNS Filter profiles and settings. Valid values: `none` `read` `read-write` .
* `emailfilter` - Email Filter and settings. Valid values: `none` `read` `read-write` .
* `endpoint_control` - FortiClient Profiles. Valid values: `none` `read` `read-write` .
* `file_filter` - File-filter profiles and settings. Valid values: `none` `read` `read-write` .
* `icap` - ICAP profiles and settings. Valid values: `none` `read` `read-write` .
* `ips` - IPS profiles and settings. Valid values: `none` `read` `read-write` .
* `videofilter` - Video filter profiles and settings. Valid values: `none` `read` `read-write` .
* `voip` - VoIP profiles and settings. Valid values: `none` `read` `read-write` .
* `waf` - Web Application Firewall profiles and settings. Valid values: `none` `read` `read-write` .
* `webfilter` - Web Filter profiles and settings. Valid values: `none` `read` `read-write` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_accprofile can be imported using:
```sh
terraform import fortios_system_accprofile.labelname {{mkey}}
```
