---
subcategory: "FortiGate Report"
layout: "fortios"
page_title: "FortiOS: fortios_report_chart"
description: |-
  Get information on a fortios Report chart widget configuration.
---

# Data Source: fortios_report_chart
Use this data source to get information on a fortios Report chart widget configuration.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Chart Widget Name
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `background` - Chart background.
* `category` - Category.
* `color_palette` - Color palette (system will pick color automatically by default).
* `comments` - Comment.
* `dataset` - Bind dataset to chart.
* `dimension` - Dimension.
* `favorite` - Favorite.
* `graph_type` - Graph type.
* `legend` - Enable/Disable Legend area.
* `legend_font_size` - Font size of legend area.
* `name` - Chart Widget Name
* `period` - Time period.
* `policy` - Used by monitor policy.
* `style` - Style.
* `title` - Chart title.
* `title_font_size` - Font size of chart title.
* `type` - Chart type.
* `category_series` - Category series of pie chart.The structure of `category_series` block is documented below.

The `category_series` block contains:

* `databind` - Category series value expression.
* `font_size` - Font size of category-series title.
* `column` - Table column definition.The structure of `column` block is documented below.

The `column` block contains:

* `detail_unit` - Detail unit of column.
* `detail_value` - Detail value of column.
* `footer_unit` - Footer unit of column.
* `footer_value` - Footer value of column.
* `header_value` - Display name of table header.
* `id` - ID.
* `mapping` - Show detail in certain display value for certain condition.The structure of `mapping` block is documented below.

The `mapping` block contains:

* `displayname` - Display name.
* `id` - id
* `op` - Comparision operater.
* `value_type` - Value type.
* `value1` - Value 1.
* `value2` - Value 2.
* `drill_down_charts` - Drill down charts.The structure of `drill_down_charts` block is documented below.

The `drill_down_charts` block contains:

* `chart_name` - Drill down chart name.
* `id` - Drill down chart ID.
* `status` - Enable/disable this drill down chart.
* `value_series` - Value series of pie chart.The structure of `value_series` block is documented below.

The `value_series` block contains:

* `databind` - Value series value expression.
* `x_series` - X-series of chart.The structure of `x_series` block is documented below.

The `x_series` block contains:

* `caption` - X-series caption.
* `caption_font_size` - X-series caption font size.
* `databind` - X-series value expression.
* `font_size` - X-series label font size.
* `is_category` - X-series represent category or not.
* `label_angle` - X-series label angle.
* `scale_direction` - Scale increase or decrease.
* `scale_format` - Date/time format.
* `scale_step` - Scale step.
* `scale_unit` - Scale unit.
* `unit` - X-series unit.
* `y_series` - Y-series of chart.The structure of `y_series` block is documented below.

The `y_series` block contains:

* `caption` - Y-series caption.
* `caption_font_size` - Y-series caption font size.
* `databind` - Y-series value expression.
* `extra_databind` - Extra Y-series value.
* `extra_y` - Allow another Y-series value
* `extra_y_legend` - Extra Y-series legend type/name.
* `font_size` - Y-series label font size.
* `group` - Y-series group option.
* `label_angle` - Y-series label angle.
* `unit` - Y-series unit.
* `y_legend` - First Y-series legend type/name.
