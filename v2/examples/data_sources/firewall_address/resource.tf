data "fortios_firewall_address" example {
    vdomparam = "root"
    
  name = "all"
}

output example {
  value = data.fortios_firewall_address.example
}