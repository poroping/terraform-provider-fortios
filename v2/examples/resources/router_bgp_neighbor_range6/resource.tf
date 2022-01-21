resource "fortios_router_bgp_neighbor_group" "example" {
  name      = "foobar"
  remote_as = 65000
}

resource "fortios_router_bgp_neighbor_range6" "example" {
  prefix6        = "2003::/48"
  neighbor_group = fortios_router_bgp_neighbor_group.example.name
}