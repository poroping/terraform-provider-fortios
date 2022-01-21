---
subcategory: "FortiGate Report"
layout: "fortios"
page_title: "FortiOS: fortios_report_layout"
description: |-
  Report layout configuration.
---

## fortios_report_layout
Report layout configuration.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `cutoff_option` - Cutoff-option is either run-time or custom. Valid values: `run-time` `custom` .
* `cutoff_time` - Custom cutoff time to generate report [hh:mm].
* `day` - Schedule days of week to generate report. Valid values: `sunday` `monday` `tuesday` `wednesday` `thursday` `friday` `saturday` .
* `description` - Description.
* `email_recipients` - Email recipients for generated reports.
* `email_send` - Enable/disable sending emails after reports are generated. Valid values: `enable` `disable` .
* `format` - Report format. Valid values: `pdf` .
* `max_pdf_report` - Maximum number of PDF reports to keep at one time (oldest report is overwritten).
* `name` - Report layout name.
* `options` - Report layout options. Valid values: `include-table-of-content` `auto-numbering-heading` `view-chart-as-heading` `show-html-navbar-before-heading` `dummy-option` .
* `schedule_type` - Report schedule type. Valid values: `demand` `daily` `weekly` .
* `style_theme` - Report style theme.
* `subtitle` - Report subtitle.
* `time` - Schedule time to generate report [hh:mm].
* `title` - Report title.
* `body_item` - Configure report body item. The structure of `body_item` block is documented below.

The `body_item` block contains:

* `chart` - Report item chart name.
* `chart_options` - Report chart options. Valid values: `include-no-data` `hide-title` `show-caption` .
* `column` - Report section column number.
* `content` - Report item text content.
* `description` - Description.
* `drill_down_items` - Control how drill down charts are shown.
* `drill_down_types` - Control whether keys from the parent being combined or not.
* `hide` - Enable/disable hide item in report. Valid values: `enable` `disable` .
* `id` - Report item ID.
* `img_src` - Report item image file name.
* `list_component` - Report item list component. Valid values: `bullet` `numbered` .
* `misc_component` - Report item miscellaneous component. Valid values: `hline` `page-break` `column-break` `section-start` .
* `style` - Report item style.
* `table_caption_style` - Table chart caption style.
* `table_column_widths` - Report item table column widths.
* `table_even_row_style` - Table chart even row style.
* `table_head_style` - Table chart head style.
* `table_odd_row_style` - Table chart odd row style.
* `text_component` - Report item text component. Valid values: `text` `heading1` `heading2` `heading3` .
* `title` - Report section title.
* `top_n` - Value of top.
* `type` - Report item type. Valid values: `text` `image` `chart` `misc` .
* `list` - Configure report list item. The structure of `list` block is documented below.

The `list` block contains:

* `content` - List entry content.
* `id` - List entry ID.
* `parameters` - Parameters. The structure of `parameters` block is documented below.

The `parameters` block contains:

* `id` - ID.
* `name` - Field name that match field of parameters defined in dataset.
* `value` - Value to replace corresponding field of parameters defined in dataset.
* `page` - Configure report page. The structure of `page` block is documented below.

The `page` block contains:

* `column_break_before` - Report page auto column break before heading. Valid values: `heading1` `heading2` `heading3` .
* `options` - Report page options. Valid values: `header-on-first-page` `footer-on-first-page` .
* `page_break_before` - Report page auto page break before heading. Valid values: `heading1` `heading2` `heading3` .
* `paper` - Report page paper. Valid values: `a4` `letter` .
* `footer` - Configure report page footer. The structure of `footer` block is documented below.

The `footer` block contains:

* `style` - Report footer style.
* `footer_item` - Configure report footer item. The structure of `footer_item` block is documented below.

The `footer_item` block contains:

* `content` - Report item text content.
* `description` - Description.
* `id` - Report item ID.
* `img_src` - Report item image file name.
* `style` - Report item style.
* `type` - Report item type. Valid values: `text` `image` .
* `header` - Configure report page header. The structure of `header` block is documented below.

The `header` block contains:

* `style` - Report header style.
* `header_item` - Configure report header item. The structure of `header_item` block is documented below.

The `header_item` block contains:

* `content` - Report item text content.
* `description` - Description.
* `id` - Report item ID.
* `img_src` - Report item image file name.
* `style` - Report item style.
* `type` - Report item type. Valid values: `text` `image` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_report_layout can be imported using:
```sh
terraform import fortios_report_layout.labelname {{mkey}}
```
