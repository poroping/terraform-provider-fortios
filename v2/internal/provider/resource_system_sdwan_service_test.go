// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSystemSdwanService_basic(t *testing.T) {
	rName := "system_sdwan_service"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			testAccCreateResourceFromExampleStep(rName),
		},
	})
}
