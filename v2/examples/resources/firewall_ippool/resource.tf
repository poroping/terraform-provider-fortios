resource "fortios_firewall_ippool" "example" {
  name    = "foobar"
  startip = "1.1.1.1"
  endip   = "1.1.1.2"
}