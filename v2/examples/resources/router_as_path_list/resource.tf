resource "fortios_router_aspathlist" "example" {
  name = "aspath1"

  rule {
    action = "deny"
    regexp = "/d+/n"
  }
}