resource "fortios_system_interface" "example" {
  vdom = "BUTT"
  type = "loopback"
  name = "foobar234"
}

resource "fortios_routerospf_ospfinterface" "example" {
  name      = fortios_system_interface.example.name
  interface = fortios_system_interface.example.name
}