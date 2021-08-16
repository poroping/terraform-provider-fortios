---
layout: "fortios"
page_title: "Provider: FortiOS"
sidebar_current: "docs-fortios-index"
description: |-
  The FortiOS provider interacts with FortiGate.
---

# FortiOS Provider

->	This is a fork of the Official FortiOS provider with a bunch of added functionality and bug fixes. It is a drop-in replacement with one breaking change related to the dynamic_sort_subtable attribute.

If you have used this attribute you just need to change it from a string to a bool.

The FortiOS provider is used to interact with the resources supported by FortiOS. We need to configure the provider with the proper credentials before it can be used.

## Configuration for FortiGate

### Example Usage

```hcl
# Configure the FortiOS Provider for FortiGate
provider "fortios" {
  hostname     = "192.168.52.177"
  token        = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
  insecure     = "false"
  cabundlefile = "/path/yourCA.crt"
}

# Create a Static Route Item
resource "fortios_networking_route_static" "test1" {
  dst     = "110.2.2.122/32"
  gateway = "2.2.2.2"
  # ...
}
```

If it is used for testing, you can set `insecure` to "true" and unset `cabundlefile` to quickly set the provider up, for example:

```hcl
provider "fortios" {
  hostname = "192.168.52.177"
  token    = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
  insecure = "true"
}
```

Please refer to the Argument Reference below for more help on `insecure` and `cabundlefile`.

### Authentication

The FortiOS provider offers a means of providing credentials for authentication. The following methods are supported:

- Static credentials
- Environment variables

#### Static credentials

Static credentials can be provided by adding a `token` key in-line in the FortiOS provider block.

Usage:

```hcl
provider "fortios" {
  hostname     = "192.168.52.177"
  token        = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
  insecure     = "false"
  cabundlefile = "/path/yourCA.crt"
}
```

#### Generate an API token for FortiOS

See the left navigation: `Guides` -> `Generate an API token for FortiOS`.

#### Environment variables

You can provide your credentials via the `FORTIOS_ACCESS_HOSTNAME`, `FORTIOS_ACCESS_TOKEN`, `FORTIOS_INSECURE` and `FORTIOS_CA_CABUNDLE` environment variables. Note that setting your FortiOS credentials using static credentials variables will override the environment variables.

Usage:

```shell
$ export "FORTIOS_ACCESS_HOSTNAME"="192.168.52.177"
$ export "FORTIOS_INSECURE"="false"
$ export "FORTIOS_ACCESS_TOKEN"="09m441wrwc10yGyrtQ4nk6mjbqcfz9"
$ export "FORTIOS_CA_CABUNDLE"="/path/yourCA.crt"
```

Then configure the FortiOS Provider as following:

```hcl
provider "fortios" {}

# Create a Static Route Item
resource "fortios_networking_route_static" "test1" {
  dst       = "110.2.2.122/32"
  gateway   = "2.2.2.2"
  blackhole = "disable"
  distance  = "22"
  weight    = "3"
  # …
}
```

### VDOM

If the FortiGate unit is running in VDOM mode, the `vdom` configuration needs to be added.

Usage:

```hcl
provider "fortios" {
  hostname     = "192.168.52.177"
  token        = "q3Hs49jxts195gkd9Hjsxnjtmr6k39"
  insecure     = "false"
  cabundlefile = "/path/yourCA.crt"
  vdom         = "vdomtest"
}

resource "fortios_networking_route_static" "test1" {
  dst       = "120.2.2.122/32"
  gateway   = "2.2.2.2"
  blackhole = "disable"
  distance  = "22"
  weight    = "3"
  priority  = "3"
  device    = "lbforvdomtest"
  comment   = "Terraform test"
}
```

By default, each resource inherits the provider's global vdom settings, but it can also set its own vdom through the `vdomparam` of each resource. See the `vdomparam` argument of each resource for details.


### Argument Reference

The following arguments are supported:

* `hostname` - (Optional) The hostname or IP address of FortiOS unit. It must be provided, but it can also be sourced from the `FORTIOS_ACCESS_HOSTNAME` environment variable.

* `token` - (Optional) The token of FortiOS unit. It must be provided, but it can also be sourced from the `FORTIOS_ACCESS_TOKEN` environment variable.

* `insecure` - (Optional) Control whether the Provider to perform insecure SSL requests. If omitted, the `FORTIOS_INSECURE` environment variable is used. If neither is set, default value is `false`.

* `cabundlefile` - (Optional) The path of a custom CA bundle file. You can specify a path to the file, or you can specify it by the `FORTIOS_CA_CABUNDLE` environment variable.

* `vdom` - (Optional) If the FortiGate unit is running in VDOM mode, you can use this argument to specify the name of the vdom to be set .


## Release
Check out the FortiOS provider release notes and additional information from: [the FortiOS provider releases](https://github.com/poroping/terraform-provider-fortios/releases).

## Versioning

This forked provider has been tested against FortiOS 6.4, 7.0 versions, the configuration of all parameters should be based on the relevant FortiOS version manual.