---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_sslsshprofile"
description: |-
  Configure SSL/SSH protocol options.
---

## fortios_firewall_sslsshprofile
Configure SSL/SSH protocol options.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `allowlist` - Enable/disable exempting servers by FortiGuard allowlist. Valid values: `enable` `disable` .
* `block_blacklisted_certificates` - Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blacklist. Valid values: `disable` `enable` .
* `block_blocklisted_certificates` - Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blocklist. Valid values: `disable` `enable` .
* `caname` - CA certificate used by SSL Inspection. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `comment` - Optional comments.
* `mapi_over_https` - Enable/disable inspection of MAPI over HTTPS. Valid values: `enable` `disable` .
* `name` - Name.
* `rpc_over_https` - Enable/disable inspection of RPC over HTTPS. Valid values: `enable` `disable` .
* `server_cert_mode` - Re-sign or replace the server's certificate. Valid values: `re-sign` `replace` .
* `ssl_anomalies_log` - Enable/disable logging SSL anomalies. Valid values: `disable` `enable` .
* `ssl_anomaly_log` - Enable/disable logging of SSL anomalies. Valid values: `disable` `enable` .
* `ssl_exemption_ip_rating` - Enable/disable IP based URL rating. Valid values: `enable` `disable` .
* `ssl_exemption_log` - Enable/disable logging SSL exemptions. Valid values: `disable` `enable` .
* `ssl_exemptions_log` - Enable/disable logging SSL exemptions. Valid values: `disable` `enable` .
* `ssl_handshake_log` - Enable/disable logging of TLS handshakes. Valid values: `disable` `enable` .
* `ssl_negotiation_log` - Enable/disable logging SSL negotiation. Valid values: `disable` `enable` .
* `ssl_server_cert_log` - Enable/disable logging of server certificate information. Valid values: `disable` `enable` .
* `supported_alpn` - Configure ALPN option. Valid values: `http1-1` `http2` `all` `none` .
* `untrusted_caname` - Untrusted CA certificate used by SSL Inspection. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `use_ssl_server` - Enable/disable the use of SSL server table for SSL offloading. Valid values: `disable` `enable` .
* `whitelist` - Enable/disable exempting servers by FortiGuard whitelist. Valid values: `enable` `disable` .
* `dot` - Configure DNS over TLS options. The structure of `dot` block is documented below.

The `dot` block contains:

* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow` `block` `ignore` .
* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow` `block` `ignore` .
* `client_certificate` - Action based on received client certificate. Valid values: `bypass` `inspect` `block` .
* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow` `block` `ignore` .
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable` `disable` .
* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow` `block` `ignore` .
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `enable` `strict` `disable` .
* `status` - Configure protocol inspection status. Valid values: `disable` `deep-inspection` .
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow` `block` .
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow` `block` .
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported. Valid values: `allow` `block` .
* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA. Valid values: `allow` `block` `ignore` .
* `ftps` - Configure FTPS options. The structure of `ftps` block is documented below.

The `ftps` block contains:

* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow` `block` `ignore` .
* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow` `block` `ignore` .
* `client_cert_request` - Action based on client certificate request. Valid values: `bypass` `inspect` `block` .
* `client_certificate` - Action based on received client certificate. Valid values: `bypass` `inspect` `block` .
* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow` `block` `ignore` .
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate. Valid values: `allow` `block` .
* `min_allowed_ssl_version` - Minimum SSL version to be allowed. Valid values: `ssl-3.0` `tls-1.0` `tls-1.1` `tls-1.2` `tls-1.3` .
* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow` `block` `ignore` .
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `enable` `strict` `disable` .
* `status` - Configure protocol inspection status. Valid values: `disable` `deep-inspection` .
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported. Valid values: `bypass` `inspect` `block` .
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow` `block` .
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow` `block` .
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported. Valid values: `allow` `block` .
* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA. Valid values: `allow` `block` `ignore` .
* `https` - Configure HTTPS options. The structure of `https` block is documented below.

The `https` block contains:

* `cert_probe_failure` - Action based on certificate probe failure. Valid values: `allow` `block` .
* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow` `block` `ignore` .
* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow` `block` `ignore` .
* `client_cert_request` - Action based on client certificate request. Valid values: `bypass` `inspect` `block` .
* `client_certificate` - Action based on received client certificate. Valid values: `bypass` `inspect` `block` .
* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow` `block` `ignore` .
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate. Valid values: `allow` `block` .
* `min_allowed_ssl_version` - Minimum SSL version to be allowed. Valid values: `ssl-3.0` `tls-1.0` `tls-1.1` `tls-1.2` `tls-1.3` .
* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable` `disable` .
* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow` `block` `ignore` .
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `enable` `strict` `disable` .
* `status` - Configure protocol inspection status. Valid values: `disable` `certificate-inspection` `deep-inspection` .
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported. Valid values: `bypass` `inspect` `block` .
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow` `block` .
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow` `block` .
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported. Valid values: `allow` `block` .
* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA. Valid values: `allow` `block` `ignore` .
* `imaps` - Configure IMAPS options. The structure of `imaps` block is documented below.

