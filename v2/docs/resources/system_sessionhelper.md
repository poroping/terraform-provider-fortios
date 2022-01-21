---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sessionhelper"
description: |-
  Configure session helper.
---

## fortios_system_sessionhelper
Configure session helper.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.

* `id` - Session helper ID.
* `name` - Helper name. Valid values: `ftp` `tftp` `ras` `h323` `tns` `mms` `sip` `pptp` `rtsp` `dns-udp` `dns-tcp` `pmap` `rsh` `dcerpc` `mgcp` .
* `port` - Protocol port.
* `protocol` - Protocol number.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_sessionhelper can be imported using:
```sh
terraform import fortios_system_sessionhelper.labelname {{mkey}}
```
