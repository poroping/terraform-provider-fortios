---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_emailserver"
description: |-
  Get information on a fortios Configure the email server used by the FortiGate various things. For example, for sending email messages to users to support user authentication features.
---

# Data Source: fortios_system_emailserver
Use this data source to get information on a fortios Configure the email server used by the FortiGate various things. For example, for sending email messages to users to support user authentication features.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `authenticate` - Enable/disable authentication.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `password` - SMTP server user password for authentication.
* `port` - SMTP server port.
* `reply_to` - Reply-To email address.
* `security` - Connection security used by the email server.
* `server` - SMTP server IP address or hostname.
* `source_ip` - SMTP server IPv4 source IP.
* `source_ip6` - SMTP server IPv6 source IP.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `type` - Use FortiGuard Message service or custom email server.
* `username` - SMTP server user name for authentication.
* `validate_server` - Enable/disable validation of server certificate.
