---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_hsprofile"
description: |-
  Get information on a fortios Configure hotspot profile.
---

# Data Source: fortios_wirelesscontrollerhotspot20_hsprofile
Use this data source to get information on a fortios Configure hotspot profile.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Hotspot profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `3gpp_plmn` - 3GPP PLMN name.
* `access_network_asra` - Enable/disable additional step required for access (ASRA).
* `access_network_esr` - Enable/disable emergency services reachable (ESR).
* `access_network_internet` - Enable/disable connectivity to the Internet.
* `access_network_type` - Access network type.
* `access_network_uesa` - Enable/disable unauthenticated emergency service accessible (UESA).
* `advice_of_charge` - Advice of charge.
* `anqp_domain_id` - ANQP Domain ID (0-65535).
* `bss_transition` - Enable/disable basic service set (BSS) transition Support.
* `conn_cap` - Connection capability name.
* `deauth_request_timeout` - Deauthentication request timeout (in seconds).
* `dgaf` - Enable/disable downstream group-addressed forwarding (DGAF).
* `domain_name` - Domain name.
* `gas_comeback_delay` - GAS comeback delay (0 or 100 - 10000 milliseconds, default = 500).
* `gas_fragmentation_limit` - GAS fragmentation limit (512 - 4096, default = 1024).
* `hessid` - Homogeneous extended service set identifier (HESSID).
* `ip_addr_type` - IP address type name.
* `l2tif` - Enable/disable Layer 2 traffic inspection and filtering.
* `nai_realm` - NAI realm list name.
* `name` - Hotspot profile name.
* `network_auth` - Network authentication name.
* `oper_friendly_name` - Operator friendly name.
* `oper_icon` - Operator icon.
* `osu_provider_nai` - OSU Provider NAI.
* `osu_ssid` - Online sign up (OSU) SSID.
* `pame_bi` - Enable/disable Pre-Association Message Exchange BSSID Independent (PAME-BI).
* `proxy_arp` - Enable/disable Proxy ARP.
* `qos_map` - QoS MAP set ID.
* `release` - Hotspot 2.0 Release number (1, 2, 3, default = 2).
* `roaming_consortium` - Roaming consortium list name.
* `terms_and_conditions` - Terms and conditions.
* `venue_group` - Venue group.
* `venue_name` - Venue name.
* `venue_type` - Venue type.
* `venue_url` - Venue name.
* `wan_metrics` - WAN metric name.
* `wnm_sleep_mode` - Enable/disable wireless network management (WNM) sleep mode.
* `osu_provider` - Manually selected list of OSU provider(s).The structure of `osu_provider` block is documented below.

The `osu_provider` block contains:

* `name` - OSU provider name.
