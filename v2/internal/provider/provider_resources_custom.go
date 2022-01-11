package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func providerResourcesCustom() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"fortios_firewall_policy_sort": resourceFirewallPolicySort(),
	}
}
