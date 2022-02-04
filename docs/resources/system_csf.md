---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_csf"
description: |-
  Add this FortiGate to a Security Fabric or set up a new Security Fabric on this FortiGate.
---

## fortios_system_csf
Add this FortiGate to a Security Fabric or set up a new Security Fabric on this FortiGate.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `accept_auth_by_cert` - Accept connections with unknown certificates and ask admin for approval. Valid values: `disable` `enable` .
* `authorization_request_type` - Authorization request type. Valid values: `serial` `certificate` .
* `certificate` - Certificate. This attribute must reference one of the following datasources: `certificate.local.name` .
* `configuration_sync` - Configuration sync mode. Valid values: `default` `local` .
* `downstream_access` - Enable/disable downstream device access to this device's configuration and data. Valid values: `enable` `disable` .
* `downstream_accprofile` - Default access profile for requests from downstream devices. This attribute must reference one of the following datasources: `system.accprofile.name` .
* `fabric_object_unification` - Fabric CMDB Object Unification. Valid values: `default` `local` .
* `fabric_workers` - Number of worker processes for Security Fabric daemon.
* `forticloud_account_enforcement` - Fabric FortiCloud account unification. Valid values: `enable` `disable` .
* `group_name` - Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.
* `group_password` - Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.
* `log_unification` - Enable/disable broadcast of discovery messages for log unification. Valid values: `disable` `enable` .
* `management_ip` - Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
* `management_port` - Overriding port for management connection (Overrides admin port).
* `saml_configuration_sync` - SAML setting configuration synchronization. Valid values: `default` `local` .
* `status` - Enable/disable Security Fabric. Valid values: `enable` `disable` .
* `upstream` - IP/FQDN of the FortiGate upstream from this FortiGate in the Security Fabric.
* `upstream_ip` - IP address of the FortiGate upstream from this FortiGate in the Security Fabric.
* `upstream_port` - The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).
* `fabric_connector` - Fabric connector configuration. The structure of `fabric_connector` block is documented below.

The `fabric_connector` block contains:

* `accprofile` - Override access profile. This attribute must reference one of the following datasources: `system.accprofile.name` .
* `configuration_write_access` - Enable/disable downstream device write access to configuration. Valid values: `enable` `disable` .
* `serial` - Serial.
* `fabric_device` - Fabric device configuration. The structure of `fabric_device` block is documented below.

The `fabric_device` block contains:

* `access_token` - Device access token.
* `device_ip` - Device IP.
* `https_port` - HTTPS port for fabric device.
* `name` - Device name.
* `trusted_list` - Pre-authorized and blocked security fabric nodes. The structure of `trusted_list` block is documented below.

The `trusted_list` block contains:

* `action` - Security fabric authorization action. Valid values: `accept` `deny` .
* `authorization_type` - Authorization type. Valid values: `serial` `certificate` .
* `certificate` - Certificate.
* `downstream_authorization` - Trust authorizations by this node's administrator. Valid values: `enable` `disable` .
* `ha_members` - HA members.
* `name` - Name.
* `serial` - Serial.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_csf can be imported using:
```sh
terraform import fortios_system_csf.labelname {{mkey}}
```
