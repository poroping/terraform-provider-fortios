resource "fortios_router_accesslist6" "example" {
  comments = "foobar"
  name     = "acl_test"

  rule {
      id = 5
      action = "permit"
      prefix6 = "any"
  }

  rule {
      id = 3
      action = "deny"
      prefix6 = "2001::/64"
  }
}