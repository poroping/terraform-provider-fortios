// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.6,v7.0.2,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Spanning Tree Protocol (STP).

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSystemStp_basic(t *testing.T) {
	rName := "system_stp"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			testAccCreateResourceFromExampleStep(rName),
		},
	})
}
