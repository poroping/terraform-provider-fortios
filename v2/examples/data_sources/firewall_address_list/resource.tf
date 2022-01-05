data "fortios_firewall_addresslist" example {
    vdomparam = "root"
    
  filter = "name!=fakeyjakey"
}

output example {
  value = data.fortios_firewall_addresslist.example.namelist
}