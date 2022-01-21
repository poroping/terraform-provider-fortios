---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipam"
description: |-
  Configure IP address management services.
---

## fortios_system_ipam
Configure IP address management services.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `pool_subnet` - Configure IPAM pool subnet, Class A - Class B subnet.
* `server_type` - Configure the type of IPAM server to use. Valid values: `cloud` `fabric-root` .
* `status` - Enable/disable IP address management services. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_ipam can be imported using:
```sh
terraform import fortios_system_ipam.labelname {{mkey}}
```
