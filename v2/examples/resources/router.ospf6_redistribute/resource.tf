resource "fortios_routerospf6_redistribute" "example" {
  allow_append = true

  name   = "connected"
  status = "enable"
}