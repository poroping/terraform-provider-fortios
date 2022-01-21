---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_mobiletunnel"
description: |-
  Get information on a fortios Configure Mobile tunnels, an implementation of Network Mobility (NEMO) extensions for Mobile IPv4 RFC5177.
---

# Data Source: fortios_system_mobiletunnel
Use this data source to get information on a fortios Configure Mobile tunnels, an implementation of Network Mobility (NEMO) extensions for Mobile IPv4 RFC5177.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Tunnel name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `hash_algorithm` - Hash Algorithm (Keyed MD5).
* `home_address` - Home IP address (Format: xxx.xxx.xxx.xxx).
* `home_agent` - IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).
* `lifetime` - NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).
* `n_mhae_key` - NEMO authentication key.
* `n_mhae_key_type` - NEMO authentication key type (ascii or base64).
* `n_mhae_spi` - NEMO authentication SPI (default: 256).
* `name` - Tunnel name.
* `reg_interval` - NMMO HA registration interval (5 - 300, default = 5).
* `reg_retry` - Maximum number of NMMO HA registration retries (1 to 30, default = 3).
* `renew_interval` - Time before lifetime expiraton to send NMMO HA re-registration (5 - 60, default = 60).
* `roaming_interface` - Select the associated interface name from available options.
* `status` - Enable/disable this mobile tunnel.
* `tunnel_mode` - NEMO tunnnel mode (GRE tunnel).
* `network` - NEMO network configuration.The structure of `network` block is documented below.

The `network` block contains:

* `id` - Network entry ID.
* `interface` - Select the associated interface name from available options.
* `prefix` - Class IP and Netmask with correction (Format:xxx.xxx.xxx.xxx xxx.xxx.xxx.xxx or xxx.xxx.xxx.xxx/x).
