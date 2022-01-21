---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_replacemsggroup"
description: |-
  Get information on a fortios Configure replacement message groups.
---

# Data Source: fortios_system_replacemsggroup
Use this data source to get information on a fortios Configure replacement message groups.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Group name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `group_type` - Group type.
* `name` - Group name.
* `admin` - Replacement message table entries.The structure of `admin` block is documented below.

The `admin` block contains:

* `buffer` - Message string.
* `format` - Format flag.
* `header` - Header flag.
* `msg_type` - Message type.
* `alertmail` - Replacement message table entries.The structure of `alertmail` block is documented below.

The `alertmail` block contains:

* `buffer` - Message string.
* `format` - Format flag.
* `header` - Header flag.
* `msg_type` - Message type.
* `auth` - Replacement message table entries.The structure of `auth` block is documented below.

The `auth` block contains:

* `buffer` - Message string.
* `format` - Format flag.
* `header` - Header flag.
* `msg_type` - Message type.
* `automation` - Replacement message table entries.The structure of `automation` block is documented below.

The `automation` block contains:

* `buffer` - Message string.
* `format` - Format flag.
* `header` - Header flag.
* `msg_type` - Message type.
* `custom_message` - Replacement message table entries.The structure of `custom_message` block is documented below.

The `custom_message` block contains:

* `buffer` - Message string.
* `format` - Format flag.
* `header` - Header flag.
* `msg_type` - Message type.
* `device_detection_portal` - Replacement message table entries.The structure of `device_detection_portal` block is documented below.

The `device_detection_portal` block contains:

* `buffer` - Message string.
* `format` - Format flag.
* `header` - Header flag.
* `msg_type` - Message type.
* `fortiguard_wf` - Replacement message table entries.The structure of `fortiguard_wf` block is documented below.

The `fortiguard_wf` block contains:

* `buffer` - Message string.
* `format` - Format flag.
* `header` - Header flag.
* `msg_type` - Message type.
* `ftp` - Replacement message table entries.The structure of `ftp` block is documented below.

The `ftp` block contains:

* `buffer` - Message string.
* `format` - Format flag.
* `header` - Header flag.
* `msg_type` - Message type.
* `http` - Replacement message table entries.The structure of `http` block is documented below.

The `http` block contains:

* `buffer` - Message string.
* `format` - Format flag.
* `header` - Header flag.
* `msg_type` - Message type.
* `icap` - Replacement message table entries.The structure of `icap` block is documented below.

The `icap` block contains:

* `buffer` - Message string.
* `format` - Format flag.
* `header` - Header flag.
* `msg_type` - Message type.
* `mail` - Replacement message table entries.The structure of `mail` block is documented below.

The `mail` block contains:

* `buffer` - Message string.
* `format` - Format flag.
* `header` - Header flag.
* `msg_type` - Message type.
* `nac_quar` - Replacement message table entries.The structure of `nac_quar` block is documented below.

The `nac_quar` block contains:

* `buffer` - Message string.
* `format` - Format flag.
* `header` - Header flag.
* `msg_type` - Message type.
* `nntp` - Replacement message table entries.The structure of `nntp` block is documented below.

The `nntp` block contains:

* `buffer` - Message string.
* `format` - Format flag.
* `header` - Header flag.
* `msg_type` - Message type.
* `spam` - Replacement message table entries.The structure of `spam` block is documented below.

The `spam` block contains:

* `buffer` - Message string.
* `format` - Format flag.
* `header` - Header flag.
* `msg_type` - Message type.
* `sslvpn` - Replacement message table entries.The structure of `sslvpn` block is documented below.

The `sslvpn` block contains:

* `buffer` - Message string.
* `format` - Format flag.
* `header` - Header flag.
* `msg_type` - Message type.
* `traffic_quota` - Replacement message table entries.The structure of `traffic_quota` block is documented below.

The `traffic_quota` block contains:

* `buffer` - Message string.
* `format` - Format flag.
* `header` - Header flag.
* `msg_type` - Message type.
* `utm` - Replacement message table entries.The structure of `utm` block is documented below.

The `utm` block contains:

* `buffer` - Message string.
* `format` - Format flag.
* `header` - Header flag.
* `msg_type` - Message type.
* `webproxy` - Replacement message table entries.The structure of `webproxy` block is documented below.

The `webproxy` block contains:

* `buffer` - Message string.
* `format` - Format flag.
* `header` - Header flag.
* `msg_type` - Message type.
