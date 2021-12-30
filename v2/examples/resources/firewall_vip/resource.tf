resource "fortios_firewall_vip" "example" {
  name    = "foobar"
  extip = "1.1.1.1"
  extintf = "any"

  mappedip {
    range = "192.168.1.1-192.168.1.1"
  }
}