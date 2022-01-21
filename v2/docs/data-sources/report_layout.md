---
subcategory: "FortiGate Report"
layout: "fortios"
page_title: "FortiOS: fortios_report_layout"
description: |-
  Get information on a fortios Report layout configuration.
---

# Data Source: fortios_report_layout
Use this data source to get information on a fortios Report layout configuration.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Report layout name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `cutoff_option` - Cutoff-option is either run-time or custom.
* `cutoff_time` - Custom cutoff time to generate report [hh:mm].
* `day` - Schedule days of week to generate report.
* `description` - Description.
* `email_recipients` - Email recipients for generated reports.
* `email_send` - Enable/disable sending emails after reports are generated.
* `format` - Report format.
* `max_pdf_report` - Maximum number of PDF reports to keep at one time (oldest report is overwritten).
* `name` - Report layout name.
* `options` - Report layout options.
* `schedule_type` - Report schedule type.
* `style_theme` - Report style theme.
* `subtitle` - Report subtitle.
* `time` - Schedule time to generate report [hh:mm].
* `title` - Report title.
* `body_item` - Configure report body item.The structure of `body_item` block is documented below.

The `body_item` block contains:

* `chart` - Report item chart name.
* `chart_options` - Report chart options.
* `column` - Report section column number.
* `content` - Report item text content.
* `description` - Description.
* `drill_down_items` - Control how drill down charts are shown.
* `drill_down_types` - Control whether keys from the parent being combined or not.
* `hide` - Enable/disable hide item in report.
* `id` - Report item ID.
* `img_src` - Report item image file name.
* `list_component` - Report item list component.
* `misc_component` - Report item miscellaneous component.
* `style` - Report item style.
* `table_caption_style` - Table chart caption style.
* `table_column_widths` - Report item table column widths.
* `table_even_row_style` - Table chart even row style.
* `table_head_style` - Table chart head style.
* `table_odd_row_style` - Table chart odd row style.
* `text_component` - Report item text component.
* `title` - Report section title.
* `top_n` - Value of top.
* `type` - Report item type.
* `list` - Configure report list item.The structure of `list` block is documented below.

The `list` block contains:

* `content` - List entry content.
* `id` - List entry ID.
* `parameters` - Parameters.The structure of `parameters` block is documented below.

The `parameters` block contains:

* `id` - ID.
* `name` - Field name that match field of parameters defined in dataset.
* `value` - Value to replace corresponding field of parameters defined in dataset.
* `page` - Configure report page.The structure of `page` block is documented below.

The `page` block contains:

* `column_break_before` - Report page auto column break before heading.
* `options` - Report page options.
* `page_break_before` - Report page auto page break before heading.
* `paper` - Report page paper.
* `footer` - Configure report page footer.The structure of `footer` block is documented below.

The `footer` block contains:

* `style` - Report footer style.
* `footer_item` - Configure report footer item.The structure of `footer_item` block is documented below.

The `footer_item` block contains:

* `content` - Report item text content.
* `description` - Description.
* `id` - Report item ID.
* `img_src` - Report item image file name.
* `style` - Report item style.
* `type` - Report item type.
* `header` - Configure report page header.The structure of `header` block is documented below.

The `header` block contains:

* `style` - Report header style.
* `header_item` - Configure report header item.The structure of `header_item` block is documented below.

The `header_item` block contains:

* `content` - Report item text content.
* `description` - Description.
* `id` - Report item ID.
* `img_src` - Report item image file name.
* `style` - Report item style.
* `type` - Report item type.
