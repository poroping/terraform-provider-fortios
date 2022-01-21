---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_accprofile"
description: |-
  Get information on a fortios Configure access profiles for system administrators.
---

# Data Source: fortios_system_accprofile
Use this data source to get information on a fortios Configure access profiles for system administrators.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `admintimeout` - Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
* `admintimeout_override` - Enable/disable overriding the global administrator idle timeout.
* `authgrp` - Administrator access to Users and Devices.
* `comments` - Comment.
* `ftviewgrp` - FortiView.
* `fwgrp` - Administrator access to the Firewall configuration.
* `loggrp` - Administrator access to Logging and Reporting including viewing log messages.
* `name` - Profile name.
* `netgrp` - Network Configuration.
* `scope` - Scope of admin access: global or specific VDOM(s).
* `secfabgrp` - Security Fabric.
* `sysgrp` - System Configuration.
* `system_diagnostics` - Enable/disable permission to run system diagnostic commands.
* `utmgrp` - Administrator access to Security Profiles.
* `vpngrp` - Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
* `wanoptgrp` - Administrator access to WAN Opt & Cache.
* `wifi` - Administrator access to the WiFi controller and Switch controller.
* `fwgrp_permission` - Custom firewall permission.The structure of `fwgrp_permission` block is documented below.

The `fwgrp_permission` block contains:

* `address` - Address Configuration.
* `others` - Other Firewall Configuration.
* `policy` - Policy Configuration.
* `schedule` - Schedule Configuration.
* `service` - Service Configuration.
* `loggrp_permission` - Custom Log & Report permission.The structure of `loggrp_permission` block is documented below.

The `loggrp_permission` block contains:

* `config` - Log & Report configuration.
* `data_access` - Log & Report Data Access.
* `report_access` - Log & Report Report Access.
* `threat_weight` - Log & Report Threat Weight.
* `netgrp_permission` - Custom network permission.The structure of `netgrp_permission` block is documented below.

The `netgrp_permission` block contains:

* `cfg` - Network Configuration.
* `packet_capture` - Packet Capture Configuration.
* `route_cfg` - Router Configuration.
* `sysgrp_permission` - Custom system permission.The structure of `sysgrp_permission` block is documented below.

The `sysgrp_permission` block contains:

* `admin` - Administrator Users.
* `cfg` - System Configuration.
* `mnt` - Maintenance.
* `upd` - FortiGuard Updates.
* `utmgrp_permission` - Custom Security Profile permissions.The structure of `utmgrp_permission` block is documented below.

The `utmgrp_permission` block contains:

* `antivirus` - Antivirus profiles and settings.
* `application_control` - Application Control profiles and settings.
* `data_loss_prevention` - DLP profiles and settings.
* `dnsfilter` - DNS Filter profiles and settings.
* `emailfilter` - Email Filter and settings.
* `endpoint_control` - FortiClient Profiles.
* `file_filter` - File-filter profiles and settings.
* `icap` - ICAP profiles and settings.
* `ips` - IPS profiles and settings.
* `voip` - VoIP profiles and settings.
* `waf` - Web Application Firewall profiles and settings.
* `webfilter` - Web Filter profiles and settings.
