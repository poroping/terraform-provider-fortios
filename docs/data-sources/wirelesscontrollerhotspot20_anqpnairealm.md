---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_anqpnairealm"
description: |-
  Get information on a fortios Configure network access identifier (NAI) realm.
---

# Data Source: fortios_wirelesscontrollerhotspot20_anqpnairealm
Use this data source to get information on a fortios Configure network access identifier (NAI) realm.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) NAI realm list name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - NAI realm list name.
* `nai_list` - NAI list.The structure of `nai_list` block is documented below.

The `nai_list` block contains:

* `encoding` - Enable/disable format in accordance with IETF RFC 4282.
* `nai_realm` - Configure NAI realms (delimited by a semi-colon character).
* `name` - NAI realm name.
* `eap_method` - EAP Methods.The structure of `eap_method` block is documented below.

The `eap_method` block contains:

* `index` - EAP method index.
* `method` - EAP method type.
* `auth_param` - EAP auth param.The structure of `auth_param` block is documented below.

The `auth_param` block contains:

* `id` - ID of authentication parameter.
* `index` - Param index.
* `val` - Value of authentication parameter.
