resource "fortios_router_static" "example" {
  dst       = "1.1.1.50/32"
  blackhole = "enable"
}