resource "fortios_routerbgp_neighborgroup" "example" {
  name      = "foobar"
  remote_as = 65000
}

resource "fortios_routerbgp_neighborrange6" "example" {
  prefix6        = "2003::/48"
  neighbor_group = fortios_routerbgp_neighborgroup.example.name
}