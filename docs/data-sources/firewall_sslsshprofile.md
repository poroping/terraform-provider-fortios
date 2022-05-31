---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_sslsshprofile"
description: |-
  Get information on a fortios Configure SSL/SSH protocol options.
---

# Data Source: fortios_firewall_sslsshprofile
Use this data source to get information on a fortios Configure SSL/SSH protocol options.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `allowlist` - Enable/disable exempting servers by FortiGuard allowlist.
* `block_blacklisted_certificates` - Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blacklist.
* `block_blocklisted_certificates` - Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blocklist.
* `caname` - CA certificate used by SSL Inspection.
* `comment` - Optional comments.
* `mapi_over_https` - Enable/disable inspection of MAPI over HTTPS.
* `name` - Name.
* `rpc_over_https` - Enable/disable inspection of RPC over HTTPS.
* `server_cert_mode` - Re-sign or replace the server's certificate.
* `ssl_anomalies_log` - Enable/disable logging SSL anomalies.
* `ssl_anomaly_log` - Enable/disable logging of SSL anomalies.
* `ssl_exemption_ip_rating` - Enable/disable IP based URL rating.
* `ssl_exemption_log` - Enable/disable logging SSL exemptions.
* `ssl_exemptions_log` - Enable/disable logging SSL exemptions.
* `ssl_handshake_log` - Enable/disable logging of TLS handshakes.
* `ssl_negotiation_log` - Enable/disable logging SSL negotiation.
* `ssl_server_cert_log` - Enable/disable logging of server certificate information.
* `supported_alpn` - Configure ALPN option.
* `untrusted_caname` - Untrusted CA certificate used by SSL Inspection.
* `use_ssl_server` - Enable/disable the use of SSL server table for SSL offloading.
* `whitelist` - Enable/disable exempting servers by FortiGuard whitelist.
* `dot` - Configure DNS over TLS options.The structure of `dot` block is documented below.

The `dot` block contains:

* `cert_validation_failure` - Action based on certificate validation failure.
* `cert_validation_timeout` - Action based on certificate validation timeout.
* `client_certificate` - Action based on received client certificate.
* `expired_server_cert` - Action based on server certificate is expired.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before).
* `revoked_server_cert` - Action based on server certificate is revoked.
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.
* `status` - Configure protocol inspection status.
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported.
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported.
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported.
* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA.
* `ftps` - Configure FTPS options.The structure of `ftps` block is documented below.

The `ftps` block contains:

* `cert_validation_failure` - Action based on certificate validation failure.
* `cert_validation_timeout` - Action based on certificate validation timeout.
* `client_cert_request` - Action based on client certificate request.
* `client_certificate` - Action based on received client certificate.
* `expired_server_cert` - Action based on server certificate is expired.
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate.
* `min_allowed_ssl_version` - Minimum SSL version to be allowed.
* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `revoked_server_cert` - Action based on server certificate is revoked.
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.
* `status` - Configure protocol inspection status.
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported.
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported.
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported.
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported.
* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA.
* `https` - Configure HTTPS options.The structure of `https` block is documented below.

The `https` block contains:

* `cert_probe_failure` - Action based on certificate probe failure.
* `cert_validation_failure` - Action based on certificate validation failure.
* `cert_validation_timeout` - Action based on certificate validation timeout.
* `client_cert_request` - Action based on client certificate request.
* `client_certificate` - Action based on received client certificate.
* `expired_server_cert` - Action based on server certificate is expired.
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate.
* `min_allowed_ssl_version` - Minimum SSL version to be allowed.
* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before).
* `revoked_server_cert` - Action based on server certificate is revoked.
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.
* `status` - Configure protocol inspection status.
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported.
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported.
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported.
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported.
* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA.
* `imaps` - Configure IMAPS options.The structure of `imaps` block is documented below.

The `imaps` block contains:

* `cert_validation_failure` - Action based on certificate validation failure.
* `cert_validation_timeout` - Action based on certificate validation timeout.
* `client_cert_request` - Action based on client certificate request.
* `client_certificate` - Action based on received client certificate.
* `expired_server_cert` - Action based on server certificate is expired.
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate.
* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before).
* `revoked_server_cert` - Action based on server certificate is revoked.
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.
* `status` - Configure protocol inspection status.
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported.
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported.
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported.
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported.
* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA.
* `pop3s` - Configure POP3S options.The structure of `pop3s` block is documented below.

