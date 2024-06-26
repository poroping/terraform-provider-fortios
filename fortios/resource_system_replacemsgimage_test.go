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

func TestAccFortiOSSystemReplacemsgImage_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemReplacemsgImage_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemReplacemsgImageConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemReplacemsgImageExists("fortios_system_replacemsgimage.trname"),
					resource.TestCheckResourceAttr("fortios_system_replacemsgimage.trname", "image_type", "png"),
					resource.TestCheckResourceAttr("fortios_system_replacemsgimage.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_system_replacemsgimage.trname", "image_base64", "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAAEWAAABFgAVshLGQAAAAMSURBVBhXY/j//z8ABf4C/qc1gYQAAAAASUVORK5CYII="),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemReplacemsgImageExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemReplacemsgImage: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemReplacemsgImage is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemReplacemsgImage(i, "root", make(map[string][]string), 0)

		if err != nil {
			return fmt.Errorf("error reading SystemReplacemsgImage: %s", err)
		}

		if o == nil {
			return fmt.Errorf("error creating SystemReplacemsgImage: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemReplacemsgImageDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_replacemsgimage" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemReplacemsgImage(i, "root", make(map[string][]string), 0)

		if err == nil {
			if o != nil {
				return fmt.Errorf("error SystemReplacemsgImage %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemReplacemsgImageConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_replacemsgimage" "trname" {
  image_type   = "png"
  name         = "%[1]s"
  image_base64 = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAAEWAAABFgAVshLGQAAAAMSURBVBhXY/j//z8ABf4C/qc1gYQAAAAASUVORK5CYII="
}
`, name)
}
