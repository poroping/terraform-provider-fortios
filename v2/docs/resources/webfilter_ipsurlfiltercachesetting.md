---
subcategory: "FortiGate Webfilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_ipsurlfiltercachesetting"
description: |-
  Configure IPS URL filter cache settings.
---

## fortios_webfilter_ipsurlfiltercachesetting
Configure IPS URL filter cache settings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `dns_retry_interval` - Retry interval. Refresh DNS faster than TTL to capture multiple IPs for hosts. 0 means use DNS server's TTL only.
* `extended_ttl` - Extend time to live beyond reported by DNS. 0 means use DNS server's TTL

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_webfilter_ipsurlfiltercachesetting can be imported using:
```sh
terraform import fortios_webfilter_ipsurlfiltercachesetting.labelname {{mkey}}
```
