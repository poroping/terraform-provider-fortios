---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_log_eventfilter"
description: |-
  Get information on a fortios Configure log event filters.
---

# Data Source: fortios_log_eventfilter
Use this data source to get information on a fortios Configure log event filters.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `cifs` - Enable/disable CIFS logging.
* `connector` - Enable/disable SDN connector logging.
* `endpoint` - Enable/disable endpoint event logging.
* `event` - Enable/disable event logging.
* `fortiextender` - Enable/disable FortiExtender logging.
* `ha` - Enable/disable ha event logging.
* `rest_api` - Enable/disable REST API logging.
* `router` - Enable/disable router event logging.
* `sdwan` - Enable/disable SD-WAN logging.
* `security_rating` - Enable/disable Security Rating result logging.
* `switch_controller` - Enable/disable Switch-Controller logging.
* `system` - Enable/disable system event logging.
* `user` - Enable/disable user authentication event logging.
* `vpn` - Enable/disable VPN event logging.
* `wan_opt` - Enable/disable WAN optimization event logging.
* `webproxy` - Enable/disable web proxy event logging.
* `wireless_activity` - Enable/disable wireless event logging.
