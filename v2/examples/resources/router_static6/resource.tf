resource "fortios_router_static6" "example" {
  dst       = "2001:b00b:fa57::/64"
  blackhole = "enable"
}