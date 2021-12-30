resource "fortios_firewall_ippool6" "example" {
  name    = "foobar"
  startip = "2001::1"
  endip   = "2001::3"
}