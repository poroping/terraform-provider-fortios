---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_h2qpconncapability"
description: |-
  Get information on a fortios Configure connection capability.
---

# Data Source: fortios_wirelesscontrollerhotspot20_h2qpconncapability
Use this data source to get information on a fortios Configure connection capability.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Connection capability name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `esp_port` - Set ESP port service (used by IPsec VPNs) status.
* `ftp_port` - Set FTP port service status.
* `http_port` - Set HTTP port service status.
* `icmp_port` - Set ICMP port service status.
* `ikev2_port` - Set IKEv2 port service for IPsec VPN status.
* `ikev2_xx_port` - Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status.
* `name` - Connection capability name.
* `pptp_vpn_port` - Set Point to Point Tunneling Protocol (PPTP) VPN port service status.
* `ssh_port` - Set SSH port service status.
* `tls_port` - Set TLS VPN (HTTPS) port service status.
* `voip_tcp_port` - Set VoIP TCP port service status.
* `voip_udp_port` - Set VoIP UDP port service status.
