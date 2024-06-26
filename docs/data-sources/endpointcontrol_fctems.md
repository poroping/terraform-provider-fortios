---
subcategory: "FortiGate Endpoint-Control"
layout: "fortios"
page_title: "FortiOS: fortios_endpointcontrol_fctems"
description: |-
  Get information on a fortios Configure FortiClient Enterprise Management Server (EMS) entries.
---

# Data Source: fortios_endpointcontrol_fctems
Use this data source to get information on a fortios Configure FortiClient Enterprise Management Server (EMS) entries.


## Example Usage

```hcl

```

## Argument Reference

* `ems_id` - (Required) EMS ID in order (1 - 7).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `admin_password` - FortiClient EMS admin password.
* `admin_username` - FortiClient EMS admin username.
* `call_timeout` - FortiClient EMS call timeout in seconds (1 - 180 seconds, default = 30).
* `capabilities` - List of EMS capabilities.
* `certificate` - FortiClient EMS certificate.
* `cloud_server_type` - Cloud server type.
* `dirty_reason` - Dirty Reason for FortiClient EMS.
* `ems_id` - EMS ID in order (1 - 7).
* `fortinetone_cloud_authentication` - Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account.
* `https_port` - FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `name` - FortiClient Enterprise Management Server (EMS) name.
* `out_of_sync_threshold` - Outdated resource threshold in seconds (10 - 3600, default = 180).
* `preserve_ssl_session` - Enable/disable preservation of EMS SSL session connection. Warning, most users should not touch this setting.
* `pull_avatars` - Enable/disable pulling avatars from EMS.
* `pull_malware_hash` - Enable/disable pulling FortiClient malware hash from EMS.
* `pull_sysinfo` - Enable/disable pulling SysInfo from EMS.
* `pull_tags` - Enable/disable pulling FortiClient user tags from EMS.
* `pull_vulnerabilities` - Enable/disable pulling vulnerabilities from EMS.
* `serial_number` - EMS Serial Number.
* `server` - FortiClient EMS FQDN or IPv4 address.
* `source_ip` - REST API call source IP.
* `status` - Enable or disable this EMS configuration.
* `status_check_interval` - FortiClient EMS call timeout in seconds (1 - 120 seconds, default = 5).
* `tenant_id` - EMS Tenant ID.
* `trust_ca_cn` - Enable/disable trust of the EMS certificate issuer(CA) and common name(CN) for certificate auto-renewal.
* `websocket_override` - Enable/disable override behavior for how this FortiGate unit connects to EMS using a WebSocket connection.
