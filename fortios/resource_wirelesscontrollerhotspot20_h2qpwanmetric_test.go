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

func TestAccFortiOSWirelessControllerHotspot20H2QpWanMetric_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWirelessControllerHotspot20H2QpWanMetric_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWirelessControllerHotspot20H2QpWanMetricConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWirelessControllerHotspot20H2QpWanMetricExists("fortios_wirelesscontrollerhotspot20_h2qpwanmetric.trname"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpwanmetric.trname", "downlink_load", "0"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpwanmetric.trname", "downlink_speed", "2400"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpwanmetric.trname", "link_at_capacity", "disable"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpwanmetric.trname", "link_status", "up"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpwanmetric.trname", "load_measurement_duration", "0"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpwanmetric.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpwanmetric.trname", "symmetric_wan_link", "symmetric"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpwanmetric.trname", "uplink_load", "0"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpwanmetric.trname", "uplink_speed", "2400"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWirelessControllerHotspot20H2QpWanMetricExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WirelessControllerHotspot20H2QpWanMetric: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WirelessControllerHotspot20H2QpWanMetric is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWirelessControllerHotspot20H2QpWanMetric(i, "root", make(map[string][]string), 0)

		if err != nil {
			return fmt.Errorf("error reading WirelessControllerHotspot20H2QpWanMetric: %s", err)
		}

		if o == nil {
			return fmt.Errorf("error creating WirelessControllerHotspot20H2QpWanMetric: %s", n)
		}

		return nil
	}
}

func testAccCheckWirelessControllerHotspot20H2QpWanMetricDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_wirelesscontrollerhotspot20_h2qpwanmetric" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWirelessControllerHotspot20H2QpWanMetric(i, "root", make(map[string][]string), 0)

		if err == nil {
			if o != nil {
				return fmt.Errorf("error WirelessControllerHotspot20H2QpWanMetric %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWirelessControllerHotspot20H2QpWanMetricConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_wirelesscontrollerhotspot20_h2qpwanmetric" "trname" {
  downlink_load             = 0
  downlink_speed            = 2400
  link_at_capacity          = "disable"
  link_status               = "up"
  load_measurement_duration = 0
  name                      = "%[1]s"
  symmetric_wan_link        = "symmetric"
  uplink_load               = 0
  uplink_speed              = 2400
}
`, name)
}
