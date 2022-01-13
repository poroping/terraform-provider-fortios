resource "fortios_system_interface" "example" {
  allow_append = true
  vdomparam    = "root"

  name = "TESTACCINT2"
  type = "aggregate"
  vdom = "root"
}

resource "fortios_system_zone" "example" {
  vdomparam = fortios_system_interface.example.vdom

  name = "TESTZONE"
  interface {
    interface_name = fortios_system_interface.example.name
  }
}