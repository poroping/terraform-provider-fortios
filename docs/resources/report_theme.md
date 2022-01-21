---
subcategory: "FortiGate Report"
layout: "fortios"
page_title: "FortiOS: fortios_report_theme"
description: |-
  Report themes configuration
---

## fortios_report_theme
Report themes configuration

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `bullet_list_style` - Bullet list style.
* `column_count` - Report page column count. Valid values: `1` `2` `3` .
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
* `page_orient` - Report page orientation. Valid values: `portrait` `landscape` .
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

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_report_theme can be imported using:
```sh
terraform import fortios_report_theme.labelname {{mkey}}
```
