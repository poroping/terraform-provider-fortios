---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemdhcp_server"
description: |-
  Get information on a fortios Configure DHCP servers.
---

# Data Source: fortios_systemdhcp_server
Use this data source to get information on a fortios Configure DHCP servers.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `auto_configuration` - Enable/disable auto configuration.
* `auto_managed_status` - Enable/disable use of this DHCP server once this interface has been assigned an IP address from FortiIPAM.
* `conflicted_ip_timeout` - Time in seconds to wait after a conflicted IP address is removed from the DHCP range before it can be reused.
* `ddns_auth` - DDNS authentication mode.
* `ddns_key` - DDNS update key (base 64 encoding).
* `ddns_keyname` - DDNS update key name.
* `ddns_server_ip` - DDNS server IP.
* `ddns_ttl` - TTL.
* `ddns_update` - Enable/disable DDNS update for DHCP.
* `ddns_update_override` - Enable/disable DDNS update override for DHCP.
* `ddns_zone` - Zone of your domain name (ex. DDNS.com).
* `default_gateway` - Default gateway IP address assigned by the DHCP server.
* `dhcp_settings_from_fortiipam` - Enable/disable populating of DHCP server settings from FortiIPAM.
* `dns_server1` - DNS server 1.
* `dns_server2` - DNS server 2.
* `dns_server3` - DNS server 3.
* `dns_server4` - DNS server 4.
* `dns_service` - Options for assigning DNS servers to DHCP clients.
* `domain` - Domain name suffix for the IP addresses that the DHCP server assigns to clients.
* `filename` - Name of the boot file on the TFTP server.
* `forticlient_on_net_status` - Enable/disable FortiClient-On-Net service for this DHCP server.
* `id` - ID.
* `interface` - DHCP server can assign IP configurations to clients connected to this interface.
* `ip_mode` - Method used to assign client IP.
* `ipsec_lease_hold` - DHCP over IPsec leases expire this many seconds after tunnel down (0 to disable forced-expiry).
* `lease_time` - Lease time in seconds, 0 means unlimited.
* `mac_acl_default_action` - MAC access control default action (allow or block assigning IP settings).
* `netmask` - Netmask assigned by the DHCP server.
* `next_server` - IP address of a server (for example, a TFTP sever) that DHCP clients can download a boot file from.
* `ntp_server1` - NTP server 1.
* `ntp_server2` - NTP server 2.
* `ntp_server3` - NTP server 3.
* `ntp_service` - Options for assigning Network Time Protocol (NTP) servers to DHCP clients.
* `server_type` - DHCP server can be a normal DHCP server or an IPsec DHCP server.
* `status` - Enable/disable this DHCP configuration.
* `timezone` - Select the time zone to be assigned to DHCP clients.
* `timezone_option` - Options for the DHCP server to set the client's time zone.
* `vci_match` - Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served.
* `wifi_ac_service` - Options for assigning WiFi access controllers to DHCP clients.
* `wifi_ac1` - WiFi Access Controller 1 IP address (DHCP option 138, RFC 5417).
* `wifi_ac2` - WiFi Access Controller 2 IP address (DHCP option 138, RFC 5417).
* `wifi_ac3` - WiFi Access Controller 3 IP address (DHCP option 138, RFC 5417).
* `wins_server1` - WINS server 1.
* `wins_server2` - WINS server 2.
* `exclude_range` - Exclude one or more ranges of IP addresses from being assigned to clients.The structure of `exclude_range` block is documented below.

The `exclude_range` block contains:

* `end_ip` - End of IP range.
* `id` - ID.
* `lease_time` - Lease time in seconds, 0 means default lease time.
* `start_ip` - Start of IP range.
* `uci_match` - Enable/disable user class identifier (UCI) matching. When enabled only DHCP requests with a matching UCI are served with this range.
* `vci_match` - Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this range.
* `uci_string` - One or more UCI strings in quotes separated by spaces.The structure of `uci_string` block is documented below.

The `uci_string` block contains:

* `uci_string` - UCI strings.
* `vci_string` - One or more VCI strings in quotes separated by spaces.The structure of `vci_string` block is documented below.

The `vci_string` block contains:

* `vci_string` - VCI strings.
* `ip_range` - DHCP IP range configuration.The structure of `ip_range` block is documented below.

The `ip_range` block contains:

* `end_ip` - End of IP range.
* `id` - ID.
* `lease_time` - Lease time in seconds, 0 means default lease time.
* `start_ip` - Start of IP range.
* `uci_match` - Enable/disable user class identifier (UCI) matching. When enabled only DHCP requests with a matching UCI are served with this range.
* `vci_match` - Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this range.
* `uci_string` - One or more UCI strings in quotes separated by spaces.The structure of `uci_string` block is documented below.

The `uci_string` block contains:

* `uci_string` - UCI strings.
* `vci_string` - One or more VCI strings in quotes separated by spaces.The structure of `vci_string` block is documented below.

The `vci_string` block contains:

* `vci_string` - VCI strings.
* `options` - DHCP options.The structure of `options` block is documented below.

The `options` block contains:

* `code` - DHCP option code.
* `id` - ID.
* `ip` - DHCP option IPs.
* `type` - DHCP option type.
* `uci_match` - Enable/disable user class identifier (UCI) matching. When enabled only DHCP requests with a matching UCI are served with this option.
* `value` - DHCP option value.
* `vci_match` - Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this option.
* `uci_string` - One or more UCI strings in quotes separated by spaces.The structure of `uci_string` block is documented below.

The `uci_string` block contains:

* `uci_string` - UCI strings.
* `vci_string` - One or more VCI strings in quotes separated by spaces.The structure of `vci_string` block is documented below.

The `vci_string` block contains:

* `vci_string` - VCI strings.
* `reserved_address` - Options for the DHCP server to assign IP settings to specific MAC addresses.The structure of `reserved_address` block is documented below.

The `reserved_address` block contains:

* `action` - Options for the DHCP server to configure the client with the reserved MAC address.
* `circuit_id` - Option 82 circuit-ID of the client that will get the reserved IP address.
* `circuit_id_type` - DHCP option type.
* `description` - Description.
* `id` - ID.
* `ip` - IP address to be reserved for the MAC address.
* `mac` - MAC address of the client that will get the reserved IP address.
* `remote_id` - Option 82 remote-ID of the client that will get the reserved IP address.
* `remote_id_type` - DHCP option type.
* `type` - DHCP reserved-address type.
* `tftp_server` - One or more hostnames or IP addresses of the TFTP servers in quotes separated by spaces.The structure of `tftp_server` block is documented below.

The `tftp_server` block contains:

* `tftp_server` - TFTP server.
* `vci_string` - One or more VCI strings in quotes separated by spaces.The structure of `vci_string` block is documented below.

The `vci_string` block contains:

* `vci_string` - VCI strings.
