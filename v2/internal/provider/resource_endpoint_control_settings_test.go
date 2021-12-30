// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure endpoint control settings.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccEndpointControlSettings_basic(t *testing.T) {
	rName := "endpoint_control_settings"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			testAccCreateResourceFromExampleStep(rName),
		},
	})
}
