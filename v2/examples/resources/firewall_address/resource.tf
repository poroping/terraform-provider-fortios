resource "fortios_firewall_address" "example" {
  name    = "foobar"
  type    = "ipmask"
  subnet  = "10.0.1.0/24"
  color   = 6
  comment = "acc testing"
}

resource "fortios_firewall_address" "example2" {
  name    = "fqdn"
  type    = "fqdn"
  fqdn    = "example.com"
  color   = 3
  comment = "acc testing"
}

resource "fortios_firewall_address" "example3" {
  name     = "iprange"
  type     = "iprange"
  start_ip = "192.168.1.4"
  end_ip   = "192.168.1.76"
  color    = 32
  comment  = "acc testing"
}

resource "fortios_firewall_address" "example4" {
  name    = "country"
  type    = "geography"
  country = "IS"
  color   = 15
  comment = "acc testing"
}

resource "fortios_firewall_address" "specialchars" {
  name    = "BLA**^^$$@@&&++-10.1.1.1/32"
  type    = "ipmask"
  subnet  = "10.1.1.1/32"
  color   = 6
  comment = "acc testing"
}