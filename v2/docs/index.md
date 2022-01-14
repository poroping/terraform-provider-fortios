---
layout: "fortios"
page_title: "Provider: FortiOS"
sidebar_current: "docs-fortios-index"
description: |-
  The FortiOS provider interacts with a FortiGate.
---

# FortiOS Provider

## Why use over official FortiOS Provider?

- Ability to use CIDR notation everywhere instead of IP MASK. Makes interoperability much simpler.
- Includes validation at the resource schema level. Fail at Plan, not Apply!
- Deals with changes made 'out-of-band'. Resources deleted outside of Terraform will no longer cause an error and will be recreated/state updated as needed.
- Easier to onboard brownfield deployments through the use of allow_append. This attribute will perform a check if the resource already exists upon resource creation and just deal with it.
- Custom DiffSuppression functions to deal with some formatting quirks.
- Uses SDK written with models instead of interface{} everywhere. Why does that matter? Speed!
- New SDK implements better Error reporting. No more HTTP500 everywhere!
- Separate resources for tables within Complex resources (router/sdwan). Read the usage caveats!
- Implements Plugin SDK v2 allowing errors to 'bubble up' and returned as diags.
- Detect current FortiOS version on client setup and generates warnings if fields are used that are not supported in that version.
- Generated from templates from FortiOS API schema so should be able to implement new versions quickly.
- Fixes bugs that have been present in the official provider for some time.

## Coming from official FortiOS Provider?

Effort has been made to make this a drop-in replacement. However there may be some gotchas.

- Some old resources are missing, assuming these are from 6.0 track which is EOL anyway. But there maybe some other endpoints that were missed.
- This provider will enforce network addresses for resources such as router_bgp_network. Even though FortiOS API will convert on the fly. This will fail during planning phase and should be obvious and easy to resolve.
- Sorting resources were re-written. Raise an Issue on GitHub if this is something that needs to be backported.
- PKI auth has not been tested.

### Example Usage

```hcl
# Configure the FortiOS Provider for FortiGate
provider "fortios" {
  hostname     = "192.168.1.99"
  token        = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
  vdom         = "root"
  insecure     = "true"
}

# Create a Static Route
resource "fortios_router_static" "example" {
  dst       = "110.2.2.122/32"
  blackhole = "enable"
  # …
}
```

Please refer to the Argument Reference below for more help on additional arguments.

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
  insecure     = "true"
}
```

#### Generate an API token for FortiOS

See the left navigation: `Guides` -> `Generate an API token for FortiOS`.

#### Environment variables

You can provide your credentials via environment variables. 

->Setting FortiOS credentials statically will override the environment variables.

Usage:

```shell
$ export FORTIOS_ACCESS_HOSTNAME=192.168.52.177
$ export FORTIOS_INSECURE=true
$ export FORTIOS_ACCESS_TOKEN=09m441wrwc10yGyrtQ4nk6mjbqcfz9
$ export FORTIOS_VDOM=root
```

Then configure the FortiOS Provider as following:

```hcl
provider "fortios" {}

# Create a Static Route
resource "fortios_router_static" "example" {
  dst       = "110.2.2.122/32"
  blackhole = "enable"
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

* `hostname` - The hostname/IP address of the FortiOS to be connected. If omitted, `FORTIOS_ACCESS_HOSTNAME` environment variable is used.

* `token` - (Optional) The FortiOS API access token. If omitted, `FORTIOS_ACCESS_TOKEN` environment variable is used.

* `insecure` - (Optional) Control whether the Provider to perform insecure SSL requests. If omitted, the `FORTIOS_INSECURE` environment variable is used. Default value is `false`.

* `cabundlefile` - (Optional) The path of a custom CA bundle file. If omitted, `FORTIOS_CA_CABUNDLE` environment variable is used.

* `vdom` - (Optional) Default VDOM for API calls. If omitted, `FORTIOS_VDOM` environment variable is used.

* `auto_version` - (Optional) Will determine FortiOS version on client creation. If omitted, `FORTIOS_AUTO_VERSION` environment variable is used. Default value is `true`.

* `os_version` - (Optional) Manually set the FortiOS version. Will overwrite auto_version. If omitted, `FORTIOS_OS_VERSION` environment variable is used.

* `peerauth` - (Optional) Use PKI authentication. If omitted, `FORTIOS_CA_PEERAUTH` environment variable is used.

* `cacert` - (Optional) CA certificate for PKI authentication. If omitted, `FORTIOS_CA_CACERT` environment variable is used.

* `clientcert` - (Optional) Client certificate for PKI authentication. If omitted, `FORTIOS_CA_CLIENTCERT` environment variable is used.

* `clientkey` - (Optional) Client key for PKI authentication. If omitted, `FORTIOS_CA_CLIENTKEY` environment variable is used.


## Found a bug or want to Contribute?
Help is welcome!

See `Guides` -> `Debugging Provider` for information on how to provide good debug logs.

[GitHub](https://github.com/poroping/terraform-provider-fortios/issues).

## Versioning

The provider was generated from schemas covering 6.2, 6.4 and 7.0. Acceptance testing is performed with the latest GA release.