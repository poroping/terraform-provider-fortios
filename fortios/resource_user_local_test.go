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

func TestAccFortiOSUserLocal_basic(t *testing.T) {
	rname := acctest.RandString(8)
	var0 := "var0" + rname
	log.Printf(var0)
	log.Printf("TestAccFortiOSUserLocal_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSUserLocalConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSUserLocalExists("fortios_user_local.trname"),
					resource.TestCheckResourceAttr("fortios_user_local.trname", "auth_concurrent_override", "disable"),
					resource.TestCheckResourceAttr("fortios_user_local.trname", "auth_concurrent_value", "0"),
					resource.TestCheckResourceAttr("fortios_user_local.trname", "authtimeout", "0"),
					resource.TestCheckResourceAttr("fortios_user_local.trname", "ldap_server", var0),
					resource.TestCheckResourceAttr("fortios_user_local.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_user_local.trname", "passwd_time", "0000-00-00 00:00:00"),
					resource.TestCheckResourceAttr("fortios_user_local.trname", "sms_server", "fortiguard"),
					resource.TestCheckResourceAttr("fortios_user_local.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_user_local.trname", "two_factor", "disable"),
					resource.TestCheckResourceAttr("fortios_user_local.trname", "type", "ldap"),
				),
			},
		},
	})
}

func testAccCheckFortiOSUserLocalExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found UserLocal: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No UserLocal is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadUserLocal(i, "root", make(map[string][]string), 0)

		if err != nil {
			return fmt.Errorf("error reading UserLocal: %s", err)
		}

		if o == nil {
			return fmt.Errorf("error creating UserLocal: %s", n)
		}

		return nil
	}
}

func testAccCheckUserLocalDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_user_local" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadUserLocal(i, "root", make(map[string][]string), 0)

		if err == nil {
			if o != nil {
				return fmt.Errorf("error UserLocal %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSUserLocalConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_user_ldap" "trname3" {
  account_key_filter      = "(&(userPrincipalName=%s)(!(UserAccountControl:1.2.840.113556.1.4.803:=2)))"
  account_key_processing  = "same"
  cnid                    = "cn"
  dn                      = "EIWNCIEW"
  group_member_check      = "user-attr"
  group_object_filter     = "(&(objectcategory=group)(member=*))"
  member_attr             = "memberOf"
  name                    = "var0%[1]s"
  password_expiry_warning = "disable"
  password_renewal        = "disable"
  port                    = 389
  secure                  = "disable"
  server                  = "1.1.1.1"
  server_identity_check   = "disable"
  source_ip               = "0.0.0.0"
  ssl_min_proto_version   = "default"
  type                    = "simple"
}

resource "fortios_user_local" "trname" {
  auth_concurrent_override = "disable"
  auth_concurrent_value    = 0
  authtimeout              = 0
  ldap_server              = fortios_user_ldap.trname3.name
  name                     = "%[1]s"
  passwd_time              = "0000-00-00 00:00:00"
  sms_server               = "fortiguard"
  status                   = "enable"
  two_factor               = "disable"
  type                     = "ldap"
}
`, name)
}
