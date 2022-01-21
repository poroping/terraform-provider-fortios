---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_hsprofile"
description: |-
  Configure hotspot profile.
---

## fortios_wirelesscontrollerhotspot20_hsprofile
Configure hotspot profile.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `3gpp_plmn` - 3GPP PLMN name. This attribute must reference one of the following datasources: `wireless-controller.hotspot20.anqp-3gpp-cellular.name` .
* `access_network_asra` - Enable/disable additional step required for access (ASRA). Valid values: `enable` `disable` .
* `access_network_esr` - Enable/disable emergency services reachable (ESR). Valid values: `enable` `disable` .
* `access_network_internet` - Enable/disable connectivity to the Internet. Valid values: `enable` `disable` .
* `access_network_type` - Access network type. Valid values: `private-network` `private-network-with-guest-access` `chargeable-public-network` `free-public-network` `personal-device-network` `emergency-services-only-network` `test-or-experimental` `wildcard` .
* `access_network_uesa` - Enable/disable unauthenticated emergency service accessible (UESA). Valid values: `enable` `disable` .
* `advice_of_charge` - Advice of charge. This attribute must reference one of the following datasources: `wireless-controller.hotspot20.h2qp-advice-of-charge.name` .
* `anqp_domain_id` - ANQP Domain ID (0-65535).
* `bss_transition` - Enable/disable basic service set (BSS) transition Support. Valid values: `enable` `disable` .
* `conn_cap` - Connection capability name. This attribute must reference one of the following datasources: `wireless-controller.hotspot20.h2qp-conn-capability.name` .
* `deauth_request_timeout` - Deauthentication request timeout (in seconds).
* `dgaf` - Enable/disable downstream group-addressed forwarding (DGAF). Valid values: `enable` `disable` .
* `domain_name` - Domain name.
* `gas_comeback_delay` - GAS comeback delay (0 or 100 - 10000 milliseconds, default = 500).
* `gas_fragmentation_limit` - GAS fragmentation limit (512 - 4096, default = 1024).
* `hessid` - Homogeneous extended service set identifier (HESSID).
* `ip_addr_type` - IP address type name. This attribute must reference one of the following datasources: `wireless-controller.hotspot20.anqp-ip-address-type.name` .
* `l2tif` - Enable/disable Layer 2 traffic inspection and filtering. Valid values: `enable` `disable` .
* `nai_realm` - NAI realm list name. This attribute must reference one of the following datasources: `wireless-controller.hotspot20.anqp-nai-realm.name` .
* `name` - Hotspot profile name.
* `network_auth` - Network authentication name. This attribute must reference one of the following datasources: `wireless-controller.hotspot20.anqp-network-auth-type.name` .
* `oper_friendly_name` - Operator friendly name. This attribute must reference one of the following datasources: `wireless-controller.hotspot20.h2qp-operator-name.name` .
* `oper_icon` - Operator icon. This attribute must reference one of the following datasources: `wireless-controller.hotspot20.icon.name` .
* `osu_provider_nai` - OSU Provider NAI. This attribute must reference one of the following datasources: `wireless-controller.hotspot20.h2qp-osu-provider-nai.name` .
* `osu_ssid` - Online sign up (OSU) SSID.
* `pame_bi` - Enable/disable Pre-Association Message Exchange BSSID Independent (PAME-BI). Valid values: `disable` `enable` .
* `proxy_arp` - Enable/disable Proxy ARP. Valid values: `enable` `disable` .
* `qos_map` - QoS MAP set ID. This attribute must reference one of the following datasources: `wireless-controller.hotspot20.qos-map.name` .
* `release` - Hotspot 2.0 Release number (1, 2, 3, default = 2).
* `roaming_consortium` - Roaming consortium list name. This attribute must reference one of the following datasources: `wireless-controller.hotspot20.anqp-roaming-consortium.name` .
* `terms_and_conditions` - Terms and conditions. This attribute must reference one of the following datasources: `wireless-controller.hotspot20.h2qp-terms-and-conditions.name` .
* `venue_group` - Venue group. Valid values: `unspecified` `assembly` `business` `educational` `factory` `institutional` `mercantile` `residential` `storage` `utility` `vehicular` `outdoor` .
* `venue_name` - Venue name. This attribute must reference one of the following datasources: `wireless-controller.hotspot20.anqp-venue-name.name` .
* `venue_type` - Venue type. Valid values: `unspecified` `arena` `stadium` `passenger-terminal` `amphitheater` `amusement-park` `place-of-worship` `convention-center` `library` `museum` `restaurant` `theater` `bar` `coffee-shop` `zoo-or-aquarium` `emergency-center` `doctor-office` `bank` `fire-station` `police-station` `post-office` `professional-office` `research-facility` `attorney-office` `primary-school` `secondary-school` `university-or-college` `factory` `hospital` `long-term-care-facility` `rehab-center` `group-home` `prison-or-jail` `retail-store` `grocery-market` `auto-service-station` `shopping-mall` `gas-station` `private` `hotel-or-motel` `dormitory` `boarding-house` `automobile` `airplane` `bus` `ferry` `ship-or-boat` `train` `motor-bike` `muni-mesh-network` `city-park` `rest-area` `traffic-control` `bus-stop` `kiosk` .
* `venue_url` - Venue name. This attribute must reference one of the following datasources: `wireless-controller.hotspot20.anqp-venue-url.name` .
* `wan_metrics` - WAN metric name. This attribute must reference one of the following datasources: `wireless-controller.hotspot20.h2qp-wan-metric.name` .
* `wnm_sleep_mode` - Enable/disable wireless network management (WNM) sleep mode. Valid values: `enable` `disable` .
* `osu_provider` - Manually selected list of OSU provider(s). The structure of `osu_provider` block is documented below.

The `osu_provider` block contains:

* `name` - OSU provider name. This attribute must reference one of the following datasources: `wireless-controller.hotspot20.h2qp-osu-provider.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontrollerhotspot20_hsprofile can be imported using:
```sh
terraform import fortios_wirelesscontrollerhotspot20_hsprofile.labelname {{mkey}}
```
