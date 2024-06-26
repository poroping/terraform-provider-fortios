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

func TestAccFortiOSSystemGeneve_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemGeneve_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemGeneveConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemGeneveExists("fortios_system_geneve.trname"),
					resource.TestCheckResourceAttr("fortios_system_geneve.trname", "dstport", "22"),
					resource.TestCheckResourceAttr("fortios_system_geneve.trname", "interface", "port2"),
					resource.TestCheckResourceAttr("fortios_system_geneve.trname", "ip_version", "ipv4-unicast"),
					resource.TestCheckResourceAttr("fortios_system_geneve.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_system_geneve.trname", "remote_ip", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_system_geneve.trname", "remote_ip6", "::"),
					resource.TestCheckResourceAttr("fortios_system_geneve.trname", "vni", "0"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemGeneveExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemGeneve: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemGeneve is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemGeneve(i, "root", make(map[string][]string), 0)

		if err != nil {
			return fmt.Errorf("error reading SystemGeneve: %s", err)
		}

		if o == nil {
			return fmt.Errorf("error creating SystemGeneve: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemGeneveDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_geneve" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemGeneve(i, "root", make(map[string][]string), 0)

		if err == nil {
			if o != nil {
				return fmt.Errorf("error SystemGeneve %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemGeneveConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_geneve" "trname" {
  dstport    = 22
  interface  = "port2"
  ip_version = "ipv4-unicast"
  name       = "%[1]s"
  remote_ip  = "1.1.1.1"
  remote_ip6 = "::"
  vni        = 0
}
`, name)
}
