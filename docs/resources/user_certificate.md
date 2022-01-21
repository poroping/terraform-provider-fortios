---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_certificate"
description: |-
  Configure certificate users.
---

## fortios_user_certificate
Configure certificate users.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `common_name` - Certificate common name.
* `id` - User ID.
* `issuer` - CA certificate used for client certificate verification. This attribute must reference one of the following datasources: `vpn.certificate.ca.name` .
* `name` - User name.
* `status` - Enable/disable allowing the certificate user to authenticate with the FortiGate unit. Valid values: `enable` `disable` .
* `type` - Type of certificate authentication method. Valid values: `single-certificate` `trusted-issuer` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_certificate can be imported using:
```sh
terraform import fortios_user_certificate.labelname {{mkey}}
```
