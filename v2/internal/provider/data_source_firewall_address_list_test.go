package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceFirewallAddressList_basic(t *testing.T) {
	rName := "firewall_address_list"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			testAccCreateDataSourceFromExampleStep(rName),
		},
	})
}
