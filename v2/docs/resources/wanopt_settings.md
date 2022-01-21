---
subcategory: "FortiGate Wanopt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_settings"
description: |-
  Configure WAN optimization settings.
---

## fortios_wanopt_settings
Configure WAN optimization settings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `auto_detect_algorithm` - Auto detection algorithms used in tunnel negotiations. Valid values: `simple` `diff-req-resp` .
* `host_id` - Local host ID (must also be entered in the remote FortiGate's peer list).
* `tunnel_ssl_algorithm` - Relative strength of encryption algorithms accepted during tunnel negotiation. Valid values: `high` `medium` `low` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wanopt_settings can be imported using:
```sh
terraform import fortios_wanopt_settings.labelname {{mkey}}
```
