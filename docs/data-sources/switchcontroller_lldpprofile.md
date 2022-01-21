---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_lldpprofile"
description: |-
  Get information on a fortios Configure FortiSwitch LLDP profiles.
---

# Data Source: fortios_switchcontroller_lldpprofile
Use this data source to get information on a fortios Configure FortiSwitch LLDP profiles.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `8021_tlvs` - Transmitted IEEE 802.1 TLVs.
* `8023_tlvs` - Transmitted IEEE 802.3 TLVs.
* `auto_isl` - Enable/disable auto inter-switch LAG.
* `auto_isl_hello_timer` - Auto inter-switch LAG hello timer duration (1 - 30 sec, default = 3).
* `auto_isl_port_group` - Auto inter-switch LAG port group ID (0 - 9).
* `auto_isl_receive_timeout` - Auto inter-switch LAG timeout if no response is received (3 - 90 sec, default = 9).
* `auto_mclag_icl` - Enable/disable MCLAG inter chassis link.
* `med_tlvs` - Transmitted LLDP-MED TLVs (type-length-value descriptions).
* `name` - Profile name.
* `custom_tlvs` - Configuration method to edit custom TLV entries.The structure of `custom_tlvs` block is documented below.

The `custom_tlvs` block contains:

* `information_string` - Organizationally defined information string (0 - 507 hexadecimal bytes).
* `name` - TLV name (not sent).
* `oui` - Organizationally unique identifier (OUI), a 3-byte hexadecimal number, for this TLV.
* `subtype` - Organizationally defined subtype (0 - 255).
* `med_location_service` - Configuration method to edit Media Endpoint Discovery (MED) location service type-length-value (TLV) categories.The structure of `med_location_service` block is documented below.

The `med_location_service` block contains:

* `name` - Location service type name.
* `status` - Enable or disable this TLV.
* `sys_location_id` - Location service ID.
* `med_network_policy` - Configuration method to edit Media Endpoint Discovery (MED) network policy type-length-value (TLV) categories.The structure of `med_network_policy` block is documented below.

The `med_network_policy` block contains:

* `assign_vlan` - Enable/disable VLAN assignment when this profile is applied on managed FortiSwitch port.
* `dscp` - Advertised Differentiated Services Code Point (DSCP) value, a packet header value indicating the level of service requested for traffic, such as high priority or best effort delivery.
* `name` - Policy type name.
* `priority` - Advertised Layer 2 priority (0 - 7; from lowest to highest priority).
* `status` - Enable or disable this TLV.
* `vlan_intf` - VLAN interface to advertise; if configured on port.
