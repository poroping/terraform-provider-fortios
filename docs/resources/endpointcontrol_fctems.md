---
subcategory: "FortiGate Endpoint-Control"
layout: "fortios"
page_title: "FortiOS: fortios_endpointcontrol_fctems"
description: |-
  Configure FortiClient Enterprise Management Server (EMS) entries.
---

## fortios_endpointcontrol_fctems
Configure FortiClient Enterprise Management Server (EMS) entries.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `admin_password` - FortiClient EMS admin password.
* `admin_username` - FortiClient EMS admin username.
* `call_timeout` - FortiClient EMS call timeout in seconds (1 - 180 seconds, default = 30).
* `capabilities` - List of EMS capabilities. Valid values: `fabric-auth` `silent-approval` `websocket` `websocket-malware` `push-ca-certs` `common-tags-api` .
* `certificate` - FortiClient EMS certificate. This attribute must reference one of the following datasources: `certificate.remote.name` .
* `cloud_server_type` - Cloud server type. Valid values: `production` `alpha` `beta` .
* `fortinetone_cloud_authentication` - Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `enable` `disable` .
* `https_port` - FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
* `name` - FortiClient Enterprise Management Server (EMS) name.
* `preserve_ssl_session` - Enable/disable preservation of EMS SSL session connection. Warning, most users should not touch this setting. Valid values: `enable` `disable` .
* `pull_avatars` - Enable/disable pulling avatars from EMS. Valid values: `enable` `disable` .
* `pull_malware_hash` - Enable/disable pulling FortiClient malware hash from EMS. Valid values: `enable` `disable` .
* `pull_sysinfo` - Enable/disable pulling SysInfo from EMS. Valid values: `enable` `disable` .
* `pull_tags` - Enable/disable pulling FortiClient user tags from EMS. Valid values: `enable` `disable` .
* `pull_vulnerabilities` - Enable/disable pulling vulnerabilities from EMS. Valid values: `enable` `disable` .
* `serial_number` - FortiClient EMS Serial Number.
* `server` - FortiClient EMS FQDN or IPv4 address.
* `source_ip` - REST API call source IP.
* `status_check_interval` - FortiClient EMS call timeout in seconds (1 - 120 seconds, default = 5).
* `websocket_override` - Enable/disable override behavior for how this FortiGate unit connects to EMS using a WebSocket connection. Valid values: `disable` `enable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_endpointcontrol_fctems can be imported using:
```sh
terraform import fortios_endpointcontrol_fctems.labelname {{mkey}}
```
