resource "fortios_routerbgp_neighbor" "example" {
  ip        = "1.1.1.1"
  remote_as = 65000
}

resource "fortios_routerbgp_neighbor" "example6" {
  ip        = "2001::c"
  remote_as = 65000
}