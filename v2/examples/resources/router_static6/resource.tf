resource "fortios_router_static6" "example" {
  dst       = "2001:boob:fa57::/64"
  blackhole = "enable"
}