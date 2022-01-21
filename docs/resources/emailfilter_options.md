---
subcategory: "FortiGate Emailfilter"
layout: "fortios"
page_title: "FortiOS: fortios_emailfilter_options"
description: |-
  Configure AntiSpam options.
---

## fortios_emailfilter_options
Configure AntiSpam options.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `dns_timeout` - DNS query time out (1 - 30 sec).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_emailfilter_options can be imported using:
```sh
terraform import fortios_emailfilter_options.labelname {{mkey}}
```
