---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_csf"
description: |-
  Get information on a fortios Add this FortiGate to a Security Fabric or set up a new Security Fabric on this FortiGate.
---

# Data Source: fortios_system_csf
Use this data source to get information on a fortios Add this FortiGate to a Security Fabric or set up a new Security Fabric on this FortiGate.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `accept_auth_by_cert` - Accept connections with unknown certificates and ask admin for approval.
* `authorization_request_type` - Authorization request type.
* `certificate` - Certificate.
* `configuration_sync` - Configuration sync mode.
* `downstream_access` - Enable/disable downstream device access to this device's configuration and data.
* `downstream_accprofile` - Default access profile for requests from downstream devices.
* `fabric_object_unification` - Fabric CMDB Object Unification.
* `fabric_workers` - Number of worker processes for Security Fabric daemon.
* `group_name` - Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.
* `group_password` - Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.
* `log_unification` - Enable/disable broadcast of discovery messages for log unification.
* `management_ip` - Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
* `management_port` - Overriding port for management connection (Overrides admin port).
* `saml_configuration_sync` - SAML setting configuration synchronization.
* `status` - Enable/disable Security Fabric.
* `upstream` - IP/FQDN of the FortiGate upstream from this FortiGate in the Security Fabric.
* `upstream_ip` - IP address of the FortiGate upstream from this FortiGate in the Security Fabric.
* `upstream_port` - The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).
* `fabric_connector` - Fabric connector configuration.The structure of `fabric_connector` block is documented below.

The `fabric_connector` block contains:

* `accprofile` - Override access profile.
* `configuration_write_access` - Enable/disable downstream device write access to configuration.
* `serial` - Serial.
* `fabric_device` - Fabric device configuration.The structure of `fabric_device` block is documented below.

The `fabric_device` block contains:

* `access_token` - Device access token.
* `device_ip` - Device IP.
* `https_port` - HTTPS port for fabric device.
* `name` - Device name.
* `trusted_list` - Pre-authorized and blocked security fabric nodes.The structure of `trusted_list` block is documented below.

The `trusted_list` block contains:

* `action` - Security fabric authorization action.
* `authorization_type` - Authorization type.
* `certificate` - Certificate.
* `downstream_authorization` - Trust authorizations by this node's administrator.
* `ha_members` - HA members.
* `name` - Name.
* `serial` - Serial.
