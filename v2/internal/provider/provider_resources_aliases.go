package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// Aliases to resources with different naming conventions. For backwards compatibility.
func providerResourcesAliases() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"fortios_vpnssl_settings_authentication_rule": resourceVpnsslsettingsAuthenticationRule(),
		"fortios_system_sdwan_duplication":            resourceSystemsdwanDuplication(),
		"fortios_system_sdwan_health_check":           resourceSystemsdwanHealthCheck(),
		"fortios_system_sdwan_members":                resourceSystemsdwanMembers(),
		"fortios_system_sdwan_neighbor":               resourceSystemsdwanNeighbor(),
		"fortios_system_sdwan_service":                resourceSystemsdwanService(),
		"fortios_system_sdwan_zone":                   resourceSystemsdwanZone(),
		"fortios_routerbgp_neighbor_group":            resourceRouterbgpNeighborGroup(),
		"fortios_routerbgp_neighbor_range":            resourceRouterbgpNeighborRange(),
		"fortios_firewall_access_proxy":               resourceFirewallAccessProxy(),
		"fortios_firewall_access_proxy_virtual_host":  resourceFirewallAccessProxyVirtualHost(),
	}
}
