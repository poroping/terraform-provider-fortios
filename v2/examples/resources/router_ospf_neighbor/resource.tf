resource "fortios_system_interface" "example" {
  vdom = "BUTT"
  type = "loopback"
  name = "foobar123"
  ip   = "5.5.5.0/31"
}

resource "fortios_routerospf_ospfinterface" "example" {
  name         = fortios_system_interface.example.name
  interface    = fortios_system_interface.example.name
  network_type = "non-broadcast"
}

resource "fortios_routerospf_neighbor" "example" {
  ip = "5.5.5.1"

  depends_on = [
    fortios_routerospf_ospfinterface.example
  ]
}