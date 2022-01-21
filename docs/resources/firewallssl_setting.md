---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallssl_setting"
description: |-
  SSL proxy settings.
---

## fortios_firewallssl_setting
SSL proxy settings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `abbreviate_handshake` - Enable/disable use of SSL abbreviated handshake. Valid values: `enable` `disable` .
* `cert_cache_capacity` - Maximum capacity of the host certificate cache (0 - 500, default = 200).
* `cert_cache_timeout` - Time limit to keep certificate cache (1 - 120 min, default = 10).
* `kxp_queue_threshold` - Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
* `no_matching_cipher_action` - Bypass or drop the connection when no matching cipher is found. Valid values: `bypass` `drop` .
* `proxy_connect_timeout` - Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
* `session_cache_capacity` - Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
* `session_cache_timeout` - Time limit to keep SSL session state (1 - 60 min, default = 20).
* `ssl_dh_bits` - Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768` `1024` `1536` `2048` .
* `ssl_queue_threshold` - Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
* `ssl_send_empty_frags` - Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewallssl_setting can be imported using:
```sh
terraform import fortios_firewallssl_setting.labelname {{mkey}}
```
