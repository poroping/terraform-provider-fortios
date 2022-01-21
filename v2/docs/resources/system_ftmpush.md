---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ftmpush"
description: |-
  Configure FortiToken Mobile push services.
---

## fortios_system_ftmpush
Configure FortiToken Mobile push services.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `server` - IPv4 address or domain name of FortiToken Mobile push services server.
* `server_cert` - Name of the server certificate to be used for SSL (default = Fortinet_Factory). This attribute must reference one of the following datasources: `certificate.local.name` .
* `server_ip` - IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).
* `server_port` - Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).
* `status` - Enable/disable the use of FortiToken Mobile push services. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_ftmpush can be imported using:
```sh
terraform import fortios_system_ftmpush.labelname {{mkey}}
```
