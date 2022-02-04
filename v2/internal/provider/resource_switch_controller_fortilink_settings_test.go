// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure integrated FortiLink settings for FortiSwitch.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSwitchControllerFortilinkSettings_basic(t *testing.T) {
	rName := "switch_controller_fortilink_settings"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			testAccCreateResourceFromExampleStep(rName),
		},
	})
}