The `imaps` block contains:

* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow` `block` `ignore` .
* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow` `block` `ignore` .
* `client_cert_request` - Action based on client certificate request. Valid values: `bypass` `inspect` `block` .
* `client_certificate` - Action based on received client certificate. Valid values: `bypass` `inspect` `block` .
* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow` `block` `ignore` .
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate. Valid values: `allow` `block` .
* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable` `disable` .
* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow` `block` `ignore` .
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `enable` `strict` `disable` .
* `status` - Configure protocol inspection status. Valid values: `disable` `deep-inspection` .
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported. Valid values: `bypass` `inspect` `block` .
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow` `block` .
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow` `block` .
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported. Valid values: `allow` `block` .
* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA. Valid values: `allow` `block` `ignore` .
* `pop3s` - Configure POP3S options. The structure of `pop3s` block is documented below.

The `pop3s` block contains:

* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow` `block` `ignore` .
* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow` `block` `ignore` .
* `client_cert_request` - Action based on client certificate request. Valid values: `bypass` `inspect` `block` .
* `client_certificate` - Action based on received client certificate. Valid values: `bypass` `inspect` `block` .
* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow` `block` `ignore` .
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate. Valid values: `allow` `block` .
* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable` `disable` .
* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow` `block` `ignore` .
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `enable` `strict` `disable` .
* `status` - Configure protocol inspection status. Valid values: `disable` `deep-inspection` .
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported. Valid values: `bypass` `inspect` `block` .
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow` `block` .
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow` `block` .
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported. Valid values: `allow` `block` .
* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA. Valid values: `allow` `block` `ignore` .
* `server_cert` - Certificate used by SSL Inspection to replace server certificate. The structure of `server_cert` block is documented below.

The `server_cert` block contains:

* `name` - Certificate list. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `smtps` - Configure SMTPS options. The structure of `smtps` block is documented below.

The `smtps` block contains:

* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow` `block` `ignore` .
* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow` `block` `ignore` .
* `client_cert_request` - Action based on client certificate request. Valid values: `bypass` `inspect` `block` .
* `client_certificate` - Action based on received client certificate. Valid values: `bypass` `inspect` `block` .
* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow` `block` `ignore` .
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate. Valid values: `allow` `block` .
* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable` `disable` .
* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow` `block` `ignore` .
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `enable` `strict` `disable` .
* `status` - Configure protocol inspection status. Valid values: `disable` `deep-inspection` .
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported. Valid values: `bypass` `inspect` `block` .
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow` `block` .
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow` `block` .
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported. Valid values: `allow` `block` .
* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA. Valid values: `allow` `block` `ignore` .
* `ssh` - Configure SSH options. The structure of `ssh` block is documented below.

The `ssh` block contains:

* `inspect_all` - Level of SSL inspection. Valid values: `disable` `deep-inspection` .
* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable` `disable` .
* `ssh_algorithm` - Relative strength of encryption algorithms accepted during negotiation. Valid values: `compatible` `high-encryption` .
* `ssh_tun_policy_check` - Enable/disable SSH tunnel policy check. Valid values: `disable` `enable` .
* `status` - Configure protocol inspection status. Valid values: `disable` `deep-inspection` .
* `unsupported_version` - Action based on SSH version being unsupported. Valid values: `bypass` `block` .
* `ssl` - Configure SSL options. The structure of `ssl` block is documented below.

