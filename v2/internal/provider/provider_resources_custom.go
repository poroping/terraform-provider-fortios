package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func providerResourcesCustom() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"fortios_firewall_policy_sort":          resourceFirewallPolicySort(),
		"fortios_firewall_policy6_sort":         resourceFirewallPolicy6Sort(),
		"fortios_firewall_policy46_sort":        resourceFirewallPolicy46Sort(),
		"fortios_firewall_policy64_sort":        resourceFirewallPolicy64Sort(),
		"fortios_firewall_security_policy_sort": resourceFirewallSecurityPolicySort(),
		"fortios_firewall_proxy_policy_sort":    resourceFirewallProxyPolicySort(),
		"fortios_firewall_shaping_policy_sort":  resourceFirewallShapingPolicySort(),
		"fortios_certificate_management_local":  resourceCertificateManagementLocal(),
		"fortios_certificate_management_remote": resourceCertificateManagementRemote(),
		"fortios_generic_api":                   resourceGenericApi(),
	}
}
