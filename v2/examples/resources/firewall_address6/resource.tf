resource "fortios_firewall_address6" "example" {
  name    = "foobar6"
  ip6     = "3004:1234:b00b::/64"
  comment = "acc testing"
}