resource "fortios_router_ospf" "example" {
  area {
    id = "0.0.0.0"
  }

  lifecycle { # hack to deal with API bug https://github.com/poroping/terraform-provider-fortios/issues/7
    ignore_changes = [
      "redistribute",
      "network",
      "neighbor",
      "ospf_interface",
      "summary_address",
      "distribute_list"
    ]
  }

}

resource "fortios_router_ospf_network" "example" {
  prefix = "10.0.0.0/24"
  area   = fortios_router_ospf.example.area[0].id
}