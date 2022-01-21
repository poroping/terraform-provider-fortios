---
subcategory: "FortiGate Dlp"
layout: "fortios"
page_title: "FortiOS: fortios_dlp_filepattern"
description: |-
  Configure file patterns used by DLP blocking.
---

## fortios_dlp_filepattern
Configure file patterns used by DLP blocking.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comment` - Optional comments.
* `id` - ID.
* `name` - Name of table containing the file pattern list.
* `entries` - Configure file patterns used by DLP blocking. The structure of `entries` block is documented below.

The `entries` block contains:

* `file_type` - Select a file type. Valid values: `7z` `arj` `cab` `lzh` `rar` `tar` `zip` `bzip` `gzip` `bzip2` `xz` `bat` `uue` `mime` `base64` `binhex` `elf` `exe` `hta` `html` `jad` `class` `cod` `javascript` `msoffice` `msofficex` `fsg` `upx` `petite` `aspack` `sis` `hlp` `activemime` `jpeg` `gif` `tiff` `png` `bmp` `unknown` `mpeg` `mov` `mp3` `wma` `wav` `pdf` `avi` `rm` `torrent` `hibun` `msi` `mach-o` `dmg` `.net` `xar` `chm` `iso` `crx` `flac` .
* `filter_type` - Filter by file name pattern or by file type. Valid values: `pattern` `type` .
* `pattern` - Add a file name pattern.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_dlp_filepattern can be imported using:
```sh
terraform import fortios_dlp_filepattern.labelname {{mkey}}
```
