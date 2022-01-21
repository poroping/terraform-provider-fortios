resource "fortios_router_bgp_neighbor_group" "example" {
  name      = "foobar"
  remote_as = 65000
}

resource "fortios_router_bgp_neighbor_range" "example" {
  prefix         = "192.168.1.0/24"
  neighbor_group = fortios_router_bgp_neighbor_group.example.name
}