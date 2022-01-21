# resource "fortios_router_ospf_area" "example" {
#   fosid = "0.0.0.0"
# }

### This resource may not work due to an API issue
### https://github.com/poroping/terraform-provider-fortios/issues/7
### WORKAROUND:

resource "fortios_router_ospf" "example" {
  area {
    id = "0.0.0.0"
  }

  lifecycle {
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