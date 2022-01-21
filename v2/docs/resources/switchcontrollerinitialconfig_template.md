---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerinitialconfig_template"
description: |-
  Configure template for auto-generated VLANs.
---

## fortios_switchcontrollerinitialconfig_template
Configure template for auto-generated VLANs.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `allowaccess` - Permitted types of management access to this interface. Valid values: `ping` `https` `ssh` `snmp` `http` `telnet` `fgfm` `radius-acct` `probe-response` `fabric` `ftm` .
* `auto_ip` - Automatically allocate interface address and subnet block. Valid values: `enable` `disable` .
* `dhcp_server` - Enable/disable a DHCP server on this interface. Valid values: `enable` `disable` .
* `ip` - Interface IPv4 address and subnet mask.
* `name` - Initial config template name
* `vlanid` - Unique VLAN ID.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_switchcontrollerinitialconfig_template can be imported using:
```sh
terraform import fortios_switchcontrollerinitialconfig_template.labelname {{mkey}}
```
