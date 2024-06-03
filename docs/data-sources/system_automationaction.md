---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_automationaction"
description: |-
  Get information on a fortios Action for automation stitches.
---

# Data Source: fortios_system_automationaction
Use this data source to get information on a fortios Action for automation stitches.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `accprofile` - Access profile for CLI script action to access FortiGate features.
* `action_type` - Action type.
* `alicloud_access_key_id` - AliCloud AccessKey ID.
* `alicloud_access_key_secret` - AliCloud AccessKey secret.
* `alicloud_account_id` - AliCloud account ID.
* `alicloud_function` - AliCloud function name.
* `alicloud_function_authorization` - AliCloud function authorization type.
* `alicloud_function_domain` - AliCloud function domain.
* `alicloud_region` - AliCloud region.
* `alicloud_service` - AliCloud service name.
* `alicloud_version` - AliCloud version.
* `aws_api_id` - AWS API Gateway ID.
* `aws_api_key` - AWS API Gateway API key.
* `aws_api_path` - AWS API Gateway path.
* `aws_api_stage` - AWS API Gateway deployment stage name.
* `aws_domain` - AWS domain.
* `aws_region` - AWS region.
* `azure_api_key` - Azure function API key.
* `azure_app` - Azure function application name.
* `azure_domain` - Azure function domain.
* `azure_function` - Azure function name.
* `azure_function_authorization` - Azure function authorization level.
* `delay` - Delay before execution (in seconds).
* `description` - Description.
* `email_body` - Email body.
* `email_from` - Email sender name.
* `email_subject` - Email subject.
* `execute_security_fabric` - Enable/disable execution of CLI script on all or only one FortiGate unit in the Security Fabric.
* `gcp_function` - Google Cloud function name.
* `gcp_function_domain` - Google Cloud function domain.
* `gcp_function_region` - Google Cloud function region.
* `gcp_project` - Google Cloud Platform project name.
* `http_body` - Request body (if necessary). Should be serialized json string.
* `message` - Message content.
* `message_type` - Message type.
* `method` - Request method (POST, PUT, GET, PATCH or DELETE).
* `minimum_interval` - Limit execution to no more than once in this interval (in seconds).
* `name` - Name.
* `output_size` - Number of megabytes to limit script output to (1 - 1024, default = 10).
* `port` - Protocol port.
* `protocol` - Request protocol.
* `replacement_message` - Enable/disable replacement message.
* `replacemsg_group` - Replacement message group.
* `required` - Required in action chain.
* `script` - CLI script.
* `security_tag` - NSX security tag.
* `system_action` - System action type.
* `timeout` - Maximum running time for this script in seconds (0 = no timeout).
* `tls_certificate` - Custom TLS certificate for API request.
* `uri` - Request API URI.
* `verify_host_cert` - Enable/disable verification of the remote host certificate.
* `email_to` - Email addresses.The structure of `email_to` block is documented below.

The `email_to` block contains:

* `name` - Email address.
* `headers` - Request headers.The structure of `headers` block is documented below.

The `headers` block contains:

* `header` - Request header.
* `http_headers` - Request headers.The structure of `http_headers` block is documented below.

The `http_headers` block contains:

* `id` - Entry ID.
* `key` - Request header key.
* `value` - Request header value.
* `sdn_connector` - NSX SDN connector names.The structure of `sdn_connector` block is documented below.

The `sdn_connector` block contains:

* `name` - SDN connector name.
