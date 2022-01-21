---
subcategory: "FortiGate Report"
layout: "fortios"
page_title: "FortiOS: fortios_report_theme"
description: |-
  Get information on a fortios Report themes configuration
---

# Data Source: fortios_report_theme
Use this data source to get information on a fortios Report themes configuration


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Report theme name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `bullet_list_style` - Bullet list style.
* `column_count` - Report page column count.
* `default_html_style` - Default HTML report style.
* `default_pdf_style` - Default PDF report style.
* `graph_chart_style` - Graph chart style.
* `heading1_style` - Report heading style.
* `heading2_style` - Report heading style.
* `heading3_style` - Report heading style.
* `heading4_style` - Report heading style.
* `hline_style` - Horizontal line style.
* `image_style` - Image style.
* `name` - Report theme name.
* `normal_text_style` - Normal text style.
* `numbered_list_style` - Numbered list style.
* `page_footer_style` - Report page footer style.
* `page_header_style` - Report page header style.
* `page_orient` - Report page orientation.
* `page_style` - Report page style.
* `report_subtitle_style` - Report subtitle style.
* `report_title_style` - Report title style.
* `table_chart_caption_style` - Table chart caption style.
* `table_chart_even_row_style` - Table chart even row style.
* `table_chart_head_style` - Table chart head row style.
* `table_chart_odd_row_style` - Table chart odd row style.
* `table_chart_style` - Table chart style.
* `toc_heading1_style` - Table of contents heading style.
* `toc_heading2_style` - Table of contents heading style.
* `toc_heading3_style` - Table of contents heading style.
* `toc_heading4_style` - Table of contents heading style.
* `toc_title_style` - Table of contents title style.
