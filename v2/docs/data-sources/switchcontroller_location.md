---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_location"
description: |-
  Get information on a fortios Configure FortiSwitch location services.
---

# Data Source: fortios_switchcontroller_location
Use this data source to get information on a fortios Configure FortiSwitch location services.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Unique location item name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - Unique location item name.
* `address_civic` - Configure location civic address.The structure of `address_civic` block is documented below.

The `address_civic` block contains:

* `additional` - Location additional details.
* `additional_code` - Location additional code details.
* `block` - Location block details.
* `branch_road` - Location branch road details.
* `building` - Location building details.
* `city` - Location city details.
* `city_division` - Location city division details.
* `country` - The two-letter ISO 3166 country code in capital ASCII letters eg. US, CA, DK, DE.
* `country_subdivision` - National subdivisions (state, canton, region, province, or prefecture).
* `county` - County, parish, gun (JP), or district (IN).
* `direction` - Leading street direction.
* `floor` - Floor.
* `landmark` - Landmark or vanity address.
* `language` - Language.
* `name` - Name (residence and office occupant).
* `number` - House number.
* `number_suffix` - House number suffix.
* `parent_key` - Parent key name.
* `place_type` - Placetype.
* `post_office_box` - Post office box (P.O. box).
* `postal_community` - Postal community name.
* `primary_road` - Primary road name.
* `road_section` - Road section.
* `room` - Room number.
* `script` - Script used to present the address information.
* `seat` - Seat number.
* `street` - Street.
* `street_name_post_mod` - Street name post modifier.
* `street_name_pre_mod` - Street name pre modifier.
* `street_suffix` - Street suffix.
* `sub_branch_road` - Sub branch road name.
* `trailing_str_suffix` - Trailing street suffix.
* `unit` - Unit (apartment, suite).
* `zip` - Postal/zip code.
* `coordinates` - Configure location GPS coordinates.The structure of `coordinates` block is documented below.

The `coordinates` block contains:

* `altitude` - +/- Floating point no. eg. 117.47.
* `altitude_unit` - m ( meters), f ( floors).
* `datum` - WGS84, NAD83, NAD83/MLLW.
* `latitude` - Floating point start with ( +/- )  or end with ( N or S ) eg. +/-16.67 or 16.67N.
* `longitude` - Floating point start with ( +/- )  or end with ( E or W ) eg. +/-26.789 or 26.789E.
* `parent_key` - Parent key name.
* `elin_number` - Configure location ELIN number.The structure of `elin_number` block is documented below.

The `elin_number` block contains:

* `elin_num` - Configure ELIN callback number.
* `parent_key` - Parent key name.
