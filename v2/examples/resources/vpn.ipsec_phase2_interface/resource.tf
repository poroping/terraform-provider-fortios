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

resource "fortios_vpnipsec_phase2interface" "example" {
  name       = fortios_vpnipsec_phase1interface.example.name
  phase1name = fortios_vpnipsec_phase1interface.example.name
  proposal   = "aes256-sha256 aes256-sha512"
  dhgrp      = "5 14"
  src_subnet = "0.0.0.0/0"
  dst_subnet = "0.0.0.0/0"
}

resource "fortios_system_interface" "vpn_example" {
  allow_append = true // needed as interface is created during phase1 creation

  name      = fortios_vpnipsec_phase1interface.example.name
  ip        = "169.254.3.0/32"
  remote_ip = "169.254.3.1/31"
}
