resource "fortios_system_sdwan" "example" {
  status = "enable"

  // ignore defaults
  lifecycle {
    ignore_changes = [
      service,
      health_check,
      members,
      zone
    ]
  }
}

resource "fortios_system_sdwan_health_check" "example" {
  name   = "example"
  server = "192.0.2.0"
  
  // members = 0 is a default
  members {
    seq_num = 0
  }
}
