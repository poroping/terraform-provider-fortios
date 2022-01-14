resource "fortios_router_ospf6_redistribute" "example" {
  allow_append = true

  name   = "connected"
  status = "enable"
}