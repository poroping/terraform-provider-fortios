// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: BGP VRF leaking table.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccRouterBgpVrf_basic(t *testing.T) {
	rName := "router_bgp_vrf"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			testAccCreateResourceFromExampleStep(rName),
		},
	})
}
