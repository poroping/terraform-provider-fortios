---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_mobiletunnel"
description: |-
  Configure Mobile tunnels, an implementation of Network Mobility (NEMO) extensions for Mobile IPv4 RFC5177.
---

## fortios_system_mobiletunnel
Configure Mobile tunnels, an implementation of Network Mobility (NEMO) extensions for Mobile IPv4 RFC5177.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `hash_algorithm` - Hash Algorithm (Keyed MD5). Valid values: `hmac-md5` .
* `home_address` - Home IP address (Format: xxx.xxx.xxx.xxx).
* `home_agent` - IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).
* `lifetime` - NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).
* `n_mhae_key` - NEMO authentication key.
* `n_mhae_key_type` - NEMO authentication key type (ASCII or base64). Valid values: `ascii` `base64` .
* `n_mhae_spi` - NEMO authentication SPI (default: 256).
* `name` - Tunnel name.
* `reg_interval` - NMMO HA registration interval (5 - 300, default = 5).
* `reg_retry` - Maximum number of NMMO HA registration retries (1 to 30, default = 3).
* `renew_interval` - Time before lifetime expiration to send NMMO HA re-registration (5 - 60, default = 60).
* `roaming_interface` - Select the associated interface name from available options. This attribute must reference one of the following datasources: `system.interface.name` .
* `status` - Enable/disable this mobile tunnel. Valid values: `disable` `enable` .
* `tunnel_mode` - NEMO tunnel mode (GRE tunnel). Valid values: `gre` .
* `network` - NEMO network configuration. The structure of `network` block is documented below.

The `network` block contains:

* `id` - Network entry ID.
* `interface` - Select the associated interface name from available options. This attribute must reference one of the following datasources: `system.interface.name` .
* `prefix` - Class IP and Netmask with correction (Format:xxx.xxx.xxx.xxx xxx.xxx.xxx.xxx or xxx.xxx.xxx.xxx/x).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_mobiletunnel can be imported using:
```sh
terraform import fortios_system_mobiletunnel.labelname {{mkey}}
```
