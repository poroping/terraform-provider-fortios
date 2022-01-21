---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomexception"
description: |-
  Global configuration objects that can be configured independently across different ha peers for all VDOMs or for the defined VDOM scope.
---

## fortios_system_vdomexception
Global configuration objects that can be configured independently across different ha peers for all VDOMs or for the defined VDOM scope.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `id` - Index <1-4096>.
* `object` - Name of the configuration object that can be configured independently for all VDOMs. Valid values: `log.fortianalyzer.setting` `log.fortianalyzer.override-setting` `log.fortianalyzer2.setting` `log.fortianalyzer2.override-setting` `log.fortianalyzer3.setting` `log.fortianalyzer3.override-setting` `log.fortianalyzer-cloud.setting` `log.fortianalyzer-cloud.override-setting` `log.syslogd.setting` `log.syslogd.override-setting` `log.syslogd2.setting` `log.syslogd2.override-setting` `log.syslogd3.setting` `log.syslogd3.override-setting` `log.syslogd4.setting` `log.syslogd4.override-setting` `system.gre-tunnel` `system.central-management` `system.csf` `user.radius` `system.cluster-sync` `system.standalone-cluster` `system.interface` `vpn.ipsec.phase1-interface` `vpn.ipsec.phase2-interface` `router.bgp` `router.route-map` `router.prefix-list` `firewall.ippool` `firewall.ippool6` `router.static` `router.static6` `firewall.vip` `firewall.vip6` `system.sdwan` `system.saml` `router.policy` `router.policy6` .
* `scope` - Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration. Valid values: `all` `inclusive` `exclusive` .
* `vdom` - Names of the VDOMs. The structure of `vdom` block is documented below.

The `vdom` block contains:

* `name` - VDOM name. This attribute must reference one of the following datasources: `system.vdom.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_vdomexception can be imported using:
```sh
terraform import fortios_system_vdomexception.labelname {{mkey}}
```
