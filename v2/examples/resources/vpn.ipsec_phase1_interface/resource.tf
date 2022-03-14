resource "fortios_system_interface" "example" {
  name = "TESTACCINT"
  type = "loopback"
  vdom = "BUTT"
}

resource "fortios_vpnipsec_phase1interface" "example" {
  name      = "vpn-example"
  interface = fortios_system_interface.example.name
  psksecret = "supersecret"
  remote_gw = "192.0.2.1"
  proposal  = "aes256-sha256 aes256-sha512"
}
