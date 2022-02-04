---
subcategory: "FortiGate Cifs"
layout: "fortios"
page_title: "FortiOS: fortios_cifs_profile"
description: |-
  Configure CIFS profile.
---

## fortios_cifs_profile
Configure CIFS profile.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `domain_controller` - Domain for which to decrypt CIFS traffic. This attribute must reference one of the following datasources: `credential-store.domain-controller.server-name` .
* `name` - Profile name.
* `server_credential_type` - CIFS server credential type. Valid values: `none` `credential-replication` `credential-keytab` .
* `file_filter` - File filter. The structure of `file_filter` block is documented below.

The `file_filter` block contains:

* `log` - Enable/disable file filter logging. Valid values: `enable` `disable` .
* `status` - Enable/disable file filter. Valid values: `enable` `disable` .
* `entries` - File filter entries. The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Action taken for matched file. Valid values: `log` `block` .
* `comment` - Comment.
* `direction` - Match files transmitted in the session's originating or reply direction. Valid values: `incoming` `outgoing` `any` .
* `filter` - Add a file filter.
* `file_type` - Select file type. The structure of `file_type` block is documented below.

The `file_type` block contains:

* `name` - File type name. This attribute must reference one of the following datasources: `antivirus.filetype.name` .
* `server_keytab` - Server keytab. The structure of `server_keytab` block is documented below.

The `server_keytab` block contains:

* `keytab` - Base64 encoded keytab file containing credential of the server.
* `principal` - Service principal. For example, host/cifsserver.example.com@example.com.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_cifs_profile can be imported using:
```sh
terraform import fortios_cifs_profile.labelname {{mkey}}
```
