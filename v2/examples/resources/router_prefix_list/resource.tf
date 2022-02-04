resource "fortios_router_prefixlist" "example" {
  dynamic_sort_subtable = true

  name     = "pl_public_only"

  rule {
    id     = 1
    action = "deny"
    prefix = "10.0.0.0/8"
    le     = 32
  }

  rule {
    id     = 2
    action = "deny"
    prefix = "172.16.0.0/12"
    le     = 32
  }

  rule {
    id     = 3
    action = "deny"
    prefix = "192.168.0.0/16"
    le     = 32
  }

  rule {
    id     = 4
    action = "deny"
    prefix = "0.0.0.0/0"
  }

  rule {
    id     = 10
    action = "permit"
    prefix = "0.0.0.0/0"
    le     = 32
  }
}