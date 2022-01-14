resource "fortios_system_interface" "example" {
  vdom = "BUTT"
  type = "loopback"
  name = "foobar123"
  ip   = "5.5.5.0/31"
}

resource "fortios_router_ospf_ospf_interface" "example" {
  name         = fortios_system_interface.example.name
  interface    = fortios_system_interface.example.name
  network_type = "non-broadcast"
}

resource "fortios_router_ospf_neighbor" "example" {
  ip = "5.5.5.1"

  depends_on = [
    fortios_router_ospf_ospf_interface.example
  ]
}