The `pop3s` block contains:

* `cert_validation_failure` - Action based on certificate validation failure.
* `cert_validation_timeout` - Action based on certificate validation timeout.
* `client_cert_request` - Action based on client certificate request.
* `client_certificate` - Action based on received client certificate.
* `expired_server_cert` - Action based on server certificate is expired.
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate.
* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before).
* `revoked_server_cert` - Action based on server certificate is revoked.
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.
* `status` - Configure protocol inspection status.
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported.
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported.
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported.
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported.
* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA.
* `server_cert` - Certificate used by SSL Inspection to replace server certificate.The structure of `server_cert` block is documented below.

The `server_cert` block contains:

* `name` - Certificate list.
* `smtps` - Configure SMTPS options.The structure of `smtps` block is documented below.

The `smtps` block contains:

* `cert_validation_failure` - Action based on certificate validation failure.
* `cert_validation_timeout` - Action based on certificate validation timeout.
* `client_cert_request` - Action based on client certificate request.
* `client_certificate` - Action based on received client certificate.
* `expired_server_cert` - Action based on server certificate is expired.
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate.
* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before).
* `revoked_server_cert` - Action based on server certificate is revoked.
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.
* `status` - Configure protocol inspection status.
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported.
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported.
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported.
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported.
* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA.
* `ssh` - Configure SSH options.The structure of `ssh` block is documented below.

The `ssh` block contains:

* `inspect_all` - Level of SSL inspection.
* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before).
* `ssh_algorithm` - Relative strength of encryption algorithms accepted during negotiation.
* `ssh_tun_policy_check` - Enable/disable SSH tunnel policy check.
* `status` - Configure protocol inspection status.
* `unsupported_version` - Action based on SSH version being unsupported.
* `ssl` - Configure SSL options.The structure of `ssl` block is documented below.

The `ssl` block contains:

* `cert_probe_failure` - Action based on certificate probe failure.
* `cert_validation_failure` - Action based on certificate validation failure.
* `cert_validation_timeout` - Action based on certificate validation timeout.
* `client_cert_request` - Action based on client certificate request.
* `client_certificate` - Action based on received client certificate.
* `expired_server_cert` - Action based on server certificate is expired.
* `inspect_all` - Level of SSL inspection.
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate.
* `min_allowed_ssl_version` - Minimum SSL version to be allowed.
* `revoked_server_cert` - Action based on server certificate is revoked.
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported.
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported.
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported.
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported.
* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA.
* `ssl_exempt` - Servers to exempt from SSL inspection.The structure of `ssl_exempt` block is documented below.

The `ssl_exempt` block contains:

* `address` - IPv4 address object.
* `address6` - IPv6 address object.
* `fortiguard_category` - FortiGuard category ID.
* `id` - ID number.
* `regex` - Exempt servers by regular expression.
* `type` - Type of address object (IPv4 or IPv6) or FortiGuard category.
* `wildcard_fqdn` - Exempt servers by wildcard FQDN.
* `ssl_server` - SSL server settings used for client certificate request.The structure of `ssl_server` block is documented below.

The `ssl_server` block contains:

* `ftps_client_cert_request` - Action based on client certificate request during the FTPS handshake.
* `ftps_client_certificate` - Action based on received client certificate during the FTPS handshake.
* `https_client_cert_request` - Action based on client certificate request during the HTTPS handshake.
* `https_client_certificate` - Action based on received client certificate during the HTTPS handshake.
* `id` - SSL server ID.
* `imaps_client_cert_request` - Action based on client certificate request during the IMAPS handshake.
* `imaps_client_certificate` - Action based on received client certificate during the IMAPS handshake.
* `ip` - IPv4 address of the SSL server.
* `pop3s_client_cert_request` - Action based on client certificate request during the POP3S handshake.
* `pop3s_client_certificate` - Action based on received client certificate during the POP3S handshake.
* `smtps_client_cert_request` - Action based on client certificate request during the SMTPS handshake.
* `smtps_client_certificate` - Action based on received client certificate during the SMTPS handshake.
* `ssl_other_client_cert_request` - Action based on client certificate request during an SSL protocol handshake.
* `ssl_other_client_certificate` - Action based on received client certificate during an SSL protocol handshake.
