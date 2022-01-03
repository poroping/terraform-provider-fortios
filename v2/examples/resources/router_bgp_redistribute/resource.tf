resource "fortios_routerbgp_redistribute" "example" {
  allow_append = true

  name   = "connected"
  status = "enable"
}