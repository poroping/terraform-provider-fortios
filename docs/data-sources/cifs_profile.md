---
subcategory: "FortiGate Cifs"
layout: "fortios"
page_title: "FortiOS: fortios_cifs_profile"
description: |-
  Get information on a fortios Configure CIFS profile.
---

# Data Source: fortios_cifs_profile
Use this data source to get information on a fortios Configure CIFS profile.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `domain_controller` - Domain for which to decrypt CIFS traffic.
* `name` - Profile name.
* `server_credential_type` - CIFS server credential type.
* `file_filter` - File filter.The structure of `file_filter` block is documented below.

The `file_filter` block contains:

* `log` - Enable/disable file filter logging.
* `status` - Enable/disable file filter.
* `entries` - File filter entries.The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Action taken for matched file.
* `comment` - Comment.
* `direction` - Match files transmitted in the session's originating or reply direction.
* `filter` - Add a file filter.
* `file_type` - Select file type.The structure of `file_type` block is documented below.

The `file_type` block contains:

* `name` - File type name.
* `server_keytab` - Server keytab.The structure of `server_keytab` block is documented below.

The `server_keytab` block contains:

* `keytab` - Base64 encoded keytab file containing credential of the server.
* `principal` - Service principal.  For example, "host/cifsserver.example.com@example.com".
