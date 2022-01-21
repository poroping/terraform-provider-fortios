---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_h2qpconncapability"
description: |-
  Configure connection capability.
---

## fortios_wirelesscontrollerhotspot20_h2qpconncapability
Configure connection capability.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `esp_port` - Set ESP port service (used by IPsec VPNs) status. Valid values: `closed` `open` `unknown` .
* `ftp_port` - Set FTP port service status. Valid values: `closed` `open` `unknown` .
* `http_port` - Set HTTP port service status. Valid values: `closed` `open` `unknown` .
* `icmp_port` - Set ICMP port service status. Valid values: `closed` `open` `unknown` .
* `ikev2_port` - Set IKEv2 port service for IPsec VPN status. Valid values: `closed` `open` `unknown` .
* `ikev2_xx_port` - Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed` `open` `unknown` .
* `name` - Connection capability name.
* `pptp_vpn_port` - Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed` `open` `unknown` .
* `ssh_port` - Set SSH port service status. Valid values: `closed` `open` `unknown` .
* `tls_port` - Set TLS VPN (HTTPS) port service status. Valid values: `closed` `open` `unknown` .
* `voip_tcp_port` - Set VoIP TCP port service status. Valid values: `closed` `open` `unknown` .
* `voip_udp_port` - Set VoIP UDP port service status. Valid values: `closed` `open` `unknown` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontrollerhotspot20_h2qpconncapability can be imported using:
```sh
terraform import fortios_wirelesscontrollerhotspot20_h2qpconncapability.labelname {{mkey}}
```
