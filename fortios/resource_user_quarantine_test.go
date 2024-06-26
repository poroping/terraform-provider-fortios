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

func TestAccFortiOSUserQuarantine_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSUserQuarantine_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSUserQuarantineConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSUserQuarantineExists("fortios_user_quarantine.trname"),
					resource.TestCheckResourceAttr("fortios_user_quarantine.trname", "quarantine", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSUserQuarantineExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found UserQuarantine: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No UserQuarantine is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadUserQuarantine(i, "root", make(map[string][]string), 0)

		if err != nil {
			return fmt.Errorf("error reading UserQuarantine: %s", err)
		}

		if o == nil {
			return fmt.Errorf("error creating UserQuarantine: %s", n)
		}

		return nil
	}
}

func testAccCheckUserQuarantineDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_user_quarantine" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadUserQuarantine(i, "root", make(map[string][]string), 0)

		if err == nil {
			if o != nil {
				return fmt.Errorf("error UserQuarantine %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSUserQuarantineConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_user_quarantine" "trname" {
  quarantine = "enable"
}
`)
}
