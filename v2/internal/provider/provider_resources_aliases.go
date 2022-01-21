package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// Aliases to resources with different naming conventions. For backwards compatibility.
func providerResourcesAliases() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"fortios_routerbgp_neighbor":                 resourceRouterBgpNeighbor(),
		"fortios_routerbgp_network":                  resourceRouterBgpNetwork(),
		"fortios_routerbgp_network6":                 resourceRouterBgpNetwork6(),
		"fortios_routerbgp_neighbor_group":           resourceRouterBgpNeighborGroup(),
		"fortios_routerbgp_neighbor_range":           resourceRouterBgpNeighborRange(),
		"fortios_routerospf_ospfinterface":           resourceRouterOspfOspfInterface(),
		"fortios_routerospf6_ospf6interface":         resourceRouterOspf6Ospf6Interface(),
		"fortios_routerospf_network":                 resourceRouterOspfNetwork(),
		"fortios_routerospf_neighbor":                resourceRouterOspfNeighbor(),
		"fortios_routerospf_area":                    resourceRouterOspfArea(),
		"fortios_json_generic_api":                   resourceGenericApi(),
		"fortios_firewall_DoSpolicy":                 resourceFirewallDoSPolicy(),
		"fortios_firewall_DoSpolicy6":                resourceFirewallDoSPolicy6(),
		"fortios_firewall_access_proxy":              resourceFirewallAccessProxy(),
		"fortios_firewall_access_proxy_virtual_host": resourceFirewallAccessProxyVirtualHost(),
	}
}
