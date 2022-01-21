---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_dospolicy"
description: |-
  Get information on a fortios Configure IPv4 DoS policies.
---

# Data Source: fortios_firewall_dospolicy
Use this data source to get information on a fortios Configure IPv4 DoS policies.


## Example Usage

```hcl

```

## Argument Reference

* `policyid` - (Required) Policy ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comments` - Comment.
* `interface` - Incoming interface name from available interfaces.
* `name` - Policy name.
* `policyid` - Policy ID.
* `status` - Enable/disable this policy.
* `anomaly` - Anomaly name.The structure of `anomaly` block is documented below.

The `anomaly` block contains:

* `action` - Action taken when the threshold is reached.
* `log` - Enable/disable anomaly logging.
* `name` - Anomaly name.
* `quarantine` - Quarantine method.
* `quarantine_expiry` - Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.
* `quarantine_log` - Enable/disable quarantine logging.
* `status` - Enable/disable this anomaly.
* `threshold` - Anomaly threshold. Number of detected instances per minute that triggers the anomaly action.
* `thresholddefault` - Number of detected instances per minute which triggers action (1 - 2147483647, default = 1000). Note that each anomaly has a different threshold value assigned to it.
* `dstaddr` - Destination address name from available addresses.The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name.
* `service` - Service object from available options.The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service name.
* `srcaddr` - Source address name from available addresses.The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name.
