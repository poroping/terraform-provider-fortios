resource "fortios_router_ospf_redistribute" "example" {
  allow_append = true

  name   = "connected"
  status = "enable"
}