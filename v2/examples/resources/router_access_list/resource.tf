resource "fortios_router_accesslist" "example" {
  comments = "foobar"
  name     = "acl_test"

  rule {
      id = 5
      action = "deny"
      prefix = "10.0.0.0/8"
  }
  rule {
      id = 20
      action = "permit"
      prefix = "any"
  }
}