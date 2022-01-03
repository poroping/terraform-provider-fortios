resource "fortios_system_interface" "example" {
  allow_append = true
  vdomparam = "root"

  name   = "TESTACCINT"
  type   = "loopback"
  ip     = "169.254.72.255/32"
  vdom   = "root"
}