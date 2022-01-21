---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_pop3"
description: |-
  POP3 server entry configuration.
---

## fortios_user_pop3
POP3 server entry configuration.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `name` - POP3 server entry name.
* `port` - POP3 service port number.
* `secure` - SSL connection. Valid values: `none` `starttls` `pop3s` .
* `server` - {<name_str|ip_str>} server domain name or IP.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default` `SSLv3` `TLSv1` `TLSv1-1` `TLSv1-2` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_pop3 can be imported using:
```sh
terraform import fortios_user_pop3.labelname {{mkey}}
```
