package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// Aliases to resources with different naming conventions. For backwards compatibility.
func providerDataSourcesAliases() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"fortios_routerbgp_neighbor":         dataSourceRouterBgpNeighbor(),
		"fortios_routerbgp_network":          dataSourceRouterBgpNetwork(),
		"fortios_routerbgp_network6":         dataSourceRouterBgpNetwork6(),
		"fortios_routerbgp_neighbor_group":   dataSourceRouterBgpNeighborGroup(),
		"fortios_routerbgp_neighbor_range":   dataSourceRouterBgpNeighborRange(),
		"fortios_routerospf_ospfinterface":   dataSourceRouterOspfOspfInterface(),
		"fortios_routerospf6_ospf6interface": dataSourceRouterOspf6Ospf6Interface(),
		"fortios_routerospf_network":         dataSourceRouterOspfNetwork(),
		"fortios_routerospf_neighbor":        dataSourceRouterOspfNeighbor(),
		"fortios_routerospf_area":            dataSourceRouterOspfArea(),
	}
}
