resource "fortios_firewall_address6" "example" {
  name    = "foobar60"
  ip6     = "3004:1234:b00b::/64"
  comment = "acc testing"
}

resource "fortios_firewall_addrgrp6" "example" {
  name = "groupygroup"
  member {
    name = fortios_firewall_address6.example.name
  }
}