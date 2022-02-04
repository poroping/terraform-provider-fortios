---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_pop3"
description: |-
  Get information on a fortios POP3 server entry configuration.
---

# Data Source: fortios_user_pop3
Use this data source to get information on a fortios POP3 server entry configuration.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) POP3 server entry name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - POP3 server entry name.
* `port` - POP3 service port number.
* `secure` - SSL connection.
* `server` - Server domain name or IP address.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
