---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_location"
description: |-
  Configure FortiSwitch location services.
---

## fortios_switchcontroller_location
Configure FortiSwitch location services.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `name` - Unique location item name.
* `address_civic` - Configure location civic address. The structure of `address_civic` block is documented below.

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
* `place_type` - Place type.
* `post_office_box` - Post office box.
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
* `coordinates` - Configure location GPS coordinates. The structure of `coordinates` block is documented below.

The `coordinates` block contains:

* `altitude` - Plus or minus floating point number. For example, 117.47.
* `altitude_unit` - Configure the unit for which the altitude is to (m = meters, f = floors of a building). Valid values: `m` `f` .
* `datum` - WGS84, NAD83, NAD83/MLLW. Valid values: `WGS84` `NAD83` `NAD83/MLLW` .
* `latitude` - Floating point starting with +/- or ending with (N or S). For example, +/-16.67 or 16.67N.
* `longitude` - Floating point starting with +/- or ending with (N or S). For example, +/-26.789 or 26.789E.
* `parent_key` - Parent key name.
* `elin_number` - Configure location ELIN number. The structure of `elin_number` block is documented below.

The `elin_number` block contains:

* `elin_num` - Configure ELIN callback number.
* `parent_key` - Parent key name.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_switchcontroller_location can be imported using:
```sh
terraform import fortios_switchcontroller_location.labelname {{mkey}}
```
