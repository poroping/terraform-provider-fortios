---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_anqpnairealm"
description: |-
  Configure network access identifier (NAI) realm.
---

## fortios_wirelesscontrollerhotspot20_anqpnairealm
Configure network access identifier (NAI) realm.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `name` - NAI realm list name.
* `nai_list` - NAI list. The structure of `nai_list` block is documented below.

The `nai_list` block contains:

* `encoding` - Enable/disable format in accordance with IETF RFC 4282. Valid values: `disable` `enable` .
* `nai_realm` - Configure NAI realms (delimited by a semi-colon character).
* `name` - NAI realm name.
* `eap_method` - EAP Methods. The structure of `eap_method` block is documented below.

The `eap_method` block contains:

* `index` - EAP method index.
* `method` - EAP method type. Valid values: `eap-identity` `eap-md5` `eap-tls` `eap-ttls` `eap-peap` `eap-sim` `eap-aka` `eap-aka-prime` .
* `auth_param` - EAP auth param. The structure of `auth_param` block is documented below.

The `auth_param` block contains:

* `id` - ID of authentication parameter. Valid values: `non-eap-inner-auth` `inner-auth-eap` `credential` `tunneled-credential` .
* `index` - Param index.
* `val` - Value of authentication parameter. Valid values: `eap-identity` `eap-md5` `eap-tls` `eap-ttls` `eap-peap` `eap-sim` `eap-aka` `eap-aka-prime` `non-eap-pap` `non-eap-chap` `non-eap-mschap` `non-eap-mschapv2` `cred-sim` `cred-usim` `cred-nfc` `cred-hardware-token` `cred-softoken` `cred-certificate` `cred-user-pwd` `cred-none` `cred-vendor-specific` `tun-cred-sim` `tun-cred-usim` `tun-cred-nfc` `tun-cred-hardware-token` `tun-cred-softoken` `tun-cred-certificate` `tun-cred-user-pwd` `tun-cred-anonymous` `tun-cred-vendor-specific` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontrollerhotspot20_anqpnairealm can be imported using:
```sh
terraform import fortios_wirelesscontrollerhotspot20_anqpnairealm.labelname {{mkey}}
```
