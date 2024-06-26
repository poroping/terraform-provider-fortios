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

func TestAccFortiOSIcapServer_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSIcapServer_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSIcapServerConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSIcapServerExists("fortios_icap_server.trname"),
					resource.TestCheckResourceAttr("fortios_icap_server.trname", "ip6_address", "::"),
					resource.TestCheckResourceAttr("fortios_icap_server.trname", "ip_address", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_icap_server.trname", "ip_version", "4"),
					resource.TestCheckResourceAttr("fortios_icap_server.trname", "max_connections", "100"),
					resource.TestCheckResourceAttr("fortios_icap_server.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_icap_server.trname", "port", "22"),
				),
			},
		},
	})
}

func testAccCheckFortiOSIcapServerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found IcapServer: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No IcapServer is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadIcapServer(i, "root", make(map[string][]string), 0)

		if err != nil {
			return fmt.Errorf("error reading IcapServer: %s", err)
		}

		if o == nil {
			return fmt.Errorf("error creating IcapServer: %s", n)
		}

		return nil
	}
}

func testAccCheckIcapServerDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_icap_server" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadIcapServer(i, "root", make(map[string][]string), 0)

		if err == nil {
			if o != nil {
				return fmt.Errorf("error IcapServer %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSIcapServerConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_icap_server" "trname" {
  ip6_address     = "::"
  ip_address      = "1.1.1.1"
  ip_version      = "4"
  max_connections = 100
  name            = "%[1]s"
  port            = 22
}
`, name)
}
