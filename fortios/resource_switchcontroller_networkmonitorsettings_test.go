// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSSwitchControllerNetworkMonitorSettings_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSwitchControllerNetworkMonitorSettings_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSwitchControllerNetworkMonitorSettingsConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSwitchControllerNetworkMonitorSettingsExists("fortios_switchcontroller_networkmonitorsettings.trname"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_networkmonitorsettings.trname", "network_monitoring", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSwitchControllerNetworkMonitorSettingsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SwitchControllerNetworkMonitorSettings: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SwitchControllerNetworkMonitorSettings is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerNetworkMonitorSettings(i, "root", make(map[string][]string), 0)

		if err != nil {
			return fmt.Errorf("error reading SwitchControllerNetworkMonitorSettings: %s", err)
		}

		if o == nil {
			return fmt.Errorf("error creating SwitchControllerNetworkMonitorSettings: %s", n)
		}

		return nil
	}
}

func testAccCheckSwitchControllerNetworkMonitorSettingsDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_switchcontroller_networkmonitorsettings" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerNetworkMonitorSettings(i, "root", make(map[string][]string), 0)

		if err == nil {
			if o != nil {
				return fmt.Errorf("error SwitchControllerNetworkMonitorSettings %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSwitchControllerNetworkMonitorSettingsConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_switchcontroller_networkmonitorsettings" "trname" {
  network_monitoring = "disable"
}
`)
}
