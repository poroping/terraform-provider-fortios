---
subcategory: "FortiGate Report"
layout: "fortios"
page_title: "FortiOS: fortios_report_chart"
description: |-
  Report chart widget configuration.
---

## fortios_report_chart
Report chart widget configuration.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `background` - Chart background.
* `category` - Category. Valid values: `misc` `traffic` `event` `virus` `webfilter` `attack` `spam` `dlp` `app-ctrl` `vulnerability` .
* `color_palette` - Color palette (system will pick color automatically by default).
* `comments` - Comment.
* `dataset` - Bind dataset to chart.
* `dimension` - Dimension. Valid values: `2D` `3D` .
* `favorite` - Favorite. Valid values: `no` `yes` .
* `graph_type` - Graph type. Valid values: `none` `bar` `pie` `line` `flow` .
* `legend` - Enable/Disable Legend area. Valid values: `enable` `disable` .
* `legend_font_size` - Font size of legend area.
* `name` - Chart Widget Name
* `period` - Time period. Valid values: `last24h` `last7d` .
* `policy` - Used by monitor policy.
* `style` - Style. Valid values: `auto` `manual` .
* `title` - Chart title.
* `title_font_size` - Font size of chart title.
* `type` - Chart type. Valid values: `graph` `table` .
* `category_series` - Category series of pie chart. The structure of `category_series` block is documented below.

The `category_series` block contains:

* `databind` - Category series value expression.
* `font_size` - Font size of category-series title.
* `column` - Table column definition. The structure of `column` block is documented below.

The `column` block contains:

* `detail_unit` - Detail unit of column.
* `detail_value` - Detail value of column.
* `footer_unit` - Footer unit of column.
* `footer_value` - Footer value of column.
* `header_value` - Display name of table header.
* `id` - ID.
* `mapping` - Show detail in certain display value for certain condition. The structure of `mapping` block is documented below.

The `mapping` block contains:

* `displayname` - Display name.
* `id` - id
* `op` - Comparision operater. Valid values: `none` `greater` `greater-equal` `less` `less-equal` `equal` `between` .
* `value_type` - Value type. Valid values: `integer` `string` .
* `value1` - Value 1.
* `value2` - Value 2.
* `drill_down_charts` - Drill down charts. The structure of `drill_down_charts` block is documented below.

The `drill_down_charts` block contains:

* `chart_name` - Drill down chart name.
* `id` - Drill down chart ID.
* `status` - Enable/disable this drill down chart. Valid values: `enable` `disable` .
* `value_series` - Value series of pie chart. The structure of `value_series` block is documented below.

The `value_series` block contains:

* `databind` - Value series value expression.
* `x_series` - X-series of chart. The structure of `x_series` block is documented below.

The `x_series` block contains:

* `caption` - X-series caption.
* `caption_font_size` - X-series caption font size.
* `databind` - X-series value expression.
* `font_size` - X-series label font size.
* `is_category` - X-series represent category or not. Valid values: `yes` `no` .
* `label_angle` - X-series label angle. Valid values: `45-degree` `vertical` `horizontal` .
* `scale_direction` - Scale increase or decrease. Valid values: `decrease` `increase` .
* `scale_format` - Date/time format. Valid values: `YYYY-MM-DD-HH-MM` `YYYY-MM-DD HH` `YYYY-MM-DD` `YYYY-MM` `YYYY` `HH-MM` `MM-DD` .
* `scale_step` - Scale step.
* `scale_unit` - Scale unit. Valid values: `minute` `hour` `day` `month` `year` .
* `unit` - X-series unit.
* `y_series` - Y-series of chart. The structure of `y_series` block is documented below.

The `y_series` block contains:

* `caption` - Y-series caption.
* `caption_font_size` - Y-series caption font size.
* `databind` - Y-series value expression.
* `extra_databind` - Extra Y-series value.
* `extra_y` - Allow another Y-series value Valid values: `enable` `disable` .
* `extra_y_legend` - Extra Y-series legend type/name.
* `font_size` - Y-series label font size.
* `group` - Y-series group option.
* `label_angle` - Y-series label angle. Valid values: `45-degree` `vertical` `horizontal` .
* `unit` - Y-series unit.
* `y_legend` - First Y-series legend type/name.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_report_chart can be imported using:
```sh
terraform import fortios_report_chart.labelname {{mkey}}
```
