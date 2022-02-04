---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fortiguard"
description: |-
  Get information on a fortios Configure FortiGuard services.
---

# Data Source: fortios_system_fortiguard
Use this data source to get information on a fortios Configure FortiGuard services.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `antispam_cache` - Enable/disable FortiGuard antispam request caching. Uses a small amount of memory but improves performance.
* `antispam_cache_mpercent` - Maximum percentage of FortiGate memory the antispam cache is allowed to use (1 - 15).
* `antispam_cache_ttl` - Time-to-live for antispam cache entries in seconds (300 - 86400). Lower times reduce the cache size. Higher times may improve performance since the cache will have more entries.
* `antispam_expiration` - Expiration date of the FortiGuard antispam contract.
* `antispam_force_off` - Enable/disable turning off the FortiGuard antispam service.
* `antispam_license` - Interval of time between license checks for the FortiGuard antispam contract.
* `antispam_timeout` - Antispam query time out (1 - 30 sec, default = 7).
* `anycast_sdns_server_ip` - IP address of the FortiGuard anycast DNS rating server.
* `anycast_sdns_server_port` - Port to connect to on the FortiGuard anycast DNS rating server.
* `auto_join_forticloud` - Automatically connect to and login to FortiCloud.
* `ddns_server_ip` - IP address of the FortiDDNS server.
* `ddns_server_ip6` - IPv6 address of the FortiDDNS server.
* `ddns_server_port` - Port used to communicate with FortiDDNS servers.
* `fortiguard_anycast` - Enable/disable use of FortiGuard's Anycast network.
* `fortiguard_anycast_source` - Configure which of Fortinet's servers to provide FortiGuard services in FortiGuard's anycast network. Default is Fortinet.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `load_balance_servers` - Number of servers to alternate between as first FortiGuard option.
* `outbreak_prevention_cache` - Enable/disable FortiGuard Virus Outbreak Prevention cache.
* `outbreak_prevention_cache_mpercent` - Maximum percent of memory FortiGuard Virus Outbreak Prevention cache can use (1 - 15%, default = 2).
* `outbreak_prevention_cache_ttl` - Time-to-live for FortiGuard Virus Outbreak Prevention cache entries (300 - 86400 sec, default = 300).
* `outbreak_prevention_expiration` - Expiration date of FortiGuard Virus Outbreak Prevention contract.
* `outbreak_prevention_force_off` - Turn off FortiGuard Virus Outbreak Prevention service.
* `outbreak_prevention_license` - Interval of time between license checks for FortiGuard Virus Outbreak Prevention contract.
* `outbreak_prevention_timeout` - FortiGuard Virus Outbreak Prevention time out (1 - 30 sec, default = 7).
* `persistent_connection` - Enable/disable use of persistent connection to receive update notification from FortiGuard.
* `port` - Port used to communicate with the FortiGuard servers.
* `protocol` - Protocol used to communicate with the FortiGuard servers.
* `proxy_password` - Proxy user password.
* `proxy_server_ip` - IP address of the proxy server.
* `proxy_server_port` - Port used to communicate with the proxy server.
* `proxy_username` - Proxy user name.
* `sandbox_region` - Cloud sandbox region.
* `sdns_options` - Customization options for the FortiGuard DNS service.
* `sdns_server_ip` - IP address of the FortiGuard DNS rating server.
* `sdns_server_port` - Port to connect to on the FortiGuard DNS rating server.
* `source_ip` - Source IPv4 address used to communicate with FortiGuard.
* `source_ip6` - Source IPv6 address used to communicate with FortiGuard.
* `update_build_proxy` - Enable/disable proxy dictionary rebuild.
* `update_extdb` - Enable/disable external resource update.
* `update_ffdb` - Enable/disable Internet Service Database update.
* `update_server_location` - Location from which to receive FortiGuard updates.
* `update_uwdb` - Enable/disable allowlist update.
* `videofilter_expiration` - Expiration date of the FortiGuard video filter contract.
* `videofilter_license` - Interval of time between license checks for the FortiGuard video filter contract.
* `webfilter_cache` - Enable/disable FortiGuard web filter caching.
* `webfilter_cache_ttl` - Time-to-live for web filter cache entries in seconds (300 - 86400).
* `webfilter_expiration` - Expiration date of the FortiGuard web filter contract.
* `webfilter_force_off` - Enable/disable turning off the FortiGuard web filtering service.
* `webfilter_license` - Interval of time between license checks for the FortiGuard web filter contract.
* `webfilter_timeout` - Web filter query time out (1 - 30 sec, default = 7).
