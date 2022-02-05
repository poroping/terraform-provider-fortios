resource "fortios_system_interface" "example" {
  allow_append = true
  vdomparam    = "root"

  name = "TESTACCINT"
  type = "loopback"
  ip   = "169.254.72.255/32"
  vdom = "root"
}

resource "fortios_system_interface" "example2" {
  allow_append = true
  vdomparam    = "root"

  name = "TESTV6"
  type = "loopback"
  ip   = "169.254.71.255/32"
  vdom = "root"
  secondary_ip = "enable"

  ipv6 {
    ip6_address = "2001:beef::01/128"
  }

  secondaryip {
    ip = "3.3.3.3/32"
  }
}