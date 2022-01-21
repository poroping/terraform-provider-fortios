---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_lldpprofile"
description: |-
  Configure FortiSwitch LLDP profiles.
---

## fortios_switchcontroller_lldpprofile
Configure FortiSwitch LLDP profiles.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `8021_tlvs` - Transmitted IEEE 802.1 TLVs. Valid values: `port-vlan-id` .
* `8023_tlvs` - Transmitted IEEE 802.3 TLVs. Valid values: `max-frame-size` `power-negotiation` .
* `auto_isl` - Enable/disable auto inter-switch LAG. Valid values: `disable` `enable` .
* `auto_isl_hello_timer` - Auto inter-switch LAG hello timer duration (1 - 30 sec, default = 3).
* `auto_isl_port_group` - Auto inter-switch LAG port group ID (0 - 9).
* `auto_isl_receive_timeout` - Auto inter-switch LAG timeout if no response is received (3 - 90 sec, default = 9).
* `auto_mclag_icl` - Enable/disable MCLAG inter chassis link. Valid values: `disable` `enable` .
* `med_tlvs` - Transmitted LLDP-MED TLVs (type-length-value descriptions). Valid values: `inventory-management` `network-policy` `power-management` `location-identification` .
* `name` - Profile name.
* `custom_tlvs` - Configuration method to edit custom TLV entries. The structure of `custom_tlvs` block is documented below.

The `custom_tlvs` block contains:

* `information_string` - Organizationally defined information string (0 - 507 hexadecimal bytes).
* `name` - TLV name (not sent).
* `oui` - Organizationally unique identifier (OUI), a 3-byte hexadecimal number, for this TLV.
* `subtype` - Organizationally defined subtype (0 - 255).
* `med_location_service` - Configuration method to edit Media Endpoint Discovery (MED) location service type-length-value (TLV) categories. The structure of `med_location_service` block is documented below.

The `med_location_service` block contains:

* `name` - Location service type name.
* `status` - Enable or disable this TLV. Valid values: `disable` `enable` .
* `sys_location_id` - Location service ID. This attribute must reference one of the following datasources: `switch-controller.location.name` .
* `med_network_policy` - Configuration method to edit Media Endpoint Discovery (MED) network policy type-length-value (TLV) categories. The structure of `med_network_policy` block is documented below.

The `med_network_policy` block contains:

* `assign_vlan` - Enable/disable VLAN assignment when this profile is applied on managed FortiSwitch port. Valid values: `disable` `enable` .
* `dscp` - Advertised Differentiated Services Code Point (DSCP) value, a packet header value indicating the level of service requested for traffic, such as high priority or best effort delivery.
* `name` - Policy type name.
* `priority` - Advertised Layer 2 priority (0 - 7; from lowest to highest priority).
* `status` - Enable or disable this TLV. Valid values: `disable` `enable` .
* `vlan_intf` - VLAN interface to advertise; if configured on port. This attribute must reference one of the following datasources: `system.interface.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_switchcontroller_lldpprofile can be imported using:
```sh
terraform import fortios_switchcontroller_lldpprofile.labelname {{mkey}}
```
