---
subcategory: "FortiGate Ips"
layout: "fortios"
page_title: "FortiOS: fortios_ips_custom"
description: |-
  Get information on a fortios Configure IPS custom signature.
---

# Data Source: fortios_ips_custom
Use this data source to get information on a fortios Configure IPS custom signature.


## Example Usage

```hcl

```

## Argument Reference

* `tag` - (Required) Signature tag.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `action` - Default action (pass or block) for this signature.
* `application` - Applications to be protected. Blank for all applications.
* `comment` - Comment.
* `location` - Protect client or server traffic.
* `log` - Enable/disable logging.
* `log_packet` - Enable/disable packet logging.
* `os` - Operating system(s) that the signature protects. Blank for all operating systems.
* `protocol` - Protocol(s) that the signature scans. Blank for all protocols.
* `rule_id` - Signature ID.
* `severity` - Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.
* `signature` - Custom signature enclosed in single quotes.
* `status` - Enable/disable this signature.
* `tag` - Signature tag.
