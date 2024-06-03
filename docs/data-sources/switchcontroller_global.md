---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_global"
description: |-
  Get information on a fortios Configure FortiSwitch global settings.
---

# Data Source: fortios_switchcontroller_global
Use this data source to get information on a fortios Configure FortiSwitch global settings.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `allow_multiple_interfaces` - Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate.
* `bounce_quarantined_link` - Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device.
* `default_virtual_switch_vlan` - Default VLAN for ports when added to the virtual-switch.
* `dhcp_server_access_list` - Enable/disable DHCP snooping server access list.
* `fips_enforce` - Enable/disable enforcement of FIPS on managed FortiSwitch devices.
* `firmware_provision_on_authorization` - Enable/disable automatic provisioning of latest firmware on authorization.
* `https_image_push` - Enable/disable image push to FortiSwitch using HTTPS.
* `log_mac_limit_violations` - Enable/disable logs for Learning Limit Violations.
* `mac_aging_interval` - Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
* `mac_event_logging` - Enable/disable MAC address event logging.
* `mac_retention_period` - Time in hours after which an inactive MAC is removed from client DB (0 = aged out based on mac-aging-interval).
* `mac_violation_timer` - Set timeout for Learning Limit Violations (0 = disabled).
* `quarantine_mode` - Quarantine mode.
* `sn_dns_resolution` - Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number.
* `update_user_device` - Control which sources update the device user list.
* `vlan_all_mode` - VLAN configuration mode, user-defined-vlans or all-possible-vlans.
* `vlan_identity` - Identity of the VLAN. Commonly used for RADIUS Tunnel-Private-Group-Id.
* `vlan_optimization` - FortiLink VLAN optimization.
* `custom_command` - List of custom commands to be pushed to all FortiSwitches in the VDOM.The structure of `custom_command` block is documented below.

The `custom_command` block contains:

* `command_entry` - List of FortiSwitch commands.
* `command_name` - Name of custom command to push to all FortiSwitches in VDOM.
* `disable_discovery` - Prevent this FortiSwitch from discovering.The structure of `disable_discovery` block is documented below.

The `disable_discovery` block contains:

* `name` - Managed device ID.
