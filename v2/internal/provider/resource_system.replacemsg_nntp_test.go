// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Replacement messages.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSystemreplacemsgNntp_basic(t *testing.T) {
	rName := "system.replacemsg_nntp"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			testAccCreateResourceFromExampleStep(rName),
		},
	})
}
