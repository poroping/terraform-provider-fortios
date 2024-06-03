---
subcategory: "FortiGate Wanopt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_profile"
description: |-
  Configure WAN optimization profiles.
---

## fortios_wanopt_profile
Configure WAN optimization profiles.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `auth_group` - Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
* `comments` - Comment.
* `name` - Profile name.
* `transparent` - Enable/disable transparent mode. Valid values: `enable` `disable` .
* `cifs` - Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.

The `cifs` block contains:

* `byte_caching` - Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `enable` `disable` .
* `log_traffic` - Enable/disable logging. Valid values: `enable` `disable` .
* `port` - Single port number or port number range for CIFS. Only packets with a destination port number that matches this port number or range are accepted by this profile.
* `prefer_chunking` - Select dynamic or fixed-size data chunking for WAN Optimization. Valid values: `dynamic` `fix` .
* `protocol_opt` - Select protocol specific optimization or generic TCP optimization. Valid values: `protocol` `tcp` .
* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `enable` `disable` .
* `status` - Enable/disable WAN Optimization. Valid values: `enable` `disable` .
* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `shared` `express-shared` `private` .
* `ftp` - Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.

The `ftp` block contains:

* `byte_caching` - Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `enable` `disable` .
* `log_traffic` - Enable/disable logging. Valid values: `enable` `disable` .
* `port` - Single port number or port number range for FTP. Only packets with a destination port number that matches this port number or range are accepted by this profile.
* `prefer_chunking` - Select dynamic or fixed-size data chunking for WAN Optimization. Valid values: `dynamic` `fix` .
* `protocol_opt` - Select protocol specific optimization or generic TCP optimization. Valid values: `protocol` `tcp` .
* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `enable` `disable` .
* `ssl` - Enable/disable SSL/TLS offloading (hardware acceleration) for traffic in this tunnel. Valid values: `enable` `disable` .
* `status` - Enable/disable WAN Optimization. Valid values: `enable` `disable` .
* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `shared` `express-shared` `private` .
* `http` - Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.

The `http` block contains:

* `byte_caching` - Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `enable` `disable` .
* `log_traffic` - Enable/disable logging. Valid values: `enable` `disable` .
* `port` - Single port number or port number range for HTTP. Only packets with a destination port number that matches this port number or range are accepted by this profile.
* `prefer_chunking` - Select dynamic or fixed-size data chunking for WAN Optimization. Valid values: `dynamic` `fix` .
* `protocol_opt` - Select protocol specific optimization or generic TCP optimization. Valid values: `protocol` `tcp` .
* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `enable` `disable` .
* `ssl` - Enable/disable SSL/TLS offloading (hardware acceleration) for traffic in this tunnel. Valid values: `enable` `disable` .
* `ssl_port` - Port on which to expect HTTPS traffic for SSL/TLS offloading.
* `status` - Enable/disable WAN Optimization. Valid values: `enable` `disable` .
* `tunnel_non_http` - Configure how to process non-HTTP traffic when a profile configured for HTTP traffic accepts a non-HTTP session. Can occur if an application sends non-HTTP traffic using an HTTP destination port. Valid values: `enable` `disable` .
* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `shared` `express-shared` `private` .
* `unknown_http_version` - How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1. Valid values: `reject` `tunnel` `best-effort` .
* `mapi` - Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.

The `mapi` block contains:

* `byte_caching` - Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `enable` `disable` .
* `log_traffic` - Enable/disable logging. Valid values: `enable` `disable` .
* `port` - Single port number or port number range for MAPI. Only packets with a destination port number that matches this port number or range are accepted by this profile.
* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `enable` `disable` .
* `status` - Enable/disable WAN Optimization. Valid values: `enable` `disable` .
* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `shared` `express-shared` `private` .
* `tcp` - Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.

The `tcp` block contains:

* `byte_caching` - Enable/disable byte-caching. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `enable` `disable` .
* `byte_caching_opt` - Select whether TCP byte-caching uses system memory only or both memory and disk space. Valid values: `mem-only` `mem-disk` .
* `log_traffic` - Enable/disable logging. Valid values: `enable` `disable` .
* `port` - Port numbers or port number ranges for TCP. Only packets with a destination port number that matches this port number or range are accepted by this profile.
* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `enable` `disable` .
* `ssl` - Enable/disable SSL/TLS offloading (hardware acceleration) for traffic in this tunnel. Valid values: `enable` `disable` .
* `ssl_port` - Port numbers or port number ranges on which to expect HTTPS traffic for SSL/TLS offloading.
* `status` - Enable/disable WAN Optimization. Valid values: `enable` `disable` .
* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `shared` `express-shared` `private` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wanopt_profile can be imported using:
```sh
terraform import fortios_wanopt_profile.labelname {{mkey}}
```
