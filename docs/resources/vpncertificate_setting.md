---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpncertificate_setting"
description: |-
  VPN certificate setting.
---

## fortios_vpncertificate_setting
VPN certificate setting.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `certname_dsa1024` - 1024 bit DSA key certificate for re-signing server certificates for SSL inspection. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `certname_dsa2048` - 2048 bit DSA key certificate for re-signing server certificates for SSL inspection. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `certname_ecdsa256` - 256 bit ECDSA key certificate for re-signing server certificates for SSL inspection. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `certname_ecdsa384` - 384 bit ECDSA key certificate for re-signing server certificates for SSL inspection. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `certname_ecdsa521` - 521 bit ECDSA key certificate for re-signing server certificates for SSL inspection. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `certname_ed25519` - 253 bit EdDSA key certificate for re-signing server certificates for SSL inspection. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `certname_ed448` - 456 bit EdDSA key certificate for re-signing server certificates for SSL inspection. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `certname_rsa1024` - 1024 bit RSA key certificate for re-signing server certificates for SSL inspection. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `certname_rsa2048` - 2048 bit RSA key certificate for re-signing server certificates for SSL inspection. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `certname_rsa4096` - 4096 bit RSA key certificate for re-signing server certificates for SSL inspection. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `check_ca_cert` - Enable/disable verification of the user certificate and pass authentication if any CA in the chain is trusted (default = enable). Valid values: `enable` `disable` .
* `check_ca_chain` - Enable/disable verification of the entire certificate chain and pass authentication only if the chain is complete and all of the CAs in the chain are trusted (default = disable). Valid values: `enable` `disable` .
* `cmp_key_usage_checking` - Enable/disable server certificate key usage checking in CMP mode (default = enable). Valid values: `enable` `disable` .
* `cmp_save_extra_certs` - Enable/disable saving extra certificates in CMP mode (default = disable). Valid values: `enable` `disable` .
* `cn_allow_multi` - When searching for a matching certificate, allow multiple CN fields in certificate subject name (default = enable). Valid values: `disable` `enable` .
* `cn_match` - When searching for a matching certificate, control how to do CN value matching with certificate subject name (default = substring). Valid values: `substring` `value` .
* `interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `ocsp_default_server` - Default OCSP server. This attribute must reference one of the following datasources: `vpn.certificate.ocsp-server.name` .
* `ocsp_option` - Specify whether the OCSP URL is from certificate or configured OCSP server. Valid values: `certificate` `server` .
* `ocsp_status` - Enable/disable receiving certificates using the OCSP. Valid values: `enable` `disable` .
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default` `SSLv3` `TLSv1` `TLSv1-1` `TLSv1-2` .
* `ssl_ocsp_source_ip` - Source IP address to use to communicate with the OCSP server.
* `strict_crl_check` - Enable/disable strict mode CRL checking. Valid values: `enable` `disable` .
* `strict_ocsp_check` - Enable/disable strict mode OCSP checking. Valid values: `enable` `disable` .
* `subject_match` - When searching for a matching certificate, control how to do RDN value matching with certificate subject name (default = substring). Valid values: `substring` `value` .
* `subject_set` - When searching for a matching certificate, control how to do RDN set matching with certificate subject name (default = subset). Valid values: `subset` `superset` .
* `crl_verification` - CRL verification options. The structure of `crl_verification` block is documented below.

The `crl_verification` block contains:

* `chain_crl_absence` - CRL verification option when CRL of any certificate in chain is absent (default = ignore). Valid values: `ignore` `revoke` .
* `expiry` - CRL verification option when CRL is expired (default = ignore). Valid values: `ignore` `revoke` .
* `leaf_crl_absence` - CRL verification option when leaf CRL is absent (default = ignore). Valid values: `ignore` `revoke` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_vpncertificate_setting can be imported using:
```sh
terraform import fortios_vpncertificate_setting.labelname {{mkey}}
```
