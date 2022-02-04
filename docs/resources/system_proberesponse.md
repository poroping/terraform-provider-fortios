---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_proberesponse"
description: |-
  Configure system probe response.
---

## fortios_system_proberesponse
Configure system probe response.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `http_probe_value` - Value to respond to the monitoring server.
* `mode` - SLA response mode. Valid values: `none` `http-probe` `twamp` .
* `password` - TWAMP responder password in authentication mode.
* `port` - Port number to response.
* `security_mode` - TWAMP responder security mode. Valid values: `none` `authentication` .
* `timeout` - An inactivity timer for a twamp test session.
* `ttl_mode` - Mode for TWAMP packet TTL modification. Valid values: `reinit` `decrease` `retain` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_proberesponse can be imported using:
```sh
terraform import fortios_system_proberesponse.labelname {{mkey}}
```
