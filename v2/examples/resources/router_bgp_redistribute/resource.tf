resource "fortios_router_bgp_redistribute" "example" {
  allow_append = true

  name   = "connected"
  status = "enable"
}