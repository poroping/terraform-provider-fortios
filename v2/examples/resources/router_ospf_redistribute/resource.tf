resource "fortios_routerospf_redistribute" "example" {
  allow_append = true

  name   = "connected"
  status = "enable"
}