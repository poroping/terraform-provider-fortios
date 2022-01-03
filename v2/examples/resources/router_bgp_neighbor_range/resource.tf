resource "fortios_routerbgp_neighborgroup" "example" {
  name      = "foobar"
  remote_as = 65000
}

resource "fortios_routerbgp_neighborrange" "example" {
  prefix         = "192.168.1.0/24"
  neighbor_group = fortios_routerbgp_neighborgroup.example.name
}