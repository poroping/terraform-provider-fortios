resource "fortios_firewall_policy_sort" "example" {
  sortby       = "name"
  sortdirection = "descending"
  auto_sort = true
}