---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_proberesponse"
description: |-
  Get information on a fortios Configure system probe response.
---

# Data Source: fortios_system_proberesponse
Use this data source to get information on a fortios Configure system probe response.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `http_probe_value` - Value to respond to the monitoring server.
* `mode` - SLA response mode.
* `password` - TWAMP responder password in authentication mode.
* `port` - Port number to response.
* `security_mode` - TWAMP responder security mode.
* `timeout` - An inactivity timer for a twamp test session.
* `ttl_mode` - Mode for TWAMP packet TTL modification.
