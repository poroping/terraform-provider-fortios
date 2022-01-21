---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sdnconnector"
description: |-
  Get information on a fortios Configure connection to SDN Connector.
---

# Data Source: fortios_system_sdnconnector
Use this data source to get information on a fortios Configure connection to SDN Connector.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) SDN connector name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `access_key` - AWS / ACS access key ID.
* `api_key` - IBM cloud API key or service ID API key.
* `azure_region` - Azure server region.
* `client_id` - Azure client ID (application ID).
* `client_secret` - Azure client secret (application key).
* `compartment_id` - Compartment ID.
* `compute_generation` - Compute generation for IBM cloud infrastructure.
* `domain` - Domain name.
* `gcp_project` - GCP project name.
* `group_name` - Group name of computers.
* `ha_status` - Enable/disable use for FortiGate HA service.
* `ibm_region` - IBM cloud region name.
* `login_endpoint` - Azure Stack login endpoint.
* `name` - SDN connector name.
* `oci_cert` - OCI certificate.
* `oci_fingerprint` - OCI pubkey fingerprint.
* `oci_region` - OCI server region.
* `oci_region_type` - OCI region type.
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
* `status` - Enable/disable connection to the remote SDN connector.
* `subscription_id` - Azure subscription ID.
* `tenant_id` - Tenant ID (directory ID).
* `type` - Type of SDN connector.
* `update_interval` - Dynamic object update interval (30 - 3600 sec, default = 60, 0 = disabled).
* `use_metadata_iam` - Enable/disable use of IAM role from metadata to call API.
* `user_id` - User ID.
* `username` - Username of the remote SDN connector as login credentials.
* `vcenter_password` - vCenter server password for NSX quarantine.
* `vcenter_server` - vCenter server address for NSX quarantine.
* `vcenter_username` - vCenter server username for NSX quarantine.
* `verify_certificate` - Enable/disable server certificate verification.
* `vpc_id` - AWS VPC ID.
* `external_ip` - Configure GCP external IP.The structure of `external_ip` block is documented below.

The `external_ip` block contains:

* `name` - External IP name.
* `forwarding_rule` - Configure GCP forwarding rule.The structure of `forwarding_rule` block is documented below.

The `forwarding_rule` block contains:

* `rule_name` - Forwarding rule name.
* `target` - Target instance name.
* `gcp_project_list` - Configure GCP project list.The structure of `gcp_project_list` block is documented below.

The `gcp_project_list` block contains:

* `id` - GCP project ID.
* `gcp_zone_list` - Configure GCP zone list.The structure of `gcp_zone_list` block is documented below.

The `gcp_zone_list` block contains:

* `name` - GCP zone name.
* `nic` - Configure Azure network interface.The structure of `nic` block is documented below.

The `nic` block contains:

* `name` - Network interface name.
* `ip` - Configure IP configuration.The structure of `ip` block is documented below.

The `ip` block contains:

* `name` - IP configuration name.
* `public_ip` - Public IP name.
* `resource_group` - Resource group of Azure public IP.
* `route` - Configure GCP route.The structure of `route` block is documented below.

The `route` block contains:

* `name` - Route name.
* `route_table` - Configure Azure route table.The structure of `route_table` block is documented below.

The `route_table` block contains:

* `name` - Route table name.
* `resource_group` - Resource group of Azure route table.
* `subscription_id` - Subscription ID of Azure route table.
* `route` - Configure Azure route.The structure of `route` block is documented below.

The `route` block contains:

* `name` - Route name.
* `next_hop` - Next hop address.
* `server_list` - Server address list of the remote SDN connector.The structure of `server_list` block is documented below.

The `server_list` block contains:

* `ip` - IPv4 address.
