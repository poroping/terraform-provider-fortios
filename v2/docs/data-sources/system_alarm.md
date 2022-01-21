---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_alarm"
description: |-
  Get information on a fortios Configure alarm.
---

# Data Source: fortios_system_alarm
Use this data source to get information on a fortios Configure alarm.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `audible` - Enable/disable audible alarm.
* `status` - Enable/disable alarm.
* `groups` - Alarm groups.The structure of `groups` block is documented below.

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
* `fw_policy_violations` - Firewall policy violations.The structure of `fw_policy_violations` block is documented below.

The `fw_policy_violations` block contains:

* `dst_ip` - Destination IP (0=all).
* `dst_port` - Destination port (0=all).
* `id` - Firewall policy violations ID.
* `src_ip` - Source IP (0=all).
* `src_port` - Source port (0=all).
* `threshold` - Firewall policy violation threshold.
