---
subcategory: "FortiGate Webfilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_ipsurlfiltercachesetting"
description: |-
  Get information on a fortios Configure IPS URL filter cache settings.
---

# Data Source: fortios_webfilter_ipsurlfiltercachesetting
Use this data source to get information on a fortios Configure IPS URL filter cache settings.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `dns_retry_interval` - Retry interval. Refresh DNS faster than TTL to capture multiple IPs for hosts. 0 means use DNS server's TTL only.
* `extended_ttl` - Extend time to live beyond reported by DNS. 0 means use DNS server's TTL
