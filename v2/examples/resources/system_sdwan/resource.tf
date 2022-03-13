resource "fortios_system_sdwan" "example" {
    status = "enable"

    // ignore defaults
    lifecycle {
        ignore_changes = [
          health_check,
          members,
          zone
        ]
    }
}