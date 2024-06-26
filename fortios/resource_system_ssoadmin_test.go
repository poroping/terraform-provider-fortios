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

func TestAccFortiOSSystemSsoAdmin_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemSsoAdmin_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemSsoAdminConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemSsoAdminExists("fortios_system_ssoadmin.trname"),
					resource.TestCheckResourceAttr("fortios_system_ssoadmin.trname", "accprofile", "super_admin"),
					resource.TestCheckResourceAttr("fortios_system_ssoadmin.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_system_ssoadmin.trname", "vdom.0.name", "root"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemSsoAdminExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemSsoAdmin: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemSsoAdmin is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemSsoAdmin(i, "root", make(map[string][]string), 0)

		if err != nil {
			return fmt.Errorf("error reading SystemSsoAdmin: %s", err)
		}

		if o == nil {
			return fmt.Errorf("error creating SystemSsoAdmin: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemSsoAdminDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_ssoadmin" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemSsoAdmin(i, "root", make(map[string][]string), 0)

		if err == nil {
			if o != nil {
				return fmt.Errorf("error SystemSsoAdmin %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemSsoAdminConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_ssoadmin" "trname" {
  accprofile = "super_admin"
  name       = "%[1]s"

  vdom {
    name = "root"
  }
}
`, name)
}
