---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fortiguard"
description: |-
  Configure FortiGuard services.
---

## fortios_system_fortiguard
Configure FortiGuard services.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `antispam_cache` - Enable/disable FortiGuard antispam request caching. Uses a small amount of memory but improves performance. Valid values: `enable` `disable` .
* `antispam_cache_mpercent` - Maximum percentage of FortiGate memory the antispam cache is allowed to use (1 - 15).
* `antispam_cache_ttl` - Time-to-live for antispam cache entries in seconds (300 - 86400). Lower times reduce the cache size. Higher times may improve performance since the cache will have more entries.
* `antispam_expiration` - Expiration date of the FortiGuard antispam contract.
* `antispam_force_off` - Enable/disable turning off the FortiGuard antispam service. Valid values: `enable` `disable` .
* `antispam_license` - Interval of time between license checks for the FortiGuard antispam contract.
* `antispam_timeout` - Antispam query time out (1 - 30 sec, default = 7).
* `anycast_sdns_server_ip` - IP address of the FortiGuard anycast DNS rating server.
* `anycast_sdns_server_port` - Port to connect to on the FortiGuard anycast DNS rating server.
* `auto_join_forticloud` - Automatically connect to and login to FortiCloud. Valid values: `enable` `disable` .
* `ddns_server_ip` - IP address of the FortiDDNS server.
* `ddns_server_ip6` - IPv6 address of the FortiDDNS server.
* `ddns_server_port` - Port used to communicate with FortiDDNS servers.
* `fortiguard_anycast` - Enable/disable use of FortiGuard's Anycast network. Valid values: `enable` `disable` .
* `fortiguard_anycast_source` - Configure which of Fortinet's servers to provide FortiGuard services in FortiGuard's anycast network. Default is Fortinet. Valid values: `fortinet` `aws` `debug` .
* `interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `load_balance_servers` - Number of servers to alternate between as first FortiGuard option.
* `outbreak_prevention_cache` - Enable/disable FortiGuard Virus Outbreak Prevention cache. Valid values: `enable` `disable` .
* `outbreak_prevention_cache_mpercent` - Maximum percent of memory FortiGuard Virus Outbreak Prevention cache can use (1 - 15%, default = 2).
* `outbreak_prevention_cache_ttl` - Time-to-live for FortiGuard Virus Outbreak Prevention cache entries (300 - 86400 sec, default = 300).
* `outbreak_prevention_expiration` - Expiration date of FortiGuard Virus Outbreak Prevention contract.
* `outbreak_prevention_force_off` - Turn off FortiGuard Virus Outbreak Prevention service. Valid values: `enable` `disable` .
* `outbreak_prevention_license` - Interval of time between license checks for FortiGuard Virus Outbreak Prevention contract.
* `outbreak_prevention_timeout` - FortiGuard Virus Outbreak Prevention time out (1 - 30 sec, default = 7).
* `persistent_connection` - Enable/disable use of persistent connection to receive update notification from FortiGuard. Valid values: `enable` `disable` .
* `port` - Port used to communicate with the FortiGuard servers. Valid values: `8888` `53` `80` `443` .
* `protocol` - Protocol used to communicate with the FortiGuard servers. Valid values: `udp` `http` `https` .
* `proxy_password` - Proxy user password.
* `proxy_server_ip` - IP address of the proxy server.
* `proxy_server_port` - Port used to communicate with the proxy server.
* `proxy_username` - Proxy user name.
* `sandbox_region` - Cloud sandbox region.
* `sdns_options` - Customization options for the FortiGuard DNS service. Valid values: `include-question-section` .
* `sdns_server_ip` - IP address of the FortiGuard DNS rating server.
* `sdns_server_port` - Port to connect to on the FortiGuard DNS rating server.
* `source_ip` - Source IPv4 address used to communicate with FortiGuard.
* `source_ip6` - Source IPv6 address used to communicate with FortiGuard.
* `update_build_proxy` - Enable/disable proxy dictionary rebuild. Valid values: `enable` `disable` .
* `update_extdb` - Enable/disable external resource update. Valid values: `enable` `disable` .
* `update_ffdb` - Enable/disable Internet Service Database update. Valid values: `enable` `disable` .
* `update_server_location` - Location from which to receive FortiGuard updates. Valid values: `automatic` `usa` `eu` .
* `update_uwdb` - Enable/disable allowlist update. Valid values: `enable` `disable` .
* `vdom` - FortiGuard Service virtual domain name. This attribute must reference one of the following datasources: `system.vdom.name` .
* `videofilter_expiration` - Expiration date of the FortiGuard video filter contract.
* `videofilter_license` - Interval of time between license checks for the FortiGuard video filter contract.
* `webfilter_cache` - Enable/disable FortiGuard web filter caching. Valid values: `enable` `disable` .
* `webfilter_cache_ttl` - Time-to-live for web filter cache entries in seconds (300 - 86400).
* `webfilter_expiration` - Expiration date of the FortiGuard web filter contract.
* `webfilter_force_off` - Enable/disable turning off the FortiGuard web filtering service. Valid values: `enable` `disable` .
* `webfilter_license` - Interval of time between license checks for the FortiGuard web filter contract.
* `webfilter_timeout` - Web filter query time out (1 - 30 sec, default = 7).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_fortiguard can be imported using:
```sh
terraform import fortios_system_fortiguard.labelname {{mkey}}
```
