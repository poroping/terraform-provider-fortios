resource "fortios_system_interface" "example" {
  vdom = "BUTT"
  type = "loopback"
  name = "foobar234"
}

resource "fortios_router_ospf_ospf_interface" "example" {
  name      = fortios_system_interface.example.name
  interface = fortios_system_interface.example.name
}