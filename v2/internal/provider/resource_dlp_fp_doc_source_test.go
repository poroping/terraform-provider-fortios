// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Create a DLP fingerprint database by allowing the FortiGate to access a file server containing files from which to create fingerprints.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDlpFpDocSource_basic(t *testing.T) {
	rName := "dlp_fp_doc_source"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			testAccCreateResourceFromExampleStep(rName),
		},
	})
}
