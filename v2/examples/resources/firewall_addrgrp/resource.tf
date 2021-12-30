resource "fortios_firewall_address" "example" {
  name    = "foobar"
  type    = "ipmask"
  subnet  = "10.0.1.0/24"
  color   = 6
  comment = "acc testing"
}

resource "fortios_firewall_addrgrp" "example" {
  name = "groupygroup"
  member {
    name = fortios_firewall_address.example.name
  }
}