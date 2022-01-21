---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_automationaction"
description: |-
  Action for automation stitches.
---

## fortios_system_automationaction
Action for automation stitches.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `accprofile` - Access profile for CLI script action to access FortiGate features. This attribute must reference one of the following datasources: `system.accprofile.name` .
* `action_type` - Action type. Valid values: `email` `fortiexplorer-notification` `alert` `disable-ssid` `quarantine` `quarantine-forticlient` `quarantine-nsx` `quarantine-fortinac` `ban-ip` `aws-lambda` `azure-function` `google-cloud-function` `alicloud-function` `webhook` `cli-script` `slack-notification` `microsoft-teams-notification` .
* `alicloud_access_key_id` - AliCloud AccessKey ID.
* `alicloud_access_key_secret` - AliCloud AccessKey secret.
* `alicloud_account_id` - AliCloud account ID.
* `alicloud_function` - AliCloud function name.
* `alicloud_function_authorization` - AliCloud function authorization type. Valid values: `anonymous` `function` .
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
* `azure_function_authorization` - Azure function authorization level. Valid values: `anonymous` `function` `admin` .
* `delay` - Delay before execution (in seconds).
* `description` - Description.
* `email_body` - Email body.
* `email_from` - Email sender name.
* `email_subject` - Email subject.
* `execute_security_fabric` - Enable/disable execution of CLI script on all or only one FortiGate unit in the Security Fabric. Valid values: `enable` `disable` .
* `gcp_function` - Google Cloud function name.
* `gcp_function_domain` - Google Cloud function domain.
* `gcp_function_region` - Google Cloud function region.
* `gcp_project` - Google Cloud Platform project name.
* `http_body` - Request body (if necessary). Should be serialized json string.
* `message` - Message content.
* `message_type` - Message type. Valid values: `text` `json` .
* `method` - Request method (POST, PUT, GET, PATCH or DELETE). Valid values: `post` `put` `get` `patch` `delete` .
* `minimum_interval` - Limit execution to no more than once in this interval (in seconds).
* `name` - Name.
* `port` - Protocol port.
* `protocol` - Request protocol. Valid values: `http` `https` .
* `replacement_message` - Enable/disable replacement message. Valid values: `enable` `disable` .
* `replacemsg_group` - Replacement message group. This attribute must reference one of the following datasources: `system.replacemsg-group.name` .
* `required` - Required in action chain. Valid values: `enable` `disable` .
* `script` - CLI script.
* `security_tag` - NSX security tag.
* `tls_certificate` - Custom TLS certificate for API request. This attribute must reference one of the following datasources: `certificate.local.name` .
* `uri` - Request API URI.
* `verify_host_cert` - Enable/disable verification of the remote host certificate. Valid values: `enable` `disable` .
* `email_to` - Email addresses. The structure of `email_to` block is documented below.

The `email_to` block contains:

* `name` - Email address.
* `headers` - Request headers. The structure of `headers` block is documented below.

The `headers` block contains:

* `header` - Request header.
* `sdn_connector` - NSX SDN connector names. The structure of `sdn_connector` block is documented below.

The `sdn_connector` block contains:

* `name` - SDN connector name. This attribute must reference one of the following datasources: `system.sdn-connector.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_automationaction can be imported using:
```sh
terraform import fortios_system_automationaction.labelname {{mkey}}
```
