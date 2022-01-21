---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_emailserver"
description: |-
  Configure the email server used by the FortiGate various things. For example, for sending email messages to users to support user authentication features.
---

## fortios_system_emailserver
Configure the email server used by the FortiGate various things. For example, for sending email messages to users to support user authentication features.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `authenticate` - Enable/disable authentication. Valid values: `enable` `disable` .
* `interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `password` - SMTP server user password for authentication.
* `port` - SMTP server port.
* `reply_to` - Reply-To email address.
* `security` - Connection security used by the email server. Valid values: `none` `starttls` `smtps` .
* `server` - SMTP server IP address or hostname.
* `source_ip` - SMTP server IPv4 source IP.
* `source_ip6` - SMTP server IPv6 source IP.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default` `SSLv3` `TLSv1` `TLSv1-1` `TLSv1-2` .
* `type` - Use FortiGuard Message service or custom email server. Valid values: `custom` .
* `username` - SMTP server user name for authentication.
* `validate_server` - Enable/disable validation of server certificate. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_emailserver can be imported using:
```sh
terraform import fortios_system_emailserver.labelname {{mkey}}
```
