---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_alarm"
description: |-
  Configure alarm.
---

## fortios_system_alarm
Configure alarm.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `audible` - Enable/disable audible alarm. Valid values: `enable` `disable` .
* `status` - Enable/disable alarm. Valid values: `enable` `disable` .
* `groups` - Alarm groups. The structure of `groups` block is documented below.

The `groups` block contains:

* `admin_auth_failure_threshold` - Admin authentication failure threshold.
* `admin_auth_lockout_threshold` - Admin authentication lockout threshold.
* `decryption_failure_threshold` - Decryption failure threshold.
* `encryption_failure_threshold` - Encryption failure threshold.
* `fw_policy_id` - Firewall policy ID.
* `fw_policy_id_threshold` - Firewall policy ID threshold.
* `id` - Group ID.
* `log_full_warning_threshold` - Log full warning threshold.
* `period` - Time period in seconds (0 = from start up).
* `replay_attempt_threshold` - Replay attempt threshold.
* `self_test_failure_threshold` - Self-test failure threshold.
* `user_auth_failure_threshold` - User authentication failure threshold.
* `user_auth_lockout_threshold` - User authentication lockout threshold.
* `fw_policy_violations` - Firewall policy violations. The structure of `fw_policy_violations` block is documented below.

The `fw_policy_violations` block contains:

* `dst_ip` - Destination IP (0=all).
* `dst_port` - Destination port (0=all).
* `id` - Firewall policy violations ID.
* `src_ip` - Source IP (0=all).
* `src_port` - Source port (0=all).
* `threshold` - Firewall policy violation threshold.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_alarm can be imported using:
```sh
terraform import fortios_system_alarm.labelname {{mkey}}
```
