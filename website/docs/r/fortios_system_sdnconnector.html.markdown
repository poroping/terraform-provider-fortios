---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sdn_connector"
description: |-
  Configure connection to SDN Connector.
---

## fortios_system_sdn_connector
Configure connection to SDN Connector.
## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `access_key` - AWS / ACS access key ID.
* `api_key` - IBM cloud API key or service ID API key.
* `azure_region` - Azure server region. Valid values: `global` `china` `germany` `usgov` `local` .
* `client_id` - Azure client ID (application ID).
* `client_secret` - Azure client secret (application key).
* `compartment_id` - Compartment ID.
* `compute_generation` - Compute generation for IBM cloud infrastructure.
* `domain` - Domain name.
* `gcp_project` - GCP project name.
* `group_name` - Group name of computers.
* `ha_status` - Enable/disable use for FortiGate HA service. Valid values: `disable` `enable` .
* `ibm_region` - IBM cloud region name. Valid values: `us-south` `us-east` `germany` `great-britain` `japan` `australia` .
* `login_endpoint` - Azure Stack login endpoint.
* `name` - SDN connector name.
* `oci_cert` - OCI certificate. This attribute must reference one of the following datasources: `certificate.local.name` .
* `oci_fingerprint` - OCI pubkey fingerprint.
* `oci_region` - OCI server region.
* `oci_region_type` - OCI region type. Valid values: `commercial` `government` .
* `password` - Password of the remote SDN connector as login credentials.
* `private_key` - Private key of GCP service account.
* `region` - AWS / ACS region name.
* `resource_group` - Azure resource group.
* `resource_url` - Azure Stack resource URL.
* `secret_key` - AWS / ACS secret access key.
* `secret_token` - Secret token of Kubernetes service account.
* `server` - Server address of the remote SDN connector.
* `server_port` - Port number of the remote SDN connector.
* `service_account` - GCP service account email.
* `status` - Enable/disable connection to the remote SDN connector. Valid values: `disable` `enable` .
* `subscription_id` - Azure subscription ID.
* `tenant_id` - Tenant ID (directory ID).
* `type` - Type of SDN connector. Valid values: `aci` `alicloud` `aws` `azure` `gcp` `nsx` `nuage` `oci` `openstack` `kubernetes` `vmware` `sepm` `aci-direct` `ibm` `nutanix` .
* `update_interval` - Dynamic object update interval (30 - 3600 sec, default = 60, 0 = disabled).
* `use_metadata_iam` - Enable/disable use of IAM role from metadata to call API. Valid values: `disable` `enable` .
* `user_id` - User ID.
* `username` - Username of the remote SDN connector as login credentials.
* `vcenter_password` - vCenter server password for NSX quarantine.
* `vcenter_server` - vCenter server address for NSX quarantine.
* `vcenter_username` - vCenter server username for NSX quarantine.
* `verify_certificate` - Enable/disable server certificate verification. Valid values: `disable` `enable` .
* `vpc_id` - AWS VPC ID.
* `external_ip` - Configure GCP external IP. The structure of `external_ip` block is documented below.

The `external_ip` block contains:

* `name` - External IP name.
* `nic` - Configure Azure network interface. The structure of `nic` block is documented below.

The `nic` block contains:

* `name` - Network interface name.
* `ip` - Configure IP configuration. The structure of `ip` block is documented below.

The `ip` block contains:

* `name` - IP configuration name.
* `public_ip` - Public IP name.
* `resource_group` - Resource group of Azure public IP.
* `route` - Configure GCP route. The structure of `route` block is documented below.

The `route` block contains:

* `name` - Route name.
* `route_table` - Configure Azure route table. The structure of `route_table` block is documented below.

The `route_table` block contains:

* `name` - Route table name.
* `resource_group` - Resource group of Azure route table.
* `subscription_id` - Subscription ID of Azure route table.
* `route` - Configure Azure route. The structure of `route` block is documented below.

The `route` block contains:

* `name` - Route name.
* `next_hop` - Next hop address.
* `server_list` - Server address list of the remote SDN connector. The structure of `server_list` block is documented below.

The `server_list` block contains:

* `ip` - IPv4 address.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

system_sdn_connector can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_sdn_connector.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
