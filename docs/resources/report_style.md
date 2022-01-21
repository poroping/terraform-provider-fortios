---
subcategory: "FortiGate Report"
layout: "fortios"
page_title: "FortiOS: fortios_report_style"
description: |-
  Report style configuration.
---

## fortios_report_style
Report style configuration.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `align` - Alignment. Valid values: `left` `center` `right` `justify` .
* `bg_color` - Background color.
* `border_bottom` - Border bottom.
* `border_left` - Border left.
* `border_right` - Border right.
* `border_top` - Border top.
* `column_gap` - Column gap.
* `column_span` - Column span. Valid values: `none` `all` .
* `fg_color` - Foreground color.
* `font_family` - Font family. Valid values: `Verdana` `Arial` `Helvetica` `Courier` `Times` .
* `font_size` - Font size.
* `font_style` - Font style. Valid values: `normal` `italic` .
* `font_weight` - Font weight. Valid values: `normal` `bold` .
* `height` - Height.
* `line_height` - Text line height.
* `margin_bottom` - Margin bottom.
* `margin_left` - Margin left.
* `margin_right` - Margin right.
* `margin_top` - Margin top.
* `name` - Report style name.
* `options` - Report style options. Valid values: `font` `text` `color` `align` `size` `margin` `border` `padding` `column` .
* `padding_bottom` - Padding bottom.
* `padding_left` - Padding left.
* `padding_right` - Padding right.
* `padding_top` - Padding top.
* `width` - Width.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_report_style can be imported using:
```sh
terraform import fortios_report_style.labelname {{mkey}}
```
