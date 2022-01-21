---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_wtp"
description: |-
  Configure Wireless Termination Points (WTPs), that is, FortiAPs or APs to be managed by FortiGate.
---

## fortios_wirelesscontroller_wtp
Configure Wireless Termination Points (WTPs), that is, FortiAPs or APs to be managed by FortiGate.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `wtp_id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `admin` - Configure how the FortiGate operating as a wireless controller discovers and manages this WTP, AP or FortiAP. Valid values: `discovered` `disable` `enable` .
* `allowaccess` - Control management access to the managed WTP, FortiAP, or AP. Separate entries with a space. Valid values: `https` `ssh` `snmp` .
* `apcfg_profile` - AP local configuration profile name. This attribute must reference one of the following datasources: `wireless-controller.apcfg-profile.name` .
* `bonjour_profile` - Bonjour profile name. This attribute must reference one of the following datasources: `wireless-controller.bonjour-profile.name` .
* `coordinate_latitude` - WTP latitude coordinate.
* `coordinate_longitude` - WTP longitude coordinate.
* `firmware_provision` - Firmware version to provision to this FortiAP on bootup (major.minor.build, i.e. 6.2.1234).
* `firmware_provision_latest` - Enable/disable one-time automatic provisioning of the latest firmware version. Valid values: `disable` `once` .
* `image_download` - Enable/disable WTP image download. Valid values: `enable` `disable` .
* `index` - Index (0 - 4294967295).
* `ip_fragment_preventing` - Method(s) by which IP fragmentation is prevented for control and data packets through CAPWAP tunnel (default = tcp-mss-adjust). Valid values: `tcp-mss-adjust` `icmp-unreachable` .
* `led_state` - Enable to allow the FortiAPs LEDs to light. Disable to keep the LEDs off. You may want to keep the LEDs off so they are not distracting in low light areas etc. Valid values: `enable` `disable` .
* `location` - Field for describing the physical location of the WTP, AP or FortiAP.
* `login_passwd` - Set the managed WTP, FortiAP, or AP's administrator password.
* `login_passwd_change` - Change or reset the administrator password of a managed WTP, FortiAP or AP (yes, default, or no, default = no). Valid values: `yes` `default` `no` .
* `mesh_bridge_enable` - Enable/disable mesh Ethernet bridge when WTP is configured as a mesh branch/leaf AP. Valid values: `default` `enable` `disable` .
* `name` - WTP, AP or FortiAP configuration name.
* `override_allowaccess` - Enable to override the WTP profile management access configuration. Valid values: `enable` `disable` .
* `override_ip_fragment` - Enable/disable overriding the WTP profile IP fragment prevention setting. Valid values: `enable` `disable` .
* `override_lan` - Enable to override the WTP profile LAN port setting. Valid values: `enable` `disable` .
* `override_led_state` - Enable to override the profile LED state setting for this FortiAP. You must enable this option to use the led-state command to turn off the FortiAP's LEDs. Valid values: `enable` `disable` .
* `override_login_passwd_change` - Enable to override the WTP profile login-password (administrator password) setting. Valid values: `enable` `disable` .
* `override_split_tunnel` - Enable/disable overriding the WTP profile split tunneling setting. Valid values: `enable` `disable` .
* `override_wan_port_mode` - Enable/disable overriding the wan-port-mode in the WTP profile. Valid values: `enable` `disable` .
* `region` - Region name WTP is associated with. This attribute must reference one of the following datasources: `wireless-controller.region.name` .
* `region_x` - Relative horizontal region coordinate (between 0 and 1).
* `region_y` - Relative vertical region coordinate (between 0 and 1).
* `split_tunneling_acl_local_ap_subnet` - Enable/disable automatically adding local subnetwork of FortiAP to split-tunneling ACL (default = disable). Valid values: `enable` `disable` .
* `split_tunneling_acl_path` - Split tunneling ACL path is local/tunnel. Valid values: `tunnel` `local` .
* `tun_mtu_downlink` - The MTU of downlink CAPWAP tunnel (576 - 1500 bytes or 0; 0 means the local MTU of FortiAP; default = 0).
* `tun_mtu_uplink` - The maximum transmission unit (MTU) of uplink CAPWAP tunnel (576 - 1500 bytes or 0; 0 means the local MTU of FortiAP; default = 0).
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `wan_port_mode` - Enable/disable using the FortiAP WAN port as a LAN port. Valid values: `wan-lan` `wan-only` .
* `wtp_id` - WTP ID.
* `wtp_mode` - WTP, AP, or FortiAP operating mode; normal (by default) or remote. A tunnel mode SSID can be assigned to an AP in normal mode but not remote mode, while a local-bridge mode SSID can be assigned to an AP in either normal mode or remote mode. Valid values: `normal` `remote` .
* `wtp_profile` - WTP profile name to apply to this WTP, AP or FortiAP. This attribute must reference one of the following datasources: `wireless-controller.wtp-profile.name` .
* `lan` - WTP LAN port mapping. The structure of `lan` block is documented below.

The `lan` block contains:

* `port_esl_mode` - ESL port mode. Valid values: `offline` `nat-to-wan` `bridge-to-wan` `bridge-to-ssid` .
* `port_esl_ssid` - Bridge ESL port to SSID. This attribute must reference one of the following datasources: `system.interface.name` .
* `port_mode` - LAN port mode. Valid values: `offline` `nat-to-wan` `bridge-to-wan` `bridge-to-ssid` .
* `port_ssid` - Bridge LAN port to SSID. This attribute must reference one of the following datasources: `system.interface.name` .
* `port1_mode` - LAN port 1 mode. Valid values: `offline` `nat-to-wan` `bridge-to-wan` `bridge-to-ssid` .
* `port1_ssid` - Bridge LAN port 1 to SSID. This attribute must reference one of the following datasources: `system.interface.name` .
* `port2_mode` - LAN port 2 mode. Valid values: `offline` `nat-to-wan` `bridge-to-wan` `bridge-to-ssid` .
* `port2_ssid` - Bridge LAN port 2 to SSID. This attribute must reference one of the following datasources: `system.interface.name` .
* `port3_mode` - LAN port 3 mode. Valid values: `offline` `nat-to-wan` `bridge-to-wan` `bridge-to-ssid` .
* `port3_ssid` - Bridge LAN port 3 to SSID. This attribute must reference one of the following datasources: `system.interface.name` .
* `port4_mode` - LAN port 4 mode. Valid values: `offline` `nat-to-wan` `bridge-to-wan` `bridge-to-ssid` .
* `port4_ssid` - Bridge LAN port 4 to SSID. This attribute must reference one of the following datasources: `system.interface.name` .
* `port5_mode` - LAN port 5 mode. Valid values: `offline` `nat-to-wan` `bridge-to-wan` `bridge-to-ssid` .
* `port5_ssid` - Bridge LAN port 5 to SSID. This attribute must reference one of the following datasources: `system.interface.name` .
* `port6_mode` - LAN port 6 mode. Valid values: `offline` `nat-to-wan` `bridge-to-wan` `bridge-to-ssid` .
* `port6_ssid` - Bridge LAN port 6 to SSID. This attribute must reference one of the following datasources: `system.interface.name` .
* `port7_mode` - LAN port 7 mode. Valid values: `offline` `nat-to-wan` `bridge-to-wan` `bridge-to-ssid` .
* `port7_ssid` - Bridge LAN port 7 to SSID. This attribute must reference one of the following datasources: `system.interface.name` .
* `port8_mode` - LAN port 8 mode. Valid values: `offline` `nat-to-wan` `bridge-to-wan` `bridge-to-ssid` .
* `port8_ssid` - Bridge LAN port 8 to SSID. This attribute must reference one of the following datasources: `system.interface.name` .
* `radio_1` - Configuration options for radio 1. The structure of `radio_1` block is documented below.

The `radio_1` block contains:

* `auto_power_high` - The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable). Valid values: `enable` `disable` .
* `auto_power_low` - The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - The target of automatic transmit power adjustment in dBm. (-95 to -20, default = -70).
* `band` - WiFi band that Radio 1 operates on. Valid values: `802.11a` `802.11b` `802.11g` `802.11n` `802.11n-5G` `802.11ac` `802.11ax-5G` `802.11ax` `802.11ac-2G` `802.11n,g-only` `802.11g-only` `802.11n-only` `802.11n-5G-only` `802.11ac,n-only` `802.11ac-only` `802.11ax,ac-only` `802.11ax,ac,n-only` `802.11ax-5G-only` `802.11ax,n-only` `802.11ax,n,g-only` `802.11ax-only` .
* `drma_manual_mode` - Radio mode to be used for DRMA manual mode (default = ncf). Valid values: `ap` `monitor` `ncf` `ncf-peek` .
* `override_analysis` - Enable to override the WTP profile spectrum analysis configuration. Valid values: `enable` `disable` .
* `override_band` - Enable to override the WTP profile band setting. Valid values: `enable` `disable` .
* `override_channel` - Enable to override WTP profile channel settings. Valid values: `enable` `disable` .
* `override_txpower` - Enable to override the WTP profile power level configuration. Valid values: `enable` `disable` .
* `override_vaps` - Enable to override WTP profile Virtual Access Point (VAP) settings. Valid values: `enable` `disable` .
* `power_level` - Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).
* `power_mode` - Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities. Valid values: `dBm` `percentage` .
* `power_value` - Radio EIRP power in dBm (1 - 33, default = 27).
* `spectrum_analysis` - Enable/disable spectrum analysis to find interference that would negatively impact wireless performance. Valid values: `enable` `scan-only` `disable` .
* `vap_all` - Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs). Valid values: `tunnel` `bridge` `manual` .
* `channel` - Selected list of wireless radio channels. The structure of `channel` block is documented below.

The `channel` block contains:

* `chan` - Channel number.
* `vaps` - Manually selected list of Virtual Access Points (VAPs). The structure of `vaps` block is documented below.

The `vaps` block contains:

* `name` - Virtual Access Point (VAP) name. This attribute must reference one of the following datasources: `wireless-controller.vap-group.name` `system.interface.name` .
* `radio_2` - Configuration options for radio 2. The structure of `radio_2` block is documented below.

The `radio_2` block contains:

* `auto_power_high` - The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable). Valid values: `enable` `disable` .
* `auto_power_low` - The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - The target of automatic transmit power adjustment in dBm. (-95 to -20, default = -70).
* `band` - WiFi band that Radio 2 operates on. Valid values: `802.11a` `802.11b` `802.11g` `802.11n` `802.11n-5G` `802.11ac` `802.11ax-5G` `802.11ax` `802.11ac-2G` `802.11n,g-only` `802.11g-only` `802.11n-only` `802.11n-5G-only` `802.11ac,n-only` `802.11ac-only` `802.11ax,ac-only` `802.11ax,ac,n-only` `802.11ax-5G-only` `802.11ax,n-only` `802.11ax,n,g-only` `802.11ax-only` .
* `drma_manual_mode` - Radio mode to be used for DRMA manual mode (default = ncf). Valid values: `ap` `monitor` `ncf` `ncf-peek` .
* `override_analysis` - Enable to override the WTP profile spectrum analysis configuration. Valid values: `enable` `disable` .
* `override_band` - Enable to override the WTP profile band setting. Valid values: `enable` `disable` .
* `override_channel` - Enable to override WTP profile channel settings. Valid values: `enable` `disable` .
* `override_txpower` - Enable to override the WTP profile power level configuration. Valid values: `enable` `disable` .
* `override_vaps` - Enable to override WTP profile Virtual Access Point (VAP) settings. Valid values: `enable` `disable` .
* `power_level` - Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).
* `power_mode` - Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities. Valid values: `dBm` `percentage` .
* `power_value` - Radio EIRP power in dBm (1 - 33, default = 27).
* `spectrum_analysis` - Enable/disable spectrum analysis to find interference that would negatively impact wireless performance. Valid values: `enable` `scan-only` `disable` .
* `vap_all` - Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs). Valid values: `tunnel` `bridge` `manual` .
* `channel` - Selected list of wireless radio channels. The structure of `channel` block is documented below.

The `channel` block contains:

* `chan` - Channel number.
* `vaps` - Manually selected list of Virtual Access Points (VAPs). The structure of `vaps` block is documented below.

The `vaps` block contains:

* `name` - Virtual Access Point (VAP) name. This attribute must reference one of the following datasources: `wireless-controller.vap-group.name` `system.interface.name` .
* `radio_3` - Configuration options for radio 3. The structure of `radio_3` block is documented below.

The `radio_3` block contains:

* `auto_power_high` - The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable). Valid values: `enable` `disable` .
* `auto_power_low` - The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - The target of automatic transmit power adjustment in dBm. (-95 to -20, default = -70).
* `band` - WiFi band that Radio 3 operates on. Valid values: `802.11a` `802.11b` `802.11g` `802.11n` `802.11n-5G` `802.11ac` `802.11ax-5G` `802.11ax` `802.11ac-2G` `802.11n,g-only` `802.11g-only` `802.11n-only` `802.11n-5G-only` `802.11ac,n-only` `802.11ac-only` `802.11ax,ac-only` `802.11ax,ac,n-only` `802.11ax-5G-only` `802.11ax,n-only` `802.11ax,n,g-only` `802.11ax-only` .
* `drma_manual_mode` - Radio mode to be used for DRMA manual mode (default = ncf). Valid values: `ap` `monitor` `ncf` `ncf-peek` .
* `override_analysis` - Enable to override the WTP profile spectrum analysis configuration. Valid values: `enable` `disable` .
* `override_band` - Enable to override the WTP profile band setting. Valid values: `enable` `disable` .
* `override_channel` - Enable to override WTP profile channel settings. Valid values: `enable` `disable` .
* `override_txpower` - Enable to override the WTP profile power level configuration. Valid values: `enable` `disable` .
* `override_vaps` - Enable to override WTP profile Virtual Access Point (VAP) settings. Valid values: `enable` `disable` .
* `power_level` - Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).
* `power_mode` - Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities. Valid values: `dBm` `percentage` .
* `power_value` - Radio EIRP power in dBm (1 - 33, default = 27).
* `spectrum_analysis` - Enable/disable spectrum analysis to find interference that would negatively impact wireless performance. Valid values: `enable` `scan-only` `disable` .
* `vap_all` - Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs). Valid values: `tunnel` `bridge` `manual` .
* `channel` - Selected list of wireless radio channels. The structure of `channel` block is documented below.

The `channel` block contains:

* `chan` - Channel number.
* `vaps` - Manually selected list of Virtual Access Points (VAPs). The structure of `vaps` block is documented below.

The `vaps` block contains:

* `name` - Virtual Access Point (VAP) name. This attribute must reference one of the following datasources: `wireless-controller.vap-group.name` `system.interface.name` .
* `radio_4` - Configuration options for radio 4. The structure of `radio_4` block is documented below.

The `radio_4` block contains:

* `auto_power_high` - The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable). Valid values: `enable` `disable` .
* `auto_power_low` - The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - The target of automatic transmit power adjustment in dBm. (-95 to -20, default = -70).
* `band` - WiFi band that Radio 4 operates on. Valid values: `802.11a` `802.11b` `802.11g` `802.11n` `802.11n-5G` `802.11ac` `802.11ax-5G` `802.11ax` `802.11ac-2G` `802.11n,g-only` `802.11g-only` `802.11n-only` `802.11n-5G-only` `802.11ac,n-only` `802.11ac-only` `802.11ax,ac-only` `802.11ax,ac,n-only` `802.11ax-5G-only` `802.11ax,n-only` `802.11ax,n,g-only` `802.11ax-only` .
* `drma_manual_mode` - Radio mode to be used for DRMA manual mode (default = ncf). Valid values: `ap` `monitor` `ncf` `ncf-peek` .
* `override_analysis` - Enable to override the WTP profile spectrum analysis configuration. Valid values: `enable` `disable` .
* `override_band` - Enable to override the WTP profile band setting. Valid values: `enable` `disable` .
* `override_channel` - Enable to override WTP profile channel settings. Valid values: `enable` `disable` .
* `override_txpower` - Enable to override the WTP profile power level configuration. Valid values: `enable` `disable` .
* `override_vaps` - Enable to override WTP profile Virtual Access Point (VAP) settings. Valid values: `enable` `disable` .
* `power_level` - Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).
* `power_mode` - Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities. Valid values: `dBm` `percentage` .
* `power_value` - Radio EIRP power in dBm (1 - 33, default = 27).
* `spectrum_analysis` - Enable/disable spectrum analysis to find interference that would negatively impact wireless performance. Valid values: `enable` `scan-only` `disable` .
* `vap_all` - Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs). Valid values: `tunnel` `bridge` `manual` .
* `channel` - Selected list of wireless radio channels. The structure of `channel` block is documented below.

The `channel` block contains:

* `chan` - Channel number.
* `vaps` - Manually selected list of Virtual Access Points (VAPs). The structure of `vaps` block is documented below.

The `vaps` block contains:

* `name` - Virtual Access Point (VAP) name. This attribute must reference one of the following datasources: `wireless-controller.vap-group.name` `system.interface.name` .
* `split_tunneling_acl` - Split tunneling ACL filter list. The structure of `split_tunneling_acl` block is documented below.

The `split_tunneling_acl` block contains:

* `dest_ip` - Destination IP and mask for the split-tunneling subnet.
* `id` - ID.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_wtp can be imported using:
```sh
terraform import fortios_wirelesscontroller_wtp.labelname {{mkey}}
```
