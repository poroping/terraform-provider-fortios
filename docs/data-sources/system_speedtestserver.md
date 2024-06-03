---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_speedtestserver"
description: |-
  Get information on a fortios Configure speed test server list.
---

# Data Source: fortios_system_speedtestserver
Use this data source to get information on a fortios Configure speed test server list.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Speed test server name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - Speed test server name.
* `timestamp` - Speed test server timestamp.
* `host` - Hosts of the server.The structure of `host` block is documented below.

The `host` block contains:

* `distance` - Speed test host distance.
* `id` - Server host ID.
* `ip` - Server host IPv4 address.
* `latitude` - Speed test host latitude.
* `longitude` - Speed test host longitude.
* `password` - Speed test host password.
* `port` - Server host port number to communicate with client.
* `user` - Speed test host user name.
