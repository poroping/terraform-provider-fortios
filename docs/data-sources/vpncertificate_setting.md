---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpncertificate_setting"
description: |-
  Get information on a fortios VPN certificate setting.
---

# Data Source: fortios_vpncertificate_setting
Use this data source to get information on a fortios VPN certificate setting.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `cert_expire_warning` - Number of days before a certificate expires to send a warning. Set to 0 to disable sending of the warning (0 - 100, default = 14).
* `certname_dsa1024` - 1024 bit DSA key certificate for re-signing server certificates for SSL inspection.
* `certname_dsa2048` - 2048 bit DSA key certificate for re-signing server certificates for SSL inspection.
* `certname_ecdsa256` - 256 bit ECDSA key certificate for re-signing server certificates for SSL inspection.
* `certname_ecdsa384` - 384 bit ECDSA key certificate for re-signing server certificates for SSL inspection.
* `certname_ecdsa521` - 521 bit ECDSA key certificate for re-signing server certificates for SSL inspection.
* `certname_ed25519` - 253 bit EdDSA key certificate for re-signing server certificates for SSL inspection.
* `certname_ed448` - 456 bit EdDSA key certificate for re-signing server certificates for SSL inspection.
* `certname_rsa1024` - 1024 bit RSA key certificate for re-signing server certificates for SSL inspection.
* `certname_rsa2048` - 2048 bit RSA key certificate for re-signing server certificates for SSL inspection.
* `certname_rsa4096` - 4096 bit RSA key certificate for re-signing server certificates for SSL inspection.
* `check_ca_cert` - Enable/disable verification of the user certificate and pass authentication if any CA in the chain is trusted (default = enable).
* `check_ca_chain` - Enable/disable verification of the entire certificate chain and pass authentication only if the chain is complete and all of the CAs in the chain are trusted (default = disable).
* `cmp_key_usage_checking` - Enable/disable server certificate key usage checking in CMP mode (default = enable).
* `cmp_save_extra_certs` - Enable/disable saving extra certificates in CMP mode (default = disable).
* `cn_allow_multi` - When searching for a matching certificate, allow multiple CN fields in certificate subject name (default = enable).
* `cn_match` - When searching for a matching certificate, control how to do CN value matching with certificate subject name (default = substring).
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `ocsp_default_server` - Default OCSP server.
* `ocsp_option` - Specify whether the OCSP URL is from certificate or configured OCSP server.
* `ocsp_status` - Enable/disable receiving certificates using the OCSP.
* `proxy` - Proxy server FQDN or IP for OCSP/CA queries during certificate verification.
* `proxy_password` - Proxy server password.
* `proxy_port` - Proxy server port (1 - 65535, default = 8080).
* `proxy_username` - Proxy server user name.
* `source_ip` - Source IP address for dynamic AIA and OCSP queries.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `ssl_ocsp_source_ip` - Source IP address to use to communicate with the OCSP server.
* `strict_crl_check` - Enable/disable strict mode CRL checking.
* `strict_ocsp_check` - Enable/disable strict mode OCSP checking.
* `subject_match` - When searching for a matching certificate, control how to do RDN value matching with certificate subject name (default = substring).
* `subject_set` - When searching for a matching certificate, control how to do RDN set matching with certificate subject name (default = subset).
* `crl_verification` - CRL verification options.The structure of `crl_verification` block is documented below.

The `crl_verification` block contains:

* `chain_crl_absence` - CRL verification option when CRL of any certificate in chain is absent (default = ignore).
* `expiry` - CRL verification option when CRL is expired (default = ignore).
* `leaf_crl_absence` - CRL verification option when leaf CRL is absent (default = ignore).
