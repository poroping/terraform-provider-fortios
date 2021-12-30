// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure redundant internet connections using SD-WAN (formerly virtual WAN link).

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSystemVirtualWanLink_basic(t *testing.T) {
	rName := "system_virtual_wan_link"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			testAccCreateResourceFromExampleStep(rName),
		},
	})
}
