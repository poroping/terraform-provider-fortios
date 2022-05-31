// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Access Proxy virtual hosts.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFirewallAccessProxyVirtualHost_basic(t *testing.T) {
	rName := "firewall_access_proxy_virtual_host"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			testAccCreateResourceFromExampleStep(rName),
		},
	})
}