The `ssl` block contains:

* `cert_probe_failure` - Action based on certificate probe failure. Valid values: `allow` `block` .
* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow` `block` `ignore` .
* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow` `block` `ignore` .
* `client_cert_request` - Action based on client certificate request. Valid values: `bypass` `inspect` `block` .
* `client_certificate` - Action based on received client certificate. Valid values: `bypass` `inspect` `block` .
* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow` `block` `ignore` .
* `inspect_all` - Level of SSL inspection. Valid values: `disable` `certificate-inspection` `deep-inspection` .
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate. Valid values: `allow` `block` .
* `min_allowed_ssl_version` - Minimum SSL version to be allowed. Valid values: `ssl-3.0` `tls-1.0` `tls-1.1` `tls-1.2` `tls-1.3` .
* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow` `block` `ignore` .
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `enable` `strict` `disable` .
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported. Valid values: `bypass` `inspect` `block` .
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow` `block` .
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow` `block` .
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported. Valid values: `allow` `block` .
* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA. Valid values: `allow` `block` `ignore` .
* `ssl_exempt` - Servers to exempt from SSL inspection. The structure of `ssl_exempt` block is documented below.

The `ssl_exempt` block contains:

* `address` - IPv4 address object. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `address6` - IPv6 address object. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `fortiguard_category` - FortiGuard category ID.
* `id` - ID number.
* `regex` - Exempt servers by regular expression.
* `type` - Type of address object (IPv4 or IPv6) or FortiGuard category. Valid values: `fortiguard-category` `address` `address6` `wildcard-fqdn` `regex` .
* `wildcard_fqdn` - Exempt servers by wildcard FQDN. This attribute must reference one of the following datasources: `firewall.wildcard-fqdn.custom.name` `firewall.wildcard-fqdn.group.name` .
* `ssl_server` - SSL server settings used for client certificate request. The structure of `ssl_server` block is documented below.

The `ssl_server` block contains:

* `ftps_client_cert_request` - Action based on client certificate request during the FTPS handshake. Valid values: `bypass` `inspect` `block` .
* `ftps_client_certificate` - Action based on received client certificate during the FTPS handshake. Valid values: `bypass` `inspect` `block` .
* `https_client_cert_request` - Action based on client certificate request during the HTTPS handshake. Valid values: `bypass` `inspect` `block` .
* `https_client_certificate` - Action based on received client certificate during the HTTPS handshake. Valid values: `bypass` `inspect` `block` .
* `id` - SSL server ID.
* `imaps_client_cert_request` - Action based on client certificate request during the IMAPS handshake. Valid values: `bypass` `inspect` `block` .
* `imaps_client_certificate` - Action based on received client certificate during the IMAPS handshake. Valid values: `bypass` `inspect` `block` .
* `ip` - IPv4 address of the SSL server.
* `pop3s_client_cert_request` - Action based on client certificate request during the POP3S handshake. Valid values: `bypass` `inspect` `block` .
* `pop3s_client_certificate` - Action based on received client certificate during the POP3S handshake. Valid values: `bypass` `inspect` `block` .
* `smtps_client_cert_request` - Action based on client certificate request during the SMTPS handshake. Valid values: `bypass` `inspect` `block` .
* `smtps_client_certificate` - Action based on received client certificate during the SMTPS handshake. Valid values: `bypass` `inspect` `block` .
* `ssl_other_client_cert_request` - Action based on client certificate request during an SSL protocol handshake. Valid values: `bypass` `inspect` `block` .
* `ssl_other_client_certificate` - Action based on received client certificate during an SSL protocol handshake. Valid values: `bypass` `inspect` `block` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_sslsshprofile can be imported using:
```sh
terraform import fortios_firewall_sslsshprofile.labelname {{mkey}}
```
