---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemlldp_networkpolicy"
description: |-
  Get information on a fortios Configure LLDP network policy.
---

# Data Source: fortios_systemlldp_networkpolicy
Use this data source to get information on a fortios Configure LLDP network policy.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) LLDP network policy name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `name` - LLDP network policy name.
* `guest` - Guest.The structure of `guest` block is documented below.

The `guest` block contains:

* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `guest_voice_signaling` - Guest Voice Signaling.The structure of `guest_voice_signaling` block is documented below.

The `guest_voice_signaling` block contains:

* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `softphone` - Softphone.The structure of `softphone` block is documented below.

The `softphone` block contains:

* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `streaming_video` - Streaming Video.The structure of `streaming_video` block is documented below.

The `streaming_video` block contains:

* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `video_conferencing` - Video Conferencing.The structure of `video_conferencing` block is documented below.

The `video_conferencing` block contains:

* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `video_signaling` - Video Signaling.The structure of `video_signaling` block is documented below.

The `video_signaling` block contains:

* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `voice` - Voice.The structure of `voice` block is documented below.

The `voice` block contains:

* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
* `voice_signaling` - Voice signaling.The structure of `voice_signaling` block is documented below.

The `voice_signaling` block contains:

* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `status` - Enable/disable advertising this policy.
* `tag` - Advertise tagged or untagged traffic.
* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).
